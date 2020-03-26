<template>
  <v-container>
    <v-card flat class="mx-auto mt-5 d-flex justify-center" max-width="800">
      <router-link tag="div" to="/new-post"><v-btn color="success" tile block >New Post</v-btn></router-link>
      
    </v-card>
    <Card v-for="post in posts" :key="post.id" @goster="goster" :post="post" />
    <v-dialog class="text-center dialog" v-model="display" max-width="800">
      <Details :message="message" class="dialog"/>
    </v-dialog>
  </v-container>
</template>

<script>
import Card from "./Card";
import Details from "./Details"
export default {
  components: {
    Card,
    Details,
  },
  data: () => {
    return {
      
      display: false,
      
      message: {
        id: "",
        author: "",
        body: "",
        data: "",
        isFavorite: false
      }
    };
  },
  methods: {
    goster(e) {
      this.message = e;
      this.display = true;
    },
    newPost() {
      if(this.$store.getters.isAuthenticated){
      this.$router.push("/new-post");
      }else{
        this.$router.push("/login")
      }
    },
   
    
  },
  computed: {
    posts() {
      return this.$store.getters.getPosts;
    },
    
  }
};
</script>

<style>
.dialog{
  overflow: hidden;
}
</style>