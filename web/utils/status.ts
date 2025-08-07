import { ReadingStatus } from '#shared/types'

/**
 * Status and progress utility functions
 */

export const getStatusLabel = (status: ReadingStatus): string => {
  switch (status) {
    case ReadingStatus.READING_STATUS_TO_READ:
      return 'To Read'
    case ReadingStatus.READING_STATUS_READING:
      return 'Reading'
    case ReadingStatus.READING_STATUS_READ:
      return 'Completed'
    case ReadingStatus.READING_STATUS_ABANDONED:
      return 'Abandoned'
    default:
      return 'Unknown'
  }
}

export const getStatusColor = (status: ReadingStatus): string => {
  switch (status) {
    case ReadingStatus.READING_STATUS_TO_READ:
      return 'text-blue-400'
    case ReadingStatus.READING_STATUS_READING:
      return 'text-orange-400'
    case ReadingStatus.READING_STATUS_READ:
      return 'text-green-400'
    case ReadingStatus.READING_STATUS_ABANDONED:
      return 'text-gray-400'
    default:
      return 'text-gray-400'
  }
}

export const getStatusIcon = (status: ReadingStatus): string => {
  switch (status) {
    case ReadingStatus.READING_STATUS_TO_READ:
      return 'heroicons:bookmark'
    case ReadingStatus.READING_STATUS_READING:
      return 'heroicons:book-open'
    case ReadingStatus.READING_STATUS_READ:
      return 'heroicons:check-circle'
    case ReadingStatus.READING_STATUS_ABANDONED:
      return 'heroicons:x-circle'
    default:
      return 'heroicons:question-mark-circle'
  }
}

export const calculateProgress = (articles: Array<{ readingStatus: ReadingStatus }>): {
  total: number
  completed: number
  inProgress: number
  toRead: number
  abandoned: number
  completionRate: number
} => {
  if (!articles || articles.length === 0) {
    return {
      total: 0,
      completed: 0,
      inProgress: 0,
      toRead: 0,
      abandoned: 0,
      completionRate: 0
    }
  }

  const stats = articles.reduce((acc, article) => {
    switch (article.readingStatus) {
      case ReadingStatus.READING_STATUS_READ:
        acc.completed++
        break
      case ReadingStatus.READING_STATUS_READING:
        acc.inProgress++
        break
      case ReadingStatus.READING_STATUS_TO_READ:
        acc.toRead++
        break
      case ReadingStatus.READING_STATUS_ABANDONED:
        acc.abandoned++
        break
    }
    return acc
  }, {
    completed: 0,
    inProgress: 0,
    toRead: 0,
    abandoned: 0
  })

  return {
    ...stats,
    total: articles.length,
    completionRate: articles.length > 0 ? Math.round((stats.completed / articles.length) * 100) : 0
  }
}

export const getNextStatus = (currentStatus: ReadingStatus): ReadingStatus => {
  switch (currentStatus) {
    case ReadingStatus.READING_STATUS_TO_READ:
      return ReadingStatus.READING_STATUS_READING
    case ReadingStatus.READING_STATUS_READING:
      return ReadingStatus.READING_STATUS_READ
    case ReadingStatus.READING_STATUS_READ:
      return ReadingStatus.READING_STATUS_TO_READ
    case ReadingStatus.READING_STATUS_ABANDONED:
      return ReadingStatus.READING_STATUS_TO_READ
    default:
      return ReadingStatus.READING_STATUS_TO_READ
  }
}
