package main

import (
	"os"
	"strconv"

	"github.com/atotto/clipboard"
	// "github.com/go-vgo/robotgo/clipboard"
	"github.com/therecipe/qt/core"
	"github.com/therecipe/qt/gui"
	"github.com/therecipe/qt/widgets"
)

func InitUI() *widgets.QMainWindow {
	var timeCN, timeEN, restaurant, mealCN, mealEN, hour, minute, apm string
	var googleSheetLink, drdLink string
	var resultCN, resultEN string
	var report string
	// create window
	app := widgets.NewQMainWindow(nil, 0)

	// set window title
	app.SetWindowTitle("酱酱の饭")

	// set window size
	app.SetGeometry2(20, 20, 1200, 960)

	// set app icon
	app.SetWindowIcon(gui.NewQIcon5("icon.ico"))

	// 布局窗口组件载体
	widget := widgets.NewQWidget(app, core.Qt__Widget)
	widgets.QApplication_SetFont(gui.NewQFont2("Microsoft JhengHei", 14, 1, false), "")

	// widget.SetGeometry(core.NewQRect4(300, 300, 300, 220))
	widget.SetGeometry2(0, 0, 1200, 960)
	app.SetCentralWidget(widget)

	// state
	app.StatusBar()

	// tag
	lbl := widgets.NewQLabel2("", widget, 0)
	lbl.Move2(20, 300)

	gridLayout := widgets.NewQGridLayout(widget)
	// 下拉列表
	// 选择中餐/晚餐
	meal := widgets.NewQComboBox(widget)
	meal.AddItem("Lunch", core.NewQVariant15("Lunch"))
	meal.AddItem("Dinner", core.NewQVariant15("Dinner"))
	// meal.Move2(50, 20)
	gridLayout.AddWidget2(widgets.NewQLabel2("中餐/晚餐: ", widget, 0), 0, 0, 0)
	gridLayout.AddWidget2(meal, 0, 1, 0)

	meal.ConnectActivated2(func(text string) {
		// lbl.SetText(text)
		if text == "Lunch" {
			timeCN = "中午"
			timeEN = "Morning"
			mealCN = "午餐"
			mealEN = "Lunch"
		} else {
			timeCN = "晚上"
			timeEN = "Evening"
			mealCN = "晚餐"
			mealEN = "Dinner"
		}
		lbl.AdjustSize()
	})

	// 选择饭店
	header := Converter()
	restaurants := widgets.NewQComboBox(widget)
	for _, v := range header {
		res := v.EN + " " + v.CN
		restaurants.AddItem(res, core.NewQVariant15(res))
	}

	// restaurants.Move2(50, 50)
	gridLayout.AddWidget2(widgets.NewQLabel2("饭店: ", widget, 0), 1, 0, 0)
	gridLayout.AddWidget2(restaurants, 1, 1, 0)

	// 当选中某个条目时会调用方法
	restaurants.ConnectActivated2(func(text string) {
		// lbl.SetText(text)
		restaurant = text
		lbl.AdjustSize()
	})

	// 截单时间
	// endtimeSelect := widgets.NewQLineEdit(widget)
	// gridLayout.AddWidget2(widgets.NewQLabel2("截单时间: ", widget, 0), 2, 0, 0)
	// gridLayout.AddWidget2(endtimeSelect, 2, 1, 0)
	// endtimeSelect.
	gridLayout.AddWidget2(widgets.NewQLabel2("截单时间: ", widget, 0), 2, 0, 0)
	hourSelect := widgets.NewQComboBox(widget)
	for i := 1; i < 13; i++ {
		hourSelect.AddItem(strconv.Itoa(i), core.NewQVariant15(strconv.Itoa(i)))
	}
	hourSelect.ConnectActivated2(func(text string) {
		hour = text
		lbl.AdjustSize()
	})
	
	gridLayout.AddWidget2(hourSelect, 2, 1, 1)

	minuteSelect := widgets.NewQComboBox(widget)
	for i := 1; i < 61; i++ {
		minuteSelect.AddItem(strconv.Itoa(i), core.NewQVariant15(strconv.Itoa(i)))
	}
	minuteSelect.ConnectActivated2(func(text string) {
		minute = text
		lbl.AdjustSize()
	})
	gridLayout.AddWidget2(minuteSelect, 3, 1, 1)

	apmSelect := widgets.NewQComboBox(widget)
	apmSelect.AddItem("am", core.NewQVariant15("am"))
	apmSelect.AddItem("pm", core.NewQVariant15("pm"))
	apmSelect.ConnectActivated2(func(text string) {
		apm = text
		lbl.AdjustSize()
	})
	gridLayout.AddWidget2(apmSelect, 4, 1, 1)

	// googleSheet链接
	googleSheet := widgets.NewQLineEdit(widget)
	gridLayout.AddWidget2(widgets.NewQLabel2("Google Sheet 链接: ", widget, 0), 5, 0, 0)
	gridLayout.AddWidget2(googleSheet, 5, 1, 0)
	googleSheet.ConnectTextChanged(func(text string) {
		googleSheetLink = text
		lbl.AdjustSize()
	})

	// drd链接
	drd := widgets.NewQLineEdit(widget)
	gridLayout.AddWidget2(widgets.NewQLabel2("doordash 链接: ", widget, 0), 6, 0, 0)
	gridLayout.AddWidget2(drd, 6, 1, 0)
	drd.ConnectTextChanged(func(text string) {
		drdLink = text
		lbl.AdjustSize()
	})

	// 提交按钮 - 显示最后结果
	btn := widgets.NewQPushButton2("提交☆(≧∀≦*)ﾉ ", app)
	btn.Resize(btn.SizeHint()) //设置按钮大小
	// btn.Move2(50, 100)         //设置按钮位置
	gridLayout.AddWidget2(widgets.NewQLabel2("", widget, 0), 8, 0, 0)
	gridLayout.AddWidget2(btn, 8, 1, 0)
	// 设置按钮触发，触发退出程序
	btn.ConnectClicked(func(checked bool) {
		resultCN = "@channel\n哈喽大家" + timeCN + "好 :blush: ，大家加班辛苦啦，我们" + mealCN + "OT Meal目的地是：" + restaurant +
			"，\n需要上车点餐的老板请在下面的链接里sign up your full name在G列：（" + hour + ":" + minute + " " + apm + " " + "截单）\n" +
			googleSheetLink + "\n这是菜单链接: \n" + drdLink + "\nThanks and enjoy it！~:yum:\n\n"
		resultEN = "@channel\nGood " + timeEN + " guys :blush: , we're going to order the OT meal for " + mealEN + " from: " + restaurant +
			", \nplease sign up your full name in column G below the OT Meal For " + mealEN + " if you need, here is the link for the name sheet:(" + hour + ":" + minute + " " + apm + " end)\n" +
			googleSheetLink + "\nand menu: \n" + drdLink + "\nThanks and enjoy it！~:yum:\n"
		// report := "已将上述内容复制到剪贴板"
		lbl.SetText(resultCN + resultEN)
		lbl.AdjustSize()
		// output := widgets.NewQLineEdit(widget)
		// output.SetText(resultCN + resultEN)
		// output.AdjustSize()
		// // gridLayout.AddWidget2(output, 6, 1, 0)
		// output.Move2(100, 300)
		// clipboard.WriteAll(resultCN + resultEN)
	})
	
	// gridLayout.AddWidget2(widgets.NewQLabel2("", widget, 0), 6, 0, 0)
	gridLayout.AddWidget2(btn, 9, 0, 0)

	vbox := widgets.NewQVBoxLayout2(widget)
	vbox.AddWidget(lbl, 0, 0)

	output := widgets.NewQLabel2("", widget, 0)

	copyBtnCN := widgets.NewQPushButton2("(つ•̀ω•́)つ复制中文", app)
	copyBtnCN.Resize(copyBtnCN.SizeHint())
	gridLayout.AddWidget2(widgets.NewQLabel2("", widget, 0), 10, 0, 0)
	gridLayout.AddWidget2(copyBtnCN, 10, 0, 0)
	copyBtnCN.ConnectClicked(func(checked bool) {
		clipboard.WriteAll(resultCN)
		report = "复制了d(￣▽￣*)b "
		output.SetText(report)
		output.AdjustSize()
		gridLayout.AddWidget2(output, 12, 1, 1)
	})
	

	copyBtnEN := widgets.NewQPushButton2("(σﾟ∀ﾟ)σ copy English", app)
	copyBtnEN.Resize(copyBtnEN.SizeHint())
	gridLayout.AddWidget2(widgets.NewQLabel2("", widget, 0), 11, 0, 0)
	gridLayout.AddWidget2(copyBtnEN, 11, 0, 0)
	copyBtnEN.ConnectClicked(func(checked bool) {
		clipboard.WriteAll(resultEN)
		report = "Over! _(:ι」∠)_"
		output.SetText(report)
		output.AdjustSize()
		gridLayout.AddWidget2(output, 12, 1, 1)
	})

	cancelBtn := widgets.NewQPushButton2("(≧ω≦)/清除已复制内容", app)
	cancelBtn.Resize(cancelBtn.SizeHint())
	gridLayout.AddWidget2(widgets.NewQLabel2("", widget, 0), 12, 0, 0)
	gridLayout.AddWidget2(cancelBtn, 12, 0, 0)
	cancelBtn.ConnectClicked(func(checked bool) {
		clipboard.WriteAll("")
		report = "∑(っ°Д°;)っ卧槽，不见了"
		output.SetText(report)
		output.AdjustSize()
		gridLayout.AddWidget2(output, 12, 1, 1)
	})

	return app
}

func main() {
	widgets.NewQApplication(len(os.Args), os.Args)
	app := InitUI()

	app.Show()

	widgets.QApplication_Exec()
}
