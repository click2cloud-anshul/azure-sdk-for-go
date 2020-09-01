package main

import (
	"context"
	"fmt"
	"os"

	"github.com/Azure/azure-sdk-for-go/services/resources/mgmt/2018-05-01/resources"
	"github.com/Azure/go-autorest/autorest/azure/auth"
)

func main() {

	authorizer, err := auth.NewAuthorizerFromEnvironment()
	if err != nil {
		panic(err)
	}
	subscriptionID := os.Getenv("AZURE_SUBSCRIPTION_ID")

	// list all resource groups within the given subscription
	rgClient := resources.NewGroupsClient(subscriptionID)
	rgClient.Authorizer = authorizer
	groups, err := rgClient.List(context.Background(), "", nil)
	if err != nil {
		panic(err)
	}

	for _, group := range groups.Values() {
		fmt.Println(*group.Name)
	}

}
