// antlr4 -Dlanguage=Go uno.g4 -package uno
grammar uno;
options { caseInsensitive=true; }
BRACKET_OPEN        : '(';
BRACKET_CLOSE       : ')';
SQUARE_OPEN         : '[';
SQUARE_CLOSE        : ']';
DOT                 : '.';
COMMA: ',';
QUOTA: '"';
T_ADD                   : '+' ;
T_SUB                   : '-' ;
T_MUL                   : '*' ;
T_DIV                   : '/' ;
T_MOD                   : '%' ;


start                   : boolean_expression EOF;

boolean_expression      : boolean_expression T_AND boolean_expression                                       # AndBooleanExpression
                        | boolean_expression T_OR boolean_expression                                        # OrBooleanExpression
                        | arithmetic_expression T_COMPARE arithmetic_expression                             # CmpBooleanExpression
                        | T_NOT boolean_expression                                                          # NotBooleanExpression
                        | arithmetic_expression T_IN (INTEGER_LIST|STRING_LIST|DECIMAL_LIST)                # InBooleanExpression
                        | arithmetic_expression T_NOT T_IN (INTEGER_LIST|STRING_LIST|DECIMAL_LIST)          # NotInBooleanExpression
                        | BRACKET_OPEN boolean_expression BRACKET_CLOSE                                                        # PlainBooleanExpression
                        | T_TRUE                                                                            # TrueBooleanExpression
                        | T_FALSE                                                                           # FalseBooleanExpression
                        ;

arithmetic_expression   : arithmetic_expression T_MOD arithmetic_expression                                 # ModArithmeticExpression
                        | arithmetic_expression T_MUL arithmetic_expression                                 # MulArithmeticExpression
                        | arithmetic_expression T_DIV arithmetic_expression                                 # DivArithmeticExpression
                        | arithmetic_expression T_ADD arithmetic_expression                                 # AddArithmeticExpression
                        | arithmetic_expression T_SUB arithmetic_expression                                 # SubArithmeticExpression
                        | IDENTIFIER BRACKET_OPEN BRACKET_CLOSE                                                                # RuntTimeFuncArithmeticExpression
                        | IDENTIFIER BRACKET_OPEN arithmetic_expression (COMMA arithmetic_expression)* BRACKET_CLOSE             # FuncArithmeticExpression
                        | IDENTIFIER type_marker                                                            # ColumnArithmeticExpression
                        | IDENTIFIER DOT IDENTIFIER type_marker                                            # FieldColumnArithmeticExpression
                        | STRING                                                                            # StringArithmeticExpression
                        | INTEGER                                                                           # IntegerArithmeticExpression
                        | DECIMAL                                                                           # DecimalArithmeticExpression
                        | BRACKET_OPEN arithmetic_expression BRACKET_CLOSE                                                     # PlainArithmeticExpression
                        ;

type_marker             : SQUARE_OPEN (T_INT|T_FLOAT|T_STRING|T_INTS|T_FLOATS|T_STRINGS) SQUARE_CLOSE ;



// reserved keywords
T_INT                   : 'int64';
T_INTS                  : 'int64s' ;
T_FLOAT                 : 'float32';
T_FLOATS                : 'float32s';
T_STRING                : 'string';
T_STRINGS               : 'strings';
T_ON                    : 'ON' ;
T_AND                   : 'and' ;
T_OR                    : 'or' ;
T_NOT                   : 'not';
T_IN                    : 'in' ;
T_TRUE                  : 'true' ;
T_FALSE                 : 'false' ;


// Support case-insensitive keywords and allowing case-sensitive identifiers    

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

IDENTIFIER              options { caseInsensitive=false; } : [_a-zA-Z][_a-zA-Z0-9]*; // 变量

INTEGER_LIST            : BRACKET_OPEN (INTEGER COMMA)* INTEGER BRACKET_CLOSE ;
INTEGER                 : '-'? '0' | [1-9] [0-9]* ;

DECIMAL_LIST            : BRACKET_OPEN (DECIMAL COMMA)* DECIMAL BRACKET_CLOSE ;
DECIMAL                 : '-'? ('0' | [1-9] [0-9]*) '.' [0-9]* ;

STRING_LIST             : BRACKET_OPEN (STRING COMMA)* STRING BRACKET_CLOSE ;

STRING options { caseInsensitive=false; } : QUOTA (ESC | SAFECODEPOINT)* QUOTA;

fragment ESC: '\\' (["\\/bfnrt] | UNICODE);
fragment UNICODE: 'u' HEX HEX HEX HEX;
fragment HEX options { caseInsensitive=false; }  : [0-9a-fA-F];
fragment SAFECODEPOINT: ~ ["\\\u0000-\u001F];

// Ignore whitespace
WS                      : [ \t\n\r] + -> skip ;