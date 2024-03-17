// Copyright 2023 Aiden Fox Ivey, Alisya K. All rights reserved.
// Use of this source code is governed by an epic-style
// license that can be found in the LICENSE file.

package merkle

import (
	"bytes"
	"crypto/sha256"
	"testing"
)

func Test_NewMerkleTree(t *testing.T) {
	_, err := NewMerkleTree([]byte{}, 1)
	if err == nil {
		t.Fatal("Expected error, got no error.")
	}
}

func Test_HashListFromBytes(t *testing.T) {
	b := []byte("hello")
	h, err := HashListFromBytes(b, 4)
	if err != nil {
		t.Fatal("Expected no error, got error.")
	}

	expectedBlocks := [][]byte{
		[]byte("hell"),
		[]byte("o"),
	}

	if len(h.nodes) != len(expectedBlocks) {
		t.Fatalf("Expected %d blocks, got %d.\n", len(h.nodes), len(expectedBlocks))
	}

	for i, node := range h.nodes {
		shaSum := sha256.Sum256(expectedBlocks[i])
		if !bytes.Equal(node.hash[:], shaSum[:]) {
			t.Fatal("Mismatching bytes.")
		}
	}
}

func Test_Bytes(t *testing.T) {
	b := []byte("hello")
	nodes := splitIntoDataBlocks(b, 4)

	expectedBlocks := [][]byte{
		[]byte("hell"),
		[]byte("o"),
	}

	if len(nodes) != len(expectedBlocks) {
		t.Fatalf("Expected %d blocks, got %d.\n", len(nodes), len(expectedBlocks))
	}

	for i, node := range nodes {
		shaSum := sha256.Sum256(expectedBlocks[i])
		if !bytes.Equal(node.value, shaSum[:]) {
			t.Fatal("Mismatching bytes.")
		}
	}
}
