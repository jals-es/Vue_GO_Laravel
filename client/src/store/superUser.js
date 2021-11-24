import {createStore} from 'vuex'

let user = createStore({
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
            }
        },
        actions: {},
        getters: {
            getUserid: state => state.user.id,
            getBarFromId: state => id => state.bars.find(bar => bar.id = id),
            getOrderFromId: state => id => state.order.find(order => order.id = id),
        }
    }
)

export default user