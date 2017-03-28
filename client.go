package main

import (
	"fmt"
	"io"
	"net"
	"os"
	"strconv"
)

const BUFFERSIZE = 1024

func main() {
	connection, err := net.Dial("tcp", "localhost:27001")
	if err != nil {
		panic(err)
	}
	// defer connection.Close()
	fmt.Println("Connected to server")
	// go sendFileToClient(connection)
	file, err := os.Open("data/person.csv")
	if err != nil {
		fmt.Println(err)
		return
	}
	fileInfo, err := file.Stat()
	if err != nil {
		fmt.Println(err)
		return
	}
	fileSize := fillString(strconv.FormatInt(fileInfo.Size(), 10), 10)
	fileName := fillString(fileInfo.Name(), 64)
	fmt.Println("Sending filename and filesize!")
	connection.Write([]byte(fileSize))
	connection.Write([]byte(fileName))
	sendBuffer := make([]byte, BUFFERSIZE)
	fmt.Println("Start sending file!")
	for {
		_, err = file.Read(sendBuffer)
		if err == io.EOF {
			break
		}
		connection.Write(sendBuffer)
	}
	fmt.Println("File has been sent, closing connection!")
	return

}

//
// func sendFileToClient(connection net.Conn) {
// 	fmt.Println("A client has connected!")
// 	defer connection.Close()
// 	file, err := os.Open("data/person.csv")
// 	if err != nil {
// 		fmt.Println(err)
// 		return
// 	}
// 	fileInfo, err := file.Stat()
// 	if err != nil {
// 		fmt.Println(err)
// 		return
// 	}
// 	fileSize := fillString(strconv.FormatInt(fileInfo.Size(), 10), 10)
// 	fileName := fillString(fileInfo.Name(), 64)
// 	fmt.Println("Sending filename and filesize!")
// 	connection.Write([]byte(fileSize))
// 	connection.Write([]byte(fileName))
// 	sendBuffer := make([]byte, BUFFERSIZE)
// 	fmt.Println("Start sending file!")
// 	for {
// 		_, err = file.Read(sendBuffer)
// 		if err == io.EOF {
// 			break
// 		}
// 		connection.Write(sendBuffer)
// 	}
// 	fmt.Println("File has been sent, closing connection!")
// 	return
// }

func fillString(retunString string, toLength int) string {
	for {
		lengtString := len(retunString)
		if lengtString < toLength {
			retunString = retunString + ":"
			continue
		}
		break
	}
	return retunString
}

func checkError(err error) {
	if err != nil {
		fmt.Fprintf(os.Stderr, "Fatal error: %s", err.Error())
		os.Exit(1)
	}
}
