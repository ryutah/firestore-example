package main

import (
	"context"
	"log"
	"os"

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

	log.Println("Write parent...")
	ref, ret, err := client.Collection("parent").Add(ctx, map[string]interface{}{
		"value": "foo",
	})
	if err != nil {
		log.Fatalf("failed to add parent: %v", err)
	}
	log.Printf("Add id: %v, Path: %v", ref.ID, ref.Path)
	log.Printf("Ret: %v", ret)

	ref, _, err = ref.Collection("child").Add(ctx, map[string]interface{}{
		"value": "bar",
	})
	if err != nil {
		log.Fatalf("failed to add child: %v", err)
	}
	log.Printf("Add id: %v, Path: %v", ref.ID, ref.Path)
	log.Printf("Ret: %v", ret)
}
