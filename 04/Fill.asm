// This file is part of www.nand2tetris.org
// and the book "The Elements of Computing Systems"
// by Nisan and Schocken, MIT Press.
// File name: projects/04/Fill.asm

// Runs an infinite loop that listens to the keyboard input.
// When a key is pressed (any key), the program blackens the screen,
// i.e. writes "black" in every pixel;
// the screen should remain fully black as long as the key is pressed. 
// When no key is pressed, the program clears the screen, i.e. writes
// "white" in every pixel;
// the screen should remain fully clear as long as no key is pressed.

// Put your code here.

(LOOP)
  @SCREEN
  D=A
  @addr
  M=D           //reset addr variable to value of SCREEN address which is the first register 
                //representing the screen and is at 16384 in Hack CPU
  @24576
  D=A
  @screenend
  M=D      //end of screen is represented by register at constant SCREEN + 8191 + 1

  @KBD
  D=M
  @WHITE
  D;JEQ         // goto WHITE if KBD value is 0

  (BLACK)       //paint screen black
    @addr
    A=M         //set address register to what is in addr variable
    M=-1        //Set the register pointed to by A register to 1111 1111 1111 1111

    @addr
    M=M+1       //increment addr variable
    D=M         //store new value in data register

    @screenend
    D=M-D       //set data register to value of screenend-addr

    @BLACK
    D;JNE       //if value is greater than 0 then whole screen is not painted to continue painting

    @LOOP
    0;JEQ       //otherwise jmp to beginning of loop

  (WHITE)       //paint screen white
    @addr
    A=M         //set address register to what is in addr variable
    M=0         //Set the register pointed to by A register to 0000 0000 0000 0000

    @addr
    M=M+1       //increment addr variable
    D=M         //store new value in data register

    @screenend
    D=M-D       //set data register to value of screenend-addr

    @WHITE
    D;JNE       //if value is greater than 0 then whole screen is not painted to continue painting

    @LOOP
    0;JEQ       //otherwise jmp to beginning of loop