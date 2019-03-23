// Package ontology provides a way of representing process/ nodes
// as go types. This provides the ability to write methods on top
// of Resource types/ collections of nodes. Node types in ontology are able
// to translate themselves into proto-api types and be correctly sent
// to the front end.
package ontology

import (
	"fmt"
)

// TODO(hdh) write transformation logic so that a process can be "written" to db (updates and creates)
// TODO(hdh) write transformation logic so that a process can be read from the db
// TODO(hdh) write transformation logic so that a process can be sent along via proto

// Any field that is designated with a predicate (pred:foo)
// Must be an addressable Resource, or a collection of addressableNodes
// Any field that is designated with a prop must the correct type for that predicate
// (haven't done this yet)
type Resource interface {
	GetIri() string
}

func iriFromId(id string) string {
	return fmt.Sprintf("%s:id/%s", dl, id)
}

type Process struct {
	Id          string  `class:"process"`
	Label       string  `prop:"label"`
	Description string  `prop:"description"`
	Owns        []*Step `pred:"owns"`
	Start       *Step   `pred:"next"`
	CanBeDoneBy []*Role `pred:"canBeDoneBy"`
}

func (v *Process) GetIri() string {
	return iriFromId(v.Id)
}

type Effect struct {
	Id         string                `class:"effect"`
	RequiredBy []*Step               `pred:"requiredBy"`
	CausedBy   []*Step               `pred:"causedBy"`
	HasSubject *Thing                `pred:"hasSubject"`
	StateThing map[string]stateThing `predMap:"state"`
}

func (v *Effect) GetIri() string {
	return iriFromId(v.Id)
}

type Media struct {
	Id       string `class:"media"`
	Label    string `prop:"label"`
	Url      string `prop:"url"`
	Next     *Media `pred:"next"`
	Previous *Media `pred:"previous"`
}

func (v *Media) GetIri() string {
	return iriFromId(v.Id)
}

type Person struct {
	Id          string  `class:"person"`
	Label       string  `prop:"label"`
	Description string  `prop:"description"`
	CanFill     []*Role `pred:"canFill"`
}

func (v *Person) GetIri() string {
	return iriFromId(v.Id)
}

type Role struct {
	Id            string    `class:"role"`
	Label         string    `prop:"label"`
	Description   string    `prop:"description"`
	CanBeFilledBy []*Person `pred:"canBeFilledBy"`
}

func (v *Role) GetIri() string {
	return iriFromId(v.Id)
}

type stateThing interface {
	isStateThing()
}

type State struct {
	Id    string `class:"state"`
	Label string `prop:"label"`
}

func (v *State) GetIri() string {
	return iriFromId(v.Id)
}

func (v *State) isStateThing() {}

type Thing struct {
	Id          string                 `class:"thing"`
	Label       string                 `prop:"label"`
	Description string                 `prop:"description"`
	SubjectOf   []*Effect              `pred:"subjectOf"`
	Properties  map[string]interface{} `propMap:""`
}

func (v *Thing) GetIri() string {
	return iriFromId(v.Id)
}

func (v *Thing) isStateThing() {}

type Workflow struct {
	Type        class
	Id          string     `class:"workflow"`
	Label       string     `prop:"label"`
	Description string     `prop:"description"`
	Owns        []*Process `pred:"owns"`
}

func (v *Workflow) GetIri() string {
	return iriFromId(v.Id)
}

type StepLike interface {
	Next() []struct {
		Choice string // Can be blank
		Next   StepLike
	}
	Previous() StepLike
}

type Step struct {
	Id          string     `class:"step"`
	Label       string     `prop:"label"`
	Description string     `prop:"description"`
	Next        []StepLike `pred:"next"`
	OwnedBy     *Process   `pred:"ownedBy"`
	Requires    []*Effect  `pred:"requires"`
	Causes      []*Effect  `pred:"causes"`
	HasMedia    *Media     `pred:"hasMedia"`
}

func (v *Step) GetIri() string {
	return iriFromId(v.Id)
}

type Decision struct {
	Id          string              `class:"step"`
	Label       string              `prop:"label"`
	Description string              `prop:"description"`
	OwnedBy     *Process            `pred:"ownedBy"`
	Next        map[string]StepLike `predMap:"next"`
}

func (v *Decision) GetIri() string {
	return iriFromId(v.Id)
}
