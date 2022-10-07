package pipelines

import "fmt"

type DispatchTask struct {
	file string
}

type DispatchPipeline struct {
	DipatchChannel chan DispatchTask
}

func (p *DispatchPipeline) Run() {
	dc := make(chan DispatchTask)
	go func() {
		t := <-dc
		fmt.Println(t.file)
	}()
	p.DipatchChannel = dc
}
