<template>
    <div class="wrapper">
        <div class="container">
            <h1>Sign Up in APPBAR</h1>
            
            <form class="form">
                <input type="text" placeholder="Name">
                <input type="email" placeholder="Email">
                <input type="password" placeholder="Password">
                <input type="password" placeholder="Repeat Password">
                <button type="submit" id="login-button">Register</button>
                <p>Are you already registered? <router-link to="/login">LogIn</router-link></p>
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
export default {
    
    
}
</script>

<style>
@import '../assets/styles/login.css';
</style>