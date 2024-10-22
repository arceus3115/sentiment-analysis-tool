package main

import (
	"bufio"
	"fmt"
	"log"
	"os"

	"github.com/cdipaolo/sentiment"
)

func main() {
	// Restore the pre-trained sentiment analysis model
	model, err := sentiment.Restore()
	if err != nil {
		log.Fatalf("Could not restore model: %v\n", err)
	}

	fmt.Println("Welcome to the Mental Health Sentiment Analysis tool!")
	fmt.Println("Please enter your thoughts or feelings (type 'exit' to quit):")

	scanner := bufio.NewScanner(os.Stdin)

	for {
		fmt.Print("> ")
		scanner.Scan()
		input := scanner.Text()

		// Exit if user types 'exit'
		if input == "exit" {
			fmt.Println("Exiting the program. Take care!")
			break
		}

		// Perform sentiment analysis
		analysis := model.SentimentAnalysis(input, sentiment.English)

		// Display the results
		fmt.Printf("Sentiment Score: %d\n", analysis.Score)
		provideFeedback(analysis.Score)
	}
}

// provideFeedback gives feedback based on the sentiment score
func provideFeedback(score uint8) {
	switch score {
	case 0:
		fmt.Println("It seems like you're feeling negative. It's okay to seek help. Please consider talking to someone.")
	case 1:
		fmt.Println("You're expressing some negative sentiments. Remember, it's okay to share your feelings.")
	case 2:
		fmt.Println("You're neutral. It's a good opportunity to reflect on your feelings.")
	case 3:
		fmt.Println("You're feeling somewhat positive! That's great! Keep focusing on the good things.")
	case 4:
		fmt.Println("You're feeling very positive! Wonderful to hear! Keep up the positivity.")
	default:
		fmt.Println("Sentiment analysis score not recognized.")
	}
}
