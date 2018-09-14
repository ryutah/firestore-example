package main

import (
	"context"
	"log"
	"os"

	"cloud.google.com/go/firestore"
)

var projectID = os.Getenv("PROJECT_ID")

type foo struct {
	Name string
}

type bar struct {
	Name string
	Foo1 foo
	Foo2 *foo
}

func main() {
	ctx := context.Background()

	client, err := firestore.NewClient(ctx, projectID)
	if err != nil {
		log.Fatalf("failed to create client: %v", err)
	}
	defer client.Close()

	var (
		b = bar{
			Name: "bar",
			Foo1: foo{Name: "foo1"},
			Foo2: &foo{Name: "foo2"},
		}
		b2 = bar{
			Name: "bar",
		}
	)

	log.Println("Write bar...")
	ref, ret, err := client.Collection("bar").Add(ctx, b)
	if err != nil {
		log.Fatalf("failed to add first users: %v", err)
	}
	log.Printf("Add id: %v, Path: %v", ref.ID, ref.Path)
	log.Printf("Ret: %v", ret)

	if _, _, err := client.Collection("bar").Add(ctx, b2); err != nil {
		log.Fatalf("failed to add second users: %v", err)
	}
}
