<template>
  <v-card
    class="mx-auto pt-5"
    max-width="800"
    @mouseover="raised=true"
    @mouseout="raised=false"
    :raised="raised"
  >
    <v-row>
      <v-col lg="10">
        <v-card-text class="headline text--secondary">"{{message.body}}"</v-card-text>
      </v-col>
      <v-col>
        <v-icon raised>chat_bubble_outline</v-icon>
        <small class="text--secondary mr-2">{{message.commentCount}}</small>
        <a class="favorite" @click.prevent="toggleFavorite" :disabled="!isUser">
          <v-icon
            :color="message.isFavorite?'#FF4F5A':'secondary'"
          >{{message.isFavorite?"favorite":"favorite_outline"}}</v-icon>
        </a>

        <small class="text--secondary">{{message.favoriteCount}}</small>
      </v-col>
    </v-row>

    <v-card-actions>
      <v-list-item class="grow">
        <v-list-item-avatar color="grey darken-3">
          <v-img class="elevation-6" :src="message.avatar"></v-img>
        </v-list-item-avatar>

        <v-list-item-content>
          <v-list-item-title class="text--secondary">{{message.author}}</v-list-item-title>
        </v-list-item-content>

        <v-row align="center" justify="end">
          <span class="subheading mr-2 text--disabled">{{message.date}}</span>
        </v-row>
      </v-list-item>
    </v-card-actions>
    <v-divider class="ml-4 mt-3 mr-4 mb-3" />
    <v-timeline dense class="mr-5">
      <v-timeline-item class="white--text" large v-if="isUser">
        <template v-slot:icon>
          <v-avatar>
            <img :src="useravatar" />
          </v-avatar>
        </template>
        <v-text-field
          v-model="input"
          hide-details
          flat
          label="Leave a comment..."
          solo
          @keydown.enter="send"
        >
          <template v-slot:append>
            <v-btn class="mx-0" depressed @click="send">Post</v-btn>
          </template>
        </v-text-field>
        <v-divider class="mt-5" />
      </v-timeline-item>

      <v-timeline-item v-for="comment in comments" :key="comment.id" large>
        <template v-slot:icon>
          <v-avatar>
            <img :src="comment.avatar" />
          </v-avatar>
        </template>
        <!-- <template v-slot:opposite>
          <span>Tus eu perfecto</span>
        </template>-->
        <v-card class="elevation-2">
          <v-card-title class="headline text--secondary">"{{comment.comment}}"</v-card-title>
          <v-list-item class="grow">
            <v-list-item-content>
              <v-list-item-title class="text--secondary">{{comment.username}}</v-list-item-title>
            </v-list-item-content>

            <v-row align="center" justify="end">
              <span class="subheading mr-2 text--disabled">{{comment.date}}</span>
            </v-row>
          </v-list-item>
        </v-card>
      </v-timeline-item>
    </v-timeline>
  </v-card>
</template>

<script>
export default {
  created() {
    this.$store.dispatch("getComments", this.message.id);
  },
  props: ["message"],
  data: () => {
    return {
      input: "",
      raised: false
    };
  },
  methods: {
    send() {
      const user = this.$store.getters.getUser;
      const comment = {
        id: this.comments.length>0 ? this.comments[0].id + 1 : 1,
        comment: this.input,
        username: user.username,
        avatar: user.avatar,
        date: new Date().getTime().toString(),
        postid: this.message.id,
        commentCount: +this.message.commentCount
      };
      this.$store.dispatch("sendComment", comment).then(() => {
        this.input = "";
        this.message.commentCount++;
      });
    },
    toggleFavorite() {
      if (this.$store.getters.getUser) {
        this.message.isFavorite = !this.message.isFavorite;
        if (this.message.isFavorite) {
          this.message.favoriteCount++;
        } else {
          this.message.favoriteCount--;
        }
        this.$store.dispatch("changeFavorite", {
          isFavorite: this.message.isFavorite,
          count: this.message.favoriteCount,
          postid: +this.message.id
        });
      }

      // this.$store.dispatch("updateFavorite")
    }
  },
  computed: {
    useravatar() {
      const user = this.$store.getters.getUser;
      return user.avatar;
    },
    isUser() {
      return this.$store.getters.isAuthenticated
    },
    comments() {
      let comments = this.$store.getters.getComments;
      return comments.reverse();
    }
  },
  watch: {
    message() {
      this.$store.dispatch("getComments", this.message.id);
    }
  }
};
</script>

<style>
a.favorite {
  text-decoration: none;
}

a.favorite[disabled] {
  pointer-events: none;
}
</style>