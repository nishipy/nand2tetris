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
//reference: https://github.com/indragiek/tecs/blob/master/Project4/Fill.asm

@8192   // Hackでは、横512 * 縦256 = 131072ピクセルのスクリーン
        // 131072 / 16 = 8192
D=A     // D=8192
@count
M=D     // count = 8192 (# of bytes)

(LOOP)
    @index
    M=0     // index = 0
(INNER)
    @KBD    // KBD: RAMアドレスの24576(0x6000)を参照するように定義されている
            // キーボードを入力するとRAM[24576]にASCIIコードが現れ、何も入力していないと0
    D=M     // D=KBD     
    @WHITE
    D;JEQ   // KBDの値が0(i.e.何も入力していない)ならば、WITEへ
(BLACK)
    @index
    D=M     //D=index
    @SCREEN // SCREEN: RAMアドレスの16384(0x4000)を参照するように定義されている
    A=A+D
    M=-1    // 2の補数: 16bitで-1を表す場合、"11......1"(16bit)となる
            // つまり全ビット1となるので、黒く塗ることになる
    @END
    0;JMP   // ENDへ
(WHITE)
    @index
    D=M     // D=index
    @SCREEN
    A=A+D   // バイトアドレスの計算
    M=0     // 白く塗る
(END)   
    @index
    MD=M+1  // indexをインクリメント
    @count  // indexが8192ならばLOOPへ行き、indexを初期化
            // それ以外は、INNERへ
    D=D-M
    @LOOP
    D;JEQ
    @INNER
    0;JMP