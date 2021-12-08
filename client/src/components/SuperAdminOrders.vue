<template>
  <tr>
    <td>{{order.slug}}</td>
    <td>{{order.id_table}}</td>
    <td>{{order.date}}</td>
    <td>{{order.name}}</td>
    <td>{{order.price}}</td>

    <!-- <li v-bind:key="order.id" v-for="(item, key, index) in order">
      {{ item }} - {{ key }} - {{ index }}
    </li> -->
  </tr>
</template>
<script>
export default {
  name: "SuperAdminOrders",
  props: {
    order: Object,
    bar: String,
    // stats: Object,
  },
  setup(order) {
    console.log(order);
  },
};
</script>