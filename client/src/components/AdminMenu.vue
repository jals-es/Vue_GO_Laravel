<template>
  <ion-menu menu-id="admin-left-menu" auto-hide="true" content-id="main-content" @click="closeMenu()">
    <ion-header>
      <ion-toolbar>
        <ion-title>Close! X</ion-title>
      </ion-toolbar>
    </ion-header>
    <ion-content>
      <ion-list>
        <ion-item>
          <ion-label>Home</ion-label>
        </ion-item>
        <ion-item>
          <ion-label>About</ion-label>
        </ion-item>
      </ion-list>
    </ion-content>
  </ion-menu>
</template>
<script>
import {menuController} from "@ionic/vue";

import {
  IonMenu,IonHeader,IonContent,IonToolbar,IonTitle,IonItem,IonList,IonLabel,
} from "@ionic/vue";
export default {
      methods:{
      closeMenu(){
        console.log("asdasd")
        menuController.close("admin-left-menu")
      }
    },
  components: {
    IonMenu,IonHeader,IonContent,IonToolbar,IonTitle,IonItem,IonList,IonLabel,
  },
  
};
</script>