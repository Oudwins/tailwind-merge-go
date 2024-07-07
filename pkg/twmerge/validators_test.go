package twmerge

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestArbitraryShadow(t *testing.T) {
	assert.Equal(t, IsArbitraryShadow("[inset_0_1px_0,inset_0_-1px_0]"), true)
	assert.Equal(t, IsArbitraryShadow("[0_35px_60px_-15px_rgba(0,0,0,0.3)]"), true)
	assert.Equal(t, IsArbitraryShadow("[inset_0_1px_0,inset_0_-1px_0]"), true)
	assert.Equal(t, IsArbitraryShadow("[0_0_#00f]"), true)
	assert.Equal(t, IsArbitraryShadow("[.5rem_0_rgba(5,5,5,5)]"), true)
	assert.Equal(t, IsArbitraryShadow("[-.5rem_0_#123456]"), true)
	assert.Equal(t, IsArbitraryShadow("[0.5rem_-0_#123456]"), true)
	assert.Equal(t, IsArbitraryShadow("[0.5rem_-0.005vh_#123456]"), true)
	assert.Equal(t, IsArbitraryShadow("[0.5rem_-0.005vh]"), true)

	assert.Equal(t, IsArbitraryShadow("[rgba(5,5,5,5)]"), false)
	assert.Equal(t, IsArbitraryShadow("[#00f]"), false)
	assert.Equal(t, IsArbitraryShadow("[something-else]"), false)
}
