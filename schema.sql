-- Profiles of users/researchers
CREATE TABLE profiles
(
    id          BIGINT AUTO_INCREMENT PRIMARY KEY,
    name        VARCHAR(100) NOT NULL,
    bio         TEXT,
    institution VARCHAR(100),
    created_at  TIMESTAMP DEFAULT CURRENT_TIMESTAMP,
    updated_at  TIMESTAMP DEFAULT CURRENT_TIMESTAMP ON UPDATE CURRENT_TIMESTAMP
);

-- Authors (can be linked to a profile, but don't have to be)
CREATE TABLE authors
(
    id         BIGINT AUTO_INCREMENT PRIMARY KEY,
    name       VARCHAR(100) NOT NULL,
    -- An author might have a profile, or might be an external author not in the system.
    profile_id BIGINT       NULL,                                                                     -- Changed to BIGINT to match profiles.id, and NULLABLE
    created_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP,
    updated_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP ON UPDATE CURRENT_TIMESTAMP,
    CONSTRAINT fk_author_profile FOREIGN KEY (profile_id) REFERENCES profiles (id) ON DELETE SET NULL -- If profile deleted, unlink author
);

-- Academic articles/papers
CREATE TABLE articles
(
    id               BIGINT AUTO_INCREMENT PRIMARY KEY,
    doi              VARCHAR(100) NOT NULL,
    title            VARCHAR(255) NOT NULL,
    abstract         TEXT,
    publication_year INT,               -- Added for common metadata
    journal_name     VARCHAR(255),      -- Added for common metadata
    created_at       TIMESTAMP DEFAULT CURRENT_TIMESTAMP,
    updated_at       TIMESTAMP DEFAULT CURRENT_TIMESTAMP ON UPDATE CURRENT_TIMESTAMP,
    UNIQUE INDEX idx_articles_doi (doi) -- DOI should be unique
);

-- Junction table for many-to-many relationship between articles and authors
CREATE TABLE article_authors
(
    article_id   BIGINT NOT NULL,
    author_id    BIGINT NOT NULL,
    author_order INT       DEFAULT 0, -- To maintain the order of authors on a paper
    created_at   TIMESTAMP DEFAULT CURRENT_TIMESTAMP,
    PRIMARY KEY (article_id, author_id),
    CONSTRAINT fk_articleauthors_article FOREIGN KEY (article_id) REFERENCES articles (id) ON DELETE CASCADE,
    CONSTRAINT fk_articleauthors_author FOREIGN KEY (author_id) REFERENCES authors (id) ON DELETE CASCADE
);

-- User's personal library
CREATE TABLE library
(
    id         BIGINT AUTO_INCREMENT PRIMARY KEY,
    user_id    BIGINT NOT NULL,                   -- This should reference profiles.id
    name       VARCHAR(255) DEFAULT 'My Library', -- Optional: Allow users to name libraries if they can have multiple
    created_at TIMESTAMP    DEFAULT CURRENT_TIMESTAMP,
    updated_at TIMESTAMP    DEFAULT CURRENT_TIMESTAMP ON UPDATE CURRENT_TIMESTAMP,
    CONSTRAINT fk_library_user FOREIGN KEY (user_id) REFERENCES profiles (id) ON DELETE CASCADE
);

-- Junction table linking articles to a user's library, with reading status
CREATE TABLE library_articles
(
    id             BIGINT AUTO_INCREMENT PRIMARY KEY,                -- Optional, composite PK (library_id, article_id) is also fine
    library_id     BIGINT NOT NULL,
    article_id     BIGINT NOT NULL,
    -- 0: Unspecified, 1: To Read, 2: Reading, 3: Read, 4: Abandoned
    reading_status TINYINT   DEFAULT 0 COMMENT '0:Unspecified, 1:ToRead, 2:Reading, 3:Read, 4:Abandoned',
    added_at       TIMESTAMP DEFAULT CURRENT_TIMESTAMP,              -- Renamed from created_at for clarity
    notes          TEXT,                                             -- User's personal notes on this article in their library
    created_at     TIMESTAMP DEFAULT CURRENT_TIMESTAMP,
    updated_at     TIMESTAMP DEFAULT CURRENT_TIMESTAMP ON UPDATE CURRENT_TIMESTAMP,
    CONSTRAINT fk_libraryarticles_library FOREIGN KEY (library_id) REFERENCES library (id) ON DELETE CASCADE,
    CONSTRAINT fk_libraryarticles_article FOREIGN KEY (article_id) REFERENCES articles (id) ON DELETE CASCADE,
    UNIQUE INDEX idx_library_article_unique (library_id, article_id) -- Prevent adding same article multiple times to same library
);


-- Indexes for performance
CREATE INDEX idx_authors_name ON authors (name);
CREATE INDEX idx_articles_title ON articles (title);
CREATE INDEX idx_library_user_id ON library (user_id);
CREATE INDEX idx_library_articles_reading_status ON library_articles (reading_status);