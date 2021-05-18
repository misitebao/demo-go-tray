package main

import (
	"fmt"
	"io/ioutil"
	"time"

	"github.com/riftbit/go-systray"
)

var (
	timezone string
)

func main() {
	//systray.SetCustomLeftClickAction()
	//systray.SetCustomRightClickAction()
	systray.Run(onReady, onExit)
}

func onReady() {
	timezone = "默认文字"
	systray.SetIcon(getIcon("./favicon_misitebao.ico"))

	submenu := systray.AddSubMenu("子菜单")
	_ = submenu.AddSubMenuItem("开始", "", 0)
	_ = submenu.AddSubMenuItem("结束", "", 0)

	localTime := systray.AddMenuItem("炎龙", "炎龙", 0)
	hcmcTime := systray.AddMenuItem("风鹰", "风鹰", 0)
	sydTime := systray.AddMenuItem("黑犀", "黑犀", 0)
	gdlTime := systray.AddMenuItem("地虎", "地虎", 0)
	sfTime := systray.AddMenuItem("雪獒", "雪獒", 0)

	fmt.Printf("%#v", localTime)

	systray.AddSeparator()
	mQuit := systray.AddMenuItem("退出", "退出", 0)

	go func() {
		for {
			systray.SetTitle(timezone)
			systray.SetTooltip(timezone)
			time.Sleep(1 * time.Second)
		}
	}()

	go func() {
		for {
			select {
			case <-localTime.OnClickCh():
				timezone = "炎龙"
			case <-hcmcTime.OnClickCh():
				timezone = "风鹰"
			case <-sydTime.OnClickCh():
				timezone = "黑犀"
			case <-gdlTime.OnClickCh():
				timezone = "地虎"
			case <-sfTime.OnClickCh():
				timezone = "雪獒"
			case <-mQuit.OnClickCh():
				systray.Quit()
				return
			}
		}
	}()
}

func onExit() {
	// 清除销毁
}

func getIcon(s string) []byte {
	b, err := ioutil.ReadFile(s)
	if err != nil {
		fmt.Print(err)
	}
	return b
}
