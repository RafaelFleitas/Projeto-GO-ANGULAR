import { Routes } from '@angular/router';
import { Login } from './components/login/login';
import { Dashboard } from './components/dashboard/dashboard';

export const routes: Routes = [
    {
        path: '',
        redirectTo: 'login',
        pathMatch: 'full'

    },
    {
        path: 'login',
        component: Login,
        title: 'Dashboard'
    },
    {
        path: 'dashboard',
        component: Dashboard,
        title: 'Dashboard'
    },
    

];
