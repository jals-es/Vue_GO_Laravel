<template>
  <div class="card" style="border-radius: 15px; flex-basis: 45%; margin-bottom: 10px">
    <div
      class="bg-image hover-overlay ripple ripple-surface ripple-surface-light"
      data-mdb-ripple-color="light"
    >
      <img
        v-bind:src="product.photo"
        style="border-top-left-radius: 15px; border-top-right-radius: 15px"
        class="img-fluid"
      />
    </div>
    <div class="card-body pb-0">
      <div class="d-flex justify-content-between">
        <div>
          <p class="text-dark h4">{{ product.name }}</p>
          <p class="text-dark">{{ product.descr }}</p>
        </div>
      </div>
    </div>
    <hr class="my-0" />
    <div class="card-body pb-0">
      <p class="h3">Tipos</p>
      <div
        class="d-flex justify-content-between"
        v-for="type in product.types"
        :key="type.id"
      >
        <p class="text-dark h6">{{ type.name }}</p>
        <p class="text-dark">{{ type.descr }}</p>
        <p class="text-dark">{{ type.price }}€</p>
      </div>
    </div>
    <hr class="my-0" v-if="product.extras.length > 0" />
    <div class="card-body" v-if="product.extras.length > 0">
      <p class="h3">Extras</p>
      <div
        class="d-flex justify-content-between align-items-center"
        v-for="extra in product.extras"
        :key="extra.id"
      >
        <p class="h6">{{extra.name}}</p>
        <p>{{extra.price}}€</p>
      </div>
    </div>
  </div>
</template>
<script>
export default {
  name: "SuperAdminProduct",
  props: {
    product: Object,
  },
  setup(product) {
    console.log(product);
  },
};
</script>