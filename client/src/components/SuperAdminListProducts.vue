<template>
  <div>
    <h1>Llistat de productes</h1>
    <div class="" style="display:flex; align-items:flex-start; justify-content: space-between; flex-wrap: wrap;">
    <SuperAdminProduct
      v-for="prod in products"
      :product="prod"
      :key="prod.id"
    ></SuperAdminProduct>
    </div>
  </div>
</template>
<script>
import { useStore } from "vuex";
import { computed } from "vue";
import SuperAdminProduct from "./SuperAdminProduct.vue";
export default {
  components: { SuperAdminProduct },
  name: "SuperAdminListProducts",
  props: {
    idBar: String,
  },
  setup(idBar) {
    const store = useStore();
    store.dispatch("superUser/getProducts", idBar);
    const products = computed(() => store.getters["superUser/getProducts"]);
    return {
      products,
    };
  },
};
</script>
