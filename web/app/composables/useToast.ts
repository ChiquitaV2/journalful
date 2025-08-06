interface Toast {
  id: string
  type: 'success' | 'error' | 'warning' | 'info'
  title: string
  message?: string
  duration?: number
  persistent?: boolean
}

interface ToastOptions {
  type?: Toast['type']
  title: string
  message?: string
  duration?: number
  persistent?: boolean
}

export const useToast = () => {
  const toasts = useState<Toast[]>('toasts', () => [])

  const addToast = (options: ToastOptions): string => {
    const id = Date.now().toString() + Math.random().toString(36).substr(2, 9)
    
    const toast: Toast = {
      id,
      type: options.type || 'info',
      title: options.title,
      message: options.message,
      duration: options.duration || 5000,
      persistent: options.persistent || false
    }

    toasts.value.push(toast)

    // Auto-remove toast after duration (unless persistent)
    if (!toast.persistent && toast.duration && toast.duration > 0) {
      setTimeout(() => {
        removeToast(id)
      }, toast.duration)
    }

    return id
  }

  const removeToast = (id: string) => {
    const index = toasts.value.findIndex(toast => toast.id === id)
    if (index > -1) {
      toasts.value.splice(index, 1)
    }
  }

  const clearAll = () => {
    toasts.value = []
  }

  // Convenience methods
  const success = (title: string, message?: string, duration?: number) => {
    return addToast({ type: 'success', title, message, duration })
  }

  const error = (title: string, message?: string, persistent = false) => {
    return addToast({ type: 'error', title, message, persistent, duration: persistent ? 0 : 8000 })
  }

  const warning = (title: string, message?: string, duration?: number) => {
    return addToast({ type: 'warning', title, message, duration: duration || 6000 })
  }

  const info = (title: string, message?: string, duration?: number) => {
    return addToast({ type: 'info', title, message, duration })
  }

  return {
    toasts: readonly(toasts),
    addToast,
    removeToast,
    clearAll,
    success,
    error,
    warning,
    info
  }
}
