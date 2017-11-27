// This file is part of www.nand2tetris.org
// and the book "The Elements of Computing Systems"
// by Nisan and Schocken, MIT Press.
// File name: projects/04/Mult.asm

// Multiplies R0 and R1 and stores the result in R2.
// (R0, R1, R2 refer to RAM[0], RAM[1], and RAM[2], respectively.)

// Put your code here.

@R0
D=M
@multiplier
M=D     //multiplier = R0

@R1
D=M
@multiplicand
M=D     //multiplicand = R1

@R2
M=0    //set RAM[2] = 0

@counter
M=1     //counter = 1

@product
M=0     //product = 0

(LOOP)
  @counter
  D=M
  @multiplier
  D=D-M
  @STOP
  D;JGT //if counter > multiplier goto STOP


  @multiplicand
  D=M
  @product
  M=M+D   //product += multiplicand

  @counter
  M=M+1   //counter += 1

  @LOOP
  0;JMP

(STOP)
  @product
  D=M
  @R2
  M=D   //RAM[2] = product

(END)
  @END
  0;JMP
