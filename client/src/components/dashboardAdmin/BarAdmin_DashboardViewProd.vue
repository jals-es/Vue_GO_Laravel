<template>
    <div class="modal fade" :id="prod.slug" tabindex="-1" aria-labelledby="exampleModalLabel" aria-hidden="true">
        <div class="modal-dialog">
            <div class="modal-content">
            <div class="modal-header">
                <h5 class="modal-title" id="exampleModalLabel">{{ prod.name }}</h5>
                <button type="button" class="btn-close" data-bs-dismiss="modal" aria-label="Close"></button>
            </div>
            <div class="modal-body">
                <div>
                    <img class="img-fluid" :src="prod.photo">
                </div>
                <p class="my-2">Descripción: {{ prod.descr }}</p>
                <h4 class="mt-3">Tipos:</h4>
                <div class="tipo mb-2" v-for="thisType in prod.types" :key="thisType.name">
                    <p class="my-0">Nombre: {{ thisType.name }}</p>
                    <p class="my-0">Descripción: {{ thisType.descr }}</p>
                    <p class="my-0">Precio: {{ thisType.price }}€</p>
                </div>
                <h4 class="mt-3">Extras:</h4>
                <div class="tipo mb-2" v-for="extra in prod.extras" :key="extra.name">
                    <p class="my-0">Nombre: {{ extra.name }}</p>
                    <p class="my-0">Descripción: {{ extra.descr }}</p>
                    <p class="my-0">Precio: {{ extra.price }}€</p>
                </div>
            </div>
            </div>
        </div>
    </div>
</template>

<script>
export default {
  name: "DashboardViewProd",
  props: {
    prod: Object
  },
  setup() {
    // document.getElementById(`menu${active}`).classList.add("active");
  },
};
</script>

<style>
    
</style>