package db

import (
	"DeviceModelService/devicemodelservice/domain/model"
	"fmt"
	"github.com/aws/aws-sdk-go/aws"
	"github.com/aws/aws-sdk-go/aws/credentials"
	"github.com/aws/aws-sdk-go/aws/session"
	"github.com/aws/aws-sdk-go/service/dynamodb"
	"github.com/aws/aws-sdk-go/service/dynamodb/dynamodbattribute"
	"github.com/aws/aws-sdk-go/service/dynamodb/dynamodbiface"
	"github.com/google/uuid"
	"log"
	"os"
)

// a dynamo db based device model repository
type DeviceModelDBRepository struct {
	Db dynamodbiface.DynamoDBAPI // a dynamo db client
}

const tname string = "DeviceModels"
const modelNameIndex string = "ModelName-index"

//constructor
func New() *DeviceModelDBRepository {
	log.Print("Instantiating DeviceModelDBRepository.")
	// Initialize a session that the SDK will use to load
	// credentials from the shared credentials file ~/.aws/credentials
	// and region from the shared configuration file ~/.aws/config.
	sess := session.Must(session.NewSessionWithOptions(session.Options{
		SharedConfigState: session.SharedConfigEnable,
	}))

	// Create DynamoDB client
	svc := dynamodb.New(sess)
	if svc != nil {
		log.Print("Initialized successfully DynamoDb client.")
	}
	// send back a new ref to DeviceModelDBRepository
	dmd := &DeviceModelDBRepository{
		Db: dynamodbiface.DynamoDBAPI(svc),
	}
	return dmd
}

//get all the models from dynamo db
func (r *DeviceModelDBRepository) FindAll() ([]*model.DeviceModel, error) {
	log.Printf("Fetching all items...from %s table\n", tname)
	// fetch all items
	params := &dynamodb.ScanInput{
		TableName: aws.String(tname),
	}
	allItems, err := r.Db.Scan(params)
	if err != nil {
		fmt.Errorf("failed to make Query API call, %v", err)
		return nil, err
	}

	var devmdls []model.DeviceModel
	err = dynamodbattribute.UnmarshalListOfMaps(allItems.Items, &devmdls)

	if err != nil {
		fmt.Errorf("failed to unmarshal Query result items, %v", err)
		return nil, err
	}
	deviceModels := make([]*model.DeviceModel, len(devmdls))
	i := 0
	for _, deviceModel := range devmdls {
		deviceModels[i] = model.NewDeviceModel(deviceModel.ModelName, deviceModel.ModelNumber, deviceModel.DeviceType,
			deviceModel.DeviceTypeUUID, deviceModel.Vendor, deviceModel.TypePatterns)
		i++
	}
	return deviceModels, nil
}

// find the device model by name
func (r *DeviceModelDBRepository) FindByModelName(name string) (*model.DeviceModel, error) {
	//var queryInput = &dynamodb.QueryInput{
	//	TableName: aws.String(tname),
	//	IndexName: aws.String(modelNameIndex),
	//	KeyConditions: map[string]*dynamodb.Condition{
	//		"modelName": {
	//			ComparisonOperator: aws.String("EQ"),
	//			AttributeValueList: []*dynamodb.AttributeValue{
	//				{
	//					S: aws.String(name),
	//				},
	//			},
	//		},
	//	},
	//}
	//var resp1, err1 = r.Db.Query(queryInput)
	//
	//deviceModels := []model.DeviceModel{}
	//
	//if err1 != nil {
	//	fmt.Println(err1)
	//	return nil, err1
	//} else {
	//	var err = dynamodbattribute.UnmarshalListOfMaps(resp1.Items, &deviceModels)
	//	log.Println(deviceModels)
	//	if err != nil {
	//		return nil, err
	//	}
	//}
	//for _, deviceModel := range deviceModels {
	//	if deviceModel.ModelName == name {
	//		return &deviceModel, nil
	//	}
	//}
	//return nil, nil
	//TODO not using GSI since it might cost money figure this out later. For querying all items
	// and programtically finding the one with match
	log.Printf("Fetching all items...from %s table\n", tname)
	// fetch all items
	params := &dynamodb.ScanInput{
		TableName: aws.String(tname),
	}
	allItems, err := r.Db.Scan(params)
	if err != nil {
		fmt.Errorf("failed to make Query API call, %v", err)
		return nil, err
	}

	var devmdls []model.DeviceModel
	err = dynamodbattribute.UnmarshalListOfMaps(allItems.Items, &devmdls)

	if err != nil {
		fmt.Errorf("failed to unmarshal Query result items, %v", err)
		return nil, err
	}

	i := 0
	for _, deviceModel := range devmdls {
		if deviceModel.ModelName == name { // check if the input requested matches
			fmt.Println(deviceModel)
			return &deviceModel, nil
		}
		i++
	}
	return nil, nil
}

// find the device model by name
func (r *DeviceModelDBRepository) Save(model *model.DeviceModel) (error, model.DeviceModel) {
	u := uuid.New()
	model.DeviceTypeUUID = u.String() // create and update UUID
	av, err := dynamodbattribute.MarshalMap(model)
	if err != nil {
		fmt.Println("Got error marshalling map:")
		fmt.Println(err.Error())
		os.Exit(1)
	}

	fmt.Printf("UUID %s generated for model  : %s\n", u, model.ModelName)
	fmt.Printf("Saving %s.. \n", model.ModelName)
	// Create item in table Device Model
	input := &dynamodb.PutItemInput{
		Item:      av,
		TableName: aws.String(tname),
	}
	_, err = r.Db.PutItem(input)
	if err != nil {
		fmt.Println("Got error calling DynamoDB PutItem:")
		fmt.Println(err.Error())
		return err, *model
	}
	fmt.Printf("Success saving device model %s typeUUID %s \n", model.ModelName, model.DeviceTypeUUID)
	return nil, *model
}

// find the device model by name
func (r *DeviceModelDBRepository) UpdateModel(modelUpdate model.DeviceModel, oldModel model.DeviceModel) error {
	// NOTE this is a simple update where old resource is replaced by new one
	if modelUpdate.DeviceTypeUUID == "" {
		modelUpdate.DeviceTypeUUID = oldModel.DeviceTypeUUID // update UUID from pre existing model copy in db
	}
	fmt.Println(modelUpdate)
	av, err := dynamodbattribute.MarshalMap(modelUpdate)
	if err != nil {
		fmt.Println("Got error marshalling map:")
		fmt.Println(err.Error())
		os.Exit(1)
	}
	fmt.Printf("Updating %s.. \n", modelUpdate.ModelName)
	// Create item in table Device Model
	input := &dynamodb.PutItemInput{
		Item:      av,
		TableName: aws.String(tname),
	}
	_, err = r.Db.PutItem(input)
	if err != nil {
		fmt.Println("Got error calling DynamoDB PutItem:")
		fmt.Println(err.Error())
		return err
	}
	fmt.Printf("Success updating device model %s typeUUID %s \n", modelUpdate.ModelName, modelUpdate.DeviceTypeUUID)
	return nil
}

func (r *DeviceModelDBRepository) DeleteItem(modelName string) (error, string) {
	// Create item in table Device Model
	input := &dynamodb.DeleteItemInput{
		Key: map[string]*dynamodb.AttributeValue{
			"ModelName": {
				N: aws.String(modelName),
			},
		},
		TableName: aws.String(tname),
	}
	_, err := r.Db.DeleteItem(input)
	if err != nil {
		fmt.Println("Got error calling DynamoDB PutItem:")
		fmt.Println(err.Error())
		return err, "Got error calling DynamoDB PutItem"
	}
	fmt.Printf("Success saving device model %s \n", modelName)
	return nil, "Success saving device model"
}

func main() {
	sess, err := session.NewSession(&aws.Config{
		Region:      aws.String("us-east-1"),
		Credentials: credentials.NewSharedCredentials("/Users/mramakrishnan/.aws/credentials", ""),
	})
	svc := dynamodb.New(sess)

	if err != nil {
		log.Fatal("Error initializing Dynamo DB", err)
	}
	// print describe details related to a table
	req := &dynamodb.DescribeTableInput{
		TableName: aws.String("DeviceModels"),
	}
	result, err := svc.DescribeTable(req)
	if err != nil {
		fmt.Printf("%s", err)
	}
	table := result.Table
	fmt.Printf("done", table)

	fmt.Println("Fetching all items...from DeviceModel table")
	// fetch all items
	params := &dynamodb.ScanInput{
		TableName: aws.String("DeviceModels"),
	}
	allItems, err := svc.Scan(params)
	if err != nil {
		fmt.Errorf("failed to make Query API call, %v", err)
	}

	obj := []model.DeviceModel{}
	err = dynamodbattribute.UnmarshalListOfMaps(allItems.Items, &obj)

	if err != nil {
		fmt.Errorf("failed to unmarshal Query result items, %v", err)
	}
	//Following will printout the whole object ,
	//this is just for testing,If you have lots of data, then probably don’t try this.
	fmt.Println(obj)

}
