Say you have a file with test grades of students it may look something like this

Bob 95
Tom 75
Rod 41
Ned 91
Sal 85

{ print $2 }

awk '{print "You have scored", $2}'
