import Vuex from "vuex";

import { superUser } from './superUser'
import {userStore} from "@/store/user";

export default new Vuex.Store({
    modules: {
        superUser,
        userStore
    }

})
