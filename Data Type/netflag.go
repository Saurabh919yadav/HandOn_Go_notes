package main

import "fmt"

type Flags uint

const (
	FlagUp Flags = 1 << iota //is up
	FlagBroadcast
	FlagLoopback
	FlagPointToPoint
	FlagMulticast
)

func IsUp(v Flags) bool {
	return v&FlagUp == FlagUp
}

func TurnDown(v *Flags) {
	*v &^= FlagUp
}

func SetBroadcast(v *Flags) {
	*v |= FlagBroadcast
}

func IsCast(v Flags) bool {
	return v & (FlagBroadcast | FlagMulticast)
}

func main() {
	var v Flags = FlagMulticast | FlagUp
	fmt.Printf("%b %t\n", IsUp(v))
	TurnDown(&v)
	fmt.Printf("%b %t\n", IsUp(v))
	SetBroadcast(&v)
	fmt.Printf("%b %t\n", v, IsUp(v))
	fmt.Printf("%b %t\n", v, IsCast(v))
}
