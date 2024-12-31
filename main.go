package depcalldemoa

import (
	"fmt"

	"github.com/csguojin/depcalldemob"
	"github.com/csguojin/depcalldemoc"
	"github.com/csguojin/depcalldemoe"
	"github.com/csguojin/depcalldemoh"
)

func MyFunc1() {
	fmt.Println("depcalldemoA.MyFunc1")
	depcalldemob.MyFunc1()
	depcalldemoc.MyFunc1()
}

func MyFunc2() {
	fmt.Println("depcalldemoA.MyFunc2")
	depcalldemoe.MyFunc1()
	depcalldemoh.MyFunc1()
}
