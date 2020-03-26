<template>
  <v-container class="fill-height" fluid>
    <v-row align="center" justify="center">
      <v-col cols="12" sm="12" md="8">
        <v-card class="elevation-12">
          <v-toolbar color="warning" dark flat>
            <v-toolbar-title>New Post</v-toolbar-title>
          </v-toolbar>
          <v-card-text>
            <v-form v-model="valid" ref="form">
              <v-textarea label="Text" :rules="rules" v-model="message" :counter="144"></v-textarea>
            </v-form>
          </v-card-text>
          <v-card-actions>
            <v-spacer />
            <v-btn color="warning" :disabled="!valid" @click="send">Send</v-btn>
          </v-card-actions>
        </v-card>
      </v-col>
    </v-row>
  </v-container>
</template>

<script>
export default {
  data: () => {
    return {
      valid: true,
      rules: [
        v => v.length <= 144 || "Max 144 characters",
        v => !!v || "This field can't be empty"
      ],
      message: ""
    };
  },
  methods: {
    send() {
      if (this.message.length <= 144 && this.message.length > 0) {
        const date = new Date().getTime().toString();
        const user = this.$store.getters.getUser
        const post = {
          body: this.message,
          author: user.username,
          date: date,
          avatar: user.avatar
        };
        this.$store.dispatch("sendPost", post).then(() => {
          this.$router.push("/")
          
        });
      }
    }
  }
};
</script>

<style>
</style>