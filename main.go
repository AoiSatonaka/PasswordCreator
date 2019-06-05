package main

import (
  crand "crypto/rand"
  "flag"
  "fmt"
  "math/big"
  "os"
  "strconv"
  "sync"
)

var exitCode = 0

const letters = "abcdefghijkmnopqrstuvwxyzABCDEFGHJKLMNPQRSTUVWXYZ1234567890"

func main()  {
	pwcMain()
	os.Exit(exitCode)
}

func pwcMain(){
	//get command line aruguments
	flag.Parse()

	pwLength := getIntFlag(0,1,"パスワード長が不正です。")
	pwQuantity := getIntFlag(1,2,"パスワード生成数が不正です。")

	fmt.Println(pwLength,pwQuantity)

	//create password
	pwChan := make(chan string,pwQuantity)
  wg := new(sync.WaitGroup)

	for i:=0;i<pwQuantity;i++{
		wg.Add(1)
		//channel writer goroutine
	  go func() {
	    defer wg.Done()
      generatePassword(pwLength,pwChan)
    }()
	}

	//close channel goroutine
	go func() {
	  wg.Wait()
	  close(pwChan)
  }()

	//channel reader
	for receive := range pwChan{
	  println("length=",receive)
	}
	println(len(letters))
	return
}

func getIntFlag(flagIndex,errCode int,errMsg string) int {
	retFlag,err := strconv.Atoi(flag.Arg(flagIndex))
	if err!=nil {
		exitCode = errCode
		fmt.Println(errMsg)
		os.Exit(exitCode)
	}
	return retFlag
}

func generatePassword(length int, c chan string){
  password := ""
  for i:=0; i<length; i++ {
    password += string(letters[generateCryptoRand()])
  }
  c <- password
}

func generateCryptoRand()int64{
  nBig,err := crand.Int(crand.Reader,big.NewInt(int64(len(letters))))
  if err != nil {
    exitCode = 3
    fmt.Println("パスワード生成に失敗しました。")
    os.Exit(exitCode)
  }
  return nBig.Int64()
}