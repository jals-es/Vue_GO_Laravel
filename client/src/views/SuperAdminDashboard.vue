<template>
  <div class="superadmindashboard">
    <Sidebar></Sidebar>

    <section id="wrapper">
      <div class="p-4">
        <div class="welcome">
          <div class="content rounded-3 p-3">
            <h1 class="fs-3">Welcome to Dashboard</h1>
            <p class="mb-0">Hello Jone Doe, welcome to your awesome dashboard!</p>
          </div>
        </div>

        <DashStats></DashStats>


        <section class="charts mt-4">
          <div class="row">
            <div class="col-lg-6">
              <div class="chart-container rounded-2 p-3">
                <h3 class="fs-6 mb-3">Chart title number one</h3>
                <SuperAdminDashChart1></SuperAdminDashChart1>
              </div>
            </div>
            <div class="col-lg-6">
              <div class="chart-container rounded-2 p-3">
                <h3 class="fs-6 mb-3">Chart title number two</h3>
                <canvas id="myChart2"></canvas>
              </div>
            </div>
          </div>
        </section>

        <section class="admins mt-4">
          <div class="row">
            <div class="col-md-6">
              <div class="box">
                <!-- <h4>Admins:</h4> -->
                <div class="admin d-flex align-items-center rounded-2 p-3 mb-4">
                  <div class="img">
                    <img class="img-fluid rounded-pill"
                         width="75" height="75"
                         src="https://uniim1.shutterfly.com/ng/services/mediarender/THISLIFE/021036514417/media/23148906966/small/1501685402/enhance"
                         alt="admin">
                  </div>
                  <div class="ms-3">
                    <h3 class="fs-5 mb-1">Joge Lucky</h3>
                    <p class="mb-0">Lorem ipsum dolor sit amet consectetur elit.</p>
                  </div>
                </div>
                <div class="admin d-flex align-items-center rounded-2 p-3 mb-4">
                  <div class="img">
                    <img class="img-fluid rounded-pill"
                         width="75" height="75"
                         src="https://uniim1.shutterfly.com/ng/services/mediarender/THISLIFE/021036514417/media/23148907137/small/1501685404/enhance"
                         alt="admin">
                  </div>
                  <div class="ms-3">
                    <h3 class="fs-5 mb-1">Joge Lucky</h3>
                    <p class="mb-0">Lorem ipsum dolor sit amet consectetur elit.</p>
                  </div>
                </div>
                <div class="admin d-flex align-items-center rounded-2 p-3">
                  <div class="img">
                    <img class="img-fluid rounded-pill"
                         width="75" height="75"
                         src="https://uniim1.shutterfly.com/ng/services/mediarender/THISLIFE/021036514417/media/23148907019/small/1501685403/enhance"
                         alt="admin">
                  </div>
                  <div class="ms-3">
                    <h3 class="fs-5 mb-1">Joge Lucky</h3>
                    <p class="mb-0">Lorem ipsum dolor sit amet consectetur elit.</p>
                  </div>
                </div>
              </div>
            </div>
            <div class="col-md-6">
              <div class="box">
                <!-- <h4>Moderators:</h4> -->
                <div class="admin d-flex align-items-center rounded-2 p-3 mb-4">
                  <div class="img">
                    <img class="img-fluid rounded-pill"
                         width="75" height="75"
                         src="https://uniim1.shutterfly.com/ng/services/mediarender/THISLIFE/021036514417/media/23148907114/small/1501685404/enhance"
                         alt="admin">
                  </div>
                  <div class="ms-3">
                    <h3 class="fs-5 mb-1">Joge Lucky</h3>
                    <p class="mb-0">Lorem ipsum dolor sit amet consectetur elit.</p>
                  </div>
                </div>
                <div class="admin d-flex align-items-center rounded-2 p-3 mb-4">
                  <div class="img">
                    <img class="img-fluid rounded-pill"
                         width="75" height="75"
                         src="https://uniim1.shutterfly.com/ng/services/mediarender/THISLIFE/021036514417/media/23148907086/small/1501685404/enhance"
                         alt="admin">
                  </div>
                  <div class="ms-3">
                    <h3 class="fs-5 mb-1">Joge Lucky</h3>
                    <p class="mb-0">Lorem ipsum dolor sit amet consectetur elit.</p>
                  </div>
                </div>
                <div class="admin d-flex align-items-center rounded-2 p-3">
                  <div class="img">
                    <img class="img-fluid rounded-pill"
                         width="75" height="75"
                         src="https://uniim1.shutterfly.com/ng/services/mediarender/THISLIFE/021036514417/media/23148907008/medium/1501685726/enhance"
                         alt="admin">
                  </div>
                  <div class="ms-3">
                    <h3 class="fs-5 mb-1">Joge Lucky</h3>
                    <p class="mb-0">Lorem ipsum dolor sit amet consectetur elit.</p>
                  </div>
                </div>
              </div>
            </div>
          </div>
        </section>

        <section class="statis mt-4 text-center">
          <div class="row">
            <div class="col-md-6 col-lg-3 mb-4 mb-lg-0">
              <div class="box bg-primary p-3">
                <Icon name="Eye"></Icon>
                <h3>5,154</h3>
                <p class="lead">Page views</p>

              </div>
            </div>
            <div class="col-md-6 col-lg-3 mb-4 mb-lg-0">
              <div class="box bg-danger p-3">
                <Icon name="Users"></Icon>
                <h3>245</h3>
                <p class="lead">User registered</p>
              </div>
            </div>
            <div class="col-md-6 col-lg-3 mb-4 mb-md-0">
              <div class="box bg-warning p-3">
                <Icon name="TrendingUp"></Icon>
                <h3>5,154</h3>
                <p class="lead">Product sales</p>
              </div>
            </div>
            <div class="col-md-6 col-lg-3">
              <div class="box bg-success p-3">
               <icon name="Landmark"></icon>
                <h3>5,154</h3>
                <p class="lead">Transactions</p>
              </div>
            </div>
          </div>
        </section>
        <section class="charts mt-4">
          <div class="chart-container p-3">
            <h3 class="fs-6 mb-3">Chart title number three </h3>
            <div style="height: 300px">
              <canvas id="chart3" width="100%"></canvas>
            </div>
          </div>
        </section>
      </div>
    </section>
    <SuperAdminDashChart1></SuperAdminDashChart1>

  </div>

</template>

<script>
import DashStats from "../components/SuperAdminDashStats.vue";
import SuperAdminDashChart1 from "../components/SuperAdminChart1.vue"
import Sidebar from "@/components/Sidebar";
import Icon from "@/components/Icon"
// @ is an alias to /src
export default {
  components: {DashStats, SuperAdminDashChart1, Sidebar, Icon},
  name: 'SuperAdminDashboard',
  props: {
    algo: String
  },


}

</script>
<style scoped>
@import 'https://fonts.googleapis.com/css2?family=Inter:wght@300;400;500;600&display=swap" rel="stylesheet';

:root {
  --dk-gray-100: #F3F4F6;
  --dk-gray-200: #E5E7EB;
  --dk-gray-300: #D1D5DB;
  --dk-gray-400: #9CA3AF;
  --dk-gray-500: #6B7280;
  --dk-gray-600: #4B5563;
  --dk-gray-700: #374151;
  --dk-gray-800: #1F2937;
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
  font-family: 'Inter', sans-serif;
  background-color: #2a2b3d;
  font-size: .925rem;
}

.col-lg-4 {
  width: 30%;
}

.statistics .row {
  justify-content: space-between
}

#wrapper {
  margin-left: 250px;
  transition: all .3s ease-in-out;
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
  transition: all .3s ease-in-out;
  transform: translateX(0);
  z-index: 9999999
}

.sidebar .close-aside {
  position: absolute;
  top: 7px;
  right: 7px;
  cursor: pointer;
  color: #EEE;
}

.sidebar .sidebar-header {
  border-bottom: 1px solid #2a2b3c
}

.sidebar .sidebar-header h5 a {
  color: #D1D5DB
}

.sidebar .sidebar-header p {
  color: #9CA3AF;
  font-size: .825rem;
}

.sidebar .search .form-control ~ i {
  color: #2b2f3a;
  right: 40px;
  top: 22px;
}

.sidebar > ul > li {
  padding: .7rem 1.75rem;
}

.sidebar ul > li > a {
  color: #9CA3AF;
  text-decoration: none;
}

/* Start numbers */
.sidebar ul > li > a > .num {
  line-height: 0;
  border-radius: 3px;
  font-size: 14px;
  padding: 0px 5px
}

.sidebar ul > li > i {
  font-size: 18px;
  margin-right: .7rem;
  color: #6B7280;
}

.sidebar ul > li.has-dropdown > a:after {
  content: '\eb3a';
  font-family: unicons-line;
  font-size: 1rem;
  line-height: 1.8;
  float: right;
  color: #6B7280;
  transition: all .3s ease-in-out;
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
  font-size: .85rem;
  padding: .5rem 0;
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
    padding: 10px 0 10px 30px
  }
}


/** --------------------------------
 -- welcome
-------------------------------- */
.welcome {
  color: #D1D5DB;
}

.welcome .content {
  background-color: #313348;
}

.welcome p {
  color: #9CA3AF;
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
  color: #9CA3AF
}


/** --------------------------------
 -- users
-------------------------------- */
.admins .box .admin {
  background-color: #313348;
}

.admins .box h3 {
  color: #D1D5DB;
}

.admins .box p {
  color: #9CA3AF
}


/** --------------------------------
 -- statis
-------------------------------- */
.statis {
  color: #F3F4F6;
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
  color: #ffc107
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
  color: #EEE !important;
  line-height: 55px !important;
  padding: 0 10px !important;
}

.navbar .navbar-brand {
  color: #FFF !important
}

.navbar .navbar-nav > li > a:focus,
.navbar .navbar-nav > li > a:hover {
  color: #EEE !important
}

.navbar .navbar-nav > .open > a,
.navbar .navbar-nav > .open > a:focus,
.navbar .navbar-nav > .open > a:hover {
  background-color: transparent !important;
  color: #FFF !important
}

.navbar .navbar-brand {
  line-height: 55px !important;
  padding: 0 !important
}

.navbar .navbar-brand:focus,
.navbar .navbar-brand:hover {
  color: #FFF !important
}

.navbar > .container .navbar-brand, .navbar > .container-fluid .navbar-brand {
  margin: 0 !important
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
  right: 5px !important
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
    margin: 0 !important
  }

  .statistics .box {
    margin-bottom: 25px !important;
  }

  .navbar .navbar-nav .open .dropdown-menu > li > a {
    color: #CCC !important
  }

  .navbar .navbar-nav .open .dropdown-menu > li > a:hover {
    color: #FFF !important
  }

  .navbar .navbar-toggle {
    border: none !important;
    color: #EEE !important;
    font-size: 18px !important;
  }

  .navbar .navbar-toggle:focus, .navbar .navbar-toggle:hover {
    background-color: transparent !important
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