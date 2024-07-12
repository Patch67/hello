package mynodes

import (
	"reflect"
	"testing"
)

func TestNewLabel(t *testing.T) {
	type args struct {
		name       string
		properties []Attribute
	}
	tests := []struct {
		name string
		args args
		want *Label
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := NewLabel(tt.args.name, tt.args.properties); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("NewLabel() = %v, want %v", got, tt.want)
			}
		})
	}
}
