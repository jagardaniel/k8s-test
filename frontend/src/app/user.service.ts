import { Injectable } from '@angular/core';
import { HttpClient, HttpHeaders } from '@angular/common/http';

import { Observable } from 'rxjs';

import { User } from './user';

@Injectable({
  providedIn: 'root'
})
export class UserService {
  private usersUrl = '/api/users';

  httpOptions = {
    headers: new HttpHeaders({ 'Content-Type': 'application/json' })
  };

  /** GET all users from the server **/
  getUsers(): Observable<User[]> {
    return this.http.get<User[]>(this.usersUrl)
  }

  constructor(private http: HttpClient) { }
}
