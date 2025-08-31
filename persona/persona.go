package persona

import (
	"fmt"

	"trpc.group/trpc-go/trpc-a2a-go/protocol"
	"trpc.group/trpc-go/trpc-a2a-go/server"
)

// ToSystemPrompt renders the persona into a system prompt string for LLM calls.
func (p Persona) ToSystemPrompt() string {
	out := fmt.Sprintf("You are %s.\n\n", p.Role)

	if len(p.Goals) > 0 {
		out += formatList("Your goals:", p.Goals) + "\n"
	}

	if len(p.Backstory) > 0 {
		out += formatList("Your backstory:", p.Backstory) + "\n"
	}

	if len(p.Guidelines) > 0 {
		out += formatList("Guidelines:", p.Guidelines) + "\n"
	}
	return out
}

// helper to format a string slice as Markdown bullet points
func formatList(sectionTitle string, items []string) string {
	result := ""
	if sectionTitle != "" {
		result = sectionTitle
	}
	if len(items) == 0 {
		return "(none)"
	}
	for _, i := range items {
		result += "\n- " + i
	}
	// trim trailing "- "
	return result
}

// ToAgentCard converts a Persona into an A2A AgentCard.
// Skills should be passed in since they're task-specific and not persona-specific.
func (p Persona) ToAgentCard(
	name string,
	url string,
	version string,
	skills []server.AgentSkill,
) server.AgentCard {

	desc := p.ToSystemPrompt()

	return server.AgentCard{
		Name:        name,
		Description: desc,
		URL:         url,
		Version:     version,
		Capabilities: server.AgentCapabilities{
			Streaming: boolPtr(true),
		},
		DefaultInputModes:  []string{protocol.KindText},
		DefaultOutputModes: []string{protocol.KindText},
		Skills:             skills,
	}
}

/*
// ToAgentCardJSON returns an AgentCard with the description formatted as JSON.
func (p *Persona) ToAgentCardJSON(name string) (AgentCard, error) {
	jsonBytes, err := json.MarshalIndent(p, "", "  ")
	if err != nil {
		return AgentCard{}, err
	}
	return AgentCard{
		Name:        name,
		Description: string(jsonBytes),
	}, nil
}
*/

func boolPtr(b bool) *bool { return &b }
