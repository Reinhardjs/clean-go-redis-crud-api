![image](https://user-images.githubusercontent.com/7758970/205461470-544322a0-3577-4bf8-980b-39b45762fd5e.png)


## Architecture

In this project, I implemented clean architecture, that the main structure consist of model, repository, usecase, and controller

I'm using this architecture because it provide us to be more convenience to implement Open/Close Principle, Single Responsibility Principle, and also Unit Test.

<br> 

#### Built With

* Go (Mux)
* Gorm
* Redis
* PostgreDB
* Docker
* Kubernetes

<br>

## Installation
`go get -d -v ./...`

`go mod download`

`go run main.go`

<br>

## Endpoints

API Endpoint Host : http://103.134.154.18:30033/

I've deployed this project to my personal VPS, and deployed to `single-node kubernetes`

This project is configured to be able to containerized using docker (Dockerfile). <br> And deployed on a single-node kubernetes cluster. (Deployment.yaml)

Here is what i mean by single-node cluster, go checkout my story on medium here: <br>
https://reinhardjsilalahi.medium.com/beginners-guide-simple-hello-kubernetes-all-in-one-on-a-single-vps-fcfdfee9edfc

<br> 

### Create Post
`POST` http://103.134.154.18:30033/posts

Example Request Payload:
```
{
    "title": "This is title",
    "description": "This is description"
}
```

<br> 

Example Response Payload:
```
{
    "status": 201,
    "message": "success",
    "data": {
        "id": 62,
        "title": "This is title",
        "description": "This is description",
        "created_at": "2022-12-04T03:23:55.0987237+07:00",
        "updated_at": "2022-12-04T03:23:55.0987237+07:00",
        "comments": null
    }
}
```

<br>

### Get post list
`GET` http://103.134.154.18:30033/posts

<br>

### Get single post
`GET` http://103.134.154.18:30033/posts/{post-id}

<br> 

### Update post
`PUT` http://103.134.154.18:30033/posts/{post-id}

`PATCH` http://103.134.154.18:30033/posts/{post-id}

Example Request Payload:
```
{
    "title": "This is updated title",
    "description": "This is updated description"
}
```

Example Response Payload:
```
{
    "status": 200,
    "message": "success",
    "data": {
        "id": 61,
        "title": "This is updated title",
        "description": "This is updated description",
        "created_at": "2022-12-03T20:18:26.090668Z",
        "updated_at": "2022-12-03T20:40:17.473926Z"
    }
}
```

<br> 

### Delete post
`DELETE` http://103.134.154.18:30033/posts/{post-id}

<br>

Demo and explanation on youtube (bahasa) :

https://youtu.be/p-QQj4LtuD8
