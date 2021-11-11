<template>
  <ion-header>
    <ion-toolbar>
      <ion-title>Menu - Basic</ion-title>
    </ion-toolbar>
    <ion-button @click="openMenu()">Open Menu</ion-button>
  </ion-header>
</template>
<script>
import { IonHeader, IonToolbar, IonTitle, IonButton, menuController } from "@ionic/vue";
export default {
  methods: {
    openMenu() {
      menuController.enable(true, 'admin-left-menu');
      menuController.open('admin-left-menu');
    },
  },
  name: "AdminHeader",
  components: {
    IonHeader,
    IonToolbar,
    IonTitle,
    IonButton,
  },
};
</script>