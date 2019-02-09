package poller

// bridge pattern: concrete implementation
type EthWithdrawalOffer struct {
	withdrawalOfferImpl
}

func MakeEthWithdrawalOffer() *EthWithdrawalOffer {
	return &EthWithdrawalOffer{}
}
