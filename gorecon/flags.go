package main

import (
	"flag"
	"fmt"
)

func main() {
	stringFlagPtr := flag.String("StringflagName", "default value", "example usgae statement")
	intFlagPtr := flag.Int("IntflagName", 45, "integer flag usage")
	boolBoolPtr := flag.Bool("BoolfagName", false, "bool flag usage")

	flag.Parse()

	fmt.Println("string flag: ", *stringFlagPtr)
	fmt.Println("int flag: ", *intFlagPtr)
	fmt.Println("bool flag: ", *boolBoolPtr)

}
