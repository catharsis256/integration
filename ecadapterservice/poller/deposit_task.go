package poller

type DepositTask interface {
	Task
	append()
	synchronize()
	clear()
}
