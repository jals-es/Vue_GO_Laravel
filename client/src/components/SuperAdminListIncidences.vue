<template>
    <div>
        <SuperAdminIncidence v-for="incidence in incidences" :key="incidence.id" :incidence="incidence"></SuperAdminIncidence>
    </div>
</template>
<script>
import SuperAdminIncidence from "./SuperAdminIncidence.vue"
import { useStore } from "vuex";
import { computed } from "vue";
export default {
    name:"SuperAdminListIncidences",
    components: {SuperAdminIncidence},
    setup() {
        const store = useStore();
        store.dispatch("superUser/getIncidences");
        const incidences = computed(() => store.getters["superUser/getIncidences"]);
        return {
            incidences
        }
    },
}
</script>