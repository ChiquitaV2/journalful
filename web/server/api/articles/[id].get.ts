import {useServices} from "~~/server/proto/useServices";
import {GetArticleRequest} from "~~/server/proto/grpc/articles/v1/article";
import {generateMockArticle} from "~~/server/api/mock/data";

export default defineEventHandler(async event => {
  return generateMockArticle(1, {
    title: "Machine Learning Approaches to Academic Literature Analysis",
    authors: [
      { id: 1, name: "Dr. Sarah Chen", profileId: 1 },
      { id: 2, name: "Prof. Michael Rodriguez", profileId: 2 }
    ],
    abstract: "This comprehensive study explores the application of machine learning techniques to analyze patterns in academic literature. We present novel approaches for citation analysis, topic modeling, and research trend prediction that significantly improve upon existing methodologies.",
    url: "https://example.com/article",
    pdfUrl: "https://example.com/article.pdf",
    publishedDate: "2024-03-15",
    journalName: "Journal of Academic Computing",
    tags: ["Machine Learning", "Literature Analysis", "NLP", "Academia"],
    readingStatus: "in-progress",
    readingProgress: 65,
    dateStarted: "2024-07-20",
    dateAdded: "2024-07-15",
    isFavorite: true,
    notes: [
      {
        id: "1",
        content: "Interesting approach to citation network analysis. The methodology seems robust and could be applicable to other domains.",
        createdAt: "2024-07-22"
      },
      {
        id: "2",
        content: "Key insight: The temporal analysis of research trends shows cyclical patterns that weren't previously identified.",
        createdAt: "2024-07-25"
      }
    ]
  })

  const articlesSvr = await useServices().getArticlesServiceClient(event)
  const id = Number(event.context.params?.id)
  if (!id) createError({
    statusCode: 400,
    statusMessage: 'ID is required'
  })
  return await articlesSvr.GetArticle(GetArticleRequest.fromJSON({id}))

})
