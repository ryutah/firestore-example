package main

import (
	"context"
	"log"
	"os"

	"google.golang.org/api/iterator"

	"cloud.google.com/go/firestore"
)

var projectID = os.Getenv("PROJECT_ID")

func main() {
	ctx := context.Background()

	client, err := firestore.NewClient(ctx, projectID)
	if err != nil {
		log.Fatalf("failed to create client: %v", err)
	}
	defer client.Close()

	log.Println("Write users...")
	ref, ret, err := client.Collection("users").Add(ctx, map[string]interface{}{
		"first": "Ada",
		"last":  "Lovelace",
		"born":  1815,
	})
	if err != nil {
		log.Fatalf("failed to add first users: %v", err)
	}
	log.Printf("Add id: %v, Path: %v", ref.ID, ref.Path)
	log.Printf("Ret: %v", ret)

	if _, _, err := client.Collection("users").Add(ctx, map[string]interface{}{
		"first":  "Alan",
		"middle": "Mathison",
		"last":   "Turing",
		"born":   1912,
	}); err != nil {
		log.Fatalf("failed to add second users: %v", err)
	}

	log.Println("Read users...")
	iter := client.Collection("users").Documents(ctx)
	for {
		doc, err := iter.Next()
		if err == iterator.Done {
			break
		} else if err != nil {
			log.Fatalf("failed to read document: %v", err)
		}
		log.Printf("%v:%v; %v\n", doc.Ref.Path, doc.Ref.ID, doc.Data())
	}
}
