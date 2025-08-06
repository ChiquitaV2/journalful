/**
 * Date utility functions
 */

export const formatDate = (date: string | Date, options?: Intl.DateTimeFormatOptions): string => {
  if (!date) return ''
  
  const defaultOptions: Intl.DateTimeFormatOptions = {
    year: 'numeric',
    month: 'short',
    day: 'numeric'
  }
  
  return new Date(date).toLocaleDateString('en-US', { ...defaultOptions, ...options })
}

export const formatDateLong = (date: string | Date): string => {
  return formatDate(date, {
    year: 'numeric',
    month: 'long',
    day: 'numeric'
  })
}

export const formatRelativeTime = (date: string | Date): string => {
  if (!date) return ''
  
  const now = new Date()
  const past = new Date(date)
  const diffMs = now.getTime() - past.getTime()
  
  const diffMinutes = Math.floor(diffMs / (1000 * 60))
  const diffHours = Math.floor(diffMs / (1000 * 60 * 60))
  const diffDays = Math.floor(diffMs / (1000 * 60 * 60 * 24))
  const diffWeeks = Math.floor(diffDays / 7)
  const diffMonths = Math.floor(diffDays / 30)
  
  if (diffMinutes < 1) return 'Just now'
  if (diffMinutes < 60) return `${diffMinutes} minute${diffMinutes > 1 ? 's' : ''} ago`
  if (diffHours < 24) return `${diffHours} hour${diffHours > 1 ? 's' : ''} ago`
  if (diffDays < 7) return `${diffDays} day${diffDays > 1 ? 's' : ''} ago`
  if (diffWeeks < 4) return `${diffWeeks} week${diffWeeks > 1 ? 's' : ''} ago`
  if (diffMonths < 12) return `${diffMonths} month${diffMonths > 1 ? 's' : ''} ago`
  
  return formatDate(date)
}

export const isToday = (date: string | Date): boolean => {
  const today = new Date()
  const checkDate = new Date(date)
  
  return today.toDateString() === checkDate.toDateString()
}

export const isWithinWeek = (date: string | Date): boolean => {
  const now = new Date()
  const checkDate = new Date(date)
  const diffMs = now.getTime() - checkDate.getTime()
  const diffDays = Math.floor(diffMs / (1000 * 60 * 60 * 24))
  
  return diffDays <= 7
}
