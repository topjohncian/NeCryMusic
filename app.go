package main

import (
	"github.com/Tnze/CoolQ-Golang-SDK/cqp"
	"github.com/topjohncian/NeCryMusic/command"
)

//go:generate cqcfg -c .
// cqp: 名称: 网抑云音乐语句获取
// cqp: 版本: 1.0.0:1
// cqp: 作者: topjohncian
// cqp: 简介: 网抑云音乐语录 感谢API:https://cqp.cc/t/51719
func main() { /*此处应当留空*/ }

func init() {
	cqp.AppID = "com.github.topjohncian.necrymusic"
	cqp.PrivateMsg = onPrivateMsg
	cqp.GroupMsg = onGroupMsg
}

func onPrivateMsg(subType, msgID int32, fromQQ int64, msg string, font int32) int32 {
	if msg == "网抑云" {
		cqp.SendPrivateMsg(fromQQ, command.GetOne())
	}

	return 0
}

func onGroupMsg(subType, msgID int32, fromGroup, fromQQ int64, fromAnonymous, msg string, font int32) int32 {
	if msg == "网抑云" {
		cqp.SendGroupMsg(fromGroup, command.GetOne())
	}

	return 0
}
