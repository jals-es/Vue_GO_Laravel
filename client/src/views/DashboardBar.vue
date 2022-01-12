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
            <tr class="prod" v-for="(prod, index) in products.prods" :key="prod.slug">
              <td>{{ prod.slug }}</td>
              <td>{{ prod.name }}</td>
              <td>{{ prod.descr }}</td>
              <td>{{ prod.types.length }}</td>
              <td>{{ prod.extras.length }}</td>
              <td><button class="btn btn-success" data-bs-toggle="modal" :data-bs-target="`#${prod.slug}`">Ver</button></td>
              <td>
                <button class="btn btn-outline-danger" @click="delProd(prod.slug, bar_slug, index)">Borrar</button>
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

        const delProd = (slug, bar_slug, index) => {
            deleteProd(slug, bar_slug, index)
        }

        return {
            bar_slug,
            products,
            delProd
        }
    }
}
</script>

<style>



.prod td{
    width: 20%;
}

.prod td img{
  width: 100%;
  height: auto;
}
</style>