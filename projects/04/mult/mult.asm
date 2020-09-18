// This file is part of www.nand2tetris.org
// and the book "The Elements of Computing Systems"
// by Nisan and Schocken, MIT Press.
// File name: projects/04/Mult.asm

// Multiplies R0 and R1 and stores the result in R2.
// (R0, R1, R2 refer to RAM[0], RAM[1], and RAM[2], respectively.)

// Put your code here.
//答えの初期化
@2
M=0

//R0==0 || R1==0 ならば、@END にジャンプする
@0
D=M
@END
D;JEQ

@1
D=M
@END
D;JEQ

(LOOP)
    @2
    D=M

    @1
    D=D+M

    //RAM[2]を書き換える
    @2
    M=D

    //RAM[0]の値を1減らす
    @0
    D=M-1
    M=D

    //RAM[0]>0である限りループする
    @LOOP
    D;JGT


(END)
    @END
    0;JMP //これは、無限ループを表す
