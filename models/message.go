package models

import (
	"encoding/json"
	"fmt"
	"github.com/gorilla/websocket"
	"gopkg.in/fatih/set.v0"
	"gorm.io/gorm"
	"net"
	"net/http"
	"strconv"
	"sync"
	"time"
)

type UserMessage struct {
	gorm.Model
	UserID     int64  //发送者
	TargetID   int64  //接受者
	SendType   int    //发送类型  1私聊  2群聊  3心跳
	MsgType    int    //消息类型  1文字 2表情包 3语音 4图片 /表情包
	Content    string //消息内容
	CreateTime uint64 //创建时间
	ReadTime   uint64 //读取时间
	Pic        string
	Url        string
	Desc       string
	Amount     int //其他数字统计
}

func (table *UserMessage) TableName() string {
	return "user_message"
}

type Node struct {
	Conn          *websocket.Conn //连接
	Addr          string          //客户端地址
	FirstTime     uint64          //首次连接时间
	HeartbeatTime uint64          //心跳时间
	LoginTime     uint64          //登录时间
	DataQueue     chan []byte     //消息
	GroupSets     set.Interface   //好友 / 群
}

// 映射关系
var clientMap = make(map[int64]*Node, 0)

// 读写锁
var rwLocker sync.RWMutex

func Chat(w http.ResponseWriter, r *http.Request) {
	query := r.URL.Query()
	id := query.Get("userID")
	userID, err := strconv.ParseInt(id, 10, 64)
	if err != nil {
		fmt.Println(err)
	}
	isValida := true
	// 升级http为websocket
	conn, err := (&websocket.Upgrader{
		// 解决跨域
		CheckOrigin: func(r *http.Request) bool {
			return isValida
		},
	}).Upgrade(w, r, nil)
	if err != nil {
		panic(err)
	}
	//创建节点
	currentTime := uint64(time.Now().Unix())
	node := &Node{
		Conn:          conn,
		Addr:          conn.RemoteAddr().String(), //客户端地址
		HeartbeatTime: currentTime,                //心跳时间
		LoginTime:     currentTime,                //登录时间
		DataQueue:     make(chan []byte, 50),
		GroupSets:     set.New(set.ThreadSafe),
	}
	//3. 用户关系
	//4. userid 跟 node绑定 并加锁
	rwLocker.Lock()
	clientMap[userID] = node
	rwLocker.Unlock()
	//发送
	go sendProc(node)
	//接受
	go recvProc(node)
	////7.加入在线用户到缓存
	//SetUserOnlineInfo("online_"+Id, []byte(node.Addr), time.Duration(viper.GetInt("timeout.RedisOnlineTime"))*time.Hour)
	sendMsg(userID, []byte("欢迎进入聊天系统"))
}

func sendProc(node *Node) {
	for {
		select {
		case data := <-node.DataQueue:
			err := node.Conn.WriteMessage(websocket.TextMessage, data)
			if err != nil {
				panic(err)
			}
		}
	}
}

func recvProc(node *Node) {
	for {
		_, data, err := node.Conn.ReadMessage()
		if err != nil {
			fmt.Println(err)
			return
		}
		broadMsg(data)
		fmt.Println("[ws] <<<<<", data)
	}
}

var udpsendChan = make(chan []byte, 1024)

func broadMsg(data []byte) {
	udpsendChan <- data
}

func init() {
	go udpsendProc()
	go udpRecvProc()
}

func udpsendProc() {
	con, err := net.DialUDP("udp", nil, &net.UDPAddr{
		IP:   net.IPv4(192, 168, 1, 1),
		Port: 3000,
	})
	defer func() {
		err = con.Close()
		if err != nil {
			fmt.Println(err)
		}
	}()
	if err != nil {
		fmt.Println(err)
		return
	}
	for {
		select {
		case data := <-udpsendChan:
			_, err := con.Write(data)
			if err != nil {
				fmt.Println(err)
				return
			}
		}
	}
}

func udpRecvProc() {
	con, err := net.ListenUDP("udp", &net.UDPAddr{
		IP:   net.IPv4zero,
		Port: 3000,
	})
	if err != nil {
		fmt.Println(err)
		return
	}
	defer func() {
		err = con.Close()
		if err != nil {
			fmt.Println(err)
		}
	}()
	for {
		var buf [512]byte
		n, err := con.Read(buf[:])
		if err != nil {
			fmt.Println(err)
			return
		}
		dispatch(buf[:n])
	}
}

func dispatch(data []byte) {
	msg := UserMessage{}
	err := json.Unmarshal(data, &msg)
	if err != nil {
		fmt.Println(err)
		return
	}
	switch msg.SendType {
	case 1: //私信
		sendMsg(msg.TargetID, data)
		//case 2:sendGroupMsg()//群发
		//case 3:sendAllMsg()//广播
	}
}

func sendMsg(userID int64, msg []byte) {
	rwLocker.RLock()
	node, ok := clientMap[userID]
	rwLocker.RUnlock()
	if ok {
		node.DataQueue <- msg
	}
}
