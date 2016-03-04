# reinhardt

This is a golang MVC modeled loosely off of Django. Still in pre-alpha.

### Getting Started

##### Installation

    go get github.com/r0fls/reinhardt
    cd $GOPATH/src/github.com/r0fls/reinhardt
    go build
    
##### Starting a new project

    ./reinhardt new <projectname>
    cd <projectname>
    go build
    ./<projectname> runserver
