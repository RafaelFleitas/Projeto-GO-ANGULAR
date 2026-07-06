import { Component, OnInit, signal, computed } from '@angular/core'
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
  searchTerm= signal('')

  filteredUsers = computed(() => {
    const term = this.searchTerm().trim().toLowerCase()
  

    if (!term) {
      return this.users()
    }

    return this.users().filter(user =>
      user.name.toLowerCase().includes(term) ||
      user.email.toLowerCase().includes(term)
    )
  })

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

  deleteUser(id: number) {
    this.userService.deleteUser(id).subscribe({
      next: () => {
        this.atualizarPagina()
      },
      error: (error) => {
        this.errorMessage.set('Erro ao deletar usuário')
        console.error(error)
      }
    })
  }

  
  onSearchChange(event: Event) {
    const value = (event.target as HTMLInputElement).value
    this.searchTerm.set(value)
  }

  goToDashboard() {
    this.router.navigate(['/dashboard'])
  }

  atualizarPagina(){
    this.loadUsers();
  }
}