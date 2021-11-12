package main

import (
	"fmt"
)

type pc struct {
	ram   int
	disk  int
	brand string
}

func (myPC pc) String() string {
	return fmt.Sprintf("Este PC tiene ram: %d, disk: %d, brand: %s", myPC.ram, myPC.disk, myPC.brand)
}

func main() {
	myPC := pc{ram: 8, disk: 500, brand: "Dell"}

	fmt.Println(myPC)

}
