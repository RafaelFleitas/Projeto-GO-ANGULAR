import { Component, OnInit, signal } from '@angular/core'
import { CommonModule } from '@angular/common'
import { Router } from '@angular/router'
import { UserService } from '../../services/user.services'
import { User } from '../../models/user.model'

@Component({
  selector: 'app-user-list',
  imports: [CommonModule],
  templateUrl: './user-list.html',
  styleUrl: './user-list.css'
})

export class UserList implements OnInit {
  users = signal<User[]>([])
  isLoading= signal(true)
  errorMessage= signal('')

  constructor(
    private userService: UserService,
    private router: Router
  ) {}

  ngOnInit() {
    this.loadUsers()
  }

  loadUsers() {
    this.isLoading.set(true)

    this.userService.getAllUsers().subscribe({
      next: (users) => {
        this.users.set(users)
        this.isLoading.set(false)
      },
      error: (error) => {
        this.errorMessage.set('Erro ao carregar usuários')
        console.error(error)
        this.isLoading.set(false)
      }
    })
  }

  goToDashboard() {
    this.router.navigate(['/dashboard'])
  }
}