'use strict'
const path = require('path')
const glob = require('glob')
const webpack = require("webpack")

function resolve(dir) {
    return path.join(__dirname, dir)
}

/***
 * 1、pages/[name]/main.js =》 name.html
 * 2、pages/templates/main.js => ......html(里面所有的模板按照其路径生成对应的html文件)
 */
function getEntry(url) {
    let entrys = {}
    glob.sync(url).forEach(item => {
        // splice(-3)取数组后三项
        let urlArr = item.split('/').splice(-3)
        if (urlArr[1] === "templates") {
            // 处理templates中的html文件
            let path = item.substring(0, item.length - 7)
            glob.sync(`${path}**/*.html`).forEach(template_item => {
                let file_path = template_item.replace(path, "")
                entrys[file_path] = {
                    entry: item,
                    template: template_item,
                    filename: 'templates/' + file_path,
                    title: 'pages-' + file_path
                }
            })
        }
        else {
            entrys[urlArr[1]] = {
                entry: './src/pages/' + urlArr[1] + '/' + 'main.js',
                template: './src/pages/' + urlArr[1] + '/index.html',
                filename: 'templates/' + urlArr[1] + '.html',
                title: 'pages-' + urlArr[1]
            }
        }
    })
    return entrys
}

module.exports = {
    publicPath: '/',
    pages: getEntry('./src/pages/*/main.js'),
    outputDir: 'dist',
    assetsDir: 'static',
    lintOnSave: process.env.NODE_ENV === 'development',
    productionSourceMap: false,
    devServer: {
        port: 8888, // 运行的端口
        open: true,  // 设置为true，当应用启动后会打开默认的浏览器
        overlay: {  // 当出现编译异常或者错误的时候是否全屏显示的默认行为
            warnings: false,
            errors: true
        },
        // before: require('./mock-server'),  // 提供所有服务器中间件之前的行为
        // proxy: {
        //     '/api': {
        //         target: 'http://localhost',
        //         changeOrigin: true,
        //         pathRewrite: {
        //             '^/api': ''
        //         }
        //     }
        // }
    },
    configureWebpack: {
        // provide the app's title in webpack's name field, so that
        // it can be accessed in index.html to inject the correct title.
        resolve: {
            alias: {
                '@': resolve('src')
            }
        },
    },
    chainWebpack(config) {
        // when there are many pages, it will cause too many meaningless requests
        config.plugins.delete('prefetch')

        // 避免基础svgloder加载
        config.module
            .rule('svg')
            .exclude.add(resolve('src/assets/icons'))
            .end()

        // 自定义svgloader，在svg-sprite-loader基础上
        config.module
            .rule('icons')
            .test(/\.svg$/)
            .include.add(resolve('src/assets/icons'))
            .end()
            .use('svg-sprite-loader')
            .loader('svg-sprite-loader')
            .options({
                symbolId: 'icon-[name]'
            })
            .end()
        // 修改preload配置项
        // config.plugin('preload').tap(() => [
        //     {
        //         rel: 'preload',
        //         // to ignore runtime.js
        //         // https://github.com/vuejs/vue-cli/blob/dev/packages/@vue/cli-service/lib/config/app.js#L171
        //         fileBlacklist: [/\.map$/, /hot-update\.js$/, /runtime\..*\.js$/],
        //         include: 'initial'
        //     }
        // ])

        // // 删除prefetch配置项，when there are many pages, it will cause too many meaningless requests
        // config.plugins.delete('prefetch')

        // // set svg-sprite-loader
        // config.module
        //     .rule('svg')
        //     .exclude.add(resolve('src/icons'))
        //     .end()
        // config.module
        //     .rule('icons')
        //     .test(/\.svg$/)
        //     .include.add(resolve('src/icons'))
        //     .end()
        //     .use('svg-sprite-loader')
        //     .loader('svg-sprite-loader')
        //     .options({
        //         symbolId: 'icon-[name]'
        //     })
        //     .end()

        // config
        //     .when(process.env.NODE_ENV !== 'development',
        //         config => {
        //             config
        //                 .plugin('ScriptExtHtmlWebpackPlugin')
        //                 .after('html')
        //                 .use('script-ext-html-webpack-plugin', [{
        //                     // `runtime` must same as runtimeChunk name. default is `runtime`
        //                     inline: /runtime\..*\.js$/
        //                 }])
        //                 .end()
        //             config
        //                 .optimization.splitChunks({
        //                     chunks: 'all',
        //                     cacheGroups: {
        //                         libs: {
        //                             name: 'chunk-libs',
        //                             test: /[\\/]node_modules[\\/]/,
        //                             priority: 10,
        //                             chunks: 'initial' // only package third parties that are initially dependent
        //                         },
        //                         elementUI: {
        //                             name: 'chunk-elementUI', // split elementUI into a single package
        //                             priority: 20, // the weight needs to be larger than libs and app or it will be packaged into libs or app
        //                             test: /[\\/]node_modules[\\/]_?element-ui(.*)/ // in order to adapt to cnpm
        //                         },
        //                         commons: {
        //                             name: 'chunk-commons',
        //                             test: resolve('src/components'), // can customize your rules
        //                             minChunks: 3, //  minimum common number
        //                             priority: 5,
        //                             reuseExistingChunk: true
        //                         }
        //                     }
        //                 })
        //             // https:// webpack.js.org/configuration/optimization/#optimizationruntimechunk
        //             config.optimization.runtimeChunk('single')
        //         }
        //     )
    }
}
