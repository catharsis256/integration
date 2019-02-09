package poller

type baseDepositTask struct{}

func (bi *baseDepositTask) Run() {
	bi.append()
	bi.synchronize()
	bi.clear()
}

func (bi *baseDepositTask) append() {

}

func (bi *baseDepositTask) synchronize() {

}

func (bi *baseDepositTask) clear() {

}
