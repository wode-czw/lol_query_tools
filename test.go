package main

import (
	"czw_lol_query_tools/get_port_token"
	"czw_lol_query_tools/tools"
)

func main() {
	tools.Get_MTD_by_gameID_v1(6147811485, get_port_token.Return_port_token())

}
