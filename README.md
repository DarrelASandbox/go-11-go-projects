<details>
  <summary>Table of Contents</summary>
  <ol>
    <li><a href="#about-the-project">About The Project</a></li>
    <li><a href="#simple-http-server">simple-http-server</a></li>
    <li><a href="#movies-crud">movies-crud</a></li>
  </ol>
</details>

&nbsp;

## About The Project

- [Learn Go Programming by Building 11 Projects – Full Course](https://www.youtube.com/watch?v=jFfo23yIWac)
- [freeCodeCamp](https://www.freecodecamp.org/)
- [Akhil Sharma](https://github.com/AkhilSharma90)

1. [Original Repo: simple-http-server](https://github.com/AkhilSharma90/simple-http-server-GO)
2. [Original Repo: movies-crud](https://github.com/AkhilSharma90?tab=repositories&type=source)
3. [Original Repo: ]()
4. [Original Repo: ]()
5. [Original Repo: ]()
6. [Original Repo: ]()
7. [Original Repo: ]()
8. [Original Repo: ]()
9. [Original Repo: ]()
10. [Original Repo: ]()
11. [Original Repo: ]()

&nbsp;

---

&nbsp;

## simple-http-server

```
        -> /      -> index.html
Server  -> /hello -> hello func
        -> /form  -> form func  -> form.html
```

&nbsp;

---

&nbsp;

## movies-crud

- MySQL + CRUD
- No database

```
                                          ROUTES        FUNCTIONS     ENDPOINTS       METHODS

                                        -> GET ALL      getMovies     -> /movies      -> GET      <--|
                                        -> GET BY ID    getMovie      -> /movies/id   -> GET      <--|
Movies Server serving using Gorilla MUX -> CREATE       creteMovie    -> /movies      -> POST     <--- POSTMAN
                                        -> UPDATE       updateMovie   -> /movies/id   -> PUT      <--|
                                        -> DELETE       deleteMovie   -> /movies/id   -> DELETE   <--|
```

- [Error message "go: go.mod file not found in current directory or any parent directory; see 'go help modules'"](https://stackoverflow.com/questions/66894200/error-message-go-go-mod-file-not-found-in-current-directory-or-any-parent-dire)

```sh
go env
go env -w GO111MODULE=auto
go mod init 02-movies-crud/main.go
go mod tidy
go get -u github.com/gorilla/mux
```

&nbsp;

---

&nbsp;
