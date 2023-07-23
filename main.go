package main

import (
	"fmt"
	"os"
	"time"
)

// Application name
const APP_NAME string = "Billing and Invoice Generation"

// Product Struct
type Product struct {
	name     string
	price    float64
	quantity float64
}

// Product Struct function
func (pd Product) add_product() Product {
	// // Ask for product name
	fmt.Printf("Please enter product name\n")
	fmt.Scan(&pd.name)

	// // Ask for product Unit price
	fmt.Printf("Please enter Unit price of the product\n")
	fmt.Scan(&pd.price)

	// // Ask for Quantity to be purchased
	fmt.Printf("Please enter Quantity to be purchased\n")
	fmt.Scan(&pd.quantity)

	return pd
}

func main() {

	// Display Application name
	fmt.Println(APP_NAME, "\n\n")

	// Instantiate Product struct
	pd := Product{}

	productList := []Product{}

	for {
		pd = pd.add_product()

		productList = append(productList, pd)

		fmt.Println("Product added. Press q to exit and any other key to Continue!")

		var q string
		fmt.Scanln(&q)

		if q == "q" {
			break
		}

	}

	fmt.Println()
	var amountDue float64

	invoiceData := fmt.Sprintf("%v\n", APP_NAME)
	invoiceData += fmt.Sprintf("Date and Time: %v\n\n", time.Now())
	for i, product := range productList {

		invoiceData += fmt.Sprintf("%v)\t Product:\"%v\" \t Rate:\"%v\" \t Quantity: %v \n", i+1, product.name, product.price, product.quantity)

		amountDue += product.price * product.quantity

	}

	invoiceData += fmt.Sprintf("\nAmount Due: %v\n", amountDue)

	err := os.WriteFile("invoice.txt", []byte(invoiceData), 0644)

	if err != nil {
		fmt.Println(err)
	}

	fmt.Println("Invoice successfully generated.")
}
