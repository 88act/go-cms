package utils

//读取服务器端硬件信息  github.com/super-l/machine-code/machine

import (
	"fmt"

	"github.com/super-l/machine-code/machine"
)

// @author: [linjd] 10212203@qq.com
// @function: GetCpuInfo
// @description: 获取硬件信息 cupid + serialNumber
// @return: c Cpu, err error
func GetDevInfo() (physicalID string, err error) {
	// serialNumber, err := machine.GetSerialNumber()
	// if err != nil {
	// 	fmt.Println(err.Error())
	// }
	// //fmt.Println("serialNumber = ", serialNumber)

	uuid, err := machine.GetPlatformUUID()
	if err != nil {
		fmt.Println(err.Error())
	}
	//fmt.Println("uuid = ", uuid)

	cpuid, err := machine.GetCpuId()
	if err != nil {
		fmt.Println(err.Error())
	}
	//fmt.Println("cpuid = ", cpuid)
	// macInfo, err := machine.GetMACAddress()
	// if err != nil {
	// 	fmt.Println(err.Error())
	// }
	// fmt.Println("mac = ", macInfo)
	physicalID = uuid + "_" + cpuid
	return physicalID, err

}
