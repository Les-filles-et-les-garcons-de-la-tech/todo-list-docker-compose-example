using WebApp.Models;
using MongoDB.Driver;
using System.Collections.Generic;
using System.Linq;
using System;

namespace WebApp.Services
{
    public class TodoService
    {
        private readonly IMongoCollection<TodoItem> _todoItems;

        public TodoService(ITodolistDatabaseSettings settings)
        {
            MongoClient client;
            if(Environment.GetEnvironmentVariable("DB_ADDRESS") != null && 
                Environment.GetEnvironmentVariable("DB_PORT") != null && 
                Environment.GetEnvironmentVariable("DB_USERNAME") != null && 
                Environment.GetEnvironmentVariable("DB_PASSWORD") != null)
            {
                client = new MongoClient("mongodb://"+
                                Environment.GetEnvironmentVariable("DB_USERNAME") +
                                ":" +
                                Environment.GetEnvironmentVariable("DB_PASSWORD") +
                                "@" + 
                                Environment.GetEnvironmentVariable("DB_ADDRESS") + 
                                ":" +
                                Environment.GetEnvironmentVariable("DB_PORT") +
                                "/");
            }
            else
            {
                client = new MongoClient("mongodb://root:example@mongo:27017/");
            }
            
            var database = client.GetDatabase(settings.DatabaseName);

            _todoItems = database.GetCollection<TodoItem>(settings.TodoCollectionName);
        }

        public List<TodoItem> Get() =>
            _todoItems.Find(todoItem => true).ToList();

        public TodoItem Get(string id) =>
            _todoItems.Find<TodoItem>(todoItem => todoItem.Id == id).FirstOrDefault();

        public TodoItem Create(TodoItem todoItem)
        {
            _todoItems.InsertOne(todoItem);
            return todoItem;
        }

        public void Update(string id, TodoItem todoItemIn) =>
            _todoItems.ReplaceOne(todoItem => todoItem.Id == id, todoItemIn);

        public void Remove(TodoItem todoItemIn) =>
            _todoItems.DeleteOne(todoItem => todoItem.Id == todoItemIn.Id);

        public void Remove(string id) =>
            _todoItems.DeleteOne(todoItem => todoItem.Id == id);
    }
}