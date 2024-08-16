# Don't Do This With Recursion (!!!)

Up: [Recursion in Java](recursion_in_java)
Brother(s):
TARGET DECK
Leaves of the Tree of Ohara::Branch I::Semester II::1206 Programming II::Recursion

**DON'T** make a recursive method that is equivalent to `while (true)`.

…*just don't.*

WHAT SHOULDN'T YOU DO IN JAVA WITH RECURSION!!! (And why?) #flashcard 
DO NOT MAKE A RECURSIVE METHOD THAT IS THE SAME AS `while(true)` OR ELSE I WILL EAT YOU!!!
Because every time the method calls itself, a new [The Java Stack Frame|stack frame](the_java_stack_frame|stack_frame) is created, and if this looks like it will go on forever, this will lead to a stack overflow (KABOOM)
<!--ID: 1716281947117-->




































#### Why:
#### How:









