package simulation

import (
	"math/rand"

	"github.com/andReyM228/one/x/core/keeper"
	"github.com/andReyM228/one/x/core/types"
	"github.com/cosmos/cosmos-sdk/baseapp"
	sdk "github.com/cosmos/cosmos-sdk/types"
	simtypes "github.com/cosmos/cosmos-sdk/types/simulation"
)

func SimulateMsgIssue(
	ak types.AccountKeeper,
	bk types.BankKeeper,
	k keeper.Keeper,
) simtypes.Operation {
	return func(r *rand.Rand, app *baseapp.BaseApp, ctx sdk.Context, accs []simtypes.Account, chainID string,
	) (simtypes.OperationMsg, []simtypes.FutureOperation, error) {
		simAccount, _ := simtypes.RandomAcc(r, accs)
		msg := &types.MsgIssue{
			Creator: simAccount.Address.String(),
		}

		// TODO: Handling the Issue simulation

		return simtypes.NoOpMsg(types.ModuleName, msg.Type(), "Issue simulation not implemented"), nil, nil
	}
}
