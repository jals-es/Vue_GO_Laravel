// import {createStore} from 'vuex'
import laravelApiService from '../core/http/laravel.api.service'
export const superUser = {
        namespaced: true,
        state: {
            user: {},
            bars: [],
            bar_info: {},
            orders: [],
            stats: [],
            graph_one: {},
            graph_two: {},
            graph_three: {},
            graph_four: {},
            graph_five: {},
            graph_six: {}

        },
        mutations: {
            addBar(state, bar) {
                state.bars.push(bar)
            },
            addOrder(state, order) {
                state.order.push(order)
            },
            getStats(state, payload) {
                state.stats = payload
            },
            getUser(state, payload) {
                state.user = payload
            },
            getFirstChart(state, payload) {
                state.graph_one = payload
            },
            getSecondChart(state, payload) {
                state.graph_two = payload
            },
            getBars(state,payload) {
                state.bars = payload
            },
            getBarInfo(state,payload) {
                state.bar_info = payload
            },
            getBarStats(state,payload) {
                state.bar_info.stats = payload
                // console.log({...state.bar_info});
            }

            
        },
        actions: {
            getUser(store) {
                laravelApiService.get('/api/userData')
                .then(({ data }) => {
                    console.log(data);
                    store.commit("getUser", data.message[0]);
                    // state.user=data.message[0];
                  })
                  .catch((error) => {
                    console.log("ERROR: userData");
                    console.log(error);
                  });
            },
            getStats(store) {
                laravelApiService.get('/api/bars/stats')
                .then(({ data }) => {
                    console.log(data);
                    store.commit("getStats", data.data);
                    // state.user=data.message[0];
                  })
                  .catch((error) => {
                    console.log("ERROR: barsStats");
                    console.log(error);
                  });
            },
            getFirstChart(store) {
                laravelApiService.get('/api/charts/first')
                .then(({ data }) => {
                    console.log(data);
                    store.commit("getFirstChart", data.data);
                  })
                  .catch((error) => {
                    console.log("ERROR: chartsFirst");
                    console.log(error);
                  });
            },
            getSecondChart(store) {
                laravelApiService.get('/api/charts/second')
                .then(({ data }) => {
                    console.log(data);
                    store.commit("getSecondChart", data.data);
                  })
                  .catch((error) => {
                    console.log("ERROR: chartsSecond");
                    console.log(error);
                  });
            },
            getBars(store,name = "%") {
                laravelApiService.get('/api/bars/search/'+name)
                .then(({ data }) => {
                    console.log(data);
                    store.commit("getBars", data.data);
                  })
                  .catch((error) => {
                    console.log("ERROR: getBars");
                    console.log(error);
                  });
            },
            getBarInfo(store,slug) {
                laravelApiService.get('/api/bars/info/'+slug)
                .then(({ data }) => {
                    console.log(data.data[0]);
                    store.commit("getBarInfo", data.data[0]);
                  })
                  .catch((error) => {
                    console.log("ERROR: getBarInfo");
                    console.log(error);
                  });
                laravelApiService.get('/api/bars/stats/'+slug)
                .then(({ data }) => {
                    console.log(data.data);
                    store.commit("getBarStats", data.data);
                  })
                  .catch((error) => {
                    console.log("ERROR: getBarStats");
                    console.log(error);
                  });
            }
        },
        getters: {
            getUser(state) {
                return state.user;
            },
            getStats(state) {
                return state.stats;
            },
            getFirstChart(state) {
                return state.graph_one;
            },
            getSecondChart(state) {
                return state.graph_two;
            },
            getBars(state) {
                return state.bars;
            },
            getBarInfo(state) {
                return state.bar_info;
            }


            // getBarFromId: state => id => state.bars.find(bar => bar.id = id),
            // getOrderFromId: state => id => state.order.find(order => order.id = id),
        }
    }
