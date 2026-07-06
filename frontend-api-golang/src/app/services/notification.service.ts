import { Injectable } from '@angular/core'
import { Notification } from '../models/notification.model'
import { signal } from '@angular/core'


@Injectable({
  providedIn: 'root'
})

export class NotificationService {
    notifications = signal<Notification[]>([])

    show(
        message: string, 
        type: 'success' | 'error' | 'warning' | 'info',
        duration: number = 3000
    ) {
        const id = `notification-${Date.now()}-${Math.random()}`

        const notification: Notification = {
            id,
            message,
            type,
            duration,
            timestamp: Date.now()
        }

        this.notifications.update(notifications => [...notifications, notification])

        setTimeout(() => {
            this.remove(id)
        }, duration)
    }

  success(message: string, duration?: number) {
    this.show(message, 'success', duration)
  }

  error(message: string, duration?: number) {
    this.show(message, 'error', duration)
  }

  info(message: string, duration?: number) {
    this.show(message, 'info', duration)
  }

  warning(message: string, duration?: number) {
    this.show(message, 'warning', duration)
  }

  remove(id: string) {
    this.notifications.update(list =>
      list.filter(n => n.id !== id)
    )
  }
}