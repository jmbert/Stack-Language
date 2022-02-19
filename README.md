# Stack-Language

Type number to add it to the stack
Seperate Commands by a new line
+ to add, - to subtract, * to multiply, / to divide (rounded), ^ to exponentiate, PRINT to print, STDIN to read from standard input
POP removes the top element and puts it in the current value slot (will overwrite if popped twice)
PUSH pushes the current value onto the stack
DUPE duplicates the top element
JUMP() jumps to the number in the parentheses. Use dunder TOP to signify the top element, and use dunder TOPREPLACE to use the top element without pulling it off the stack
CONDJUMP() jumps if the condition is true : e.g CONDJUMP(>0:1) jumps to line 1 if the top element is greater then 0
INPUT() prints the string in the parentheses, then takes input
SLEEP() sleeps for the number of seconds in the parentheses
CLEAR clears the stack
CONCAT concatanates the last two elements
LEN pushes the length of the stack onto the stack

for example :
INPUT(Starting number: )
PRINT
1
-
SLEEP(1)
CONDJUMP(>0:1)
PRINT
JUMP(0)

this would take an input, then subtract one until it reached 0, then take another input, etc.

NOTE:
floats are not allowed, only integers