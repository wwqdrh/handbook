// https://www.kubernetes.org.cn/
// k8s中文社区
const https = require('https');
const xpath = require('xpath');
const { DOMParser, XMLSerializer } = require('xmldom')


const URL = "https://www.kubernetes.org.cn/"
const NAME = "k8s中文社区"
const DESCRIPTION = "kubernetes的中文社区，提供一些文章"


module.exports = {
    GetNewArticle: GetNewArticle,
    GetArticleContent: GetArticleContent
}

// GetNewArticle 获取当前最新文章
async function GetNewArticle() {
    const url = URL
    const xarticle = "/html/body/section/div[1]/div/article/header/h2/a"
    const body = await new Promise((resolve, reject) => {
        https.get(url, function (res) {
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
    });

    if (body === "") {
        return []
    }

    // 解析article
    let res = []
    let doc = new DOMParser().parseFromString(body)
    let nodes = xpath.select(xarticle, doc)
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

// GetArticleContent 获取文章内容
async function GetArticleContent(url) {
    // const xheader = "string(/html/body/section/div/div/header)"
    // const xcontent = "string(/html/body/section/div/div/article)"
    const xheader = "/html/body/section/div/div/header"
    const xcontent = "/html/body/section/div/div/article"

    const body  = await new Promise((resolve, reject) => {
        https.get(url, function (res) {
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
    if (body === "") {
        return {}
    }

    let doc = new DOMParser().parseFromString(body)
    let header = xpath.select(xheader, doc)[0]
    let content = xpath.select(xcontent, doc)[0]
    return {
        header: header.toString(),
        content: content.toString(),
    }
}