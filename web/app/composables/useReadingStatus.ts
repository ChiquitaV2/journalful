import { ReadingStatus } from '#shared/types'
import { getStatusLabel, getStatusColor, getStatusIcon, calculateProgress, getNextStatus } from '~/utils'

/**
 * Composable for reading status management
 */
export const useReadingStatus = () => {
  const getStatusDisplay = (status: ReadingStatus) => ({
    label: getStatusLabel(status),
    color: getStatusColor(status),
    icon: getStatusIcon(status)
  })

  const getLibraryStats = (articles: Array<{ readingStatus: ReadingStatus }>) => {
    return calculateProgress(articles)
  }

  const cycleStatus = (currentStatus: ReadingStatus): ReadingStatus => {
    return getNextStatus(currentStatus)
  }

  const getStatusOptions = () => [
    {
      value: ReadingStatus.READING_STATUS_TO_READ,
      label: 'To Read',
      icon: 'heroicons:bookmark',
      color: 'text-blue-400'
    },
    {
      value: ReadingStatus.READING_STATUS_READING,
      label: 'Currently Reading',
      icon: 'heroicons:book-open',
      color: 'text-orange-400'
    },
    {
      value: ReadingStatus.READING_STATUS_READ,
      label: 'Completed',
      icon: 'heroicons:check-circle',
      color: 'text-green-400'
    },
    {
      value: ReadingStatus.READING_STATUS_ABANDONED,
      label: 'Abandoned',
      icon: 'heroicons:x-circle',
      color: 'text-gray-400'
    }
  ]

  const updateArticleStatus = async (articleId: number, newStatus: ReadingStatus) => {
    // Mock API call - replace with actual implementation
    await new Promise(resolve => setTimeout(resolve, 300))
    
    // In a real app, this would make an API call
    console.log(`Updating article ${articleId} to status ${newStatus}`)
    
    return { success: true }
  }

  return {
    getStatusDisplay,
    getLibraryStats,
    cycleStatus,
    getStatusOptions,
    updateArticleStatus
  }
}
