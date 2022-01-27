import laravelApiService from "@/core/http/laravel.api.service";
import { reactive } from 'vue';
export function incidences() {

    const incidences = reactive({ incidences: [] })
    const getIncidences = async() => {
        let getIncidences = await laravelApiService.get('/api/incidence')
        if (getIncidences) {
            console.log(getIncidences.data.incidences)
            console.log(getIncidences.data)

            Object.assign(incidences.incidences, getIncidences.data.incidences)
        }
    }

    return {
        getIncidences,
        incidences
    }
}