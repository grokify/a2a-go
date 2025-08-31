package persona

// AgentCard represents the A2A protocol AgentCard.
// This is a simplified placeholder; adapt to match trpc-a2a-go’s AgentCard struct.
type AgentCard struct {
	Name        string `json:"name"`
	Description string `json:"description"`
}

// Persona represents a higher-level abstraction developers can use
// when defining LLM-based agents with Role, Goals, Backstory, and Guidelines.
// Role: a single string identity
// Goals: list of mission objectives
// Backstory: narrative background
// Guidelines: style/behavior constraints
type Persona struct {
	Role       string
	Goals      []string
	Backstory  []string
	Guidelines []string
}

// New creates a new Persona with required Role and optional attributes.
func New(role string, opts ...Option) *Persona {
	p := &Persona{Role: role}
	for _, opt := range opts {
		opt(p)
	}
	return p
}

// Option pattern for extensibility
type Option func(*Persona)

func WithGoals(goals ...string) Option {
	return func(p *Persona) { p.Goals = goals }
}

func WithBackstory(backstory ...string) Option {
	return func(p *Persona) { p.Backstory = backstory }
}

func WithGuidelines(guidelines ...string) Option {
	return func(p *Persona) { p.Guidelines = guidelines }
}
