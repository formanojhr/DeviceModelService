.PHONY: build clean deploy

#This command specifically compiles your golang binary for Linux, as that's the Lambda runtime you get (specifically AWS Linux).
# If you forget to do this on a Mac (like I did originally) you'll get a really strange error.
build:
	dep ensure -v
#    env GOOS=linux go build -ldflags="-s -w" -o .aws-sam/build/serverless serverless/main.go
	env GOOS=linux go build -ldflags="-s -w" -o .aws-sam/build/devicemodelfunction/devicemodel devicemodel/main.go

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

deploy: clean build
	sls deploy --verbose

# compile the code to run in Lambda (local or real)
.PHONY: lambda
lambda:
	GOOS=linux GOARCH=amd64 $(MAKE) main

#Calls the sam package command (which is an alias for aws cloudformation package) to push your code to S3 for deployment.
.PHONY: package
package: build
	sam package --template-file $(TEMPLATE) --s3-bucket $(S3_BUCKET) --output-template-file $(PACKAGED_TEMPLATE)

.PHONY: api
api: build
	sam local start-api

.PHONY: deploy
deploy: package
	sam deploy --stack-name $(STACK_NAME) --template-file $(PACKAGED_TEMPLATE) --capabilities CAPABILITY_IAM