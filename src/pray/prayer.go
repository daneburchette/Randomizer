package pray

import (
	"fmt"
)

func Prayer(choice string) {
	fmt.Println("Place your head on the BatterUp peripheral,")
	fmt.Println("Point your No-No hole in a random direction")
	fmt.Println("and say the prayer be all love to say:")
	fmt.Scanln()
	for i := 0; i < 3; i++ {
		fmt.Println("Bee-da-bud-a-bud-a")
		fmt.Scanln()
	}
	for i := 0; i < 3; i++ {
		fmt.Println("Boop")
		fmt.Scanln()
	}
	fmt.Println("No whammies...")
	fmt.Scanln()
	fmt.Println("No Whammies!")
	fmt.Scanln()
	fmt.Println("NO WHAMMIES!")
	fmt.Scanln()
	fmt.Println("STOP!!!!")
	fmt.Scanln()
	fmt.Println(choice)
	fmt.Println("(Return to exit)")
	fmt.Scanln()
}
