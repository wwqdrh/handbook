export default {
    namespaced: true,
    state: {
        tabs: [], // 需要保存当前打开的tab信息{url: ..., icon: ...}
        showTabs: [],
        status: new Map(), // 保存当前的tab状态，show、offset等
    },
    getters: {
        tabs: state => {
            if (!state.tabs) {
                try {
                    state.tabs = JSON.parse(sessionStorage.getItem("HUI_TAB_ITEMS"))
                } catch (error) {
                    return state.tabs
                }
            }
            return state.tabs
        },
        showTabs: state => {
            let res = []
            for (let i in state.tabs) {
                let item = state.tabs[i]
                if (state.status.get(item.url)) {
                    res.push(item)
                }
            }
            return res
        },
        isShow: state => (key) => {
            return state.status[key]
        }
    },
    mutations: {
        addItem(state, tabs) {
            state.tabs = tabs
            sessionStorage.setItem("HUI_TAB_ITEMS", JSON.stringify(tabs))
        }
    },
    actions: {
        addItem(ctx, data) {
            let origin_data = ctx.getters['tabs']
            ctx.state.status.set(data['url'], true)
            origin_data.push({ url: data['url'], icon: data['icon'] })
            console.log(ctx.getters['showTabs'])
            ctx.commit("addItem", origin_data)
        },
        minItem(ctx, key) {
            ctx.state.status.set(key, false)
            console.log(ctx.state.status)
        },
        maxItem(ctx, key) {
            ctx.state.status.set(key, true)
        },
        removeItem(ctx, key) {
            let origin_data = ctx.getters['tabs']
            let new_data = []
            ctx.state.status.remove(key)
            for (let i in origin_data) {
                let item = origin_data[i]
                if (item.url == key) {
                    continue
                }
                new_data.push(item)
            }
            ctx.commit("addItem", new_data)
        }
    },
}