package main

import (
  "flag"
  "fmt"
  "os"
  "strconv"
)

var exitCode = 0;

const(
  letters = "abcdefghijkmnopqrstuvwxyzABCDEFGHJKLMNPQRSTUVWXYZ1234567890"
)

func main()  {
  pwcMain()
  os.Exit(exitCode)
}

func pwcMain(){
  //get command line aruguments
  flag.Parse()

  pwLength := getFlag(0,1,"パスワード長が不正です。")
  pwQuantity := getFlag(1,2,"パスワード生成数が不正です。")

  fmt.Println(pwLength,pwQuantity)

  //create password
  pwChan := make(chan int)

  for i:=0;i<pwQuantity;i++{
     generatePassword(i,pwChan)
  }

  //for receive := range pwChan{
  // println("length=",receive)
  //}
close(pwChan)
  return
}

func getFlag(flagIndex,errCode int,errMsg string) int {
  retFlag,err := strconv.Atoi(flag.Arg(flagIndex))
  if err!=nil {
    exitCode = errCode
    fmt.Println(errMsg)
    os.Exit(exitCode)
  }
  return retFlag
}

func generatePassword(length int, c chan int){
    c <- length
    fmt.Println("length",length)
}