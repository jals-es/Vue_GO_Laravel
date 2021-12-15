import golangApiService from "@/core/http/golang.api.service";

export const userStore = {
    namespaced: true,
    state: {
        user: {}
    },
    mutations: {
        fillUser (state, data){
            delete data.token
            state.user = data
        }
    },
    getters: {
        getSuperAdmin: state => {
            return state.user.superadmin
        }
    },
    actions: {
        registerUser(state, data) {
            const user = {users: {}}
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
        }, loginUser(state, data) {
            const user = {users: {email: data.email, passwd: data.passwd}}
            return new Promise(((resolve, reject) => {
                golangApiService.put('/api/user/', user)
                    .then(result => {
                        resolve(result)
                        localStorage.setItem('token', result.data.user.token)

                    })
                    .catch(error => reject(error))
            }))
        }
    }
}
