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
        <p
          v-else-if="incidence.status === 1"
          class="status text-white bg-success"
        >
          Closed
        </p>
        <p
          v-else-if="incidence.status === 2"
          class="status text-white bg-secondary"
        >
          Deleted
        </p>
        <p v-else class="status text-white bg-warning">Unknown</p>
      </div>
      <div class="col-2">
        <button
          class="btn btn-dark h-100"
          data-bs-toggle="collapse"
          v-bind:data-bs-target="'#cll' + incidence.id"
        >
          Show More
        </button>
      </div>
    </div>
    <div class="row collapse" v-bind:id="'cll' + incidence.id">
      <div class="row" style="width: 82%">
        <div class="col" v-for="photo in incidence.photos" :key="photo.name">
          <img
            class="img-fluid"
            v-bind:src="photo.path"
            style="max-height: 280px"
          />
        </div>
      </div>
      <div class="row" style="width: 82%">
        <h2>{{ incidence.name }}</h2>
        <h4>{{ incidence.descr }}</h4>
      </div>
      <div v-if="incidence.fix.length != 0" class="row" style="width: 82%">
        <div class="alert alert-success" role="alert">Fixed!</div>
        <h4>{{ incidence.fix }}</h4>
      </div>
      <div v-else class="row" style="width: 82%">
        <div class="form-floating">
          <textarea
            class="form-control"
            placeholder="Explain how you solved it!"
            id="incidenceFix"
            style="height: 120px"
            v-model="incidenceFix"
          ></textarea>
          <label for="incidenceFix" style="padding: 1rem 1.75rem"
            >Explain how you solved it!</label
          >
        </div>
        <div class="row">
          <button
            class="col btn btn-success"
            @click="closeIncidence(incidence.id)"
          >
            Close
          </button>
          <button class="col btn btn-danger">Remove</button>
        </div>
      </div>
    </div>
  </div>
</template>
<script>
import {useStore} from "vuex";
import {  reactive, toRefs } from 'vue';
export default {
  name: "SuperAdminIncidence",
  props: {
    incidence: Object,
  },
  setup() {
    const store = useStore();
    const incidenceFix = reactive({
      incidenceFix: "",
    });
    function deleteIncidence(id) {
      store.dispatch("superUser/deleteIncidence", id);
    }
    function closeIncidence(id) {
      store.dispatch("superUser/closeIncidence", id, incidenceFix.incidenceFix);
    }
    return {
      deleteIncidence,
      closeIncidence,
      ...toRefs(incidenceFix)
    };
  },
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