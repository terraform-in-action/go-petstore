package main

import (
	"fmt"
	"log"

	"github.com/davecgh/go-spew/spew"
	sdk "github.com/scottwinkler/go-petstore"
)

func main() {
	config := &sdk.Config{
		//Address: "insert-your-petstore-address-here",
		Address: "https://g6pny6dke9.execute-api.us-west-2.amazonaws.com/petstore",
	}

	client, err := sdk.NewClient(config)
	if err != nil {
		log.Fatal(err)
	}

	// List all Pets
	petL, _ := client.Pets.List(sdk.PetListOptions{})
	spew.Printf("petList: %v", petL)

	// Create a new pet
	options := sdk.PetCreateOptions{
		Name:    "mittens",
		Species: "cat",
		Age:     2,
	}

	pet, err := client.Pets.Create(options)
	if err != nil {
		log.Fatal(err)
	}

	// Update pet
	pet, _ = client.Pets.Update(pet.ID, sdk.PetUpdateOptions{Age: 3})

	// Read pet by ID
	pet, _ = client.Pets.Read(pet.ID)
	fmt.Printf("pet: %v \n", pet)

	// Delete a pet
	err = client.Pets.Delete(pet.ID)
	if err != nil {
		log.Fatal(err)
	}

	// List all Pets
	petL, _ = client.Pets.List(sdk.PetListOptions{})
	spew.Printf("petList: %v", petL)
}
