package keeper_test

import (
	"testing"

	"github.com/stretchr/testify/require"
	testkeeper "himitsu/testutil/keeper"
	"himitsu/x/himitsu/types"
)

func TestGetParams(t *testing.T) {
	k, ctx := testkeeper.HimitsuKeeper(t)
	params := types.DefaultParams()

	k.SetParams(ctx, params)

	require.EqualValues(t, params, k.GetParams(ctx))
}
