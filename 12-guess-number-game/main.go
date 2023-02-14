package main

import (
	"bufio"
	"fmt"
	"math/rand"
	"os"
	"time"
)

const prompt = "dan jangan ketikan angka kamu hanya tekan Enter ketika siap."

func main() {
	// seed the random number generator
	rand.Seed(time.Now().UnixNano())

	// rand generates a number between 0 and whatever is passed as a paramter
	// we add 2 to it because we want the number used to be at least 2, and no
	// greater than 10 (multiplying by 1 makes the game a bit silly)
	var firstNumber = rand.Intn(8) + 2
	var secondNumber = rand.Intn(8) + 2
	var subtraction = rand.Intn(8) + 2
	var answer = firstNumber*secondNumber - subtraction

	playTheGame(firstNumber, secondNumber, subtraction, answer)
}

func playTheGame(firstNumber, secondNumber, subtraction, answer int) {
	// create our reader variable, which reads input from standard in (the keyboard)
	reader := bufio.NewReader(os.Stdin)

	// display a welcome/instructions
	fmt.Println("Permainan tebak angka")
	fmt.Println("----------------------")
	fmt.Println("")

	fmt.Println("Pikirkan 1 angka antara 1 sampai 10", prompt)
	reader.ReadString('\n')

	// take them through the games
	fmt.Println("Kalikan nomor Anda dengan", firstNumber, prompt)
	reader.ReadString('\n')

	fmt.Println("Sekarang kalikan hasilnya dengan", secondNumber, prompt)
	reader.ReadString('\n')

	fmt.Println("Bagilah hasilnya dengan angka yang awalnya Anda pikirkan", prompt)
	reader.ReadString('\n')

	fmt.Println("Sekarang kurangi", subtraction, prompt)
	reader.ReadString('\n')

	// give them the answer
	fmt.Println("Jawabannya adalah", answer)
}
