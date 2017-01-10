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
	"strconv"
	"strings"
)

func GetChance(s string) []int {
	fmt.Println("GetChance: %s", s)
	var ia []int
	values := strings.Split(s, ",")
	fmt.Printf("%v\n", values)
	for _, v := range values {
		i, err := strconv.ParseInt(v, 0, 32)
		if err == nil {
			ia = append(ia, int(i))
		}
	}
	fmt.Printf("%v\n", ia)
	return ia
}

func main() {
	var ss string = " 1,2,3,3,3,4"
	fmt.Println(GetChance(ss))
}
