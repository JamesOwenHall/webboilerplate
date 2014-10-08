# webboilerplate

webboilerplate is the boilerplate I use for web apps I write in Go.

## Features

* MVC
* public file serving
* optional template caching
* sessions (via [Gorilla Toolkit](http://www.gorillatoolkit.org))
* logging with [logrus](https://github.com/Sirupsen/logrus)

## Building & Running

You can quickly run the project using `go run main.go` from the project's root.
Since it contains relative import paths, you need to include `main.go` in the
build command: `go build main.go`.

## Structure

At the root of the project, you'll find `main.go`.  Here we capture command
line arguments and start the server.

### public

This is the default directory for files that should be served publicly.

### server

All Go code (other than main.go) is found in this package and its subpackages.

* `server.go` defines our server struct.
* `routes.go` matches routes to controllers.
* `database.go` sets up the database handle.

None of its subpackages should import this package.

#### server/common

This package contains types that should be accessible from any other package
in the project.

* `config.go` contains site configuration information.
* `context.go` defines the context that is passed to the controller on each
request.  This holds useful variables such as a templater, session and database
handle.
* `controller.go` defines the interface for a controller.
* `not_found.go` defines a function that can be called to display a 404 page.
* `templater.go` defines a struct that renders templates with optional caching.
* `view_data.go` defines a common structure to be passed to all templates.

#### server/controllers

This package is where you define your controllers.

* `index.go` is the controller for the `/` route.  It handles displaying the
home page, serving public files and displaying a 404 for undefined routes.

#### server/models

This package is where you define your models.

#### views

This is the directory where you keep your views.  There shouldn't be any Go
code (not including Go's templating code).
