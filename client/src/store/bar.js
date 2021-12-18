import golangApiService from "@/core/http/golang.api.service";

export const barStore = {
    namespaced: true,
    state: {
        bars: [],
    },
    mutations: {
        getBars(state, payload) {
            state.bars = payload
        }
    },
    actions: {
        getBars(store) {
            golangApiService.get('/api/bar/')
                .then(({ data }) => {
                    console.log(data);
                    store.commit("getBars", data);
                })
                .catch((error) => {
                    console.log("ERROR: getBars");
                    console.log(error);
                });
        }

    },
    getters: {
        getBars(state) {
            return state.bars;
        }
    }
}