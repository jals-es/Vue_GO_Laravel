<template>
  <tr>
    <td><router-link :key="order.slug" :to="'/admin/bars/' + order.slug">{{order.slug}}</router-link></td>
    <td>{{order.id_table}}</td>
    <td>{{order.date}}</td>
    <td>{{order.name}}</td>
    <td>{{order.price}} €</td>

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

  setup() {
  },
};
</script>