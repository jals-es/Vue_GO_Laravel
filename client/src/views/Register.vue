<template>
  <div class="main">
    <div class="text-center">
      <main class="form-signin">
        <form @submit.prevent="submit">
          <h1 class="h3 mb-3 fw-normal">Welcome to AppBar!</h1>
          <div class="form-floating">
            <input type="text" class="form-control"
                   :class="{'is-invalid': this.v$.form.name.$error, 'is-valid': !this.v$.form.name.$error && !this.v$.form.name.$invalid}"
                   id="floatingNameInput" placeholder="John Wick"
                   v-model="state.form.name">
            <label for="floatingNameInput">Full Name</label>
            <div class="invalid-feedback" v-if="this.v$.form.name.$error">
              <p v-for="(error, index) in this.v$.form.name.$errors" :key="index">
                {{ error.$message }}
              </p>
            </div>
          </div>
          <div class="form-floating">
            <input type="email" class="form-control"
                   :class="{'is-invalid': this.v$.form.email.$error}"
                   id="floatingMailInput" placeholder="name@example.com"
                   v-model="state.form.email">
            <label for="floatingMailInput">Email address</label>
            <div class="invalid-feedback" v-if="v$.form.email.$error">
              <p v-for="(error, index) in this.v$.form.email.$errors" :key="index">
                {{ error.$message }}
              </p>
            </div>

            <input type="password" class="form-control"
                   :class="{'is-invalid': this.v$.form.passwd.$error, 'is-valid': !this.v$.form.passwd.$error && !this.v$.form.passwd.$invalid}"
                   @change.capture="this.v$.$touch()"
                   id="floatingPassword" placeholder="Password"
                   v-model="state.form.passwd" @change="this.v$.$touch()">
            <label for="floatingPassword">Password</label>
            <div class="invalid-feedback" v-if="v$.form.passwd.$error">
              <p v-for="(error, index) in this.v$.form.passwd.$errors" :key="index">
                {{ error.$message }}
              </p>
            </div>
          </div>
          <div class="form-floating">
            <input type="password" class="form-control last"
                   @change.capture="this.v$.$touch()"
                   :class="{'is-invalid': this.v$.form.passwdCheck.$error, 'is-valid': !this.v$.form.passwdCheck.$error && !this.v$.form.passwdCheck.$invalid}"
                   id="floatingPasswordCheck" placeholder="Password"
                   v-model="state.form.passwdCheck" @change="this.v$.$touch()">
            <label for="floatingPasswordCheck">Confirm Password</label>
            <div class="invalid-feedback" v-if="v$.form.passwdCheck.$error">
              <p v-for="(error, index) in this.v$.form.passwdCheck.$errors" :key="index">
                {{ error.$message }}
              </p>
            </div>
          </div>

          <div class="mb-3">
            <p>Already have an account?
              <router-link to="/login">Login Here</router-link>
            </p>
          </div>
          <button class="w-100 btn btn-lg btn-primary" type="submit">Sign Up</button>
        </form>
      </main>
    </div>
  </div>
  <Alert :alert-data="state.alertData" v-if="state.alertData.open"></Alert>
</template>

<script>
import useVuelidate from '@vuelidate/core'
import {reactive} from "vue";
import Alert from "@/components/Alert";
import {email, required, minLength, helpers} from "@vuelidate/validators"
import {useStore} from "vuex";

export default {
  name: "Register",
  components: {Alert},
  setup() {
    const store = useStore()
    const state = reactive({
      form: {
        name: '',
        email: '',
        passwd: '',
        passwdCheck: ''
      },
      alertData: {
        open: false,
        message: '',
        status: 0
      }
    })
    const ensurePasswordIsSame = (passwd) => passwd === state.form.passwd
    const validations = {
      form: {
        name: {required},
        email: {required, email},
        passwd: {required, minLength: minLength(8)},
        passwdCheck: {
          required,
          minLength: minLength(8),
          ensurePasswordIsSame: helpers.withMessage('Password does not match', ensurePasswordIsSame)
        },
      }
    }
    const v$ = useVuelidate(validations, state)

    function submit() {
      if (!this.v$.$error) {
        store.dispatch('userStore/registerUser', state.form).then(data => {
          state.alertData.open = true
          state.alertData.message = data.data.message
          state.alertData.status = data.status
        })
      }

    }

    return {
      state,
      v$,
      validations,
      submit
    }
  },
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
  margin-bottom: -1px;
  border-radius: 0;
}

.last {
  margin-bottom: 10px !important;
  border-bottom-right-radius: 0.25rem !important;
  border-bottom-left-radius: 0.25rem !important;
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
