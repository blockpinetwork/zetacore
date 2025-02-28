package smoketests

import "github.com/zeta-chain/zetacore/contrib/localnet/orchestrator/smoketest/runner"

const (
	TestContextUpgradeName              = "context_upgrade"
	TestDepositAndCallRefundName        = "deposit_and_call_refund"
	TestMultipleERC20DepositName        = "erc20_multiple_deposit"
	TestWithdrawERC20Name               = "erc20_withdraw"
	TestMultipleWithdrawsName           = "erc20_multiple_withdraw"
	TestSendZetaOutName                 = "send_zeta_out"
	TestSendZetaOutBTCRevertName        = "send_zeta_out_btc_revert" // #nosec G101 - not a hardcoded password
	TestMessagePassingName              = "message_passing"
	TestZRC20SwapName                   = "zrc20_swap"
	TestBitcoinWithdrawName             = "bitcoin_withdraw"
	TestCrosschainSwapName              = "crosschain_swap"
	TestMessagePassingRevertFailName    = "message_passing_revert_fail"
	TestMessagePassingRevertSuccessName = "message_passing_revert_success"
	TestPauseZRC20Name                  = "pause_zrc20"
	TestERC20DepositAndCallRefundName   = "erc20_deposit_and_call_refund"
	TestUpdateBytecodeName              = "update_bytecode"
	TestEtherDepositAndCallName         = "eth_deposit_and_call"
	TestDepositEtherLiquidityCapName    = "deposit_eth_liquidity_cap"
	TestMyTestName                      = "my_test"

	TestERC20DepositName = "erc20_deposit"
	TestEtherDepositName = "eth_deposit"
)

// AllSmokeTests is an ordered list of all smoke tests
var AllSmokeTests = []runner.SmokeTest{
	{
		TestContextUpgradeName,
		"tests sending ETH on ZEVM and check context data using ContextApp",
		TestContextUpgrade,
	},
	{
		TestDepositAndCallRefundName,
		"deposit ZRC20 into ZEVM and call a contract that reverts; should refund",
		TestDepositAndCallRefund,
	},
	{
		TestMultipleERC20DepositName,
		"deposit USDT ERC20 into ZEVM in multiple deposits",
		TestMultipleERC20Deposit,
	},
	{
		TestWithdrawERC20Name,
		"withdraw USDT ERC20 from ZEVM",
		TestWithdrawERC20,
	},
	{
		TestMultipleWithdrawsName,
		"withdraw USDT ERC20 from ZEVM in multiple deposits",
		TestMultipleWithdraws,
	},
	{
		TestSendZetaOutName,
		"sending ZETA from ZEVM to Ethereum",
		TestSendZetaOut,
	},
	{
		TestSendZetaOutBTCRevertName,
		"sending ZETA from ZEVM to Bitcoin; should revert when ",
		TestSendZetaOutBTCRevert,
	},
	{
		TestMessagePassingName,
		"goerli->goerli message passing (sending ZETA only)",
		TestMessagePassing,
	},
	{
		TestZRC20SwapName,
		"swap ZRC20 USDT for ZRC20 ETH",
		TestZRC20Swap,
	},
	{
		TestBitcoinWithdrawName,
		"withdraw BTC from ZEVM",
		TestBitcoinWithdraw,
	},
	{
		TestCrosschainSwapName,
		"testing Bitcoin ERC20 cross-chain swap",
		TestCrosschainSwap,
	},
	{
		TestMessagePassingRevertFailName,
		"goerli->goerli message passing (revert fail)",
		TestMessagePassingRevertFail,
	},
	{
		TestMessagePassingRevertSuccessName,
		"goerli->goerli message passing (revert success)",
		TestMessagePassingRevertSuccess,
	},
	{
		TestPauseZRC20Name,
		"pausing ZRC20 on ZetaChain",
		TestPauseZRC20,
	},
	{
		TestERC20DepositAndCallRefundName,
		"deposit a non-gas ZRC20 into ZEVM and call a contract that reverts; should refund on ZetaChain if no liquidity pool, should refund on origin if liquidity pool",
		TestERC20DepositAndCallRefund,
	},
	{
		TestUpdateBytecodeName,
		"update ZRC20 bytecode swap",
		TestUpdateBytecode,
	},
	{
		TestEtherDepositAndCallName,
		"deposit ZRC20 into ZEVM and call a contract",
		TestEtherDepositAndCall,
	},
	{
		TestDepositEtherLiquidityCapName,
		"deposit Ethers into ZEVM with a liquidity cap",
		TestDepositEtherLiquidityCap,
	},
	{
		TestMyTestName,
		"performing custom test",
		TestMyTest,
	},
	{
		TestERC20DepositName,
		"deposit ERC20 into ZEVM",
		TestERC20Deposit,
	},
	{
		TestEtherDepositName,
		"deposit Ether into ZEVM",
		TestEtherDeposit,
	},
}
