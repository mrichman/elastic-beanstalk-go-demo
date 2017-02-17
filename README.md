# AWS Elastic Beanstalk Go Demo

Video demo: https://markrichman.com/2017/02/17/aws-elastic-beanstalk-go/

## Creating a source bundle

Create a zip file from the root of this package, for example:

`cd $GOPATH/src/github.com/mrichman/elastic-beanstalk-go-demo`

Then, either zip it from the command line:

`zip ../eb.zip -r * .[^.]*`

or use `git`:

`git archive -o ../eb.zip HEAD`

The git command above will zip up the last commit on master.

## Deploying a Go Application on Elastic Beanstalk

Browse to the [Elastic Beanstalk console](https://console.aws.amazon.com/elasticbeanstalk/), and click "Create New Application".

Give it a name, for example "Go Demo". Then create a new web server environment. For "Platform", select "Go", then select "Upload your code".

For simple Go applications, there are two ways to deploy your application:

* **Method 1:** Provide a source bundle with a source file at the root called application.go that contains the main package for your application. Elastic Beanstalk automatically builds the binary using the following command at deployment time:

    `go build -o bin/application application.go`

* **Method 2:** Provide a source bundle with a binary file called application. The binary file can be located either at the root of the source bundle or in the bin/ directory of the source bundle. If you place the application binary file in both locations, Elastic Beanstalk uses the file in the bin/ directory.

By default, Elastic Beanstalk configures the nginx proxy to forward requests to your application on port 5000. You can override the default port by setting the `PORT` system property to the port on which your main application listens.

For "Source code origin", choose "Local file" and pick the `eb.zip` file created above. Enter anything for "Version label".

Click "Upload", then back at the "Create a new environment" screen, click "Create environment".

Your application will begin to deploy. This process can take several minutes.

When the deployment is complete, you'll be redirected to your Elastic Beanstalk application's overview page. In the breadcrumb trail at the top, next to your Environment ID, you'll see a URL in the control panel ending in `elasticbeanstalk.com` where you can browse your application.

To clean up after yourself, and terminate your application, click the Actions menu and select "Terminate Environment". This will terminate any EC2 instances created during deployment. You can restart your application by selecting Actions -> Create Environment again.
