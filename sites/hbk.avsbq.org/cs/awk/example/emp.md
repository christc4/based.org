# Awk by example

This is straight from the book, these are the examples I am following and I am manually typing all of this so excuse any typos, I have condensed where I can

Beth	21	0
Dan	19	0
Kathy	15.50	10
Mark	25	20
Mary	22.50	22
Susie	17	18

awk '$3 > 0 { print $1, $2 * $3 }' emp.data

Output will be

Kathy 155
Mark 500
Mary 495
Susie 306

What is inside the quotes is the complete awk program, it consists of a single pattern action statement.

The pattern $3 > 0 matches every input line in which the third column/field is greater than zero

and the action

{ print $1, $2 * $3 }

prints the first field and the product of  the second and third fields of each matched line

To print non-working employee names, run

awk '$3 == 0 { print $1 }' emp.data

this is a single pattern-action statement, for every line in which the third field is zero, the first field is printed

The output is 

Beth
Dan
