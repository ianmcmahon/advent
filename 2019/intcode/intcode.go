package intcode

import (
	"fmt"
	"io"
	"os"
	"sync"
)

type ParamMode int

const (
	POSITION  ParamMode = 0
	IMMEDIATE           = 1
	RELATIVE            = 2
)

// gets the param at ip, fetched based on mode
func getParam(memory []int, ip int, rb int, mode ParamMode) int {
	switch mode {
	case POSITION:
		return memory[memory[ip]]
	case IMMEDIATE:
		return memory[ip]
	case RELATIVE:
		return memory[rb+memory[ip]]
	default:
		panic(fmt.Errorf("Unknown parameter mode: %d", mode))
	}
	return 0
}

// splits the parameter modes value into A and B input modes
// turns out the "output" digit is the leading 0 of the opcode, badly documented
func iioParamModes(modes int) (a, b, o ParamMode, err error) {
	a = ParamMode(modes % 10)
	modes /= 10
	b = ParamMode(modes % 10)
	o = ParamMode(modes / 10)
	if o == IMMEDIATE {
		err = fmt.Errorf("modes %d, 3rd param (output) should never be immediate", o)
	}
	return
}

func Process(program []int, userInput ...int) (memory, output []int) {
	inCh := make(chan int, 0)
	outCh := make(chan int, 0)

	wg := sync.WaitGroup{}

	go func() {
		wg.Add(1)
		for {
			o, more := <-outCh
			if !more {
				break
			}
			output = append(output, o)
		}
		wg.Done()
	}()

	go func() {
		for _, v := range userInput {
			inCh <- v
		}
	}()

	memory = Run(program, inCh, outCh, os.Stdout)

	wg.Wait()

	return memory, output
}

func Run(program []int, input <-chan int, output chan<- int, console io.Writer) (memory []int) {
	memory = make([]int, 1024*1024)
	copy(memory, program)
	defer close(output)

	ip := 0
	rb := 0
	for {
		//fmt.Printf("ip: %d, instruction: %d \n", ip, memory[ip])

		instruction := memory[ip]
		opcode := instruction % 100
		paramModes := instruction / 100

		switch opcode {
		case 99:
			// HALT
			//fmt.Fprintf(console, "HALT\n")
			return memory

		case 1:
			// ADD
			am, bm, om, err := iioParamModes(paramModes)
			if err != nil {
				fmt.Printf("ip: %d instruction: %d err: %v\n", ip, instruction, err)
			}
			// getparam dereferences the pointers if appropriate
			a, b := getParam(memory, ip+1, rb, am), getParam(memory, ip+2, rb, bm)
			// output is always a pointer
			po := memory[ip+3]
			// but sometimes relative
			if om == RELATIVE {
				po += rb
			}

			memory[po] = a + b
			ip += 4

		case 2:
			// MUL
			am, bm, om, err := iioParamModes(paramModes)
			if err != nil {
				fmt.Printf("ip: %d instruction: %d err: %v\n", ip, instruction, err)
			}
			// getparam dereferences the pointers if appropriate
			a, b := getParam(memory, ip+1, rb, am), getParam(memory, ip+2, rb, bm)
			// output is always a pointer
			po := memory[ip+3]
			// but sometimes relative
			if om == RELATIVE {
				po += rb
			}

			memory[po] = a * b
			ip += 4

		case 3:
			// INPUT
			// this needs to handle relative mode
			pa := memory[ip+1]
			if paramModes%10 == RELATIVE {
				fmt.Printf("the dreaded 203: rb %d + %d\n", rb, pa)
				pa += rb
			}
			// block on input channel until we have input
			//fmt.Fprintf(console, "blocking on input\n")
			memory[pa] = <-input
			fmt.Fprintf(console, "stored at %d input %d\n", pa, memory[pa])
			ip += 2

		case 4:
			// OUTPUT
			a := getParam(memory, ip+1, rb, ParamMode(paramModes%10))
			//fmt.Fprintf(console, "blocking on output\n")
			output <- a
			//fmt.Fprintf(console, "sent output %d\n", a)
			ip += 2

		case 5:
			// JNZ
			am, bm, _, _ := iioParamModes(paramModes)
			// getparam dereferences the pointers if appropriate
			a, b := getParam(memory, ip+1, rb, am), getParam(memory, ip+2, rb, bm)
			if a != 0 {
				ip = b
			} else {
				ip += 3
			}

		case 6:
			// JZ
			am, bm, _, _ := iioParamModes(paramModes)
			// getparam dereferences the pointers if appropriate
			a, b := getParam(memory, ip+1, rb, am), getParam(memory, ip+2, rb, bm)
			if a == 0 {
				ip = b
			} else {
				ip += 3
			}

		case 7:
			// LT    C = (A < B ? 1 : 0)
			am, bm, om, err := iioParamModes(paramModes)
			if err != nil {
				fmt.Printf("ip: %d instruction: %d err: %v\n", ip, instruction, err)
			}
			// getparam dereferences the pointers if appropriate
			a, b := getParam(memory, ip+1, rb, am), getParam(memory, ip+2, rb, bm)
			// output is always a pointer
			po := memory[ip+3]
			// but sometimes relative
			if om == RELATIVE {
				po += rb
			}
			if a < b {
				memory[po] = 1
			} else {
				memory[po] = 0
			}
			ip += 4

		case 8:
			// EQ    C = (A == B ? 1 : 0)
			am, bm, om, err := iioParamModes(paramModes)
			if err != nil {
				fmt.Printf("ip: %d instruction: %d err: %v\n", ip, instruction, err)
			}
			// getparam dereferences the pointers if appropriate
			a, b := getParam(memory, ip+1, rb, am), getParam(memory, ip+2, rb, bm)
			// output is always a pointer
			po := memory[ip+3]
			// but sometimes relative
			if om == RELATIVE {
				po += rb
			}
			if a == b {
				memory[po] = 1
			} else {
				memory[po] = 0
			}
			ip += 4

		case 9:
			// RB OFFSET
			a := getParam(memory, ip+1, rb, ParamMode(paramModes%10))
			rb += a
			ip += 2

		default:
			fmt.Printf("Bad opcode at instruction %d: %d\n", ip, memory[ip])
			return memory
		}
	}

	return memory
}
