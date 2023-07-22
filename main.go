package main

import "fmt"

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

	pd = pd.add_product()

	fmt.Printf("Amount for the Product \"%v\" with \"%v\" Quantity is \n", pd.name, pd.quantity)

	fmt.Printf("Amount Due: %v\n", pd.price*pd.quantity)
}
