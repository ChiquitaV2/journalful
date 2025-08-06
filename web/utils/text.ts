/**
 * Text and string utility functions
 */

export const formatAuthors = (authors: Array<{ name: string }> | undefined): string => {
  if (!authors || authors.length === 0) return 'Unknown'
  if (authors.length === 1) return authors[0].name
  if (authors.length === 2) return `${authors[0].name} & ${authors[1].name}`
  return `${authors[0].name} et al.`
}

export const truncateText = (text: string, maxLength: number = 100): string => {
  if (!text || text.length <= maxLength) return text
  return text.slice(0, maxLength).trim() + '...'
}

export const generateSlug = (text: string): string => {
  return text
    .toLowerCase()
    .replace(/[^\w\s-]/g, '') // Remove special characters
    .replace(/\s+/g, '-') // Replace spaces with hyphens
    .replace(/-+/g, '-') // Replace multiple hyphens with single
    .trim()
}

export const capitalizeFirst = (text: string): string => {
  if (!text) return ''
  return text.charAt(0).toUpperCase() + text.slice(1)
}

export const generateInitials = (name: string): string => {
  const names = name.split(' ')
  return names.map(name => name.charAt(0)).join('').toUpperCase()
}

export const cleanDOI = (doi: string): string => {
  if (!doi) return ''
  // Remove common prefixes and clean the DOI
  return doi.replace(/^(doi:|DOI:|https?:\/\/(dx\.)?doi\.org\/)/i, '').trim()
}

export const formatDOI = (doi: string): string => {
  const cleaned = cleanDOI(doi)
  return cleaned ? `https://doi.org/${cleaned}` : ''
}

export const extractKeywords = (text: string): string[] => {
  if (!text) return []
  
  // Simple keyword extraction - split by common delimiters and clean
  return text
    .split(/[,;|]/)
    .map(keyword => keyword.trim())
    .filter(keyword => keyword.length > 0)
    .slice(0, 10) // Limit to 10 keywords
}
