package main

import (
	"fmt"

	"github.com/grokify/a2a-go/persona"
	"trpc.group/trpc-go/trpc-a2a-go/protocol"
	"trpc.group/trpc-go/trpc-a2a-go/server"
)

func main() {
	// Define persona
	persona := persona.Persona{
		Role: "Technology News Curator and Summarizer",
		Goals: []string{
			"Continuously monitor trusted tech sources",
			"Summarize the most relevant stories",
			"Prioritize articles by novelty and business impact",
		},
		Backstory: []string{
			"Veteran technology journalist with 10 years of experience",
			"Skilled at spotting patterns and distilling complex stories",
		},
		Guidelines: []string{
			"Write in clear, concise bullet points",
			"Highlight the most novel insights first",
		},
	}

	// Define skills
	skills := []server.AgentSkill{
		{
			ID:          "summarize_news",
			Name:        "Summarize News",
			Description: strPtr("Summarize multiple articles into concise bullet points."),
			InputModes:  []string{protocol.KindText},
			OutputModes: []string{protocol.KindText},
		},
		{
			ID:          "prioritize_news",
			Name:        "Prioritize News",
			Description: strPtr("Rank summarized stories by novelty and business impact."),
			InputModes:  []string{protocol.KindText},
			OutputModes: []string{protocol.KindText},
		},
	}

	// Convert to A2A AgentCard
	card := persona.ToAgentCard(
		"Tech News Curator Agent",
		"http://localhost:8080/",
		"0.1.0",
		skills,
	)

	// Print for demo
	fmt.Println("=== System Prompt ===")
	fmt.Println(persona.ToSystemPrompt())
	fmt.Println("\n=== AgentCard ===")
	fmt.Printf("%+v\n", card)
}

func strPtr(s string) *string { return &s }
