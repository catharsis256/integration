package poller

type BtcDepositTask struct {
	baseDepositTask
}

func MakeBtcDepositTask() DepositTask {
	return &BtcDepositTask{}
}
