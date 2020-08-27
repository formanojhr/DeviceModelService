.PHONY: build clean deploy

#This command specifically compiles your golang binary for Linux, as that's the Lambda runtime you get (specifically AWS Linux).
# If you forget to do this on a Mac (like I did originally) you'll get a really strange error.
build:
	dep ensure -v
#    env GOOS=linux go build -ldflags="-s -w" -o .aws-sam/build/serverless serverless/main.go
	env GOOS=linux go build -ldflags="-d -s -w" -o .aws-sam/build/devicemodelfunction/devicemodel devicemodelservice/devicemodel/main.go
	chmod 777 .aws-sam/build/devicemodelfunction/devicemodel
	cp template.yaml .aws-sam/build/devicemodelfunction
	zip -j .aws-sam/build/devicemodelfunction/devicemodelservice.zip .aws-sam/build/devicemodelfunction/devicemodel
#	chmod 777 .aws-sam/build/devicemodelfunction/devicemodelservice.zip
	#sam build

.PHONY: build
#build:
#	sam build
	#env GOOS=linux go build -ldflags="-s -w" -o bin/world world/main.go
#    ${GOPATH}\bin\build-lambda-zip.exe -o bin/serverless serverless/main.go
#	${GOPATH}\bin\build-lambda-zip.exe -o bin/serverless serverless/main.go

#Runs all the unit tests.
.PHONY: test
test:
	go test ./...

clean:
	rm -rf ./bin ./vendor Gopkg.lock
	rm -rf .aws-sam/build

#deploy: clean build
#	sls deploy --verbose

# compile the code to run in Lambda (local or real)
.PHONY: lambdals
lambda:
	GOOS=linux GOARCH=amd64 $(MAKE) main

#Calls the sam package command (which is an alias for aws cloudformation package) to push your code to S3 for deployment.
.PHONY: package
# package the folder with saml template and upload to an s3 bucket. NOTE: sam deploy now implicitly performs the functionality of sam package.
#You can use the sam deploy command directly to package and deploy your application. https://docs.aws.amazon.com/serverless-application-model/latest/developerguide/sam-cli-command-reference-sam-package.html
package: build
	sam package --template-file template.yaml --s3-bucket device-model-lambda --output-template-file packaged.yaml
	#sam package --template-file $(TEMPLATE) --s3-bucket $(S3_BUCKET) --output-template-file $(PACKAGED_TEMPLATE)

.PHONY: api
api: build
	sam local start-api

.PHONY: deploy
deploy: package
	sam deploy --template-file packaged.yaml --stack-name devicemodelapi --capabilities CAPABILITY_IAM

createrole:
	aws iam create-role --role-name lambda-devicemodel-executor --assume-role-policy-document file://./trust-policy.json