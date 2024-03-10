namespace API.Models;

public class TodoList {
    public int Id { get; set; }
    public required List<TodoItem> List { get; set; }

    public static implicit operator TodoList(TodoListDTO todoList) {
        return new TodoList {
            Id = todoList.Id,
            List = todoList.List.Select(item => (TodoItem) item).ToList()
        };
    }
}

public class TodoListDTO {
    public int Id { get; set; }
    public required List<TodoItemDTO> List { get; set; }

    public static implicit operator TodoListDTO(TodoList todoList) {
        return new TodoListDTO {
            Id = todoList.Id,
            List = todoList.List.Select(item => (TodoItemDTO) item).ToList()
        };
    }
}