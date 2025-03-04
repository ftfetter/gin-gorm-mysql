# gin-gorm-mysql

Simple CRUD just to practicing Gin + GORM + MySQL

To test the application, please go to project root folder and run

```
$ docker compose up -d
```

Once running, you can test using REST

- **GET** `/api/v1/users`: List all users
- **POST** `/api/v1/users`: Create a user
- **GET** `/api/v1/users/{id}`: Find a user by ID
- **PUT** `/api/v1/users/{id}`: Update a user
- **DELETE** `/api/v1/users/{id}`: Delete a user

Or even access `localhost:8080/swagger/index.html` for an interactive documentation

User Model:

```
{
    "name": "test",
    "email": "test@test.com"
}
```
