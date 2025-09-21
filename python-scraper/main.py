import asyncio
import csv
from time import sleep

from playwright.async_api import async_playwright
from init_db import init_db, save_news

SITE_URL = "https://radio.gov.pk"

async def main():
    news_list = []
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
            news_list.append({
                "slug": slug,
                "headline": headline,
                "published_on": published_on,
                "news_url": SITE_URL + news_link,
                "image_url": image_url,
            })
        await browser.close()
    for news in news_list:
        await save_news(news)
        print(f"Slug: {news['slug']} Saved!")
    # with open(csv_file, mode="w", encoding="utf-8") as file:
    #     writer = csv.DictWriter(file, fieldnames=["headline","news_link","image_url"])
    #     writer.writeheader()
    #     writer.writerows(news_list)
    # print(f"Data saved to {csv_file}")

if __name__ == '__main__':
    init_db()
    asyncio.run(main())

