package tools

import (
	"log"
	"testing"
)

type PublicParam struct {
	MsgType     string `form:"msg_type"`
	MsgId       string `form:"msg_id"`
	ProviderId  string `form:"logistic_provider_id"`
	FromCode    string `form:"from_code"`
	ToCode      string `form:"to_code"`
	PartnerCode string `form:"partner_code"`
	Digest      string `form:"data_digest"`
	Interface   string `form:"logistics_interface"`
}

func TestFunc(t *testing.T) {
	s := PublicParam{
		MsgType:     "123",
		MsgId:       "234",
		ProviderId:  "345",
		FromCode:    "456",
		ToCode:      "567",
		PartnerCode: "678",
		Digest:      "789",
		Interface:   "890",
	}

	sss := Struct2Param(&s, "form", nil, true)
	log.Printf("--dsa111-----%v", sss.String())
	ss := GenerateSignString(&s, "form", "", "", nil, true, true)
	log.Printf("--dsa222-----%v", ss)
}
