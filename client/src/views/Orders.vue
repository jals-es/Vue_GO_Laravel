<template>
    <ion-content>
      <ion-button @click="openMenu()" expand="full">Open Menu</ion-button>
      <h1>Orders</h1>
    </ion-content>
</template>
<!-- Code hidden for display purpose -->
<script>
import {menuController} from "@ionic/vue";
export default {
    methods:{
      openMenu(){
        menuController.open("app-menu")
      }
    }
};
</script>