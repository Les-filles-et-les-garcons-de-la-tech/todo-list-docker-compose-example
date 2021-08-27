import { Component } from '@angular/core';
import { Todo } from './models/app.models.todo';
import { TodoService } from './todo.service';
@Component({
  selector: 'app-root',
  templateUrl: './app.component.html',
  styleUrls: ['./app.component.scss']
})
export class AppComponent {
  public todos: Todo[] = [];
  todoName: string = "";
  constructor(private todoService: TodoService ) { }

  ngOnInit() {
    this.todoService.todosSubject.subscribe(todos => {
      this.todos = todos
    })
    this.loadAllTodoList();
  }

  loadAllTodoList() {
    this.todoService.getAllTodos();
  }
  
  onClickTodoDelete(id="d") {
    this.todoService.deleteTodo(id);
  }

  onClickTodoAdd(){

    this.todoService.addTodo({
      name: this.todoName,
      done: false
    });
  }
}
