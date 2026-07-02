import { Injectable } from '@angular/core';
import { Observable } from 'rxjs/internal/Observable';
import { HttpClient } from '@angular/common/http';
import { User, LoginRequest } from '../models/user.model';
import { tap } from 'rxjs/operators';

@Injectable({
    providedIn: 'root'
})

export class AuthService {

    private tokenKey = 'jwt_token'
    private userKey = 'current_user'
    private apiUrl = 'http://localhost:8000'


    //Função que roda quando a classe é criada para receber o HttpClient do Angular
    constructor(private http: HttpClient) { }


    //Faz o login do usuário e retorna um Observable<User> com os dados do usuário logado
    login(email: string, password: string): Observable<User>{
        const request: LoginRequest = { email, password }

        // Faz a requisição POST para a API de login e armazena o usuário logado no localStorage
        return this.http.post<User>(`${this.apiUrl}/login`, request).pipe(
            tap((user: User) => {
                localStorage.setItem(this.userKey, JSON.stringify(user))
            })
        )
    }

    //Verifica se o usuário está autenticado
    isAuthenticated(): boolean {
        return !!localStorage.getItem(this.tokenKey)
    }

    //Retorna o token do usuário
    getToken(): string | null {
        return localStorage.getItem(this.tokenKey)
    }

    //Faz o logout do usuário removendo suas informações do localStorage
    logout(): void {
        localStorage.removeItem(this.tokenKey)
        localStorage.removeItem(this.userKey)
    }

    //Retorna o usuário logado
    getCurrentUser(): User | null {
        const user = localStorage.getItem(this.userKey)
        return user ? JSON.parse(user) : null
    }
}
