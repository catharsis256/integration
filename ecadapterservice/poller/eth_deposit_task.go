package poller

type EthDepositTask struct {
	baseDepositTask
}

func MakeEthDepositTask() DepositTask {
	return &EthDepositTask{}
}
