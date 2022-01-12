<template>
<!-- Header -->
<DashboardHeader :barSlug="bar_slug"></DashboardHeader>
<div class="container-fluid">
  <div class="row">
    <!-- Menu -->
    <DashboardMenu :barSlug="bar_slug"></DashboardMenu>

    <main class="col-md-9 ms-sm-auto col-lg-10 px-md-4">
      <h2 class="mt-3">Productos</h2>
      <div class="table-responsive">
        <table class="table table-striped table-sm">
          <thead>
            <tr>
              <th scope="col">#</th>
              <th scope="col">Nombre</th>
              <th scope="col">Descripcion</th>
              <th scope="col">Tipos</th>
              <th scope="col">Extras</th>
            </tr>
          </thead>
          <tbody>
            <tr class="prod" v-for="prod in products.prods" :key="prod.slug" @click="goProd(prod.slug)">
              <td>{{ prod.slug }}</td>
              <td>{{ prod.name }}</td>
              <td>{{ prod.descr }}</td>
              <td>{{ prod.types.length }}</td>
              <td>{{ prod.extras.length }}</td>
              <td><button class="btn btn-success" data-bs-toggle="modal" :data-bs-target="`#${prod.slug}`">Ver</button></td>
              <td>
                <button class="btn btn-outline-danger" @click="delProd(prod.slug, bar_slug)">Borrar</button>
              </td>
              <DashboardViewProd :prod="prod"></DashboardViewProd>
            </tr>
            <tr v-if="products.prods.length == 0">
              <td colspan="7" class="text-center">No tienes productos</td>
            </tr>
          </tbody>
        </table>
      </div>
    </main>
  </div>
</div>
</template>

<script>
import { useRoute } from "vue-router";
import { prods } from '@/composables/prods'
import DashboardMenu from '@/components/dashboardAdmin/BarAdmin_DashboardMenu'
import DashboardHeader from '@/components/dashboardAdmin/BarAdmin_DashboardHeader'
import DashboardViewProd from '@/components/dashboardAdmin/BarAdmin_DashboardViewProd'
export default {
    components: {
      DashboardMenu,
      DashboardHeader,
      DashboardViewProd
    },
    setup() {
        const route = useRoute();
        const bar_slug = route.params.slug;

        const { getProds, products, deleteProd } = prods();

        getProds(bar_slug)

        const goProd = (/*prod_slug*/) => {
            // router.push(`/${bar_slug}/admin/${prod_slug}/`)
        }

        const delProd = (slug, bar_slug) => {
            deleteProd(slug, bar_slug)
        }

        return {
            bar_slug,
            products,
            goProd,
            delProd
        }
    }
}
</script>

<style>

/*
 * Sidebar
 */

.sidebar {
  position: fixed;
  top: 0;
  /* rtl:raw:
  right: 0;
  */
  bottom: 0;
  /* rtl:remove */
  left: 0;
  z-index: 100; /* Behind the navbar */
  padding: 48px 0 0; /* Height of navbar */
  box-shadow: inset -1px 0 0 rgba(0, 0, 0, .1);
}

@media (max-width: 767.98px) {
  .sidebar {
    top: 5rem;
  }
}

.sidebar-sticky {
  position: relative;
  top: 0;
  height: calc(100vh - 48px);
  padding-top: .5rem;
  overflow-x: hidden;
  overflow-y: auto; /* Scrollable contents if viewport is shorter than content. */
}

.sidebar .nav-link {
  font-weight: 500;
  color: #333;
}

.sidebar .nav-link i {
  margin-right: 4px;
  padding: 20px;
  color: #727272;
}

.sidebar .nav-link.active {
  color: #2470dc;
}

.sidebar .nav-link:hover .feather,
.sidebar .nav-link.active .feather {
  color: inherit;
}

.sidebar-heading {
  font-size: .75rem;
  text-transform: uppercase;
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

.prod td{
    width: 20%;
}

.prod td img{
  width: 100%;
  height: auto;
}
</style>