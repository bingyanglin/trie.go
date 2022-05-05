package trie256p

import (
	"fmt"
	trie_go "github.com/iotaledger/trie.go"
)

// CommitmentModel abstracts 256+ Trie logic from the commitment logic/cryptography
type CommitmentModel interface {
	// NewVectorCommitment creates empty trie_go.VCommitment
	NewVectorCommitment() trie_go.VCommitment
	// NewTerminalCommitment creates empty trie_go.TCommitment
	NewTerminalCommitment() trie_go.TCommitment
	// CommitToData calculates terminal commitment to an arbitrary data
	CommitToData([]byte) trie_go.TCommitment
	// CalcNodeCommitment calculates commitment of the node data
	CalcNodeCommitment(*NodeData) trie_go.VCommitment
	// UpdateNodeCommitment updates mutable NodeData with the update information.
	// It also (optionally, if 'update' != nil) updates previous commitment to the node
	// If update != nil and *update != nil, parameter calcDelta specifies if commitment is calculated
	// from scratch using CalcNodeCommitment, or it can be calculated by applying additive delta
	// I can be used by implementation to optimize the computation of update. For example KZG implementation
	// can be made dramatically faster this way than strictly computing each time whole expensive vector commitment
	// This interface takes into account different ways how updates are propagated in the trie
	UpdateNodeCommitment(mutate *NodeData, childUpdates map[byte]trie_go.VCommitment, calcDelta bool, terminal trie_go.TCommitment, update *trie_go.VCommitment)
	// Description return description of the implementation
	Description() string
}
type PathArity byte

const (
	PathArity256 = PathArity(255)
	PathArity16  = PathArity(15)
	PathArity2   = PathArity(1)
)

func (a PathArity) String() string {
	switch a {
	case PathArity256, PathArity16, PathArity2:
		return fmt.Sprintf("PathArity(%d)", a)
	default:
		return "PathArity(wrong)"
	}
}
