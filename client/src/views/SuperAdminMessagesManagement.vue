<template>
  <div>
    <SuperAdminListMessages :key="toReload.toReload" :messages="messages"></SuperAdminListMessages> 
    <SuperAdminSendMessage></SuperAdminSendMessage>
  </div>
</template>
<script>
import SuperAdminSendMessage from "../components/SuperAdminSendMessage.vue";
import SuperAdminListMessages from "../components/SuperAdminListMessages.vue";
import Pusher from "pusher-js"; // import Pusher
import { reactive } from "vue";

export default {
  components: { SuperAdminSendMessage, SuperAdminListMessages },
  setup() {
    Pusher.logToConsole = true;
    var messages = [];
    const toReload = reactive({ // Creem un component reactiu per a que vue siga capaç de detectar els canvis sols ja que el que retorna pusher no es reactiu
      toReload: 0,
    });
    var pusher = new Pusher("3b155703aa62c7c63521", { // Creem el pusher
      cluster: "eu",
    });
    var channel = pusher.subscribe("messages"); // Assignem pusher al canal de messages
    channel.bind("App\\Events\\NewMessage", function (data) { // Assignem pusher als events de "App\\Events\\NewMessage"
      messages.push(data); // Actualitzem les dades 
      toReload.toReload++; // Canviem el component reactiu per a que ho detecte Vue i torne a renderitzar el component fill amb els nous datos de vue
    });
    return {
      messages,
      toReload
    };
  },
};
</script>