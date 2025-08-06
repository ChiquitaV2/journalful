import { formatAuthors, formatDate, formatRelativeTime } from '../../utils'

/**
 * Composable for common formatting functions
 */
export const useFormatters = () => {
  const formatArticleAuthors = (authors: Array<{ name: string }> | undefined): string => {
    return formatAuthors(authors)
  }

  const formatArticleDate = (date: string | Date, relative = false): string => {
    if (!date) return ''
    return relative ? formatRelativeTime(date) : formatDate(date)
  }

  const formatProgress = (progress: number | undefined): string => {
    if (!progress || progress <= 0) return '0%'
    if (progress >= 100) return '100%'
    return `${Math.round(progress)}%`
  }

  const formatCount = (count: number, singular: string, plural?: string): string => {
    const p = plural || `${singular}s`
    return `${count} ${count === 1 ? singular : p}`
  }

  const formatFileSize = (bytes: number): string => {
    if (bytes === 0) return '0 Bytes'
    
    const k = 1024
    const sizes = ['Bytes', 'KB', 'MB', 'GB']
    const i = Math.floor(Math.log(bytes) / Math.log(k))
    
    return parseFloat((bytes / Math.pow(k, i)).toFixed(2)) + ' ' + sizes[i]
  }

  const formatNumber = (num: number): string => {
    if (num >= 1000000) return (num / 1000000).toFixed(1) + 'M'
    if (num >= 1000) return (num / 1000).toFixed(1) + 'K'
    return num.toString()
  }

  return {
    formatArticleAuthors,
    formatArticleDate,
    formatProgress,
    formatCount,
    formatFileSize,
    formatNumber
  }
}
