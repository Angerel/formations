namespace API.Models;

public class TodoItem {
    public int Id { get; set; }
    public string? Label { get; set; }
    public string? Secret { get; set; }

    public static implicit operator TodoItem(TodoItemDTO todoItem) {
        return new TodoItem { Id = todoItem.Id, Label = todoItem.Label };
    }
}

public class TodoItemDTO {
    public int Id { get; set; }
    public string? Label { get; set; }

    public static implicit operator TodoItemDTO(TodoItem todoItem) {
        return new TodoItemDTO { Id = todoItem.Id, Label = todoItem.Label };
    }
}