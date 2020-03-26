# Project: Post Dashboard
> Dashboard project for Kartaca.  

## Core technologies
> Html, Css, Js, Nodejs, Vuejs, Golang.

## Plugins
> vuex, axios, vuetify, jwt, gorilla/mux, gorilla/handlers, json, regexp.

## Database
> SQLITE3  

&nbsp;
<hr>
&nbsp;

## Instructions
### For testing on Docker: 
I'm using Windows Home Os. My Docker won't connect to localhost from EXPOSE port. So I'm using docker-machine IP address for root access. My docker-machine IP address is `192.168.99.100`. You can learn yours as write `docker-machine ip` to terminal. If your docker-machine IP address diffrent than mine you should change axios baseURL from `app/private/keys.js`. You have to change `baseURL` value with your docker-machine IP address.

Everything is fine? Then run the app with these command on your terminal inside of project folder:
```
    docker-compose up --build -d
```
After build application will be listen on `http://192.168.99.100:8080` or `http://[your-docker-machine-ip]:8080`. You can connect to app from your browser. Thats it.