<template>
<div class="incidencesManagement">
    <SuperAdminCreateIncidence></SuperAdminCreateIncidence>
</div>
</template>
<script>
import SuperAdminCreateIncidence from '../components/SuperAdminCreateIncidence.vue'
export default {
    components: {SuperAdminCreateIncidence},
    setup() {
        
    },
}
</script>
<style scoped>
.incidencesManagement {
  padding-top: 25px;
  margin-left: 260px;
}
</style>