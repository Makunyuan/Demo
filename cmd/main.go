package main

import (
	"demo/internal/module"
	"fmt"
	"io/ioutil"
	"net"
	"os"
	"path/filepath"
	"text/template"
)

const udevRulePath = "/etc/udev/rules.d/70-persistent-net.rules"
const configuratioFilePath = "/etc/sysconfig/network-scripts"

func getDeviceList(interfaceNameList []string) []module.Device {
	//var interfaceNameList = [...]string{"enp10s0", "enp9s0", "enp8s0", "enp12s0", "enp7s0", "enp13s0", "enp6s0", "enp14s0", "enp15s0", "enp17s0",
	//	"enp16s0"}

	deviceSlice := make([]module.Device, 0)

	for i, v := range interfaceNameList {
		var device module.Device
		if i == 0 {
			device.Ipaddr = "192.168.0.1"
			device.Netmask = "255.255.255.0"
		}
		device.Name = fmt.Sprintf("eth%v", i)
		deviceHardwareAddr, err := getNetworkInfo(v)
		if err != nil {
			fmt.Printf("getNetworkInfo err : %v\n", err)
		}
		device.HardwareAddr = deviceHardwareAddr
		deviceSlice = append(deviceSlice, device)
	}
	return deviceSlice
}

func getNetworkInfo(interfaceName string) (string, error) {
	inf, err := net.InterfaceByName(interfaceName)
	if err != nil {
		return "", err
	}
	return inf.HardwareAddr.String(), nil

}

func generateUdevRule(rule string) {

	err := ioutil.WriteFile(udevRulePath, []byte(rule), 0666)
	if err != nil {
		fmt.Printf("generateUdevRule err : %v\n", err)

	}
}

func generateNetworkConfiguratioFile(device module.Device) {
	getTemplate, err := GetTemplate("ifcfg-template")
	var t = template.Must(template.New("ifcfg").Parse(getTemplate))
	fp, err := os.OpenFile(GetConfigPath(device.Name), os.O_RDWR|os.O_CREATE|os.O_TRUNC, 0o644)
	if err != nil {
		fmt.Printf("generate NetworkConfiguratio %v error :%v\n", device.Name, err)
	}
	defer func() {
		_ = fp.Close()
	}()

	err = t.Execute(fp, device)
}

func checkFileIsExist(filename string) bool {
	var exist = true
	if _, err := os.Stat(filename); os.IsNotExist(err) {
		exist = false
	}
	return exist
}

func GetTemplate(name string) (string, error) {
	filename := fmt.Sprintf("/usr/local/las/data/%s", name)
	out, err := ioutil.ReadFile(filename)
	if err != nil {
		return "", err
	}
	return string(out), nil
}
func GetConfigPath(dev string) string {
	return filepath.Join(configuratioFilePath, fmt.Sprintf("ifcfg-%s", dev))
}

func checkIsEth() bool {
	_, err := net.InterfaceByName("eth0")
	if err != nil {
		return false
	}
	return true
}

func main() {
	var interfaceNameList = []string{"eno16780032", "eno67111936", "eno50332672", "eno33561344"}[:]
	if !checkIsEth() {
		ruleStr := ""
		for i, v := range getDeviceList(interfaceNameList) {
			generateNetworkConfiguratioFile(v)
			oldFilePath := filepath.Join(configuratioFilePath, fmt.Sprintf("ifcfg-%s", interfaceNameList[i]))
			if checkFileIsExist(oldFilePath) {
				err := os.Remove(oldFilePath)
				if err != nil {
					fmt.Printf("Deleting file error: %v\n", err)
				}
			}
			fmt.Println(GetConfigPath(v.Name) + "  generation completed")
			ruleStr += fmt.Sprintf("ACTION==\"add\", SUBSYSTEM==\"net\", DRIVERS==\"?*\", ATTR{type}==\"1\", ATTR{address}==\"%v\", NAME=\"%v\"\n", v.HardwareAddr, v.Name)
		}
		generateUdevRule(ruleStr)
		fmt.Println(udevRulePath + " generation completed")
	} else {
		fmt.Println("eth0 already exists")
	}

}
