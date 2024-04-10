package helper

import (
	"crypto/md5"
	"encoding/hex"
	"encoding/json"
	"fmt"
	"io/ioutil"
	"log"
	"os"
	"os/exec"
	"runtime"
	"time"

	"github.com/farismnrr/go-auth-api-consume/model"
)

func ClearScreen() {
	osName := runtime.GOOS

	switch osName {
	case "linux", "darwin": // Untuk Linux dan MacOS
		cmd := exec.Command("clear")
		cmd.Stdout = os.Stdout
		cmd.Run()
	case "windows": // Untuk Windows 10
		cmd := exec.Command("cmd", "/c", "cls")
		cmd.Stdout = os.Stdout
		cmd.Run()
	default:
		fmt.Println("Clear screen tidak didukung pada sistem operasi ini")
	}
}

func Delay(duration int) {
	for i := duration; i >= 1; i-- {
		fmt.Printf("\r%d seconds...", i)
		time.Sleep(1 * time.Second)
	}
	fmt.Print("\r")
}

func GenerateHash(text string) string {
	hasher := md5.New()
	hasher.Write([]byte(text))
	return hex.EncodeToString(hasher.Sum(nil))
}

func ReadJsonFile() (string, string) {
	authData, err := ioutil.ReadFile("Authorization.json")
	if err != nil {
		log.Fatal("Please insert the authorization file!")
	}

	var auth model.AuthorizationData
	json.Unmarshal(authData, &auth)

	return auth.Username, auth.PrivateKey
}
