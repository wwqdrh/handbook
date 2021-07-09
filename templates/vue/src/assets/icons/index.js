import Vue from 'vue'
import SvgIcon from './icon.vue'// svg component

// register globally
Vue.component('svg-icon', SvgIcon)

/**
 * require.context
 * 
 * @param directory {String} -读取文件的路径
 * @param useSubdirectories {Boolean} -是否遍历文件的子目录
 * @param regExp {RegExp} -匹配文件的正则
 * @return {function} 在方法属性上注册了有keys方法，resolve方法，id属性
 */
const req = require.context('./svg', false, /\.svg$/)
// console.log(req)
const requireAll = requireContext => requireContext.keys().map(requireContext)//应该可以理解为加载在内存中
requireAll(req)
