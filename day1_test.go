package main

import (
	"testing"
)

func TestCalculateTax(t *testing.T) {
	type args struct {
		income uint64
	}
	tests := []struct {
		name string
		args args
		want uint64
	}{
		{
			name: "step_1",
			args: args{income:10000},
			want: 0,
		},
		{
			name: "step_2",
			args: args{income:15000},
			want: 150,
		},
		{
			name: "step_3",
			args: args{income:25000},
			want: 500,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := CalculateTax(tt.args.income); got != tt.want {
				t.Errorf("CalculateTax() = %v, want %v", got, tt.want)
			}
		})
	}
}
