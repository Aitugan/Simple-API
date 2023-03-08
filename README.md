# Simple Golang & Gorilla Mux REST API

## Completed requirements

### API requirements:

- Get a list of all books
- Get a single book by ID
- Add a new book
- Update an existing book
- Delete a book

### General requirements:

- Use the standard Go library for HTTP server and JSON encoding/decoding.
- Use the mux package for routing.
- Write clear and concise code with meaningful variable and function names.
- Use git for version control and commit your work regularly.
- Implement authentication and authorization for the API using JSON Web Tokens (JWT).
- Containerize your application using Docker.
- Use the testify package for unit testing.

## Launching the API:

Simply run the following commands:

```
$ docker build -t techt-task .
$ docker run techt-task
```

## Sample API requests

## Users:

### User Login:

```
POST http://localhost:8080/user/login
{
    "email":"James@mail.com",
    "password": "password1"
}
```

### User Registration:

```
POST http://localhost:8080/user/register
{
    "name":"James",
    "email":"James@mail.com",
    "password": "password1"
}
```

## Books:

### All books:

```
GET http://localhost:8080/book
```

### Book by id:

```
GET http://localhost:8080/book/3c28de98-e09a-4f3e-884b-624694b13c96
```

### Book Update:

```
PUT http://localhost:8080/book/3c28de98-e09a-4f3e-884b-624694b13c96
{
    "publicationDate": "27.09.2017"
}
```

### Delete book by id:

```
DELETE http://localhost:8080/book/3c28de98-e09a-4f3e-884b-624694b13c96
```

### Create book:

```
POST http://localhost:8080/book
{
    "title":"Alchemist",
    "author":"Paulo Coelho",
    "publicationDate": "01.01.2001"
}
```

### My books:

```
GET http://localhost:8080/book/my
```
