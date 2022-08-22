package pool

type TaskPool struct {
	work chan func()   // 任务
	sem  chan struct{} // 数量
}

func NewTaskPool(size int) *TaskPool {
	return &TaskPool{
		work: make(chan func()),
		sem:  make(chan struct{}, size),
	}
}

func (p *TaskPool) NewTask(task func()) {
	select {
	case p.work <- task:
	case p.sem <- struct{}{}:
		go p.worker(task)
	}
}

func (p *TaskPool) worker(task func()) {
	defer func() {
		<-p.sem
	}()

	for {
		task()
		task = <-p.work
	}
}
