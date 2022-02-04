<template>
    <p>{{msg}} <span>{{author}}</span></p>
</template>
<script>
export default {
    props: {
        msg: String,
        author: String,
    },
    setup() {
        
    },
}
</script>