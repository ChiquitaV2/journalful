import { ReadingStatus } from '#shared/types'

/**
 * Mock data generator for development and testing
 */

export const generateMockArticle = (id?: number, overrides?: any) => ({
  id: id || Math.floor(Math.random() * 10000),
  doi: `10.1000/${Math.random().toString(36).substr(2, 9)}`,
  title: overrides?.title || generateMockTitle(),
  authors: overrides?.authors || generateMockAuthors(),
  abstract: overrides?.abstract || generateMockAbstract(),
  publicationYear: overrides?.publicationYear || (2020 + Math.floor(Math.random() * 5)),
  journalName: overrides?.journalName || generateMockJournal(),
  createdAt: new Date(Date.now() - Math.random() * 365 * 24 * 60 * 60 * 1000),
  updatedAt: new Date(),
  ...overrides
})

export const generateMockLibraryArticle = (articleId?: number, overrides?: any) => ({
  id: Math.floor(Math.random() * 10000),
  articleId: articleId || Math.floor(Math.random() * 10000),
  readingStatus: overrides?.readingStatus || getRandomStatus(),
  addedAt: new Date(Date.now() - Math.random() * 30 * 24 * 60 * 60 * 1000),
  notes: overrides?.notes || (Math.random() > 0.7 ? generateMockNotes() : undefined),
  articleTitle: overrides?.articleTitle || generateMockTitle(),
  doi: `10.1000/${Math.random().toString(36).substr(2, 9)}`,
  publicationYear: 2020 + Math.floor(Math.random() * 5),
  ...overrides
})

export const generateMockLibrary = (id?: number, overrides?: any) => {
  const articleCount = overrides?.articleCount || Math.floor(Math.random() * 50) + 5
  const articles = Array.from({ length: articleCount }, () => generateMockLibraryArticle())
  
  return {
    id: id || Math.floor(Math.random() * 10000),
    userId: 1,
    name: overrides?.name || generateMockLibraryName(),
    description: overrides?.description || generateMockDescription(),
    articles,
    createdAt: new Date(Date.now() - Math.random() * 365 * 24 * 60 * 60 * 1000),
    updatedAt: new Date(),
    ...overrides
  }
}

export const generateMockProfile = (id?: number, overrides?: any) => ({
  id: id || 1,
  name: overrides?.name || 'Dr. Alex Johnson',
  bio: overrides?.bio || 'Machine Learning Researcher focused on academic literature analysis and citation networks.',
  institution: overrides?.institution || 'University of Research Excellence',
  createdAt: new Date('2024-01-01'),
  updatedAt: new Date(),
  ...overrides
})

// Helper functions for generating realistic mock data
const titles = [
  'Advances in Machine Learning for Academic Research',
  'Deep Learning Applications in Scientific Literature Analysis',
  'Neural Networks for Citation Pattern Recognition',
  'Transformer Models in Academic Text Processing',
  'Computer Vision Techniques for Document Analysis',
  'Natural Language Processing in Research Papers',
  'Automated Literature Review Using AI',
  'Semantic Analysis of Scientific Publications',
  'Knowledge Graph Construction from Academic Papers',
  'Multi-modal Learning for Research Document Understanding'
]

const authors = [
  'Dr. Sarah Chen',
  'Prof. Michael Rodriguez',
  'Dr. Emily Watson',
  'Prof. David Kim',
  'Dr. Anna Kowalski',
  'Prof. James Thompson',
  'Dr. Maria Garcia',
  'Prof. Robert Wilson',
  'Dr. Lisa Zhang',
  'Prof. Ahmed Hassan'
]

const journals = [
  'Nature Machine Intelligence',
  'Journal of Academic Computing',
  'IEEE Transactions on Knowledge and Data Engineering',
  'ACM Computing Surveys',
  'Artificial Intelligence Review',
  'Machine Learning Journal',
  'Neural Networks and Deep Learning',
  'Information Sciences',
  'Pattern Recognition Letters',
  'Computer Science Review'
]

const libraryNames = [
  'Machine Learning Research',
  'Computer Vision Papers',
  'Natural Language Processing',
  'Deep Learning Foundations',
  'AI Ethics and Society',
  'Reinforcement Learning',
  'Neural Network Architectures',
  'Data Mining Techniques',
  'Information Retrieval',
  'Computational Linguistics'
]

function generateMockTitle(): string {
  return titles[Math.floor(Math.random() * titles.length)]
}

function generateMockAuthors(): Array<{ id: number; name: string; profileId: number }> {
  const count = Math.floor(Math.random() * 4) + 1
  const selectedAuthors = []
  
  for (let i = 0; i < count; i++) {
    selectedAuthors.push({
      id: Math.floor(Math.random() * 10000),
      name: authors[Math.floor(Math.random() * authors.length)],
      profileId: Math.floor(Math.random() * 100)
    })
  }
  
  return selectedAuthors
}

function generateMockAbstract(): string {
  const abstracts = [
    'This paper presents novel approaches to machine learning applications in academic research, focusing on automated literature analysis and citation pattern recognition.',
    'We introduce a comprehensive framework for analyzing scientific publications using deep learning techniques, demonstrating significant improvements in accuracy and efficiency.',
    'Our research explores the application of transformer models to academic text processing, with particular emphasis on semantic understanding and knowledge extraction.',
    'This study investigates the use of neural networks for pattern recognition in citation networks, providing insights into research collaboration and knowledge flow.',
    'We propose a multi-modal approach to research document understanding, combining text analysis with visual element recognition for comprehensive academic literature processing.'
  ]
  
  return abstracts[Math.floor(Math.random() * abstracts.length)]
}

function generateMockJournal(): string {
  return journals[Math.floor(Math.random() * journals.length)]
}

function generateMockLibraryName(): string {
  return libraryNames[Math.floor(Math.random() * libraryNames.length)]
}

function generateMockDescription(): string {
  const descriptions = [
    'A curated collection of cutting-edge research papers in the field.',
    'Essential readings for understanding current developments and trends.',
    'Comprehensive resource for researchers and practitioners.',
    'Selected papers covering both theoretical foundations and practical applications.',
    'Current research papers addressing key challenges and innovations.'
  ]
  
  return descriptions[Math.floor(Math.random() * descriptions.length)]
}

function generateMockNotes(): string {
  const notes = [
    'Interesting methodology - could be applicable to our current project.',
    'Key insights on page 15 regarding the experimental setup.',
    'Important findings that contradict previous research in this area.',
    'Potential collaboration opportunity with the authors.',
    'Novel approach that could inspire our next research direction.'
  ]
  
  return notes[Math.floor(Math.random() * notes.length)]
}

function getRandomStatus(): ReadingStatus {
  const statuses = [
    ReadingStatus.READING_STATUS_TO_READ,
    ReadingStatus.READING_STATUS_READING,
    ReadingStatus.READING_STATUS_READ,
    ReadingStatus.READING_STATUS_ABANDONED
  ]
  
  // Weight towards completed and reading statuses
  const weights = [0.3, 0.3, 0.35, 0.05]
  const random = Math.random()
  let cumulative = 0
  
  for (let i = 0; i < weights.length; i++) {
    cumulative += weights[i]
    if (random <= cumulative) {
      return statuses[i]
    }
  }
  
  return ReadingStatus.READING_STATUS_TO_READ
}
