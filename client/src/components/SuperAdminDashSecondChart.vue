<template>
  <canvas id="dashChart2"></canvas>
</template>

<script>
import Chart from "chart.js";
import { onMounted, computed, watch } from "vue";
import { useStore } from "vuex";
export default {
  name: "SuperAdminDashSecondChart",
  setup() {
    // Retornem els datos
    const store = useStore();

    // Omplim el state
    store.dispatch("superUser/getSecondChart");

    //Agafem els datos
    var interruptor = true;
    const secondChart = computed(
      () => store.getters["superUser/getSecondChart"]
    );
    function random_rgba() {
      var o = Math.round,
        r = Math.random,
        s = 255;
      return (
        "rgba(" +
        o(r() * s) +
        "," +
        o(r() * s) +
        "," +
        o(r() * s) +
        "," +
        r().toFixed(1) +
        ")"
      );
    }
    console.log(random_rgba())
    onMounted(() => {
      // Cridem al store

      // Wait for the proxy to retrive the store data
      watch(
        () => secondChart.value,
        (first) => {
          // When proxy retrieves the data we change the chart value and update it (only for the first time to evade reactivity to fire this multiple times)
          if (interruptor) {
            interruptor = false;
            var newData = Object.values(first.total);
            console.log(newData, Object.values(first.total));
            chart.data.labels = Object.values(first.days);
            chart.data.datasets[0] = {
              label: "Pedidos totales por dia de la semana",
              data: newData,
              backgroundColor: [random_rgba(),random_rgba(),random_rgba(),random_rgba(),random_rgba(),random_rgba(),random_rgba()],
            };
            chart.update();
          }
        }
      );

      // Initialize an empty chart
      const ctx = document.getElementById("dashChart2");
      const chart = new Chart(ctx, {
        type: "pie",
        data: {
          labels: [],
          datasets: [],
        },
        options: {
          responsive: true,
          plugins: {
            legend: {
              position: "top",
            },
            title: {
              display: true,
              text: "Chart.js Pie Chart",
            },
          },
        },
      });
    });
  },
};
</script>