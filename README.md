# AWS Elastic Beanstalk Go Demo

## Creating a source bundle

Create a zip file from the root of this package, for example:

`cd $GOPATH/src/github.com/mrichman/elastic-beanstalk-go-demo`

Then, zip from the command line:

`zip ../elastic-beanstalk-go-demo.zip -r * .[^.]*`

or using `git`:

`git archive --format=zip HEAD > ../elastic-beanstalk-go-demo.zip`

## Deploying Go Application on Elastic Beanstalk

For simple Go applications, there are two ways to deploy your application:

* Provide a source bundle with a source file at the root called application.go that contains the main package for your application. Elastic Beanstalk builds the binary using the following command:

`go build -o bin/application application.go`

After the application is built, Elastic Beanstalk starts it on port 5000.

* Provide a source bundle with a binary file called application. The binary file can be located either at the root of the source bundle or in the bin/ directory of the source bundle. If you place the application binary file in both locations, Elastic Beanstalk uses the file in the bin/ directory.

Elastic Beanstalk launches this application on port 5000.

## Local development note

This application attempts to log to `/var/log/golang`. If this directory exists, ensure that it's writable by your user, or run this application using `sudo go run application.go`.

