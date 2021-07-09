import request from './index'

export function GetSiteInfo() {
    return request({
        url: '/data/siteinfo.json',
        method: "get",
    })
}

export function GetSiteDetail() {
    return request({
        url: '/data/sitedetail.json',
        method: "get",
    })
}