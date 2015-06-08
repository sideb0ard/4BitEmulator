4Bit MicroProcessor emulator based on Richard Buckland lectures:
https://www.youtube.com/watch?v=gTeDX4yAdyU&list=PL3E12E78D91CA229E

The 4917 Microprocessor
(Hypothetical processor designed for course). It is 4bit, has 16 memory locations and 4 registers (instruction pointer, instruction store, general register 0 and general register 1). Each memory location could store a number between 0 and 15, and there were 16 instructions.

Instruction Set
=======================

1-byte Instructions
-------------------

0 = Halt

1 = Add (R0 = R0 + R1)

2 = Subtract (R0 = R0 – R1)

3 = Increment R0 (R0 = R0 + 1)

4 = Increment R1 (R1 = R1 + 1)

5 = Decrement R0 (R0 = R0 – 1)

6 = Decrement R1 (R1 = R1 – 1)

7 = Ring Bell

2-byte Instructions, value of the second byte is called \<data\>
--------------------------------------------------------------

8 = Print <data> (The numerical value of <data> is printed)

9 = Load value at address <data> into R0

10 = Load value at address <data> into R1

11 = Store R0 into address <data>

12 = Store R1 into address <data>

13 = Jump to address <data>

14 = Jump to address <data> if R0 == 0

15 = Jump to address <data> if R0 != 0


Startup
================
It had a simple start up (registers set to 0, all memory locations set to 0, fetch-execute cycle begins), and a fetch-execute cyle,

EvalLoop
==================

* The instruction at the address given by the instruction pointer is loaded into the instruction store.
* The instruction pointer is incremented by 1 or 2 depending on whether the instruction store is a 1 or 2 byte instruction.
* The instruction in the instruction store is executed.
* This is repeated until the instruction store equals 0 (HALT)

