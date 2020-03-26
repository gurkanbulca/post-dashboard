<template>
  <v-card
    class="mx-auto mt-5 pt-5"
    max-width="800"
    @mouseover="raised=true"
    @mouseout="raised=false"
    :raised="raised"
    @click="on"
  >
    <v-row>
      <v-col md="10"><v-card-text class="headline text--secondary" >"{{post.body}}"</v-card-text></v-col>
      <v-col md="2"><v-icon>chat_bubble_outline</v-icon><small class="text--secondary mr-2">{{post.commentCount}}</small><v-icon :color="post.isFavorite?'#FF4F5A':'secondary'">{{post.isFavorite?"favorite":"favorite_outline"}}</v-icon><small class="text--secondary">{{post.favoriteCount}}</small></v-col>
      
    </v-row>
    

    <v-card-actions>
      <v-list-item class="grow">
        <v-list-item-avatar color="grey darken-3">
          <v-img class="elevation-6" :src="post.avatar"></v-img>
        </v-list-item-avatar>

        <v-list-item-content>
          <v-list-item-title class="text--secondary">{{post.author}}</v-list-item-title>
        </v-list-item-content>

        <v-row align="center" justify="end">
          <span class="subheading mr-2 text--disabled">{{post.date}}</span>
        </v-row>
      </v-list-item>
    </v-card-actions>
    <hr class="mt-3 mb-3 ml-3 mr-3" opacity="0.1">
    
  </v-card>
</template>

<script>
export default {
  props: ["post"],
  data: () => {
    return {
      raised: false
    };
  },
  methods: {
    on() {
      this.$emit("goster", this.post);
    }
  }
};
</script>

<style scoped>
.v-card:hover {
  cursor: pointer;
}
</style>