<template>
  <div class="container-xl">
    <div class="table-responsive">
      <div class="table-wrapper">
        <div class="table-title">
          <div class="row">
            <div class="col-sm-5">
              <span class="h2">Bar <b>Management</b></span>
            </div>
            <div class="col-sm-5">
                
              <div class="input-group">
                <input
                  type="search"
                  class="form-control rounded"
                  placeholder="Search"
                  aria-label="Search"
                  aria-describedby="search-addon"
                />
                <button type="button" class="btn btn-outline-primary">
                  Search
                </button>
              </div>
            </div>
          </div>
        </div>
        <table class="table table-striped table-hover">
          <thead>
            <tr>
              <th>Bar</th>
              <th>Link</th>
              <th>Location</th>
              <th>Date Created</th>
              <th>Owner</th>
              <th>Status</th>
              <th>Action</th>
            </tr>
          </thead>
          <tbody>
            <Bar v-for="b in bars" :key="b.id" :bar="b"> </Bar>
          </tbody>
        </table>
      </div>
    </div>
  </div>
</template>
<script>
import Bar from "../components/bars/Bar.vue";
import { computed } from "vue";
import { useStore } from "vuex";
export default {
  components: { Bar },

  setup() {
    // Cridem al store
    const store = useStore();

    // Omplim el state
    store.dispatch("superUser/getBars");

    //Agafem els datos
    const bars = computed(() => store.getters["superUser/getBars"]);
    return {
      bars,
    };
  },
};
</script>