import aiosqlite

DB_PATH = "../data/news.db"

async def get_connection_to_db():
    return await aiosqlite.connect(DB_PATH)

async def init_db():
    async with aiosqlite.connect(DB_PATH) as conn:
        await conn.execute("""
            CREATE TABLE IF NOT EXISTS news (
                slug TEXT NOT NULL UNIQUE PRIMARY KEY,
                headline TEXT NOT NULL,
                published_on TEXT NOT NULL,
                news_url TEXT NOT NULL,
                image_url TEXT
            )
        """)
        await conn.commit()

async def save_news(news):
    async with aiosqlite.connect(DB_PATH) as conn:
        await conn.execute("""
            INSERT INTO news (slug, headline, published_on, news_url, image_url)
            VALUES (?, ?, ?, ?, ?)
        """, (
            news['slug'],
            news['headline'],
            news['published_on'],
            news['news_url'],
            news['image_url']
        ))
        await conn.commit()
