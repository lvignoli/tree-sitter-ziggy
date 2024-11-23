package tree_sitter_ziggy_test

import (
	"testing"

	tree_sitter "github.com/tree-sitter/go-tree-sitter"
	tree_sitter_ziggy "github.com/tree-sitter/tree-sitter-ziggy/bindings/go"
)

func TestCanLoadGrammar(t *testing.T) {
	language := tree_sitter.NewLanguage(tree_sitter_ziggy.Language())
	if language == nil {
		t.Errorf("Error loading Ziggy grammar")
	}
}