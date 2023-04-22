package mode

import "fmt"

type ConsolePrint struct {
}

func NewConsolePrint() *ConsolePrint {
	return &ConsolePrint{}
}

func (cp *ConsolePrint) PrintDataOnConsole(data string) {
	fmt.Println(data)
}
