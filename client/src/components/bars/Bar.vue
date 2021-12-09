<template>
    <tr v-if="type === 'list'">
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
        <a href="assd"></a>
        <a href="#"></a>
      </td>
    </tr>
    <div v-if="type === 'header'">
      <div class="container">
            <div class="profile-header">
        <div class="profile-header-cover"></div>
        <div class="profile-header-content">
            <div class="profile-header-img mb-4">
                <img src="https://bootdey.com/img/Content/avatar/avatar7.png" class="mb-4" alt="" />
            </div>

            <div class="profile-header-info">
                <h4 class="m-t-sm">Clyde Stanley</h4>
                <p class="m-b-sm">UXUI + Frontend Developer</p>
                <a href="#" class="btn btn-xs btn-primary mb-2">Edit Profile</a>
            </div>
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