/**
 * Browser and environment utility functions
 */

export const copyToClipboard = async (text: string): Promise<boolean> => {
  try {
    if (navigator.clipboard && window.isSecureContext) {
      await navigator.clipboard.writeText(text)
      return true
    } else {
      // Fallback for older browsers
      const textArea = document.createElement('textarea')
      textArea.value = text
      textArea.style.position = 'fixed'
      textArea.style.left = '-999999px'
      textArea.style.top = '-999999px'
      document.body.appendChild(textArea)
      textArea.focus()
      textArea.select()
      
      const result = document.execCommand('copy')
      document.body.removeChild(textArea)
      return result
    }
  } catch (error) {
    console.error('Failed to copy to clipboard:', error)
    return false
  }
}

export const shareContent = async (content: {
  title?: string
  text?: string
  url?: string
}): Promise<boolean> => {
  try {
    if (navigator.share) {
      await navigator.share(content)
      return true
    } else {
      // Fallback to clipboard
      const shareText = `${content.title || ''}\n${content.text || ''}\n${content.url || ''}`.trim()
      return await copyToClipboard(shareText)
    }
  } catch (error) {
    console.error('Failed to share content:', error)
    return false
  }
}

export const generateUniqueId = (): string => {
  return Date.now().toString(36) + Math.random().toString(36).substr(2)
}

export const downloadFile = (content: string, filename: string, contentType: string = 'text/plain'): void => {
  const blob = new Blob([content], { type: contentType })
  const url = URL.createObjectURL(blob)
  
  const link = document.createElement('a')
  link.href = url
  link.download = filename
  document.body.appendChild(link)
  link.click()
  
  document.body.removeChild(link)
  URL.revokeObjectURL(url)
}

export const isMobileDevice = (): boolean => {
  return /Android|webOS|iPhone|iPad|iPod|BlackBerry|IEMobile|Opera Mini/i.test(navigator.userAgent)
}

export const isTouchDevice = (): boolean => {
  return 'ontouchstart' in window || navigator.maxTouchPoints > 0
}

export const debounce = <T extends (...args: any[]) => any>(
  func: T,
  wait: number
): (...args: Parameters<T>) => void => {
  let timeout: number | null = null
  
  return (...args: Parameters<T>) => {
    if (timeout) clearTimeout(timeout)
    timeout = setTimeout(() => func(...args), wait)
  }
}

export const throttle = <T extends (...args: any[]) => any>(
  func: T,
  limit: number
): (...args: Parameters<T>) => void => {
  let inThrottle: boolean
  
  return (...args: Parameters<T>) => {
    if (!inThrottle) {
      func(...args)
      inThrottle = true
      setTimeout(() => inThrottle = false, limit)
    }
  }
}
