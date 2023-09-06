package net

import "github.com/3JoB/telebot/internal/json"

type NetFrame interface{
	SetClient()
	SetJson(v json.Json)
	GETFile()
	POSTFile()
	GETJson()
	POSTJson()
}

var header = map[string]string{

}