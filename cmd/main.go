package main

import "github.com/beeceej/wire-sandbox/pkg/thing"

func main() {
	t1 := thing.InjectThing()
	t1.Say("stdout: Hello, World\n")
	t2 := thing.InjectStderrThing()
	t2.Say("stderr: Hello, World")
}
