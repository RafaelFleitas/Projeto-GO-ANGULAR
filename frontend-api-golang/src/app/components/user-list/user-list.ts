import { Component, OnInit, signal, computed } from '@angular/core'
import { Router } from '@angular/router'
import { UserService } from '../../services/user.services'
import { User } from '../../models/user.model'
import { NotificationService } from '../../services/notification.service'
import { AuthService } from '../../services/auth'


@Component({
  selector: 'app-user-list',
  imports: [],
  templateUrl: './user-list.html',
  styleUrl: './user-list.css'
})

export class UserList implements OnInit {
  users = signal<User[]>([])
  isLoading = signal(true)
  errorMessage = signal('')
  searchTerm = signal('')

  currentPage = signal(1)
  pageSize = signal(10)
  total = signal(0)
  totalPages = signal(0)

  filteredUsers = computed(() => {
    const term = this.searchTerm().trim().toLowerCase()

    if (!term) {
      return this.users()
    }

    return this.users().filter(user =>
      user.name.toLowerCase().includes(term) || user.email.toLowerCase().includes(term)
    )
  })

  constructor(
    private authService: AuthService,
    private userService: UserService,
    private router: Router,
    private notificationService: NotificationService
  ) {}

  ngOnInit() {
    if (!this.authService.isAuthenticated()) {
      this.router.navigate(['/login'])
    }
    this.loadUsers()
  }

  loadUsers() {
    this.isLoading.set(true)

    this.userService.getAllUsers(this.currentPage(), this.pageSize()).subscribe({
      next: (response) => {
        this.users.set(response.users)
        this.total.set(response.total)
        this.totalPages.set(response.totalPages)
        this.isLoading.set(false)
      },
      error: (error) => {
        this.notificationService.error('Erro ao carregar usuários')
        console.error(error)
        this.isLoading.set(false)
      }
    })
  }

  deleteUser(id: number) {
    this.userService.deleteUser(id).subscribe({
      next: () => {
        this.notificationService.success('Usuário deletado com sucesso')
        this.atualizarPagina()
      },
      error: (error) => {
        this.notificationService.error('Erro ao deletar usuários')
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

  nextPage() {
    if (this.currentPage() < this.totalPages()) {
      this.currentPage.set(this.currentPage() + 1)
      this.loadUsers()
    }
  }

  previousPage() {
    if (this.currentPage() > 1) {
      this.currentPage.set(this.currentPage() - 1)
      this.loadUsers()
    }
  }
}
