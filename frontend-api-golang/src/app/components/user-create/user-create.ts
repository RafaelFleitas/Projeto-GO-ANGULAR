import { Component } from '@angular/core';
import { Router } from '@angular/router';
import { signal } from '@angular/core';
import { UserService } from '../../services/user.services';
import { FormsModule } from '@angular/forms';




@Component({
  selector: 'app-user-create',
  imports: [FormsModule],
  templateUrl: './user-create.html',
  styleUrl: './user-create.css',
})
export class UserCreate {
  userName = signal('')
  userEmail = signal('')
  userPassword = signal('')
  userAge = signal('')
  isLoading = signal(false)
  errorMessage = signal('')

  constructor(
    private router: Router,
    private userService: UserService
  ) { }

  createUser(){
    this.userService.createUser({
      name: this.userName(),
      email: this.userEmail(),
      password: this.userPassword(),
      age: parseInt(this.userAge())
      }).subscribe({
          next: () =>  {
            this.isLoading.set(false)
            this.router.navigate(['/users'])
          },
          error: (error) => {
            this.errorMessage.set('Erro ao criar usuário')
            console.error(error)
            this.isLoading.set(false)
          }
        })
  }



  goToDashboard() {
    this.router.navigate(['/dashboard'])
  }
}
