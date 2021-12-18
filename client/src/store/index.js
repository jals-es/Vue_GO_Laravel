import Vuex from "vuex";

import { superUser } from './superUser'
import { userStore } from "@/store/user";
import { barStore } from '@/store/bar';

export default new Vuex.Store({
    modules: {
        superUser,
        userStore,
        barStore
    }

})