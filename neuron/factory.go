package neuron

import "neural-network/math"

func NewReLUNeuron() Neuron {
	return &neuron{
		activation: math.ReLU,
	}
}

func NewPerceptron() Neuron {
	return &neuron{
		activation: math.Step,
	}
}

func NewSigmoidNeuron() Neuron {
	return &neuron{
		activation: math.Sigmoid,
	}
}
