package resources




type resource struct {
	namespace namespace
	class class
	id string
}

type namespace string
const (
	pb namespace = "pb"
)


type class string
const (
	pbProcesses class = "process"
	pbStep class = "step"
	pbTool class = "tool"
	pbQualification class = "qualification"
	pbExecutor class = "executor"
	pbMachine class = "machine"
)


type relation string
const (
	pbType = "type"
	pbHas = "has"
	pbAre = "are"
	pbIn = "in"
)