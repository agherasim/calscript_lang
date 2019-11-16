lexer grammar CommonLexer;

A: [Aa];
B: [Bb];
C: [Cc];
D: [Dd];
E: [Ee];
F: [Ff];
G: [Gg];
H: [Hh];
I: [Ii];
J: [Jj];
K: [Kk];
L: [Ll];
M: [Mm];
N: [Nn];
O: [Oo];
P: [Pp];
Q: [Qq];
R: [Rr];
S: [Ss];
T: [Tt];
U: [Uu];
V: [Vv];
W: [Ww];
X: [Xx];
Y: [Yy];
Z: [Zz];

COMMA: ',';
COLON: ':';
SLASH: '/';
DASH: '-';
DOT: '.';

// Complex types
IDENTIFIER: LETTER (LETTER | DIGIT | '_')+; 
STRING: '"' ( '\\"' | . )*? '"';

// Base types
UNSIGNED_INT: DIGIT+; 
SIGNED_INT: SIGN DIGIT+;
FLOAT: SIGN? DIGIT+ ('.' DIGIT+);

fragment LETTER: UCLETTER | LCLETTER;
fragment LCLETTER: [a-z];
fragment UCLETTER: [A-Z];
fragment DIGIT: [0-9];
fragment SIGN: '+' | '-';