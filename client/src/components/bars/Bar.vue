<template>
  <tr v-if="type === 'list'">
    <td>
      <b>{{ bar.name }}</b>
    </td>
    <td v-if="bar">
      <router-link :to="'/admin/bars/' + bar.slug">{{ bar.slug }}</router-link>
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
    <td v-else>
      <span class="status text-white bg-warning">Unknown</span>
    </td>

    <td>
      <a href="assd"></a>
      <a href="#"></a>
    </td>
  </tr>
  <div v-else-if="type === 'header'">
    <div class="container header-bar-admin">
      <div class="profile-header">
        <div class="row justify-content-start">
          <div class="col-4">
            <h4 class="m-t-sm">
              {{ bar.name }}
              <span v-if="bar.status === 0" class="status bg-danger text-white"
                >Deleted</span
              >
              <span
                v-else-if="bar.status === 1"
                class="status text-white bg-success"
                >Active</span
              >
              <span
                v-else-if="bar.status === 2"
                class="status text-white bg-secondary"
                >Blocked</span
              >
              <span v-else class="status text-white bg-warning">Unknown</span>
            </h4>
            <p class="m-b-sm">{{ bar.address }} , {{ bar.city }}</p>
          </div>
          <div class="col-4">
            <p v-if="bar.stats">
              <span>Monthly Orders: {{ bar.stats.monthlyOrders }}</span>
            </p>
            <p v-if="bar.stats">
              <span>Daily Orders: {{ bar.stats.dailyOrders }}</span>
            </p>
          </div>
          <div class="col-4">
            <p>Owner: {{ bar.ownerName }}</p>
            <p>Since: {{ bar.created_at }}</p>
          </div>
          <!-- <div class="col-4"></div> -->
        </div>
      </div>
    </div>
  </div>
</template>
     
 <script>
export default {
  name: "Bar",
  props: {
    bar: Object,
    type: String,
    stats: Object
  },
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
.header-bar-admin {
    max-width: 100%;
    padding-top: 1em;
    color: white;
    background-color: #313348;
    font-weight: bold;
    width: 95%;
    border-radius: 5px
;
}
</style>