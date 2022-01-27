<template>
        <!-- Header -->
<DashboardHeader :barSlug="bar_slug"></DashboardHeader>
<div class="container-fluid">
    <div class="row">
        <!-- Menu -->
        <DashboardMenu :barSlug="bar_slug"></DashboardMenu>
        <main class="col-md-9 ms-sm-auto col-lg-10 px-md-4">
            <form class="my-5 mx-auto p-3" @submit.prevent="submit">
                <h2 class="mb-3">Crear Producto</h2>
                <h5 class="mt-3 mb-2">Datos:</h5>
                <div class="form-floating">
                    <input v-model="state.prod.name" type="text" class="form-control my-1" id="floatingName" placeholder="Nombre">
                    <label for="floatingName">Nombre</label>
                </div>
                <div class="form-floating">
                    <input v-model="state.prod.descr" type="text" class="form-control my-1" id="floatingDescr" placeholder="Descripcion">
                    <label for="floatingDescr">Descripción</label>
                </div>
                <div class="form-floating">
                    <input v-model="state.prod.catego" type="text" class="form-control my-1" id="floatingCatego" placeholder="Categoria">
                    <label for="floatingCatego">Categoría</label>
                </div>
                <div class="form-floating">
                    <input v-model="state.prod.photo" type="url" class="form-control my-1" id="floatingPhoto" placeholder="URL Foto">
                    <label for="floatingPhoto">URL Foto</label>
                </div>
                <h5 class="mt-3 mb-2">Tipos:</h5>
                <div class="form-floating">
                    <a class="btn btn-outline-dark rounded-circle" data-bs-toggle="modal" data-bs-target="#createType">+</a>
                    <a class="btn btn-outline-primary del mx-2" v-for="(thisType, index) in state.prod.types" :key="thisType.name" @click="deleteType(index)">{{ thisType.name }}</a>
                </div>
                <h5 class="mt-3 mb-2">Extras:</h5>
                <div class="form-floating">
                    <a class="btn btn-outline-dark rounded-circle" data-bs-toggle="modal" data-bs-target="#createExtra">+</a>
                    <a class="btn btn-outline-primary del mx-2" v-for="(thisExtra, index) in state.prod.extras" :key="thisExtra.name" @click="deleteExtra(index)">{{thisExtra.name}}</a>
                </div>
                <div class="mt-4">
                    <button class="w-100 btn btn-lg btn-primary" type="submit">Crear</button>
                </div>
            </form>
            <form class="modal fade p-3" id="createType" tabindex="-1" aria-labelledby="exampleModalLabel" aria-hidden="true" @submit.prevent="submitType">
                <div class="modal-dialog">
                    <div class="modal-content">
                        <div class="modal-header">
                            <h5 class="modal-title" id="exampleModalLabel">Crear Tipo</h5>
                            <button type="button" class="btn-close" data-bs-dismiss="modal" aria-label="Close"></button>
                        </div>
                        <div class="modal-body">
                            <div class="form-floating">
                                <input v-model="models.type.name" type="text" class="form-control my-1" id="floatingName" placeholder="URL Foto">
                                <label for="floatingName">Nombre</label>
                            </div>
                            <div class="form-floating">
                                <input v-model="models.type.descr" type="text" class="form-control my-1" id="floatingDescr" placeholder="URL Foto">
                                <label for="floatingDescr">Descripcion</label>
                            </div>
                            <div class="form-floating">
                                <input v-model="models.type.price" type="number" step='0.01' class="form-control my-1" id="floatingPrice" placeholder="URL Foto">
                                <label for="floatingPrice">Precio</label>
                            </div>
                        </div>
                        <div class="modal-footer">
                            <button class="w-100 btn btn-lg btn-primary" type="submit" data-bs-dismiss="modal" aria-label="Close">Crear</button>
                        </div>
                    </div>
                </div>
            </form>
            <form class="modal fade p-3" id="createExtra" tabindex="-1" aria-labelledby="exampleModalLabel" aria-hidden="true" @submit.prevent="submitExtra">
                <div class="modal-dialog">
                    <div class="modal-content">
                        <div class="modal-header">
                            <h5 class="modal-title" id="exampleModalLabel">Crear Extra</h5>
                            <button type="button" class="btn-close" data-bs-dismiss="modal" aria-label="Close"></button>
                        </div>
                        <div class="modal-body">
                            <div class="form-floating">
                                <input v-model="models.extra.name" type="text" class="form-control my-1" id="floatingName" placeholder="URL Foto">
                                <label for="floatingName">Nombre</label>
                            </div>
                            <div class="form-floating">
                                <input v-model="models.extra.descr" type="text" class="form-control my-1" id="floatingDescr" placeholder="URL Foto">
                                <label for="floatingDescr">Descripcion</label>
                            </div>
                            <div class="form-floating">
                                <input v-model="models.extra.price" type="number" step='0.01' class="form-control my-1" id="floatingPrice" placeholder="URL Foto">
                                <label for="floatingPrice">Precio</label>
                            </div>
                        </div>
                        <div class="modal-footer">
                            <button class="w-100 btn btn-lg btn-primary" data-bs-dismiss="modal" aria-label="Close" type="submit">Crear</button>
                        </div>
                    </div>
                </div>
            </form>
        </main>
    </div>
</div>
</template>

<script>
import { useRoute } from "vue-router";
import {reactive} from "vue";
import useVuelidate from '@vuelidate/core'
import {required, minLength} from "@vuelidate/validators"
import DashboardMenu from '@/components/dashboardAdmin/BarAdmin_DashboardMenu'
import DashboardHeader from '@/components/dashboardAdmin/BarAdmin_DashboardHeader'
import { prods } from '@/composables/prods'
export default {
    components: {
      DashboardMenu,
      DashboardHeader
    },
    setup() {
        const route = useRoute();
        const bar_slug = route.params.slug;
        const { createProd } = prods()

        const state = reactive({
            prod: {
                name: '',
                descr: '',
                catego: '',
                photo: '',
                types: [],
                extras: []
            }
        })

        const rules = {
            prod: {
                name: {required},
                descr: {required},
                catego: {required},
                photo: {required},
                types: {
                    required,
                    minLength: minLength(1)
                },
                extras: {
                    required,
                    minLength: minLength(1)
                },
            }
        }

        const models = {
            type: {
                name: '',
                descr: '',
                price: ''
            },
            extra: {
                name: '',
                descr: '',
                price: ''
            }
        }

        const deleteType = (index) => {
            state.prod.types.splice(index, 1)
        }

        const deleteExtra = (index) => {
            state.prod.extras.splice(index, 1)
        }

        const submitType = () => {
            state.prod.types.push(models.type)
            models.type = {
                name: '',
                descr: '',
                price: ''
            }
        }

        const submitExtra = () => {
            state.prod.extras.push(models.extra)
            models.extra = {
                name: '',
                descr: '',
                price: ''
            }
        }

        const v$ = useVuelidate(rules, state)
        const submit = () => {
            // v$.$validate();
            // console.log("submit")
            // if(!v$.$error) {
            //     console.log("no error submit")
                createProd(state.prod, bar_slug);
            // }
        }

        return {
            bar_slug,
            state,
            deleteType,
            deleteExtra,
            submitType,
            submitExtra,
            submit,
            v$,
            models
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

.del:hover{
    background-color: red;
    border-color: red;
}

.del:hover:before{
    content:"\01F5D1";
}

</style>