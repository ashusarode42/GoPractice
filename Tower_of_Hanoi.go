package main

import (
	"fmt"
)

func main() {
	toi(4, "A", "C", "B")
}

/* function to run recursively all the moves required for tower of hanoi. It is a repeation of 3 moves */
func toi(count int, from_tower, to_tower, auxiliary_tower string) {

	if count == 1 {
		fmt.Printf("Moved disk %v %v --> %v\n", count, from_tower, to_tower)
		return
	}
	//Call toi() method which runs other successive moves accordingly.
	toi(count-1, from_tower, auxiliary_tower, to_tower)
	//Display the moves made by disks.
	fmt.Printf("Moved disk from %v %v => %v\n", count, from_tower, to_tower)
	toi(count-1, auxiliary_tower, to_tower, from_tower)

}
