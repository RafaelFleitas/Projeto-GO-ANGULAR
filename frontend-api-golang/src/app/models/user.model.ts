export interface User {
  id: number;
  name: string;
  email: string;
  age: number;

}

export interface LoginRequest{
  email: string;
  password: string;
}

export interface LoginResponse{
  user: User;
  token: string; //Será extraído do header
}