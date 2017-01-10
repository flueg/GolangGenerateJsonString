/****************************************************************************
* 版权所有  ： (C)2017 深圳响巢看看信息技术有限公司
* 设计部门  : 明星空间
* 系统名称  : 明星空间-直播机器人观众
* 文件名称  : template.go
* 内容摘要  :
* 当前版本  : 1.0
* 作    者 :
* 设计日期  : 2016年01月0x日
* 修改记录  :
* 修改记录  ：
  *1）、版本号
       *日  期：
       *修改人：FluegLau
       *摘  要：
****************************************************************************/

package main

import (
	"fmt"
)

func main2() {
	var ss string = "1005964138"
	var str interface{} = ss
	switch str.(type) {
	case string:
		fmt.Println("type of string")
	case []byte:
	case int:
		fmt.Println("type of int")
	case int64:
		fmt.Println("type of int64")
	case float64:
	case bool:
	case nil:
	default:
	}

}
