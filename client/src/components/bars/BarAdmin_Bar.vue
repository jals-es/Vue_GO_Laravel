<template>
    <router-link :to="`/${bar.slug}/admin/`" class="bar">
        <span class="name">{{ bar.name }}</span>
        <span class="city">{{ bar.city }}</span>
        <span class="rol">{{ bar.rol }}</span>
    </router-link>
</template>

<script>
// import router from "@/router"

export default {
    name: "BarAdminBar",
    props: {
        bar: Object
    },
    setup() {
        
    },
};

</script>