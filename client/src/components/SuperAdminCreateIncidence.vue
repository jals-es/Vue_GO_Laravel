<template>
    <button type="button" class="btn btn-primary" data-bs-toggle="modal" data-bs-target="#createIncidenceModal">Create a new incidence</button>

<div class="modal fade" id="createIncidenceModal" tabindex="-1" aria-labelledby="exampleModalLabel" aria-hidden="true">
  <div class="modal-dialog">
    <div class="modal-content">
      <div class="modal-header">
        <h5 class="modal-title" id="exampleModalLabel">New Incidence</h5>
        <button type="button" class="btn-close" data-bs-dismiss="modal" aria-label="Close"></button>
      </div>
      <div class="modal-body">
        <form>
          <div class="mb-3">
            <label for="recipient-name" class="col-form-label">Name:</label>
            <input type="text" v-model="nameInput" class="form-control" id="recipient-name">
          </div>
          <div class="mb-3">
            <label for="message-text"  class="col-form-label">Description:</label>
            <textarea v-model="descrInput" class="form-control" id="message-text"></textarea>
          </div>
          <div class="custom-file">
            <input type="file" name="file[]" class="custom-file-input" id="customFileInput" multiple aria-describedby="customFileInput" @change="saveFile">
            <label class="custom-file-label" for="customFileInput" id="fileInput" >Select file</label>
          </div>
        </form>
      </div>
      <div class="modal-footer">
        <button type="button" class="btn btn-secondary"  data-bs-dismiss="modal">Close</button>
        <button type="button" class="btn btn-primary" @click="createIncidence">Create</button>
      </div>
    </div>
  </div>
</div>
</template>
<script>
import {  reactive, toRefs } from 'vue';
import {useStore} from "vuex";

export default {
    name: "SuperAdminCreateIncidence",
    setup() {
        const incidenceState = reactive({
            nameInput: '',
            descrInput: ''
        })
        function validateForm() {
            if ( incidenceState.nameInput.replace(" ","").length > 0 && incidenceState.nameInput.replace(" ","").length > 0) { //Length validation
                return true
            }
            return false
        }
        const store = useStore();

        function createIncidence() {
            if (validateForm()) {
                const formData = new FormData();
                formData.append("name",incidenceState.nameInput)
                formData.append("descr",incidenceState.descrInput)
                // Agafem el input de les files
                const inputFiles = document.getElementById("customFileInput").files
                for (let index = 0; index < inputFiles.length; index++) {
                    formData.append("file[]", inputFiles[index]);
                }
                // Omplim el state
                store.dispatch("superUser/createIncidence", formData);
            }
        }
        return {
            createIncidence,
            ...toRefs(incidenceState)
        }
    },
}
</script>