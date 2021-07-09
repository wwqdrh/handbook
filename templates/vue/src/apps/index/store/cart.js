import { GetSiteInfo, GetSiteDetail } from "@/api/cart"

export default {
    namespaced: true,
    state: {
        activeID: null,

        siteInfo: null,
        siteDetail: null,
    },
    getters: {
        siteInfo: state => {
            if (!state.siteInfo) {
                try {
                    state.siteInfo = JSON.parse(sessionStorage.getItem("HUI_SITE_INFO"))
                } catch {
                    return
                }
            }
            return state.siteInfo
        },
        siteDetail: state => {
            if (!state.siteDetail) {
                try {
                    state.siteDetail = JSON.parse(sessionStorage.getItem("HUI_SITE_DETAIL"))
                } catch {
                    return
                }
            }
            return state.siteDetail
        },
        activeDetail: state => {
            if (state.activeID === null) {
                return null
            }
            for (let item of state.siteDetail) {
                if (item.id === state.activeID) {
                    return item
                }
            }
            return null
        }
    },
    mutations: {
        initSiteInfo(state, data) {
            try {
                state.siteInfo = data
                sessionStorage.setItem("HUI_SITE_INFO", JSON.stringify(data))
            } catch {
                return
            }
        },
        initSiteDetail(state, data) {
            try {
                state.siteDetail = data
                sessionStorage.setItem("HUI_SITE_DETAIL", JSON.stringify(data))
            } catch {
                return
            }
        }
    },
    actions: {
        initSiteInfo(ctx) {
            // 通过网络获取，主要是为了降低一次性打包资源，但是这些数据还是通过文件导入避免别人直接使用
            return new Promise((resolve, reject) => {
                // 通过网络获取，主要是为了降低一次性打包资源，但是这些数据还是通过文件导入避免别人直接使用
                // GetSiteInfo()
                // .then(data => {
                //     ctx.commit("initSiteInfo", data.data)
                //     resolve(data)
                // })
                // .catch(err => {
                //     reject(err)
                // })
                try {
                    let data = require("@/configs/siteinfo.js").DATA
                    ctx.commit("initSiteInfo", data)
                    resolve(data)
                } catch (error) {
                    reject(error)
                }
            })

        },
        initSiteDetail(ctx) {
            return new Promise((resolve, reject) => {
                // GetSiteDetail()
                //     .then(data => {
                //         ctx.commit("initSiteDetail", data.data)
                //         resolve(data)
                //     })
                //     .catch(err => {
                //         reject(err)
                //     })
                try {
                    let data = require("@/configs/sitedetail.js").DATA
                    ctx.commit("initSiteDetail", data)
                    resolve(data)
                } catch (error) {
                    reject(error)
                }
            })
        },
        activeID(ctx, data) {
            ctx.state.activeID = data
        }
    },
}