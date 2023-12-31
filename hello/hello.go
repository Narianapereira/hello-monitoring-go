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

const monitoringTimes = 5
const delay = 5

func main() {

	introduction()
	for {
		showMenu()

		cmd := readCmd()

		if cmd == 1 {
			monitoring()
		} else if cmd == 2 {
			fmt.Println("Showing logs")
			printLogs()
		} else if cmd == 0 {
			fmt.Println("Exiting")
			os.Exit(0)
		} else {
			fmt.Println("Wrong cmd")
			os.Exit(-1)
		}
	}
}

func monitoring() {
	fmt.Println("Monitoring...")

	sites := readFileSites()

	for i := 0; i < monitoringTimes; i++ {
		for _, site := range sites {
			testSite(site)
		}
		time.Sleep(delay * time.Second)
	}
}

func testSite(site string) {

	resp, err := http.Get(site)

	if err != nil {
		fmt.Println("The error was:", err)
	}

	if resp.StatusCode == 200 {
		fmt.Println("Site:", site, "Carregado com sucesso!")
		logRegister(site, true)
	} else {
		logRegister(site, false)
		fmt.Println("Erro no site", site, resp.StatusCode)
	}

}
func showMenu() {

	fmt.Println("1 - Start monitoring")
	fmt.Println("2 - Show logs")
	fmt.Println("0 - Exit")
}

func introduction() {
	name := "Nari"
	version := 1.1

	fmt.Println("Hello", name)
	fmt.Println("This program is version", version)
}

func readCmd() int {
	var cmd int
	fmt.Scan(&cmd)

	fmt.Println("The chosen cmd was", cmd)
	return cmd
}

func readFileSites() []string {
	var sites []string
	newFile, err := os.Open("sites.txt")
	if err != nil {
		fmt.Println("The error was:", err)
	}

	reading := bufio.NewReader(newFile)
	for {
		line, err := reading.ReadString('\n')
		line = strings.TrimSpace(line)

		sites = append(sites, line)
		if err == io.EOF {
			break
		}
	}
	newFile.Close()
	return sites
}

func logRegister(site string, status bool) {
	newFile, err := os.OpenFile("log.txt", os.O_RDWR|os.O_CREATE|os.O_APPEND, 0666)

	if err != nil {
		fmt.Println("Error:", err)
	}
	newFile.WriteString(time.Now().Format("02/01/2006 15:04:05") + " - " + site + " -online: " + strconv.FormatBool(status) + "\n")
	newFile.Close()
}

func printLogs() {
	newFile, err := ioutil.ReadFile("log.txt")

	if err != nil {
		fmt.Println(err)
	}

	fmt.Println(string(newFile))
}
