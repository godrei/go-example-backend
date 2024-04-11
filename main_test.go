package main

import (
	"github.com/stretchr/testify/require"
	"testing"
)

func TestStoreName(t *testing.T) {
	wantName := "Elliot"

	client := createClient()
	err := writeName(client, wantName)
	require.NoError(t, err)

	gotName, err := readName(client)
	require.NoError(t, err)
	require.Equal(t, wantName, gotName)
}
