package main

/*
#include <stdio.h>
#include <math.h>
#include "lib/emd.h"
#cgo LDFLAGS: -L. -llib
*/
import "C"
import "fmt"

func main() {
	f1 := []C.feature_t{{100, 40, 22}, {211, 20, 2}, {32, 190, 150}, {2, 100, 100}}
	f2 := []C.feature_t{{0, 0, 0}, {50, 100, 80}, {255, 255, 255}}
	w1 := []C.float{0.4, 0.3, 0.2, 0.1}
	w2 := []C.float{0.5, 0.3, 0.2}

	s1 := C.signature_t{n: 4, Weights: &w1[0], Features: &f1[0]}
	s2 := C.signature_t{n: 3, Weights: &w2[0], Features: &f2[0]}

	var flow C.flow_t
	var flowsize C.int

	edist := C.emd(&s1, &s2, (*[0]byte)(C.dist), &flow, &flowsize)
	fmt.Printf("emd = %f\n", edist)
}
