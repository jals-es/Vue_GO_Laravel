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
            }
        },
        getters: {
            getUser(state) {
                return state.user;
            }
            // getBarFromId: state => id => state.bars.find(bar => bar.id = id),
            // getOrderFromId: state => id => state.order.find(order => order.id = id),
        }
    }
