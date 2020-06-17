package main
import (
	"fmt"
	"os"
	"net"
	"bufio"
	"strings"
)

func main() {
	conn, err := net.Dial("tcp", ":4000")
	if err != nil {
		fmt.Println("Error connecting", err.Error())
	}
	inputReader := bufio.NewReader(os.Stdin)
	fmt.Println("Enter your name")
	clientName, _ := inputReader.ReadString('\n')
	trimmedClient := strings.Trim(clientName, "\n")
	for {
		fmt.Println("Enter Message or Q to quit")
		input, _ := inputReader.ReadString('\n')
		trimmedInput := strings.Trim(input, "\n")
		if trimmedInput == "Q" {
			return
		}
		_, err = conn.Write([]byte(trimmedClient + " says: " + trimmedInput +"\n"))
	}
}