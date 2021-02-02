package handler

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"os"
)

func (h *Handler) jsonInit() {
	file, err := os.OpenFile("users.json", os.O_RDONLY|os.O_CREATE, 0777)
	if err != nil {
		fmt.Println("error: ", err)
	}
	defer file.Close()

	jsonData, _ := ioutil.ReadAll(file)
	json.Unmarshal(jsonData, &h.Users)

	for _, u := range h.Users {
		if h.Seq <= u.ID {
			h.Seq = u.ID + 1
		}
	}
}

func (h *Handler) jsonUpdate() error {
	file, err := os.OpenFile("users.json", os.O_WRONLY, 0777)
	if err != nil {
		return err
	}
	defer file.Close()

	jsonData, err := json.Marshal(&h.Users)
	if err != nil {
		return err
	}
	file.Truncate(0)
	file.Write(jsonData) //Не смог придумать ничего лучше, чем полностью переписывать файл
	return nil
}
