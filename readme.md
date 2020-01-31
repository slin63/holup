# Holup
Grabs a PID or process name and suspends it until the program is `SIGINT`'d (`ctrl+c`).

Usage:
1. `go build -o /usr/local/bin/holup main.go`
2. `holup -n com.docker.hyperkit`

_or_

2. `pgrep com.docker.hyperkit`
3. `holup -p <identified PID>`
