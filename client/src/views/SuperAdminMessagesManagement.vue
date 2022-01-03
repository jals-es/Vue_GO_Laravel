<template>
    <p v-for="msg in messages">{{msg}}</p>
</template>
  <script src="https://js.pusher.com/7.0/pusher.min.js"></script>
  <script src="https://cdn.jsdelivr.net/npm/vue/dist/vue.js"></script>
<script>

export default {
    
    // Vue application
    const app = new Vue({
      el: '#app',
      data: {
        messages: [],
      },
    });
    setup() {
        Pusher.logToConsole = true;
        var messages = [];
        var pusher = new Pusher('3b155703aa62c7c63521', {
            cluster: 'eu'
        });

        var channel = pusher.subscribe('messages');
        channel.bind('newMessage', function(data) {
          console.log(JSON.stringify(data));
          messages.push(JSON.stringify(data));
        });
        return {
            messages
        }    
    },
}
</script>