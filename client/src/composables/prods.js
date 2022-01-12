import golangApiService from "@/core/http/golang.api.service";
import router from '@/router'
import { reactive } from 'vue';
export function prods() {

    const products = reactive({ prods: [] })

    const getProds = async(slug) => {
        let getProds = await golangApiService.get("/api/prod/" + slug + "/")
        if (getProds) {
            products.prods = getProds.data.prods
        }
    }

    const deleteProd = async(slug, bar_slug, index) => {
        let delProds = await golangApiService.delete("/api/prod/" + slug + "/")
        if (delProds) {
            console.log(delProds.data.message)
            console.log(bar_slug)
            products.prods.splice(index, 1)
        }
    }

    const createProd = async(prod, bar_slug) => {
        let create = await golangApiService.post("/api/prod/", { barslug: bar_slug, prod: prod })
        if (create.data.message) {
            products.prods.push(prod)
            router.push(`/${bar_slug}/admin/`)
        }
    }

    return {
        getProds,
        deleteProd,
        createProd,
        products
    }
}