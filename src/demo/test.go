package main

import (
	"encoding/json"
	"fmt"
	"github.com/gostudys/huobi_contract_golang/src/ws/api"
	"github.com/gostudys/huobi_contract_golang/src/ws/response"
	"strings"
	"time"
)

// socket 服务 链接火币网，获取数据缓存到redis

func init() {
	var cstZone = time.FixedZone("CST", 8*3600) // 东八
	time.Local = cstZone
}

func main() {
	type topics struct {
		Topic string `json:"topic"` // K线图代码
	}
	var Topics []topics
	// 收集所有要订阅处理的 topic
	Topics = append(Topics,
		// BTC-USD
		// 1分钟k线图
		topics{
			Topic: "market.BTC-USD.index.1min",
		},
		// 5分钟k线图
		topics{
			Topic: "market.BTC-USD.index.5min",
		},
		// 15分钟k线图
		topics{
			Topic: "market.BTC-USD.index.15min",
		},
		// ETH-USD
		// 1分钟k线图
		topics{
			Topic: "market.ETH-USD.index.1min",
		},
		// 5分钟k线图
		topics{
			Topic: "market.ETH-USD.index.5min",
		},
		// 15分钟k线图
		topics{
			Topic: "market.ETH-USD.index.15min",
		})
	start := "ok"
	for {
		if start == "ok" {
			for _, obj := range Topics {
				go SocketHuoBi(obj.Topic)
			}
			start = "no"
		}
		time.Sleep(time.Second)
	}
}

// 链接火币网

func SocketHuoBi(parentTopic string) {
	Topic := strings.Split(parentTopic, ".")
	callback := make(chan struct{}, 1)
	// websocket api
	wsmkClient := new(api.WSIndexClient).Init("", callback)
	wsmkClient.SubIndexKLine(Topic[1], Topic[3], func(data *response.SubIndexKLineResponse) {
		// 火币网推送的回来的数据转换为字符串
		msgData, MarshalJSONErr := json.Marshal(data) //转换成JSON返回的是byte[]
		if MarshalJSONErr != nil {
			fmt.Println(MarshalJSONErr)
			return
		}
		if len(msgData) > 0 {
			fmt.Println(string(msgData))
		} else {
			return
		}

	}, "")
	select {
	case <-callback:
		fmt.Println(parentTopic + "：已关闭")
		fmt.Println(parentTopic + "：正在重新启动 Goroutine")
		go SocketHuoBi(parentTopic)
	}
}
