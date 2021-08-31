package metrics

import "testing"

type testSink struct{}
func (t testSink) Put(s string, v interface{}) error {
	return nil

}


//interface check
var _ Sink = (*testSink)(nil)

func TestCollector_ReportMetrics(t *testing.T) {
	type fields struct {
		name string
	}
	type args struct {
		s Sink
	}
	tests := []struct {
		name    string
		fields  fields
		args    args
		wantErr bool
	}{
		{
			args:args{
				s: testSink{},
			},
		},
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			c := Collector{
				name: tt.fields.name,
			}
			if err := c.ReportMetrics(tt.args.s); (err != nil) != tt.wantErr {
				t.Errorf("ReportMetrics() error = %v, wantErr %v", err, tt.wantErr)
			}
		})
	}
}
