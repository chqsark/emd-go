package main

/*
#include <stdio.h>
#include <math.h>
#include "lib/emd.h"
#cgo LDFLAGS: -L. -llib
*/
import "C"
import (
	"bufio"
	"fmt"
	"os"
	"strings"

	wordmodel "github.com/thinxer/go-word2vec"
)

func edist(sent1, sent2 string, m *wordmodel.Model) C.float {
	words1 := strings.Split(sent1, " ")
	words2 := strings.Split(sent2, " ")
	vec1 := [][]float32{}
	vec2 := [][]float32{}

	for _, w1 := range words1 {
		vec := wordmodel.Vector(make([]float32, m.Layer1Size))
		if wordId, ok := m.Vocab[w1]; !ok {
			fmt.Printf("word not found: %s\n", w1)
		} else {
			vec.Add(1, m.Vector(wordId))
		}
		vec.Normalize()
		vec1 = append(vec1, vec)
	}
	for _, w2 := range words2 {
		vec := wordmodel.Vector(make([]float32, m.Layer1Size))
		if wordId, ok := m.Vocab[w2]; !ok {
			fmt.Printf("word not found: %s\n", w2)
		} else {
			vec.Add(1, m.Vector(wordId))
		}
		vec.Normalize()
		vec2 = append(vec2, vec)
	}

	f1, f2 := make([]C.feature_t, len(vec1)), make([]C.feature_t, len(vec2))
	for i, v := range vec1 {
		f1[i].arr = (*C.float)(&v[0])
	}
	for i, v := range vec2 {
		f2[i].arr = (*C.float)(&v[0])
	}

	wt1, wt2 := make([]C.float, len(vec1[0])), make([]C.float, len(vec2[0]))

	for i := 0; i < len(wt1); i++ {
		wt1[i], wt2[i] = 1.0, 1.0
	}

	s1 := C.signature_t{n: (C.int)(len(words1)), Weights: &wt1[0], Features: &f1[0]}
	s2 := C.signature_t{n: (C.int)(len(words2)), Weights: &wt2[0], Features: &f2[0]}

	var flow C.flow_t
	var flowsize C.int

	emd_dist := C.emd(&s1, &s2, (*[0]byte)(C.dist), &flow, &flowsize)
	return emd_dist
}

func main() {
	model, err := wordmodel.Load(os.Args[1])
	if err != nil {
		fmt.Println(err)
		return
	}
	scanner := bufio.NewScanner(os.Stdin)
	sents := make([]string, 2)
	for {
		fmt.Println("input:")
		i := 0
		for scanner.Scan() && i <= 1 {
			fmt.Printf("sent%d:\n", i)
			line := scanner.Text()
			sents[i] = line
			i += 1
		}

		fmt.Printf("%s | %s\n", sents[0], sents[1])
		if sents[0] == "EXIT" || sents[1] == "EXIT" {
			break
		}
		fmt.Printf("distance: %f\n", edist(sents[0], sents[1], model))
	}

}

/*
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
*/
