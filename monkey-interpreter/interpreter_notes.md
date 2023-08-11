## Interpreter
Take source code and evaluate it without producing any intermediate result(compiled code) that can later be executed  
## Compiler
Take source code and produce output in another language that the underlying system can understand.  
### Types of Interpreters 
- no parsing step they just interpret the input right away.
- Highly Optimized advance parsing and evaluation techniques. Some of them compile it to internal representation called bytecode and evalute them.
- Even more advanced are JIT Interpreters that compile the input just-in-time into native machine code that get's executed.  
- between these two, some parse the source code, build an abstract tree out of it and then evaluate this tree. This type of interpreter is called "tree-walking" interpreter because it walks the tree and interprets it 

### Features
Expressed as a list of features, Monkey has the following:  
• C-like syntax
• variable bindings
• integers and booleans
• arithmetic expressions
• built-in functions
• first-class and higher-order functions
• closures
• a string data structure
• an array data structure
• a hash data structure

### Major Steps
• the lexer
• the parser
• the Abstract Syntax Tree (AST)
• the internal object system
• the evaluator

# Lexing

