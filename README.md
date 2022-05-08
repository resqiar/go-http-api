# What's this?

This is an example of a simple Golang HTTP API that I used to get to know more about how Golang works. The structure of this project is deeply inspired by NestJS (a framework for NodeJS) where each layer has its own responsibility, for example, the repository layer is responsible for speaking with the database, the service layer is responsible for business logic, and the controller layer forwards user requests to the related service layers.

# Table of Contents

- [What's this?](#what-s-this-)
- [Table of Contents](#table-of-contents)
- [Installation](#installation)
- [Exposed API Endpoints](#exposed-api-endpoints)
  - [Authentication API](#authentication-api)
    - [Login](#login)
    - [Register](#register)
  - [Users API](#users-api)
    - [Find User by Username](#find-user-by-username)
    - [Find User by ID](#find-user-by-id)
    - [Get All Users (TEST PURPOSE ONLY)](#get-all-users--test-purpose-only-)
  - [Question API](#question-api)
    - [Create](#create)
    - [Read All](#read-all)
    - [Read By ID](#read-by-id)
    - [Search](#search)
    - [Update](#update)
    - [Delete](#delete)
  - [Answer API](#answer-api)
    - [Create](#create-1)
    - [Read All](#read-all-1)
    - [Update](#update-1)
    - [Delete](#delete-1)

# Installation

**NOTE: YOU NEED TO ADJUST THE DATABASE CONNECTION FIRST, GO TO `main.go` AND FILL THE DSN STRING WITH YOURS.**

After you adjust the database, simply run this

```bash
    go run main.go
```

Server will run on port 8000

# Exposed API Endpoints

Below are all the exposed endpoint for the APIs.

## Authentication API

The authentication endpoint is used to authenticate the API access for cetain user.

### Login

This API will match input (username and password) with the existing user.

#### Path

```
    /v1/auth/login
```

#### Method

```
    POST
```

#### Expected Body

```json
{
  "email": string, // required
  "password": string // required
}
```

### Register

This API will register and create a new instance of user with the given email and password.

#### Path

```
    /v1/auth/register
```

#### Method

```
    POST
```

#### Expected Body

```json
{
  "username": string, // required
  "email": string, // required
  "password": string // required
}
```

#### Return value

```json
{
  "data": {
    "access_token": string,
    "email": string,
    "username": string,
  },
  "response_ms": number,
  "status": number,
  "timestamp": string
}
```

## Users API

### Find User by Username

This API query an instance of user with the given username.

#### Path

```
    /v1/user/:username
```

#### Method

```
    GET
```

#### Return value

```json
{
  "data": {
    "ID": number,
    "Username": string,
    "Bio": string,
    "Questions": []Question,
    "CreatedAt": string,
    "UpdatedAt": string,
  },
  "response_ms": number,
  "status": number,
  "timestamp": string
}
```

### Find User by ID

This API query an instance of user with the given user ID.

#### Path

```
    /v1/userid/:id
```

#### Method

```
    GET
```

#### Return value

```json
{
  "data": {
    "ID": number,
    "Username": string,
    "Bio": string,
    "Questions": []Question,
    "CreatedAt": string,
    "UpdatedAt": string,
  },
  "response_ms": number,
  "status": number,
  "timestamp": string
}
```

### Get All Users (TEST PURPOSE ONLY)

This API query all instances of users from the database.

**IMPORTANT NOTE: YOU MUST NOT EXPOSE THIS API TO PRODUCTION OR YOU WILL GET A TROUBLE.**

#### Path

```
    /v1/users
```

#### Method

```
    GET
```

#### Return value

```json
{
  "data": [
    {
      "ID": number,
      "Username": string,
      "Bio": string,
      "Questions": []Question,
      "CreatedAt": string,
      "UpdatedAt": string,
    },
    {...}
  ],
  "response_ms": number,
  "status": number,
  "timestamp": string
}
```

## Question API

The Question API is used as a front-end communication bridge to the database, please be aware of the expected path and request body input.

### Create

This API will create a new question based on the input.

> THIS API IS PROTECTED, YOU CAN ONLY ACCESS IT BY PROVIDING BEARER TOKEN

#### Path

```
    /v1/question/create
```

#### Method

```
    POST
```

#### Expected Body

```json
{
  "title": string, // required
  "desc": string // required
}
```

#### Return value

```json
{
  "response_ms": number,
  "status": number,
  "timestamp": string
}
```

### Read All

This API return all instance of questions inside of the database table.

#### Path

```
    /v1/questions
```

#### Method

```
    GET
```

#### Return value

```json
{
  "data": [
    {
      "ID": number,
      "Title": string,
      "Desc": string,
      "AuthorID": 1,
      "Answers": []Answer,
      "CreatedAt": string,
      "UpdatedAt": string,
      "Deleted": string || null,
    },
  ],
  "response_ms": number,
  "status": number,
  "timestamp": string
}
```

### Read By ID

This API return the specific question based on the provided ID.

#### Path

```
    /v1/question/:id
```

#### Method

```
    GET
```

#### Return value

```json
{
  "data": Question,
  "response_ms": number,
  "status": number,
  "timestamp": string
}
```

### Search

This API return all instance of questions that match with search query.

#### Path

```
    /v1/question/search?query=
```

#### Method

```
    GET
```

#### Return value

```json
{
  "data": Question,
  "response_ms": number,
  "status": number,
  "timestamp": string
}
```

### Update

This API will update existing question based on the input.

> THIS API IS PROTECTED, YOU CAN ONLY ACCESS IT BY PROVIDING BEARER TOKEN

#### Path

```
    /v1/question/update
```

#### Method

```
    POST
```

#### Expected Body

```json
{
  "id": number, // required
  "title": string,
  "desc": string
}
```

#### Return value

```json
{
  "response_ms": number,
  "status": number,
  "timestamp": string
}
```

### Delete

This API will SOFT DELETE existing question based on the input.

> THIS API IS PROTECTED, YOU CAN ONLY ACCESS IT BY PROVIDING BEARER TOKEN

#### Path

```
    /v1/question/delete
```

#### Method

```
    POST
```

#### Expected Body

```json
{
  "id": number, // required
}
```

#### Return value

```json
{
  "response_ms": number,
  "status": number,
  "timestamp": string
}
```

## Answer API

The Answer API is used as a front-end communication bridge to the answer table, please be aware of the expected path and request body input.

### Create

This API will create a new question based on the input.

> THIS API IS PROTECTED, YOU CAN ONLY ACCESS IT BY PROVIDING BEARER TOKEN

#### Path

```
    /v1/answer/create
```

#### Method

```
    POST
```

#### Expected Body

```json
{
  "answer": string, // required
  "question_id": number // required
}
```

#### Return value

```json
{
  "response_ms": number,
  "status": number,
  "timestamp": string
}
```

### Read All

This API return all instance of answers inside of the database table.

#### Path

```
    /v1/answers
```

#### Method

```
    GET
```

#### Return value

```json
{
  "data": [
    {
      "ID": number,
      "AnswerText": string,
      "AuthorID": number,
      "QuestionID": number,
      "CreatedAt": string,
      "UpdatedAt": string,
      "Deleted": string || null
    }
  ],
  "response_ms": number,
  "status": number,
  "timestamp": string
}
```

### Update

This API will update existing answer based on the input.

> THIS API IS PROTECTED, YOU CAN ONLY ACCESS IT BY PROVIDING BEARER TOKEN

#### Path

```
    /v1/answer/update
```

#### Method

```
    POST
```

#### Expected Body

```json
{
  "id": number, // required
  "answer": string,
}
```

#### Return value

```json
{
  "response_ms": number,
  "status": number,
  "timestamp": string
}
```

### Delete

This API will SOFT DELETE existing answer based on the input.

> THIS API IS PROTECTED, YOU CAN ONLY ACCESS IT BY PROVIDING BEARER TOKEN

#### Path

```
    /v1/answer/delete
```

#### Method

```
    POST
```

#### Expected Body

```json
{
  "id": number, // required
}
```

#### Return value

```json
{
  "response_ms": number,
  "status": number,
  "timestamp": string
}
```
