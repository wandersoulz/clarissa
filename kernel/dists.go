package kernel

import "github.com/wandersoulz/godes"

var peopleIndexDist *godes.UniformDistr
var numPeopleDist *godes.TriangularDistr
var randomEncounter *godes.UniformDistr
var randomTime *godes.TriangularDistr

var friendEncDist *godes.FunctionalDistr

var friendshipDist *godes.TriangularDistr
var scalarDist *godes.NormalDistr

var sleepTimeDist *godes.TriangularDistr

var compsci *godes.UniformDistr
var market *godes.UniformDistr
var astro *godes.UniformDistr
var visart *godes.UniformDistr
var biology *godes.UniformDistr
var prefDist *godes.UniformDistr

// InitMajorDists Must be called first
func InitMajorDists() {
	compsci = godes.NewUniformDistr(true)
	market = godes.NewUniformDistr(true)
	astro = godes.NewUniformDistr(true)
	visart = godes.NewUniformDistr(true)
	biology = godes.NewUniformDistr(true)
	prefDist = godes.NewUniformDistr(true)
}

func InitBusDists() {
	peopleIndexDist = godes.NewUniformDistr(true)
	numPeopleDist = godes.NewTriangularDistr(true)
	randomEncounter = godes.NewUniformDistr(true)
	randomTime = godes.NewTriangularDistr(true)
}

func InitHomeDists() {
	friendEncDist = godes.NewFunctionalDistr(true)
	sleepTimeDist = godes.NewTriangularDistr(true)
}
