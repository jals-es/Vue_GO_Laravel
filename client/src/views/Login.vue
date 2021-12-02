<template>
    <div class="wrapper">
        <div class="container">
            <h1>Welcome to APPBAR</h1>
            
            <form class="form">
                <input type="email" v-model="email" placeholder="Email">
                <input type="password" v-model="password" placeholder="Password">
                <button type="submit" id="login-button">Login</button>
                <p>Are you already registered? <router-link to="/register">SignUp</router-link></p>
            </form>
        </div>
        
        <ul class="bg-bubbles">
            <li></li>
            <li></li>
            <li></li>
            <li></li>
            <li></li>
            <li></li>
            <li></li>
            <li></li>
            <li></li>
            <li></li>
        </ul>
    </div>
</template>

<script>
// import { mapActions } from 'vuex'
import useVuelidate from '@vuelidate/core'
import { email, required, minLength } from '@vuelidate/validators'
// import { getStore } from  '@/store/'
// import router from '@/router/'
export default {
    
    name: 'login',
    setup() {
        return { v$: useVuelidate() }
    },
    data(){
        return {
            email: '',
            password: ''
        }
    },
    validations() {
        return {
            email: { email, required },
            password: { required, minLength: minLength(6) }
        };
    },
    methods: {
        onSubmit() {
            this.v$.$touch();
            if (this.v$.$error) {
                console.log("hola"); 
                return;
            }
            console.log("entra");
            return

            // if (this.v$.$error) return;
            // let user = getStore('user')
            // console.log(user)
            // if (user) {
            //     if (user.email == this.email && user.password == this.password) {
            //     this.userLogin(user)
            //     router.push('/todo')
            //     return
            //     }
            // }
            // this.loginFailed = true
        }
    }
}
</script>

<style>
@import '../assets/styles/login.css';
</style>