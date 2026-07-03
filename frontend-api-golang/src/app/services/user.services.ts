import { Injectable } from '@angular/core'
import { HttpClient } from '@angular/common/http'
import { Observable } from 'rxjs'
import { User } from '../models/user.model'
import { CreateRequest } from '../models/user.model'


@Injectable({
  providedIn: 'root' //Significa que o serviço será injetado no root do projeto, ou seja, em qualquer lugar do projeto podemos usar ele
})

export class UserService {
    private apiUrl = 'http://localhost:8000'

    constructor(private http: HttpClient) {}

    createUser(user: CreateRequest): Observable<User> {
        return this.http.post<User>(`${this.apiUrl}/createUser`, user)
    }

    getUserById(userId: number): Observable<User> {
        return this.http.get<User>(`${this.apiUrl}/getUserById/${userId}`)
    }

    getUserByEmail(userEmail: string): Observable<User> {
        return this.http.get<User>(`${this.apiUrl}/getUserByEmail/${userEmail}`)
    }
    getAllUsers(): Observable<User[]> {
        return this.http.get<User[]>(`${this.apiUrl}/getAllUsers`)
    }
    updateUser(userId: number, user: Omit<User, 'id'>): Observable<User> {
        return this.http.put<User>(`${this.apiUrl}/updateUser/${userId}`, user)
    }

    deleteUser(userId: number): Observable<void>{
        return this.http.delete<void>(`${this.apiUrl}/deleteUser/${userId}`)
    }

}