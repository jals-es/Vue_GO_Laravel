import golangApiService from "@/core/http/golang.api.service";
import { reactive } from 'vue';
export function prods() {

    const products = reactive({ prods: [] })

    const getProds = async(slug) => {
        let getProds = await golangApiService.get("/api/prod/" + slug + "/")
        if (getProds) {
            console.log(getProds.data.prods)
            Object.assign(products.prods, getProds.data.prods)
        }
    }

    return {
        getProds,
        products
    }
}