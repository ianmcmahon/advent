package day2

import "fmt"

func process(input []int) []int {
	ip := 0
	for {
		//fmt.Printf("ip: %d, instruction: %d ", ip, input[ip])
		switch input[ip] {
		case 99:
			//fmt.Println("end")
			return input
		case 1:
			pa, pb, po := input[ip+1], input[ip+2], input[ip+3]
			input[po] = input[pa] + input[pb]
			//fmt.Printf("i[%d] %d  +  i[%d] %d  =  i[%d] %d\n", pa, input[pa], pb, input[pb], po, input[po])
			ip += 4
		case 2:
			pa, pb, po := input[ip+1], input[ip+2], input[ip+3]
			input[po] = input[pa] * input[pb]
			//fmt.Printf("i[%d] %d  *  i[%d] %d  =  i[%d] %d\n", pa, input[pa], pb, input[pb], po, input[po])
			ip += 4
		default:
			fmt.Printf("Bad opcode at instruction %d: %d\n", ip, input[ip])
			return input
		}
	}

	return input
}
