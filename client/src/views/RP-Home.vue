<template>
  <div class="card w-50" style="height: 350px">
    <div class="card-header">
      <h1>
        Role List
      </h1>
    </div>
    <div class="card-body overflow-auto h-50">
      <div class="row gap-2 p-2">
          <RolePill v-for="role in compData.roles" :role="role" :key="role.id"></RolePill>
      </div>

    </div>
    <div class="card-footer">
      <div class="d-flex justify-content-between">
        <p class="mt-2 mb-0">Total roles: {{compData.roles.length}}</p>
        <button type="submit" class="btn btn-success">Add Role</button>
      </div>
    </div>
  </div>
</template>

<script>
import RolePill from "@/components/RolePill"
import {roles} from "@/composables/role_permission";

export default {
  name: "RPHome",
  components: {RolePill},
  setup() {
   const {compData, getRoles} = roles()
    getRoles()

    return {
     compData
    }
  }
}
</script>

<style scoped>

</style>
