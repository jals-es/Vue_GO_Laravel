<template>
  <ion-page>
    <ion-header>
      <ion-toolbar>
        <ion-title>Stats</ion-title>
      </ion-toolbar>
    </ion-header>
    <ion-content :fullscreen="true">
      <ion-header collapse="condense">
        <ion-toolbar>
          <ion-title size="large">Stats</ion-title>
        </ion-toolbar>
      </ion-header>
    </ion-content>
  <h1>stats</h1>
  </ion-page>
  
</template>

<script lang="ts">
import {
  IonPage,
  IonHeader,
  IonToolbar,
  IonTitle,
  IonContent,
} from "@ionic/vue";
// import ExploreContainer from '@/components/ExploreContainer.vue';
import { menuController } from "@ionic/vue";
export default {
  name: "Tab1",
  components: { IonHeader, IonToolbar, IonTitle, IonContent, IonPage },
  methods: {
    openMenu() {
      menuController.open("app-menu");
    },
  },
};
</script>