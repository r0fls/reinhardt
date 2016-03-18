# reinhardt

This is a golang MVC modeled loosely off of Django. Still in pre-alpha.

### Getting Started

#### Installation
```bash
go get github.com/r0fls/reinhardt
cd $GOPATH/src/github.com/r0fls/reinhardt
go build
```

#### Starting a new project

```bash
./reinhardt new <projectname>
cd <projectname>
go build
./<projectname> runserver
```
At which point you'll have a folder named `projectname` with the following structure:

    projectname
    ├─settings.json   
    ├─manager.go      
    ├─app
    │  └─urls.go
    │  └─views
    │      └─views.go
    │  └─models
    │      └─models.go
    │  └─temps
    │      └─home.html
    

#### Running a project
From within the project type:
```
go build
./projectname runserver
```
