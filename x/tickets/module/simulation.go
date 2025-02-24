package tickets

import (
	"math/rand"

	sdk "github.com/cosmos/cosmos-sdk/types"
	"github.com/cosmos/cosmos-sdk/types/module"
	simtypes "github.com/cosmos/cosmos-sdk/types/simulation"
	"github.com/cosmos/cosmos-sdk/x/simulation"

	"tickets/testutil/sample"
	ticketssimulation "tickets/x/tickets/simulation"
	"tickets/x/tickets/types"
)

// avoid unused import issue
var (
	_ = ticketssimulation.FindAccount
	_ = rand.Rand{}
	_ = sample.AccAddress
	_ = sdk.AccAddress{}
	_ = simulation.MsgEntryKind
)

const (
	opWeightMsgCreateTicket = "op_weight_msg_create_ticket"
	// TODO: Determine the simulation weight value
	defaultWeightMsgCreateTicket int = 100

	opWeightMsgUpdateTicket = "op_weight_msg_update_ticket"
	// TODO: Determine the simulation weight value
	defaultWeightMsgUpdateTicket int = 100

	opWeightMsgDeleteTicket = "op_weight_msg_delete_ticket"
	// TODO: Determine the simulation weight value
	defaultWeightMsgDeleteTicket int = 100

	// this line is used by starport scaffolding # simapp/module/const
)

// GenerateGenesisState creates a randomized GenState of the module.
func (AppModule) GenerateGenesisState(simState *module.SimulationState) {
	accs := make([]string, len(simState.Accounts))
	for i, acc := range simState.Accounts {
		accs[i] = acc.Address.String()
	}
	ticketsGenesis := types.GenesisState{
		Params: types.DefaultParams(),
		// this line is used by starport scaffolding # simapp/module/genesisState
	}
	simState.GenState[types.ModuleName] = simState.Cdc.MustMarshalJSON(&ticketsGenesis)
}

// RegisterStoreDecoder registers a decoder.
func (am AppModule) RegisterStoreDecoder(_ simtypes.StoreDecoderRegistry) {}

// WeightedOperations returns the all the gov module operations with their respective weights.
func (am AppModule) WeightedOperations(simState module.SimulationState) []simtypes.WeightedOperation {
	operations := make([]simtypes.WeightedOperation, 0)

	var weightMsgCreateTicket int
	simState.AppParams.GetOrGenerate(opWeightMsgCreateTicket, &weightMsgCreateTicket, nil,
		func(_ *rand.Rand) {
			weightMsgCreateTicket = defaultWeightMsgCreateTicket
		},
	)
	operations = append(operations, simulation.NewWeightedOperation(
		weightMsgCreateTicket,
		ticketssimulation.SimulateMsgCreateTicket(am.accountKeeper, am.bankKeeper, am.keeper),
	))

	var weightMsgUpdateTicket int
	simState.AppParams.GetOrGenerate(opWeightMsgUpdateTicket, &weightMsgUpdateTicket, nil,
		func(_ *rand.Rand) {
			weightMsgUpdateTicket = defaultWeightMsgUpdateTicket
		},
	)
	operations = append(operations, simulation.NewWeightedOperation(
		weightMsgUpdateTicket,
		ticketssimulation.SimulateMsgUpdateTicket(am.accountKeeper, am.bankKeeper, am.keeper),
	))

	var weightMsgDeleteTicket int
	simState.AppParams.GetOrGenerate(opWeightMsgDeleteTicket, &weightMsgDeleteTicket, nil,
		func(_ *rand.Rand) {
			weightMsgDeleteTicket = defaultWeightMsgDeleteTicket
		},
	)
	operations = append(operations, simulation.NewWeightedOperation(
		weightMsgDeleteTicket,
		ticketssimulation.SimulateMsgDeleteTicket(am.accountKeeper, am.bankKeeper, am.keeper),
	))

	// this line is used by starport scaffolding # simapp/module/operation

	return operations
}

// ProposalMsgs returns msgs used for governance proposals for simulations.
func (am AppModule) ProposalMsgs(simState module.SimulationState) []simtypes.WeightedProposalMsg {
	return []simtypes.WeightedProposalMsg{
		simulation.NewWeightedProposalMsg(
			opWeightMsgCreateTicket,
			defaultWeightMsgCreateTicket,
			func(r *rand.Rand, ctx sdk.Context, accs []simtypes.Account) sdk.Msg {
				ticketssimulation.SimulateMsgCreateTicket(am.accountKeeper, am.bankKeeper, am.keeper)
				return nil
			},
		),
		simulation.NewWeightedProposalMsg(
			opWeightMsgUpdateTicket,
			defaultWeightMsgUpdateTicket,
			func(r *rand.Rand, ctx sdk.Context, accs []simtypes.Account) sdk.Msg {
				ticketssimulation.SimulateMsgUpdateTicket(am.accountKeeper, am.bankKeeper, am.keeper)
				return nil
			},
		),
		simulation.NewWeightedProposalMsg(
			opWeightMsgDeleteTicket,
			defaultWeightMsgDeleteTicket,
			func(r *rand.Rand, ctx sdk.Context, accs []simtypes.Account) sdk.Msg {
				ticketssimulation.SimulateMsgDeleteTicket(am.accountKeeper, am.bankKeeper, am.keeper)
				return nil
			},
		),
		// this line is used by starport scaffolding # simapp/module/OpMsg
	}
}
