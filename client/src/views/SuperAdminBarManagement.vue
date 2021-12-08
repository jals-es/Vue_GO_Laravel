<template>
  <div>
    <Bar :type="'header'" :bar="bar"></Bar>
    <div class="row" style="margin: 0 auto; width: 95%">
      <div class="col">
        <SuperAdminListOrders :bar="bar" class="col"></SuperAdminListOrders>
      </div>
      <div class="col bg-danger"></div>
    </div>
  </div>
</template>
<script>
import Bar from "../components/bars/Bar.vue";
import SuperAdminListOrders from "../components/SuperAdminListOrders.vue";
import { useStore } from "vuex";
import { computed } from "vue";
import { useRoute } from "vue-router";
export default {
  // name: SuperAdminBarManagement,
  components: { Bar, SuperAdminListOrders },
  setup() {
    const route = useRoute();
    const ruta = route.params.id;

    // Retornem els datos
    const store = useStore();

    // Omplim el state
    // onMounted(() => {
    store.dispatch("superUser/getBarInfo", route.params.id);
    // })
    //Agafem els datos
    const bar = computed(() => store.getters["superUser/getBarInfo"]);
    return {
      bar,
      ruta,
    };
  },
};
</script>