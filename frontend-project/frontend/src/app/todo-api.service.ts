import { Injectable } from '@angular/core';
import { HttpClient } from '@angular/common/http';
import { Observable } from 'rxjs';
import { environment } from 'src/environments/environment';
import { ColorModel, TodoItem } from './models';

@Injectable({
  providedIn: 'root'
})
export class TodoApiService {

  private readonly baseUrl = environment.apiBaseUrl;

  constructor(private http: HttpClient) {}

  getBackgroundColor(): Observable<ColorModel> {
    return this.http.get<ColorModel>(`${this.baseUrl}/api/color`);
  }

  getTodos(): Observable<TodoItem[]> {
    return this.http.get<TodoItem[]>(`${this.baseUrl}/api/todo`);
  }

  createTodo(name: string): Observable<TodoItem> {
    const body = { name, done: false };
    return this.http.post<TodoItem>(`${this.baseUrl}/api/todo`, body);
  }

  deleteTodo(id: string): Observable<void> {
    return this.http.delete<void>(`${this.baseUrl}/api/todo/${id}`);
  }
}
