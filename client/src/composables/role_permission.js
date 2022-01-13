import golangApiService from "@/core/http/golang.api.service";
import {reactive} from "vue";

export function roles() {
    let compData = reactive({roles: [], role: {}, permissions:{}})

    const getRoles = async () => {
        await golangApiService.get('/api/pr/roles')
            .then(data => Object.assign(compData.roles, data.data.roles))

    }

    const getOneRole = async (id) => {
        await golangApiService.get(`/api/pr/role/${id}`).then(data => Object.assign(compData.role, data.data.role))
    }

    const disableRole = async (id) => {
        await golangApiService.delete(`/api/pr/role/${id}`)
    }
    const enableRole = async (id) => {
        await golangApiService.get(`/api/pr/enable/${id}`)
    }

    const saveChanges = async (data) => {
        const index = compData.roles.map(element => element.id).indexOf(data.id)
        compData.roles[index] = data
        await golangApiService.post('/api/pr/save', {role: data})

    }

    const getPermissions = async () => {
        await golangApiService.get('/api/pr/permissions').then(data => Object.assign(compData.permissions, data.data.permissions))
    }

    return {
        getOneRole,
        getRoles,
        disableRole,
        enableRole,
        saveChanges,
        getPermissions,
        compData
    }
}
