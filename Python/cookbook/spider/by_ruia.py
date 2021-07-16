"""
看了下，不太行

入口是Spider的spider，会需要传递一个loop或者新建一个loop，无法关闭日志
"""
import aiofiles

from ruia import Spider, Item, TextField, AttrField


class HackerNewsItem(Item):
    target_item = TextField(css_select="tr.athing")
    title = TextField(css_select="a.storylink")
    url = AttrField(css_select="a.storylink", attr="href")


class HackerNewsSpider(Spider):
    start_urls = [f"https://news.ycombinator.com/news?p={index}" for index in range(3)]

    async def parse(self, response):
        async for item in HackerNewsItem.get_items(html=response.html):
            yield item

    async def process_item(self, item: HackerNewsItem):
        """Ruia build-in method"""
        async with aiofiles.open("./hacker_news.txt", "a") as f:
            await f.write(str(item.title) + "\n")


if __name__ == "__main__":
    HackerNewsSpider.start()