import asyncio
from time import sleep
from typing import List, TypedDict

from playwright.async_api import async_playwright
from init_db import init_db, save_news

SITE_URL = "https://radio.gov.pk"

class News(TypedDict):
    slug: str
    headline: str | None
    published_on: str
    news_url: str
    image_url: str | None

async def main():
    news_list : List[News] = []
    async with async_playwright() as p:
        browser = await p.chromium.launch(headless=False)
        page = await browser.new_page()
        print(f"Going to {SITE_URL}...")
        await page.goto(url=SITE_URL, wait_until="domcontentloaded", timeout=60000)
        print("Getting news...")
        sleep(1)
        print("Locating elements...")
        link_elements = await page.locator("a:has([class*='title'])").all()
        print(f"Located {len(link_elements)} news")
        for el in link_elements:
            news_link = await el.get_attribute("href")
            image_el = el.locator("img")
            image_doesnt_exist = await image_el.count() == 0
            headline = await el.locator("[class*='title']").text_content()
            image_url = None if image_doesnt_exist else "https:"+ await image_el.get_attribute("src")
            published_on, slug = news_link.split("/")[-2:]
            news_item:News = {
                "slug": slug,
                "headline": headline,
                "published_on": published_on,
                "news_url": SITE_URL + news_link,
                "image_url": image_url,
            }
            news_list.append(news_item)
        await browser.close()
    for news in news_list:
        try:
            await save_news(news)
        except Exception as error:
            print(f"Error in slug: {news['slug']}, {error}")

if __name__ == '__main__':
    init_db()
    asyncio.run(main())

