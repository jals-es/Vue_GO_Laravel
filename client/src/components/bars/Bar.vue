<template>
  <tr>
    <td>
      <b>{{ bar.name }}</b>
    </td>
    <td>
      <a href="">{{ bar.slug }}</a>
    </td>
    <td>{{ bar.city }} {{ bar.address }}</td>
    <td>{{ bar.created_at }}</td>
    <td>{{ bar.ownerName }}</td>
    <td v-if="bar.status === 0">
      <span class="status bg-danger text-white">Deleted</span>
    </td>
    <td v-else-if="bar.status === 1">
      <span class="status text-white bg-success">Active</span>
    </td>
    <td v-else-if="bar.status === 2">
      <span class="status text-white bg-secondary">Blocked</span>
    </td>
    <td v-else><span class="status text-warning">&#11044;Unknown</span></td>

    <td>
      <a href="asdasd"></a>
      <a href="#"></a>
    </td>
  </tr>
</template>
     
 <script>
export default {
  name: "Bar",
  props: {
    bar: Object,
  },
  // props: ["key"],
  setup() {
    //   console.log(bar);
  },
};
</script>
<style scoped>
.status {
    padding: 6px;
    border-radius: 8px;
}
</style>