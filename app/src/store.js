import Vue from "vue"
import Vuex from "vuex"
import axios from "axios"
import jwt from "jsonwebtoken"
import { JWTsecretKey,baseURL } from "../private/keys"


Vue.use(Vuex)


axios.defaults.baseURL = baseURL

export const store = new Vuex.Store({
    state: {
        posts: [],
        comments: [],
        user: {},
    },
    getters: {
        getPosts(state) {
            return state.posts;
        },
        getComments(state) {
            return state.comments
        },
        getUser(state) {
            return state.user
        },
        getToken(state) {
            return state.user.token
        },
        isAuthenticated(state) {
            return (state.user.token != null && state.user.expire > new Date().getTime())
        }
    },
    mutations: {
        setPosts(state, posts) {
            state.posts = posts;
        },
        setComments(state, comments) {
            state.comments = comments
        },
        formatDate(vuexContext, post) {
            const date = new Date(+post.date)
            let hour = date.getHours()
            let mins = date.getMinutes()
            const year = date.getFullYear()
            let month = date.getMonth() + 1
            let day = date.getDate()
            if (day < 10) {
                day = "0" + day
            }
            if (month < 10) {
                month = "0" + month
            }
            if (mins < 10) {
                mins = "0" + mins
            }
            if (hour < 10) {
                hour = "0" + hour
            }

            post.date = hour + ":" + mins + " " + day + "." + month + "." + year
        },

        setToken(state, token) {
            state.user.token = token
        },
        setUser(state, user) {
            if (user.favorites == null) user.favorites = []
            state.user = user
        },
        getTokenFromCookies(state) {
            state.user.token = Vue.$cookies.get("token")
            if (state.user.token) {
                try {
                    jwt.verify(state.user.token, JWTsecretKey, function (err, decoded) {
                        // console.log(decoded);
                        state.user = { ...state.user, username: decoded.username, avatar: decoded.avatar, id: decoded.id, favorites: decoded.favorites, expire: decoded.exp }
                        if (err) {
                            throw new Error()
                        }
                    })
                } catch (err) {
                    console.log(err);
                }
            }


        }

    },
    actions: {
        setPosts({ commit, getters }) {
            return axios.get("/posts", {

            })
                .then(response => {
                    response.data.map(post => {
                        commit("formatDate", post)
                        const user = getters.getUser
                        if (user && user.favorites) {
                            if (user.favorites.indexOf(post.id) > -1) {
                                post.isFavorite = true
                            } else {
                                post.isFavorite = false
                            }
                        }

                    })
                    let posts = response.data.reverse()
                    commit("setPosts", posts)
                })
        },


        sendPost(vuexContext, post) {
            return axios.post("/posts", { ...post, id: 1 })
                .then(() => {
                    vuexContext.commit("formatDate", post)
                    let posts = vuexContext.getters.getPosts
                    post.commentCount = 0
                    post.favoriteCount = 0
                    posts = [post, ...posts]
                    vuexContext.commit("setPosts", posts)
                })

        },
        getComments(vuexContext, postid) {
            return axios.get("/comments?id=" + postid)
                .then(comments => {
                    if (comments.data) {
                        comments.data.map(comment => {
                            vuexContext.commit("formatDate", comment)
                        })
                        vuexContext.commit("setComments", comments.data)
                    }
                    else {
                        vuexContext.commit("setComments", [])
                    }

                })
        },
        sendComment(vuexContext, comment) {
            return axios.post("/comments", comment)
                .then(() => {
                    vuexContext.commit("formatDate", comment)
                    let comments = vuexContext.getters.getComments
                    comments.push(comment)
                    vuexContext.commit("setComments", comments)
                })
        },

        changeFavorite(vuexContext, favorite) {
            let user = vuexContext.getters.getUser
            favorite = { ...favorite, userid: user.id }
            // console.log(favorite);
            return axios.post("/favorite", favorite)
                .then(() => {
                    // console.log(user.favorites);
                    if (favorite.isFavorite) {
                        user.favorites.push(favorite.postid)
                    } else {
                        user.favorites.splice(user.favorites.indexOf(favorite.postid), 1)
                    }
                    // console.log(new Date().getTime() + 1800 * 1000);
                    let token = jwt.sign({ authorized: true, username: user.username, avatar: user.avatar, id: user.id, favorites: user.favorites, exp: new Date().getTime() + 1800 * 1000 }, JWTsecretKey)
                    // console.log(token);
                    vuexContext.commit("setToken", token)
                    Vue.$cookies.set("token", token)
                })
        },

        login(vuexContext, user) {
            if (user.username &&
                user.username.length >= 6 &&
                user.username.length <= 16 &&
                user.password &&
                user.password.length >= 8 &&
                user.password.length <= 16) {
                return axios.post("/login", user)
                    .then(res => {

                        if (res.data == "wrong") throw new Error()
                        // console.log(res.data);
                        jwt.verify(res.data.token, JWTsecretKey, function (err, decoded) {
                            const user = { username: decoded.username, avatar: decoded.avatar, expire: decoded.exp, id: decoded.id, favorites: decoded.favorites, token: res.data.token }
                            vuexContext.commit("setUser", user)
                        })
                        Vue.$cookies.set("token", res.data.token)
                        let posts = vuexContext.getters.getPosts
                        posts.map(post => {
                            if (res.data.favorites.indexOf(post.id) > -1) {
                                post.isFavorite = true
                            } else {
                                post.isFavorite = false
                            }
                        })
                    })

            }
        },
        register(vuexContext, user) {
            return axios.post("/register", { ...user })
                .then(res => {
                    if (res.data == "wrong") throw new Error()
                    jwt.verify(res.data.token, JWTsecretKey, function (err, decoded) {
                        const user = { username: decoded.username, avatar: decoded.avatar, expire: decoded.exp, id: decoded.id, favorites: decoded.favorites, token: res.data.token }
                        vuexContext.commit("setUser", user)
                    })
                    Vue.$cookies.set("token", res.data.token)
                    let posts = vuexContext.getters.getPosts
                    posts.map(post => {
                        if (res.data.favorites.indexOf(post.id) > -1) {
                            post.isFavorite = true
                        } else {
                            post.isFavorite = false
                        }
                    })
                })
        }

    }

})