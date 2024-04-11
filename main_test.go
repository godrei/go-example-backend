package main

import (
	"testing"

	"github.com/stretchr/testify/require"
)

func TestPing(t *testing.T) {
	client := createClient()
	gotPong, err := pingDatabase(client)
	require.NoError(t, err)
	require.Equal(t, "PONG", gotPong)
}

func TestStoreName(t *testing.T) {
	wantName := "Elliot"

	client := createClient()
	err := writeName(client, wantName)
	require.NoError(t, err)

	gotName, err := readName(client)
	require.NoError(t, err)
	require.Equal(t, wantName, gotName)
}
