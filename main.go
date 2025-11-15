package main

import (
	"fmt"
	"image/color"

	"fyne.io/fyne/v2"
	"fyne.io/fyne/v2/app"
	"fyne.io/fyne/v2/canvas"
	"fyne.io/fyne/v2/container"
	"fyne.io/fyne/v2/theme"
	"fyne.io/fyne/v2/widget"
)

// 自定义主题
type myTheme struct {
	fyne.Theme
}

func (m myTheme) Color(name fyne.ThemeColorName, variant fyne.ThemeVariant) color.Color {
	if name == theme.ColorNameButton {
		return color.RGBA{R: 70, G: 130, B: 180, A: 255} // 蓝色按钮
	}
	if name == theme.ColorNamePrimary {
		return color.RGBA{R: 60, G: 120, B: 216, A: 255} // 亮蓝色主色调
	}
	if name == theme.ColorNameHover {
		return color.RGBA{R: 80, G: 140, B: 200, A: 255} // 悬停颜色
	}
	return theme.DarkTheme().Color(name, variant)
}

func (m myTheme) Font(style fyne.TextStyle) fyne.Resource {
	return theme.DefaultTheme().Font(style)
}

func (m myTheme) Icon(name fyne.ThemeIconName) fyne.Resource {
	return theme.DefaultTheme().Icon(name)
}

func (m myTheme) Size(name fyne.ThemeSizeName) float32 {
	if name == theme.SizeNameText {
		return 14 // 稍大的字体
	}
	return theme.DefaultTheme().Size(name)
}

func main() {
	// 创建应用
	a := app.New()
	a.Settings().SetTheme(&myTheme{}) // 使用自定义主题

	// 创建窗口
	w := a.NewWindow("Fyne组件演示")
	w.Resize(fyne.NewSize(800, 600))
	w.CenterOnScreen() // 窗口居中显示

	// 创建标题标签
	titleLabel := widget.NewLabelWithStyle("Fyne组件演示", fyne.TextAlignCenter, fyne.TextStyle{Bold: true})

	// 创建标签组件
	label := widget.NewLabelWithStyle("标签", fyne.TextAlignCenter, fyne.TextStyle{Italic: true})

	// 创建编辑框组件
	entry := widget.NewEntry()
	entry.SetPlaceHolder("输入文本")

	// 创建按钮组件
	button := widget.NewButtonWithIcon("点击", theme.ConfirmIcon(), func() {
		// 按钮点击事件：将编辑框的内容设置到标签上
		label.SetText(entry.Text)
	})

	// 创建表格组件
	// 准备表格数据
	headers := []string{"ID", "姓名", "年龄"}
	data := [][]string{
		{"1", "张三", "28"},
		{"2", "李四", "32"},
		{"3", "王五", "25"},
	}

	// 创建表格表头
	headerContainer := container.NewHBox()
	for _, header := range headers {
		headerLabel := widget.NewLabelWithStyle(header, fyne.TextAlignLeading, fyne.TextStyle{Bold: true})
		headerContainer.Add(headerLabel)
	}

	// 创建表格内容
	tableContainer := container.NewVBox()
	for i, row := range data {
		rowContainer := container.NewHBox()
		for _, cell := range row {
			rowContainer.Add(widget.NewLabel(cell))
		}
		// 为奇数行添加背景色
		if i%2 == 1 {
			bg := canvas.NewRectangle(color.RGBA{R: 40, G: 40, B: 40, A: 255})
			bgContainer := container.NewStack(bg, rowContainer)
			tableContainer.Add(bgContainer)
		} else {
			tableContainer.Add(rowContainer)
		}
	}

	// 创建渐变图片组件
	gradient := canvas.NewLinearGradient(color.RGBA{R: 100, G: 150, B: 200, A: 255}, color.RGBA{R: 50, G: 100, B: 150, A: 255}, 0)
	gradient.SetMinSize(fyne.NewSize(50, 50))

	// 创建下拉框组件
	selectOptions := []string{"选项1", "选项2", "选项3"}
	selectEntry := widget.NewSelect(selectOptions, func(selected string) {
		label.SetText("选择: " + selected)
	})
	selectEntry.SetSelected("选项1")

	// 创建单选框组件
	radioOptions := []string{"男性", "女性"}
	radio := widget.NewRadioGroup(radioOptions, func(selected string) {
		label.SetText("性别: " + selected)
	})
	radio.SetSelected("男性")

	// 创建复选框组件
	check1 := widget.NewCheck("通知", func(checked bool) {
		status := "关"
		if checked {
			status = "开"
		}
		label.SetText("通知" + status)
	})
	check1.SetChecked(true)

	check2 := widget.NewCheck("自动保存", func(checked bool) {
		status := "关"
		if checked {
			status = "开"
		}
		label.SetText("自动保存" + status)
	})

	// 创建进度条组件
	progress := widget.NewProgressBar()
	progress.SetValue(0.7)

	// 创建滑块组件
	slider := widget.NewSlider(0, 100)
	slider.SetValue(50)
	sliderLabel := widget.NewLabel("50")
	slider.OnChanged = func(value float64) {
		sliderLabel.SetText(fmt.Sprintf("%.0f", value))
	}

	// 创建列表组件
	listData := []string{"项目1", "项目2", "项目3", "项目4"}
	list := widget.NewList(
		func() int {
			return len(listData)
		},
		func() fyne.CanvasObject {
			return widget.NewLabel("项目")
		},
		func(id widget.ListItemID, obj fyne.CanvasObject) {
			obj.(*widget.Label).SetText(listData[id])
		},
	)
	list.Resize(fyne.NewSize(150, 80))

	// 创建多行文本框组件
	multiEntry := widget.NewMultiLineEntry()
	multiEntry.SetPlaceHolder("多行文本...")
	multiEntry.Resize(fyne.NewSize(200, 60))

	// 创建选项卡组件
	tabContent := container.NewVBox(
		widget.NewLabelWithStyle("选项卡内容1", fyne.TextAlignCenter, fyne.TextStyle{Bold: true}),
		widget.NewButtonWithIcon("按钮", theme.ConfirmIcon(), func() {
			label.SetText("点击选项卡按钮")
		}),
	)

	tabContent2 := container.NewVBox(
		widget.NewLabelWithStyle("选项卡内容2", fyne.TextAlignCenter, fyne.TextStyle{Bold: true}),
		widget.NewEntry(),
	)

	// 创建选项卡容器
	tabs := container.NewAppTabs(
		container.NewTabItemWithIcon("基本", theme.HomeIcon(), tabContent),
		container.NewTabItemWithIcon("高级", theme.SettingsIcon(), tabContent2),
	)

	// 创建工具栏
	toolbar := container.NewHBox(
		widget.NewButtonWithIcon("", theme.FileIcon(), func() {
			label.SetText("文件")
		}),
		widget.NewButtonWithIcon("", theme.FolderOpenIcon(), func() {
			label.SetText("文件夹")
		}),
		widget.NewSeparator(),
		widget.NewButtonWithIcon("", theme.ContentCutIcon(), func() {
			label.SetText("剪切")
		}),
		widget.NewButtonWithIcon("", theme.ContentCopyIcon(), func() {
			label.SetText("复制")
		}),
		widget.NewButtonWithIcon("", theme.ContentPasteIcon(), func() {
			label.SetText("粘贴")
		}),
	)

	// 创建表单
	form := widget.NewForm(
		widget.NewFormItem("用户名", widget.NewEntry()),
		widget.NewFormItem("密码", widget.NewPasswordEntry()),
		widget.NewFormItem("邮箱", widget.NewEntry()),
	)
	form.OnSubmit = func() {
		label.SetText("表单已提交")
	}
	form.OnCancel = func() {
		label.SetText("表单已取消")
	}

	// 创建菜单栏
	menu := fyne.NewMainMenu(
		fyne.NewMenu("文件",
			fyne.NewMenuItem("新建", func() { label.SetText("新建文件") }),
			fyne.NewMenuItem("打开", func() { label.SetText("打开文件") }),
			fyne.NewMenuItem("保存", func() { label.SetText("保存文件") }),
			fyne.NewMenuItemSeparator(),
			fyne.NewMenuItem("退出", func() { a.Quit() }),
		),
		fyne.NewMenu("编辑",
			fyne.NewMenuItem("剪切", func() { label.SetText("剪切") }),
			fyne.NewMenuItem("复制", func() { label.SetText("复制") }),
			fyne.NewMenuItem("粘贴", func() { label.SetText("粘贴") }),
		),
		fyne.NewMenu("帮助",
			fyne.NewMenuItem("关于", func() { label.SetText("关于应用") }),
		),
	)

	// 设置窗口菜单
	w.SetMainMenu(menu)

	// 创建带边框的卡片容器
	createCard := func(title string, content fyne.CanvasObject) fyne.CanvasObject {
		card := widget.NewCard(title, "", content)
		return card
	}

	// 创建美化型布局 - 使用网格布局
	// 第一行：基本信息和选择器
	row1 := container.NewGridWithColumns(2,
		// 左侧：基本信息
		createCard("基本信息", container.NewVBox(
			label,
			entry,
			button,
		)),
		// 右侧：选择器
		createCard("选择器", container.NewVBox(
			container.NewGridWithColumns(2,
				widget.NewLabel("下拉框:"),
				selectEntry,
			),
			container.NewGridWithColumns(2,
				widget.NewLabel("单选框:"),
				radio,
			),
		)),
	)

	// 第二行：复选框和进度指示器
	row2 := container.NewGridWithColumns(2,
		// 左侧：复选框
		createCard("复选框", container.NewVBox(check1, check2)),
		// 右侧：进度指示器
		createCard("进度指示器", container.NewVBox(
			progress,
			container.NewHBox(slider, sliderLabel),
		)),
	)

	// 第三行：列表和图片
	row3 := container.NewGridWithColumns(2,
		// 左侧：列表
		createCard("列表", list),
		// 右侧：图片
		createCard("图片", gradient),
	)

	// 第四行：多行文本框和表格
	row4 := container.NewGridWithColumns(2,
		// 左侧：多行文本框
		createCard("多行文本框", multiEntry),
		// 右侧：表格
		createCard("表格", container.NewVBox(headerContainer, tableContainer)),
	)

	// 第五行：表单和选项卡
	row5 := container.NewGridWithColumns(2,
		// 左侧：表单
		createCard("表单", form),
		// 右侧：选项卡
		createCard("选项卡", tabs),
	)

	// 创建整体布局
	content := container.NewVBox(
		titleLabel,
		widget.NewSeparator(),
		toolbar,
		widget.NewSeparator(),
		row1,
		widget.NewSeparator(),
		row2,
		widget.NewSeparator(),
		row3,
		widget.NewSeparator(),
		row4,
		widget.NewSeparator(),
		row5,
	)

	// 使用滚动容器包装内容
	scrollContainer := container.NewScroll(content)

	// 设置窗口内容
	w.SetContent(scrollContainer)

	// 显示窗口
	w.ShowAndRun()
}
