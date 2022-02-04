<template>
  <div class="superadmindashboard">
    <Sidebar></Sidebar>

    <section id="wrapper">
      <div class="p-4">
        <div class="welcome">
          <div class="content rounded-3 p-3">
            <h1 class="fs-3">Welcome to Dashboard</h1>
            <p class="mb-0">
              Hello Jone Doe, welcome to your awesome dashboard!
            </p>
          </div>
        </div>

        <DashStats></DashStats>

        <section class="charts mt-4">
          <div class="row">
            <div class="col-lg-6">
              <div class="chart-container rounded-2 p-3">
                <h3 class="fs-6 mb-3">Chart title number one</h3>
                <SuperAdminDashFirstChart></SuperAdminDashFirstChart>
              </div>
            </div>
            <div class="col-lg-6">
              <div class="chart-container rounded-2 p-3">
                <h3 class="fs-6 mb-3">Chart title number two</h3>
                <SuperAdminDashSecondChart></SuperAdminDashSecondChart>
              </div>
            </div>
          </div>
        </section>
        <section class="misc">
          <div class="row">
            <div class="col">
              <SuperAdminListOrders v-bind:id_bar="'%'"></SuperAdminListOrders>
            </div>
            <div class="col"></div>
          </div>
        </section>
      </div>
    </section>
  </div>
</template>

<script>
import DashStats from "../components/SuperAdminDashStats.vue";
import SuperAdminDashFirstChart from "../components/SuperAdminDashFirstChart.vue";
import SuperAdminDashSecondChart from "../components/SuperAdminDashSecondChart.vue";
import Sidebar from "../components/Sidebar";
import SuperAdminListOrders from "../components/SuperAdminListOrders.vue";

import { computed } from "vue";
import { useStore } from "vuex";

// @ is an alias to /src
export default {
  components: {
    DashStats,
    SuperAdminListOrders,
    SuperAdminDashFirstChart,
    SuperAdminDashSecondChart,
    Sidebar,
  },
  name: "SuperAdminDashboard",

  setup() {
    // Cridem al store
    const store = useStore();

    // Omplim el state
    store.dispatch("superUser/getUser");

    //Agafem els datos
    const user = computed(() => store.getters["superUser/getUser"]);

    // Retornem els datos
    return {
      user,
    };
  },
};
</script>
<style scoped>
@import 'https://fonts.googleapis.com/css2?family=Inter:wght@300;400;500;600&display=swap" rel="stylesheet';

:root {
  --dk-gray-100: #f3f4f6;
  --dk-gray-200: #e5e7eb;
  --dk-gray-300: #d1d5db;
  --dk-gray-400: #9ca3af;
  --dk-gray-500: #6b7280;
  --dk-gray-600: #4b5563;
  --dk-gray-700: #374151;
  --dk-gray-800: #1f2937;
  --dk-gray-900: #111827;
  --dk-dark-bg: #313348;
  --dk-darker-bg: #2a2b3d;
  --navbar-bg-color: #6f6486;
  --sidebar-bg-color: #252636;
  --sidebar-width: 250px;
}

* {
  margin: 0;
  padding: 0;
  box-sizing: border-box;
}

body {
  font-family: "Inter", sans-serif;
  background-color: #2a2b3d;
  font-size: 0.925rem;
}

.col-lg-4 {
  width: 30%;
}

.statistics .row {
  justify-content: space-between;
}

#wrapper {
  margin-left: 250px;
  transition: all 0.3s ease-in-out;
}

#wrapper.fullwidth {
  margin-left: 0;
}

/** --------------------------------
 -- Sidebar
-------------------------------- */
.sidebar {
  background-color: #252636;
  width: 250px;
  transition: all 0.3s ease-in-out;
  transform: translateX(0);
  z-index: 9999999;
}

.sidebar .close-aside {
  position: absolute;
  top: 7px;
  right: 7px;
  cursor: pointer;
  color: #eee;
}

.sidebar .sidebar-header {
  border-bottom: 1px solid #2a2b3c;
}

.sidebar .sidebar-header h5 a {
  color: #d1d5db;
}

.sidebar .sidebar-header p {
  color: #9ca3af;
  font-size: 0.825rem;
}

.sidebar .search .form-control ~ i {
  color: #2b2f3a;
  right: 40px;
  top: 22px;
}

.sidebar > ul > li {
  padding: 0.7rem 1.75rem;
}

.sidebar ul > li > a {
  color: #9ca3af;
  text-decoration: none;
}

/* Start numbers */
.sidebar ul > li > a > .num {
  line-height: 0;
  border-radius: 3px;
  font-size: 14px;
  padding: 0px 5px;
}

.sidebar ul > li > i {
  font-size: 18px;
  margin-right: 0.7rem;
  color: #6b7280;
}

.sidebar ul > li.has-dropdown > a:after {
  content: "\eb3a";
  font-family: unicons-line;
  font-size: 1rem;
  line-height: 1.8;
  float: right;
  color: #6b7280;
  transition: all 0.3s ease-in-out;
}

.sidebar ul .opened > a:after {
  transform: rotate(-90deg);
}

/* Start dropdown menu */
.sidebar ul .sidebar-dropdown {
  padding-top: 10px;
  padding-left: 30px;
  display: none;
}

.sidebar ul .sidebar-dropdown.active {
  display: block;
}

.sidebar ul .sidebar-dropdown > li > a {
  font-size: 0.85rem;
  padding: 0.5rem 0;
  display: block;
}

/* End dropdown menu */

.show-sidebar {
  transform: translateX(-270px);
}

@media (max-width: 767px) {
  .sidebar ul > li {
    padding-top: 12px;
    padding-bottom: 12px;
  }

  .sidebar .search {
    padding: 10px 0 10px 30px;
  }
}

/** --------------------------------
 -- welcome
-------------------------------- */
.welcome {
  color: #d1d5db;
}

.welcome .content {
  background-color: #313348;
}

.welcome p {
  color: #9ca3af;
}

/** --------------------------------
 -- Statistics
-------------------------------- */

/** --------------------------------
 -- Charts
-------------------------------- */
.charts .chart-container {
  background-color: #313348;
}

.charts .chart-container h3 {
  color: #9ca3af;
}

/** --------------------------------
 -- users
-------------------------------- */
.admins .box .admin {
  background-color: #313348;
}

.admins .box h3 {
  color: #d1d5db;
}

.admins .box p {
  color: #9ca3af;
}

/** --------------------------------
 -- statis
-------------------------------- */
.statis {
  color: #f3f4f6;
}

.statis .box {
  position: relative;
  overflow: hidden;
  border-radius: 3px;
}

.statis .box h3:after {
  content: "";
  height: 2px;
  width: 70%;
  margin: auto;
  background-color: rgba(255, 255, 255, 0.12);
  display: block;
  margin-top: 10px;
}

.statis .box i {
  position: absolute;
  height: 70px;
  width: 70px;
  font-size: 22px;
  padding: 15px;
  top: -25px;
  left: -25px;
  background-color: rgba(255, 255, 255, 0.15);
  line-height: 60px;
  text-align: right;
  border-radius: 50%;
}

.main-color {
  color: #ffc107;
}

/** --------------------------------
 -- Please don't do that in real-world projects!
 -- overwrite Bootstrap variables instead.
-------------------------------- */

.navbar {
  background-color: #6f6486 !important;
  border: none !important;
}

.navbar .dropdown-menu {
  right: auto !important;
  left: 0 !important;
}

.navbar .navbar-nav > li > a {
  color: #eee !important;
  line-height: 55px !important;
  padding: 0 10px !important;
}

.navbar .navbar-brand {
  color: #fff !important;
}

.navbar .navbar-nav > li > a:focus,
.navbar .navbar-nav > li > a:hover {
  color: #eee !important;
}

.navbar .navbar-nav > .open > a,
.navbar .navbar-nav > .open > a:focus,
.navbar .navbar-nav > .open > a:hover {
  background-color: transparent !important;
  color: #fff !important;
}

.navbar .navbar-brand {
  line-height: 55px !important;
  padding: 0 !important;
}

.navbar .navbar-brand:focus,
.navbar .navbar-brand:hover {
  color: #fff !important;
}

.navbar > .container .navbar-brand,
.navbar > .container-fluid .navbar-brand {
  margin: 0 !important;
}

@media (max-width: 767px) {
  .navbar > .container-fluid .navbar-brand {
    margin-left: 15px !important;
  }

  .navbar .navbar-nav > li > a {
    padding-left: 0 !important;
  }

  .navbar-nav {
    margin: 0 !important;
  }

  .navbar .navbar-collapse,
  .navbar .navbar-form {
    border: none !important;
  }
}

.navbar .navbar-nav > li > a {
  float: left !important;
}

.navbar .navbar-nav > li > a > span:not(.caret) {
  background-color: #e74c3c !important;
  border-radius: 50% !important;
  height: 25px !important;
  width: 25px !important;
  padding: 2px !important;
  font-size: 11px !important;
  position: relative !important;
  top: -10px !important;
  right: 5px !important;
}

.dropdown-menu > li > a {
  padding-top: 5px !important;
  padding-right: 5px !important;
}

.navbar .navbar-nav > li > a > i {
  font-size: 18px !important;
}

/* Start media query */

@media (max-width: 767px) {
  #wrapper {
    margin: 0 !important;
  }

  .statistics .box {
    margin-bottom: 25px !important;
  }

  .navbar .navbar-nav .open .dropdown-menu > li > a {
    color: #ccc !important;
  }

  .navbar .navbar-nav .open .dropdown-menu > li > a:hover {
    color: #fff !important;
  }

  .navbar .navbar-toggle {
    border: none !important;
    color: #eee !important;
    font-size: 18px !important;
  }

  .navbar .navbar-toggle:focus,
  .navbar .navbar-toggle:hover {
    background-color: transparent !important;
  }
}

::-webkit-scrollbar {
  background: transparent;
  width: 5px;
  height: 5px;
}

::-webkit-scrollbar-thumb {
  background-color: #3c3f58;
}

::-webkit-scrollbar-thumb:hover {
  background-color: rgba(0, 0, 0, 0.3);
}
</style>