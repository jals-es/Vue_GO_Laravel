<template>
    <div>
        <Bar :type="'header'" :bar="bar"></Bar>
        <p>hola</p>
    </div>
</template>
<script>
import Bar from '../components/bars/Bar.vue'
import { useStore } from "vuex";
import { computed } from "vue"
import { useRoute } from 'vue-router';
export default {
    // name: SuperAdminBarManagement,
    components: {Bar},
    setup() {
            const route = useRoute();
            console.log(route.params.id);
            
    // Retornem els datos
    const store = useStore();

    // Omplim el state
    // onMounted(() => {
        store.dispatch("superUser/getBarInfo", route.params.id);
    // })
    //Agafem els datos
    const bar = computed(() => store.getters["superUser/getBarInfo"]);
    return {
        bar
    } 
},
}
</script>