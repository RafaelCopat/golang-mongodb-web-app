package main

import (
	"context"
	"fmt"
	"html/template"
	"log"
	"net/http"

	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

type Produto struct {
	Nome       string
	Descricao  string
	Preco      float64
	Quantidade int
}

var templates = template.Must(template.ParseGlob("templates/*.html"))

func main() {
	http.HandleFunc("/", index)
	http.ListenAndServe(":8000", nil)

}

func index(w http.ResponseWriter, r *http.Request) {

	// Set client options
	// Connect to MongoDB
	// Check the connection
	client := ConnectMongoDB()

	// prints 'document is nil'
	// create a value into which the single document can be decoded
	results := readProdutos(client)

	templates.ExecuteTemplate(w, "Index", results)

	CloseMongoDB(client)
}

func readProdutos(client *mongo.Client) []*Produto {
	coll := client.Database("test").Collection("produtos")
	cur, err := coll.Find(context.TODO(), bson.D{})
	if err != nil {
		fmt.Println(err)
	}

	var results []*Produto
	for cur.Next(context.TODO()) {

		var elem Produto
		err := cur.Decode(&elem)
		if err != nil {
			log.Fatal(err)
		}

		results = append(results, &elem)
	}
	return results
}

func ConnectMongoDB() *mongo.Client {
	clientOptions := options.Client().ApplyURI("mongodb://localhost:27017")

	client, err := mongo.Connect(context.TODO(), clientOptions)

	if err != nil {
		log.Fatal(err)
	}

	err = client.Ping(context.TODO(), nil)

	if err != nil {
		log.Fatal(err)
	}

	fmt.Println("Connected to MongoDB!")
	return client
}

func CloseMongoDB(client *mongo.Client) {
	var err = client.Disconnect(context.TODO())

	if err != nil {
		log.Fatal(err)
	}
	fmt.Println("Connection to MongoDB closed.")
}
