package neuron

import "neural-network/math"

type Activation func(input float64) (output float64)

type Neuron interface {
	Compute(inputs []float64, weights []float64, bias float64) float64
}

type neuron struct {
	activation Activation
}

func (n *neuron) Compute(inputs []float64, weights []float64, bias float64) float64 {
	sum := math.DotProduct(inputs, weights)
	sum += bias

	res := n.activation(sum)
	return res
}
