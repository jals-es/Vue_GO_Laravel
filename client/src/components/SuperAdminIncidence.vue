<template>
  <div class="container">
    <div class="principalIncidence row">
      <div class="col-2 colPhoto">
        <img
          v-bind:src="incidence.photos[0].path"
          style="border-top-left-radius: 8px; border-bottom-left-radius: 8px"
          class="img-fluid"
        />
      </div>
      <div class="col-4">
        <p class="text-white" style="font-size: 1.6em">
          {{ incidence.name }}
        </p>
      </div>
      <div class="col-4">
        <p v-if="incidence.status === 0" class="status bg-danger text-white">
          Active
        </p>
        <p v-else-if="incidence.status === 1" class="status text-white bg-success">
          Closed
        </p>
        <p v-else-if="incidence.status === 2" class="status text-white bg-secondary">
          Deleted
        </p>
        <p v-else class="status text-white bg-warning">Unknown</p>
      </div>
      <div class="col-2">
        <button class="btn btn-dark h-100">Show More</button>
      </div>
      <div></div>
    </div>
  </div>
</template>
<script>
export default {
  name: "SuperAdminIncidence",
  props: {
    incidence: Object,
  },
  setup() {},
};
</script>
<style scoped>
.colPhoto {
  padding-left: 0 !important;
}
.principalIncidence {
  width: 82%;
  border-radius: 8px;
  border: 2px solid #313348;
  margin-top: 8px;
  margin-bottom: 8px;
  background-color: #313348;
}
.status {
  padding: 4px;
  border-radius: 8px;
  position: relative;
  margin: 0 auto;
  margin-top: 1%;
  font-size: 1.6em;
  display: inline-block;
  /* width: 10%; */
}
</style>