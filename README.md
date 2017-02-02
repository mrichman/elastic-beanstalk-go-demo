# AWS Elastic Beanstalk Go Demo

## Creating a source bundle

Create a zip file from the root of this package, for example:

`cd $GOPATH/src/github.com/mrichman/elastic-beanstalk-go-demo`

Then, either zip it from the command line:

`zip ../eb.zip -r * .[^.]*`

or use `git`:

`git archive --format=zip HEAD > ../eb.zip`

The git command above will zip up the last commit on master.

## Deploying a Go Application on Elastic Beanstalk

Browse to the [Elastic Beanstalk console](https://console.aws.amazon.com/elasticbeanstalk/), and click "Create New Application".

For simple Go applications, there are two ways to deploy your application:

* **Method 1:** Provide a source bundle with a source file at the root called application.go that contains the main package for your application. Elastic Beanstalk automatically builds the binary using the following command at deployment time:

    `go build -o bin/application application.go`

* **Method 2:** Provide a source bundle with a binary file called application. The binary file can be located either at the root of the source bundle or in the bin/ directory of the source bundle. If you place the application binary file in both locations, Elastic Beanstalk uses the file in the bin/ directory.

By default, Elastic Beanstalk configures the nginx proxy to forward requests to your application on port 5000. You can override the default port by setting the `PORT` system property to the port on which your main application listens.
