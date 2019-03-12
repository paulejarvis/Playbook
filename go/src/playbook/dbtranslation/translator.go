// It is not possible to have a direct 1:1 dgraph-json-structure
// to proto conversion. It requires some work to get it right.
// The tradeoff is that we get type safety, and reasonable
// guarantees about the shape of process that enters our system.
package dbtranslation

import (
	"encoding/json"
	"github.com/pkg/errors"
	api "playbook/apis/procstore"
	"playbook/linker"
)

func NewTranslator() Translator {
	return &translator{}
}

type Translator interface {
	// TODO(hdh) rethink how this interface works for serializing mutations
	ProcessToDbJson(*api.Process) ([]byte, error)
}

type translator struct {}

func (t *translator) ProcessToDbJson(p *api.Process) ([]byte, error) {
	if err := linker.AnnotateProcess(p); err != nil {
		return nil, errors.Wrap(err, "failed to annotate process")
	}

	var out []map[string]interface{}
	var toDbJson = func(s *api.Step) error {
		iStep, err := NewStepFromProto(s)
		if err != nil {
			return errors.Wrap(err, "failed to get step from proto")
		}
		out = append(out, iStep.ToMaps()...)

		for _, state := range append(s.Causes, s.Requires...) {
			ic, err := NewStateFromProto(state)
			if err != nil {
				return errors.Wrap(err, "failed to get state from proto")
			}
			out = append(out, ic.ToMaps()...)

			io, err := NewObjectFromProto(state.Subject)
			if err != nil {
				return errors.Wrap(err, "failed to get state object from proto")
			}
			out = append(out, io.ToMaps()...)

			if so := state.GetStateObject(); so != nil {
				iso, err := NewObjectFromProto(so)
				if err != nil {
					return errors.Wrap(err, "failed to get state object from proto")
				}
				out = append(out, iso.ToMaps()...)
			}

			// TODO(hdh) include people once we figure out how to link them....

		}
		return nil
	}

	ip, err := NewProcessFromProto(p)
	if err != nil {
		return nil, errors.Wrap(err, "failed to get process from proto")
	}
	out = append(out, ip.ToMaps()...)
	if err:= linker.TraverseStepGraph(p.StartsWith, toDbJson); err != nil {
		return nil, errors.Wrap(err, "failed to traverse step graph")
	}

	return json.MarshalIndent(out, " ", " ")
}




