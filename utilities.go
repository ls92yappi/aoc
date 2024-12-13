package aoc

import (
	"fmt"
	"io/ioutil"
	"log"
	"os"
	"strconv"
	"strings"
)

////////////////////////////////////////////////////////////////////
//                                                                //
//   Many of these functions below operate on or return           //
//   int slices. None have any import dependencies.               //
//                                                                //
////////////////////////////////////////////////////////////////////
//                                                                //
// func die(msg interface{})                                      //
// func If[T any](cond bool, vtrue, vfalse T) T                   //
// func InputFileName() string                                    //
// func ReadWholeFile(fname string) ([]string, int, error)        //
// func IntSlice(s string, delim string) (ints []int, int, error) //
//                                                                //
////////////////////////////////////////////////////////////////////

// If returns vtrue if cond is true, vfalse otherwise.
//
// Useful to avoid an if statement when initializing variables, for example:
//
//	min := If(i > 0, i, 0)
// If -- the best Ternary operator ever!!! from https://github.com/icza/gox/blob/main/gox/gox.go
func If[T any](cond bool, vtrue, vfalse T) T {
	if cond {
		return vtrue
	}
	return vfalse
}

// msg is interface{} because cannot convert error to string
func die(msg interface{}) {
	log.Println(msg)
	os.Exit(1)
}

func InputFileName() string {
	var argc int = len(os.Args)
	var argv []string = os.Args

	if argc < 2 {
		fmt.Printf("Usage: %s [inputfile]\n", argv[0])
		os.Exit(1)
	}

	inputFileName := argv[1]
	return inputFileName
}

func ReadWholeFile(fname string) (lines []string, numLines int, err error) {
	emptySlice := make([]string, 0)
	// open file
	file, err := os.Open(fname)
	if err != nil { die(err) }
	defer file.Close()

	// read the whole file in
	srcbuf, err := ioutil.ReadAll(file)
	if err != nil { return emptySlice, 0, err }
	src := string(srcbuf)

	lines = strings.Split(src, "\n")
	numLines = len(lines)
	return lines, numLines, nil
}

func IntSlice(s string, delim string) (ints []int, numItems int, err error) {
	emptySlice := make([]int, 0)
	sl := strings.Split(s, delim)
	numItems = len(sl)
	ints = make([]int, numItems)
	for i := range(numItems) {
		val, err := strconv.Atoi(sl[i])
		if err != nil {
			return emptySlice, 0, err
		}
		ints[i] = val
	}
	return ints, numItems, nil
}
