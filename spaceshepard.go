package main

import (
	"fmt"
	"github.com/MalloZup/spaceshepard/ssh"
)

func main() {
	fmt.Printf("SpaceShepard.\n")
	ssh.Run("uptime")
}
