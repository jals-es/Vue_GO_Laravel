import golangApiService from "@/core/http/golang.api.service";
import { reactive } from 'vue';
export function prods() {

    const products = reactive({ prods: [] })

    const getProds = async(slug) => {
        let getProds = await golangApiService.get("/api/prod/" + slug + "/")
        if (getProds) {
            products.prods = getProds.data.prods
        }
    }

    const deleteProd = async(slug, bar_slug) => {
        let delProds = await golangApiService.delete("/api/prod/" + slug + "/")
        if (delProds) {
            console.log(delProds.data.message)
            console.log(bar_slug)
                // getProds(bar_slug)
        }
    }

    return {
        getProds,
        deleteProd,
        products
    }
}