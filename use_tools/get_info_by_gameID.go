package main

import (
	"encoding/json"
	"fmt"
	"os"
	"strconv"
	"time"
)

//func main() {
//	Get_MTD_by_gameID_v2(6847919402, 4046768662, lcu.Get_port_token())
//}

//拿到结构体并写到文件里
func Get_MTD_by_gameID_v1(gameID int, port_token string) {
	czw_client := New_client()
	query_command := "lol-match-history/v1/games/" + strconv.Itoa(gameID)
	my_url := port_token + "/" + query_command
	fmt.Println(my_url)
	czw_data_struct := body_to_struct(czw_client, my_url)

	resp2, _ := json.MarshalIndent(czw_data_struct, "", "    ")
	file_path := "../data/gameID_info/"
	file_name := time.Now().Format("2006-01-02 15:04:05 Mon Jan")[11:13] + "_" + time.Now().Format("2006-01-02 15:04:05 Mon Jan")[14:16] + "_gameID_" + strconv.Itoa(gameID) + ".json" //hour_min
	fileobj, _ := os.OpenFile(file_path+file_name, os.O_APPEND|os.O_CREATE, 0644)
	fileobj.Write([]byte(resp2))

}
