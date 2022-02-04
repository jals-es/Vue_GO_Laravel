import golangApiService from "@/core/http/golang.api.service";
import router from "../router";

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
                    store.commit("getBars", data);
                })
                .catch((error) => {
                    console.log("ERROR: getBars");
                    console.log(error);
                });
        },
        createBar(store, data) {
            const bar = { bar: { name: data.name, descr: data.descr, lat: String(data.lat), lon: String(data.lon), city: data.city, address: data.address } }
            golangApiService.post('/api/bar/', bar)
                .then(({ data }) => {
                    if (data.message) {
                        if (data.message === "Bar creado correctamente") {
                            router.push("/bars");
                        }
                    }
                })
                .catch((error) => {
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