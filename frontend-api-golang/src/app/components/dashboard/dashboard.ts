import { Component, OnInit } from '@angular/core'
import { CommonModule } from '@angular/common'
import { Router } from '@angular/router'
import { AuthService } from '../../services/auth'
import { User } from '../../models/user.model'

@Component({
  selector: 'app-dashboard',
  imports: [CommonModule],
  templateUrl: './dashboard.html',
  styleUrl: './dashboard.css'
})

export class Dashboard implements OnInit {
  currentUser: User | null = null

  constructor(
    private authService: AuthService,
    private router: Router
  ) {}

  ngOnInit() {
    // Quando o componente carrega, pega o usuário logado do localStorage
    this.currentUser = this.authService.getCurrentUser()

    // Se não houver usuário, redireciona para login
    if (!this.currentUser) {
      this.router.navigate(['/login'])
    }
  }

  logout() {
    this.authService.logout()
    this.router.navigate(['/login'])
  }

  //Novas funções para navegação

  goToUserList() {
    this.router.navigate(['/users'])
  }

  goToCreateUser() {
    this.router.navigate(['/users/new'])
  }
}
