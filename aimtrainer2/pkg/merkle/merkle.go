// Copyright 2023 Aiden Fox Ivey, Alisya K. All rights reserved.
// Use of this source code is governed by an epic-style
// license that can be found in the LICENSE file.

// Package merkle implements a SHA256-backed Merkle tree.

package merkle

import (
	"crypto/sha256"
	"fmt"
)

// placeholder name for a node in a hashlist
// idk how long it should be
type HashListNode struct {
	data []byte
	hash [32]byte
}

type HashList struct {
	nodes   []HashListNode
	TopHash [32]byte
}

func HashListFromBytes(data []byte, blockSize int) (HashList, error) {
	t := HashList{}

	if len(data) == 0 {
		return t, fmt.Errorf("No data was provided.")
	}

	var hashes []byte

	for i := 0; i < len(data); i += blockSize {
		end := i + blockSize

		if end > len(data) {
			end = len(data)
		}

		b := sha256.Sum256(data[i:end])
		hashes = append(hashes, b[:]...)
		t.nodes = append(t.nodes, HashListNode{data[i:end], b})
	}

	t.TopHash = sha256.Sum256(hashes)

	return t, nil
}

type MerkleTree struct {
	root   *Node
	leaves []*Node
}

type Node struct {
	left   *Node
	right  *Node
	parent *Node
	value  []byte
}

func (n Node) hash() []byte {
	hashArray := sha256.Sum256(n.value)
	return hashArray[:]
}

// Consume the bytes b of some serialized object, and return a valid Merkle tree.
func NewMerkleTree(data []byte, blockSize int) (*Node, error) {
	t := &MerkleTree{}
	if len(data) == 0 {
		return nil, fmt.Errorf("No data was provided.")
	}

	dataBlocks := splitIntoDataBlocks(data, blockSize)

	// assign leaves to tree
	t.leaves = dataBlocks

	// TODO: determine why I've init'ed at size of len(dataBlocks)
	q := make([]*Node, 0, len(dataBlocks))

	// add leaf nodes into queue
	q = append(q, dataBlocks...)

	for len(q) > 1 {
		var n1, n2 *Node
		n1, q = q[0], q[1:]
		n2, q = q[0], q[1:]

		buff := make([]byte, 0)
		buff = append(buff, n1.value...)
		buff = append(buff, n2.value...)

		sum := sha256.Sum256(buff)

		newNode := &Node{n1, n2, nil, sum[:]}

		n1.parent = newNode
		n2.parent = newNode

		q = append(q, newNode)
	}

	return q[0], nil
}

// -\_(*_*)_/-
// - Alisya K.
func splitIntoDataBlocks(b []byte, blockSize int) []*Node {
	output := make([]*Node, 0)

	for i := 0; i < len(b); i += blockSize {
		end := i + blockSize

		if end > len(b) {
			end = len(b)
		}

		b := sha256.Sum256(b[i:end])
		output = append(output, &Node{nil, nil, nil, b[:]})
	}
	return output
}

// spits out the root hash of the merkle tree
func Apply(b []byte) uint32 {
	// dataBlocks := splitIntoDataBlocks(b)
	// rootHash := recursion(dataBlocks)
	// return rootHash

	return 3
}

func makeNode(left, right *Node) *Node {
	var output []byte
	output = append(left.value, right.value...)
	f := sha256.Sum256(output)
	return &Node{left: left, right: right, value: f[:]}
}
