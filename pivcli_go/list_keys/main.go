package main

import (
	"fmt"
	"log"
	"github.com/go-piv/piv-go/piv"
)


func listKeys() {
	cardNames, err := piv.Cards()
	if err != nil {
		log.Fatalf("Failed to list cards: %v", err)
	}

	if len(cardNames) == 0 {
		fmt.Println("No PIV cards found.")
		return
	}

	fmt.Printf("Found %d PIV card(s):\n", len(cardNames))
	for _, name := range cardNames {
		fmt.Printf("- %s\n", name)
	}
}


func openFirstKey() {
	// Find all cards
	cardNames, err := piv.Cards()
	if err != nil {
		log.Fatalf("Failed to list cards: %v", err)
	}

	if len(cardNames) == 0 {
		log.Fatalln("No PIV cards found. Please insert a YubiKey or SoloKey.")
	}

	// Open the first card found
	fmt.Printf("Attempting to open card: %s\n", cardNames[0])

	tx, err := piv.Open(cardNames[0])
	if err != nil {
		log.Fatalf("Failed to open card: %v", err)
	}
	// This ensures the card connection is closed when the function finishes
	defer tx.Close()

	fmt.Println("Successfully opened the card!")

	// Now that we have the 'tx' object, we can interact with the key.
}

func main() {
	//openFirstKey()

	listKeys()
}