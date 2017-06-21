grammar Brainfuck;

TAPE_INC   : '+';
TAPE_DEC   : '-';
TAPE_LEFT  : '<';
TAPE_RIGHT : '>';
INPUT      : ',';
OUTPUT     : '.';

program  : command*;
command  : loop
         | op
         ;
loop     : '[' command* ']';
op       : cmd=( TAPE_INC
               | TAPE_DEC
               | TAPE_LEFT
               | TAPE_RIGHT
               | INPUT
               | OUTPUT
               )
               ;

END : ~[EOF] -> skip;