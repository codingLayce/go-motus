package utils

import (
	"github.com/stretchr/testify/require"
	"testing"
)

func TestContains(t *testing.T) {
	strList := []string{"a", "b", "c"}
	intList := []int{1, 2, 3}

	require.True(t, Contains(strList, "a"))
	require.True(t, Contains(strList, "b"))
	require.True(t, Contains(strList, "c"))
	require.False(t, Contains(strList, "d"))

	require.True(t, Contains(intList, 1))
	require.True(t, Contains(intList, 2))
	require.True(t, Contains(intList, 3))
	require.False(t, Contains(intList, 15))
}

func TestContainsChar(t *testing.T) {
	str := []rune("petit")

	require.Equal(t, 0, IndexOf(str, 'p', 0))
	require.Equal(t, 1, IndexOf(str, 'e', 0))
	require.Equal(t, 2, IndexOf(str, 't', 0))
	require.Equal(t, 3, IndexOf(str, 'i', 0))
	require.Equal(t, 4, IndexOf(str, 't', 1))

	require.Equal(t, -1, IndexOf(str, 'a', 0))
	require.Equal(t, -1, IndexOf(str, 'p', 15))
	require.Equal(t, -1, IndexOf(str, 'P', 0))
	require.Equal(t, -1, IndexOf(str, 'E', 0))
	require.Equal(t, -1, IndexOf(str, 'T', 0))
	require.Equal(t, -1, IndexOf(str, 'I', 0))
	require.Equal(t, -1, IndexOf(str, 'T', 1))
}
