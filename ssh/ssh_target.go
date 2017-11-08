package ssh

import (
         "golang.org/x/crypto/ssh"
         "bytes"
         "fmt"
       )
// Run a ssh command given a target
func Run(cmd string) (bytes.Buffer, error) {

	// remove this from function
	config := &ssh.ClientConfig{
		User: "root",
		Auth: []ssh.AuthMethod{
			ssh.Password("linux"),
		},
	}

	// make this configurable via json later
	port, server := ":22", "suma-refhead-srv.mgr.suse.de"
	client, err := ssh.Dial("tcp", server+port, config)
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
	sshErr := session.Run("/usr/bin/whoami")
	// just for debug
	fmt.Println(b.String())

	return b, sshErr
}
