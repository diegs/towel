# towel

Don't panic.

Towel is a program that wraps and executes another program, capturing stderr
and reformatting it into a JSON output on stdout.

Example:

```
❯ go run towel.go ./test.sh a b c
hello from stdout, argvs are a b c
2019-02-13T20:57:56.77673 {"level":"fatal","ts":155009147677.6734048,"logger":"./test.sh","caller":"main.go","msg":"Process ended with non-zero exit code","pid":27074,"json":{"error":exit status 2,"stderr":""hello from stderr, argvs are a b c\n""}}
```
