// Package linker inspects process nodes (as defined in playbook/apis/procstore)
// and correctly disambiguates them from one another, providing UUIDs for each.
// These UUIDs are then used by the dbtranslation package to correctly serialize
// Nodes to be shoved into database.
package linker

import "playbook/apis/procstore"

// Hash_ids on nodes are used for speeding up queries involving similar nodes
// They are not however universally applicable to the process of writing nodes
// to the DB. All Nodes should have a hash_id, but only some nodes should have a hash_id
// that is their UUID. Steps should never be "duplicated" in a process. By which I mean,
// No one step should have the same name, and exist more than once in a process. Because of
// This, we use it's hash_id to disambiguate it. In all other cases we assign a unique uid
// to the Node. otherwise things would in collide in odd ways
//
// All Nodes must satisfy this interface
type Node interface {
	GetUid() string
	GetHashId() string
}

// TODO(hdh) enforce "unique name per step Node in process" somehow
func GetUidReference(n Node) string {
	switch n.(type) {
	case *procstore.Step:
		return n.GetHashId()
	default:
		return n.GetUid()
	}
}