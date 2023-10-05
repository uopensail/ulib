# UNO
This is an expression engine that serves the recommender system and is mainly divided into three parts:
1.js is the parsing and optimization of expressions
2.cpp is the evaluation of the expression engine
3. go is a wrapper for libraries used externally

## Why do you need an expression engine?
In actual business, we find that conditional filtering occurs in many scenarios, such as: 
When generating an item pool, you need to filter out the item pool according to certain conditions; At the time of recall, conditional filtering is required. Therefore, we want to abstract an expression engine and handle it efficiently.

## js
Here is mainly to parse the string into AST, on the one hand, including the syntax parsing part, which is implemented using [ANTLR](https://www.antlr.org/). Another aspect is the optimization of AST, which mainly includes:
1. Not Push down
2. Constant folding and propagation
3. Pruning
   1. In AND operation, as long as a node is false, false is returned
   2. In OR operation, return true as long as a node is true

## go&cpp
Here is mainly the calculation of expressions, using JS to convert AST into a node array, and then sequentially perform node-related operations, rather than using the structure of the tree for recursive calculation.