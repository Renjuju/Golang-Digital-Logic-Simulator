# Golang-Digital-Logic-Simulator
For your final programming project you will implement a digital circuit simulator along the lines of what you designed for Exercise 8-4.  Your simulator should include support for the basic logic gates: NOT, AND, OR, NAND, NOR, and XOR.  You should also have a clock that generates a stream of pulses at some specified rate.  The final component you need to support is a D flip-flop.  The D flip-flop has two inputs and two outputs and operates as follows: when the clock input transitions from 0 to 1, the value on the D input is stored.  The Q output is the stored value of D and the Q with bar on top is the inverse of the stored D value.

Your implementation should be written in go.  Each component (gate, clock, or flip-flop) will have an associated go routine, and the go routines will communicate using channels.  Your program should as the user for the following:

Name of a file containing the description of the circuit to simulate
Initial state values for flip-flops and for any external input signals
If a clock is used:
Number of clock pulses per second
Number of clock pulses to run the simulation
In addition to your source code, you should turn in a write-up that includes at least the following:

Description of your simulation algorithm with particular attention to how you know that all the propagation effects have moved through the circuit and it's safe for another clock pulse and how you handle the potential that pulses might occur simultaneously
Description of the file format you design for your circuits.
Example circuits for a 4-bit adder (with external inputs) and an 8-bit counter (driven from the clock).
Results of running your program on those two test circuits.
