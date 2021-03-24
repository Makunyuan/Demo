package main

import (
	"fmt"
	"github.com/urfave/cli/v2"
	"gopkg.in/yaml.v2"

	//yaml "gopkg.in/yaml.v2"
	"io/ioutil"
)

type DeviceName struct {
	OldDeviceName struct {
		List []string `yaml:"list"`
	}
	NewDeviceName struct {
		List []string `yaml:"list"`
	}
}

func main() {
	app := cli.NewApp()
	app.Name = "MNN"
	app.Usage = "Modify the network card name"
	//app.Action= func(context *cli.Context) error {
	//	fmt.Println("func print")
	//	return nil
	////}
	//err := app.Run(os.Args)
	//if err != nil {
	//	log.Fatal(err)
	//}
	//addrs, _ := net.Interfaces()
	//for _,v := range addrs{
	//	fmt.Println(v.Name)
	//}
	data, err := ioutil.ReadFile("E:/go_project/demo/config/templates/deviceName.yaml")
	fmt.Println(string(data))
	if err != nil {
		fmt.Println(err)
	}
	fmt.Println()
	var t DeviceName
	err = yaml.Unmarshal(data, &t)
	if err != nil {
		fmt.Println(err)
	}
	//fmt.Println("初始数据", t)
	//d, _ := yaml.Marshal(&t)
	//fmt.Println("看看 :", string(d))
	fmt.Println(t.NewDeviceName.List)

}
