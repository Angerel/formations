using Microsoft.AspNetCore.Mvc;
using Microsoft.EntityFrameworkCore;
using API.Models;
using API.Repositories;
using API.Services;
using Humanizer;

namespace TodoApi.Controllers;

[Route("todos")]
[ApiController]
public class TodoController : ControllerBase
{
    private readonly TodoService _service;

    public TodoController(TodoRepository context)
    {
        _service = new TodoService(context);
    }

    // GET: api/todos
    [HttpGet]
    [Route("")]
    public ActionResult<IEnumerable<TodoListDTO>> GetTodoLists() {
        return _service.GetAllTodos();
    }

    // GET: api/todos/1
    [HttpGet]
    [Route("{listId}")]
    public ActionResult<TodoListDTO> GetTodoList(int listId) {
        return _service.GetTodoById(listId) is TodoListDTO todoList ? Ok(todoList) : NotFound();
    }

    // GET: api/todos/1/2
    [HttpGet]
    [Route("{listId}/{itemId}")]
    public ActionResult<TodoItemDTO> GetTodoItem(int listId, int itemId) {
        return _service.GetTodoItemById(listId, itemId) is TodoItemDTO todoItem ? Ok(todoItem) : NotFound();
    }

    // POST: api/todos
    [HttpPost]
    [Route("")]
    public ActionResult<TodoListDTO> PostTodoList(TodoListDTO todoList)
    {
        TodoList created = _service.AddTodoList(todoList);
        return CreatedAtAction(nameof(GetTodoList), new { listId = created.Id }, (TodoListDTO) created);
    }

    // POST: api/todos/1
    [HttpPost]
    [Route("{listId}")]
    public ActionResult<TodoListDTO> PostTodoItem(int listId, TodoItemDTO todoItem)
    {
        try {
            _service.AddTodoItem(listId, todoItem);
            TodoListDTO? todoList = _service.GetTodoById(listId);
            return CreatedAtAction(nameof(GetTodoList), new { listId }, todoList);
        } catch (Exception) {
            return NotFound("The requested resource was not found.");
        }
    }

    // PUT: api/todos/1
    [HttpPut]
    [Route("{listId}")]
    public ActionResult<TodoListDTO> PutTodoList(int listId, TodoListDTO todoList) {
        try {
            _service.UpdateTodoList(listId, todoList);
            return NoContent();
        } catch (Exception e) {
            return NotFound(e.Message.Humanize());
        }
    }

    // PUT: api/todos/1/2
    [HttpPut]
    [Route("{listId}/{itemId}")]
    public ActionResult<TodoItemDTO> PutTodoItem(int listId, int itemId, TodoItemDTO todoItem) {
        try {
            _service.UpdateTodoItem(listId, itemId, todoItem);
            return NoContent();
        } catch (Exception) {
            return NotFound("The requested resource was not found.");
        }
    }

    // DELETE: api/todos/1
    [HttpDelete]
    [Route("{listId}")]
    public ActionResult<TodoListDTO> DeleteTodoList(int listId) {
        try {
            _service.RemoveTodoList(listId);
            return NoContent();
        } catch (Exception) {
            return NotFound("The requested resource was not found.");
        }
    }

    // DELETE: api/todos/1/2
    [HttpDelete]
    [Route("{listId}/{itemId}")]
    public ActionResult<TodoItemDTO> DeleteTodoItem(int listId, int itemId) {
        try {
            _service.RemoveTodoItem(listId, itemId);
            return NoContent();
        } catch (Exception) {
            return NotFound("The requested resource was not found.");
        }
    }
}
