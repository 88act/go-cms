package utils

//读取服务器端硬件信息 github.com/shirou/gopsutil/cpu
//

import (
	"fmt"
	"go-cms/global"
	"strings"

	"crypto/md5"
	"crypto/rand"
	"crypto/x509"
	"encoding/base64"
	"encoding/hex"
	"encoding/pem"

	"net"
	"sort"

	"github.com/shirou/gopsutil/cpu"
	"github.com/shirou/gopsutil/host"
)

// @author: [linjd] 10212203@qq.com
// @function: GetCpuInfo
// @description: CPU信息 physicalID
// @return: c Cpu, err error
func GetCpuInfo() (physicalID string, err error) {
	fmt.Println("GetCpuInfo CPU, cpu.Info=")
	cpuinfo, err := cpu.Info()
	fmt.Println(cpuinfo)
	if err != nil {
		return "", err
	}
	physicalID = cpuinfo[0].PhysicalID
	fmt.Println("physicalID = ", physicalID)
	return physicalID, nil
}

func CheckLicense(devCode string) (isOk bool, err error) {

	filePrv := global.CONFIG.License.RsaPath + "prvKey.txt"
	//filePub := "./pubKey.txt"
	prv2, err := LoadPrivateKeyFile(filePrv)
	if err != nil {
		global.LOG.Error(err.Error())
		fmt.Println("失败1 err=", err.Error())
		return false, err
	}
	derStream := x509.MarshalPKCS1PrivateKey(prv2)
	block := &pem.Block{
		Type:  "RSA PRIVATE KEY",
		Bytes: derStream,
	}
	prvKey := pem.EncodeToMemory(block)
	rsaInfo := global.CONFIG.License.AppId + "|" + devCode + "|" + global.CONFIG.License.EndTime + "|" + global.CONFIG.License.MaxClient + "|" + global.CONFIG.License.RtcType
	//fmt.Println(" 签名 rsaInfo  =", rsaInfo)
	rsaInfoMd5 := Md5ByString(rsaInfo)
	//fmt.Println(" 签名 rsaInfoMd5 =", rsaInfoMd5)
	signData := RsaSignWithSha256([]byte(rsaInfoMd5), prvKey)
	licence := base64.StdEncoding.EncodeToString(signData) //hex.EncodeToString(signData) //
	//fmt.Println("机器签名 licence = ", licence)
	//fmt.Println("配置签名 licence = ", global.CONFIG.License.License)
	if licence == global.CONFIG.License.License {
		return true, nil
	} else {
		return false, nil
	}
}

func CheckLicense2() (isOk bool, err error) {
	key := "jiquxr88.com12345678abcd"
	//str := data.DevCode + "|" + data.AppName + "|" + data.AppEndtime.Format("2006-01-02 15:04:05")
	strEncode := Des3Decode(global.CONFIG.License.License, key)
	fmt.Println(" CheckLicense strEncode   = ", strEncode)
	strList := strings.Split(strEncode, "|")
	//fmt.Println(" CheckLicense strList  = ", strList)
	//fmt.Println("GetPhysicalID -------2----")
	physicalID := GetPhysicalID()
	//fmt.Println(physicalID)
	//fmt.Println("GetPhysicalID -------2----")
	if physicalID == strList[0] {
		fmt.Println(" 授权正常 ")
		return true, nil
	}
	fmt.Println(" 授权失败 ")
	return false, nil
}

// GetPhysicalID
func GetPhysicalID() string {
	var ids []string

	cpuinfo, err := cpu.Info()
	if err == nil {
		fmt.Println("PhysicalID=", cpuinfo[0].PhysicalID)
		fmt.Println("VendorID=", cpuinfo[0].VendorID)
		ids = append(ids, cpuinfo[0].VendorID)
		ids = append(ids, cpuinfo[0].PhysicalID)
	}
	if mac, err := getMACAddress(); err != nil {
		ids = append(ids, "mac")
	} else {
		fmt.Println("mac=", mac)
		ids = append(ids, mac)
	}

	platform, family, version, _ := host.PlatformInformation()
	fmt.Println("platform:", platform)
	fmt.Println("family:", family)
	fmt.Println("version:", version)
	ids = append(ids, platform)
	ids = append(ids, family)
	ids = append(ids, version)

	sort.Strings(ids)
	idsstr := strings.Join(ids, "-#-")
	fmt.Println("idsstr=", idsstr)
	return GetMd5String(idsstr, true, true)

	// if guid, err := getMachineGuid(); err != nil {
	// 	fmt.Println("guid=", guid)
	// 	panic(err.Error())
	// } else {
	// 	ids = append(ids, guid)
	// }
}

func getMACAddress() (string, error) {
	netInterfaces, err := net.Interfaces()
	if err != nil {
		return "", err
	}
	mac := ""
	for i := 0; i < len(netInterfaces); i++ {
		//fmt.Println(netInterfaces[i])
		if (netInterfaces[i].Flags&net.FlagUp) != 0 && (netInterfaces[i].Flags&net.FlagLoopback) == 0 {
			addrs, _ := netInterfaces[i].Addrs()
			for _, address := range addrs {
				ipnet, ok := address.(*net.IPNet)
				//fmt.Println(ipnet.IP)
				if ok && ipnet.IP.IsGlobalUnicast() {
					// 如果IP是全局单拨地址，则返回MAC地址
					mac = netInterfaces[i].HardwareAddr.String()
					return mac, nil
				}
			}
		}
	}

	return mac, nil
}

/*
func getMachineGuid() (string, error) {
	// there has been reports of issues on 32bit using golang.org/x/sys/windows/registry, see https://github.com/shirou/gopsutil/pull/312#issuecomment-277422612
	// for rationale of using windows.RegOpenKeyEx/RegQueryValueEx instead of registry.OpenKey/GetStringValue
	var h windows.Handle
	err := windows.RegOpenKeyEx(windows.HKEY_LOCAL_MACHINE, windows.StringToUTF16Ptr(`SOFTWARE\Microsoft\Cryptography`), 0, windows.KEY_READ|windows.KEY_WOW64_64KEY, &h)
	if err != nil {
		return "", err
	}
	defer windows.RegCloseKey(h)

	const windowsRegBufLen = 74 // len(`{`) + len(`abcdefgh-1234-456789012-123345456671` * 2) + len(`}`) // 2 == bytes/UTF16
	const uuidLen = 36

	var regBuf [windowsRegBufLen]uint16
	bufLen := uint32(windowsRegBufLen)
	var valType uint32
	err = windows.RegQueryValueEx(h, windows.StringToUTF16Ptr(`MachineGuid`), nil, &valType, (*byte)(unsafe.Pointer(&regBuf[0])), &bufLen)
	if err != nil {
		return "", err
	}

	hostID := windows.UTF16ToString(regBuf[:])
	hostIDLen := len(hostID)
	if hostIDLen != uuidLen {
		return "", fmt.Errorf("HostID incorrect: %q\n", hostID)
	}

	return hostID, nil
}
*/
//生成32位md5字串
func GetMd5String(s string, upper bool, half bool) string {
	h := md5.New()
	h.Write([]byte(s))
	result := hex.EncodeToString(h.Sum(nil))
	if upper == true {
		result = strings.ToUpper(result)
	}
	if half == true {
		result = result[8:24]
	}
	return result
}

// 利用随机数生成Guid字串
func UniqueId() string {
	b := make([]byte, 48)
	if _, err := rand.Read(b); err != nil {
		return ""
	}
	return GetMd5String(base64.URLEncoding.EncodeToString(b), true, false)
}
