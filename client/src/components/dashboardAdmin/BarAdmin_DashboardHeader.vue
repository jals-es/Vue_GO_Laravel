<template>
    <header class="navbar navbar-dark sticky-top bg-dark flex-md-nowrap p-0 shadow">
        <router-link class="navbar-brand col-md-3 col-lg-2 me-0 px-3" to="/bars/">{{ bar.name }} ({{ bar.city }})</router-link>
        <button class="navbar-toggler position-absolute d-md-none collapsed" type="button" data-bs-toggle="collapse" data-bs-target="#sidebarMenu" aria-controls="sidebarMenu" aria-expanded="false" aria-label="Toggle navigation">
            <span class="navbar-toggler-icon"></span>
        </button>
        <input class="form-control form-control-dark w-100" type="text" placeholder="Search" aria-label="Search">
        <div class="navbar-nav">
            <div class="nav-item text-nowrap">
            <a class="nav-link px-3 user-select-none pe-pointer" @click="logOut">Sign out</a>
            </div>
        </div>
    </header>    
</template>

<script>
import router from '@/router'
import { bars } from '@/composables/bars'
export default {
  name: "DashboardHeader",
  props: {
      barSlug: String
  },
  setup(barSlug) {
    function logOut(){
        localStorage.removeItem("token")
        router.push("/login")
    }

    console.log(barSlug.barSlug)

    const { getBarInfo, barsInfo } = bars();

    getBarInfo(barSlug.barSlug)

    var bar = barsInfo.thisBar

    return {
        logOut,
        bar
    }
  },
};
</script>

<style>
.pe-pointer{
    cursor: pointer
}

/*
 * Navbar
 */

.navbar-brand {
  padding-top: .75rem;
  padding-bottom: .75rem;
  font-size: 1rem;
  background-color: rgba(0, 0, 0, .25);
  box-shadow: inset -1px 0 0 rgba(0, 0, 0, .25);
}

.navbar .navbar-toggler {
  top: .25rem;
  right: 1rem;
}

.navbar .form-control {
  padding: .75rem 1rem;
  border-width: 0;
  border-radius: 0;
}

.form-control-dark {
  color: #fff;
  background-color: rgba(255, 255, 255, .1);
  border-color: rgba(255, 255, 255, .1);
}

.form-control-dark:focus {
  border-color: transparent;
  box-shadow: 0 0 0 3px rgba(255, 255, 255, .25);
}
</style>