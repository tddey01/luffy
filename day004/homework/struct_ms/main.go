package main

import (
	"fmt"
	"os"
)

// 学生信息管理
// 命令行菜单 fmt.
// 添加学生
//  修改学生
// 删除学生
// 展示学生
// 退出
func showMenu() {
	fmt.Println("学生信息管理系统！")
	fmt.Println("1 添加学生")
	fmt.Println("2 修改学生")
	fmt.Println("3 删除学生")
	fmt.Println("4 展示学生")
	fmt.Println("5 退出")
}

//  user and input
func userinput() *Student {
	var (
		name  string
		age   int
		id    int64
		class string
	)

	fmt.Println("请输入录入信息")
	fmt.Println("请输入命名:")
	fmt.Scanln(&name)
	fmt.Println("请输入年龄:")
	fmt.Scanln(&age)
	fmt.Println("请输入学号:")
	fmt.Scanln(&id)
	fmt.Println("请输入类型:")
	fmt.Scanln(&class)
	newStu := NewStudent(name, age, id, class)
	return newStu
}

func main() {
	stuMgr := NewStudentMgr()
	for {
		showMenu()
		//  后去用输入指令
		var i int
		fmt.Println("请输入指令")
		fmt.Scanln(&i)
		fmt.Println("输入指令是", i)

		switch i {
		case 1:
			fmt.Println("1 添加学生")
			newStu := userinput()
			stuMgr.AddStudent(newStu)
		case 2:
			fmt.Println("2 修改学生")
			newStu := userinput()
			stuMgr.EditStudent(newStu)
		case 3:
			fmt.Println("3 删除学生")
			newStu := userinput()
			stuMgr.DeleteStudent(newStu)
		case 4:
			fmt.Println("4 展示学生")
			stuMgr.ShowStudent()
		case 5:
			fmt.Println("5 退出")
			os.Exit(0)
		default:
			fmt.Println("输入无效")
		}

	}

}
