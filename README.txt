

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

Similarly to query model name:

curl http://127.0.0.1:3000/devicemodels?name=Trio
{"modelName":"Trio","ModelNumber":"8300","DeviceType":"IPPHONE","DeviceTypeUUID":"f957617b-e773-4864-95f1-df82a6fe5e09","Vendor":"Polycom","TypePatterns":["8300"]}

AWS Lambda Deployment
=====================
Use the makefile command package and deploy

First time only :
* Create a new S3 bucket for deployment artificats

aws s3 mb s3://device-model-lambda

Package the go binary into zip file and upload to the created S3 bucket

sam package --template-file template.yaml --s3-bucket device-model-lambda --output-template-file packaged.yaml


Now deployment the resources needed for the serverless function(including API Gateway)

sam deploy --template-file packaged.yaml --stack-name devicemodelapi --capabilities CAPABILITY_IAM


Output:

#sam package --template-file  --s3-bucket  --output-template-file 
sam deploy --template-file packaged.yaml --stack-name devicemodelapi --capabilities CAPABILITY_IAM

        Deploying with following values
        ===============================
        Stack name                 : devicemodelapi
        Region                     : None
        Confirm changeset          : False
        Deployment s3 bucket       : None
        Capabilities               : ["CAPABILITY_IAM"]
        Parameter overrides        : {}

Initiating deployment
=====================

Waiting for changeset to be created..

CloudFormation stack changeset
------------------------------------------------------------------------------------------------------
Operation                          LogicalResourceId                  ResourceType                     
------------------------------------------------------------------------------------------------------
+ Add                              DeviceModelFunctionGetRequestPer   AWS::Lambda::Permission          
                                   missionProd                                                         
+ Add                              DeviceModelFunctionRole            AWS::IAM::Role                   
+ Add                              DeviceModelFunction                AWS::Lambda::Function            
+ Add                              ServerlessRestApiDeployment3493b   AWS::ApiGateway::Deployment      
                                   7e1e9                                                               
+ Add                              ServerlessRestApiProdStage         AWS::ApiGateway::Stage           
+ Add                              ServerlessRestApi                  AWS::ApiGateway::RestApi         
------------------------------------------------------------------------------------------------------

Changeset created successfully. arn:aws:cloudformation:us-east-1:466137380188:changeSet/samcli-deploy1598404130/7b376e93-55b7-4d5a-915e-117745060687


2020-08-25 18:09:01 - Waiting for stack create/update to complete

CloudFormation events from changeset
-----------------------------------------------------------------------------------------------------
ResourceStatus            ResourceType              LogicalResourceId         ResourceStatusReason    
-----------------------------------------------------------------------------------------------------
CREATE_IN_PROGRESS        AWS::IAM::Role            DeviceModelFunctionRole   -                       
CREATE_IN_PROGRESS        AWS::IAM::Role            DeviceModelFunctionRole   Resource creation       
                                                                              Initiated               
CREATE_COMPLETE           AWS::IAM::Role            DeviceModelFunctionRole   -                       
CREATE_IN_PROGRESS        AWS::Lambda::Function     DeviceModelFunction       -                       
CREATE_IN_PROGRESS        AWS::Lambda::Function     DeviceModelFunction       Resource creation       
                                                                              Initiated               
CREATE_COMPLETE           AWS::Lambda::Function     DeviceModelFunction       -                       
CREATE_IN_PROGRESS        AWS::ApiGateway::RestAp   ServerlessRestApi         -                       
                          i                                                                           
CREATE_IN_PROGRESS        AWS::ApiGateway::RestAp   ServerlessRestApi         Resource creation       
                          i                                                   Initiated               
CREATE_COMPLETE           AWS::ApiGateway::RestAp   ServerlessRestApi         -                       
                          i                                                                           
CREATE_IN_PROGRESS        AWS::ApiGateway::Deploy   ServerlessRestApiDeploy   -                       
                          ment                      ment3493b7e1e9                                    
CREATE_IN_PROGRESS        AWS::Lambda::Permission   DeviceModelFunctionGetR   -                       
                                                    equestPermissionProd                              
CREATE_IN_PROGRESS        AWS::ApiGateway::Deploy   ServerlessRestApiDeploy   Resource creation       
                          ment                      ment3493b7e1e9            Initiated               
CREATE_IN_PROGRESS        AWS::Lambda::Permission   DeviceModelFunctionGetR   Resource creation       
                                                    equestPermissionProd      Initiated               
CREATE_COMPLETE           AWS::ApiGateway::Deploy   ServerlessRestApiDeploy   -                       
                          ment                      ment3493b7e1e9                                    
CREATE_IN_PROGRESS        AWS::ApiGateway::Stage    ServerlessRestApiProdSt   -                       
                                                    age                                               
CREATE_IN_PROGRESS        AWS::ApiGateway::Stage    ServerlessRestApiProdSt   Resource creation       
                                                    age                       Initiated               
CREATE_COMPLETE           AWS::ApiGateway::Stage    ServerlessRestApiProdSt   -                       
                                                    age                                               
CREATE_COMPLETE           AWS::Lambda::Permission   DeviceModelFunctionGetR   -                       
                                                    equestPermissionProd                              
CREATE_COMPLETE           AWS::CloudFormation::St   devicemodelapi            -                       
                          ack                                                                         
-----------------------------------------------------------------------------------------------------

CloudFormation outputs from deployed stack
------------------------------------------------------------------------------------------------------
Outputs                                                                                              
------------------------------------------------------------------------------------------------------
Key                 DeviceModelAPI                                                                   
Description         API Gateway endpoint URL for Dev environment for First Function                  
Value               https://a1lgy5kwhh.execute-api.us-east-1.amazonaws.com/Dev/devicemodels/         

Key                 DeviceModelFunction                                                              
Description         Device Model Dev ARN                                                             
Value               arn:aws:lambda:us-east-1:466137380188:function:devicemodelapi-                   
DeviceModelFunction-1DPVGB483T42L                                                                    

Key                 DeviceModelFunctionIamRole                                                       
Description         Implicit IAM Role created for Device Model function                              
Value               arn:aws:iam::466137380188:role/devicemodelapi-DeviceModelFunctionRole-           
TMU245643EDJ                                                                                         
------------------------------------------------------------------------------------------------------

Successfully created/updated stack - devicemodelapi in None




This was built using serverless framework which provides abstraction for serverless deployment across cloud providers.

https://www.serverless.com/blog/framework-example-golang-lambda-support
Project Created using go lambda template
$ serverless create -t aws-go-dep -p myservice


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
