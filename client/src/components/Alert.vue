<template>
  <div class="d-flex position-absolute w-100 justify-content-end top-0">
    <div class="alert d-flex mx-4"
         :class="{'alert-success': this.alertData.status === 201, 'alert-danger': this.alertData.status !== 201}"
         role="alert">
    <Icon name="Check"></Icon>
    <div>
      {{ this.alertData.message }}
    </div>
  </div>
  </div>

</template>

<script>

import Icon from "@/components/Icon"

export default {
  name: "Alert",
  components: {Icon},
  props: ['alertData']

}

</script>

<style scoped>
.alert {
  position: absolute;
  top: 10px;
  z-index: 5;
}
</style>
