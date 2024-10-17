package q

import (
	"testing"
)

func TestTree(t *testing.T) {
	rootValue := "root"
	tree := Tree(rootValue)

	if !tree.Root {
		t.Errorf("expected Root to be true, got false")
	}

	if tree.Data != rootValue {
		t.Errorf("expected Data to be %v, got %v", rootValue, tree.Data)
	}

	if len(tree.Children) != 0 {
		t.Errorf("expected Children to be empty, got %v", len(tree.Children))
	}
}

func TestAdd(t *testing.T) {
	rootValue := "root"
	childValue := "child"
	tree := Tree(rootValue)
	childNode := tree.Add(childValue)

	if len(tree.Children) != 1 {
		t.Errorf("expected Children length to be 1, got %v", len(tree.Children))
	}

	if tree.Children[0] != childNode {
		t.Errorf("expected child node to be added to Children")
	}

	if tree.Children[0].Data != childValue {
		t.Errorf("expected child Data to be %v, got %v", childValue, tree.Children[0].Data)
	}

	if tree.Children[0].Root {
		t.Errorf("expected child Root to be false, got true")
	}
}

func TestTraverse(t *testing.T) {
	// Create a tree structure
	rootValue := "root"
	tree := Tree(rootValue)

	// Add child nodes
	child1 := tree.Add("child1")
	child2 := tree.Add("child2")

	// Add grandchildren
	child1.Add("grandchild1")
	child1.Add("grandchild2")
	child2.Add("grandchild3")

	// Expected order of traversal
	expectedOrder := []string{"root", "child1", "grandchild1", "grandchild2", "child2", "grandchild3"}

	// Slice to store the order of traversal
	var traversalOrder []string

	// Callback function to capture the traversal order
	cb := func(it string) {
		traversalOrder = append(traversalOrder, it)
	}

	// Perform the traversal
	tree.Traverse(cb)

	// Check if the traversal order matches the expected order
	if len(traversalOrder) != len(expectedOrder) {
		t.Fatalf("expected traversal order length to be %d, got %d", len(expectedOrder), len(traversalOrder))
	}

	for i, v := range expectedOrder {
		if traversalOrder[i] != v {
			t.Errorf("expected traversal order at index %d to be %v, got %v", i, v, traversalOrder[i])
		}
	}
}
