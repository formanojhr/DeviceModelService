package main

import (
	deviceapi "DeviceModelService/devicemodelservice/api"
	"DeviceModelService/devicemodelservice/registry"
	"encoding/json"
	"fmt"
	"github.com/aws/aws-lambda-go/events"
	"github.com/aws/aws-lambda-go/lambda"
	"log"
	"net/http"
	"net/url"
	"os"
)

var errorLogger = log.New(os.Stderr, "ERROR ", log.Llongfile)

// Response is of type APIGatewayProxyResponse since we're leveraging the
// AWS Lambda Proxy Request functionality (default behavior)
//
// https://serverless.com/framework/docs/providers/aws/events/apigateway/#lambda-proxy-integration
type Response events.APIGatewayProxyResponse

// Handler is our lambda handler invoked by the `lambda.Start` function call
//func Handler(ctx context.Context) (Response, error) {
//	var buf bytes.Buffer
//
//	body, err := json.Marshal(map[string]interface{}{
//		"message": "Okay so your other function also executed successfully!",
//	})
//	if err != nil {
//		return Response{StatusCode: 404}, err
//	}
//	json.HTMLEscape(&buf, body)
//
//	resp := Response{
//		StatusCode:      200,
//		IsBase64Encoded: false,
//		Body:            buf.String(),
//		Headers: map[string]string{
//			"Content-Type":           "application/json",
//			"X-MyCompany-Func-Reply": "world-handler",
//		},
//	}
//
//	return resp, nil
//}

func Handler(request events.APIGatewayProxyRequest) (events.APIGatewayProxyResponse, error) {
	fmt.Println("Received body: ", request.Body)

	return events.APIGatewayProxyResponse{Body: request.Body, StatusCode: 200}, nil
}

// a REST HTTP endpoint router determines which action to take
func router(req events.APIGatewayProxyRequest) (events.APIGatewayProxyResponse, error) {
	ctn := createDeviceModelAPIInstance()
	api := ctn.Resolve("device-model-API").(*deviceapi.DeviceModelAPIInstance)
	switch req.HTTPMethod {
	case "GET":
		return get(api, req)
	//case "POST":
	//	return create(req)
	default:
		return clientError(http.StatusMethodNotAllowed)
	}
}

// create device model API container
func createDeviceModelAPIInstance() *registry.DeviceModelContainer {
	ctn, err := registry.NewContainer()
	if err != nil {
		return ctn
	}
	return ctn
}

// process HTTP Get request
func get(api *deviceapi.DeviceModelAPIInstance, req events.APIGatewayProxyRequest) (events.APIGatewayProxyResponse, error) {
	rawName, found := req.QueryStringParameters["name"]
	if found {
		// query parameters are typically URL encoded so to get the value
		name, err := url.QueryUnescape(rawName)
		if err != nil {
			//handle error
			return serverError(err)
		}
		model, err := api.GetModelByName(name)
		if err != nil {
			//handle error
			return serverError(err)
		}
		js, err := json.Marshal(model)
		if err != nil {
			return serverError(err)
		}
		return events.APIGatewayProxyResponse{
			StatusCode: http.StatusOK,
			Body:       string(js),
		}, nil
	} else {
		fmt.Println("Getting all models...")
		models, err := api.GetModels()
		if err != nil {
			return serverError(err)
		}
		js, err := json.Marshal(models)
		return events.APIGatewayProxyResponse{
			StatusCode: http.StatusOK,
			Body:       string(js),
		}, nil
	}
}

// Helper to convert server error into APIGateway error response
func serverError(err error) (events.APIGatewayProxyResponse, error) {
	errorLogger.Println(err.Error())

	return events.APIGatewayProxyResponse{
		StatusCode: http.StatusInternalServerError,
		Body:       http.StatusText(http.StatusInternalServerError),
	}, nil
}

func clientError(status int) (events.APIGatewayProxyResponse, error) {
	return events.APIGatewayProxyResponse{
		StatusCode: status,
		Body:       http.StatusText(status),
	}, nil
}

func main() {
	lambda.Start(router)
}
