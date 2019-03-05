import Vue from "vue";
import Vuex from "vuex";
import stores from '@/service/store'

Vue.use(Vuex);

export default new Vuex.Store({
    modules: {...stores}
});
