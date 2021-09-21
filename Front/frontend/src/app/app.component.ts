import { Component } from '@angular/core';
import { ColorService } from './color.service';
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
  color: string = "#EEEEEE";
  constructor(private todoService: TodoService, private colorService: ColorService ) { }

  ngOnInit() {
    this.todoService.todosSubject.subscribe(todos => {
      this.todos = todos
    })
    this.colorService.colorSubject.subscribe(color => {
      window.document.body.style.backgroundColor = color;

    })
    this.loadAllTodoList();
    this.getColor();
  }

  loadAllTodoList() {
    this.todoService.getAllTodos();
  }

  getColor() {
    this.colorService.getColor();
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
