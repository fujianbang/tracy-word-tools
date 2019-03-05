import axios from './axios'
import api from './api'

export default {
    install: (Vue, options) => {
        Vue.prototype.$api = api
        Vue.prototype.$http = axios
    }
}
