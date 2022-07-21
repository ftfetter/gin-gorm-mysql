# gin-gorm-mysql
Simple CRUD just to practicing Gin + GORM + MySQL

To test the application, please go to project root folder and run  
```
$ docker compose up
```
to setup a MySQL server.

Then,  
```
$ go run main.go
```  
to run the application itself.

Once running, you can test using REST  
- **GET** `/users`: List all users
- **POST** `/users`: Create a user
- **GET** `/users/{id}`: Find a user by ID
- **PUT** `/users/{id}`: Update a user
- **DELETE** `/users/{id}`: Delete a user

User Model:
```
{
    "name": "test",
    "email": "test@test.com"
}
