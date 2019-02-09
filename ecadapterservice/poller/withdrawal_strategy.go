package poller

type WithdrawalStrategy struct {
}

func MakeWithdrawalStrategy() PollingStrategy {
	return &WithdrawalStrategy{}
}

func (bis *WithdrawalStrategy) Poll() {

}
