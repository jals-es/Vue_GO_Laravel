<template>
<div class="incidencesManagement">
    <SuperAdminCreateIncidence></SuperAdminCreateIncidence>
    <SuperAdminListIncidences></SuperAdminListIncidences>
</div>
</template>
<script>
import SuperAdminCreateIncidence from '../components/SuperAdminCreateIncidence.vue'
import SuperAdminListIncidences from '../components/SuperAdminListIncidences.vue'
export default {
    components: {SuperAdminCreateIncidence,SuperAdminListIncidences},
    setup() {
        
    },
}
</script>
<style scoped>
.incidencesManagement {
  padding-top: 25px;
  margin-left: 280px;
}
</style>