# What's this?

This is an example of a simple Golang HTTP API that I used to get to know more about how Golang works. The structure of this project is deeply inspired by NestJS (a framework for NodeJS) where each layer has its own responsibility, for example, the repository layer is responsible for speaking with the database, the service layer is responsible for business logic, and the controller layer forwards user requests to the related service layers.

**This is not a production-ready project that you have to imitate, which solely has a role for training and college assignments.**

# Table of Contents

- [What's this?](#what-s-this-)
- [Table of Contents](#table-of-contents)
- [Exposed API Endpoints](#exposed-api-endpoints)
  - [Authentication API](#authentication-api)
    - [Login](#login)
    - [Register](#register)
  - [Question API](#question-api)
    - [Create](#create)
    - [Read All](#read-all)
    - [Read By ID](#read-by-id)
    - [Update](#update)
    - [Delete](#delete)
  - [Answer API](#answer-api)
    - [Create](#create-1)
    - [Read All](#read-all-1)
    - [Update](#update-1)
    - [Delete](#delete-1)

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
