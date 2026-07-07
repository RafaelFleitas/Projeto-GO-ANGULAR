import { Component, OnInit } from '@angular/core'
import { Router } from '@angular/router'
import { AuthService } from '../../services/auth'
import { User } from '../../models/user.model'


@Component({
  selector: 'app-dashboard',
  imports: [],
  templateUrl: './dashboard.html',
  styleUrl: './dashboard.css'
})

export class Dashboard implements OnInit {
  currentUser: User | null = null
  isMenuOpen = false


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

  //Funções do menu
  toggleMenu() {
    this.isMenuOpen = !this.isMenuOpen
  }

  closeMenu() {
    this.isMenuOpen = false
  }

  goToProfile() {
    this.closeMenu()
    this.router.navigate(['/profile'])
  }

  logout() {
    this.authService.logout()
    this.router.navigate(['/login'])
  }

  getProfileImage(): string {
    return this.currentUser?.profileImage || '/default-avatar.png'
  }

  //Funções da grid
  goToUserList() {
    this.router.navigate(['/users'])
  }

  goToCreateUser() {
    this.router.navigate(['/users/create'])
  }
  goToUpdateUser() {
    const id = prompt('Digite o ID do usuário que deseja atualizar:')
    if (id){
      this.router.navigate(['/users/update', id])
    }
  }
}
