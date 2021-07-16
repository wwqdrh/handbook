import requests
import sys

print(sys.version)
resp = requests.get("https://www.baidu.com")
print(resp.content)
