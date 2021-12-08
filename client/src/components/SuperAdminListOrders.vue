<template>
  <p>{{ bar.id }}</p>
  <table class="table table-dark">
    <thead>
      <!-- <th>Order ID</th> -->
      <tr class="table-active">
        <th>Bar</th>
        <th>Table</th>
        <th>Date</th>
        <th>Client</th>
        <th>Price</th>
      </tr>
    </thead>
    <tbody>
    <SuperAdminOrders v-for="o in orders" :order="o" :key="o.id"></SuperAdminOrders>
        
    </tbody>
  </table>
</template>
<script>
import { useStore } from "vuex";
import { computed } from "vue";
import SuperAdminOrders from "./SuperAdminOrders.vue";
export default {
  name: "SuperAdminListBars",
  components: { SuperAdminOrders },
  props: {
    bar: Object,
    // type: String,
    // stats: Object,
  },
  setup() {
    const store = useStore();

    store.dispatch("superUser/getOrders");

    const orders = computed(() => store.getters["superUser/getOrders"]);

    return {
      orders,
    };
  },
};
</script>