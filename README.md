# Description

This project is a *Go language* learning project with simple RestFul services. It uses postgres db inside with DockerFile. You can get image with dockerfile or create your own postgres database without it. 

For DockerFile help, you can review the following link.

[Dockerize Postgres](https://docs.docker.com/engine/examples/postgresql_service/)

# Database table configuration
```
      CREATE TABLE USERS (
        ID INT PRIMARY KEY,
        NAME TEXT NOT NULL,
        SURNAME TEXT NOT NULL,
        AGE INT NOT NULL
      );
      
      CREATE SEQUENCE public.users_id_seq NO MINVALUE NO MAXVALUE NO CYCLE;
      ALTER TABLE public.users ALTER COLUMN id SET DEFAULT nextval('public.users_id_seq');
      ALTER SEQUENCE public.users_id_seq OWNED BY public.users.id;
```

# Database access configuration inside code 
Under config/config.go directory in the project, you will find database access configuration. You can change it with your custom configuration.
```
      DB_USER     = "docker"
      DB_PASSWORD = "docker"
      DB_NAME     = "docker"
      PORT = "32770"
```
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
You need to following library for the postgres.
```
      go get github.com/lib/pq
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
