package main

import (
	crand "crypto/rand"
	"flag"
	"fmt"
	"math"
	"math/big"
	"os"
	"strconv"
	"sync"
)

var exitCode = 0

const letters = "abcdefghijkmnopqrstuvwxyzABCDEFGHJKLMNPQRSTUVWXYZ1234567890"

func main() {
	fmt.Println()
	pwcMain()
	os.Exit(exitCode)
}

func pwcMain() {
	//get command line arguments(flags)
	flag.Parse()
	pwLength := getIntFlag(0, 1, "パスワード長が不正です。")
	pwQuantity := getIntFlag(1, 2, "パスワード生成数が不正です。")

	//check max quantity
	maxQuantity := uint(math.Pow(float64(len(letters)), float64(pwLength)))
	if uint(pwQuantity) > maxQuantity {
		exitCode = 3
		fmt.Printf("パスワード生成数が同値パスワードを生成してしまう値です。%d文字での最大生成数は%dです。\n", pwLength, maxQuantity)
		os.Exit(exitCode)
	}

	//generate password
	pwChan := make(chan string)
	wg := new(sync.WaitGroup)

	for i := 0; i < pwQuantity; i++ {
		wg.Add(1)
		//write channel
		go func() {
			defer wg.Done()
			generatePassword(pwLength, pwChan)
		}()
	}

	//close channel
	go func() {
		wg.Wait()
		close(pwChan)
	}()

	//read channel & output
	var passwords, receives []string

	for receive := range pwChan {
		receives = append(receives, receive)
	}

receivesLoop:
	for i := 0; i < len(receives); i++ {
		switch len(passwords) {
		case 0:
			passwords = append(passwords, receives[i])
			fmt.Println("0", receives[i])

		default:
			for idx, pw := range passwords {
				if pw == receives[i] {
					receives = append(receives, generatePassword(pwLength, nil))
					continue receivesLoop
				} else if len(passwords) == idx+1 {
					passwords = append(passwords, receives[i])
					fmt.Println(idx+1, receives[i])
				}
			}
		}
	}

	return
}

func getIntFlag(flagIndex, errCode int, errMsg string) int {
	retFlag, err := strconv.Atoi(flag.Arg(flagIndex))
	if err != nil {
		exitCode = errCode
		fmt.Println(errMsg)
		os.Exit(exitCode)
	}
	return retFlag
}

func generatePassword(length int, c chan string) string {
	password := ""
	for i := 0; i < length; i++ {
		password += string(letters[generateCryptoRand()])
	}
	if c != nil {
		c <- password
	}
	return password
}

func generateCryptoRand() int64 {
	nBig, err := crand.Int(crand.Reader, big.NewInt(int64(len(letters))))
	if err != nil {
		exitCode = 4
		fmt.Println("パスワード生成に失敗しました。")
		os.Exit(exitCode)
	}
	return nBig.Int64()
}
