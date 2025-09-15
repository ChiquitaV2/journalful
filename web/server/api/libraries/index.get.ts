import {generateMockLibrary} from '../mock/data'
import {useServices} from "~~/server/proto/useServices";
import {GetUserLibraryRequest, ReadingStatus} from "~~/server/proto/grpc/library/v1/library";

export default defineEventHandler(async event => {

  const librarySvr = await useServices().getLibraryServiceClient(event)
  try {
    const libriaires = await librarySvr.GetUserLibrary(GetUserLibraryRequest.fromJSON({userId: 1}))
    console.log("libriaires", libriaires)
    const defaultLibrary = libriaires.defaultLibrary
    const privateLibraries = libriaires.privateLibraries
    if (!defaultLibrary) {
      return createError({
        statusCode: 404,
        statusMessage: "Default library not found"
      })
    }
    return {
      defaultLibrary,
      privateLibraries,
      stats: {
        total: 1 + privateLibraries.length,
        articles: defaultLibrary.articles.length + privateLibraries.reduce((sum, lib) => sum + lib.articles.length, 0),
        completed: [defaultLibrary, ...privateLibraries]
            .flatMap(lib => lib.articles)
            .filter(article => article.readingStatus === ReadingStatus.READING_STATUS_READ).length
      }
    }
  } catch (e: any) {
    console.error("lookup error", e)
    // For all errors besdies not found return the error
    if (!e.message.includes("not found")) {
      return createError({
        statusCode: 500,
        statusMessage: e.message
      })
    }
  }
  // Generate mock libraries for the index page
  const defaultLibrary = generateMockLibrary(1, {
    name: 'My Reading List',
    description: 'Your default collection for articles',
    articleCount: 42,
    isDefault: true
  })

  const privateLibraries = [
    generateMockLibrary(2, {
      name: 'Machine Learning Papers',
      description: 'Latest research in ML and AI',
      articleCount: 18,
      icon: 'heroicons:cpu-chip'
    }),
    generateMockLibrary(3, {
      name: 'Climate Research', 
      description: 'Environmental science and climate change studies',
      articleCount: 12,
      icon: 'heroicons:globe-alt'
    }),
    generateMockLibrary(4, {
      name: 'Computer Vision',
      description: 'Image processing and visual recognition research',
      articleCount: 8,
      icon: 'heroicons:camera'
    })
  ]

  return {
    defaultLibrary,
    privateLibraries,
    stats: {
      total: 1 + privateLibraries.length,
      articles: defaultLibrary.articles.length + privateLibraries.reduce((sum, lib) => sum + lib.articles.length, 0),
      completed: [defaultLibrary, ...privateLibraries]
        .flatMap(lib => lib.articles)
        .filter(article => article.readingStatus === 'completed').length
    }
  }
})
