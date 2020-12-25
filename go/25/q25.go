package main

import (
	"fmt"
	"time"
)

func main() {
	start := time.Now()

	cardPubKey := 13233401
	doorPubKey := 6552760
	//cardPubKey := 5764801
	//doorPubKey := 17807724
	q25part1(cardPubKey, doorPubKey)

	elapsed := time.Since(start)

	fmt.Printf("Main took %s", elapsed)
}

func doHandshake(value int, subjectNumber int) int {
	value = value * subjectNumber
	value = value % 20201227
	return value
}

func q25part1(cardPubKey int, doorPubKey int) {
	doorNumber := 1
	cardNumber := 1
	doorLoops := -1
	cardLoops := -1
	i := 0
	for doorLoops < 0 {
		i++
		if doorLoops == -1 {
			doorNumber = doHandshake(doorNumber, 7)
			if doorNumber == doorPubKey {
				doorLoops = i
			}
		}

	}
	i = 0
	for cardLoops < 0 {
		i++
		if cardLoops == -1 {
			cardNumber = doHandshake(cardNumber, 7)
			if cardNumber == cardPubKey {
				cardLoops = i
			}
		}

	}
	//fmt.Printf("Door Loops: %d, Card Loops: %d\n", doorLoops, cardLoops)

	doorEncryption := 1
	for i := 0; i < cardLoops; i++ {
		doorEncryption = doHandshake(doorEncryption, doorPubKey)
	}
	cardEncryption := 1
	for i := 0; i < doorLoops; i++ {
		cardEncryption = doHandshake(cardEncryption, cardPubKey)
	}

	fmt.Printf("Door Encryption: %d, card Encryption %d\n", doorEncryption, cardEncryption)
}
