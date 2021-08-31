package metrics

type Collector struct {
	name string
}

func (c Collector) ReportMetrics(s Sink) error {
	return s.Put("", "anything")
}

func (c Collector) Observe(metric string, value interface{})  {

}


type Sink interface {
	Put(k string, v interface{}) error
}
