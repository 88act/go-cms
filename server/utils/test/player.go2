package internal

import (
	"fmt"
	"time"

	"github.com/name5566/leaf/gate"
	g "github.com/name5566/leaf/go"
	"github.com/name5566/leaf/log"
	"github.com/name5566/leaf/timer"
)

type Player struct {
	sockID      string // 远程地址 id
	playerData  PlayerData
	agent       gate.Agent  //玩家 agent
	msgChan     chan string //消息
	optChan     chan string //控制
	sendHearNum int64
	onlineState int          //当前状态 0 未登录 1 在线 2离线
	saveDBTimer *timer.Timer //自动保存数据 定时器
	*g.LinearContext
}

func (player *Player) init() {
	// timer := time.NewTimer(1 * time.Second)
	// for {
	// 	select {
	// 	case <-timer.C:
	// 		sendHear(player)
	// 		timer.Reset(3 * time.Second)
	// 		//timer.Reset(2 * time.Second)
	// 		// 若要停止定时器就使用如下代码，这样会释放资源
	// 		//timer.Stop()
	// 	}
	// 	//break
	// }

	fmt.Println(" ===== player 初始化  完成 ================== ")

	player.msgChan = make(chan string, 10)
	player.optChan = make(chan string, 10)
	//go ClientStart_do(player)

}

func (player *Player) autoSaveDB() {

	//skeleton.AfterFunc(duration, func() {
	fmt.Println("  autoSaveDB 1==== ====== ", player.playerData.role.RoleName)
	//data := util.DeepClone(player.playerData)
	//_ = data.(*PlayerData).saveValue()
	//player.autoSaveDB()

	// 	// fmt.Println(data)
	// 	// player.Go(func() {
	// 	// 	fmt.Println(" skeleton.AfterFunc 2=== ====== ")
	// 	// 	_ = data.(*PlayerData).saveValue()

	// 	// }, func() {
	// 	// 	fmt.Println(" skeleton.AfterFunc 2 ======= ==== ", player.playerData.role.RoleName)
	// 	// 	player.autoSaveDB()
	// 	// })
	// })
}

func ClientStart_do(player *Player) {
	log.Debug(" 注册 Player  ClientStart_do  run.... sockID= ", player.sockID, "  RoleName =", player.playerData.role.RoleName)
	fmt.Println(" 注册 Player   ClientStart_do  run....sockID=  ", player.sockID, "  RoleName =", player.playerData.role.RoleName)
	time.Sleep(time.Second)
	for {
		select {

		//  消息
		case msg := <-player.msgChan:
			if msg == "sendHear" {
				player.sendHearNum++
			}
			log.Debug(" 玩家收到消息  msg = ", msg)
			fmt.Println("玩家收到消息  msg = ", msg)
		// 注销
		case opt := <-player.optChan:
			log.Debug(" 玩家收到消息  msg = ", opt)
			fmt.Println("玩家收到消息  msg = ", opt)
			//停止玩家进程
			if opt == "exitGame" {

				//fmt.Println("跳出for循环")
				//break Loop

				fmt.Println("玩家Player  跳出for循环")
				goto Loop
			}
		}
	}
Loop:
	fmt.Println("玩家Player 结束进程...", player.playerData.role.RoleName)
}

/*  创建新func 并带上参数
func sendHear(player *Player) {
	player.sendHearNum++
	println("sendHear .... sendHearNum =", player.sendHearNum, player.role.RoleName)
	retBuf := &msg.SCHeartBeat{
		ServerTime: player.sendHearNum,
	}
	player.agent.WriteMsg(retBuf)

	f := newFunc(player)
	_ = time.AfterFunc(3*time.Second, f)

	//time.AfterFunc(4*time.Second, sendHear)
}

func newFunc(player *Player) func() {
	//	fmt.Printf("creating func with '%s' ", bar)
	return func() {
		sendHear(player)
	}
}
*/
