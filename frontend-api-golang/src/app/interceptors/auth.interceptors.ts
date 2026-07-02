import {HttpInterceptorFn} from '@angular/common/http' //Função interceptadora
import { inject } from '@angular/core'
import { AuthService } from '../services/auth'

export const authInterceptor: HttpInterceptorFn = (req, next) => {
    const authService = inject(AuthService)
    const token = authService.getToken()

    if (token) {
        const clonedRequest = req.clone({ //faz uma cópia do header Authorization adicionado
            setHeaders: {
                Authorization: token
            }
        })
        return next(clonedRequest)
    } 
    
    return next(req)
}