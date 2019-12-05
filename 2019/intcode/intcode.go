package intcode

import "fmt"

type ParamMode int

const (
	POSITION  ParamMode = 0
	IMMEDIATE           = 1
)

// gets the param at ip, fetched based on mode
func getParam(input []int, ip int, mode ParamMode) int {
	switch mode {
	case POSITION:
		return input[input[ip]]
	case IMMEDIATE:
		return input[ip]
	default:
		panic(fmt.Errorf("Unknown parameter mode: %d", mode))
	}
	return 0
}

// splits the parameter modes value into A and B input modes
// turns out the "output" digit is the leading 0 of the opcode, badly documented
func iioParamModes(modes int) (a, b ParamMode, err error) {
	a = ParamMode(modes % 10)
	modes /= 10
	b = ParamMode(modes % 10)
	modes /= 10
	if modes == IMMEDIATE {
		err = fmt.Errorf("modes %d, 3rd param (output) should never be immediate", modes)
	}
	return
}

func Process(input []int, userInput ...int) (memory, output []int) {
	output = []int{}
	ip := 0
	for {
		//fmt.Printf("ip: %d, instruction: %d ", ip, input[ip])

		instruction := input[ip]
		opcode := instruction % 100
		paramModes := instruction / 100

		switch opcode {
		case 99:
			// HALT
			return input, output

		case 1:
			// ADD
			am, bm, err := iioParamModes(paramModes)
			if err != nil {
				fmt.Printf("ip: %d instruction: %d err: %v\n", ip, instruction, err)
			}
			// getparam dereferences the pointers if appropriate
			a, b := getParam(input, ip+1, am), getParam(input, ip+2, bm)
			// output is always a pointer
			po := input[ip+3]

			input[po] = a + b
			ip += 4

		case 2:
			// MUL
			am, bm, err := iioParamModes(paramModes)
			if err != nil {
				fmt.Printf("ip: %d instruction: %d err: %v\n", ip, instruction, err)
			}
			// getparam dereferences the pointers if appropriate
			a, b := getParam(input, ip+1, am), getParam(input, ip+2, bm)
			// output is always a pointer
			po := input[ip+3]

			input[po] = a * b
			ip += 4

		case 3:
			// INPUT
			// writes to the single param, so never immediate mode
			pa := input[ip+1]
			// shift input off from userInput array
			if len(userInput) == 0 {
				fmt.Printf("out of user input! submitting 0\n")
				input[pa] = 0
			} else {
				input[pa], userInput = userInput[0], userInput[1:]
			}
			ip += 2

		case 4:
			// OUTPUT
			// single param can be position or immediate
			a := getParam(input, ip+1, ParamMode(paramModes%10)) // mode shouldn't ever be > 10 but :)
			//fmt.Printf("OUTPUT: %d\n", a)
			output = append(output, a)
			ip += 2

		case 5:
			// JNZ
			am, bm, err := iioParamModes(paramModes)
			if err != nil {
				fmt.Printf("ip: %d instruction: %d err: %v\n", ip, instruction, err)
			}
			// getparam dereferences the pointers if appropriate
			a, b := getParam(input, ip+1, am), getParam(input, ip+2, bm)
			if a != 0 {
				ip = b
			} else {
				ip += 3
			}

		case 6:
			// JZ
			am, bm, err := iioParamModes(paramModes)
			if err != nil {
				fmt.Printf("ip: %d instruction: %d err: %v\n", ip, instruction, err)
			}
			// getparam dereferences the pointers if appropriate
			a, b := getParam(input, ip+1, am), getParam(input, ip+2, bm)
			if a == 0 {
				ip = b
			} else {
				ip += 3
			}

		case 7:
			// LT    C = (A < B ? 1 : 0)
			am, bm, err := iioParamModes(paramModes)
			if err != nil {
				fmt.Printf("ip: %d instruction: %d err: %v\n", ip, instruction, err)
			}
			// getparam dereferences the pointers if appropriate
			a, b := getParam(input, ip+1, am), getParam(input, ip+2, bm)
			// output is always a pointer
			po := input[ip+3]
			if a < b {
				input[po] = 1
			} else {
				input[po] = 0
			}
			ip += 4

		case 8:
			// EQ    C = (A == B ? 1 : 0)
			am, bm, err := iioParamModes(paramModes)
			if err != nil {
				fmt.Printf("ip: %d instruction: %d err: %v\n", ip, instruction, err)
			}
			// getparam dereferences the pointers if appropriate
			a, b := getParam(input, ip+1, am), getParam(input, ip+2, bm)
			// output is always a pointer
			po := input[ip+3]
			if a == b {
				input[po] = 1
			} else {
				input[po] = 0
			}
			ip += 4

		default:
			fmt.Printf("Bad opcode at instruction %d: %d\n", ip, input[ip])
			return input, output
		}
	}

	return input, output
}
