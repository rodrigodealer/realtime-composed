package es

import (
	"context"
	"log"
	"os"

	elastic "gopkg.in/olivere/elastic.v5"
)

type ElasticSearch interface {
	Connect()
	Ping() int
}

type EsClient struct {
	Client *elastic.Client
}

func (e *EsClient) Connect() {
	ctx := context.Background()
	var url = config(os.Getenv("ELASTICSEARCH_URL"))
	client, err := elastic.NewClient(elastic.SetURL(url), elastic.SetSniff(false))
	if err != nil {
		log.Panic(err)
	}

	info, code, _ := client.Ping(url).Do(ctx)
	log.Printf("Elasticsearch returned with code %d and version %s", code, info.Version.Number)
	e.Client = client
}

func (e *EsClient) Ping() int {
	ctx := context.Background()
	var url = config(os.Getenv("ELASTICSEARCH_URL"))
	_, code, _ := e.Client.Ping(url).Do(ctx)
	return code
}

func (e *EsClient) IndexSetup(index string) {
	ctx := context.Background()
	exists, err := e.Client.IndexExists(index).Do(ctx)
	if err != nil {
		log.Panic(err)
	}
	if !exists {
		_, err := e.Client.CreateIndex(index).Do(ctx)
		if err != nil {
			log.Panic(err)
		} else {
			log.Printf("Index %s created", index)
		}
	} else {
		log.Printf("Index %s already exists, skipping", index)
	}
}
