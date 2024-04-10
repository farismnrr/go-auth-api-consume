package helper

import (
	"bufio"
	"crypto/md5"
	"encoding/hex"
	"encoding/json"
	"fmt"
	"io/ioutil"
	"log"
	"os"
	"os/exec"
	"runtime"
	"strings"
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

func ReadEnvFile(filename string) (map[string]string, error) {
	envVars := make(map[string]string)

	file, err := os.Open(filename)
	if err != nil {
		return nil, err
	}
	defer file.Close()

	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		line := scanner.Text()
		parts := strings.SplitN(line, "=", 2)
		if len(parts) == 2 {
			envVars[parts[0]] = parts[1]
		}
	}

	if err := scanner.Err(); err != nil {
		return nil, err
	}

	return envVars, nil
}
