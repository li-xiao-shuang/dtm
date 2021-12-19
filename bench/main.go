package main

import (
	"fmt"
	"github.com/yedf/dtm/common"
	"github.com/yedf/dtm/dtmcli"
	"github.com/yedf/dtm/dtmcli/dtmimp"
	"github.com/yedf/dtm/dtmsvr"
	"github.com/yedf/dtm/examples"
	"os"
)

var hint = `To start the bench server, you need to specify the parameters:

Available commands:
    http              start bench server
`

func main() {
	if len(os.Args) <= 1 {
		fmt.Printf(hint)
		return
	}
	dtmimp.Logf("starting dtm....")
	if os.Args[1] == "http" {
		fmt.Println("start bench server")
		common.MustLoadConfig()
		// 设置当前数据库类型
		dtmcli.SetCurrentDBType(common.DtmConfig.DB["driver"])
		// 判断数据库是否启动
		common.WaitDBUp()
		// 执行建表sql
		dtmsvr.PopulateDB(true)
		// 执行案例sql
		examples.PopulateDB(true)
		dtmsvr.StartSvr()              // 启动dtmsvr的api服务
		go dtmsvr.CronExpiredTrans(-1) // 启动dtmsvr的定时过期查询
		StartSvr()
		select {}
	} else {
		fmt.Printf(hint)
	}
}
