package models

import "log"

type jsonErr struct {
	StatusCode int    `json:"status_code"`
	Text       string `json:"text"`
}

func checkError(err error) {
	if err != nil {
		log.Fatal(err)
	}
}
