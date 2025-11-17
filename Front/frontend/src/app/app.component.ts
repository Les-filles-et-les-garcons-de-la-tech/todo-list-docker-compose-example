import { Component, OnInit, OnDestroy, ChangeDetectorRef } from '@angular/core';
import { CommonModule } from '@angular/common';
import { FormsModule } from '@angular/forms';
import { CdkDragDrop, DragDropModule, moveItemInArray } from '@angular/cdk/drag-drop';

import { TodoApiService } from './todo-api.service';
import { ColorModel, TodoItem } from './models';

@Component({
  selector: 'app-root',
  standalone: true,
  imports: [
    CommonModule,   // *ngIf, *ngFor, [ngStyle]
    FormsModule,    // [(ngModel)]
    DragDropModule  // cdkDropList, cdkDrag
  ],
  templateUrl: './app.component.html',
  styleUrls: ['./app.component.css']
})
export class AppComponent implements OnInit, OnDestroy {
  title = 'Todo List';

  backgroundColor: string = '#ffffff';
  todos: TodoItem[] = [];
  newTodoName: string = '';
  loading: boolean = false;
  errorMessage: string | null = null;

  private reloadIntervalId: any;

  constructor(
    private api: TodoApiService,
    private cdr: ChangeDetectorRef
  ) {}

  ngOnInit(): void {
    // Chargement initial
    this.loadBackgroundColor();
    this.loadTodos();

    // Reload régulier toutes les 5 secondes
    this.reloadIntervalId = setInterval(() => {
      this.loadBackgroundColor();
      this.loadTodos();
    }, 30000);
  }

  ngOnDestroy(): void {
    if (this.reloadIntervalId) {
      clearInterval(this.reloadIntervalId);
    }
  }

  loadBackgroundColor(): void {
    this.api.getBackgroundColor().subscribe({
      next: (res: ColorModel) => {
        this.backgroundColor = res?.color || '#ffffff';
        this.cdr.markForCheck();
      },
      error: () => {
        this.backgroundColor = '#ffffff';
        this.cdr.markForCheck();
      }
    });
  }

  loadTodos(): void {
    this.loading = true;
    this.errorMessage = null;
    this.cdr.markForCheck();

    this.api.getTodos().subscribe({
      next: (items) => {
        this.todos = items ?? [];
        this.loading = false;
        this.cdr.markForCheck();
      },
      error: () => {
        this.loading = false;
        this.errorMessage = 'Failed to load todos.';
        this.todos = [];
        this.cdr.markForCheck();
      }
    });
  }

  createTodo(): void {
    const name = this.newTodoName.trim();
    if (!name) {
      return;
    }

    this.api.createTodo(name).subscribe({
      next: () => {
        this.newTodoName = '';
        // On recharge la liste depuis le backend après création
        this.loadTodos();
      },
      error: () => {
        this.errorMessage = 'Failed to create todo.';
        this.cdr.markForCheck();
      }
    });
  }

  deleteTodo(todo: TodoItem): void {
    this.api.deleteTodo(todo.id).subscribe({
      next: () => {
        // On recharge la liste depuis le backend après suppression
        this.loadTodos();
      },
      error: () => {
        this.errorMessage = 'Failed to delete todo.';
        this.cdr.markForCheck();
      }
    });
  }

  drop(event: CdkDragDrop<TodoItem[]>): void {
    if (event.previousIndex === event.currentIndex) {
      return;
    }

    moveItemInArray(this.todos, event.previousIndex, event.currentIndex);
    this.cdr.markForCheck();
  }
}
