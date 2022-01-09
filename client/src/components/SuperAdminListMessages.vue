<template>
      <SuperAdminMessage v-for="msg in messages"  :key="msg.message" :msg="msg.message" :author="msg.author"></SuperAdminMessage>
    
</template>
<script>
import SuperAdminMessage from "../components/SuperAdminMessage.vue";

export default {
    components: { SuperAdminMessage },
    props: {
        messages: Array,
    },
}
</script>