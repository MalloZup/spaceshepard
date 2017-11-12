package main

import (
    "bytes"
    "log"
    "os/exec"
    "sync"
)

func runCmd(remMachine string, command string, wg *sync.WaitGroup) {
    wg.Add(1)

    go func() {
    cmd := exec.Command("ssh", remMachine, command)
    var out bytes.Buffer
    cmd.Stdout = &out
    err := cmd.Run()
    if err != nil {
        log.Fatal(err)
    }
  }
   defer wg.Done()
}


func main() {
    var remoteMachine  = "root@localhost"
    var wg sync.WaitGroup
    runCmd(remoteMachine, "uptime", &wg)
    wg.Wait()
}
