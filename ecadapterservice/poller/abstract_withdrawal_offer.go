package poller

// abstraction of bridge pattern
type AbstractWithdrawalOffer struct {
	withdrawalOffer *withdrawalOfferImpl
}

func MakeAbstractWithdrawalOffer(withdrawalOffer *withdrawalOfferImpl) *AbstractWithdrawalOffer {
	return &AbstractWithdrawalOffer{withdrawalOffer: withdrawalOffer}
}

func (wo *AbstractWithdrawalOffer) ProcessPendingOffer() {
	wo.withdrawalOffer.searchOffer()
}

func (wo *AbstractWithdrawalOffer) ProcessExecutingOffer() {
	wo.withdrawalOffer.searchOffer()

}

func (wo *AbstractWithdrawalOffer) ProcessAcceptedOffer() {
	wo.withdrawalOffer.searchOffer()

}
