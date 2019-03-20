package ontology

type uid string

type NodeType string
const (
	NTProcess  NodeType = "process"

)

type Process struct {
	Type NodeType
	Id uid

	Label string // Processes display name
	Description string // Process description

	Owns []StepLike // Steps owned by this process
	Start StepLike // first step in this process

	CanBeDoneBy []*Role
}

type Effect struct {
	Type NodeType
	Id uid

	RequiredBy []*Step
	CausedBy []*Step
	HasSubject *Thing

	Predicate string // How does it relate to its object (state or thing)
	// Either one or the other is set, never both.
	State *State
	Thing *Thing
}

type Media struct {
	Type NodeType
	Id uid

	Label string
	Content string // URL of content

	Next *Media
	Previous *Media
}

type Person struct {
	Type NodeType
	Id uid

	Label string
	Description string

	CanFill []*Role
}

type Role struct {
	Type NodeType
	Id uid

	Label string
	Description string

	CanBeFilledBy []*Person
}

type State struct {
	Type NodeType
	Id uid
}

type Thing struct {
	Type NodeType
	Id uid

	Label string
	Description string
	Properties map[string]interface{}

	SubjectOf []*Effect
}

type Workflow struct {
	Type NodeType
	Id uid

	Label string
	Description string

	Owns []*Process
}


type StepLike interface {
	Next() []struct{
		Choice string // Can be blank
		Next StepLike }
	Previous() StepLike
}

type Step struct {
	Type NodeType
	Id uid

	Label string // Step display name
	Description string // Step description

	Next []StepLike // Following Steps Or Decision
	OwnedBy *Process // Process that this is part of

	Requires []*Effect
	Causes []*Effect

	HasMedia *Media

}

type Decision struct {
	Type NodeType
	Id uid

	Label string // Decision display name
	Description string // Decision description

	Next map[string]StepLike // Following Steps Or Decision
	OwnedBy *Process // Process that this is part of
}
