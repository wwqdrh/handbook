"""
使用playwright进行爬取

快发卡的登录 添加商品 添加卡密 获取店铺链接
"""
import asyncio
import csv
import os.path
import json
import contextlib
from pathlib import Path
import dataclasses
from typing import Generator, List

from playwright.async_api import async_playwright


@dataclasses.dataclass
class Product:
    id: str
    name: str
    value: str
    product_description: str
    meta: str = ""  # 商品种类
    tag: str = ""  # 商品标签
    cover: str = ""
    spec: List[str] = dataclasses.field(default_factory=list)  # 商品说明图
    secret_key: str = ""
    secret_value: str = ""
    mall_link: str = ""  # 店铺链接
    product_type: str = "编程教程"  # 商品类型，现在只有一个编程教程
    secret_type: int = 2  # 重复卡
    secret_show: int = 0  # 卡密默认展示
    use_message: str = "更多资源请查看小店官网~\n\n" + "售后请添加QQ 1343901798"


class ProductEncoder:
    class _jsonEncoder(json.JSONEncoder):
        def default(self, obj):
            if isinstance(obj, Product):
                return {
                    "id": obj.id,
                    "meta": obj.meta,
                    "tag": obj.tag,
                    "cover": obj.cover,
                    "spec": obj.spec,
                    "name": obj.name,
                    "value": obj.value,
                    "description": obj.product_description,
                    "link": obj.mall_link,
                }
            # Base class default() raises TypeError:
            return json.JSONEncoder.default(self, obj)

    @classmethod
    def csv_input(cls, csv_path: Path) -> Generator[Product, None, None]:
        """从给定的csv路径获取信息"""
        if not os.path.exists(csv_path):
            return None

        with open(csv_path, encoding="gbk", newline="") as csvfile:
            reader = csv.DictReader(csvfile)
            for row in reader:
                secret_key, secret_value = row["secret"].split("---")
                yield Product(
                    id=row["id"],
                    name=row["name"],
                    value=row["value"],
                    product_description=row["description"],
                    meta=row["meta"],  # 商品种类
                    tag=row["tag"],  # 商品标签
                    cover=row["cover"],
                    spec=row["spec"].split(";"),  # 商品说明图
                    secret_key=secret_key,
                    secret_value=secret_value,
                )

    @classmethod
    def json_output(cls, products: List[Product], json_path: Path):
        """将product写在json文件中"""
        with open(json_path, mode="w", encoding="utf8") as f:
            f.write(
                json.dumps(products, cls=cls._jsonEncoder, indent=4, ensure_ascii=False)
            )


class BrowserContext:
    @contextlib.asynccontextmanager
    async def browser_context(self, headless: bool = True):
        async with async_playwright() as p:
            try:
                browser = await p.chromium.launch(headless=headless, slow_mo=100)
                yield browser
            except Exception as e:
                print("未知错误: ", e)
            await browser.close()


class KuaifakaContext(BrowserContext):
    """
    https://www.kuaifaka.com/
    """

    async def login_handler(self, page):
        await page.goto("https://www.kuaifaka.com/login")
        await page.fill('//*[@id="login"]/div[2]/div/div/div[2]/input', "hhhui")
        await page.fill('//*[@id="login"]/div[2]/div/div/div[3]/input', "991028DRHkfk")
        await page.click('//*[@id="login"]/div[2]/div/div/div[5]')

    async def add_product_handler(self, page, product: Product):
        await page.click('//*[@id="backstage"]/div[1]/div[1]/div[1]/div/div[3]/div[4]')
        # TODO: 商品分类暂时不管
        await page.click('//*[@id="add_goods"]/div[2]/div[2]/div/div/div[1]/div/span')
        await page.fill('//*[@id="add_goods"]/div[3]/div[2]/input', product.name)
        await page.fill('//*[@id="add_goods"]/div[4]/div[2]/input', product.value)
        await page.click(
            f'//*[@id="add_goods"]/div[5]/div[2]/div[{product.secret_type+1}]'
        )
        await page.fill(
            '//*[@id="add_goods"]/div[7]/div[2]/textarea', product.product_description
        )
        await page.fill(
            '//*[@id="add_goods"]/div[8]/div[2]/div/div[1]/div', product.use_message
        )
        await page.click('//*[@id="add_goods"]/div[12]/div')  # 保存提交

    async def add_secret_handler(self, page, product: Product):
        await page.click('//*[@id="backstage"]/div[1]/div[1]/div[1]/div/div[3]/div[1]')
        await asyncio.sleep(0.5)
        await page.click(
            '//*[@id="backstage"]/div[1]/div[1]/div[1]/div/div[3]/div[6]'
        )  # 添加卡密页面
        await page.click('//*[@id="add_card"]/div[3]/div[1]/div')
        await page.click(
            f'//*[@id="add_card"]/div[3]/div[1]/div/div[2]/ul[2]/li >> text={product.product_type}'
        )

        await asyncio.sleep(1)
        await page.click('//*[@id="add_card"]/div[3]/div[2]/div')
        await page.click(
            f"#add_card > div.add_card > div:nth-child(2) > div > div.ivu-select-dropdown > ul.ivu-select-dropdown-list > li >> text={product.name}"
        )
        await asyncio.sleep(1)
        await page.click("#add_card > div.input_cards > textarea")
        await page.fill(
            "#add_card > div.input_cards > textarea",
            f"{product.secret_key} {product.secret_value}",
        )  # 添加卡密 使用空格分隔
        await asyncio.sleep(1)
        await page.click('//*[@id="set_btn"]')  # 提交

    async def get_product_link_handler(self, page, product: Product) -> bool:
        await page.click('//*[@id="backstage"]/div[1]/div[1]/div[1]/div/div[3]/div[1]')
        await asyncio.sleep(0.5)
        await page.click(
            '//*[@id="backstage"]/div[1]/div[1]/div[1]/div/div[3]/div[5]'
        )  # 商品列表页

        max_num = 6  # 一页只有6条
        next_button = page.locator('//*[@id="goods"]/div[2]/ul/li[3]')
        while True:
            for i in range(max_num):
                if (
                    await page.locator(
                        f'//*[@id="goods"]/div[2]/table/tbody/tr[{i*2+1}]/td[3]'
                    ).inner_text()
                    != product.name
                ):
                    continue

                await page.click(
                    f'//*[@id="goods"]/div[2]/table/tbody/tr[{i*2+1}]/td[10]/span'
                )
                await page.click(
                    f'//*[@id="goods"]/div[2]/table/tbody/tr[{i*2+2}]/td/div[2]'
                )
                await asyncio.sleep(1)
                product.mall_link = (
                    await page.locator(
                        '//*[@id="pay_url"]/section[1]/div[1]/div[1]/span[2]'
                    ).inner_text()
                ).lstrip("(推荐)：")
                return True
            if await next_button.is_disabled() is False:
                await next_button.click()
            else:
                break

        return False


async def kuaifaka_addproduct_link():
    kuaifaka = KuaifakaContext()
    async with kuaifaka.browser_context(False) as browser:
        page = await browser.new_page()
        await kuaifaka.login_handler(page)

        res = []
        for product in ProductEncoder.csv_input(Path(__file__).parents[0] / "mall.csv"):
            res.append(product)
            # await kuaifaka.add_product_handler(page, product)
            # print("添加商品成功")
            # await asyncio.sleep(2)
            await kuaifaka.add_secret_handler(page, product)
            print("添加秘钥成功")
            await asyncio.sleep(2)
            await kuaifaka.get_product_link_handler(page, product)
            print("获取商品链接成功")
            await asyncio.sleep(1)

        ProductEncoder.json_output(res, Path(__file__).parents[0] / "result.json")
        print("执行完成")


async def main():
    """
    需要从csv -> dataclass -> json
    """
    await kuaifaka_addproduct_link()


if __name__ == "__main__":
    if len(argv := __import__("sys").argv) > 1 and argv[1] == "test":
        pass
        # res = []
        # for i in product_csv_input(Path(__file__).parents[0] / "mall.csv"):
        #     res.append(i)
        # product_json_output(res, Path(__file__).parents[0] / "result.json")
    else:
        asyncio.run(main())
