package main

// virtual 4 bit microprocessor

import (
	"fmt"
	"os"
)

// Instruction Loop:
// The instruction at the address given by the Instruction Pointer is loaded into the Instruction Store
// The Instruction Pointer is incremented by 1 or 2 depending on whether the instruction is a 1 or 2 byte instruction
// The instruction in the Instruction Store is executed
// This is repeated until the instruction in the Instruction Store equals 0 (Halt)

type RegisterBank struct {
	IP int // instruction pointer
	IS int // instruction store
	R0 int // register zero
	R1 int // register one
}

type Program struct {
	location map[int]int
	// location = make(map[int]int)
}
type MicroProcessor struct {
	registers RegisterBank
	memory    [16]int
}

func (mp *MicroProcessor) executeInstruction() {
	switch mp.registers.IS {
	case 0:
		fmt.Println("Halting now..")
		os.Exit(0)
	case 1:
		fmt.Println("Adding")
		mp.registers.R0 = mp.registers.R0 + mp.registers.R1
	case 2:
		fmt.Println("Subtracting now..")
		mp.registers.R0 = mp.registers.R0 - mp.registers.R1
	case 3:
		fmt.Println("Incrementing R0..")
		mp.registers.R0 = mp.registers.R0 + 1
	case 4:
		fmt.Println("Incrementing R1..")
		mp.registers.R0 = mp.registers.R1 + 1
	case 5:
		fmt.Println("decrementing R0..")
		mp.registers.R0 = mp.registers.R0 - 1
	case 6:
		fmt.Println("decrementing R1..")
		mp.registers.R1 = mp.registers.R1 - 1
	case 7:
		fmt.Println("**Ring Bell**")
	case 8:
		fmt.Println("Print Data..")
		fmt.Println(mp.memory[mp.registers.IP-1])
		fmt.Println("Print R0..")
		fmt.Println(mp.registers.R0)
	case 9:
		fmt.Println("Load Address (val", mp.memory[mp.registers.IP-1], "] to R0")
		mp.registers.R0 = mp.memory[mp.registers.IP-1]
	case 10:
		fmt.Println("Load Address (val", mp.memory[mp.registers.IP-1], "] to R1")
		mp.registers.R1 = mp.memory[mp.registers.IP-1]
	case 11:
		fmt.Println("Store R0 to Address")
		mp.memory[mp.registers.IP-1] = mp.registers.R0
	case 12:
		fmt.Println("Store R1 to Address ")
		mp.memory[mp.registers.IP-1] = mp.registers.R1
	case 13:
		fmt.Println("Jump To Address..")
		mp.registers.IP = mp.memory[mp.registers.IP-1]
	case 14:
		fmt.Println("Jump To Address if R0..")
		if mp.registers.R0 == 1 {
			mp.registers.IP = mp.memory[mp.registers.IP-1]
		}
	case 15:
		fmt.Println("Jump To Address if Not R0..")
		if mp.registers.R0 != 1 {
			mp.registers.IP = mp.memory[mp.registers.IP-1]
		}
	}
}

func (mp *MicroProcessor) dumpMemory() {
	for v := range mp.memory {
		fmt.Println(v, " Val: ", mp.memory[v])
	}
}

func (mp *MicroProcessor) fetchExecuteLoop() {
	for {
		fmt.Println("IP:", mp.registers.IP, " // IS:", mp.registers.IS)
		mp.registers.IS = mp.memory[mp.registers.IP]

		if mp.registers.IS > 8 {
			mp.registers.IP += 2
		} else {
			mp.registers.IP += 1
		}
		mp.executeInstruction()
	}
}

func (mp *MicroProcessor) loadProgram(program [16]int) {
	for k, v := range program {
		fmt.Println(k, v)
		mp.memory[k] = v
	}
}

func main() {
	mp := new(MicroProcessor)
	proggy := [16]int{9, 4, 10, 3, 1, 8}
	mp.loadProgram(proggy)
	mp.fetchExecuteLoop()

}
