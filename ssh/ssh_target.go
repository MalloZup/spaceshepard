package ssh

import (
	"bytes"
	"fmt"
	"golang.org/x/crypto/ssh"
)

// Run a ssh command given a target
func Run(cmd string) (bytes.Buffer, error) {
        var hostKey ssh.PublicKey
	// remove this from function
	config := &ssh.ClientConfig{
		User: "root",
		Auth: []ssh.AuthMethod{
			ssh.Password("linux"),
		},
         HostKeyCallback: ssh.FixedHostKey(hostKey),
	}

	// make this configurable via json later
	client, err := ssh.Dial("tcp", "suma-refhead-srv.mgr.suse.de:22" , config)
	if err != nil {
		panic("Failed to dial: " + err.Error())
	}

	session, err := client.NewSession()
	if err != nil {
		panic("Failed to create session: " + err.Error())
	}
	defer session.Close()

	var b bytes.Buffer
	session.Stdout = &b
        // make this via json later
	sshErr := session.Run("/usr/bin/whoami")
	// just for debug
	fmt.Println(b.String())

	return b, sshErr
}
