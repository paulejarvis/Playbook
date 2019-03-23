package linker

import (
	"crypto/sha256"
	"fmt"
	"github.com/google/uuid"
	"github.com/pkg/errors"
	ps "dl/apis/procstore"
	"strings"
)

// Attach Node types to all nodes in the tree.
func AnnotateProcess(p *ps.Process) error {
	if err := annotateNodes(p); err != nil {
		return errors.Wrap(err, "failed to annotate process")
	}

	var stepAnnotator = func(s *ps.Step) error {
		if err := annotateNodes(s); err != nil {
			return errors.Wrap(err, "failed to annotate step")
		}

		for i := range s.Causes {
			if err := annotateState(s.Causes[i]); err != nil {
				return errors.Wrap(err, "failed to annotate caused state")
			}
		}

		for i := range s.Requires {
			if err := annotateState(s.Requires[i]); err != nil {
				return errors.Wrap(err, "failed to annotate required state")
			}
		}

		return nil
	}

	if err := TraverseStepGraph(p.StartsWith, stepAnnotator); err != nil {
		return errors.Wrap(err, "failed to traverse step graph")
	}

	return nil
}


func annotateState(s *ps.State) error {
	if err := annotateNodes(s); err != nil {
		return errors.Wrap(err, "failed to annotate state")
	}

	if err := annotateNodes(s.Subject); err != nil {
		return errors.Wrap(err, "failed to annotate subject")
	}

	switch v := s.Object.(type) {
	case *ps.State_StateObject:
		if err := annotateNodes(v.StateObject); err != nil {
			return errors.Wrap(err, "failed to annotate state object")
		}
	}

	return nil
}

// walks the step graph
// TODO(hdh) this should be moved to a package of utils for traversing process graphs
func TraverseStepGraph(step *ps.Step, f func(s *ps.Step) error) error {
	if err := f(step); err != nil {
		return errors.Wrap(err, "failed to call f on step")
	}

	switch v := step.FollowedBy.(type) {
	case *ps.Step_FollowingSteps:
		for k, s := range v.FollowingSteps.Choice {
			step := *s.Step
			if err := TraverseStepGraph(&step, f); err != nil {
				return errors.Wrapf(err, "failed to call f on step at %s", k)
			}
		}

	case *ps.Step_FollowingStep:
		return TraverseStepGraph(v.FollowingStep, f)
	}

	return nil
}


// Attach Node types, and hash_ids to nodes
func annotateNodes(n interface{}) error {
	switch v := n.(type) {
	case *ps.Object:
		v.NodeKind = ps.NodeKind_NK_OBJECT
		// Handle the sub-type
		switch i := v.InstantiatedBy.(type) {
		case *ps.Object_Thing_:
			v.ObjectType = ps.Object_OT_THING
		case *ps.Object_Tool_:
			v.ObjectType = ps.Object_OT_TOOL
		case *ps.Object_Qualification_:
			v.ObjectType = ps.Object_OT_QUALIFICATION
		case *ps.Object_Machine_:
			v.ObjectType = ps.Object_OT_MACHINE
		case *ps.Object_Employee_:
			v.ObjectType = ps.Object_OT_EMPLOYEE
		default:
			return errors.New (fmt.Sprintf("%T is a non supported type", i))
		}

		hash, err := computeObjectHash(v)
		if err != nil {
			return errors.Wrap(err, "failed to compute hash for object")
		}
		v.HashId = hash
		v.Uid = getUid()

	case *ps.Step:
		v.NodeKind = ps.NodeKind_NK_STEP
		hash, err := computeCommonHash(v)
		if err != nil {
			return errors.Wrap(err, "failed to compute Node hash")
		}
		v.HashId = hash
		v.Uid = getUid()

	case *ps.Process:
		v.NodeKind = ps.NodeKind_NK_PROCESS
		hash, err := computeCommonHash(v)
		if err != nil {
			return errors.Wrap(err, "failed to compute Node hash")
		}
		v.HashId = hash
		v.Uid = getUid()

	case *ps.State:
		if err := annotateNodes(v.Subject); err != nil {
			return errors.Wrap(err, "failed to annotate state subject")
		}

		if v.GetStateObject() != nil {
			if err := annotateNodes(v.GetStateObject()); err != nil {
				return errors.Wrap(err, "failed to get annotate object")
			}
		}

	case *ps.Workflow:
		v.NodeKind = ps.NodeKind_NK_WORKFLOW
		hash, err := computeCommonHash(v)
		if err != nil {
			return errors.Wrap(err, "failed to compute Node hash")
		}
		v.HashId = hash
		v.Uid = getUid()

	case *ps.Person:
		v.NodeKind = ps.NodeKind_NK_PERSON
		hash, err := computeCommonHash(v)
		if err != nil {
			return errors.Wrap(err, "failed to compute Node hash")
		}
		v.HashId = hash
		v.Uid = getUid()

	default:
		return errors.New (fmt.Sprintf("%T is a non supported type", v))
	}

	return nil
}




func computeObjectHash(object *ps.Object) (string, error) {
	if object.NodeKind == ps.NodeKind_NK_UNSET {
		return "", errors.New("Node kind cannot be unset")
	}

	if object.ObjectType == ps.Object_OT_UNSET {
		return "", errors.New("object type cannot be unset")
	}

	if object.Name == "" {
		return "", errors.New("object name cannot be unset")
	}


	return computeHash(ps.NodeKind_name[int32(object.NodeKind)],
		ps.Object_ObjectType_name[int32(object.ObjectType)],
		object.Name)
}

type common interface{
	GetName() string
	GetNodeKind() ps.NodeKind
}

func computeCommonHash(e common) (string, error) {
	if e.GetNodeKind() == ps.NodeKind_NK_UNSET {
		return "", errors.New("Node kind cannot be unset")
	}

	if e.GetName() == "" {
		return "", errors.New("name cannot be unset")
	}

	return computeHash(ps.NodeKind_name[int32(e.GetNodeKind())],
		e.GetName())
}


func computeHash(strs ...string) (string, error) {
	h := sha256.New() // Maybe we don't need this, but might as well...
	h.Write([]byte(strings.Join(strs, "-")))
	return fmt.Sprintf("%x", h.Sum(nil)), nil
}


func getUid() string {
	return uuid.New().String()
}
