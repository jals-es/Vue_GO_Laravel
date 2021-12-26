// import router from "@/router";
// import { computed } from "vue";
// import { useStore } from "vuex";
import golangApiService from "@/core/http/golang.api.service";

export default {
    checkAdmin(to, from, next) {
        console.log(to);

        if (localStorage.getItem("token")) {

            golangApiService.get('/api/user/')
                .then(({ data }) => {
                    console.log(data);
                    console.log('correcte');
                    if (data.data == "Usuario verificado") {
                        next();
                    } else {
                        next("/login");
                    }
                })
                .catch((error) => {
                    console.log(error)
                    next("/login");
                })

        } else {
            console.log('no troba token');
            next("/login");
        }
    }
}