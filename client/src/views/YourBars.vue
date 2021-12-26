<template>
    <div class="cerrarSession mx-3 p-1" @click="logOut">Salir</div>
    <div class="createBar mt-3 mx-auto py-1 rounded-circle" @click="createBars">+</div>
    <BarAdminListBars :bars="state.bars.bars"></BarAdminListBars>
</template>

<script>
import BarAdminListBars from '@/components/bars/BarAdmin_ListBars.vue'
import router from '@/router'
import { computed, reactive } from "vue";
import { useStore } from "vuex";
export default {
    components: { BarAdminListBars },
    setup() {
        // Cridem al store
        const store = useStore();

        // Omplim el state
        store.dispatch("barStore/getBars");

        //Agafem els datos
        const state = reactive({
            bars: computed(() => store.getters["barStore/getBars"])
        });

        function createBars(){
            router.push("/bars/createBars")
        }

        function logOut(){
            localStorage.removeItem("token")
            router.push("/login")
        }

        return {
            state,
            createBars,
            logOut
        };
    }    
}
</script>


<style>
.cerrarSession{
    position: absolute;
    text-transform: uppercase;
    cursor: pointer;
    border: 2px solid;
}

.cerrarSession:hover{
    background-color: black;
    color: white;
    border: 2px solid black;
}

.createBar{
    width: 60px;
    height: 60px;
    text-align: center;
    border: 2px solid;
    font-size: 30px;
    cursor: pointer;
    transition: 0.5s;
}

.createBar:hover{
    background-color: black;
    color: white;
    border: 2px solid black;
    transform: scale(1.2,1.2);
}

.bars{
    width: 80%;
    margin: 0px auto;
}
.bar{
    width: 40%;
    border: 2px solid black;
    border-radius: 20px;
    margin: 20px auto;
    padding: 2px 4px;
    cursor: pointer;
    display: grid; 
    grid-template-columns: 1fr 1fr 1fr; 
    grid-template-rows: 1fr 1fr; 
    gap: 0px 0px; 
    grid-template-areas: 
        "name . city"
        "rol rol rol"; 
    text-align: center;
    text-decoration: none;
    color: black;
    transition: 0.5s;
}

.bar:hover{
    background-color: #eee;
    color: black;
    border: 2px solid red;
}

.bar .name{
    grid-area: name;
    font-size: 20px;
}

.bar .city{
    grid-area: city;
}

.bar .rol{
    grid-area: rol;
    text-transform: Capitalize;
}

@media (max-width: 1200px){
    .bar{
        width: 80%;
    }
}

@media (max-width: 700px){
    .bar{
        width: 100%;
        grid-template-columns: 1fr; 
        grid-template-rows: 1fr 1fr 1fr; 
        grid-template-areas: 
            "name"
            "city"
            "rol";
    }
}
</style>
