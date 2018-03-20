package kernel

import "github.com/agoussia/godes"

var comp_sci *godes.UniformDistr
var market *godes.UniformDistr
var astro *godes.UniformDistr
var vis_art *godes.UniformDistr
var biology *godes.UniformDistr

type Major struct {
	name  string
	value float64
}

func NewMajor(name string, initial_value float64) Major {
	major := Major{}
	major.name = name
	major.value = initial_value
	return major
}

// InitMajorDists Must be called first
func InitMajorDists() {
	comp_sci = godes.NewUniformDistr(true)
	market = godes.NewUniformDistr(true)
	astro = godes.NewUniformDistr(true)
	vis_art = godes.NewUniformDistr(true)
	biology = godes.NewUniformDistr(true)
}
