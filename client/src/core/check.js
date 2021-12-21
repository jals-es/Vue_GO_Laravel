import router from "@/router";
import { computed, nextTick } from "vue";
import { useStore } from "vuex";

checkAdmin = () => {
    if (localStorage.getItem("token")) {
        // Cridem al store
        const store = useStore();

        // Omplim el state
        store.dispatch("barStore/getBars");

        //Agafem els datos
        const user = computed(() => store.getters["barStore/getBars"]);

        console.log(user);

        if (user) {
            return next();
        }
    }

    router.push("/login");
    return false;
}