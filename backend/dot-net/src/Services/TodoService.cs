using API.Models;
using API.Repositories;
using Humanizer;
using Microsoft.EntityFrameworkCore;

namespace API.Services;
public class TodoService {
    private readonly TodoRepository _context;

    public TodoService(TodoRepository context) {
        _context = context;
    }

    public List<TodoListDTO> GetAllTodos() {
        return _context.Todos.Include(tl => tl.List).Select(tl => (TodoListDTO) tl).ToList();
    }

    public TodoListDTO? GetTodoById(int listId) {
        TodoList? todoList = _context.Todos.Include(tl => tl.List).FirstOrDefault(tl => tl.Id == listId);
        if(todoList == null) return null;
        return (TodoListDTO) todoList;
    }

    public TodoItemDTO? GetTodoItemById(int listId, int itemId) {
        TodoList? todoList = _context.Todos.Include(tl => tl.List).FirstOrDefault(tl => tl.Id == listId);
        if(todoList == null) return null;
        TodoItem? todoItem = todoList.List.Find(item => item.Id == itemId);
        if(todoItem == null) return null;
        return todoItem;
    }

    public TodoList AddTodoList(TodoListDTO todoList) {
        var added = _context.Todos.Add((TodoList) todoList);
        _context.SaveChanges();
        return added.Entity;
    }

    public TodoItem AddTodoItem(int listId, TodoItemDTO todoItem) {
        TodoList? todoList = _context.Todos.Include(tl => tl.List).FirstOrDefault(tl => tl.Id == listId);
        if(todoList == null) throw new Exception("List not found");
        todoList.List.Add((TodoItem) todoItem);
        _context.SaveChanges();
        return todoList.List.ElementAt(todoList.List.Count - 1);
    }
    public void UpdateTodoList(int listId, TodoListDTO todoList) {
        TodoList? existingTodoList = _context.Todos.Include(tl => tl.List).FirstOrDefault(tl => tl.Id == listId);
        if(existingTodoList == null) throw new Exception("List not found");
       
        existingTodoList.List.Clear();
        existingTodoList.List.AddRange(todoList.List.Select(item => (TodoItem) item).ToList());
        _context.SaveChanges();
    }

    public void UpdateTodoItem(int listId, int itemId, TodoItemDTO todoItem) {
        TodoList? todoList = _context.Todos.Include(tl => tl.List).FirstOrDefault(tl => tl.Id == listId);
        if(todoList == null || todoList.Id != listId) throw new Exception("List not found");

        TodoItem? existingTodoItem = todoList.List.Find(item => item.Id == itemId);
        if(existingTodoItem == null || existingTodoItem.Id != itemId) throw new Exception("Item not found");

        existingTodoItem.Label = todoItem.Label;
        _context.SaveChanges();
        
    }

    public void RemoveTodoList(int listId) {
        TodoList? todoList = _context.Todos.FirstOrDefault(tl => tl.Id == listId);
        if(todoList == null) throw new Exception("List not found");
        _context.Todos.Remove(todoList);
        _context.SaveChanges();
    }

    public void RemoveTodoItem(int listId, int itemId) {
        TodoList? todoList = _context.Todos.Include(tl => tl.List).FirstOrDefault(tl => tl.Id == listId);
        if(todoList == null) throw new Exception("List not found");

        TodoItem? todoItem = todoList.List.Find(item => item.Id == itemId);
        if(todoItem == null) throw new Exception("Item not found");

        todoList.List.Remove(todoItem);
        _context.SaveChanges();
    }
}