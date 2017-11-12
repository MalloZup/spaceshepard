package main

import (
    "bytes"
    "log"
    "os/exec"
)




func main() {
    var remoteMachine  = "root@localhost"
    var out bytes.Buffer
    cmd := exec.Command("ssh", remoteMachine, "uptime")
    cmd.Stdout = &out
    err := cmd.Run()
    if err != nil {
        log.Fatal(err)
    }
}
