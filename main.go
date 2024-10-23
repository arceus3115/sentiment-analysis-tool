package main

import (
	"bufio"
	"fmt"
	"os"

	"github.com/jonreiter/govader"
)

type QuestionCategory struct {
	Question string
	Category string
}

func main() {
	// Initialize VADER sentiment analyzer
	analyzer := govader.NewSentimentIntensityAnalyzer()

	// Define a list of questions with their corresponding categories
	questionCategories := []QuestionCategory{
		{"What made me smile today?", "Daily Reflection"},
		{"What challenges did I face today, and how did I handle them?", "Daily Reflection"},
		{"What emotions did I experience today?", "Daily Reflection"},
		{"What is one thing I could have done differently today?", "Daily Reflection"},
		{"What am I grateful for today?", "Daily Reflection"},
		{"What personal qualities do I admire in others?", "Personal Growth"},
		{"What fears or limiting beliefs are holding me back?", "Personal Growth"},
		{"How have I changed in the past year?", "Personal Growth"},
		{"What habits do I want to develop or break?", "Personal Growth"},
		{"What does self-love look like for me?", "Personal Growth"},
		{"What are my short-term and long-term goals?", "Goal-Oriented Reflection"},
		{"What steps did I take today toward achieving my goals?", "Goal-Oriented Reflection"},
		{"What motivates me to achieve my goals?", "Goal-Oriented Reflection"},
		{"How do I define success for myself?", "Goal-Oriented Reflection"},
		{"What obstacles am I facing in reaching my goals?", "Goal-Oriented Reflection"},
		{"What moments today made me feel present and engaged?", "Mindfulness and Presence"},
		{"How can I bring more mindfulness into my daily routine?", "Mindfulness and Presence"},
		{"What do I need to let go of to feel more at peace?", "Mindfulness and Presence"},
		{"How can I simplify my life to reduce stress?", "Mindfulness and Presence"},
		{"What does my ideal day look like?", "Mindfulness and Presence"},
	}

	// Open a scanner for user input
	scanner := bufio.NewScanner(os.Stdin)

	// Create a map to hold cumulative sentiment scores for each category
	categoryScores := make(map[string][]float64)

	// Collect answers and perform sentiment analysis
	for _, qc := range questionCategories {
		fmt.Println(qc.Question)
		fmt.Print("> ")
		scanner.Scan()
		input := scanner.Text()

		// Perform sentiment analysis
		sentiment := analyzer.PolarityScores(input)

		// Display the sentiment results
		fmt.Printf("Compound score: %f\n", sentiment.Compound)
		fmt.Printf("Positive score: %f\n", sentiment.Positive)
		fmt.Printf("Neutral score: %f\n", sentiment.Neutral)
		fmt.Printf("Negative score: %f\n\n", sentiment.Negative)

		// Store the scores in the categoryScores map
		categoryScores[qc.Category] = append(categoryScores[qc.Category], sentiment.Compound)
	}

	// Summarize sentiments for each category after all questions have been answered
	for category, scores := range categoryScores {
		summarizeSentiments(scores, category)
	}
}

// Function to summarize sentiment analysis based on categories
func summarizeSentiments(scores []float64, category string) {
	total := len(scores)
	positiveCount := 0
	negativeCount := 0
	neutralCount := 0

	for _, score := range scores {
		if score >= 0.05 {
			positiveCount++
		} else if score <= -0.05 {
			negativeCount++
		} else {
			neutralCount++
		}
	}

	fmt.Printf("Summary for category '%s':\n", category)
	fmt.Printf("Total responses: %d\n", total)
	fmt.Printf("Positive responses: %d\n", positiveCount)
	fmt.Printf("Negative responses: %d\n", negativeCount)
	fmt.Printf("Neutral responses: %d\n", neutralCount)

	if positiveCount > negativeCount {
		fmt.Println("Overall sentiment: Positive")
	} else if negativeCount > positiveCount {
		fmt.Println("Overall sentiment: Negative")
	} else {
		fmt.Println("Overall sentiment: Neutral")
	}
	fmt.Println()
}
