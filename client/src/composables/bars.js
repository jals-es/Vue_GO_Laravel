import golangApiService from "@/core/http/golang.api.service";
import { reactive } from 'vue';
export function bars() {

    const barsInfo = reactive({ thisBar: {} })

    const getBarInfo = async(slug) => {
        let getBar = await golangApiService.get("/api/bar/" + slug)
        if (getBar) {
            Object.assign(barsInfo.thisBar, getBar.data.bar)
        }
    }

    return {
        getBarInfo,
        barsInfo
    }
}