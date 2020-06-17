package main
import(
	"fmt"
	"net"
)

func main() {
	fmt.Println("Starting the server ...")
	listner, err := net.Listen("tcp", "0.0.0.0:4000")
	if err != nil {
		fmt.Println("Error listening", err.Error())
		return
	}

	for {
		conn, err := listner.Accept()
		if err != nil {
			return
		}
		go printMessage(conn)
	}
}

func printMessage(conn net.Conn) {
	for {
		buf := make([]byte, 512)
		_, err := conn.Read(buf)
		if err != nil {
			fmt.Println("Error reading", err.Error())
			return
		}
		fmt.Printf("Received data: %v", string(buf))
	}
}