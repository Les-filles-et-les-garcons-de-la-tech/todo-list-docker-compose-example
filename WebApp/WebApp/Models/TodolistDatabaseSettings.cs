using System;
namespace WebApp.Models
{
    public class TodolistDatabaseSettings : ITodolistDatabaseSettings
    {
        public string TodoCollectionName { get; set; }
        public string ConnectionString { get; set; }
        public string DatabaseName { get; set; }
    }

    public interface ITodolistDatabaseSettings
    {
        string TodoCollectionName { get; set; }
        string ConnectionString { get; set; }
        string DatabaseName { get; set; }
    }
}
