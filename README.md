![Bloguero Logo](./Logo.png)
<h1 style="text-align:center;">Bloguero</h1>
This repo holds the source code for *Bloguero.net*, an open source site that allows users to posts short posts called blogs.


<h2 style="text-align:center;">Nerd Stuff</h2>
Backend was made with Go, with the majority of code made with the standard library; [GORM](https://github.com/go-gorm/gorm) & the [Go JWT package](https://github.com/golang-jwt/jwt). The frontend was made with Sveltekit

If you wish to download this code to test it or build upon it note that it uses a config.json file to store all enviroment variable(Which is stored in the top-level directory); You must edit the struct in the configJson.go in the config directory(i.e. any new env variables not already in the struct must be added).

# Current TODOs
- [ ] Finish single post page design(Under [id] folder)
    - [ ] Place a button to view post on each blog card
- [ ] Finish account page & route
- [ ] [Add api logic to all buttons & pages (Look at server handlers to see the routes)](https://learn.svelte.dev/tutorial/other-handlers)
- [ ] Implement infinite scroll on all blogs page & single blog page
    - [True Infinite scroll](https://svelte.dev/repl/aacd1a2d8eb14bb19e5cb3b0ad20fdbe?version=3.32.3)
    - [Version which uses button click](https://svelte.dev/repl/5823b6e8a30447c59ce5b770f8a84593?version=3.31.2)
- [ ] Refactor & Clean Up code/directory structure
- [ ] ...