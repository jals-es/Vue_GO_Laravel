<template>
  <div class="main">
    <div class="text-center">
      <main class="form-signin">
        <form @submit.prevent="submit">
          <h1 class="h3 mb-3 fw-normal">Welcome to AppBar!</h1>
          <div class="form-floating">
            <input type="email" class="form-control"
                   :class="{'is-invalid': this.v$.user.email.$error}"
                   id="floatingMailInput" placeholder="name@example.com"
                   v-model="state.user.email">
            <label for="floatingMailInput">Email address</label>
            <div class="invalid-feedback" v-if="v$.user.email.$error">
              <p v-for="(error, index) in this.v$.user.email.$errors" :key="index">
                {{ error.$message }}
              </p>
            </div>
          </div>
          <div class="form-floating">
            <input type="password" class="form-control"
                   @change.capture="this.v$.$touch()"
                   id="floatingPassword" placeholder="Password"
                   v-model="state.user.passwd">
            <label for="floatingPassword">Password</label>
            <div class="invalid-feedback" v-if="v$.user.passwd.$error">
              <p v-for="(error, index) in this.v$.user.passwd.$errors" :key="index">
                {{ error.$message }}
              </p>
            </div>
          </div>
          <div class="my-3">
            <p>Don't have an account?
              <router-link to="/register">Register Here</router-link>
            </p>
          </div>
          <button class="w-100 btn btn-lg btn-primary" type="submit">Sign In</button>
        </form>
      </main>
    </div>
  </div>
  <Alert :alert-data="state.alertData" v-if="state.alertData.open"></Alert>
</template>

<script>
import {useStore} from "vuex";
import useVuelidate from '@vuelidate/core'
import {reactive} from "vue";
import {email, required} from "@vuelidate/validators"
import Alert from '@/components/Alert'

export default {
  name: "Login",
  components: {Alert},
  setup() {
    const store = useStore()
    const state = reactive({
      user: {
        email: '',
        passwd: ''
      },
      alertData: {
        open: false,
        status: 0,
        message: 'Login Successful'
      }
    })

    const rules = {
      user: {
        email: {required, email},
        passwd: {required}
      }
    }
    const v$ = useVuelidate(rules, state)

    function submit() {
      this.v$.$validate()
      if (!this.v$.$error) {
        store.dispatch('userStore/loginUser', state.user)
            .then(data => {
                  console.log(data)
                  store.commit('userStore/fillUser', data.data.user)
                  state.alertData.open = true
                  state.alertData.status = data.status
                }
            )
      }

    }

    return {
      state,
      v$,
      submit
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


.form-signin input[type="email"] {
  margin-bottom: -1px;
  border-radius: 0.25rem 0.25rem 0 0;
}

.form-signin input[type="password"] {
  margin-bottom: -1px;
  border-radius: 0 0 0.25rem 0.25rem;
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
