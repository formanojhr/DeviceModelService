

This was built using serverless framework which provides abstraction for serverless deployment across cloud providers.

https://www.serverless.com/blog/framework-example-golang-lambda-support
Project Created using go lambda template
$ serverless create -t aws-go-dep -p myservice

The default command in the included Makefile will gather your dependencies and build the proper binaries for your functions.

Build
=====
1.$ make

$ serverless deploy
Serverless: Packaging service...
Serverless: Excluding development dependencies...
Serverless: Uploading CloudFormation file to S3...
Serverless: Uploading artifacts...
Serverless: Uploading service devicemodelservice.zip file to S3 (9.33 MB)...
Serverless: Validating template...
Serverless: Updating Stack...
Serverless: Checking Stack update progress...
..............
Serverless: Stack update finished...
Service Information
service: devicemodelservice
stage: dev
region: us-east-1
stack: devicemodelservice-dev
resources: 11
api keys:
  None
endpoints:
  GET - https://nyho0t9bfb.execute-api.us-east-1.amazonaws.com/dev/devicemodel


 Logging/Monitoring (from AWS)
 ==================
 serverless logs -f devicemodel -t
 START RequestId: be9fe3ab-09d9-4223-9b53-9ae2a9d382f5 Version: $LATEST
 Received body:
 END RequestId: be9fe3ab-09d9-4223-9b53-9ae2a9d382f5
 REPORT RequestId: be9fe3ab-09d9-4223-9b53-9ae2a9d382f5	Duration: 0.53 ms	Billed Duration: 100 ms	Memory Size: 1024 MB	Max Memory Used: 35 MB
