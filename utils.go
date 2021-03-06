package gitany

import (
	"bufio"
	"fmt"
	"os"
	"syscall"

	"golang.org/x/crypto/ssh/terminal"
)

func PanicIfErrorExist(err error) {
	if err != nil {
		panic(err)
	}
}

func PromptUserAndPassword(hostName string) (string, string) {
	r := bufio.NewReader(os.Stdin)
	fmt.Print(hostName, " Username: ")
	username, _ := r.ReadString('\n')

	fmt.Print(hostName, " Password: ")
	bytePassword, _ := terminal.ReadPassword(int(syscall.Stdin))
	fmt.Println()
	password := string(bytePassword)
	return username, password
}
