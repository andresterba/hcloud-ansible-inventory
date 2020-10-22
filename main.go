package main

import (
	"context"
	"errors"
	"fmt"
	"log"
	"os"

	"github.com/hetznercloud/hcloud-go/hcloud"
)

func showHelp() {
	fmt.Println(`Usage:
    ansible-hcloud-inventory [option]
	--inventory-file [path to file]`)

	os.Exit(1)
}

func getAuthToken() (string, error) {
	authTokenFromEnv := os.Getenv("HETZNER_CLOUD_API_KEY")

	if len(authTokenFromEnv) == 0 {
		return "", errors.New("Please set HETZNER_CLOUD_API_KEY!")
	}

	return authTokenFromEnv, nil
}

func checkForError(err error) {
	if err != nil {
		log.Fatal(err)

		os.Exit(1)
	}
}

func main() {
	authToken, err := getAuthToken()
	checkForError(err)

	args := os.Args

	if len(args) < 3 {
		showHelp()
	}

	pathToWriteInventory := args[2]

	if len(pathToWriteInventory) == 0 {
		log.Fatal("Please set a path to write the inventor")
		os.Exit(1)
	}

	client := hcloud.NewClient(hcloud.WithToken(authToken))

	servers, err := client.Server.All(context.Background())
	if err != nil {
		log.Fatalf("Error retrieving servers: %s\n", err)
		os.Exit(1)
	}

	f, err := os.Create(pathToWriteInventory)
	checkForError(err)

	defer f.Close()

	_, err = f.WriteString("[hetzner]\n")
	checkForError(err)

	for _, server := range servers {
		hostEntry := fmt.Sprintf("%s ansible_host=%s ansible_user=root \n", server.Name, server.PublicNet.IPv4.DNSPtr)

		_, err := f.WriteString(hostEntry)
		checkForError(err)
	}
}
