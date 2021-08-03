import Vue from 'vue'
import "tailwindcss/tailwind.css"
import "animate.css/animate.min.css"

import App from './App.vue'
import router from './router'

import '@/assets/icons'

Vue.config.productionTip = false

new Vue({
    router,
    render: h => h(App)
}).$mount('#app')