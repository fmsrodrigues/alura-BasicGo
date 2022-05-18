package main

import (
	"bufio"
	"fmt"
	"io"
	"io/ioutil"
	"net/http"
	"os"
	"strconv"
	"strings"
	"time"
)

const monitoringFrequency = 3
const monitoringDelay = 10

func main() {
	introduction()

	for {
		showMenu()
		option := getOption()

		switch option {
		case 1:
			startMonitoring()
		case 2:
			showLogs()
		case 0:
			gracefulShutdown()
		default:
			fmt.Println("Invalid option")
			ungracefulShutdown()
		}

		fmt.Println()
	}
}

func introduction() {
	name := "Watcher"
	version := 1.1

	fmt.Println("Hello,", name, "!")
	fmt.Println("Program running version", version)
}

func showMenu() {
	fmt.Println("1 - Start monitoring")
	fmt.Println("2 - Show Logs")
	fmt.Println("0 - Exit")
}

func getOption() int {
	var option int
	fmt.Scan(&option)

	return option
}

func gracefulShutdown() {
	fmt.Println("Exiting...")
	os.Exit(0)
}

func ungracefulShutdown() {
	os.Exit(-1)
}

func startMonitoring() {
	fmt.Println("Monitoring...")
	apps := readAppsToMonitorFile()

	// for index, item := range items
	for i := 0; i < monitoringFrequency; i++ {
		fmt.Println()
		for _, app := range apps {
			verifyApp(app)
		}
		time.Sleep(monitoringDelay * time.Second)
	}
}

func verifyApp(app string) {
	res, err := http.Get(app)

	if err != nil {
		fmt.Println("Error on verifying app:", err)
		ungracefulShutdown()
	}

	if res.StatusCode == 200 {
		fmt.Println("App:", app, "was successfully reached")
	} else {
		fmt.Println("App:", app, "had an error. StatusCode:", res.StatusCode)
	}

	writeLog(app, res.StatusCode == 200)
}

func readAppsToMonitorFile() []string {
	var apps []string
	file, err := os.Open("apps.txt")

	if err != nil {
		fmt.Println("Error on opening file to read:", err)
		ungracefulShutdown()
	}

	reader := bufio.NewReader(file)

	for {
		line, err := reader.ReadString('\n')
		line = strings.TrimSpace(line)

		apps = append(apps, line)

		if err == io.EOF {
			break
		}
	}

	file.Close()

	return apps
}

func writeLog(app string, online bool) {
	file, err := os.OpenFile("log.txt", os.O_RDWR|os.O_CREATE|os.O_APPEND, 0666)

	if err != nil {
		fmt.Println("Error on opening file to write:", err)
	}

	file.WriteString(time.Now().Format("02/01/2006 15:04:05 - ") + app + " - online: " + strconv.FormatBool(online) + "\n")

	file.Close()
}

func showLogs() {
	fmt.Println("Showing logs...")

	// ioutil.ReadFile automatic closes file
	file, err := ioutil.ReadFile("log.txt")

	if err != nil {
		fmt.Println("Error on opening file to read:", err)
	}

	fmt.Println(string(file))
}
