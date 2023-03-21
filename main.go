package main

import (
	"math/rand"
	"time"

	"github.com/micmonay/keybd_event"
)

func main() {
	// Seed the random number generator
	rand.Seed(time.Now().UnixNano())

	// Define the keys
	keys := []int{keybd_event.VK_SPACE, keybd_event.VK_ENTER, keybd_event.VK_SP1}

	for {
		// Create a new keybd_event instance
		kb, err := keybd_event.NewKeyBonding()
		if err != nil {
			panic(err)
		}

		// Choose a random key from the keys slice
		randomKey := keys[rand.Intn(len(keys))]

		// Set the key to press
		kb.SetKeys(randomKey)

		// Press and release the key
		err = kb.Launching()
		if err != nil {
			panic(err)
		}

		// Sleep for 1 minute
		time.Sleep(1 * time.Minute)
	}
}
