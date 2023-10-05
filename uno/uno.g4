grammar uno;
// antlr4 -Dlanguage=Go uno.g4 -package uno

start                   : boolean_expression EOF;

boolean_expression      : boolean_expression T_AND boolean_expression                                       # AndBooleanExpression
                        | boolean_expression T_OR boolean_expression                                        # OrBooleanExpression
                        | arithmetic_expression T_COMPARE arithmetic_expression                             # CmpBooleanExpression
                        | T_NOT boolean_expression                                                          # NotBooleanExpression
                        | arithmetic_expression T_IN (INTEGER_LIST|STRING_LIST|DECIMAL_LIST)                # InBooleanExpression
                        | arithmetic_expression T_NOT T_IN (INTEGER_LIST|STRING_LIST|DECIMAL_LIST)          # NotInBooleanExpression
                        | '(' boolean_expression ')'                                                        # PlainBooleanExpression
                        | T_TRUE                                                                            # TrueBooleanExpression
                        | T_FALSE                                                                           # FalseBooleanExpression
                        ;

arithmetic_expression   : arithmetic_expression T_MOD arithmetic_expression                                 # ModArithmeticExpression
                        | arithmetic_expression T_MUL arithmetic_expression                                 # MulArithmeticExpression
                        | arithmetic_expression T_DIV arithmetic_expression                                 # DivArithmeticExpression
                        | arithmetic_expression T_ADD arithmetic_expression                                 # AddArithmeticExpression
                        | arithmetic_expression T_SUB arithmetic_expression                                 # SubArithmeticExpression
                        | IDENTIFIER '(' ')'                                                                # RuntTimeFuncArithmeticExpression
                        | IDENTIFIER '(' arithmetic_expression (',' arithmetic_expression)* ')'             # FuncArithmeticExpression
                        | IDENTIFIER type_marker                                                            # ColumnArithmeticExpression
                        | IDENTIFIER '.' IDENTIFIER type_marker                                            # FieldColumnArithmeticExpression
                        | STRING                                                                            # StringArithmeticExpression
                        | INTEGER                                                                           # IntegerArithmeticExpression
                        | DECIMAL                                                                           # DecimalArithmeticExpression
                        | '(' arithmetic_expression ')'                                                     # PlainArithmeticExpression
                        ;

type_marker             : '[' (T_INT|T_FLOAT|T_STRING) ']' ;

T_ADD                   : '+' ;
T_SUB                   : '-' ;
T_MUL                   : '*' ;
T_DIV                   : '/' ;
T_MOD                   : '%' ;

// reserved keywords
T_INT                   : I N T '6' '4' ;
T_FLOAT                 : F L O A T '3' '2' ;
T_STRING                : S T R I N G ;
T_ON                    : O N ;
T_AND                   : A N D ;
T_OR                    : O R ;
T_NOT                   : N O T ;
T_IN                    : I N ;
T_TRUE                  : T R U E ;
T_FALSE                 : F A L S E ;


// Support case-insensitive keywords and allowing case-sensitive identifiers
fragment A              : ('a'|'A') ;
fragment B              : ('b'|'B') ;
fragment C              : ('c'|'C') ;
fragment D              : ('d'|'D') ;
fragment E              : ('e'|'E') ;
fragment F              : ('f'|'F') ;
fragment G              : ('g'|'G') ;
fragment H              : ('h'|'H') ;
fragment I              : ('i'|'I') ;
fragment J              : ('j'|'J') ;
fragment K              : ('k'|'K') ;
fragment L              : ('l'|'L') ;
fragment M              : ('m'|'M') ;
fragment N              : ('n'|'N') ;
fragment O              : ('o'|'O') ;
fragment P              : ('p'|'P') ;
fragment Q              : ('q'|'Q') ;
fragment R              : ('r'|'R') ;
fragment S              : ('s'|'S') ;
fragment T              : ('t'|'T') ;
fragment U              : ('u'|'U') ;
fragment V              : ('v'|'V') ;
fragment W              : ('w'|'W') ;
fragment X              : ('x'|'X') ;
fragment Y              : ('y'|'Y') ;
fragment Z              : ('z'|'Z') ;

// Comparison marks
T_COMPARE               : T_EQUAL
                        | T_EQUAL2
                        | T_NOTEQUAL
                        | T_NOTEQUAL2
                        | T_GREATER
                        | T_GREATEREQUAL
                        | T_LESS
                        | T_LESSEQUAL
                        ;

T_EQUAL                 : '=' ;
T_EQUAL2                : '==' ;
T_NOTEQUAL              : '<>' ;
T_NOTEQUAL2             : '!=' ;
T_GREATER               : '>' ;
T_GREATEREQUAL          : '>=' ;
T_LESS                  : '<' ;
T_LESSEQUAL             : '<=' ;

IDENTIFIER              : [_a-zA-Z][_a-zA-Z0-9]* ;   

INTEGER_LIST            : '(' INTEGER (',' INTEGER)* ')' ;
INTEGER                 : '-'? '0' | [1-9] [0-9]* ;

DECIMAL_LIST            : '(' DECIMAL (',' DECIMAL)* ')' ;
DECIMAL                 : '-'? ('0' | [1-9] [0-9]*) '.' [0-9]* ;

STRING_LIST             : '(' STRING (',' STRING)* ')' ;
STRING                  : '"' (ESC | SAFECODEPOINT)* '"' ;

fragment ESC            : '\\' (["\\/bfnrt] | UNICODE) ;
fragment UNICODE        : 'u' HEX HEX HEX HEX ;
fragment HEX            : [0-9a-fA-F] ;
fragment SAFECODEPOINT  : ~ ["\\\u0000-\u001F] ;

// Ignore whitespace
WS                      : [ \t\n\r] + -> skip ;