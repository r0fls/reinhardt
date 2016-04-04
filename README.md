# reinhardt

This is a golang MVC modeled loosely off of Django. Still in pre-alpha.

### Quick setup

```bash
go install github.com/r0fls/reinhardt
reinhardt new <projectname>
cd <projectname>
go install
<projectname> runserver
```

### Overview

#### Installation
```bash
go install github.com/r0fls/reinhardt
```

#### Starting a new project

```bash
reinhardt new <projectname>
cd <projectname>
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
    

#### Compiling a project
From within the project type:
```bash
go install
```
#### Running a project
```bash
<projectname> runserver
```
