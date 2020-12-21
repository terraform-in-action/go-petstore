package petstore

import (
	"context"
	"errors"
	"fmt"
	"net/url"
)

// Compile-time proof of interface implementation.
var _ Pets = (*pets)(nil)
var ctx = context.Background()

// Pets describes all the pet related methods that the Petstore
// API supports.
type Pets interface {
	// List all the pets.
	List(options PetListOptions) (*PetList, error)

	// Create a new pet with the given options.
	Create(options PetCreateOptions) (*Pet, error)

	// Read an pet by its ID.
	Read(petID string) (*Pet, error)

	// Update an pet by its ID.
	Update(petID string, options PetUpdateOptions) (*Pet, error)

	// Delete an pet by its ID.
	Delete(petID string) error
}

// pets implements Pets.
type pets struct {
	client *Client
}

// PetList represents a list of pets.
type PetList struct {
	Items []*Pet `json:"items"`
}

// Pet represents a Petstore pet.
type Pet struct {
	ID      string `json:"id"`
	Name    string `json:"name"`
	Species string `json:"species"`
	Age     int    `json:"age"`
}

// PetListOptions represents the options for listing pets.
type PetListOptions struct {
	Limit int
}

// List all pets.
func (s *pets) List(options PetListOptions) (*PetList, error) {
	if options.Limit == 0 {
		options.Limit = 100
	}
	req, err := s.client.newRequest("GET", "pets", &options)
	if err != nil {
		return nil, err
	}

	petl := &PetList{}
	err = s.client.do(ctx, req, petl)
	if err != nil {
		return nil, err
	}

	return petl, nil
}

// PetCreateOptions represents the options for creating an pet.
type PetCreateOptions struct {
	Name    string `json:"name"`
	Species string `json:"species"`
	Age     int    `json:"age"`
}

func (p PetCreateOptions) valid() error {
	if !validString(&p.Name) {
		return errors.New("pet name is required")
	}
	if !validString(&p.Species) {
		return errors.New("pet species is required")
	}
	if !validInteger(&p.Age) {
		return errors.New("age is required")
	}
	return nil
}

// Create a new pet with the given options.
func (s *pets) Create(options PetCreateOptions) (*Pet, error) {
	if err := options.valid(); err != nil {
		return nil, err
	}

	req, err := s.client.newRequest("POST", "pets", &options)
	if err != nil {
		return nil, err
	}

	ent := &Pet{}
	err = s.client.do(ctx, req, ent)
	if err != nil {
		return nil, err
	}

	return ent, nil
}

// Read an pet by id.
func (s *pets) Read(petID string) (*Pet, error) {
	if !validStringID(&petID) {
		return nil, errors.New("invalid value for petID")
	}

	u := fmt.Sprintf("pets/%s", url.QueryEscape(petID))
	req, err := s.client.newRequest("GET", u, nil)
	if err != nil {
		return nil, err
	}

	ent := &Pet{}
	err = s.client.do(ctx, req, ent)
	if err != nil {
		return nil, err
	}

	return ent, nil
}

// PetUpdateOptions represents the options for updating an pet.
type PetUpdateOptions struct {
	Name    string `json:"name"`
	Species string `json:"species"`
	Age     int    `json:"age"`
}

// Update attributes of an existing pet.
func (s *pets) Update(petID string, options PetUpdateOptions) (*Pet, error) {
	if !validStringID(&petID) {
		return nil, errors.New("invalid value for petID")
	}

	u := fmt.Sprintf("pets/%s", url.QueryEscape(petID))
	req, err := s.client.newRequest("PATCH", u, &options)
	if err != nil {
		return nil, err
	}

	ent := &Pet{}
	err = s.client.do(ctx, req, ent)
	if err != nil {
		return nil, err
	}

	return ent, nil
}

// Delete an pet by its ID.
func (s *pets) Delete(petID string) error {
	if !validStringID(&petID) {
		return errors.New("invalid value for petID")
	}

	u := fmt.Sprintf("pets/%s", url.QueryEscape(petID))
	req, err := s.client.newRequest("DELETE", u, nil)
	if err != nil {
		return err
	}

	return s.client.do(ctx, req, nil)
}
