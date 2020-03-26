<template>
  <nav>
    <v-toolbar>
      <v-toolbar-title class="text-uppercase">
        <span class="font-weight-light">dash</span>
        <span>board</span>
      </v-toolbar-title>
      <router-link to="/" class="ml-3" text tag="a" style="text-decoration:none;">
        <v-btn text>Homepage</v-btn>
      </router-link>
      <v-spacer></v-spacer>
      <v-btn text v-if="isLoggedin" @click="logout">
        <span>Sign Out</span>
        <v-icon right>exit_to_app</v-icon>
      </v-btn>
      <router-link to="/login" tag="a" style="text-decoration:none;" v-else>
        <v-btn text>
          <span>Login</span>
          <v-icon right>contacts</v-icon>
        </v-btn>
      </router-link>
    </v-toolbar>
  </nav>
</template>

<script>
export default {
  computed: {
    isLoggedin() {
      // console.log(this.$store.getters.isAuthenticated);
      return this.$store.getters.isAuthenticated;
    }
  },
  methods: {
    logout() {
      this.$store.commit("setToken", null);
      this.$store.commit("setUser", {});
      this.$cookies.remove("token");
      this.$router.push("/login")
    }
  }
};
</script>

<style>
</style>