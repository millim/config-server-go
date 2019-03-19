package main

import (
	"bufio"
	"bytes"
	"encoding/json"
	"fmt"
	"io/ioutil"
	"log"
	"os"
	"path/filepath"
	"strings"

	"golang.org/x/crypto/ssh"
)

//DeployConfig config file fields
// fields have HostName and PrivateKey
type DeployConfig struct {
	SSHHost    string `json:"sshHost"`
	SSHPort    string `json:"sshPort"`
	PrivateKey string `json:"privateKey"`
}

func getHostKey(host string) (ssh.PublicKey, error) {
	file, err := os.Open(filepath.Join(os.Getenv("HOME"), ".ssh", "known_hosts"))
	if err != nil {
		return nil, err
	}
	defer file.Close()
	scanner := bufio.NewScanner(file)
	var hostKey ssh.PublicKey
	for scanner.Scan() {
		fields := strings.Split(scanner.Text(), " ")
		if len(fields) != 3 {
			continue
		}
		if strings.Contains(fields[0], host) {
			var err error
			hostKey, _, _, _, err = ssh.ParseAuthorizedKey(scanner.Bytes())
			if err != nil {
				return nil, fmt.Errorf("error parsing %q: %v", fields[2], err)
			}
			break
		}
	}
	if hostKey == nil {
		return nil, fmt.Errorf("no hostkey for %s", host)
	}
	return hostKey, nil
}

func main() {

	plan, _ := ioutil.ReadFile("./deploy/config.json")
	configData := DeployConfig{}
	err := json.Unmarshal(plan, &configData)
	if err != nil {
		log.Fatalf("read config file is error : %v", err)
	}
	hostKey, err := getHostKey(fmt.Sprintf("[%s]:%s", configData.SSHHost, configData.SSHPort))
	if err != nil {
		log.Fatal(err)
	}
	key, err := ioutil.ReadFile(configData.PrivateKey)
	if err != nil {
		log.Fatalf("unable to read private key: %v", err)
	}
	signer, err := ssh.ParsePrivateKey(key)
	if err != nil {
		log.Fatalf("private key: %v", err)
	}
	config := &ssh.ClientConfig{
		User: "root",
		Auth: []ssh.AuthMethod{
			ssh.PublicKeys(signer),
		},
		HostKeyCallback: ssh.FixedHostKey(hostKey),
	}
	client, err := ssh.Dial("tcp", fmt.Sprintf("%s:%s", configData.SSHHost, configData.SSHPort), config)
	if err != nil {
		log.Fatalf("unable to connect: %v", err)
	}
	session, err := client.NewSession()
	if err != nil {
		log.Fatal("Failed to create session: ", err)
	}
	defer session.Close()
	var b bytes.Buffer
	session.Stdout = &b
	if err := session.Run("ps -ef|grep ssserver"); err != nil {
		log.Fatal("Failed to run: " + err.Error())
	}
	fmt.Println(b.String())
	defer client.Close()

}
