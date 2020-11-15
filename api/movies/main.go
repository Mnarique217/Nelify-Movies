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


type Movie struct {
	Id           int       `json:"_id"`
	Rank         string    `json:"rank"`
	Title		 string    `json:"title"`
	Description  int       `json:"description"`
	Director     string    `json:"director"`
	Idioma       string    `json:"idioma"`
	Tandas      []TandaRef `json:"tandas"`	
}


var items []Movie

var jsonData string = ` [
	{
		"_id": 1,
		"rank": 1,
		"title": "Guardians of the Galaxy",
		"description": "A group of intergalactic criminals are forced to work together to stop a fanatical warrior from taking control of the universe.",
		"Director": "Bin Diesel",
		"Idioma": "Español",
		"tandas": [
            {"tanda_id":1,
			"tanda_name":"Tanda de la Mañana"},
			{"tanda_id":2,
				"tanda_name":"Tanda del Medio Dia"}
        ]
	},
	{
		"_id": 2,
		"rank": 1,
		"title": "The Great Wall",
		"description": "Gloria is an out-of-work party girl forced to leave her life in New York City, and move back home. When reports surface that a giant creature is destroying Seoul, she gradually comes to the realization that she is somehow connected to this phenomenon.",
		"Director": "Bin Diesel",
		"Idioma": "Ingles",
		"tandas": [
            {"tanda_id":1,
			"tanda_name":"Mañana"},
			{"tanda_id":2,
				"tanda_name":"Tarde"}
        ]
	},
	{
		"_id": 3,
		"rank": 1,
		"title": "Fantastic Beasts and Where to Find Them",
		"description": "The adventures of writer Newt Scamander in New York's secret community of witches and wizards seventy years before Harry Potter reads his book in school.",
		"Director": "Ridley Scott",
		"Idioma": "Chino",
		"tandas": [
            {"tanda_id":3,
			"tanda_name":"Tanda de la Tarde"},
			{"tanda_id":4,
				"tanda_name":"Tanda de la Noche"}
        ]
	},
	{
		"_id": 4,
		"rank": 1,
		"title": "Assassin's Creed",
		"description": "When Callum Lynch explores the memories of his ancestor Aguilar and gains the skills of a Master Assassin, he discovers he is a descendant of the secret Assassins society.",
		"Director": "Sean Foley",
		"Idioma": "Español",
		"tandas": [
            {"tanda_id":1,
			"tanda_name":"Mañana"},
			{"tanda_id":2,
				"tanda_name":"Tarde"}
        ]
	},
	{
		"_id": 5,
		"rank": 1,
		"title": "The Beginning",
		"description": "A cold-blooded predatory couple while cruising the streets in search of their next victim, will stumble upon a 17-year-old high school girl, who will be sedated, abducted and chained in the strangers' guest room.",
		"Director": "Bin Diesel",
		"Idioma": "Español",
		"tandas": [
            {"tanda_id":3,
			"tanda_name":"Tanda de la Tarde"},
			{"tanda_id":4,
				"tanda_name":"Tanda de la Noche"}
        ]
	},
	{
		"_id": 6,
		"rank": 1,
		"title": "Awake",
		"description": "A has-been actor best known for playing the title character in the 1980s detective series  must work with the police when a serial killer says that he will only speak with Detective Mindhorn, whom he believes to be a real person.",
		"Director": "Samuel Jackson",
		"Idioma": "Aleman",
		"tandas": [
            {"tanda_id":3,
			"tanda_name":"Tanda de la Tarde"},
			{"tanda_id":4,
				"tanda_name":"Tanda de la Noche"}
        ]
	}
]`

func FindItem(id int) *Movie {
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
