# towel

Don't panic.

Binary towel runs a program specified on the command line, intercepts stderr, and reformats any invalid lines to
fatal structured log output.

Example:

```
❯ go run towel.go ./test.sh a b c
12345 {this is a valid log output, argvs are a b c}
2019-02-13T21:29:32.03462 {"level":"fatal","ts":155009337203.4628731,"logger":"./test.sh","caller":"./test.sh","msg":"Unexpected program output","pid":26780,"json":{"stderr":"this is invalid logging\non many lines\n"}}12345 {back to valid logging}
12345 {for a few lines}
2019-02-13T21:29:32.03476 {"level":"fatal","ts":155009337203.4762143,"logger":"./test.sh","caller":"./test.sh","msg":"Unexpected program output","pid":26780,"json":{"stderr":"and then a panic!\non a few lines\n"}}
```
