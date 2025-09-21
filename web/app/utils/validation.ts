/**
 * Validation utility functions
 */

export const isValidEmail = (email: string): boolean => {
  const emailRegex = /^[^\s@]+@[^\s@]+\.[^\s@]+$/
  return emailRegex.test(email)
}

export const isValidDOI = (doi: string): boolean => {
  if (!doi) return false
  
  // Clean the DOI first
  const cleaned = doi.replace(/^(doi:|DOI:|https?:\/\/(dx\.)?doi\.org\/)/i, '').trim()
  
  // Basic DOI pattern validation
  const doiRegex = /^10\.\d{4,}[\/\.][\w\.\-\(\)\/;:]+$/i
  return doiRegex.test(cleaned)
}

export const isValidURL = (url: string): boolean => {
  try {
    new URL(url)
    return true
  } catch {
    return false
  }
}

export const validateRequired = (value: any, fieldName: string): string | null => {
  if (!value || (typeof value === 'string' && value.trim() === '')) {
    return `${fieldName} is required`
  }
  return null
}

export const validateMinLength = (value: string, minLength: number, fieldName: string): string | null => {
  if (value && value.length < minLength) {
    return `${fieldName} must be at least ${minLength} characters long`
  }
  return null
}

export const validateMaxLength = (value: string, maxLength: number, fieldName: string): string | null => {
  if (value && value.length > maxLength) {
    return `${fieldName} must be no more than ${maxLength} characters long`
  }
  return null
}

export const validateYear = (year: number): boolean => {
  const currentYear = new Date().getFullYear()
  return year >= 1900 && year <= currentYear + 1
}

export const sanitizeInput = (input: string): string => {
  if (!input) return ''
  
  return input
    .trim()
    .replace(/[<>]/g, '') // Remove potential HTML brackets
    .replace(/[\r\n\t]/g, ' ') // Replace newlines and tabs with spaces
    .replace(/\s+/g, ' ') // Replace multiple spaces with single space
}
