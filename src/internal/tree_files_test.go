package internal

import (
	"testing"

	"github.com/charmbracelet/lipgloss/tree"
	"github.com/stretchr/testify/assert"
)

func TestTreeFiles(t *testing.T) {
	treeCli, err := GetTree(".", 2, 1, "")
	assert.NoError(t, err)
	assert.IsType(t, &tree.Tree{}, treeCli)

	treeCli, err = GetTree(".", 5, 1, "")
	assert.NoError(t, err)

	treeCli, err = GetTree(".", 5, 20, "")
	assert.NoError(t, err)
}
