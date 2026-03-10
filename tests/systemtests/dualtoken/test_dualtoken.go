//go:build system_test

package dualtoken

import (
        "context"
        "math/big"
	"testing"

	"github.com/stretchr/testify/require"
	"github.com/tidwall/gjson"

	"github.com/cosmos/evm/tests/systemtests/suite"
	systest "github.com/cosmos/cosmos-sdk/testutil/systemtests"
)

// RunDualTokenTests exercises the dual token model on the running testnet:
// - TABB is used for staking
// - BRE is used for transaction fees and EVM gas
// - TABB has a fixed supply
func RunDualTokenTests(t *testing.T, base *suite.BaseTestSuite) {
	t.Helper()

	base.SetupTest(t)
	sut := base.SystemUnderTest

	cli := systest.NewCLIWrapper(t, sut, systest.Verbose)

	t.Run("Test Fee Enforcement", func(t *testing.T) {
		fromAcc := base.CosmosAccount("acc0")
		toAcc := base.CosmosAccount("acc1")
		from := fromAcc.AccAddress.String()
		to := toAcc.AccAddress.String()

		// 1. Transaction trying to pay fees in TABB should be rejected 
		// (min-gas-prices requires abre). The ante handler drops it before block commit.
		cli.WithRunErrorsIgnored().WithAssertTXUncommitted().Run("tx", "bank", "send", from, to, "1abre", "--from="+from, "--gas-prices=1000000000atabb", "--gas=auto", "--gas-adjustment=1.5", "-y")

		// 2. Exact same transaction paying fees in BRE should succeed
		rsp := cli.Run("tx", "bank", "send", from, to, "1abre", "--from="+from, "--gas-prices=1000000000abre", "--gas=auto", "--gas-adjustment=1.5", "-y")
		systest.RequireTxSuccess(t, rsp)
	})

	t.Run("Test Bonding Enforcement", func(t *testing.T) {
		fromAcc := base.CosmosAccount("acc0")
		from := fromAcc.AccAddress.String()
		// Get validator address for node0
		rawValAddr := cli.CustomQuery("q", "staking", "validators", "--output=json")
		valAddr := gjson.Get(rawValAddr, "validators.0.operator_address").String()
		require.NotEmpty(t, valAddr)

		// 1. Cannot bond BRE
		rsp, _ := cli.WithRunErrorsIgnored().RunOnly("tx", "staking", "delegate", valAddr, "100abre", "--from="+from, "--gas-prices=1000000000abre", "--gas=auto", "--gas-adjustment=1.5", "-y")
		require.Contains(t, rsp, "invalid coin denomination")

		// 2. Can bond TABB (while paying fees in BRE)
		rsp = cli.Run("tx", "staking", "delegate", valAddr, "100atabb", "--from="+from, "--gas-prices=1000000000abre", "--gas=auto", "--gas-adjustment=1.5", "-y")
		systest.RequireTxSuccess(t, rsp)
	})

	t.Run("Test Supply Integrity", func(t *testing.T) {
		// Query current total supply of TABB
		rawSupplyInitial := cli.CustomQuery("q", "bank", "total", "--denom=atabb", "--output=json")
		amountInitial := gjson.Get(rawSupplyInitial, "amount").String()
		require.NotEmpty(t, amountInitial)

		// Wait for a few blocks to ensure zero inflation
		sut.AwaitNBlocks(t, 2)

		rawSupplyFinal := cli.CustomQuery("q", "bank", "total", "--denom=atabb", "--output=json")
		amountFinal := gjson.Get(rawSupplyFinal, "amount").String()
		
		require.Equal(t, amountInitial, amountFinal, "TABB total supply inflated over blocks!")
	})

	t.Run("Test EVM JSON-RPC Gas Enforcement", func(t *testing.T) {
		// EVM transaction test using standard toolkit
		ctx := context.Background()
		
		acc0 := base.EthAccount("acc0")

		cliEVM := base.EthClient.Clients["node0"]
		require.NotNil(t, cliEVM)

		// Record initial balances
		balInitial, err := cliEVM.BalanceAt(ctx, acc0.Address, nil)
		require.NoError(t, err)

                base.SendTx(t, "node0", "acc0", 0, base.GasPriceMultiplier(2), big.NewInt(100))

		sut.AwaitNBlocks(t, 2)

		// Record final balances
		balFinal, err := cliEVM.BalanceAt(ctx, acc0.Address, nil)
		require.NoError(t, err)

                // Proof that Evm correctly used underlying BRE balances. Test is successful if no error is thrown by the Go eth client.
		require.True(t, balFinal.Cmp(balInitial) < 0, "Account 0 EVM balance did not decrease after transaction")
	})
}
