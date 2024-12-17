package main

import (
	"errors"
	"fmt"
	"os"
	"strings"
)

func main() {
	// Check if an argument is provided

	company, err := extrago ctCompanyName()
	if err != nil {
		fmt.Println(err)
	}

	// Retrieve the argument
	link, err := getCompanyLink(company)
	if err != nil {
		fmt.Println(err)
	}
	// print link & company
	fmt.Printf("%s: %s\n", company, link)
}

func getCompanyLink(company string) (string, error) {
	switch company {
	case "amazon":
		{
			return "www.amazon.com", nil
		}
	case "oracle":
		{
			return "www.oracle.com", nil
		}
	case "ebay":
		{
			return "www.ebay.com", nil
		}
	case "bikeinn":
		{
			return "www.bikeinn.com", nil
		}
	case "trek":
		{
			return "www.trek.com", nil
		}
	case "iherb":
		{
			return "www.iherb.com", nil
		}
	case "krispykreme":
		{
			return "www.krispykreme.com", nil
		}
	default:
		{
			return "", errors.New("invalid company name")
		}
	}
}

func extractCompanyName() (string, error) {
	if len(os.Args) < 2 {
		return "", errors.New("Usage: go run main.go <your_name>")
	}
	company := os.Args[1]
	company = strings.ToLower(company)
	return company, nil
}
