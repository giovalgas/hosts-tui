package reader

import (
	"bufio"
	"fmt"
	"github.com/charmbracelet/bubbles/list"
	"os"
	"strings"
)

const FilePath = "/etc/hosts"

type Item struct {
	host        string
	hostname    []string
	description string
}

func (i Item) Title() string {
	return strings.Join(i.hostname, " ")
}

func (i Item) Description() string {
	return fmt.Sprintf("Host: %s, Description: %s", i.host, i.description)
}

func (i Item) FilterValue() string {
	return strings.Join(i.hostname, ",") + "," + i.host
}

func NewItem(host string, hostname []string, description string) Item {

	if description == "" {
		description = "No description provided..."
	}

	return Item{host: host, hostname: hostname, description: description}
}

func ReadHostsFile() []list.Item {
	readFile, err := os.Open(FilePath)

	if err != nil {
		fmt.Println(err)
		os.Exit(1)
	}

	fileScanner := bufio.NewScanner(readFile)
	fileScanner.Split(bufio.ScanLines)

	var items []list.Item

	for fileScanner.Scan() {
		line := fileScanner.Text()

		// Ignore empty lines
		if line == "" {
			continue
		}

		// Ignore lines that either start with "#" or "-e"
		if strings.HasPrefix(line, "#") || strings.HasPrefix(line, "-e") {
			continue
		}

		// Remove extra whitespaces from string
		line = strings.Join(strings.Fields(line), " ")

		// Split line by #
		splitLine := strings.Split(line, "#")

		// Get Host and Hostnames
		hosts := strings.Split(splitLine[0], " ")

		host := hosts[0]
		hostnames := hosts[1:]

		// Append to items slice
		items = append(items, NewItem(host, hostnames, strings.Join(splitLine[1:], ", ")))
	}

	return items
}
