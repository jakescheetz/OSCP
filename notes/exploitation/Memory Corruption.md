# Buffer Overrflows
> Tips and tricks to succesfully exploit buffer overflows


- Goal is to aim to write aribitrary code into the EIP register as this is the CPU register that stores the address of the next instruction to be executed
	- if we can overwite this address we can direct the program to execute code in any location we are able to overflow

- the ESP register stores the address of the current stack frame
	- basically is what controls the scope of local variables

- most often associated with a memory operations in code like copying a string or moving data that accidentally overwrites other existing data in an unintended area of memory



## Windows Buffer Overflows
- main mode of discovery is fuzzing
	- looking for applicatino crashes
- generally want to fuzz any input field the application offered

- Have to know exactly which part of the buffer is landing in the EIP register
	- 2 Methods:
		- Binary Tree Analysis ->  

#### Buffer Overflow Protection Mechanisms on Windows
- DEP -> Data Execution Protection
	- hardware/software that perform extra memory access checks to make sure code is not execution from data pages
- ASLR -> Address Space Layout Randomization
	- randomizes the base address of a program and it's DLL's every time it is loaded so that exploitation of a statis memory address is not feasible
	- This is not implemented on WinXP
- CFG -> Control-Flow Guard 
	- prevents code from branching and overwriting memory address pointers


