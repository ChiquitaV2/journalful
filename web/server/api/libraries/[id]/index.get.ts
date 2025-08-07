import { generateMockLibrary } from '../../mock/data'

export default defineEventHandler(event => {
  const id = Number(event.context.params?.id)
  
  // Generate mock library with consistent data for demo
  return generateMockLibrary(id, {
    name: "Machine Learning Research",
    description: "A curated collection of papers on machine learning, deep learning, and AI research.",
    icon: "🤖",
    color: "linear-gradient(135deg, #8b5cf6, #3b82f6)",
    createdAt: "2024-01-15",
    updatedAt: "2024-07-28",
    articles: [
      {
        id: "1",
        title: "Attention Is All You Need",
        authors: [{ name: "Vaswani et al." }],
        abstract: "The dominant sequence transduction models are based on complex recurrent or convolutional neural networks...",
        url: "https://arxiv.org/abs/1706.03762",
        publishedDate: "2017-06-12",
        journal: "NeurIPS",
        tags: ["Transformers", "Attention", "NLP"],
        readingStatus: "completed",
        readingProgress: 100,
        dateAdded: "2024-01-20",
        dateStarted: "2024-01-22",
        dateCompleted: "2024-01-25",
        isFavorite: true
      },
      {
        id: "2",
        title: "BERT: Pre-training of Deep Bidirectional Transformers",
        authors: [{ name: "Devlin et al." }],
        abstract: "We introduce a new language representation model called BERT...",
        url: "https://arxiv.org/abs/1810.04805",
        publishedDate: "2018-10-11",
        journal: "NAACL",
        tags: ["BERT", "Language Models", "NLP"],
        readingStatus: "in-progress",
        readingProgress: 45,
        dateAdded: "2024-02-01",
        dateStarted: "2024-02-05",
        isFavorite: false
      },
      {
        id: "3",
        title: "GPT-3: Language Models are Few-Shot Learners",
        authors: [{ name: "Brown et al." }],
        abstract: "Recent work has demonstrated substantial gains on many NLP tasks...",
        url: "https://arxiv.org/abs/2005.14165",
        publishedDate: "2020-05-28",
        journal: "NeurIPS",
        tags: ["GPT", "Language Models", "Few-Shot Learning"],
        readingStatus: "unread",
        readingProgress: 0,
        dateAdded: "2024-03-10",
        isFavorite: false
      },
      {
        id: "4",
        title: "ResNet: Deep Residual Learning for Image Recognition",
        authors: [{ name: "He et al." }],
        abstract: "Deeper neural networks are more difficult to train...",
        url: "https://arxiv.org/abs/1512.03385",
        publishedDate: "2015-12-10",
        journal: "CVPR",
        tags: ["ResNet", "Computer Vision", "CNN"],
        readingStatus: "completed",
        readingProgress: 100,
        dateAdded: "2024-01-30",
        dateCompleted: "2024-02-15",
        isFavorite: true
      }
    ]
  })
})
