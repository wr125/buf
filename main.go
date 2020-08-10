package main

import (
	"bytes"
	"encoding/json"
	"io"
	"log"
	"os"
)

type User struct {
	Name  string
	Email string
	Age   int
}

func main() {
	user := User{Name: "Jackson", Email: "jackson@gmail.com", Age: 34}

	f, err := os.Create("db.json")

	if err != nil {
		log.Fatal(err)
	}
	defer f.Close()
	buf := new(bytes.Buffer)
	json.NewEncoder(f).Encode(user)

	io.Copy(buf, f)
}
