package main

import (
	"crypto/rand"
	"flag"
	"fmt"
	"os"
)

func main() {
	username := flag.String("user", "FR4NC1SC011", "The username")
	website := flag.String("website", "Nop", "The website")
	size := flag.Int("size", 12, "Len of the password")

	flag.Parse()

	homePath, err := os.UserHomeDir()
	check(err)

	pass_file, err := os.OpenFile(homePath+"/.passwords.txt", os.O_APPEND|os.O_WRONLY, 0644)
	check(err)

	web := "Website: " + *website
	usr := "Username: " + *username
	psw := "Password: " + randPass(*size)
	nl := "\n"

	d := []string{web, usr, psw, nl}

	for _, v := range d {
		fmt.Fprintln(pass_file, v)
		check(err)
	}

	err = pass_file.Close()
	check(err)
	fmt.Println(web)
	fmt.Println(usr)
	fmt.Println(psw)

}

func GenerateRandomBytes(n int) ([]byte, error) {
	b := make([]byte, n)
	_, err := rand.Read(b)
	if err != nil {
		return nil, err
	}
	return b, nil
}

func randPass(size int) string {
	var password string
	password = `(){}[]<>/+=-_,.~€#!@£$%^&*?|0123456789abcdefghijklmnopqrstuvwxyzABCDEFGHIJKLMNOPQRSTUVWXYZ`
	//var bytes = make([]byte, size)
	var bytes, err = GenerateRandomBytes(size)
	check(err)

	rand.Read(bytes)
	for k, v := range bytes {
		bytes[k] = password[v%byte(len(password))]
	}
	return string(bytes)
}

func check(e error) {
	if e != nil {
		panic(e)
	}
}
