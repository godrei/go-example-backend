package main

import (
	"testing"

	"github.com/stretchr/testify/require"
)

func TestPing(t *testing.T) {
	client := createDBClient()
	gotPong, err := pingDB(client)
	require.NoError(t, err)
	require.Equal(t, "PONG", gotPong)
}

func TestStoreName(t *testing.T) {
	wantName := "Elliot"

	client := createDBClient()
	err := writeNameToDB(client, wantName)
	require.NoError(t, err)

	gotName, err := readNameFromDB(client)
	require.NoError(t, err)
	require.Equal(t, wantName, gotName)
}
