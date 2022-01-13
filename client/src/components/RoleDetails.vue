<template>
  <div class="modal fade" :id="this.$props.role" tabindex="-1" aria-labelledby="exampleModalLabel"
       aria-hidden="true">
    <div class="modal-dialog">
      <div class="modal-content">
        <div class="modal-header">
          <h5 class="modal-title" id="exampleModalLabel">Details for role: {{ this.compData.role.name }}</h5>
          <button type="button" class="btn-close" data-bs-dismiss="modal" aria-label="Close"></button>
        </div>
        <div class="modal-body overflow-auto" style="height: 280px;">
          <h5 class="pt-4">Role Data</h5>
          <div class="form-floating">
            <input type="text" class="form-control"
                   placeholder="ID" id="idField" disabled>
            <label for="idField">{{ this.compData.role.id }}</label>
          </div>
          <div class="form-floating">
            <input type="text" class="form-control"
                   placeholder="Name" id="nameField">
            <label for="nameField">{{ this.compData.role.name }}</label>
          </div>
          <div class="form-floating">
            <input type="text" class="form-control" placeholder="Status" id="statusField" disabled>
            <label for="statusField">{{ this.compData.role.status === 0 ? 'Active' : 'Disabled' }}</label>
          </div>
          <div class="mt-3">
            <h5 class="pt-4">Permissions Assigned</h5>
            <ul class="list-group">
              <li class="list-group-item d-flex justify-content-between align-content-center"
                  v-for="permission in this.compData.role.permissions" :key="permission.id">
                {{ permission.name }}
                <Icon name="X" class="normalColor"></Icon>
              </li>
            </ul>
          </div>
          <div class="mt-3">
            <h5 class="py-4">Assign Permission</h5>
            <div class="form-floating">
              <select class="form-select" id="permission" aria-label="Permission Selector">
                <option v-for="permission in compData.permissions" :key="permission.id">{{permission.name}}</option>
              </select>
              <label for="statusField">Permissions</label>
            </div>
            <button type="button" class="btn btn-success mt-3 w-100">Assign Permission</button>
          </div>
        </div>
        <div class="modal-footer">
          <button type="button" class="btn"
                  :class="{'btn-warning': this.compData.role.status === 0, 'btn-success': this.compData.role.status !== 0}"
                   @click="changeRoleStatus()">
            {{ this.compData.role.status === 0 ? 'Disable Role' : 'Enable Role' }}
          </button>
          <button type="button" class="btn btn-primary" data-bs-dismiss="modal" @click="saveChanges(compData.role)">Save changes</button>
        </div>
      </div>
    </div>

  </div>
</template>

<script>
import Icon from '@/components/Icon'
import {roles} from "@/composables/role_permission";
//import {reactive} from "vue";

export default {
  name: "RoleDetails",
  props: ["role"],
  emits: ["data"],
  components: {Icon},
  setup(props) {
    const {compData, getOneRole, disableRole, enableRole, saveChanges, getPermissions} = roles()
    getOneRole(props.role)
    getPermissions()

    function changeRoleStatus() {
      const disable = function() {
        compData.role.status = 1
        disableRole(props.role)
      }

      const enable = function () {
        compData.role.status = 0
        enableRole(props.role)
        console.log(compData.role)
      }

      compData.role.status === 0 ? disable() : enable()
      disableRole(props.role)
      console.log(compData.role)


    }

    return {
      saveChanges,
      changeRoleStatus,
      compData
    }
  },

}

</script>

<style scoped>
.form-floating:focus-within {
  z-index: 4;
}

#idField {
  margin-bottom: -1px;
  border-radius: 0.25rem 0.25rem 0 0;
}

#nameField {
  border-radius: 0;
  margin-bottom: -1px;
}

#statusField {
  border-radius: 0 0 0.25rem 0.25rem;
}

.normalColor {
  color: #dc3545;
}

</style>
