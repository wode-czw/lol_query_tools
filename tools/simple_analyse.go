package tools

import (
	"czw_lol_query_tools/lcu"
	"fmt"
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
func Use_simple_analyse() {

	//这个版本的任务是认为这个账号一生一共打了超过30把排位。。。。的前提条件

	port_token := lcu.Get_port_token()

	var accountID int64
	accountID = 2602095920588480
	Summoner_name := "你们别抢我补位"
	rank_title := "钻1"
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

	rank_30_info, win_num, if_30 = Get_rank30(rank_30_info, port_token, accountID, 0, 20)

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

func Simpile_show_what_you_get(rank_30_info []lcu.GameInfo, czw_champion_map map[int]string, port_token string, account_ID int64) float32 {
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
		/*
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
		*/
		//fmt.Printf("每单位获得金币 		is   :%-9.3f||%-9.3f||%-9.3f=======\n", v.Participants[0].Timeline.GoldPerMinDeltas.Ten, v.Participants[0].Timeline.GoldPerMinDeltas.Twenty, v.Participants[0].Timeline.GoldPerMinDeltas.Thirty)
		//fmt.Println("最长存活时间 is  :", v.Participants[0].Stats.LongestTimeSpentLiving)

		//------------------------

		//fmt.Println("对英雄造成的魔法伤害 is  :", v.Participants[0].Stats.MagicDamageDealtToChampions)
		//fmt.Println("对英雄造成的总伤害 is  :", v.Participants[0].Stats.TotalDamageDealtToChampions)
		//fmt.Println("获得的金钱 is  :", v.Participants[0].Stats.GoldEarned)

		magic_damage := float32(v.Participants[0].Stats.TotalDamageDealtToChampions)
		Gold_earn := float32(v.Participants[0].Stats.GoldEarned)
		money_to_damage := magic_damage / Gold_earn
		d_to_m_rate := Get_MTD_by_gameID_v2(v.GameId, account_ID, port_token)
		fmt.Printf("金钱转换比 is  :%2.2f\n", money_to_damage)
		fmt.Printf("比例的伤害金钱转换比 is  :%2.2f\n", d_to_m_rate)
		//fmt.Println("是否提前投降-------------------------             ", v.Participants[0].Stats.GameEndedInSurrender)

		average_d_to_m = average_d_to_m + money_to_damage
		average_d_to_m_rate = average_d_to_m_rate + d_to_m_rate

	}
	fmt.Println("----------------------------------------------")
	fmt.Println("平均伤害转化率为：", average_d_to_m/30)
	fmt.Println("平均伤害转化率比例为：", average_d_to_m_rate/30)

	return average_d_to_m
}
