import {useServices} from "~~/server/proto/useServices";
import {SaveArticleToLibraryRequest} from "~~/server/proto/grpc/library/v1/library";

export default defineEventHandler(async (event) => {
   const body = await readBody(event)
   const libraryId = Number(event.context.params?.id)
    if (!libraryId || body.articleId === undefined) {
        createError({
            statusCode: 400,
            statusMessage: 'Library ID and article ID are required'
        })
    }
    const librarySvr = await useServices().getLibraryServiceClient(event)
    return librarySvr.SaveArticleToLibrary(SaveArticleToLibraryRequest.fromJSON({
        libraryId,
        articleId: body.articleId,
        readingStatus: body.readingStatus ? body.readingStatus : ReadingStatus.READING_STATUS_TO_READ
        })
    )
});