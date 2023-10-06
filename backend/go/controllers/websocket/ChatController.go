package WS

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"github.com/gorilla/websocket"
	"go-api/models"
	"go-api/services"
	"net/http"
	"strconv"
	"time"
)

type ChatControllerStruct struct {
	NextId   int                `json:"nextId"`
	Upgrader websocket.Upgrader `json:"upgrader"`
}

func HandleConnectionClosure(code int, text string, conn *websocket.Conn) error {
	usr, err := services.UserService.FindByConnection(conn)
	if err != nil {
		return err
	}
	_, err = services.UserService.UnLinkConnection(usr.User.Id)
	if err != nil {
		return err
	}
	return nil
}

/* JoinChat - Live chat websocket entrypoint
 * Requests all live users and messages, then upgrades websocket connection
 * Reads the first message received as the user name
 * Writes the first message as a message history along with the user's identifier
 * Then registers every message received in the database and broadcasts them to all the listeners
 *
 * @uses MessageService
 * @uses UserService
 */
func (instance ChatControllerStruct) JoinChat(c *gin.Context) {
	/* Variables initialisation */
	//error variable
	var err error

	//websocket variables
	var w = c.Writer
	var r = c.Request

	//User list
	users, err := services.UserService.FindAll()
	if err != nil {
		fmt.Printf("Failed to get users : %+v\n", err)
		return
	}

	//Message list
	messages, err := services.MessageService.FindAll()
	if err != nil {
		fmt.Printf("Failed to get message history : %+v\n", err)
		return
	}

	/* Upgrades connection from REST to WebSocket */
	conn, err := instance.Upgrader.Upgrade(w, r, nil)
	if err != nil {
		fmt.Printf("Failed to set websocket upgrade: %+v\n", err)
		return
	}

	/* Sets how a closed connection should be handled */
	conn.SetCloseHandler(func(code int, text string) error {
		return HandleConnectionClosure(code, text, conn)
	})

	/* First message policy : Reads the user's name first */
	_, msg, err := conn.ReadMessage()
	if err != nil {
		return
	}
	var username = string(msg)

	/* Finds the user in the database
	 * If it exists, links the connection object to the user
	 * Else, creates a new user with the connection object attached
	 */
	currentUser, err := services.UserService.FindByName(username)
	if currentUser == nil {
		currentUser = &models.ConnectedUser{
			User: models.User{
				Id:   strconv.Itoa(instance.NextId),
				Name: username,
			},
			Connection: conn,
		}
		instance.NextId++
		*users = append(*users, currentUser)
	} else {
		currentUser.Connection = conn
	}

	/* Then, sends the message history to the user, along with its identifier for its own recognition */
	conn.WriteJSON(models.FirstMessage{
		History: *messages,
		UserId:  currentUser.User.Id,
	})

	/* Finally, starts a loop of receiving messages and broadcasting them to all the listeners */
	for {
		//Receives message content
		_, msg, err := conn.ReadMessage()
		if err != nil {
			break
		}

		//Formats it into a complete object with id, author and timestamp
		var newMsg = models.Message{
			Id:        strconv.Itoa(len(*messages)),
			Message:   string(msg),
			Timestamp: strconv.FormatInt(time.Now().Unix(), 10),
			User:      currentUser.User,
		}

		//Pushes the new message into the database
		services.MessageService.CreateOne(&newMsg)

		//Broadcasts it to all live connections
		broadcastJson(newMsg, *users)
	}
}

/* Sends any content to a user if it is listening to the websocket */
func sendJsonTo(message interface{}, user *models.ConnectedUser) {
	if user.Connection != nil {
		user.Connection.WriteJSON(message)
	}
}

/* Sends any content to a list of users */
func broadcastJson(message interface{}, receivers []*models.ConnectedUser) {
	for _, user := range receivers {
		sendJsonTo(message, user)
	}
}

var ChatController = ChatControllerStruct{
	Upgrader: websocket.Upgrader{
		ReadBufferSize:  1024,
		WriteBufferSize: 1024,
		CheckOrigin: func(r *http.Request) bool {
			return true
		},
	},
	NextId: 0,
}
