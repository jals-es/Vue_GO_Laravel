// import {createStore} from 'vuex'
import laravelApiService from '../core/http/laravel.api.service'
export const superUser = {
        namespaced: true,
        state: {
            user: {},
            bars: [],
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
                console.log(payload)
                if (payload) {
                    state.user = payload;
                } else {
                    state.user = {
                        name: "",
                        email: "",
                        photo: ""
                    };
                }
            },
            getFirstChart(state, payload) {
                state.graph_one = payload
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
            }



            // getBarFromId: state => id => state.bars.find(bar => bar.id = id),
            // getOrderFromId: state => id => state.order.find(order => order.id = id),
        }
    }
