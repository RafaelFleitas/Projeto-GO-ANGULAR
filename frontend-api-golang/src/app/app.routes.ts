import { Routes } from '@angular/router';
import { Login } from './components/login/login';
import { Dashboard } from './components/dashboard/dashboard';
import { UserList } from './components/user-list/user-list';


export const routes: Routes = [
    {
        path: '',
        redirectTo: 'login',
        pathMatch: 'full'

    },
    {
        path: 'login',
        component: Login,
        title: 'Login'
    },
    {
        path: 'dashboard',
        component: Dashboard,
        title: 'Dashboard'
    },
    {
        path: 'users',
        component: UserList,
        title: 'Usuários'
    },
    
    

];
