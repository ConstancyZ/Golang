package main

import (
	"bytes"
	"context"
	"encoding/csv"
	"encoding/json"
	"fmt"
	"go.etcd.io/etcd/clientv3"
	"os"
	"reflect"
	"strconv"
	"strings"
	"testing"
	"time"
)

// etcd client put/get demo
// use etcd/clientv3

func TestEtcd(t *testing.T) {
	// 引入etcd的一些操作
	cli, err := clientv3.New(clientv3.Config{
		Endpoints:   []string{"10.3.7.84:23791"},
		DialTimeout: 5 * time.Second,
	})
	if err != nil {
		// handle error!
		fmt.Printf("connect to etcd failed, err:%v\n", err)
		return
	}
	defer cli.Close()
	//put操作
	ctx, cancel := context.WithTimeout(context.Background(), time.Second)
	_, err = cli.Put(ctx, "one", "num1")
	cancel()
	if err != nil {
		fmt.Printf("put to etcd failed, err:%v\n", err)
		return
	}

	// get操作
	//ctx, cancel := context.WithTimeout(context.Background(), time.Second)
	//resp, err := cli.Get(ctx, "kod_path")
	//cancel()
	//if err != nil {
	//	fmt.Printf("get to etcd failed, err:%v\n", err)
	//	return
	//}
	//for _, ev := range resp.Kvs {
	//	fmt.Printf("%s:%s\n", ev.Key, ev.Value)
	//}

	// watch监听操作 key:zyh change
	//rch := cli.Watch(context.Background(), "zyh") // <- chan WatchResponse
	//for resp := range rch {
	//	for _, ev := range resp.Events {
	//		fmt.Printf("Type: %s Key:%s Value:%s\n", ev.Type, ev.Kv.Key, ev.Kv.Value)
	//	}
	//}

	// lease租续操作
	// 创建一个5s的租续(这个key能存活几秒)
	//resp, err := cli.Grant(context.TODO(), 5)
	//if err != nil {
	//	log.Fatalln(err)
	//}
	//// 5s后，"gn"这个key就会被移除
	//_, err = cli.Put(context.TODO(), "gn", "zyh", clientv3.WithLease(resp.ID))
	//if err != nil {
	//	log.Fatalln(err)
	//}
	//// 如果需要"gn"这个key一直续存的话，使用KeepAlive,这样就会在每次结束的时候续存5s
	//ch, keepErr := cli.KeepAlive(context.TODO(), resp.ID)
	//if keepErr != nil {
	//	log.Fatalln(keepErr)
	//}
	//for {
	//	ka := <-ch
	//	fmt.Println("ttl:", ka.TTL)
	//}

	// 基于etcd实现分布式锁

}

func TestDemo(t *testing.T) {
	str := "1,2,3,4"
	list := strings.Split(str, ",")
	fmt.Println(list)
	var ids bytes.Buffer
	for i, _ := range list {
		ids.WriteString(list[i] + ",")
	}
	hostnameIdStr, err := json.Marshal(list)
	if err != nil {

	}
	fmt.Println("ids:", string(hostnameIdStr))
	fmt.Println("ids:", ids.String())
	fmt.Println("ids:", len(ids.String()))
	fmt.Println("join:", strings.Join(list, ","))
	fmt.Println("ids:", string(hostnameIdStr))
	fmt.Println("ids:", ids.String())
	fmt.Println("ids:", len(ids.String()))
	fmt.Println("trim trim trim:", str[1:len(str)])
	fmt.Println("time:", time.UnixMilli(1658918231829).Format("2006-01-02 15:04:05"))
	dstf, err := os.Create("text.csv")
	defer dstf.Close()
	if err != nil {
		fmt.Println("err:", err)
	}

	dstf.WriteString("\xEF\xBB\xBF") // 写入UTF-8 BOM，防止中文乱码
	// 写数据到文件
	w := csv.NewWriter(dstf)
	startTime := time.Now()
	for i := 0; i < 50000; i++ {
		w.Write([]string{"交易对", "下单下限", "下单上限", "产品限额", "用户限额", "投资币种", "行权币种", "type", "剩余规模", "当前规模",
			"产品ID", "利率", "年化", "是否存活", "delta", "创建时间", "更新时间", "结算时间"})
	}
	endTime := time.Now()
	fmt.Println("time ........", endTime.Sub(startTime))
	w.Flush() // 此时才会将缓冲区数据写入
}

func TestDemo1(t *testing.T) {
	list := []animal{
		{
			name: "zz",
			age:  4,
		},
		{
			age:  5,
			name: "dd",
		},
	}
	var build strings.Builder
	for i, _ := range list {
		if i == len(list)-1 {
			build.WriteString(list[i].name)
			build.WriteString(",")
			build.WriteString(strconv.Itoa(list[i].age))
		} else {
			build.WriteString(list[i].name)
			build.WriteString(",")
			build.WriteString(strconv.Itoa(list[i].age))
			build.WriteString(";")
		}
	}

	fmt.Println(strings.TrimPrefix(build.String(), "z"))
	fmt.Println(strings.Split(build.String(), ";"))
	str := "module_management:121212133"
	fmt.Println(strings.TrimPrefix(str, "module_management:"))
}

type animal struct {
	name string
	age  int
}

func TestSql(t *testing.T) {
	ids := []int{
		1, 3,
	}
	modulesId := []int{
		1, 2, 3, 4, 6, 9, 10, 11, 12, 13, 14, 16, 17, 18,
	}
	strPrefix := "insert into data_filter(data_name, data_id, user_id, operater_id,op_permission) values"
	str1 := "('module_management'," // 模块
	str2 := "1),"                   // 权限
	// "'1','20','20',1)"
	for _, v := range modulesId {
		for _, vv := range ids {
			modulesIdStr := "'" + strconv.Itoa(v) + "',"
			idStr := "'" + strconv.Itoa(vv) + "',"
			strPrefix += str1
			strPrefix += modulesIdStr
			strPrefix += idStr
			strPrefix += idStr
			strPrefix += str2
		}
	}
	fmt.Println(strings.TrimSuffix(strPrefix, ",") + ";")
}

func TestDemo2(t *testing.T) {
	var build strings.Builder
	build.WriteString("s")
	build.WriteString(":")
	list := []int64{
		1,
	}
	for i, _ := range list {
		if i == len(list)-1 {
			build.WriteString(strconv.FormatInt(list[i], 10))
		} else {
			build.WriteString(strconv.FormatInt(list[i], 10))
			build.WriteString(",")
		}
	}
	fmt.Println(build.String())
	str := build.String()
	dataStr := strings.TrimPrefix(str, "s:")
	// 这里error不需要返回，quark服务自己会进行slice非空判断
	if len(dataStr) == 0 {
	}
	list1 := []int64{}
	strList := strings.Split(dataStr, ",")
	for i, _ := range strList {
		dataID, _ := strconv.ParseInt(strList[i], 10, 64)
		list1 = append(list1, dataID)
	}
	fmt.Println(list1)
	fmt.Println(strconv.FormatBool(false) == "false")
	fmt.Println(strconv.FormatBool(true) == "false")
}

func TestDemo3(t *testing.T) {
	leyangjun1 := []int{10, 9, 8}
	leyangjun2 := []int{4, 5, 6, 7, 8}
	var c []int
	temp := map[int]struct{}{}

	for _, val := range leyangjun1 {
		if _, ok := temp[val]; !ok {
			temp[val] = struct{}{}
		}
	}

	for _, val := range leyangjun2 {
		if _, ok := temp[val]; !ok {
			c = append(c, val)
		}
	}

	fmt.Println(c)
	time1 := time.Now()
	fmt.Println(time1.UnixMilli())
	fmt.Println(strconv.ParseInt(time1.Format("20060102"), 10, 64))
}

func TestDemo4(t *testing.T) {
	var a float64
	a = 0.22
	fmt.Println(reflect.TypeOf(a).Name())
	var longTime int64
	longTime = 1663575890045
	fmt.Println(time.UnixMicro(longTime).Format("2020-22-22"))
}

func TestDemo5(t *testing.T) {
	a := time.Now().UnixNano() / 1e6
	aTime := time.Unix(a/1000, (a%1000)*(1000*1000))
	fmt.Println(aTime.Format("2006-01-02 15:04:05"))
}

func TestDemo6(t *testing.T) {
	str := "{\"p_symbol\":\"BTC-USDT-SPT\",\"exchange\":\"kucoin_spot\",\"symbol\":\"BTC-USDT\",\"mark_price\":20137.6,\"mark_price_origin\":20137.590357819652,\"update_at\":1664269982977,\"event_at\":1664269740000,\"base_window_minutes\":10,\"adapt_window_minutes\":10,\"base_exchange_count\":5,\"adapt_exchange_count\":5,\"mark_price_type\":\"Strong\",\"day\":20220927,\"base_exchange_vwap\":{\"gate_spot\":20148.65867038689,\"binance_spot\":20144.00784844717,\"okex_spot_v5\":20144.542463408598,\"ftx_spot\":20137.478537183288,\"huobi_spot\":20131.014760383485},\"adapt_exchange_vwap\":{\"gate_spot\":20148.65867038689,\"okex_spot_v5\":20144.542463408598,\"binance_spot\":20144.00784844717,\"huobi_spot\":20131.014760383485,\"ftx_spot\":20137.478537183288}}"
	p := &Price{}
	if err := json.Unmarshal([]byte(str), p); err != nil {
	}
	fmt.Println(p)

}

type Price struct {
	MarkPrice float64 `json:"mark_price,omitempty"`
}
type GenerateStruct struct {
	PSymbol            string            `json:"p_symbol,omitempty"`
	Exchange           string            `json:"exchange,omitempty"`
	Symbol             string            `json:"symbol,omitempty"`
	MarkPrice          float64           `json:"mark_price,omitempty"`
	MarkPriceOrigin    float64           `json:"mark_price_origin,omitempty"`
	UpdateAt           int               `json:"update_at,omitempty"`
	EventAt            int               `json:"event_at,omitempty"`
	BaseWindowMinutes  int               `json:"base_window_minutes,omitempty"`
	AdaptWindowMinutes int               `json:"adapt_window_minutes,omitempty"`
	BaseExchangeCount  int               `json:"base_exchange_count,omitempty"`
	AdaptExchangeCount int               `json:"adapt_exchange_count,omitempty"`
	MarkPriceType      string            `json:"mark_price_type,omitempty"`
	Day                int               `json:"day,omitempty"`
	BaseExchangeVwap   BaseExchangeVwap  `json:"base_exchange_vwap,omitempty"`
	AdaptExchangeVwap  AdaptExchangeVwap `json:"adapt_exchange_vwap,omitempty"`
}

type BaseExchangeVwap struct {
	GateSpot    float64 `json:"gate_spot,omitempty"`
	BinanceSpot float64 `json:"binance_spot,omitempty"`
	OkexSpotV5  float64 `json:"okex_spot_v5,omitempty"`
	FtxSpot     float64 `json:"ftx_spot,omitempty"`
	HuobiSpot   float64 `json:"huobi_spot,omitempty"`
}

type AdaptExchangeVwap struct {
	GateSpot    float64 `json:"gate_spot,omitempty"`
	OkexSpotV5  float64 `json:"okex_spot_v5,omitempty"`
	BinanceSpot float64 `json:"binance_spot,omitempty"`
	HuobiSpot   float64 `json:"huobi_spot,omitempty"`
	FtxSpot     float64 `json:"ftx_spot,omitempty"`
}
