import { Component, signal } from '@angular/core'
import { FormsModule } from '@angular/forms'
import { Router } from '@angular/router'
import { AuthService } from '../../services/auth'

@Component({
  selector: 'app-login',
  imports: [FormsModule],
  templateUrl: './login.html',
  styleUrl: './login.css',
})
export class Login {
  email: string = ''
  password: string = ''
  errorMessage: string = ''
  isLoading = signal(false)

  constructor(
    private authService: AuthService,
    private router: Router,
  ){}

  onLogin(){
    if(!this.email || !this.password){
      this.errorMessage = 'Por favor, preencha todos os campos.'
      return
    }
  

  this.isLoading.set(true)
  this.authService.login(this.email, this.password).subscribe({
    next: (response) => {
      this.router.navigate(['/dashboard'])
    },
    error: (error) => {
      this.errorMessage = 'Email ou senha inválidos. Por favor, tente novamente.'
      console.error('Erro no login:', error)
      this.isLoading.set(false)
    }
  })

}
}
