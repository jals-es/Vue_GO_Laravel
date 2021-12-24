// import {createStore} from 'vuex'
import laravelApiService from '../core/http/laravel.api.service'
export const superUser = {
        namespaced: true,
        state: {
            user: {},
            bars: [],
            bar_info: {},
            orders: [],
            products: [],
            stats: [],
            graph_one: {},
            graph_two: {},
            incidences: [],

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
            },
            getOrders(state,payload) {
                state.orders = payload
            },
            getProducts(state,payload) {
                state.products = payload
            },
            getIncidences(state,payload) {
                state.incidences = payload
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
                    const first_data = data.data[0]
                    laravelApiService.get('/api/bars/stats/'+slug)
                    .then(({ data }) => {
                        first_data.stats = data.data
                        store.commit("getBarInfo", first_data);
                      })
                      .catch((error) => {
                        console.log("ERROR: getBarStats");
                        console.log(error);
                      });
                  })
                  .catch((error) => {
                    console.log("ERROR: getBarInfo");
                    console.log(error);
                  });
            },
            getOrders(store,id_bar = '"%"') {
                laravelApiService.get('/api/bars/orders/'+'"'+id_bar.id_bar+'"')
                .then(({ data }) => {
                    store.commit("getOrders", data.data);
                  })
                  .catch((error) => {
                    console.log("ERROR: getOrders");
                    console.log(error);
                  });
            },
            getProducts(store, id_bar = '%') {
                laravelApiService.get('/api/bars/products/'+id_bar.idBar)
                .then(({ data }) => {
                    console.log(data);
                    store.commit("getProducts", data.data);
                  })
                  .catch((error) => {
                    console.log("ERROR: getProducts");
                    console.log(error);
                  });
            },
            getIncidences(store) {
                laravelApiService.get('/api/incidence')
                .then(({ data }) => {
                    console.log(data);
                    store.commit("getIncidences", data.data);
                  })
                  .catch((error) => {
                    console.log("ERROR: getIncidences");
                    console.log(error);
                  });
            },
            createIncidence(store, formData){
                laravelApiService.post('api/incidence', formData)
                .then(({ data }) => {
                    console.log(data);
                })
                .catch((error) => {
                    console.log("ERROR: createIncidence");
                    console.log(error);
                  });
            },
            updateIncidence(store, formData){
                laravelApiService.patch('api/incidence', formData)
                .then(({ data }) => {
                    console.log(data);
                })
                .catch((error) => {
                    console.log("ERROR: createIncidence");
                    console.log(error);
                  });
            },
            deleteIncidence(store, formData){
                laravelApiService.post('api/incidence', formData)
                .then(({ data }) => {
                    console.log(data);
                })
                .catch((error) => {
                    console.log("ERROR: createIncidence");
                    console.log(error);
                  });
            },

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
            },
            getOrders(state) {
                return state.orders;
            },
            getProducts(state) {
                return state.products;
            },
            getIncidences(state) {
                return state.incidences;
            }
        }
    }
