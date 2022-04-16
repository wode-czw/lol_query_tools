package little_function

import (
	"czw_lol_query_tools/Only_step"
	"czw_lol_query_tools/lcu"
	"czw_lol_query_tools/models"
	"czw_lol_query_tools/tools"
	"fmt"
	"strconv"
)

//返回一个map[string]struct

/*
英雄名字 {
	30把使用的数量  		int
	kda_list	[]struct{
		kda		string
		MTD_rate		float32
		MTD				float32
	}
}


*/

type Data_in_game struct {
	Champion_ID   models.Champion
	Champion_name string
	KDA           string
	MTD_rate      float32
	MTD_number    float32
}

type Good_at_champion_data struct {
	Use_number        int
	Champion_ID       models.Champion
	Champion_name     string
	Data_in_game_list []Data_in_game
}

func Test_Get_Data(account_ID int64, gameID int64, port_token string, champion_map map[int]string) {

	czw_game_info := Get_gameInfo(gameID, port_token)
	champion_number, champion_data := Get_Data_in_game_from_rank30(account_ID, czw_game_info, port_token, champion_map)

	fmt.Println(champion_number)
	fmt.Println(champion_data.KDA)
	fmt.Println(champion_data.Champion_name)
}

//这里的rank30可不是能够直接一口气得到的，因为matchlist返回的是lcu.GameListResp
func Get_Good_at_champion_data_from_rank30(account_ID int64, rank_30_info []lcu.GameInfo, port_token string, the_champion_map map[int]string) map[string]Good_at_champion_data {
	var champion_list map[string]Good_at_champion_data

	champion_list = make(map[string]Good_at_champion_data, 30)

	//#########
	var return_game_champion_name string

	var Game_in_this_data Data_in_game
	for _, v := range rank_30_info {
		//先拿到这把的英雄名字和数据
		return_game_champion_name, Game_in_this_data = Get_Data_in_game_from_rank30(account_ID, &v, port_token, the_champion_map)

		the_champion_list, if_exist := champion_list[return_game_champion_name]

		if if_exist == false { //第一次统计这个英雄
			var the_champion_statices Good_at_champion_data
			the_champion_statices.Use_number = 0
			the_champion_statices.Champion_ID = Game_in_this_data.Champion_ID
			the_champion_statices.Champion_name = return_game_champion_name
			the_champion_statices.Data_in_game_list = append(the_champion_statices.Data_in_game_list, Game_in_this_data)

			champion_list[return_game_champion_name] = the_champion_statices

		} else { //存在这个英雄
			the_champion_list.Use_number = the_champion_list.Use_number + 1
			the_champion_list.Data_in_game_list = append(the_champion_list.Data_in_game_list, Game_in_this_data)
			champion_list[return_game_champion_name] = the_champion_list
		}

	}

	return champion_list

}

func Get_Data_in_game_from_rank30(account_ID int64, game_info *lcu.GameInfo, port_token string, champion_map map[int]string) (string, Data_in_game) {
	var this_Data Data_in_game
	var order int

	for _, v := range game_info.ParticipantIdentities {
		if account_ID == v.Player.SummonerId {
			order = v.ParticipantId
		}
	}

	this_Data.Champion_ID = game_info.Participants[order].ChampionId
	var c_int int
	c_int = int(game_info.Participants[order].ChampionId)
	this_Data.Champion_name = champion_map[c_int]
	kill_number := game_info.Participants[order].Stats.Kills
	Deaths_number := game_info.Participants[order].Stats.Deaths
	Assists_number := game_info.Participants[order].Stats.Assists
	this_Data.KDA = strconv.Itoa(kill_number) + "--" + strconv.Itoa(Deaths_number) + "--" + strconv.Itoa(Assists_number)

	d_to_m_rate, d_to_m := Only_step.Get_MTD_by_gameID_v2(game_info.GameId, account_ID, port_token)
	this_Data.MTD_rate = d_to_m_rate
	this_Data.MTD_number = d_to_m

	return this_Data.Champion_name, this_Data
}

func Get_gameInfo(gameID int64, port_token string) *lcu.GameInfo {
	czw_client := tools.New_client()
	query_command := "lol-match-history/v1/games/" + strconv.FormatInt(gameID, 10)
	my_url := port_token + query_command

	var czw_data_struct *lcu.GameInfo
	czw_data_struct = tools.Body_to_struct_return_GameInfo(czw_client, my_url) //*lcu.GameInfo
	return czw_data_struct
}
