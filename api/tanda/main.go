package main

import (
	"encoding/json"
	"github.com/aws/aws-lambda-go/events"
	"github.com/aws/aws-lambda-go/lambda"
	"strconv"
)

type MovieRef struct {
	MovieId int    `json:"movie_id"`
	Title  string `json:"title"`
}

type Tanda struct {
	Id          int        `json:"_id"`
	Title       string     `json:"title"`
	Hora		string     `json:"hora"`
	Publico     int        `json:"publico"`
	Peliculas   []MovieRef `json:"peliculas"`
}

var items []Tanda

var jsonData string = `[
	{
		"_id": 1,
		"title": "Tanda de la Mañana",
		"hora":"9:00 am",
		"publico":"PG 13",
		"peliculas": [
			{
				"movie_id": 1,
				"title": "The Great Wall"
			},
			{
				"movie_id": 2,
				"title": "Fantastic Beasts and Where to Find Them"
			}
		]
	},
	{
		"_id": 2,
		"title": "Tanda del Medio Dia",
		"hora":"12:00 pm",
		"publico":"PG 13",
		"peliculas": [
			{
				"movie_id": 1,
				"title": "The Great Wall"
			},
			{
				"movie_id": 2,
				"title": "Fantastic Beasts and Where to Find Them"
			}
		]
	},
	{
		"_id": 3,
		"title": "Tanda de la Tarde",
		"hora":"3:00 pm",
		"publico":"Rj-3",
		"peliculas": [
			{
				"movie_id": 1,
				"title": "The Great Wall"
			},
			{
				"movie_id": 2,
				"title": "Fantastic Beasts and Where to Find Them"
			}
		]
	},
	{
		"_id": 4,
		"title": "Tanda de la Noche",
		"hora":"9:00 pm",
		"publico":"PG 18",
		"peliculas": [
			{
				"movie_id": 1,
				"title": "The Great Wall"
			},
			{
				"movie_id": 2,
				"title": "Fantastic Beasts and Where to Find Them"
			}
		]
	}
]`

func FindItem(id int) *Tanda {
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
