import { Routes } from '@angular/router';
import { Login } from './components/login/login';
import { Dashboard } from './components/dashboard/dashboard';
import { UserList } from './components/user-list/user-list';
import { UserUpdate } from './components/user-update/user-update';
import { UserCreate } from './components/user-create/user-create';


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
    {
        path: 'users/update/:id',
        component: UserUpdate,
        title: 'Update'
    },
    {
        path: 'users/create',
        component: UserCreate,
        title: 'Create'
    },
    
    
    

];
