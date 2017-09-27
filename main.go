package main

import (
	"fmt"
	"log"

	"context"

	// Imports the Google Cloud BigQuery client package.
	"cloud.google.com/go/bigquery"
	"google.golang.org/api/iterator"
)

// ClientData Create a query
const ClientData = `
SELECT  *
FROM [bigquery-public-data:usa_names.usa_1910_2013]
WHERE name = 'Jan'
LIMIT 10
`

func main() {
	ctx := context.Background()

	// Sets your Google Cloud Platform project ID.
	projectID := "プロジェクト名"

	// Creates a client.
	client, err := bigquery.NewClient(ctx, projectID)
	if err != nil {
		log.Fatalf("Failed to create client: %v", err)
	}

	// To query existing tables and call its Read method:
	query := client.Query(ClientData)

	// You can also start the query running and get the results later.
	//Create the query as above, but call Run instead of Read. This returns a Job, which represents an asychronous operation.
	it, err := query.Read(ctx)

	if err != nil {
		log.Printf("Failed to Read Query:%v", err)
	}

	i := 1
	// Then iterate through the resulting rows.
	// You can store a row using anything that implements the ValueLoader interface, or with a slice or map of bigquery.Value. A slice is simplest:
	for {
		var values []bigquery.Value

		err := it.Next(&values)
		if err == iterator.Done {
			break
		}

		if err != nil {
			log.Printf("Failed to Iterate Query:%v", err)
		}

		fmt.Printf("[%d]\t%+v\n", i, values)
		i++
	}
}
