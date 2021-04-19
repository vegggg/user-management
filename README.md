# User Manual

To Run this project on localhost try: $ docker-compose up

## Directory explain

1. cmd/
   I use cobra as to implement a commandline application.
   Basicly, this project is started from cmd/server.go
   Stuff like init db connection, DI... are also there

2. main.go
   Only to init cmd/

3. init.sql
   Sql schema init script

4. Dockerfile, docker-compose
   For ease deployment

5. prototool.yaml
   A prototool plugins manager/linter

6. vendor/
   I use go vendor to manage project's dependencies

7. proto/
   Protobuf files to define service API, models, entities

8. handler/
   A handler that implement service API interface, All it does is controlling the business flow

9. otp/
   otp service: use a Redis instance to maintain user's OTP and user status of "is validated by OTP"

10. user/
    user service
    this package is defined a bit general
    Currently, I have implemented a sqlrepo for managing users. It also has a redis repo for future caching feature.
    In addition, in sqlrepo, we can easily add a psql for more flexibility.

11. entity/
    Contains entities

12. docs/
    A postman collection json to run demo

## API flow

API can be found in docs/

1. Create User

   1. POST localhost:8080/v1/users:sendOTP body: phone Response must includes a OTP since I can NOT send a SMS
   2. POST localhost:8080/v1/users:verifyOTP body: phone + otp (otp is only valid in 1 min)
   3. POST localhost:8080/v1/users body: {UserStruct} with validated otp (also expire after 1 min since otp validation)

2. Get User
   1. GET localhost:8080/v1/users/{phone}

## Run test

$ docker-compose up
$ go test ./...
