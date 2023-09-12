// 执行api
// const http = require('http');
const https = require('https');
const HttpsProxyAgent = require('https-proxy-agent')
const HttpProxyAgent = require('http-proxy-agent')
const qs = require("url");
const { DOMParser, XMLSerializer } = require('xmldom')
const xpath = require('xpath');
const { resolve } = require('path');
function sleep(ms) {
    return new Promise(resolve => setTimeout(resolve, ms))
}

// use your conf
const url = ""
const domain = ""
const proxy = ""

class Search {
    constructor() {
        this.url = url // 首页
        this.domain = domain
        this.proxy = proxy;
        this.agent = new HttpsProxyAgent({
            host: "127.0.0.1",
            port: "7890",
            secureProxy: false,
        });
        this.cur = 1; // 当前页面
        // this.idx = 0; // TODO 当前页面中li的搜索位置
        this.datalist = []; // 当前页面中所有的数据元素 {title, url}
        this.keyword = "";
        this.maxpage = 10;
        this.curPageRes = [] // 当前搜索情况的结果
    }
    // searchTitle 根据keyword搜索出当前的内容
    // @页数
    // 1、搜出10个就返回，不够就继续往后搜索
    // 2、每次最多搜索20页，如果20页都没搜满10个也会返回
    // 3、继续往后搜索需要触发时间
    async searchTitle(keyword) {
        this.curPageRes = []
        this.cur = 0
        // this.idx = 0
        this.keyword = keyword
        const data = await this.doSearch()
        return data
    }
    doSearch() {
        let this_ = this
        let page = this.maxpage
        let tasks = []
        for (let i = 0; i < page; i++) {
            var url
            if (this_.cur + i === 1) {
                url = url
            } else {
                url = url // do page + 1
            }
            tasks.push((async (url) => {
                let times = 5;
                while (times > 0) {
                    let data = await this_.getPageData(url)
                    if (data != "") {
                        let datalist = this_.parseTitleList(data) // 更新list
                        // 搜索结果
                        this_.titlePattern(datalist)
                        break
                    }
                    times--
                    await sleep(100)
                }
            })(url))
        }
        return new Promise((resolve, reject) => {
            Promise.all(tasks).then(_ => {
                this_.cur += 10
                resolve(this_.curPageRes)
            }).catch(err => {
                console.log(err)
                reject(err)
            });
        })
    }
    // titlePattern 将当前有的和没有的进行匹配
    titlePattern(datalist) {
        for (let item of datalist) {
            if (item.title.indexOf(this.keyword) !== -1) {
                this.curPageRes.push(item)
            }
        }
    }
    // nextPage 往后翻页的操作
    async nextPage() {
        // alert("下一页")
        this.curPageRes = []
        this.cur++
        // this.idx = 0
        return await this.doSearch(this.keyword)
    }
    // TODO prevPage 往前翻页的操作
    // 是需要记录之前的元素的值
    async prevPage() {
        return this.curPageRes
    }
    // [async]getPageData 获取页面的html值
    getPageData(url) {
        return new Promise((resolve, reject) => {
            let ps = new URL(url)
            let https_options = {
                "host": ps.host,
                "path": ps.pathname + ps.search,
                // "port": 443,
                "agent": this.agent,
                // "timeout": 10000,
            };
            https.get(https_options, function (res) {
                // 分段返回的 自己拼接
                let html = '';
                // 有数据产生的时候 拼接
                res.on('data', function (chunk) {
                    html += chunk;
                })
                // 拼接完成
                res.on('end', function () {
                    resolve(html)
                })
                res.on('error', function (e) {
                    console.log(e)
                    // reject(e)
                })
            }).on('error', function (err) {
                //错误处理，处理res无法处理到的错误
                resolve("")
            })
        })
    }
    // parseTitleList 根据page内容解析出页面中的li元素
    parseTitleList(page) {
        let res = []
        // 使用xpath定位
        let doc = new DOMParser().parseFromString(page)
        // console.log(new XMLSerializer().serializeToString(doc))
        let nodes = xpath.select("/html/body/table[2]/tr/td/ul/li/a", doc)
        for (let node of nodes) {
            let title = xpath.select("./text()", node)[0]
            let href = xpath.select("./@href", node)[0]
            res.push({
                title: title.nodeValue,
                href: href.nodeValue
            })
        }
        return res
    }
}

module.exports = {
    SearchAction: Search
}

const s = new Search()

