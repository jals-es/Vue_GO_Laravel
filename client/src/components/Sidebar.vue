<template>
  <div class="d-inline-flex flex-column min-vh-100 bg-dark p-3 position-fixed" style="width: 250px">
    <div class="mb-4">
      <router-link to="/" class="d-flex align-items-center link-light text-decoration-none justify-content-center">
        <img src="https://github.com/mdo.png" alt="" width="64" height="64" class="rounded-2">
        <strong class="ms-2">Admin</strong>
      </router-link>
    </div>
    <ul class="nav nav-pills flex-column mb-auto">
      <li>
        <router-link to="/admin/dashboard" class="nav-link link-light">
         <Icon name="Gauge" class="align-text-bottom"></Icon>
          Dashboard
        </router-link>
      </li>
      <li>
        <router-link to="/admin/orders" class="nav-link link-light">
          <Icon name="Box" class="align-text-bottom"></Icon>
          Orders
        </router-link>
      </li>
      <li>
        <router-link to="/admin/bars" class="nav-link link-light">
          <Icon name="Building" class="align-text-bottom"></Icon>
          Bars
        </router-link>
      </li>
    </ul>
    <hr>
    <router-link to="/home" class="nav-link link-light">
      <Icon name="LogOut" class="align-text-bottom"></Icon>
      Exit
    </router-link>
  </div>
</template>

<script>
import Icon from "@/components/Icon"
export default {
  name: "Sidebar",
  components: {
    Icon
  }
}
</script>

<style scoped>
 hr {
   color: white;
 }
</style>