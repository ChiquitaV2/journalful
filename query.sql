-- Profiles of users/researchers

-- name: GetProfile :one
SELECT * FROM profiles WHERE id = ? LIMIT 1;

-- name: ListProfiles :many
SELECT * FROM profiles ORDER BY name;

-- name: CreateProfile :execresult
INSERT INTO profiles (name, bio, institution) VALUES (?, ?, ?);

-- name: UpdateProfile :exec
UPDATE profiles SET name = ?, bio = ?, institution = ?, updated_at = CURRENT_TIMESTAMP WHERE id = ?;

-- name: DeleteProfile :exec
DELETE FROM profiles WHERE id = ?;


-- Authors

-- name: GetAuthor :one
SELECT * FROM authors WHERE id = ? LIMIT 1;

-- name: GetAuthorByName :one
SELECT * FROM authors WHERE name = ? LIMIT 1;

-- name: GetAuthorByProfileID :one
SELECT * FROM authors WHERE profile_id = ? LIMIT 1;

-- name: ListAuthors :many
SELECT * FROM authors ORDER BY name;

-- name: CreateAuthor :execresult
INSERT INTO authors (name, profile_id) VALUES (?, ?);

-- name: UpdateAuthor :exec
UPDATE authors SET name = ?, profile_id = ?, updated_at = CURRENT_TIMESTAMP WHERE id = ?;

-- name: DeleteAuthor :exec
DELETE FROM authors WHERE id = ?;



-- Academic articles/papers

-- name: GetArticle :one
SELECT * FROM articles WHERE id = ? LIMIT 1;

-- name: GetArticleByDOI :one
SELECT * FROM articles WHERE doi = ? LIMIT 1;

-- name: ListArticles :many
SELECT * FROM articles ORDER BY title;

-- name: CreateArticle :execresult
INSERT INTO articles (doi, title, abstract, url, publication_year, journal_name) VALUES (?, ?, ?, ?, ?, ?);

-- name: UpdateArticle :exec
UPDATE articles SET doi = ?, title = ?, abstract = ?, url = ?, publication_year = ?, journal_name = ?, updated_at = CURRENT_TIMESTAMP WHERE id = ?;

-- name: DeleteArticle :exec
DELETE FROM articles WHERE id = ?;


-- Junction table for many-to-many relationship between articles and authors (article_authors)

-- name: AddArticleAuthor :execresult
INSERT INTO article_authors (article_id, author_id, author_order) VALUES (?, ?, ?);

-- name: ListArticleAuthorsByArticleID :many
SELECT
    aa.author_id,
    aa.author_order,
    a.name AS author_name,
    a.profile_id
FROM article_authors aa
         JOIN authors a ON aa.author_id = a.id
WHERE aa.article_id = ?
ORDER BY aa.author_order;

-- name: ListArticleAuthorsByAuthorID :many
SELECT
    aa.article_id,
    aa.author_order,
    ar.title AS article_title,
    ar.doi
FROM article_authors aa
         JOIN articles ar ON aa.article_id = ar.id
WHERE aa.author_id = ?
ORDER BY ar.publication_year DESC, ar.title;

-- name: DeleteArticleAuthor :exec
DELETE FROM article_authors WHERE article_id = ? AND author_id = ?;


-- User's personal library

-- name: GetLibrary :one
SELECT * FROM library WHERE id = ? LIMIT 1;

-- name: ListLibrariesByUserID :many
SELECT * FROM library WHERE owner_id = ? ORDER BY created_at;

-- name: CreateLibrary :execresult
INSERT INTO library (owner_id, name, description, isPublic, isDefault) VALUES (?, ?, ?, ?, ?);

-- name: UpdateLibrary :exec
UPDATE library SET name = ?, description = ?, isPublic = ?, isDefault = ?, updated_at = CURRENT_TIMESTAMP WHERE id = ?;

-- name: DeleteLibrary :exec
DELETE FROM library WHERE id = ?;


-- Junction table linking articles to a user's library, with reading status (library_articles)

-- name: GetLibraryArticle :one
SELECT * FROM library_articles WHERE library_id = ? AND article_id = ? LIMIT 1;

-- name: AddLibraryArticle :execresult
INSERT INTO library_articles (library_id, article_id, reading_status, reading_progress, dateAdded, dateCompleted, notes, isFavorite) VALUES (?, ?, ?, ?, ?, ?, ?, ?);

-- name: ListLibraryArticlesByLibraryID :many
SELECT
    la.id,
    la.article_id,
    la.reading_status,
    la.dateAdded,
    la.notes,
    a.title AS article_title,
    a.doi,
    a.publication_year
FROM library_articles la
         JOIN articles a ON la.article_id = a.id
WHERE la.library_id = ?
ORDER BY la.dateAdded DESC;

-- name: UpdateLibraryArticleStatus :exec
UPDATE library_articles SET reading_status = ?, updated_at = CURRENT_TIMESTAMP WHERE id = ?;

-- name: UpdateLibraryArticleNotes :exec
UPDATE library_articles SET notes = ?, updated_at = CURRENT_TIMESTAMP WHERE id = ?;

-- name: DeleteLibraryArticle :exec
DELETE FROM library_articles WHERE id = ?;

-- name: UpdateSavedArticle :exec
UPDATE library_articles SET reading_status = ?, notes = ?, updated_at = CURRENT_TIMESTAMP WHERE id = ?;

-- name: DeleteSavedArticle :exec
DELETE FROM library_articles WHERE id = ?;
