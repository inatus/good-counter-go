good-counter-go
===============

Simple Ajax 'Good!' counter web app implemented with Golang and MongoDB.

---

![Good! Counter Application](http://blog.inagaki.in/wp-content/uploads/2014/02/good-counter.png)

[http://good-counter-go.herokuapp.com/](http://good-counter-go.herokuapp.com/)  

# How to deploy Go web applications with MongoDB access to Heroku

## Prerequisite

- Register an account to Heroku
- Install Heroku Toolbelt
- Set $GOPATH environment variable

## Create a Go source folder

Create a Go source folder to $GOPATH/src.

```bash
$ mkdir $GOPATH/src/demoapp
$ cd $GOPATH/src/demoapp
$ git init
```

## Create a Heroku project

Login to Heroku Toolbelt and create a new project.

```bash
$ heroku login
Enter your Heroku credentials.
Email: you@example.com
Password:
Uploading ssh public key /Users/you/.ssh/id_rsa.pub
$ heroku create -b https://github.com/kr/heroku-buildpack-go.git
Creating vast-brook-7638... done, stack is cedar
BUILDPACK_URL=https://github.com/kr/heroku-buildpack-go.git
http://vast-brook-7638.herokuapp.com/ | git@heroku.com:vast-brook-7638.git
```

## Change the project name

As the project is create with a random name, change the name to any. 

```bash
$ heroku apps:rename good-counter-go
Renaming vast-brook-7638 to good-counter-go... done
http://good-count-go.herokuapp.com/ | git@heroku.com:good-count-go.git
Git remote heroku updated
```

## Add MongoDB addon to the Heroku project

Add MongoDB addon (free) to the project. Your credit card information must be registered even though you only use the free version.

```bash
$ heroku addons:add mongohq
Adding mongohq on good-counter-go... done, v4 (free)
Use `heroku addons:docs mongohq` to view documentation.
```

## Create a Go web application

Create a Go web application connected to MongoDB. Consider the following points.

### Designate environment variable PORT as the port of the http server.

```go
http.ListenAndServe(":"+os.Getenv("PORT"), nil)
```

### Designate environment variable MONGOHQ_URL as the URL of the instance of MongoDB

Environment variable MONGOHQ_URL also includes the user name and the password.

```go
sess, err := mgo.Dial(os.Getenv("MONGOHQ_URL"))
```

### Designate the database name from MongoHQ Web Manager

Enter MongoDB Manager from Heroku Apps Web Manager and confirm the database name. Set the name as the database that Go connects to.

![Heroku Apps Manager](http://blog.inagaki.in/wp-content/uploads/2014/02/heroku-app-menu.png)

![MongoHQ Manager](http://blog.inagaki.in/wp-content/uploads/2014/02/heroku-mongohq-menu.png)

```go
const MONGO_DB_NAME = "app21817638"
c := mgoSession.DB(MONGO_DB_NAME).C("count")
```

## Insert collections and documents to the DB

### In case to add collections and documents from command line

Add a new user from User tab of Admin page and login to the console with the following command from command line.

![MongoHQ User Addition](http://blog.inagaki.in/wp-content/uploads/2014/02/heroku-mongohq-add-user.png)

```bash
$ mongo troup.mongohq.com:10084/app21817638 -u <user> -p <password>
```

### In case to add collections and documents from the web manager

Create collections from "Create a collection" at Colletions page

![MongoHQ Collection Addition](http://blog.inagaki.in/wp-content/uploads/2014/02/heroku-mongohq-add-col.png)

## Register the application to the Heroku repository

Commit the source code to the repository.

```bash
$ git add -A .
$ git commit -m 'Add source'
```

Create a setting file for Heroku, download the libraries, and commit them to the repository.

```bash
$ echo 'web: good-counter-go' > Procfile
$ go get github.com/kr/godep
$ godep save 
$ git add -A .
$ git commit -m 'Add dependencies'
```

Lastly, push everything to Heroku.

```bash
$ git push heroku master
-----> Fetching custom git buildpack... done
-----> Go app detected
-----> Using go1.1.2
-----> Running: godep go install -tags heroku ./...
-----> Discovering process types
       Procfile declares types -> web

-----> Compressing... done, 3.0MB
-----> Launching... done, v5
       http://good-counter-go.herokuapp.com deployed to Heroku
```

## Confirmation of behavior

```bash
$ heroku open
```

## In case to push the resources to Github as well

Create a new repository in Github (do not enable auto init) and execute the following commands.

```bash
$ git remote add origin https://github.com/inatus/good-counter-go.git
$ git push -u origin master
```
