package main

import (
	"encoding/json"
	"github.com/aws/aws-lambda-go/events"
	"github.com/aws/aws-lambda-go/lambda"
	"strconv"
)

type TandaRef struct {
	TandaId    int    `json:"tanda_id"`
	TandaName  string `json:"tanda_name"`
}

type Sala struct {
	Id           int       `json:"_id"`
	NumeroSala   string    `json:"numeroSala"`
	Capacidad	 string    `json:"capacidad"`
	NombreSala   int       `json:"nombreSala"`
	Tandas     []TandaRef `json:"tandas"`	
}


var items []Sala

var jsonData string = `[
	{
		"_id": 1,
		"numeroSala": 1,
		"capacidad": 15,
		"nombreSala": "Sala 1",
		"tandas": [
            {"tanda_id":1,
			"tanda_name":"Tanda de la Mañana"},
			{"tanda_id":2,
				"tanda_name":"Tanda del Medio Dia"}
        ]
	},
	{
		"_id": 2,
		"numeroSala": 1,
		"capacidad": 15,
		"nombreSala": "Sala 2",
		"tandas": [
            {"tanda_id":1,
			"tanda_name":"Mañana"},
			{"tanda_id":2,
				"tanda_name":"Tarde"}
        ]
	},
	{
		"_id": 3,
		"numeroSala": 1,
		"capacidad": 15,
		"nombreSala": "Sala 3",
		"tandas": [
            {"tanda_id":3,
			"tanda_name":"Tanda de la Tarde"},
			{"tanda_id":4,
				"tanda_name":"Tanda de la Noche"}
        ]
	},
	{
		"_id": 4,
		"numeroSala": 1,
		"capacidad": 15,
		"nombreSala": "Sala 4",
		"tandas": [
            {"tanda_id":1,
			"tanda_name":"Mañana"},
			{"tanda_id":2,
				"tanda_name":"Tarde"}
        ]
	}
]`

func FindItem(id int) *Sala {
	for _, item := range items {
		if item.Id == id {
			return &item
		}
	}
	return nil
}

func handler(req events.APIGatewayProxyRequest) (*events.APIGatewayProxyResponse, error) {
	id := req.QueryStringParameters["id"]
	var data []byte
	if id == "" {
		data, _ = json.Marshal(items)
	} else {
		param, err := strconv.Atoi(id)
		if err == nil {
			item := FindItem(param)
			if item != nil {
				data, _ = json.Marshal(*item)
			} else {
				data = []byte("error\n")
			}
		}
	}
	return &events.APIGatewayProxyResponse{
		StatusCode:      200,
		Headers:         map[string]string{"Content-Type": "application/json"},
		Body:            string(data),
		IsBase64Encoded: false,
	}, nil
}

func main() {
	_ = json.Unmarshal([]byte(jsonData), &items)
	lambda.Start(handler)
}
