package main

import "github.com/Tnze/CoolQ-Golang-SDK/cqp"

//go:generate cqcfg -c .
// cqp: 名称: 网抑云音乐
// cqp: 版本: 1.0.0:1
// cqp: 作者: topjohncian
// cqp: 简介: 网抑云音乐语录 感谢API:https://cqp.cc/t/51719
func main() { /*此处应当留空*/ }

func init() {
	cqp.AppID = "com.github.topjohncian.necrymusic" // TODO: 修改为这个插件的ID
	cqp.PrivateMsg = onPrivateMsg
}

func onPrivateMsg(subType, msgID int32, fromQQ int64, msg string, font int32) int32 {
	cqp.SendPrivateMsg(fromQQ, msg) //复读机
	return 0
}
