CREATE TABLE news (
    slug TEXT UNIQUE NOT NULL PRIMARY KEY,
    title TEXT NOT NULL,
    published_on TEXT NOT NULL,
    image_url TEXT,
    news_url TEXT NOT NULL,
    content TEXT
)