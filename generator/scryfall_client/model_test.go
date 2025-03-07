package scryfallclient_test

import (
	"testing"

	scryfallclient "github.com/djpiper28/setmoji/generator/scryfall_client"
	"github.com/stretchr/testify/require"
)

func TestGetSets(t *testing.T) {
	t.Parallel()

	sets, err := scryfallclient.GetSets()
	require.NoError(t, err)
	require.NotEmpty(t, sets)
}

func TestGetSetSvg(t *testing.T) {
	t.Parallel()

	sets, err := scryfallclient.GetSets()
	require.NoError(t, err)
	require.NotEmpty(t, sets)

	svg, err := sets[0].GetSvg()
	require.NoError(t, err)
	require.NotEmpty(t, svg)
}
