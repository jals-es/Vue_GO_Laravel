<template>
  <div class="main">
    <div class="text-center">
      <main class="form-signin">
        <form novalidate @submit.prevent="submit">
          <h1 class="h3 mb-3 fw-normal">Welcome to AppBar!</h1>
          <div class="form-floating">
            <input type="text" class="form-control" id="floatingNameInput" placeholder="John Wick" v-model="name">
            <label for="floatingNameInput">Full Name</label>
          </div>
          <div class="form-floating">
            <input type="email" class="form-control" id="floatingMailInput" placeholder="name@example.com"
                   v-model="email">
            <label for="floatingMailInput">Email address</label>
          </div>
          <div class="form-floating">
            <input type="password" class="form-control" id="floatingPassword" placeholder="Password"
                   v-model="password">
            <label for="floatingPassword">Password</label>
          </div>

          <div class="mb-3">
            <p>Already have an account?
              <router-link to="/login">Login Here</router-link>
            </p>
          </div>
          <button class="w-100 btn btn-lg btn-primary" type="submit">Sign in</button>
        </form>
      </main>
    </div>
  </div>
</template>

<script>
import golangApiService from "@/core/http/golang.api.service";

export default {
  name: "Register",
  data() {
    return {
      name: '',
      email: '',
      password: ''
    }
  },
  methods: {
    submit(){
      const user = {
        users: {
          name: this.name,
          email: this.email,
          passwd: this.password
        }
      }
      golangApiService.post('/api/user', user)
    }
  }
}
</script>

<style scoped>
.form-signin {
  width: 450px;
  padding: 15px;
  margin: auto;
}

.form-signin .checkbox {
  font-weight: 400;
}

.form-signin .form-floating:focus-within {
  z-index: 2;
}

.form-signin input[type="text"] {

  border-bottom-left-radius: 0;
  border-bottom-right-radius: 0;
  margin-bottom: -1px;
}

.form-signin input[type="email"] {
  margin-bottom: -1px;
  border-radius: 0;
}

.form-signin input[type="password"] {
  margin-bottom: 10px;
  border-top-left-radius: 0;
  border-top-right-radius: 0;
}

.main {
  display: flex;
  align-items: center;
  justify-content: center;
  padding-top: 40px;
  padding-bottom: 40px;
  height: 100vh;
}
</style>
