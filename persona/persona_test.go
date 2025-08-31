package persona

import (
	"testing"
)

func TestPersonalToSystemPrompt(t *testing.T) {
	var personaTests = []struct {
		persona *Persona
		want    string
	}{
		{New(
			"Tech News Curator",
			WithGoals("Summarize technology news", "Prioritize AI/Cloud topics"),
			WithBackstory("Reads TechCrunch, Wired, and Hacker News daily"),
			WithGuidelines("Keep summaries under 200 words", "Avoid hype"),
		), `You are Tech News Curator.

Your goals:
- Summarize technology news
- Prioritize AI/Cloud topics
Your backstory:
- Reads TechCrunch, Wired, and Hacker News daily
Guidelines:
- Keep summaries under 200 words
- Avoid hype
`},
	}

	for _, tt := range personaTests {
		sysprompt := tt.persona.ToSystemPrompt()
		if sysprompt != tt.want {
			t.Errorf("Personal.ToSystemPrompt() Mismatch: want [%v], got [%v]",
				tt.want, sysprompt)
		}
	}
}
