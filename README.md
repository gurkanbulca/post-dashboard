# Project: Post Dashboard
> Dashboard project for Kartaca.  

## Core technologies
> Html, Css, Js, Nodejs, Vuejs, Golang.

## Plugins
> vuex, axios, vuetify, jwt, gorilla/mux, gorilla/handlers, json, regexp.

## Database
> SQLITE3  

## API
> Avataaars: Every new user will get random avatar on register process. How funny is it? 

&nbsp;
<hr>
&nbsp;

## Instructions
### For testing on Docker: 
> Everything is ready for docker-compose test. Only thing you have to do is start a terminal on project main folder. Then execute the command:
```
    docker-compose up -d
```
### Created Data
> Some data(users,posts,comments...) already created for easier test surface. Every user password is '12345678'. Also you can register by yourself. Check the sqlite3 database for further information. All user data stored as unencrypted.


## KEYS
### Axios BaseURL
> You can change axios baseURL if needed under `@/app/private/keys.js` file. 

### JWTsecretKey
> For change JWTsecretKey on Vue head to `@/app/private/keys.js`. Then you will see JWTsecretKey variable inside of the file.
> For change JWTsecretKey on golang-rest-api head to `@/backend/go-rest-api/main.go`. Under 'GenerateJWT' function, you have to change string value of mySigningKey variable.
