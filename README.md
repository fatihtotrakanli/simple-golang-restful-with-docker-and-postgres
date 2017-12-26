# Description

This project is a *Go language* learning project with simple RestFul services. It has global array inside, not real DB connection so when server down added JSON's will be deleted.

# How can run?

First of all, you need to have *Go* in your computer. For Mac you can install with brew easily.

```
      brew install go
```

If everything is OK, you should encounter an output like this at terminal when wrote *go version*.

```
      go version                                    
      go version go1.9.2 darwin/amd64
```

For run the project, in the project directory you need to write following command.

```
      go run main.go
```

If everything works correctly, you can start the CRUD operations with following URL.

```
      http://127.0.0.1:3000
```

# URL's and Example

List all of user (Need To Use GET method)
```
      http://127.0.0.1:3000/getAll
```
Add new User with JSON type ((Need To Use POST method))
```
      http://127.0.0.1:3000/newUser
      
      {
      	"id": 1,
      	"name": "mockName",
      	"surname": "mockSurname",
      	"age": 30
      	}
```
List one user with the given Id (Need To Use GET method)
```
      http://127.0.0.1:3000/users/1
```
Update one user with the given Id (Need To Use PUT method)
```
      http://127.0.0.1:3000/users/1
```
Delete one user with the given Id (Need To Use DELETE method)
```
      http://127.0.0.1:3000/users/1
```
