package main

import (
	"czw_lol_query_tools/get_port_token"
	"czw_lol_query_tools/lcu"
	"czw_lol_query_tools/little_function"
	"czw_lol_query_tools/tools"
	"fmt"
)

func main() {
	fmt.Println("Hi")
	var account_ID int64

	//account_name := "你们别抢我补位"
	//account_name := "个人破产"
	//account_name := "你霸气吗"
	account_name := "荒火流霜"
	fmt.Println("I am", account_name)
	czw_port_token := get_port_token.Return_port_token() + "/"

	ID_info := tools.Return_getAccountID_by_summonerName(account_name)

	account_ID = ID_info.SummonerId

	//czw_port_token := lcu.Get_port_token()

	czw_champion_map := tools.Get_champion_map("./data/champion_files/simple_champion_list.json")
	var rank_30_data = make([]lcu.GameInfo, 0, 30)
	var win_num int
	var no_30 bool
	rank_30_data, win_num, no_30 = tools.Get_rank30(rank_30_data, czw_port_token, account_ID, 0, 20)

	fmt.Println("win number is ", win_num)

	var adverage_good_champion_MTD, use_time float32
	adverage_good_champion_MTD = 0
	use_time = 0

	if !no_30 {
		the_data := little_function.Get_Good_at_champion_data_from_rank30(account_ID, rank_30_data, czw_port_token, czw_champion_map)

		for _, vvv := range the_data {
			if len(vvv.Data_in_game_list) < 0 {
				continue
			}

			fmt.Println("=============================================")
			fmt.Println(vvv.Champion_name)

			for _, vvvv := range vvv.Data_in_game_list {

				fmt.Printf("KDA 是 %-10s,数值伤害转化率为 %-5.2f,		比例伤害转化率为%-5.2f \n", vvvv.KDA, vvvv.MTD_number, vvvv.MTD_rate)
				adverage_good_champion_MTD = adverage_good_champion_MTD + vvvv.MTD_number
				use_time = use_time + 1
			}

		}
	}
	fmt.Printf("擅长英雄的平均伤害转化率为 %-5.2f", adverage_good_champion_MTD/use_time)
}
