import { Component, OnInit, signal } from '@angular/core'
import { Router } from '@angular/router'
import { AuthService } from '../../services/auth'
import { User } from '../../models/user.model'
import { NotificationService } from '../../services/notification.service'
import { HttpClient } from '@angular/common/http'


@Component({
  selector: 'app-user-profile',
  imports: [],
  templateUrl: './user-profile.html',
  styleUrl: './user-profile.css',
})
export class UserProfile implements OnInit {
  currentUser: User | null = null
  selectedFile: File | null = null
  previewUrl= signal<string | null>(null)
  isLoading = signal(false)


  constructor(
    private authService: AuthService,
    private router: Router,
    private notificationService: NotificationService,
    private http: HttpClient
  ){}

  ngOnInit(){
    this.currentUser = this.authService.getCurrentUser()

    if(!this.currentUser){
      this.router.navigate(['/login'])
    }

  }



  onFileSelected(event: any){
    const file = event.target.files[0]

    if(file){
      if(!file.type.startsWith('image/')){
        this.notificationService.error('Por favor, selecione uma imagem válida')
        return
      }

      if(file.size > 5 * 1024 * 1024){
        this.notificationService.error('A imagem não pode ser maior que 2MB')
        return
      }
    }

    this.selectedFile = file

    const reader = new FileReader()
    reader.onload = (e: any) => {
      this.previewUrl.set(e.target.result)
    }
    reader.readAsDataURL(file)
  }

  saveProfileImage(){
    if (!this.selectedFile || !this.currentUser){
      this.notificationService.warning('Selecione uma imagem antes de salvar')
      return
    }
    this.isLoading.set(true)

    const formData = new FormData()
    formData.append('avatar', this.selectedFile)

    this.http.post<User>('http://localhost:8000/upload/avatar', formData).subscribe({
      next: (updatedUser) => {
        this.currentUser = updatedUser
        localStorage.setItem('current_user', JSON.stringify(updatedUser))
        this.notificationService.success('Foto atualizada com sucesso!')
        this.isLoading.set(false)
        this.resetForm()
      },
      error: (err) => {
        console.error('Error ao salvar: ', err)
        this.notificationService.error('Erro ao salvar imagem')
        this.isLoading.set(false)
      }
    })
  }

  resetForm() {
    this.selectedFile = null
    this.previewUrl.set(null)
  }

  goBack() {
    this.router.navigate(['/dashboard'])
  }

}
