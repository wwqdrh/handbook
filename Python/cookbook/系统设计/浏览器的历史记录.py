"""
解释：
BrowserHistory browserHistory = new BrowserHistory("leetcode.com");
browserHistory.visit("google.com"); // 你原本在浏览 "leetcode.com" 。访问 "google.com"
browserHistory.visit("facebook.com"); // 你原本在浏览 "google.com" 。访问 "facebook.com"
browserHistory.visit("youtube.com"); // 你原本在浏览 "facebook.com" 。访问 "youtube.com"
browserHistory.back(1); // 你原本在浏览 "youtube.com" ，后退到 "facebook.com" 并返回 "facebook.com"
browserHistory.back(1); // 你原本在浏览 "facebook.com" ，后退到 "google.com" 并返回 "google.com"
browserHistory.forward(1); // 你原本在浏览 "google.com" ，前进到 "facebook.com" 并返回 "facebook.com"
browserHistory.visit("linkedin.com"); // 你原本在浏览 "facebook.com" 。 访问 "linkedin.com"
browserHistory.forward(2); // 你原本在浏览 "linkedin.com" ，你无法前进任何步数。
browserHistory.back(2); // 你原本在浏览 "linkedin.com" ，后退两步依次先到 "facebook.com" ，然后到 "google.com" ，并返回 "google.com"
browserHistory.back(7); // 你原本在浏览 "google.com"， 你只能后退一步到 "leetcode.com" ，并返回 "leetcode.com"
"""
"""
使用双向链表表示浏览记录
"""

from dataclasses import dataclass


@dataclass
class ListNode:
    val: str
    pre: "ListNode" = None
    back: "ListNode" = None


class BroswerHistory:

    def __init__(self, host: str):
        self.root = ListNode(host)    # 根url
        self.cur = self.root    # 当前url

    def visit(self, url: str):
        node = ListNode(url, back=self.cur)
        self.cur.pre = node
        self.cur = node
        return url

    def forward(self, step: int) -> str:
        cur = self.cur
        while cur.pre and step > 0:
            cur = cur.pre
            step -= 1
        self.cur = cur
        return cur.val

    def back(self, step: int) -> str:
        cur = self.cur
        while cur.back and step > 0:
            cur = cur.back
            step -= 1
        self.cur = cur
        return cur.val


if __name__ == "__main__":
    browser = BroswerHistory("http://leetcode.com")
    print(browser.visit("google.com"))
    print(browser.visit("baidu.com"))
    print(browser.forward(1))
    print(browser.back(1))
    print(browser.visit("4399.com"))
    print("----------")