import { copyToClipboard, shareContent } from '../../utils'

/**
 * Composable for sharing functionality
 */
export const useSharing = () => {
  const { success, info } = useToast()

  const shareLibrary = async (library: any) => {
    const shareData = {
      title: library.name,
      text: `Check out my reading library: ${library.name}`,
      url: `${window.location.origin}/libraries/shared/${library.id}`
    }

    try {
      const shared = await shareContent(shareData)
      
      if (shared) {
        success('Shared Successfully', 'Library has been shared!')
      } else {
        // Fallback to clipboard
        const copied = await copyToClipboard(shareData.url)
        if (copied) {
          success('Link Copied', 'Share link has been copied to your clipboard.')
        } else {
          info('Share Link', `Share link: ${shareData.url}`)
        }
      }
    } catch (error) {
      console.error('Failed to share:', error)
    }
  }

  const shareArticle = async (article: any) => {
    const shareData = {
      title: article.title,
      text: `Interesting research: ${article.title}`,
      url: article.url || `${window.location.origin}/articles/${article.id}`
    }

    try {
      const shared = await shareContent(shareData)
      
      if (shared) {
        success('Shared Successfully', 'Article has been shared!')
      } else {
        const copied = await copyToClipboard(shareData.url)
        if (copied) {
          success('Link Copied', 'Article link has been copied to your clipboard.')
        } else {
          info('Share Link', `Article link: ${shareData.url}`)
        }
      }
    } catch (error) {
      console.error('Failed to share:', error)
    }
  }

  const copyDOI = async (doi: string) => {
    const copied = await copyToClipboard(doi)
    
    if (copied) {
      success('DOI Copied', 'DOI has been copied to your clipboard.')
    } else {
      info('DOI', doi)
    }
  }

  const exportLibraryData = (library: any, format: 'json' | 'csv' = 'json') => {
    // This would be implemented based on requirements
    info('Coming Soon', `Export to ${format.toUpperCase()} will be available in the next update!`)
  }

  return {
    shareLibrary,
    shareArticle,
    copyDOI,
    exportLibraryData
  }
}
