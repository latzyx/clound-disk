package models

import (
	"bytes"
	"encoding/json"
	"log"
)

func Select() string {
	data := make([]*UserBasic, 0)
	err := Engine.Find(&data)
	if err != nil {
		log.Println("Get UserBasic Error:", err)
	}
	b, err := json.Marshal(data)
	if err != nil {
		log.Println("Marshal UserBasic Error:", err)
	}
	dst := new(bytes.Buffer)
	err = json.Indent(dst, b, "", "")
	if err != nil {
		log.Println("Indent UserBasic Error:", err)
	}
	return dst.String()
}
