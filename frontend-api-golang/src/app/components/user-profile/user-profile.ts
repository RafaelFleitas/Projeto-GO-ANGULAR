import { Component, OnInit } from '@angular/core'
import { CommonModule } from '@angular/common'
import { Router } from '@angular/router'
import { AuthService } from '../../services/auth'
import { User } from '../../models/user.model'
import { NotificationService } from '../../services/notification.service'


@Component({
  selector: 'app-user-profile',
  imports: [CommonModule],
  templateUrl: './user-profile.html',
  styleUrl: './user-profile.css',
})
export class UserProfile implements OnInit {
  currentUser: User | null = null
  selectedFile: File | null = null
  previewUrl: string | null = null
  isLoading = false

  constructor(
    private authService: AuthService,
    private router: Router,
    private notificationService: NotificationService
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
      this.previewUrl = e.target.result
    }
    reader.readAsDataURL(file)
  }

  saveProfileImage(){
    if (!this.selectedFile || !this.currentUser){
      this.notificationService.warning('Selecione uma imagem antes de salvar')
      return
    }
    this.isLoading = true


    const reader = new FileReader()
    
    reader.onload = () => {
      const base64String = reader.result as string
      this.currentUser!.profileImage = base64String
      localStorage.setItem('current_user', JSON.stringify(this.currentUser))
      this.notificationService.success('Foto atualizada com sucesso!')
      this.isLoading = false
      this.resetForm()
    }

    reader.readAsDataURL(this.selectedFile)
  }

  resetForm() {
    this.selectedFile = null
    this.previewUrl = null
  }

  goBack() {
    this.router.navigate(['/dashboard'])
  }

}
