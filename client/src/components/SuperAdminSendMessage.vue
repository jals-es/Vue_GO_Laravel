<template>
  <div>
    <div class="input-group">
      <input
        v-model="message"
        id="message"
        type="text"
        name="message"
        placeholder="Type Message ..."
        class="form-control"
      />
      <span class="input-group-btn">
        <button
          @click="sendMessage()"
          type="button"
          class="btn btn-warning btn-flat"
        >
          Send
        </button>
      </span>
    </div>
  </div>
</template>
<script>
import { useStore } from "vuex";
import { reactive, toRefs } from "vue";

export default {
  setup() {
    const store = useStore();
    const message = reactive({
      message: "",
    });
    function sendMessage() {
      store.dispatch("superUser/newMessage", {
        content: message.message,
      });
    }
    return {
      message,
      sendMessage,
      ...toRefs(message),
    };
  },
};
</script>