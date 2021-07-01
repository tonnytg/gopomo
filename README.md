# gopomo
<hr>
Go Pomodoro create a prompt Pomodoro to help your tasks.

### How to use?

Keep it simple

```go run main.go```

With this will get 1 Pomodoro with 25 minutes!
Ok do you want to change?

`go run main.go -quant=X -time=Y`

`-quant` -> How many Pomodoros do you want?

`-time` -> Maybe do you wanna change for 15 minutes or more.

<hr>

### How build this binary?

`go build -o goPomo`

`./goPomo`

<hr>

### Example Output
```
$ ./gopomo 
Time: 25
Quantity: 1
Start: 19:13:54 Pomodoro: 0/1
End: 19:38:54 Pomodoro: 1/1
Take 5 minutes for a cofee!
Press Enter to continue!
```

Custom time:

```
$ ./gopomo -quant 1 -time 1
Time: 1
Quantity: 1
Start: 19:15:44 Pomodoro: 0/1
End: 19:16:44 Pomodoro: 1/1
Take 5 minutes for a cofee!
Press Enter to continue!
```

Be Happy!