package mypackage

import "fmt"

// Un pc para crear
type Pc struct {
	Brand string
	Model string
	Ram   int
	Disk  int
}

func (myPC *Pc) Duplicateramdisk(ramchek, diskchek bool) {
	if ramchek == true {
		myPC.Ram *= 2
		fmt.Println("la Ram fue duplicada", myPC.Ram)
	}
	if diskchek == true {
		myPC.Disk *= 2
		fmt.Println("el disco fue duplicado", myPC.Disk)
	}
}
