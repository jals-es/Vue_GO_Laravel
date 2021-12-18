<template>
    <div class="bars">
      <BarAdminBar v-for="bar in bars" :key="bar.slug" :bar="bar"></BarAdminBar>
    </div>
</template>

<script>
import BarAdminBar from "@/components/bars/BarAdmin_Bar.vue"
export default {
  components: { BarAdminBar },
  name: "ListBars",
  props: {
    bars: Array
  },
  setup() {
    //   console.log(bar);
  },
};
</script>
