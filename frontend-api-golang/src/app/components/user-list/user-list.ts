import { Component, OnInit } from '@angular/core'
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
  users: User[] = []
  isLoading: boolean = true
  errorMessage: string = ''

  constructor(
    private userService: UserService,
    private router: Router
  ) {}

  ngOnInit() {
    this.loadUsers()
  }

  loadUsers() {
    this.isLoading = true

    this.userService.getAllUsers().subscribe({
      next: (users) => {
        this.users = users
        this.isLoading = false
      },
      error: (error) => {
        this.errorMessage = 'Erro ao carregar usuários'
        console.error(error)
        this.isLoading = false
      }
    })
  }

  goToDashboard() {
    this.router.navigate(['/dashboard'])
  }
}