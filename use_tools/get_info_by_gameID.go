package main

import (
	"crypto/tls"
	"czw_lol_query_tools/lcu"
	"encoding/json"
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
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
	my_url := port_token + query_command
	czw_data_struct := body_to_struct(czw_client, my_url)

	resp2, _ := json.MarshalIndent(czw_data_struct, "", "    ")
	file_path := "../data/gameID_info/"
	file_name := time.Now().Format("2006-01-02 15:04:05 Mon Jan")[11:13] + "_" + time.Now().Format("2006-01-02 15:04:05 Mon Jan")[14:16] + "_gameID_" + strconv.Itoa(gameID) + ".json" //hour_min
	fileobj, _ := os.OpenFile(file_path+file_name, os.O_APPEND|os.O_CREATE, 0644)
	fileobj.Write([]byte(resp2))

}

func New_client() *http.Client {
	tr := &http.Transport{
		ForceAttemptHTTP2: true,
		TLSClientConfig:   &tls.Config{InsecureSkipVerify: true},
	}

	//初始化一个客户端的实例
	//所以这里这个大佬的做法就是构建一个支持H2的客户端，因为默认的
	client := &http.Client{Transport: tr}
	return client
}

func body_to_struct(my_client *http.Client, my_url string) *lcu.GameInfo {
	resp, err := my_client.Get(my_url)
	if err != nil {
		fmt.Println("lcu 通信失败", err)
		log.Fatal("bug")
	}

	defer resp.Body.Close()
	body, err := ioutil.ReadAll(resp.Body) //body是一个json对象
	if err != nil {

		fmt.Println("数据获取失败")
	}

	//解析json到结构体中
	data_struct := &lcu.GameInfo{}
	json.Unmarshal([]byte(string(body)), &data_struct)
	return data_struct
}
