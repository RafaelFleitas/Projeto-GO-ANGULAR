import { Component, OnInit, signal } from '@angular/core'
import { CommonModule } from '@angular/common'
import { Router, ActivatedRoute } from '@angular/router'
import { FormsModule } from '@angular/forms'
import { UserService } from '../../services/user.services'
import { User } from '../../models/user.model'
import { NotificationService } from '../../services/notification.service'


@Component({
  selector: 'app-user-update',
  imports: [CommonModule, FormsModule],
  templateUrl: './user-update.html',
  styleUrl: './user-update.css',
})
export class UserUpdate implements OnInit {
  user = signal<User | null>(null)
  isLoading = signal(true)
  errorMessage = signal('')
  userName = signal('')
  userAge = signal('')
  userId: number | null = null

  constructor(
    private userService: UserService,
    private router: Router,
    private route: ActivatedRoute,
    private notificationService: NotificationService
  ) {}

  ngOnInit() {
    this.route.params.subscribe((params) => {
      this.userId = params['id']
      if (this.userId) {
        this.loadUser()
      }
    })
  }

  loadUser() {
    this.isLoading.set(true)
    this.userService.getUserById(this.userId!).subscribe({
      next: (userData) => {
        this.user.set(userData)
        this.userName.set(userData.name)
        this.userAge.set(userData.age.toString())
        this.isLoading.set(false)
      },
      error: (error) => {
        this.errorMessage.set('Erro ao carregar usuário')
        console.error(error)
        this.isLoading.set(false)
      }
    })
  }

  updateUser() {
  if (!this.userId) return

  this.isLoading.set(true)
  this.userService.updateUser(this.userId, {
    name: this.userName(),
    age: parseInt(this.userAge()),
    email: this.user()?.email || '' 
  }).subscribe({
    next: () => {
      this.isLoading.set(false)
      this.notificationService.success('Usuário atualizado com sucesso')
      this.router.navigate(['/users'])
    },
    error: (error) => {
      this.errorMessage.set('Erro ao atualizar usuário')
      console.error(error)
      this.isLoading.set(false)
    }
  })
}


  goToDashboard() {
    this.router.navigate(['/dashboard'])
  }
}
