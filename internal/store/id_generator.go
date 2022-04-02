package store

import (
	"math/rand"
)

var allowedChars = []rune("abcdefghijklmnopqrstuvwxyz1234567890")

func generateIDs(numOfIDs int, length int) []string {
	reservedIDs := make([]string, 0)
	reservedIDsTracker := make(map[string]bool)

	generateID := func() string {
		id := make([]rune, length)
		for i := range id {
			id[i] = allowedChars[rand.Intn(len(allowedChars))]
		}

		return string(id)
	}

	for i := 0; i < numOfIDs; i++ {
		id := generateID()
		if _, exists := reservedIDsTracker[id]; exists {
			continue
		}

		reservedIDsTracker[id] = true
		reservedIDs = append(reservedIDs, id)
	}

	return reservedIDs
}
