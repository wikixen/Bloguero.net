# Blog REST API
The purpose of this repo is to show my abilty to make a REST API using as much of the Go standard library as I possibly could. The API allows for the creation of user accounts, as well as the creation of blogs by each user. Since this is a proof of concept it uses an sqlite database as it was the easiest database to set up. There were some external libraries that I had to use; Specifically [JWT](https://github.com/golang-jwt/jwt), to implement JWT token authentication, & [GORM](https://github.com/go-gorm/gorm) to interact with the database. Argon2id was used to hash passwords as it is the best hashing algorithm as the time of this repo's development(2024)

This not only shows my ability to make a REST API but my ability to do so with the stdlib of Go.


If you wish to download this code to test it or build upon it note that it uses a config.json file to store all enviroment variable(Which is stored in the top-level directory); You must edit the struct in the configJson.go in the config directory(i.e. any new env variables not already in the struct must be added).
# TODO
- [ ] Finish single post page design(Under [id] folder)
    - [ ] Place a button to view post on each blog card
- [ ] Finish account page & route
- [ ] [Add api logic to all buttons & pages (Look at server handlers to see the routes)](https://learn.svelte.dev/tutorial/other-handlers)
- [ ] Implement infinite scroll on all blogs page & single blog page
- [ ] Refactor & Clean Up code/directory structure
- [ ] ...