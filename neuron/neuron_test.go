package neuron

import "testing"

func TestPerceptron(t *testing.T) {
	p := NewPerceptron()
	type args struct {
		inputs  []float64
		weights []float64
		bias    float64
	}
	tests := []struct {
		name string
		args args
		want float64
	}{
		{
			name: "t1",
			args: args{
				inputs:  []float64{0, 1, 1},
				weights: []float64{5, 2, 2},
				bias:    -4,
			},
			want: 0,
		},
		{
			name: "t2",
			args: args{
				inputs:  []float64{0, 1, 1},
				weights: []float64{5, 2, 2},
				bias:    -3,
			},
			want: 1,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := p.Compute(tt.args.inputs, tt.args.weights, tt.args.bias); got != tt.want {
				t.Errorf("Compute() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestReLuNeuron(t *testing.T) {
	p := NewReLUNeuron()
	type args struct {
		inputs  []float64
		weights []float64
		bias    float64
	}
	tests := []struct {
		name string
		args args
		want float64
	}{
		{
			name: "t1",
			args: args{
				inputs:  []float64{1, 2, 4},
				weights: []float64{3, 2, 2},
				bias:    -10,
			},
			want: 5,
		},
		{
			name: "t2",
			args: args{
				inputs:  []float64{0, 1, 2},
				weights: []float64{5, 8, 2},
				bias:    -6,
			},
			want: 6,
		},
		{
			name: "t2",
			args: args{
				inputs:  []float64{2, 1, 2},
				weights: []float64{5, 8, 3},
				bias:    -30,
			},
			want: 0,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := p.Compute(tt.args.inputs, tt.args.weights, tt.args.bias); got != tt.want {
				t.Errorf("Compute() = %v, want %v", got, tt.want)
			}
		})
	}
}
