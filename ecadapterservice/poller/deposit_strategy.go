package poller

type DepositStrategy struct {
	task DepositTask
}

func MakeDepositStrategy(task DepositTask) PollingStrategy {
	return &DepositStrategy{task: task}
}

func (bis *DepositStrategy) Poll() {
	bis.task.Run()
}
