package db

import (
	"context"
	"fmt"
	"log"
	"rafaelcopat/models"

	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

func ReadProdutos() []*models.Produto {
	client := ConnectMongoDB()
	coll := client.Database("test").Collection("produtos")
	cur, err := coll.Find(context.TODO(), bson.D{})
	if err != nil {
		fmt.Println(err)
	}

	var results []*models.Produto
	for cur.Next(context.TODO()) {

		var elem models.Produto
		err := cur.Decode(&elem)
		if err != nil {
			log.Fatal(err)
		}
		results = append(results, &elem)
	}
	CloseMongoDB(client)
	return results
}

func CriarNovoProduto(nome, descricao string, precoConvertido float64, quantidadeConvertida int) {
	produto := models.Produto{
		Nome:       nome,
		Descricao:  descricao,
		Preco:      precoConvertido,
		Quantidade: quantidadeConvertida,
	}

	client := ConnectMongoDB()
	collection := client.Database("test").Collection("produtos")
	insertResult, err := collection.InsertOne(context.TODO(), produto)
	if err != nil {
		log.Fatal(err)
	}
	fmt.Println("Inserted a single document: ", insertResult.InsertedID)

	CloseMongoDB(client)

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

	//fmt.Println("Connected to MongoDB!")
	return client
}

func CloseMongoDB(client *mongo.Client) {
	var err = client.Disconnect(context.TODO())

	if err != nil {
		log.Fatal(err)
	}
	//fmt.Println("Connection to MongoDB closed.")
}
