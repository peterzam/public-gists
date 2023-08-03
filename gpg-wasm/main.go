package main

import (
	"log"
	"syscall/js"

	"github.com/ProtonMail/gopenpgp/v2/helper"
)

func main() {
	done := make(chan struct{}, 0)
	js.Global().Set("encrypt", js.FuncOf(encrypt))
	<-done
}

func encrypt(this js.Value, args []js.Value) interface{} {
	key := args[0].String()
	message := args[1].String()
	armor, err := helper.EncryptMessageArmored(key, message)
	if err != nil {
		log.Fatal(err)
	}
	return armor
}
