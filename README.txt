

This was built using serverless framework which provides abstraction for serverless deployment across cloud providers.

https://www.serverless.com/blog/framework-example-golang-lambda-support
Project Created using go lambda template
$ serverless create -t aws-go-dep -p myservice

The default command in the included Makefile will gather your dependencies and build the proper binaries for your functions.

Build
=====
1.devicemodelservice mramakrishnan$ make build
  Makefile:44: warning: overriding commands for target `deploy'
  Makefile:26: warning: ignoring old commands for target `deploy'
  dep ensure -v
  env GOOS=linux go build -ldflags="-s -w" -o .aws-sam/build/devicemodelfunction/devicemodel devicemodel/main.go
Now you should see an updated devicemodel binary under .aws-sam/devicemodelfunction/devicemodel


Local Deploy
============
make api
Makefile:44: warning: overriding commands for target `deploy'
Makefile:26: warning: ignoring old commands for target `deploy'
dep ensure -v
env GOOS=linux go build -ldflags="-s -w" -o .aws-sam/build/devicemodelfunction/devicemodel devicemodel/main.go
sam local start-api


Local Test API
==============
Hit curl -v http://127.0.0.1:3000/devicemodels
Response :
*   Trying 127.0.0.1...
* TCP_NODELAY set
* Connected to 127.0.0.1 (127.0.0.1) port 3000 (#0)
> GET /devicemodels HTTP/1.1
> Host: 127.0.0.1:3000
> User-Agent: curl/7.54.0
.........
[{"modelName":"Trio","ModelNumber":"8300","DeviceType":"IPPHONE","DeviceTypeUUID":"f957617b-e773-4864-95f1-df82a6fe5e09","Vendor":"Polycom","TypePatterns":null},{"modelName":"IP Phone","ModelNumber":"7800",
"DeviceType":"IPPHONE","DeviceTypeUUID":"89c4e433-34ac-4063-8c03-e32d08e1b748","Vendor":"CISCO","TypePatterns":null}]


On process console you should see the serverless go directory being mounted :


Mounting DeviceModelFunction at http://127.0.0.1:3000/devicemodels [GET]
You can now browse to the above endpoints to invoke your functions. You do not need to restart/reload SAM CLI while working on your functions, changes will be reflected instantly/automatically. You only need to restart SAM CLI if you update your AWS SAM template
2020-08-18 14:52:14  * Running on http://127.0.0.1:3000/ (Press CTRL+C to quit)
Invoking devicemodel (go1.x)

Fetching lambci/lambda:go1.x Docker container image......
Mounting /Users/mramakrishnan/Documents/workspaces/go/src/DeviceModelService/devicemodelservice/.aws-sam/build/DeviceModelFunction as /var/task:ro,delegated inside runtime container
START RequestId: 9cf64e41-3867-154b-3145-c4c8745c1a94 Version: $LATEST
Getting all models...
END RequestId: 9cf64e41-3867-154b-3145-c4c8745c1a94
REPORT RequestId: 9cf64e41-3867-154b-3145-c4c8745c1a94  Init Duration: 130.50 ms        Duration: 4.10 ms       Billed Duration: 100 ms Memory Size: 128 MB     Max Memory Used: 22 MB
No Content-Type given. Defaulting to 'application/json'.
2020-08-18 14:52:18 127.0.0.1 - - [18/Aug/2020 14:52:18] "GET /devicemodels HTTP/1.1" 200 -






SERVERLESS framework deployment

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
