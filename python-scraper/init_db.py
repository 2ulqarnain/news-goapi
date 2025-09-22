import sqlite3

DB_PATH = "../data/news.db"

def get_connection_to_db():
    return sqlite3.connect(DB_PATH)

def init_db():
    conn = get_connection_to_db()
    cur = conn.cursor()

    cur.execute("""
        CREATE TABLE IF NOT EXISTS news (
            slug TEXT NOT NULL UNIQUE PRIMARY KEY,
            headline TEXT NOT NULL,
            published_on TEXT NOT NULL,
            news_url TEXT NOT NULL,
            image_url TEXT
        )
    """)
    conn.commit()
    conn.close()

async def save_news(news):
    conn = get_connection_to_db()
    cur = conn.cursor()

    (
        slug,
        headline,
        published_on,
        news_url,
        image_url
     ) = news['slug'], news['headline'], news['published_on'], news['news_url'], news['image_url']

    cur.execute("""
        INSERT INTO news (slug, headline, published_on, news_url, image_url) VALUES (?, ?, ?, ?, ?)
    """, (slug, headline, published_on, news_url, image_url))

    conn.commit()
    conn.close()