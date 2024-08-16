# AWK Question 1

Up: [AWK Question Solutions](awk_question_solutions)
Brother(s):
TARGET DECK

```bash
$ cat addr.txt
Hello World
How are you
This game is good
Today is sunny
12345
You are funny

$ awk cat addr.txt | awk '/is/{  print }'
This game is good
Today is sunny
```

```bash
$ awk cat addr.txt | awk '!/y/{  print $1}'
This
12345
```

```bash
$ awk cat addr.txt | awk '{if (NF < 3) print}'   
Hello World
12345
```

```bash
$ awk cat addr.txt | awk 'sub(/o/, 0) {print}'
Hell0 World
H0w are you
This game is g0od
T0day is sunny
12345
Y0u are funny
```

```bash
$ cat table.txt
brown bread mat hair 42
blue cake mug shirt -7
yellow banana window shoes 3.14

$ awk cat table.txt | awk 'BEGIN {count = 1};  {if ($5 > -10) count = count * $5}  END {print count}'
-923.16
```

```bash
printf 'last\nappend\nstop\ntail\n' | awk 'BEGIN {FS = "\n";  dot = "."} ; {$1 = $1dot; print}'
```

```bash
$ cat hex.txt
start address: 0xA0, func1 address: 0xA0
end address: 0xFF, func2 address: 0xB0

$ awk cat hex.txt | awk 'BEGIN {rep1 = "0x50"; rep2="0x7F"}; { gsub(/0xA0/, rep1); gsub(/0xFF/, rep2);  print }'
start address: 0x50, func1 address: 0x50
end address: 0x7F, func2 address: 0xB0
```




































#### Why:
#### How:









