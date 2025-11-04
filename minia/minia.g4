grammar minia;

// alias antlr='java -jar $PWD/antlr-4.13.2-complete.jar'
// antlr -Dlanguage=Go minia.g4 -package minia

options { caseInsensitive=true; }

// Program structure
prog        : (start SEMI)* start EOF;

// Assignment statement
start       : IDENTIFIER T_EQUAL expr;

// Expression hierarchy (lowest to highest precedence)
expr        : logicalOrExpr;

logicalOrExpr   : logicalOrExpr T_OR logicalOrExpr # OrExpr 
                | logicalAndExpr                   # TrivialLogicalAndExpr
                ;

logicalAndExpr  : logicalAndExpr T_AND logicalAndExpr # AndExpr
                | equalityExpr                        # TrivialEqualityExpr
                ;

equalityExpr    : relationalExpr T_EQ relationalExpr    # EqualExpr
                | relationalExpr T_NEQ relationalExpr   # NotEqualExpr
                | relationalExpr                        # TrivialRelationalExpr
                ;

relationalExpr  : additiveExpr T_GT additiveExpr        # GreaterThanExpr
                | additiveExpr T_GTE additiveExpr       # GreaterThanEqualExpr      
                | additiveExpr T_LT additiveExpr        # LessThanExpr
                | additiveExpr T_LTE additiveExpr       # LessThanEqualExpr
                | additiveExpr                          # TrivialAdditiveExpr
                ;

additiveExpr    : multiplicativeExpr T_ADD multiplicativeExpr # AddExpr
                | multiplicativeExpr T_SUB multiplicativeExpr # SubExpr
                | multiplicativeExpr                          # TrivialMultiplicativeExpr
                ;

multiplicativeExpr : unaryExpr T_MUL unaryExpr # MulExpr
                   | unaryExpr T_DIV unaryExpr # DivExpr
                   | unaryExpr T_MOD unaryExpr # ModExpr
                   | unaryExpr                 # TrivialUnaryExpr
                   ;

unaryExpr       : T_NOT unaryExpr # NotExpr 
                | T_SUB unaryExpr # NegExpr
                | primaryExpr     # TrivialPrimaryExpr
                ;

primaryExpr     : '(' expr ')'              # ParenthesizedExpr
                | funcCall                  # FunctionCallExpr
                | IDENTIFIER                # ColumnExpr
                | literal                   # LiteralExpr
                | listLiteral               # ListExpr
                | T_TRUE                    # TrueExpr
                | T_FALSE                   # FalseExpr
                ;

// Maintain typed list literals
listLiteral     : STRING_LIST               # StringListExpr
                | INTEGER_LIST              # IntegerListExpr
                | DECIMAL_LIST              # DecimalListExpr
                ;

// Function call structure
funcCall        : IDENTIFIER '(' exprList? ')';
exprList        : expr (COMMA expr)*;

// Literal value types
literal         : STRING                    # StringExpr
                | INTEGER                   # IntegerExpr
                | DECIMAL                   # DecimalExpr
                ;

// Token definitions
T_AND       : '&';
T_OR        : '|';
T_NOT       : '!';
T_TRUE      : 'true';
T_FALSE     : 'false';
T_EQ        : '==';
T_NEQ       : '!=';
T_GT        : '>';
T_GTE       : '>=';
T_LT        : '<';
T_LTE       : '<=';
T_ADD       : '+';
T_SUB       : '-';
T_MUL       : '*';
T_DIV       : '/';
T_MOD       : '%';
COMMA       : ',';
SEMI        : ';';
T_EQUAL     : '=';
QUOTA       : '"';

// Improved list definitions with better whitespace handling
STRING_LIST     : '[' WS* (STRING (WS* ',' WS* STRING)*)? WS* ']';
INTEGER_LIST    : '[' WS* (INTEGER (WS* ',' WS* INTEGER)*)? WS* ']';
DECIMAL_LIST    : '[' WS* (DECIMAL (WS* ',' WS* DECIMAL)*)? WS* ']';

// Identifier with case sensitivity
IDENTIFIER  options { caseInsensitive=false; } : [_a-zA-Z][_a-zA-Z0-9]*;

// Numeric literals
INTEGER     : '0' | [1-9][0-9]*;
DECIMAL     : '-'? ([0-9]* '.' [0-9]+ | [0-9]+ '.' [0-9]*);

// String literal with escape handling
STRING      : QUOTA (ESC | SAFECODEPOINT)* QUOTA;

// Fragment rules
fragment ESC: '\\' (["\\/bfnrt] | UNICODE);
fragment UNICODE: 'u' HEX HEX HEX HEX;
fragment HEX options { caseInsensitive=false; }  : [0-9a-fA-F];
fragment SAFECODEPOINT: ~ ["\\\u0000-\u001F];

// Whitespace handling
WS          : [ \t\n\r]+ -> skip;