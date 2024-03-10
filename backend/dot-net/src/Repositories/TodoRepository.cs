namespace API.Repositories;

using API.Models;
using Microsoft.EntityFrameworkCore;
    
public class TodoRepository : DbContext
{
    public TodoRepository(DbContextOptions<TodoRepository> options) : base(options) { }

    public DbSet<TodoList> Todos { get; set; } = null!;
}