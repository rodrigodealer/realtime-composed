package es

import (
	"context"
	"log"
	"os"

	elastic "gopkg.in/olivere/elastic.v5"
)

func Connect() *elastic.Client {
	ctx := context.Background()
	var url = config(os.Getenv("ELASTICSEARCH_URL"))
	client, err := elastic.NewClient(
		elastic.SetURL(url),
		elastic.SetSniff(false))
	if err != nil {
		log.Panic(err)
	}

	info, code, _ := client.Ping(url).Do(ctx)
	log.Printf("Elasticsearch returned with code %d and version %s", code, info.Version.Number)
	return client
}

func IndexSetup(client *elastic.Client, index string) {
	ctx := context.Background()
	exists, err := client.IndexExists(index).Do(ctx)
	if err != nil {
		log.Panic(err)
	}
	if !exists {
		_, err := client.CreateIndex(index).Do(ctx)
		if err != nil {
			log.Panic(err)
		} else {
			log.Printf("Index %s created", index)
		}
	} else {
		log.Printf("Index %s already exists, skipping", index)
	}
}
