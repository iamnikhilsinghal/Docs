import { Injectable } from '@angular/core';
import { HttpClient } from '@angular/common/http';

@Injectable({
  providedIn: 'root'
})
export class AdminService {

  constructor(private http: HttpClient) { }

toRead() {
  return this.http.get('http://localhost:8080/read');
}

doDelete(clientmail) {
  return this.http.post('http://localhost:8080/delete', {
    email: clientmail
  });
}

doReadone(clientmail) {
  return this.http.post('http://localhost:8080/readone', {
    email: clientmail
  });
}

doUpdate(clientmail, form) {
  console.log("data :",form)
  return this.http.put(`http://localhost:8080/update/${clientmail}`, form);
}

}
