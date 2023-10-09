package WS

import (
	"encoding/json"
	"fmt"
	"github.com/gin-gonic/gin"
	"github.com/gorilla/websocket"
	"go-api/errors"
	"go-api/models"
	"go-api/services"
	"net/http"
	"strconv"
	"time"
)

type ChatControllerStruct struct {
	NextUserId int                `json:"nextUserId"`
	Upgrader   websocket.Upgrader `json:"upgrader"`
}

func HandleConnectionClosure(conn *websocket.Conn) error {
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

// JoinChat
/* Live chat websocket entrypoint
 * Requests all live users and messages, then upgrades websocket connection
 * Reads the first message received as the username
 * Writes the first message as a message history along with the user's identifier
 * Then registers every message received in the database and broadcasts them to all the listeners
 *
 * @uses MessageService
 * @uses UserService
 */
func (instance ChatControllerStruct) JoinChat(c *gin.Context) {

	/* Upgrades connection from REST to WebSocket */
	conn, err := instance.Upgrader.Upgrade(c.Writer, c.Request, nil)
	if err != nil {
		fmt.Printf("Failed to set websocket upgrade: %+v\n", err)
		return
	}

	/* Sets how a closed connection should be handled */
	conn.SetCloseHandler(func(code int, text string) error {
		return HandleConnectionClosure(conn)
	})

	/* Finally, starts a loop of receiving orders and treating them */
	for {
		//Receives message
		_, msg, err := conn.ReadMessage()
		if err != nil {
			break
		}

		//Decode message content and attributes current connection as sender
		var decoded models.ReceivedMessage
		_ = json.Unmarshal(msg, &decoded)
		decoded.Sender = conn

		//Handles the message
		err = handleNewMessage(decoded)
		if err != nil {
			break
		}
	}
	_ = conn.Close()
}

func handleNewMessage(msg models.ReceivedMessage) *errors.ErrorInterface {
	switch msg.Action {
	case "getHistory":
		return getHistory(msg.Sender)
	case "postMessage":
		return postMessage(msg.Sender, msg.Options)
	case "authenticate":
		return authenticate(msg.Sender, msg.Options)
	default:
		return errors.BadRequestException("Bad request")
	}
}

func getHistory(sender *websocket.Conn) *errors.ErrorInterface {
	messages, err := services.MessageService.FindAll()
	if err != nil {
		return err
	}
	messagesJson, _ := json.Marshal(messages)
	messagesJsonStr := string(messagesJson)

	user, err := services.UserService.FindByConnection(sender)
	if err != nil {
		return err
	}

	sendJsonTo(messagesJsonStr, user)
	return nil
}

func postMessage(sender *websocket.Conn, options string) *errors.ErrorInterface {

	/* Decode options as a message */
	var decodedOptions struct {
		Message string `json:"message"`
	}
	_ = json.Unmarshal([]byte(options), &decodedOptions)

	/* Get sender, or anonymous if sender is not registered */
	user, _ := services.UserService.FindByConnection(sender)
	if user == nil {
		user, _ = services.UserService.FindByName("Anonymous")
		if user == nil {
			return errors.InternalException("Something wrong happened")
		}
	}

	/* Get current history to find next id */
	messages, err := services.MessageService.FindAll()
	if err != nil {
		return err
	}

	/* Create new Message object with message content, next identifier, current time and sender */
	var newMsg = models.Message{
		Id:        strconv.Itoa(len(*messages)),
		Message:   decodedOptions.Message,
		Timestamp: strconv.FormatInt(time.Now().Unix(), 10),
		User:      &user.User,
	}

	/* Save the message in DB and return it */
	created, err := services.MessageService.CreateOne(&newMsg)
	if err != nil {
		return err
	}

	users, err := services.UserService.FindAll()
	if err != nil {
		return err
	}

	broadcastJson(created, *users)
	return nil
}

func authenticate(sender *websocket.Conn, options string) *errors.ErrorInterface {
	/* Decode options as username */
	var decodedOptions struct {
		Name string `json:"name"`
	}
	_ = json.Unmarshal([]byte(options), &decodedOptions)

	/* Find or create the user with the decoded username */
	currentUser, _ := services.UserService.FindByName(decodedOptions.Name)
	if currentUser == nil {
		currentUser = &models.ConnectedUser{
			User: models.User{
				Id:   strconv.Itoa(ChatController.NextUserId),
				Name: decodedOptions.Name,
			},
			Connection: sender,
		}
		ChatController.NextUserId++
		if _, err := services.UserService.CreateOne(currentUser); err != nil {
			return err
		}
	} else {
		if _, err := services.UserService.LinkConnection(currentUser.User.Id, sender); err != nil {
			return err
		}
	}

	/* Return nothing as it went well */
	return nil
}

/* Sends any content to a user if it is listening to the websocket */
func sendJsonTo(message any, user *models.ConnectedUser) {
	if user.Connection != nil {
		_ = user.Connection.WriteJSON(message)
	}
}

/* Sends any content to a list of users */
func broadcastJson(message any, receivers []*models.ConnectedUser) {
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
	NextUserId: 0,
}
