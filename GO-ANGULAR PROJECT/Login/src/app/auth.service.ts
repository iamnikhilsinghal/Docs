import { Injectable } from '@angular/core';
import { HttpClient } from '@angular/common/http';

@Injectable({
  providedIn: 'root'
})
export class AuthService {

  public user: any;

  constructor(private http: HttpClient) { }

  public isLoggedIn(){
    return localStorage.getItem('ACCESS_TOKEN') !== null;

  }

  public logout(){
    localStorage.removeItem('ACCESS_TOKEN');
  }

  public doLogin(credentials: any) {
    return this.http.post('http://localhost:8080/login', credentials);
  }
}
