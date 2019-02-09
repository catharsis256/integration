package poller

type PollingStrategy interface {
	Poll()
}
