package v1

import "log"

type Control struct {
	controlled any
	stopping   chan struct{}
}

func NewControl(controlled any) (Controlling, Controlled) {
	c := new(Control)
	c.controlled = controlled
	c.stopping = make(chan struct{})
	return c, c
}

func (c *Control) GetControlled() any {
	return c.controlled
}

func (c *Control) GetStopping() <-chan struct{} {
	return c.stopping
}

func (c *Control) Close() {
	log.Printf("%T.Close: closing", c.controlled)
	close(c.stopping)
}

func (c *Control) Stop(controlling any) {
	c.asyncStop(controlling)
	c.asyncCloseWait(controlling)
}

func (c *Control) asyncStop(controlling any) {
	log.Printf("%T.Stop by %T: stopping", c.controlled, controlling)
	c.stopping <- *new(struct{})
}

func (c *Control) asyncCloseWait(controlling any) {
	_, ok := <-c.stopping
	log.Printf("%T.CloseWait by %T: closed '%v'", c.controlled, controlling, !ok)
}
