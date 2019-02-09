package poller

// bridge pattern: concrete implementation
type BtcWithdrawalOffer struct {
	withdrawalOfferImpl
}

func MakeBtcWithdrawalOffer() *BtcWithdrawalOffer {
	return &BtcWithdrawalOffer{}
}
