/*
Author: Marco Acea
Contact: aceamarco@gmail.com

Citations
--------------
https://austinmorlan.com/posts/chip8_emulator/#introduction
https://multigesture.net/articles/how-to-write-an-emulator-chip-8-interpreter/
*/

package cpu

type Chip8 struct {
	/*
		The Chip 8 has 15 8-bit general purpose registers named V0,V1 up to VE.
		The 16th register is used  for the ‘carry flag’.
	*/
	V [16]uint8

	/*
		The CHIP-8 has 4096 bytes of memory, meaning the address space is from
		0x000 to 0xFFF.
	*/
	Mem [4096]uint8

	/*
		The 16-bit Index Register is a special register used to store memory addresses for
		use in operations.
	*/
	IR uint16

	/*
		The Program Counter (PC) is a special register that holds the address of the
		next instruction to execute. Again, it’s 16 bits because it has to be
		able to hold the maximum memory address (0xFFF).
	*/
	PC uint16

	/*
		The CHIP-8 has 16 levels of stack, meaning it can hold 16 different PCs.
		Multiple levels allow for one function to call another function and so
		on, until they all return to the original caller site.
	*/
	Stack [16]uint16

	/*
		Similar to how the PC is used to keep track of where in memory the CPU
		is executing, we need a Stack Pointer (SP) to tell us where in the
		16-levels of stack our most recent value was placed (i.e, the top).
	*/
	SP uint8

	// Timers
	/*
		The CHIP-8 has a simple timer used for timing. If the timer value is
		zero, it stays zero. If it is loaded with a value, it will decrement at
		a rate of 60Hz.
	*/
	DelayTimer uint8

	/*
		The CHIP-8 also has another simple timer used for sound. Its behavior is the
		same (decrementing at 60Hz if non-zero), but a single tone will buzz when
		it’s non-zero. Programmers used this for simple sound emission.
	*/
	SoundTimer uint8

	/*
		The CHIP-8 has 16 input keys that match the first 16 hex values: 0
		through F. Each key is either pressed or not pressed.
	*/
	Keypad [16]uint8

	/*
		The CHIP-8 has an additional memory buffer used for storing the graphics
		to display. It is 64 pixels wide and 32 pixels high. Each pixel is
		either on or off, so only two colors can be represented.
	*/
	Display [64 * 32]uint32

	/*
		The Chip 8 has 35 opcodes which are all two bytes long. To store the
		current opcode, we need a data type that allows us to store two bytes.
		An unsigned short has the length of two bytes and therefor fits our
		needs.
	*/
	OpCode uint16
}
