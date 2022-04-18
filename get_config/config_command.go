package get_config

type Query_config struct {
	Begin_number                       int
	End_number                         int
	Account_name                       string
	Good_champion_use_number_more_than int
	Rank_number                        int
}

//account_name := "你们别抢我补位"
//account_name := "个人破产"
//account_name := "你霸气吗"

func Return_query_config() Query_config {
	var Czw_config Query_config

	Czw_config.Begin_number = 0
	Czw_config.End_number = 20
	Czw_config.Account_name = "你们别抢我补位"
	Czw_config.Good_champion_use_number_more_than = 4
	Czw_config.Rank_number = 50 //默认30最好，我这里做个能够修改的版本
	return Czw_config
}
