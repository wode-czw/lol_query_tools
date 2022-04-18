package tools

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

//你们别抢我补位的accountID
//2602095920588480

//起名这么难我去  屏蔽了查找的那个刀妹
//2979015586
//4017882674  站台为你买几橘
//4135692180 豹女ID

//4009959910 大笨蛋你在里面吗

//2959958775  早冬春灬卧雪燕
//4131306349	上单摆烂王中王

//2957135251		浑身破绽拜你所賜 铂金svp 剑姬

//2939142365		个人破产

//4018117359		梦里也会想念

//Summoner_name := "牛顿有辆小车220-240"
//rank_title := "铂金111"
func Use_analyse_rank() {

	//这个版本的任务是认为这个账号一生一共打了超过30把排位。。。。的前提条件

	port_token := lcu.Get_port_token()

	var accountID int64
	accountID = 2602095920588480
	Summoner_name := "你们别抢我补位"
	rank_title := "铂金1111"
	query_name := "recent_30_rank"
	file_path := "../data/rank30_json/"
	query_command_time := time.Now().Format("2006-01-02 15:04:05 Mon Jan")[11:13] + "_" + time.Now().Format("2006-01-02 15:04:05 Mon Jan")[14:16] //hour_min
	//################################################################################################################################################################
	//################################################################################################################################################################
	//################################################################################################################################################################
	//################################################################################################################################################################
	file_name := file_path + Summoner_name + query_name + "_" + query_command_time + ".json"
	file_d_t_m_list := "../data/伤害转化率-召唤师名字-段位.json"

	var rank_30_info = make([]lcu.GameInfo, 0, 30)
	var win_num int
	var if_30 bool
	//begin_number = 0
	//end_number = 20

	rank_30_info, win_num, if_30 = Get_rank30(rank_30_info, port_token, accountID, 0, 20, 30)

	czw_champion_map := Get_champion_map("../data/champion_files/simple_champion_list.json")

	//返回的有总转化率
	all_damage_turn := Show_what_you_get(rank_30_info, czw_champion_map, port_token, accountID)

	//要写的东西都准备好了，要传进来的变量有		file_name, win_num, if_30 rank_30_info
	write_rank_file(file_name, win_num, if_30, rank_30_info)
	//#-----------------------------------------------------------------------------------------
	//统计的文件和这个人的段位   file_stat, rank_title, Summoner_name all_damage_turn
	write_statistics(file_d_t_m_list, rank_title, Summoner_name, all_damage_turn)

	fmt.Println("end")
}

//########################other function
//########################other function
//########################other function
//########################other function
//########################other function
//########################other function
//########################other function
//########################other function
//########################other function
//########################other function
//########################other function
//########################other function
//########################other function
//########################other function
//########################other function
//########################other function
//########################other function
//########################other function
//########################other function
//########################other function

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

func Body_to_struct_return_GameInfo(my_client *http.Client, my_url string) *lcu.GameInfo {
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

func Body_to_struct(my_client *http.Client, my_url string) *lcu.GameListResp {
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
	data_struct := &lcu.GameListResp{}
	json.Unmarshal([]byte(string(body)), &data_struct)
	return data_struct
}

func Get_champion_map(file_path string) map[int]string {
	champion_bytes, err := ioutil.ReadFile(file_path) //读取整个json文件

	if err != nil {

		fmt.Errorf("open file failed!")
	}

	var decoder_champion lcu.Simple_champion_list
	json.Unmarshal([]byte(champion_bytes), &decoder_champion)
	//你只是解析出来了而已，你还没有得到map[int]string

	my_champion_map := make(map[int]string, 160)
	for _, vvv := range decoder_champion.Data {
		//fmt.Printf("deal with %#v \n", kkk)
		my_champion_map[vvv.Id] = vvv.Name
		//fmt.Println(my_champion_map[vvv.Id])
	}

	return my_champion_map
}

func Show_what_you_get(rank_30_info []lcu.GameInfo, czw_champion_map map[int]string, port_token string, account_ID int64) float32 {
	var champion_number_string string
	var champion_number int

	var average_d_to_m float32
	average_d_to_m = 0
	var average_d_to_m_rate float32
	average_d_to_m_rate = 0
	for kk, v := range rank_30_info {
		fmt.Printf("==========================================================\n%2d\n", kk)
		fmt.Println("游戏时间 is ", v.GameCreationDate)
		fmt.Println(v.ParticipantIdentities[0].Player.SummonerName)
		//
		//
		//champion_number = v.Participants[0].ChampionId
		//fmt.Printf("%#v\n", v.Participants[0].ChampionId)	这句能出来就行
		champion_number_string = fmt.Sprintf("%#v", v.Participants[0].ChampionId)
		champion_number, _ = strconv.Atoi(champion_number_string)
		fmt.Printf("use champion :		===【%s】===		,游戏是否胜利：		%#v\n", czw_champion_map[champion_number], v.Participants[0].Stats.Win)

		fmt.Println("KDA is  :", v.Participants[0].Stats.Kills, "-", v.Participants[0].Stats.Deaths, "-", v.Participants[0].Stats.Assists)

		fmt.Println("its Lane is  :", v.Participants[0].Timeline.Lane)
		fmt.Println("its Role is  :", v.Participants[0].Timeline.Role)
		fmt.Println("视野得分 is  :", v.Participants[0].Stats.VisionScore)
		//fmt.Println("视野得分 is  :", v.Participants[0].Stats.VisionScore)

		if v.Participants[0].Timeline.CsDiffPerMinDeltas.Ten == 0 {
			fmt.Println("【【【【【】移动的对位不存在")

		} else {
			fmt.Printf("移动的对位差距 		is   :%-9.3f||%-9.3f||%-9.3f=======\n", v.Participants[0].Timeline.CsDiffPerMinDeltas.Ten, v.Participants[0].Timeline.CsDiffPerMinDeltas.Twenty, v.Participants[0].Timeline.CsDiffPerMinDeltas.Thirty)
		}

		if v.Participants[0].Timeline.DamageTakenDiffPerMinDeltas.Ten == 0 {
			fmt.Println("【【【【【】受到伤害的对位不存在")

		} else {
			fmt.Printf("每单位受到伤害 		is   :%-9.3f||%-9.3f||%-9.3f=======\n", v.Participants[0].Timeline.DamageTakenDiffPerMinDeltas.Ten, v.Participants[0].Timeline.DamageTakenDiffPerMinDeltas.Twenty, v.Participants[0].Timeline.DamageTakenDiffPerMinDeltas.Thirty)
		}

		if v.Participants[0].Timeline.XpDiffPerMinDeltas.Ten == 0 {
			fmt.Println("【【【【【】获得经验的对位不存在")

		} else {
			fmt.Printf("每单位经验差距 		is   :%-9.3f||%-9.3f||%-9.3f=======\n", v.Participants[0].Timeline.XpDiffPerMinDeltas.Ten, v.Participants[0].Timeline.XpDiffPerMinDeltas.Twenty, v.Participants[0].Timeline.XpDiffPerMinDeltas.Thirty)
		}

		fmt.Printf("每单位获得金币 		is   :%-9.3f||%-9.3f||%-9.3f=======\n", v.Participants[0].Timeline.GoldPerMinDeltas.Ten, v.Participants[0].Timeline.GoldPerMinDeltas.Twenty, v.Participants[0].Timeline.GoldPerMinDeltas.Thirty)
		fmt.Println("最长存活时间 is  :", v.Participants[0].Stats.LongestTimeSpentLiving)

		//------------------------

		//fmt.Println("对英雄造成的魔法伤害 is  :", v.Participants[0].Stats.MagicDamageDealtToChampions)
		fmt.Println("对英雄造成的总伤害 is  :", v.Participants[0].Stats.TotalDamageDealtToChampions)
		fmt.Println("获得的金钱 is  :", v.Participants[0].Stats.GoldEarned)

		magic_damage := float32(v.Participants[0].Stats.TotalDamageDealtToChampions)
		Gold_earn := float32(v.Participants[0].Stats.GoldEarned)
		money_to_damage := magic_damage / Gold_earn
		d_to_m_rate := Get_MTD_by_gameID_v2(v.GameId, account_ID, port_token)
		fmt.Printf("金钱转换比 is  :%2.2f\n", money_to_damage)
		fmt.Printf("比例的伤害金钱转换比 is  :%2.2f\n", d_to_m_rate)
		fmt.Println("是否提前投降-------------------------             ", v.Participants[0].Stats.GameEndedInSurrender)

		average_d_to_m = average_d_to_m + money_to_damage
		average_d_to_m_rate = average_d_to_m_rate + d_to_m_rate

	}
	fmt.Println("----------------------------------------------")
	fmt.Println("平均伤害转化率为：", average_d_to_m/30)
	fmt.Println("平均伤害转化率比例为：", average_d_to_m_rate/30)

	return average_d_to_m
}

func Get_rank30(rank_30_info []lcu.GameInfo, port_token string, accountID int64, bn, en, rank_cap int) ([]lcu.GameInfo, int, bool) {
	var begin_number int
	var end_number int
	begin_number = bn
	end_number = en

	win_number := 0

	fmt.Println("In Get rank30function cap is ", rank_cap)
	no_30 := false
	for {

		query_command := fmt.Sprintf("lol-match-history/v3/matchlist/account/%d?begIndex=%d&endIndex=%d", accountID, begin_number, end_number)
		url := port_token + query_command

		czw_client := New_client()                     //新建一个client
		data_struct := Body_to_struct(czw_client, url) //读取body对象到结构体中

		if len(data_struct.Games.Games) == 0 {
			//既然数据还没有停下来。就说明还没有到达我们的要求
			fmt.Errorf("单双排数据不足%3d把，程序终止", rank_cap)
			no_30 = true
			fmt.Println("不足30把")
			break
		}

		for _, value := range data_struct.Games.Games { //

			if value.QueueId == 420 {
				//fmt.Println("kda is ", value.Participants[0].Stats.Kills, value.Participants[0].Stats.Deaths, value.Participants[0].Stats.Assists)
				rank_30_info = append(rank_30_info, value)

				if value.Participants[0].Stats.Win == true {
					win_number = win_number + 1

				}

				if len(rank_30_info) >= rank_cap {

					break
				}

			}

		}

		if len(rank_30_info) >= rank_cap {

			break
		}

		begin_number = begin_number + 20
		end_number = end_number + 20

	}
	fmt.Println("I am len(rank_30_info):", len(rank_30_info))
	return rank_30_info, win_number, no_30

}

func write_rank_file(file_name string, win_number int, no_30 bool, rank_30_info []lcu.GameInfo) {
	fileobj, err := os.OpenFile(file_name, os.O_APPEND|os.O_CREATE, 0644)
	defer fileobj.Close()
	if err != nil {
		fmt.Println("os.OpenFile failed ,err: ", err)
		return

	}

	if no_30 == true {
		fileobj.Write([]byte("这个账号单双排数据不到30把\n"))
	}

	fileobj.Write([]byte(fmt.Sprintf("win number is %d\n\n", win_number)))
	for _, v := range rank_30_info {

		wtrite_string, _ := json.MarshalIndent(v, "", "    ")

		fileobj.Write([]byte(wtrite_string))
		fileobj.Write([]byte("\n\n"))
	}
}

func write_statistics(file_stat, rank_title, Summoner_name string, average_d_to_m float32) {
	file_d_t_m_obj, _ := os.OpenFile(file_stat, os.O_APPEND|os.O_CREATE, 0644)
	defer file_d_t_m_obj.Close()
	first_line := fmt.Sprintf("ID:%-20s			段位:%-7s		伤害转化率为%2.2f\n\n", Summoner_name, rank_title, average_d_to_m/30)
	file_d_t_m_obj.Write([]byte(first_line))
}

func Get_MTD_by_gameID_v2(gameID int64, account_ID int64, port_token string) float32 {
	czw_client := New_client()

	query_command := "lol-match-history/v1/games/" + strconv.FormatInt(gameID, 10)

	my_url := port_token + query_command

	czw_data_struct := Body_to_struct_return_GameInfo(czw_client, my_url)

	var all_damage int
	var all_money int

	for _, v := range czw_data_struct.Participants {
		all_damage = all_damage + v.Stats.TotalDamageDealtToChampions
		all_money = all_money + v.Stats.GoldEarned

	}

	var my_order int

	for key, value := range czw_data_struct.ParticipantIdentities {
		if value.Player.SummonerId == account_ID {
			my_order = key
			break
		}
		continue

	}

	my_damge := float32(czw_data_struct.Participants[my_order].Stats.TotalDamageDealtToChampions)
	my_money := float32(czw_data_struct.Participants[my_order].Stats.GoldEarned)

	//accountID 应该是这个

	all_money_f32 := float32(all_money)
	all_damage_f32 := float32(all_damage)

	my_damage_rate := my_damge / all_damage_f32
	my_money_rate := my_money / all_money_f32

	var MTD_f32 float32
	MTD_f32 = my_damage_rate / my_money_rate
	//fmt.Println("顺序为", my_order)
	//fmt.Println("总伤害为为", my_damge)
	//fmt.Println("总金钱为", my_money)
	//fmt.Println("伤害转化率为", MTD_f32)
	return MTD_f32

}
