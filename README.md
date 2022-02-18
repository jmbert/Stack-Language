# Stack-Language

Type number to add it to the stack
Seperate Commands by a single space
+ to add, - to subtract, * to multiply, / to divide (rounded), ^ to exponentiate, PRINT to print, STDIN to read from standard input
POP removes the top element and puts it in the current value slot (will overwrite if popped twice)
PUSH pushes the current value onto the stack
DUPE duplicates the top element

for example :
3 5 +
6 -
PRINT

this would print (3 + 5) - 6

NOTE:
floats are not allowed, only integers