package main

import (
	"bufio"
	"fmt"
	"log"
	"math/rand"
	"os"
	"strconv"
	"strings"
	"time"
)

func main() {
	rand.Seed(time.Now().Unix())
	generated := rand.Intn(100)
	fmt.Print("Guess any number below 100: ")
	reader := bufio.NewReader(os.Stdin)

out:
	for {

		number, err := reader.ReadString('\n')

		if err != nil {
			log.Fatal(err)
		}
		number = strings.TrimRight(number, "\r\n")
		guess, err := strconv.ParseInt(number, 10, 64)
		if err != nil {
			log.Fatal(err)
		}
		switch {
		case guess < int64(generated):
			fmt.Println("The number is too small")
		case guess > int64(generated):
			fmt.Println("The number is too big")
		case guess == int64(generated):
			fmt.Println("You are correct")
			break out
		}
	}
	fmt.Println("Bye")

}
