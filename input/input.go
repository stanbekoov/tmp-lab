package input

import (
	"bufio"
	"fmt"
	"log"
	"os"
)

func Read() {
	file, err := os.OpenFile("usernames.txt", os.O_RDWR|os.O_CREATE, 0777)
	if err != nil {
		log.Fatal(err.Error())
	}
	defer file.Close()

	scanner := bufio.NewScanner(os.Stdin)

	fmt.Println("Type your name here:")
	scanner.Scan()
	input := scanner.Text()

	fileScanner := bufio.NewScanner(file)

	for fileScanner.Scan() {
		if input == fileScanner.Text() {
			fmt.Println("User with this name already exists")
			return
		}
	}

	_, err = file.WriteString(input + "\n")
	if err != nil {
		log.Fatal(err.Error())
	}
	fmt.Println("SUCCESS")
}
