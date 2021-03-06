import golangApiService from "@/core/http/golang.api.service";

import { httpClient as laravelHttpClient } from "@/core/http/laravel.api.service";
import { httpClient as golangHttpClient } from "@/core/http/golang.api.service";

export const userStore = {
    namespaced: true,
    state: {
        user: {},
        checkToken: {}
    },
    mutations: {
        fillUser(state, data) {
            delete data.token
            state.user = data
        }
    },
    actions: {
        registerUser(state, data) {
            const user = { users: {} }
            Object.keys(data).forEach((key) => {
                if (key !== 'passwdCheck') {
                    user.users[key] = data[key]
                }
            })
            return new Promise(((resolve, reject) => {
                golangApiService.post('/api/user/', user)
                    .then(result => resolve(result))
                    .catch(error => reject(error))
            }))
        },
        loginUser(state, data) {
            const user = { users: { email: data.email, passwd: data.passwd } }
            return new Promise(((resolve, reject) => {
                golangApiService.put('/api/user/', user)
                    .then(result => {
                        resolve(result)
                        if (result.data.user.superadmin) {
                            localStorage.setItem('jwt', result.data.user.token);
                            laravelHttpClient.defaults.headers.common['Authorization'] = `Bearer ` + localStorage.getItem("jwt");
                        } else {
                            localStorage.setItem('token', result.data.user.token);
                            golangHttpClient.defaults.headers.common['Authorization'] = "Token " + localStorage.getItem("token");
                        }
                    })
                    .catch(error => reject(error))
            }))
        },
        checkToken(store) {
            console.log('Entra a la consulta');

            golangApiService.get('/api/user/')
                .then(({ data }) => {
                    console.log(data);
                    store.commit("checkToken", data);
                })
                .catch((error) => {
                    console.log(error)
                })
        }
    },
    getters: {
        getSuperAdmin: state => {
            return state.user.superadmin
        },
        checkToken(state) {
            return state.checkToken;
        }
    }
}