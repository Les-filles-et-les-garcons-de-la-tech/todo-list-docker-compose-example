import { HttpClient } from '@angular/common/http';
import { Injectable } from '@angular/core';
import { Subject } from 'rxjs';
import { environment } from 'src/environments/environment';
import { Color } from './models/app.models.color';

@Injectable({
  providedIn: 'root'
})
export class ColorService {
  constructor(private http: HttpClient) { }
  private URL : string = environment.apiUrl+"/api/color"
  public colorSubject : Subject<string> = new Subject();
  
  getColor(): void {

    this.http.get<Color>(this.URL).subscribe(data => {
      var color : string = "#FFFFFF";
      console.log(data.color);
      if(data!=null){
        color=data.color;
      }
      this.colorSubject.next(color);
    })
  }
}
