<template>
    <form class="my-5 mx-auto p-3" @submit.prevent="submit">
        <h2 class="mb-3">Crear Bar</h2>
        <h5 class="mt-3 mb-2">Datos:</h5>
        <div class="form-floating">
            <input v-model="state.bar.name" type="text" class="form-control my-1" id="floatingName" placeholder="Nombre">
            <label for="floatingName">Nombre</label>
        </div>
        <div class="form-floating">
            <input v-model="state.bar.descr" type="text" class="form-control my-1" id="floatingDescr" placeholder="Descripcion">
            <label for="floatingDescr">Descripción</label>
        </div>
        <div class="form-floating">
            <input v-model="state.bar.address" type="text" class="form-control my-1" id="floatingAddrs" placeholder="Direccion">
            <label for="floatingAddrs">Dirección</label>
        </div>
        <div class="form-floating">
            <input v-model="state.bar.city" type="text" class="form-control my-1" id="floatingCity" placeholder="Ciudad">
            <label for="floatingCity">Ciudad</label>
        </div>
        <h5 class="mt-3 mb-2">Coordenadas:</h5>
        <div class="form-floating">
            <input v-model="state.bar.lat" type="number" class="form-control my-1" id="floatingLat" placeholder="Latitud">
            <label for="floatingLat">Latitud</label>
        </div>
        <div class="form-floating">
            <input v-model="state.bar.lon" type="number" class="form-control my-1" id="floatingLon" placeholder="Longitud">
            <label for="floatingLon">Longitud</label>
        </div>
        <div class="mt-4">
            <router-link to="/bars" class="w-25 btn btn-lg btn-danger">Cancelar</router-link>
            <button class="w-75 btn btn-lg btn-primary" type="submit">Crear</button>
        </div>
    </form>
    <Alert :alert-data="state.alertData" v-if="state.alertData.open"></Alert>
</template>

<script>
import {useStore} from "vuex";
import useVuelidate from '@vuelidate/core'
import {reactive} from "vue";
import {required} from "@vuelidate/validators"
import Alert from '@/components/Alert'
// import router from "@/router";
export default {
    components: {Alert},
    setup() {
        const store = useStore()
        const state = reactive({
            bar: {
                name: '',
                descr: '',
                city: '',
                address: '',
                lat: '',
                lon: ''
            },
            alertData: {
                open: false,
                status: 0,
                message: 'Bar creado correctamente'
            }
        })

        const rules = {
            bar: {
                name: {required},
                descr: {required},
                city: {required},
                address: {required},
                lat: {required},
                lon: {required}
            }
        }
        const v$ = useVuelidate(rules, state)
        function submit() {
            this.v$.$validate()
            if (!this.v$.$error) {
                store.dispatch("barStore/createBar", state.bar).then((data) => {
                    console.log(data);
                });
            }else{
                state.alertData.open = true
                state.alertData.status = 403
                state.alertData.message = "Datos incorrectos"
            }
        }   

        return  {
            state,
            v$,
            submit
        }
    }
}
</script>

<style>
form{
    width: 40%;
    border-radius: 10px;
    border: 1px solid;
    transition: 0.5s;
}

@media(max-width: 1100px){
    form{
        width: 60%;
    }
}

@media(max-width: 700px){
    form{
        width: 95%;
    }
}


</style>