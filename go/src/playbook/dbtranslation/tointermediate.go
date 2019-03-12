package dbtranslation

import (
	"fmt"
	"github.com/pkg/errors"
	ps "playbook/apis/procstore"
)

func NewProcessFromProto(pType *ps.Process) (*process, error) {
	return &process{
		nref:nref{
			Uid: pType.Uid,
		},
		NodeKind:    ps.NodeKind_name[int32(pType.NodeKind)],
		HashId:      pType.HashId,
		Name:        pType.Name,
		Description: pType.Description,

		StartsWith: nref{Uid:pType.StartsWith.HashId},

	}, nil
}


func NewStepFromProto(pType *ps.Step) (*step, error) {
	causes := make([]nref, len(pType.Causes))
	for i := range causes {
		causes[i] = nref{Uid: pType.Causes[i].Subject.Uid}
	}

	requires := make([]nref, len(pType.Requires))
	for i := range requires {
		requires[i] = nref{Uid: pType.Requires[i].Subject.Uid}
	}

	followedBy := make(map[string]nref)

	switch fb := pType.FollowedBy.(type) {
	case *ps.Step_FollowingStep:
		followedBy["followingStep"] = nref{Uid: fb.FollowingStep.HashId}

	case *ps.Step_FollowingSteps:
		choices := fb.FollowingSteps.Choice
		for i := range choices {
			choice := choices[i].Choice
			step := choices[i].Step
			if _, isSet := followedBy[choice]; isSet {
				return nil,
					errors.New(fmt.Sprintf("cannot have duplicate choices %T, used more than once", choice))
			}

			followedBy[choice] = nref{Uid: step.HashId}
		}
	}

	var precedingSteps []nref
	switch fb := pType.PrecededBy.(type) {
	case *ps.Step_PrecedingStep:
		precedingSteps = append(precedingSteps, nref{Uid: fb.PrecedingStep.HashId})

	case *ps.Step_PrecedingSteps:
		steps := fb.PrecedingSteps.Steps
		for i := range steps {
			precedingSteps = append(precedingSteps, nref{Uid: steps[i].HashId})
		}
	}

	return &step{
		nref: nref{
			Uid: pType.HashId,
		},
		NodeKind:    ps.NodeKind_name[int32(pType.NodeKind)],
		HashId:      pType.HashId,
		Name:        pType.Name,
		Description: pType.Description,

		FollowedBy: followedBy,
		PrecededBy: precedingSteps,

		Causes:   causes,
		Requires: requires,
	}, nil
}


func NewObjectFromProto(pType *ps.Object) (*object, error) {
	out := &object{
		nref:nref{
			Uid: pType.Uid,
		},
		NodeKind:    ps.NodeKind_name[int32(pType.NodeKind)],
		ObjectType:  ps.Object_ObjectType_name[int32(pType.ObjectType)],
		HashId:      pType.HashId,
		Name:        pType.Name,
		Description: pType.Description,
	}

	if pType.ExistsIn != nil {
	  out.ExistsIn = nref{Uid: pType.ExistsIn.Uid}
	}

	switch state := pType.State.(type) {
	case *ps.Object_IsIn:
		out.IsIn = nref{Uid: state.IsIn.Uid}
	case *ps.Object_IsNotIn:
		out.IsNotIn = nref{Uid: state.IsNotIn.Uid}
	}

	switch instance := pType.InstantiatedBy.(type) {
	case *ps.Object_Thing_:
		out.Thing_Props = instance.Thing.Props

	case *ps.Object_Tool_:
		out.Tool_LoginPage = instance.Tool.LoginPage

	case *ps.Object_Qualification_:
		out.Qualification_CertifiedBy = nref{Uid:instance.Qualification.CertifiedBy.Uid}

	case *ps.Object_Employee_:
		canBe := make([]nref, len(instance.Employee.CanBe))
		for i := range canBe {
			canBe[i] = nref{Uid:instance.Employee.CanBe[i].Uid}
		}

		out.Employee_CanBe = canBe
		out.Employee_RoleName = instance.Employee.RoleName

	case *ps.Object_Machine_:
		out.Machine_ProgramName = instance.Machine.ProgramName

	}

	return out, nil
}


func NewStateFromProto(pType *ps.State) (*state, error) {
	var objectUid string
	var objectName string

	switch v := pType.Object.(type){
	case *ps.State_StateObject:
		objectUid = v.StateObject.Uid
	case *ps.State_StateName:
		objectName = v.StateName
	}

	return &state{
		subject:   nref{Uid: pType.Subject.Uid},
		predicate: pType.Predicate,
		object:     struct {
			nref
			Name string
		} {
			nref: nref{Uid: objectUid},
			Name: objectName,
		},
	}, nil
}
