<template>
  <canvas id="dashChart1"></canvas>
</template>

<script>
import Chart from "chart.js";
import { onMounted, computed, watch } from "vue";
import { useStore } from "vuex";
export default {
  name: "SuperAdminDashChart1",
  setup() {
    // const firstChart = ref(() => store.getters["superUser/getFirstChart"])

    // Retornem els datos
    const store = useStore();

    // Omplim el state
    store.dispatch("superUser/getFirstChart");

    //Agafem els datos
    var interruptor = true;
    const firstChart = computed(() => store.getters["superUser/getFirstChart"]);
    onMounted(() => {
      // Cridem al store

      // Wait for the proxy to retrive the store data
      watch(
        () => firstChart.value,
        (first) => {
          // When proxy retrieves the data we change the chart value and update it (only for the first time to evade reactivity to fire this multiple times)
          if (interruptor) {
            interruptor = false;
            var newData = Object.values(first.orders);
            chart.data.labels = Object.values(first.names);
            chart.data.datasets[0] = {
              label: "Pedidos por bar",
              data: newData,
              backgroundColor: "rgba(54,73,93,.5)",
              borderColor: "#36495d",
              borderWidth: 3,
            };
            chart.update();
          }
        }
      );

      // Initialize an empty chart
      const ctx = document.getElementById("dashChart1");
      const chart = new Chart(ctx, {
        type: "line",
        data: {
          labels: [],
          datasets: [],
        },
        options: {
          responsive: true,
          lineTension: 1,
          scales: {
            yAxes: [
              {
                ticks: {
                  beginAtZero: true,
                  padding: 25,
                },
              },
            ],
          },
        },
      });
    });
  },
};

// const ctx = document.getElementById('dashChart1');
// // const exampleChart = new Chart(ctx, {
// //   type: '',
// //   data: [],
// //   options: {},
// // });
</script>