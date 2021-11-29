<template>
    <section class="statistics mt-4">
      <div class="row">
        <div class="col-lg-4">
          <div class="box d-flex rounded-2 align-items-center mb-4 mb-lg-0 p-3">
            <i class="uil-envelope-shield fs-2 text-center bg-primary rounded-circle"></i>
            <div class="ms-3">
              <div class="d-flex align-items-center">
                <h3 class="mb-0">300</h3> <span class="d-block ms-2">Pedidos</span>
              </div>
              <p class="fs-normal mb-0">Realizados en las ultimas 24 horas</p>
            </div>
          </div>
        </div>
        <div class="col-lg-4">
          <div class="box d-flex rounded-2 align-items-center mb-4 mb-lg-0 p-3">
            <i class="uil-file fs-2 text-center bg-danger rounded-circle"></i>
            <div class="ms-3">
              <div class="d-flex align-items-center">
                <h3 class="mb-0">478</h3> <span class="d-block ms-2">Pedidos</span>
              </div>
              <p class="fs-normal mb-0">Realizados en los ultimos 30 dias</p>
            </div>
          </div>
        </div>
        <div class="col-lg-4">
          <div class="box d-flex rounded-2 align-items-center p-3">
            <i class="uil-users-alt fs-2 text-center bg-success rounded-circle"></i>
            <div class="ms-3">
              <div class="d-flex align-items-center">
                <h3 class="mb-0">104</h3> <span class="d-block ms-2">Bares</span>
              </div>
              <p class="fs-normal mb-0">Han registrado un pedido hoy</p>
            </div>
          </div>
        </div>
      </div>
    </section>        
</template>
    <script>
export default {
  name: "DashStats",
};
</script>
<style scoped>
.statistics {
  color: #e5e7eb;
}

.statistics .box {
  background-color: #313348;
}

.statistics .box i {
  width: 60px;
  height: 60px;
  line-height: 60px;
}

.statistics .box p {
  color: #9ca3af;
}
</style>