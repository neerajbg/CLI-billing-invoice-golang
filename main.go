package main

import "fmt"

const APP_NAME string = "Billing and Invoice Generation"

type Product struct {
	name     string
	price    float64
	quantity float64
}

func (pd Product) add_product() Product {

	return pd
}

func main() {

	// // Step1 - Ask for product name
	var product string
	fmt.Printf("Please enter product name\n")
	fmt.Scan(&product)

	// // Step2 - Ask for product Unit price
	var price float64
	fmt.Printf("Please enter Unit price of the product\n")
	fmt.Scan(&price)

	// // Step3 - Ask for  Quantity to be purchased
	var quantity float64
	fmt.Printf("Please enter Quantity to be purchased\n")
	fmt.Scan(&quantity)

	fmt.Printf("Amount for the Product \"%v\" with \"%v\" Quantity is \n", product, quantity)

	fmt.Printf("Amount Due: %v\n", price*quantity)

}
