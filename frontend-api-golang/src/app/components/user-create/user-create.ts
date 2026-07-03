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
    if (!this.validarFormulario()){
      return
    }

    this.isLoading.set(true)

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
            this.isLoading.set(false)

            if (error.error?.message === 'User email already exists') {
            this.errorMessage.set('Este email já está cadastrado')
          } 
          }
        })
  }

  validarFormulario(){
    const senha = this.userPassword()
    const idade = parseInt(this.userAge())
    const caracterEspecial = /[!@#$%*]/

    if (senha.length < 6) {
      this.errorMessage.set('A senha deve ter no mínimo 6 caracteres')
      return false
    } else if(!caracterEspecial.test(senha)){
      this.errorMessage.set('A senha deve conter pelo menos um caracter especial')
      return false
    }

    if (isNaN(idade) ||idade < 1 || idade > 100){
      this.errorMessage.set('A idade deve ser um número entre 1 e 100')
      return false
    }
    return true

  }



  goToDashboard() {
    this.router.navigate(['/dashboard'])
  }
}
