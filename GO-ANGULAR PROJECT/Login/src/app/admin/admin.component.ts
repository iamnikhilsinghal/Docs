import { Component, OnInit, ViewChild } from '@angular/core';
import { Router } from '@angular/router';
import { AuthService } from '../auth.service';
import { AdminService } from './admin.service';

import { COMMA, ENTER } from '@angular/cdk/keycodes';
import { FormGroup, FormBuilder, Validators } from '@angular/forms';

@Component({
  selector: 'app-admin',
  templateUrl: './admin.component.html',
  styleUrls: ['./admin.component.css']
})
export class AdminComponent implements OnInit {

  user: any = {};
  personList: any;
  onePerson: any;

  visible = true;
  selectable = true;
  removable = true;
  addOnBlur = true;
  myForm: FormGroup;
  @ViewChild('chipList', { static: true }) chipList;
  cityArray: any = ['Delhi', 'Gurugram', 'Noida', 'Moradabad', 'Sambhal'];
  readonly separatorKeysCodes: number[] = [ENTER, COMMA];

  constructor(private authService: AuthService, private router: Router, private adminService: AdminService, public fb: FormBuilder) { }

  ngOnInit() {
    this.user = this.authService.user;
    this.read();
  }

  reactiveForm() {
    this.myForm = this.fb.group({
      name: ['', [Validators.required]],
      email: ['', [Validators.required, Validators.email]],
      password: ['', [Validators.required]],
      gender: [''],
      dob: ['', [Validators.required]],
      city: ['', [Validators.required]],
    });
}

  date(e) {
      const convertDate = new Date(e.target.value).toISOString().substring(0, 10);
      this.myForm.get('dob').setValue(convertDate, {
        onlyself: true
      });
    }

  read() {
    this.adminService.toRead()
    .subscribe((persons: any) => {
      this.personList = persons.data;
    });
  }

  delete(clientemail) {
    this.adminService.doDelete(clientemail)
    .subscribe((resp: any) => {
      console.log('API response: ', resp);
      window.alert('Record deleted successfully')
      this.read();
    });
  }

  readForUpdate(clientemail) {
    this.adminService.doReadone(clientemail)
    .subscribe((respone: any) => {
      console.log('API response: ', respone.data);
      this.onePerson = respone.data;
      this.reactiveForm();
    });
  }

  update(totalInfo) {
console.log('totalInfo:', totalInfo);
    this.adminService.doUpdate(totalInfo.email, this.myForm.value)
    .subscribe((respUpdate: any) => {
      console.log('API response: ', respUpdate);
      this.read();
    });
  }

  logout() {
    this.authService.logout();
    this.router.navigateByUrl('/login');
  }

}

