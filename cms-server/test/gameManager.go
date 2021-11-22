package internal

import (
	"fmt"

	"sync"

	"github.com/name5566/leaf/log"
)

// GameManager 所有 websocket 信息
type GameManager struct {
	systemRoom  map[int]string // 系统房间
	userRoom    map[int]string // 用户房间
	Lock        sync.Mutex
	userList    map[string]*Player //用户
	clientCount int                //在线人数
	userAddChan chan *Player       //用户
	userDelChan chan string        //用户离线
}

// 确保 bin/gamedata 目录中存在 Test.txt 文件
// 文件名必须和此结构体名称相同（大小写敏感）
// 结构体的一个实例映射 recordfile 中的一行
// 读取 recordfile Test.txt 到内存中
// RfTest 即为 Test.txt 的内存镜像

func (manager *GameManager) start() {
	// 	// rf, err := recordfile.New(Test{})
	// 	// if err != nil {
	// 	// 	//return
	// 	// }
	// 	// // 按索引查找
	// 	// // 获取 Test.txt 中 Id 为 1 的那一行
	// 	// err = rf.Read("test.txt")
	// 	// if err != nil {
	// 	// 	//return
	// 	// }
	// 	// for i := 0; i < rf.NumRecord(); i++ {
	// 	// 	r := rf.Record(i).(*Test)
	// 	// 	fmt.Println(r.Str)
	// 	// }

	// 	// r := rf.Index(2).(*Test)
	// 	// fmt.Println(r.Str)
	// 	// if r != nil {
	// 	// 	row := r.(*Test)
	// 	// 	// 输出此行的所有列的数据
	// 	// 	log.Debug("%v %v %v", row.Id, row.Arr, row.Str)
	// 	// }
	// 	//go start_do(manager)
	manager.systemRoom = make(map[int]string)
	manager.userRoom = make(map[int]string)
	manager.userList = make(map[string]*Player)
	manager.userAddChan = make(chan *Player, 128)
	manager.userDelChan = make(chan string, 128)
	// skeleton.Go(func() {
	// 	start_do(manager)
	// }, func() {
	// 	fmt.Println("start_do(manager)  callback")
	// })
	go start_do(manager)
}
func start_do(manager *GameManager) {
	log.Debug("gameManager 开始运行 start_do() ==========")
	//fmt.Print(manager)
	fmt.Println("gameManager 开始运行 start_do() ==========")
	for {
		select {
		// 加入注册
		case client := <-manager.userAddChan:
			log.Debug("GameManager 注册 client [%s] connect ", client.sockID)
			fmt.Println("GameManager 注册 client [%s] connect ")
			fmt.Println(client.sockID)
			//fmt.Print("GameManager 注册 client [%s] connect ", client.sockID)
			// log.Debug("register client [%s] to group [%s]", client.Id, client.Group)

			manager.Lock.Lock()
			manager.userList[client.sockID] = client
			//client.ClientStart()
			// if manager.Group[client.Group] == nil {
			// 	manager.Group[client.Group] = make(map[string]*Client)
			// 	manager.groupCount += 1
			// }
			// manager.Group[client.Group][client.Id] = client
			manager.clientCount += 1
			//fmt.Println("manager.userList-----------")
			//fmt.Println(manager.userList)
			manager.Lock.Unlock()
			fmt.Println(" 在线人数 = ", manager.clientCount)

		// 注销
		case sockID := <-manager.userDelChan:
			log.Debug("GameManager client  注销  sockID = ", sockID)
			//停止玩家进程
			myPlayer := manager.userList[sockID]
			manager.Lock.Lock()
			delete(manager.userList, sockID)
			manager.clientCount -= 1
			manager.Lock.Unlock()
			myPlayer.optChan <- "exitGame"
			fmt.Println("GameManager client [%s] 注销 userDelChan [%s]", myPlayer.sockID, myPlayer.playerData.role.GetRoleName)
			fmt.Println(" 在线人数 = ", manager.clientCount)
			// 发送广播数据到某个组的 channel 变量 Send 中
			//case data := <-manager.boardCast:
			//	if groupMap, ok := manager.wsGroup[data.GroupId]; ok {
			//		for _, conn := range groupMap {
			//			conn.Send <- data.Data
			//		}
			//	}
		}
	}
}
