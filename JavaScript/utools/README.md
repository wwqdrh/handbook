plugin.json example

```js
window.exports = {
    "a name": { // 注意：键对应的是 plugin.json 中的 features.code
        mode: "none",  // 用于无需 UI 显示，执行一些简单的代码
        args: {
            // 进入插件时调用
            enter: (action) => {
                // action = { code, type, payload }
                ExecCommand()
                
                window.utools.hideMainWindow()
                // do some thing
                try {
                    addframe(
                        utools.ubrowser.goto(`https://www.baidu.com`)
                    ).run({ width: 1500, height: 900 })
                } catch (error) {
                    alert(error)
                }
                // do some thing
                // window.utools.outPlugin()
            }
        }
    },
    "do something": {
        mode: 'list',
        args: {
            enter: (action, callbackSetList) => {
                search = new SearchAction()
            },
            search: (action, searchWord, callbackSetList) => {
                // alert(Object.keys(action))
                // alert(Object.keys(action) + "|" + action.code + "|" + action.type + "|" + action.playload)
                if (!searchWord) return callbackSetList()
                searchWord = searchWord.toLowerCase()
                if (searchWord.endsWith("$")) {
                    search.searchTitle(searchWord.substr(0, searchWord.length - 1)).then(res => {
                        if(res.length === 0) {
                            return callbackSetList({title: "无数据"})
                        } else {
                            return callbackSetList(res)
                        }
                    })
                } else if (searchWord.endsWith("$-2")) {
                    search.nextPage().then(res => {
                        res.unshift({title: "下一页"})
                        return callbackSetList(res)
                    })
                } else {
                    return callbackSetList()
                }
            },
            select: (action, itemData) => {
                // window.utools.hideMainWindow()
                // alert(search.domain + itemData.href)
                utools.ubrowser.goto(search.domain + itemData.href).run({ width: 900, height: 600 })
                // window.utools.outPlugin()
            }
        }
    }
}
```