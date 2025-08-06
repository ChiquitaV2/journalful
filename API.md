# Journalful API Design

This document outlines the design of the Journalful gRPC API. The API is defined using Protocol Buffers, and the service definitions can be found in the `api` directory.

## Services

The API is divided into four main services:

*   **ArticlesService:** Manages academic articles.
*   **LibraryService:** Manages user libraries and saved articles.
*   **AuthorService:** Manages author information.
*   **ProfileService:** Manages user profiles.

---

## ArticlesService

Service for managing academic articles.

### RPCs

#### `GetArticle`

Retrieves a single article by its ID.

*   **Request:** `GetArticleRequest`
*   **Response:** `GetArticleResponse`

#### `GetArticleByDOI`

Retrieves a single article by its DOI.

*   **Request:** `GetArticleByDOIRequest`
*   **Response:** `GetArticleByDOIResponse`

#### `ListArticles`

Lists all articles, with optional pagination.

*   **Request:** `ListArticlesRequest`
*   **Response:** `ListArticlesResponse`

#### `CreateArticle`

Creates a new article.

*   **Request:** `CreateArticleRequest`
*   **Response:** `CreateArticleResponse`

#### `UpdateArticle`

Updates an existing article.

*   **Request:** `UpdateArticleRequest`
*   **Response:** `UpdateArticleResponse`

#### `DeleteArticle`

Deletes an article.

*   **Request:** `DeleteArticleRequest`
*   **Response:** `DeleteArticleResponse`

### Messages

*   `Article`: Represents an academic article.
*   `GetArticleRequest`: Request to get an article by ID.
*   `GetArticleResponse`: Response containing the article.
*   `GetArticleByDOIRequest`: Request to get an article by DOI.
*   `GetArticleByDOIResponse`: Response containing the article.
*   `ListArticlesRequest`: Request to list articles.
*   `ListArticlesResponse`: Response containing a list of articles.
*   `CreateArticleRequest`: Request to create an article.
*   `CreateArticleResponse`: Response containing the ID of the new article.
*   `UpdateArticleRequest`: Request to update an article.
*   `UpdateArticleResponse`: Response indicating the success of the update.
*   `DeleteArticleRequest`: Request to delete an article.
*   `DeleteArticleResponse`: Response indicating the success of the deletion.

---

## LibraryService

Service for managing user libraries.

### RPCs

#### `SaveArticleToLibrary`

Saves an article to a user's library.

*   **Request:** `SaveArticleToLibraryRequest`
*   **Response:** `SaveArticleToLibraryResponse`

#### `GetUserLibrary`

Retrieves a user's library.

*   **Request:** `GetUserLibraryRequest`
*   **Response:** `GetUserLibraryResponse`

### Messages

*   `Library`: Represents a user's library.
*   `SavedArticle`: Represents an article saved in a library.
*   `ReadingStatus`: Enum for the reading status of an article.
*   `SaveArticleToLibraryRequest`: Request to save an article to a library.
*   `SaveArticleToLibraryResponse`: Response indicating the success of the save.
*   `GetUserLibraryRequest`: Request to get a user's library.
*   `GetUserLibraryResponse`: Response containing the user's library.

---

## AuthorService

Service for managing authors.

### RPCs

#### `GetAuthor`

Retrieves a single author by their ID.

*   **Request:** `GetAuthorRequest`
*   **Response:** `GetAuthorResponse`

#### `GetAuthorByProfileID`

Retrieves an author by their profile ID.

*   **Request:** `GetAuthorByProfileIDRequest`
*   **Response:** `GetAuthorByProfileIDResponse`

#### `ListAuthors`

Lists all authors.

*   **Request:** `ListAuthorsRequest`
*   **Response:** `ListAuthorsResponse`

#### `CreateAuthor`

Creates a new author.

*   **Request:** `CreateAuthorRequest`
*   **Response:** `CreateAuthorResponse`

#### `UpdateAuthor`

Updates an existing author.

*   **Request:** `UpdateAuthorRequest`
*   **Response:** `UpdateAuthorResponse`

#### `DeleteAuthor`

Deletes an author.

*   **Request:** `DeleteAuthorRequest`
*   **Response:** `DeleteAuthorResponse`

### Messages

*   `Author`: Represents an author.
*   `GetAuthorRequest`: Request to get an author by ID.
*   `GetAuthorResponse`: Response containing the author.
*   `GetAuthorByProfileIDRequest`: Request to get an author by profile ID.
*   `GetAuthorByProfileIDResponse`: Response containing the author.
*   `ListAuthorsRequest`: Request to list authors.
*   `ListAuthorsResponse`: Response containing a list of authors.
*   `CreateAuthorRequest`: Request to create an author.
*   `CreateAuthorResponse`: Response containing the ID of the new author.
*   `UpdateAuthorRequest`: Request to update an author.
*   `UpdateAuthorResponse`: Response indicating the success of the update.
*   `DeleteAuthorRequest`: Request to delete an author.
*   `DeleteAuthorResponse`: Response indicating the success of the deletion.

---

## ProfileService

Service for managing user profiles.

### RPCs

#### `GetProfile`

Retrieves a user profile.

*   **Request:** `GetProfileRequest`
*   **Response:** `GetProfileResponse`

### Messages

*   `Profile`: Represents a user profile.
*   `GetProfileRequest`: Request to get a user profile.
*   `GetProfileResponse`: Response containing the user profile.
*   `ListProfilesRequest`: Request to list profiles.
*   `ListProfilesResponse`: Response containing a list of profiles.
*   `CreateProfileRequest`: Request to create a profile.
*   `CreateProfileResponse`: Response containing the ID of the new profile.
*   `UpdateProfileRequest`: Request to update a profile.
*   `UpdateProfileResponse`: Response indicating the success of the update.
*   `DeleteProfileRequest`: Request to delete a profile.
*   `DeleteProfileResponse`: Response indicating the success of the deletion.
