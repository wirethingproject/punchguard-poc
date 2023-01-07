package v1

type OnDemandFlow struct {
	BaseFlow
}

func NewOnDemandFlow(id Id, signal Signal, otr Otr, punch Punch, guard Guard) (Flow, error) {
	o := new(OnDemandFlow)
	if err := o.Init(id, signal, otr, punch, guard); err != nil {
		return nil, err
	}
	return o, nil
}

func (o *OnDemandFlow) Run() *Control {
	return nil
}
