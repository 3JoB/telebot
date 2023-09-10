package net

import "github.com/3JoB/telebot/internal/json"

type FastHTTPRequest struct {
	json json.Json
}

func (f *FastHTTPRequest) Reset() {

}

func (f *FastHTTPRequest) Release() {
	f.Reset()
	requestPool.Put(f)
}