package main

import (
	"fmt"
	"math/rand"
	"runtime"
	"time"
)

// Function to select a Chef and Waiter for you
func selectWaiterAndChef() (string, string) {
	allChefs := []string{"Jack", "Bob", "Mark"}
	allWaiters := []string{"A", "B", "C"}

	randomC := rand.Intn(len(allChefs))
	randomW := rand.Intn(len(allWaiters))
	cpick := allChefs[randomC]
	wpick := allWaiters[randomW]

	return cpick, wpick
}

// Function where the waiter takes and brings your order
func processOrder(takes string, order int, waiter string, chef string) {
	if takes == "takes" {
		fmt.Printf("Waiter %s takes order %d to chef %s \n", waiter, order, chef)
	} else {
		fmt.Printf("Waiter %s brings order %d from chef %s \n", waiter, order, chef)
	}
}

// Function where the Chef cooks your meal
func chefCooks(gotOrder int, chef string) {
	fmt.Printf("Chef %s cooks order %d \n", chef, gotOrder)
}

// Main function
func main() {
	runtime.GOMAXPROCS(1)
	totalOrders := 5

	for i := 0; i < totalOrders; i++ {
		chef, waiter := selectWaiterAndChef()
		go processOrder("takes", i, waiter, chef)
		go chefCooks(i, chef)
		go processOrder("brings", i, waiter, chef)
	}

	<-time.After(time.Second * 5)
}
