import laravelApiService from "@/core/http/laravel.api.service";

export default {
    checkSuperAdmin(to, from, next) {
        console.log(to);
        if (localStorage.getItem("jwt")) {
            laravelApiService.get('api/check/')
                .then(({ data }) => {
                    if (data.data == "authorized") {
                        next();
                    } else {
                        localStorage.removeItem("jwt");
                        next("/login");
                    }
                })
                .catch(() => {
                    localStorage.removeItem("jwt");
                    next("/login");
                })
        } else {
            next("/login");
        }
    }
}