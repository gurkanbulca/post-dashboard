<template>
  <v-container class="fill-height" fluid>
    <v-row align="center" justify="center">
      <v-col cols="12" sm="8" md="4">
        <v-card class="elevation-12">
          <v-toolbar :color="isMember?'primary':'success'" dark flat>
            <v-toolbar-title>{{isMember?"Login form":"Register form"}}</v-toolbar-title>
          </v-toolbar>
          
          <v-card-text>
            <v-form v-model="valid" ref="form">
              <v-text-field
                v-model="username"
                :rules="nameRules"
                label="Username"
                prepend-icon="person"
                required
              ></v-text-field>
              <v-text-field
                v-model="email"
                :rules="emailRules"
                label="E-mail"
                prepend-icon="mail"
                required
                v-if="!isMember"
              ></v-text-field>

              <v-text-field
                v-model="password"
                id="password"
                label="Password"
                name="password"
                :rules="passwordRules"
                prepend-icon="lock"
                :type="show ? 'text' : 'password'"
                @click:append="show = !show"
                :append-icon="show ? 'visibility' : 'visibility_off'"
              />
              <v-text-field
                v-if="!isMember"
                prepend-icon="lock"
                v-model="rePassword"
                :append-icon="show1 ? 'visibility' : 'visibility_off'"
                :rules="[rePasswordRules,passwordConfirmationRule]"
                :type="show1 ? 'text' : 'password'"
                name="input-10-1"
                label="Re-enter Password"
                @click:append="show1 = !show1"
              ></v-text-field>
            </v-form>
          </v-card-text>
          <v-alert type="error" v-if="error">{{message}}</v-alert>
          <v-card-actions>
            <v-btn
              text
              @click="isMember=!isMember"
            >{{isMember?"I don't have a account.":"I already have a account."}}</v-btn>
            <v-spacer />
            <v-btn
              :color="isMember?'primary':'success'"
              :disabled="!valid"
              @click="authorize"
            >{{isMember?"Login":"Register"}}</v-btn>
          </v-card-actions>
        </v-card>
      </v-col>
    </v-row>
  </v-container>
</template>

<script>
import { generateRandomAvatar } from "../../api/avataaars-api/generator/generateAvatar";
export default {
  created() {
    if (this.$store.getters.isAuthenticated) {
      this.$router.push("/");
    }
  },
  data: () => {
    return {
      error: false,
      valid: true,
      isMember: true,
      show: false,
      show1: false,

      username: "",
      nameRules: [
        v => !!v || "Username is required",
        v => (v && v.length >= 6) || "Username must be more than 6 characters",
        v => (v && v.length <= 16) || "Username must be less than16 characters"
      ],

      email: "",
      emailRules: [
        v => !!v || "E-mail is required",
        v => /.+@.+\..+/.test(v) || "E-mail must be valid"
      ],

      password: "",
      passwordRules: [
        v => !!v || "Password is required",
        v => (v && v.length >= 8) || "Password must be more than 8 characters",
        v => (v && v.length <= 16) || "Password must be less than 16 characters"
      ],

      rePassword: "",
      rePasswordRules: [v => !!v || "Required."]
    };
  },
  computed: {
    passwordConfirmationRule() {
      return () => this.password === this.rePassword || "Password must match";
    }
  },
  methods: {
    authorize() {
      if (this.isMember) {
        this.$store
          .dispatch("login", {
            username: this.username,
            password: this.password
          })
          .then(() => {
            this.$router.push("/");
          })
          .catch(() => {
            this.message="Wrong username or password!"
            this.error = true;
            setTimeout(() => {
              this.error = false;
            }, 3000);
          });
      } else {
        if (this.password == this.rePassword) {
          this.$store
            .dispatch("register", {
              username: this.username,
              password: this.password,
              email: this.email,
              avatar: generateRandomAvatar("Circle")
            })
            .then(() => {
              this.$router.push("/");
            })
            .catch(() => {
              this.message="This user already exists!"
              this.error = true;
              setTimeout(() => {
                this.error = false;
              }, 3000);
            });
        }
        else{
          this.message="Password must match!"
              this.error = true;
              setTimeout(() => {
                this.error = false;
              }, 3000);
        }
      }
    }
  }
};
</script>

<style>
</style>