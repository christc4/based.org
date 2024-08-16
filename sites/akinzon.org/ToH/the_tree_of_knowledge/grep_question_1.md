# GREP Question 1

Up: [GREP Question Solutions](grep_question_solutions)
Brother(s):
TARGET DECK

```bash
cat sample.txt | grep an
banana
mango
```

```sh
cat sample.txt | grep -e "\Wdo\W"
Just do-it
```

```bash
cat sample.txt | grep -e "H∥he" | grep -e "World∥Hi"
Hello World
Hi there
```

```bash
cat code.txt | grep "fruit0"
fruit[0] = 'apple'
```

```bash
cat sample.txt | grep -e "t" | head -n2
Hi there
Just do-it
```

```bash
cat sample.txt | egrep  '[he]{0}' | head -n3
Hello World

How are you
```

```bash
cat -n sample.txt | grep 'do'
     6  Just do-it
    13  Much ado about nothing
```

```bash
cat sample.txt | sed 's/ /\n/g' | grep 'H∥he' | wc -l
5
```

```bash
cat sample.txt | grep  '^$' | wc -l
4
```

```bash
$ cat terms.txt
are
not
go
fruit[0]

cat terms.txt | sed 's/\[\|\]/\\&/g' | grep -f - sample.txt code.txt
sample.txt:How are you
sample.txt:mango
sample.txt:Much ado about nothing
sample.txt:Adios amigo
code.txt:fruit[0] = 'apple'
```

```bash
grep -H "amigo" sample.txt
sample.txt:15:Adios amigo
```

```bash
grep -l "apple" sample.txt code.txt
code.txt
```

```bash
$ cat lines.txt
banana
fruit = []

cat lines.txt | sed 's/\[\|\]/\\&/g'   | grep -nxHEf - sample.txt code.txt
sample.txt:9:banana
code.txt:1:fruit = []
```

```bash
cat terms.txt | sed 's/\[\|\]/\\&/g'   | grep -nvcHEf - sample.txt code.txt
sample.txt:11
code.txt:3
```


```bash
cat sample.txt code.txt | grep -c 'banana'
2
```







#### Why:
#### How:









