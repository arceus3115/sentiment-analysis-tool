package main

import (
	"testing"

	"github.com/jonreiter/govader"
)

type TestQuestionCategory struct {
	Question string
	Category string
	Answer   string
}

func TestSentimentAnalysis(t *testing.T) {
	// Initialize VADER sentiment analyzer
	analyzer := govader.NewSentimentIntensityAnalyzer()

	// Define a list of test questions and their corresponding answers
	testQuestionCategories := []TestQuestionCategory{
		{"What made me smile today?", "Daily Reflection", "Seeing my friends made me smile."},
		{"What challenges did I face today, and how did I handle them?", "Daily Reflection", "I struggled with a difficult project at work."},
		{"What emotions did I experience today?", "Daily Reflection", "I felt happy and content."},
		{"What is one thing I could have done differently today?", "Daily Reflection", "I could have been more organized."},
		{"What am I grateful for today?", "Daily Reflection", "I'm grateful for my supportive family."},
		{"What personal qualities do I admire in others?", "Personal Growth", "I admire kindness and resilience in others."},
		{"What fears or limiting beliefs are holding me back?", "Personal Growth", "My fear of failure holds me back."},
		{"How have I changed in the past year?", "Personal Growth", "I've become more confident in my abilities."},
		{"What habits do I want to develop or break?", "Personal Growth", "I want to develop better time management skills."},
		{"What does self-love look like for me?", "Personal Growth", "Self-love means accepting myself as I am."},
		{"What are my short-term and long-term goals?", "Goal-Oriented Reflection", "My goal is to improve my fitness this year."},
		{"What steps did I take today toward achieving my goals?", "Goal-Oriented Reflection", "I took a walk today, which helped my progress."},
		{"What motivates me to achieve my goals?", "Goal-Oriented Reflection", "I'm motivated by the desire to help others."},
		{"How do I define success for myself?", "Goal-Oriented Reflection", "I define success as being happy with my choices."},
		{"What obstacles am I facing in reaching my goals?", "Goal-Oriented Reflection", "My main obstacle is procrastination."},
		{"What moments today made me feel present and engaged?", "Mindfulness and Presence", "I felt present during my morning meditation."},
		{"How can I bring more mindfulness into my daily routine?", "Mindfulness and Presence", "I can bring more mindfulness by focusing on my breath."},
		{"What do I need to let go of to feel more at peace?", "Mindfulness and Presence", "I need to let go of my need for perfection."},
		{"How can I simplify my life to reduce stress?", "Mindfulness and Presence", "I can simplify my life by decluttering my space."},
		{"What does my ideal day look like?", "Mindfulness and Presence", "My ideal day involves a mix of work and relaxation."},
	}

	// Create a map to hold cumulative sentiment scores for each category
	categoryScores := make(map[string][]float64)

	for _, qc := range testQuestionCategories {
		// Perform sentiment analysis on preset answers
		sentiment := analyzer.PolarityScores(qc.Answer)

		// Store the scores in the categoryScores map
		categoryScores[qc.Category] = append(categoryScores[qc.Category], sentiment.Compound)
	}

	// Summarize sentiments for each category
	for category, scores := range categoryScores {
		summarizeSentiments(scores, category) // Call the original function from main.go
	}
}
