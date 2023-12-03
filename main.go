package main

import "fmt"

// Component 接口定义了组合中的基本操作
type Component interface {
	Show()
}

// MenuItem 表示菜单项
type MenuItem struct {
	Name string
}

// Show 实现了 Component 接口的 Show 方法
func (m *MenuItem) Show() {
	fmt.Printf("菜单项：%s\n", m.Name)
}

// Menu 表示菜单，可以包含菜单项和子菜单
type Menu struct {
	Name       string
	Components []Component
}

// Show 实现了 Component 接口的 Show 方法，遍历显示所有菜单项和子菜单
func (m *Menu) Show() {
	fmt.Printf("菜单：%s\n", m.Name)
	for _, component := range m.Components {
		component.Show()
	}
}

// Add 添加菜单项或子菜单
func (m *Menu) Add(c Component) {
	m.Components = append(m.Components, c)
}

func main() {
	// 创建菜单项
	item1 := &MenuItem{Name: "菜单项1"}
	item2 := &MenuItem{Name: "菜单项2"}

	// 创建子菜单
	submenu := &Menu{Name: "子菜单"}
	submenu.Add(&MenuItem{Name: "子菜单项1"})
	submenu.Add(&MenuItem{Name: "子菜单项2"})

	// 创建主菜单，包含菜单项和子菜单
	menu := &Menu{Name: "主菜单"}
	menu.Add(item1)
	menu.Add(item2)
	menu.Add(submenu)

	// 显示整个菜单系统
	menu.Show()
}
