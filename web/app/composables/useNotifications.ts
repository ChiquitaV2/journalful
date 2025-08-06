export const useNotifications = () => {
  // Global notifications state using useState (Nuxt composable)
  const notifications = useState('notifications', () => [])

  const addNotification = (notification) => {
    const id = Date.now().toString() + Math.random().toString(36).substr(2, 9)
    const newNotification = {
      id,
      duration: 5000, // 5 seconds default
      ...notification
    }
    notifications.value.push(newNotification)
    return id
  }

  const removeNotification = (id) => {
    const index = notifications.value.findIndex(n => n.id === id)
    if (index > -1) {
      notifications.value.splice(index, 1)
    }
  }

  const clearAll = () => {
    notifications.value = []
  }

  // Convenience methods
  const showSuccess = (message, duration) => {
    return addNotification({ type: 'success', message, duration })
  }

  const showError = (message, duration) => {
    return addNotification({ type: 'error', message, duration })
  }

  const showWarning = (message, duration) => {
    return addNotification({ type: 'warning', message, duration })
  }

  const showInfo = (message, duration) => {
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
