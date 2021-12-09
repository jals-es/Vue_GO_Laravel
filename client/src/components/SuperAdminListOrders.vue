<template>
  <div class="orderList">
    <h3>Order List {{slug}}</h3>
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
        <SuperAdminOrders
          v-for="o in orders"
          :order="o"
          :key="o.id"
        ></SuperAdminOrders>
      </tbody>
    </table>
  </div>
</template>
<script>
import { useStore } from "vuex";
import { computed } from "vue";
import SuperAdminOrders from "./SuperAdminOrders.vue";
export default {
  name: "SuperAdminListBars",
  components: { SuperAdminOrders },
  props: {
    id_bar: String,
    // type: String,
    // stats: Object,
  },
  setup(id_bar) {
    const store = useStore();

    store.dispatch("superUser/getOrders", id_bar);
    const orders = computed(() => store.getters["superUser/getOrders"]);

    return {
      orders,
    };
  },
};
</script>
<style scoped>
  .orderList {
    margin-top: 1em ;
    padding-top: 1em;
    color: white;
    background-color: #313348;
    font-weight: bold;
    border-radius: 5px;
  }
</style>