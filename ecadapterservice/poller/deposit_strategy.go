package poller

type DepositStrategy struct {
	Task DepositTask
}

func MakeDepositStrategy(task DepositTask) PollingStrategy {
	return &DepositStrategy{Task: task}
}

func (bis *DepositStrategy) Poll() {
	bis.Task.Run()
}
