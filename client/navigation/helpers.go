package navigation

import (
	"os"
	"os/exec"
	"runtime"
	"time"
)

var clear map[string]func()

func init() {
	clear = make(map[string]func())
	clear["linux"] = func() {
		cmd := exec.Command("clear") //Linux example, its tested
		cmd.Stdout = os.Stdout
		cmd.Run()
	}
	clear["windows"] = func() {
		cmd := exec.Command("cmd", "/c", "cls") //Windows example, its tested
		cmd.Stdout = os.Stdout
		cmd.Run()
	}
}

// Clear clears the terminal screen.
func Clear() {
	value, ok := clear[runtime.GOOS]
	if ok {
		time.Sleep(100 * time.Millisecond)
		value()
		time.Sleep(100 * time.Millisecond)
	} else {
		panic("Platform is unsupported.")
	}
}
