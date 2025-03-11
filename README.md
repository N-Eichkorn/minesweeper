# Build from Source
1. ```go build main.go -o minesweeper ```
2. ```go run main.go```

:warning: Run in Windows

The console may not support ANSI codes in the batch terminal. To run the app in Windows. You must set the Virtual Terminal Level in the registry
```
[HKEY_CURRENT_USER\Console]
"VirtualTerminalLevel"=dword:00000001
```