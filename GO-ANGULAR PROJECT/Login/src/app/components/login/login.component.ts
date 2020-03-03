import { Component, OnInit } from '@angular/core';
import { FormBuilder, FormGroup, Validators } from '@angular/forms';
import { Router } from '@angular/router';
import { AuthService } from '../../auth.service';

@Component({
  selector: 'app-login',
  templateUrl: './login.component.html',
  styleUrls: ['./login.component.css']
})
export class LoginComponent implements OnInit {
  loginForm: FormGroup;
  isSubmitted = false;

  constructor(private authService: AuthService, private router: Router, private formBuilder: FormBuilder) { }

  ngOnInit() {
    this.loginForm = this.formBuilder.group({
      email: ['', Validators.required],
      password: ['', Validators.required]
    });
  }

  get formControls() { return this.loginForm.controls; }

  login() {
    this.isSubmitted = true;
    if (this.loginForm.invalid) {
      return;
    }
    // ask server for valid login
    this.authService.doLogin(this.loginForm.value)
      .subscribe((resp: any) => {
        console.log('API response: ', resp);
        if (resp.code === 200 && resp.message === 'Login successfully') {
          localStorage.setItem('ACCESS_TOKEN', resp.message);
          this.authService.user = resp.data.Name;
          console.log('here');
          this.router.navigateByUrl('/admin');
        } else {
          window.alert('Invalid credentials');
        }
      }, (err) => {
        console.log('Error: ', err);
      });
  }

}
