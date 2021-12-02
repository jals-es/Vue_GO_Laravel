import { createStore } from 'vuex'

export default createStore({
    state: {},
    mutations: {},
    actions: {},
    modules: {}
})

export const getStore = (name) => {
    if (!name) return
    return JSON.parse(window.localStorage.getItem(name))
}