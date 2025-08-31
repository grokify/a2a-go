# A2A-Go

[![Build Status][build-status-svg]][build-status-url]
[![Lint Status][lint-status-svg]][lint-status-url]
[![Go Report Card][goreport-svg]][goreport-url]
[![Docs][docs-godoc-svg]][docs-godoc-url]
[![License][license-svg]][license-url]

Go helpers for the Agent2Agent (A2A) Protocol, providing utilities to create AI agents with well-defined personas that can integrate with A2A-compliant systems.

## Overview

This package provides a higher-level abstraction for defining LLM-based agents using the **Persona** pattern. Instead of manually crafting system prompts, you can define agents using structured components like Role, Goals, Backstory, and Guidelines, then automatically convert them to A2A AgentCards or system prompts.

## Features

🎭 **Persona-based Agent Definition**: Define agents using Role, Goals, Backstory, and Guidelines
⚡ **System Prompt Generation**: Automatically convert personas to formatted system prompts
🔗 **A2A AgentCard Integration**: Convert personas to Agent2Agent protocol AgentCards
🏗️ **Flexible Construction**: Use functional options pattern for clean agent creation
🛡️ **Type Safety**: Fully typed Go structs with comprehensive test coverage

## Installation

```bash
go get github.com/grokify/a2a-go
```

## Quick Start

```go
package main

import (
    "fmt"
    "github.com/grokify/a2a-go/persona"
    "trpc.group/trpc-go/trpc-a2a-go/protocol"
    "trpc.group/trpc-go/trpc-a2a-go/server"
)

func main() {
    // Create a persona using the builder pattern
    p := persona.New(
        "Technology News Curator and Summarizer",
        persona.WithGoals(
            "Continuously monitor trusted tech sources",
            "Summarize the most relevant stories",
            "Prioritize articles by novelty and business impact",
        ),
        persona.WithBackstory(
            "Veteran technology journalist with 10 years of experience",
            "Skilled at spotting patterns and distilling complex stories",
        ),
        persona.WithGuidelines(
            "Write in clear, concise bullet points",
            "Highlight the most novel insights first",
        ),
    )

    // Generate system prompt
    fmt.Println("System Prompt:")
    fmt.Println(p.ToSystemPrompt())

    // Convert to A2A AgentCard
    skills := []server.AgentSkill{
        {
            ID:          "summarize_news",
            Name:        "Summarize News",
            Description: &description,
            InputModes:  []string{protocol.KindText},
            OutputModes: []string{protocol.KindText},
        },
    }

    card := p.ToAgentCard(
        "Tech News Curator Agent",
        "http://localhost:8080/",
        "0.1.0",
        skills,
    )
}
```

## API Reference

### Persona Struct

```go
type Persona struct {
    Role       string    // Primary identity/role of the agent
    Goals      []string  // Mission objectives and targets
    Backstory  []string  // Narrative background and context
    Guidelines []string  // Behavioral constraints and style guides
}
```

### Constructor

```go
func New(role string, opts ...Option) *Persona
```

Creates a new Persona with the specified role and optional attributes.

### Options

```go
func WithGoals(goals ...string) Option
func WithBackstory(backstory ...string) Option  
func WithGuidelines(guidelines ...string) Option
```

Functional options for building personas with clean, readable code.

### Methods

#### `ToSystemPrompt() string`

Converts the persona into a formatted system prompt string suitable for LLM calls.

**Example Output:**
```
You are Technology News Curator and Summarizer.

Your goals:
- Continuously monitor trusted tech sources
- Summarize the most relevant stories
- Prioritize articles by novelty and business impact

Your backstory:  
- Veteran technology journalist with 10 years of experience
- Skilled at spotting patterns and distilling complex stories

Guidelines:
- Write in clear, concise bullet points
- Highlight the most novel insights first
```

#### `ToAgentCard(name, url, version string, skills []server.AgentSkill) server.AgentCard`

Converts the persona into an A2A protocol AgentCard with the specified metadata and skills.

## Package Structure

```
.
├── persona/
│   ├── card.go          # Persona struct and constructor
│   ├── persona.go       # Core methods and formatting
│   └── persona_test.go  # Test suite
├── examples/
│   └── persona/
│       └── main.go      # Usage examples
├── go.mod
└── go.sum
```

## Dependencies

- `trpc.group/trpc-go/trpc-a2a-go`: Core A2A protocol implementation
- Go 1.24.5+

## Examples

See the [examples/persona](examples/persona/) directory for complete working examples.

## Testing

Run tests with:

```bash
go test ./...
```

## Contributing

Contributions welcome! Please ensure tests pass and follow Go conventions.

 [build-status-svg]: https://github.com/grokify/a2a-go/actions/workflows/ci.yaml/badge.svg?branch=main
 [build-status-url]: https://github.com/grokify/a2a-go/actions/workflows/ci.yaml
 [lint-status-svg]: https://github.com/grokify/a2a-go/actions/workflows/lint.yaml/badge.svg?branch=main
 [lint-status-url]: https://github.com/grokify/a2a-go/actions/workflows/lint.yaml
 [goreport-svg]: https://goreportcard.com/badge/github.com/grokify/a2a-go
 [goreport-url]: https://goreportcard.com/report/github.com/grokify/a2a-go
 [docs-godoc-svg]: https://pkg.go.dev/badge/github.com/grokify/a2a-go
 [docs-godoc-url]: https://pkg.go.dev/github.com/grokify/a2a-go
 [license-svg]: https://img.shields.io/badge/license-MIT-blue.svg
 [license-url]: https://github.com/grokify/a2a-go/blob/master/LICENSE
 [used-by-svg]: https://sourcegraph.com/github.com/grokify/a2a-go/-/badge.svg
 [used-by-url]: https://sourcegraph.com/github.com/grokify/a2a-go?badge