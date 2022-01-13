<template>
  <button type="button" class="col-2 btn" :class="{'btn-primary': this.$props.role.status === 0, 'btn-secondary': this.$props.role.status !== 0}" data-bs-toggle="modal" :data-bs-target="`#${this.$props.role.id}`"
          style="width:fit-content;">
    {{this.$props.role.name}}
  </button>
  <RoleDetails :role="this.$props.role.id" ></RoleDetails>
</template>
<script>
import RoleDetails from '@/components/RoleDetails'

export default {
  name: "RolePill",
  components: {RoleDetails},
  props:["role"],
  setup() {
  },

}
</script>

<style scoped>

</style>
