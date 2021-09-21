import { HttpClient } from '@angular/common/http';
import { Injectable } from '@angular/core';
import { Subject } from 'rxjs';
import { Todo } from './models/app.models.todo';
import { environment } from 'src/environments/environment';

@Injectable({
  providedIn: 'root'
})
export class TodoService {
  constructor(private http: HttpClient) { }
  private URL : string = environment.apiUrl+"/api/todo"
  public todosSubject : Subject<Todo[]> = new Subject();
  
  getAllTodos(): void {
    console.log(environment.apiUrl);
    this.http.get<Todo[]>(this.URL).subscribe(data => {
      var todos : Todo[] = [];
      data.forEach(todo => {
        todos.push(todo)
        
      });
      this.todosSubject.next(todos);
    })
  }

  addTodo(todo: Todo) {
    this.http.post<any>(this.URL, todo).subscribe({
        next: data => {
            console.log("____________ADD____________");
            console.log(data);
            this.getAllTodos();
        },
        error: error => {
            console.error('There was an error!', error);
        }
    })
  }
  deleteTodo(id: any) {
    this.http.delete(this.URL+"/"+id)
        .subscribe({
            next: data => {
                console.log('Delete successful');
                this.getAllTodos();
            },
            error: error => {
                console.error('There was an error!', error);
            }
        });
  }


}
