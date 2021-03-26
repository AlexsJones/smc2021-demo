package main

import (
	"fmt"
	"github.com/AlexsJones/go-openapi-client/client/user"
	"github.com/AlexsJones/go-openapi-client/models"
	"github.com/Pallinder/go-randomdata"
	httptransport "github.com/go-openapi/runtime/client"
	"github.com/go-openapi/strfmt"
	"os"
	"time"

	apiclient "github.com/AlexsJones/go-openapi-client/client"
)

func main() {

	// make the request to get all items
	host := os.Getenv("GO-OPENAPI-HOST")
	if host == "" {
		host = "localhost:8080"
	}
	fmt.Printf("Using host %s", host)
	transport := httptransport.New(host, "/v2", nil)

	// create the API client, with the transport
	client := apiclient.New(transport, strfmt.Default)

	for {
	u := user.NewCreateUserParams()
	u.Body = &models.User{
		Email:          randomdata.Email(),
		FirstName:      randomdata.FullName(randomdata.RandomGender),
		FoodPreference: nil,
		ID:             0,
		LastName:       randomdata.LastName(),
		Password:       randomdata.SillyName(),
		Phone:          randomdata.PhoneNumber(),
		UserStatus:     0,
		Username:       randomdata.SillyName(),
	}
	fmt.Printf("\n%v",u.Body)
	err := client.User.CreateUser(u)
	if err != nil {
		fmt.Printf("\nError:  %s\n", err.Error())
	}

	time.Sleep(time.Second * 2)
	}
}

