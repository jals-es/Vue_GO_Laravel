<template>
    <nav id="sidebarMenu" class="col-md-3 col-lg-2 d-md-block bg-light sidebar collapse">
      <div class="position-sticky pt-3">
        <ul class="nav flex-column">
          <li class="nav-item">
            <router-link :to="`/${barSlug}/admin/`" class="nav-link">Productos</router-link>
          </li>
          <li class="nav-item">
            <router-link :to="`/${barSlug}/admin/product-create/`" class="nav-link">Crear Producto</router-link>
          </li>
        </ul>
      </div>
    </nav>
</template>

<script>
export default {
  name: "DashboardMenu",
  props: {
    barSlug: String
  },
  setup() {
    // document.getElementById(`menu${active}`).classList.add("active");
  },
};
</script>
