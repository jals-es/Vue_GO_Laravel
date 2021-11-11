<template>
  <ion-page>
    <ion-tabs>
      <ion-router-outlet></ion-router-outlet>
      <h1>asdads</h1>
      <AdminHeader></AdminHeader>
      <AdminMenu></AdminMenu>
      <!-- <Stats></Stats> -->
    </ion-tabs>
  </ion-page>
</template>

<script lang="ts">
import AdminHeader from '../components/AdminHeader.vue'
import AdminMenu from '../components/AdminMenu.vue'
export default {
  name: "AdminDashboard",
  components: {
    AdminHeader,
    AdminMenu
  },
};
</script>