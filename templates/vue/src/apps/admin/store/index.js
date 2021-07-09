import Vue from "vue"
import Vuex from "vuex"

import cart from "./cart"
import tab from "./tab"

import { ValidatorPassword } from "@/utils/sec"

Vue.use(Vuex)

export default new Vuex.Store({
    modules: {
        cart,
        tab
    },
    state: {
        tabOpen: false, // tab窗口的状态
        tabUrl: "", // tab浏览器的url地址
        searchTab: false, // search的状态

        secSIG: null,
        secPWD: null,
    },
    getters: {
        tabOpen: state => state.tabOpen,
        tabUrl: state => state.tabUrl,
        searchTab: state => state.searchTab,

        secSIG: state => {
            if (!state.secSIG) {
                state.secSIG = sessionStorage.getItem("HUI_SEC_SIG")
            }
            return state.secSIG
        },
        secPWD: state => {
            if (!state.secPWD) {
                state.secPWD = sessionStorage.getItem("HUI_SEC_P")
            }
            return state.secPWD
        },
    },
    mutations: {
        OpenTab(state, url) {
            state.tabOpen = true
            state.tabUrl = url
        },
        CloseTab(state) {
            state.tabOpen = false
            state.tabUrl = ""
        },
        OpenSearch(state) {
            state.searchTab = true
        },
        CloseSearch(state) {
            state.searchTab = false
        },

        setSIG(state, data) {
            state.secSIG = data[0]
            state.secPWD = data[1]
            sessionStorage.setItem("HUI_SEC_SIG", data[0])
            sessionStorage.setItem("HUI_SEC_P", data[1])
        }
    },
    actions: {
        OpenTab(ctx, url) {
            ctx.commit("OpenTab", url)
        },
        CloseTab(ctx) {
            ctx.commit("CloseTab")
        },
        OpenSearch(ctx) {
            ctx.commit("OpenSearch")
        },
        CloseSearch(ctx) {
            ctx.commit("CloseSearch")
        },

        setSIG(ctx, data) {
            let res = ValidatorPassword(data)
            if (res !== false) {
                ctx.commit("setSIG", res)
            } else {
                throw Error("密码错误")
            }
        },
    }
})