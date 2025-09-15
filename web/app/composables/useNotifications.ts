interface Notification {
  id?: string
  type: 'success' | 'error' | 'warning' | 'info'
  message: string
  duration?: number
}
export const useNotifications = () => {
  // Global notifications state using useState (Nuxt composable)
  const notifications = useState<Notification[]>('notifications', () => [])

  const addNotification = (notification: Notification) => {
    const id = Date.now().toString() + Math.random().toString(36).substr(2, 9)
    const newNotification = {
      id,
      // duration: 5000, // 5 seconds default
      ...notification
    }
    notifications.value.push(newNotification)
    return id
  }

  const removeNotification = (id: string) => {
    const index = notifications.value.findIndex(n => n.id === id)
    if (index > -1) {
      notifications.value.splice(index, 1)
    }
  }

  const clearAll = () => {
    notifications.value = []
  }

  // Convenience methods
  const showSuccess = (message: string, duration: number) => {
    return addNotification({ type: 'success', message, duration })
  }

  const showError = (message: string, duration: number) => {
    return addNotification({ type: 'error', message, duration })
  }

  const showWarning = (message: string, duration: number) => {
    return addNotification({ type: 'warning', message, duration })
  }

  const showInfo = (message: string, duration: number) => {
    return addNotification({ type: 'info', message, duration })
  }

  return {
    notifications,
    addNotification,
    removeNotification,
    clearAll,
    showSuccess,
    showError,
    showWarning,
    showInfo
  }
}
