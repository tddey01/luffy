package main

import "fmt"

// Student 学生信息管理
type Student struct {
	Name  string
	Age   int
	ID    int64
	Class string
}

// NewStudent 构造函数 创建一个学生对象的构造函数
func NewStudent(name string, age int, id int64, class string) *Student {
	return &Student{
		Name:  name,
		Age:   age,
		ID:    id,
		Class: class,
	}
}

// StudetMgr 定义一个学生信息管理结构体
type StudetMgr struct {
	AllStudents []*Student
}

// NewStudentMgr 初始化
func NewStudentMgr() *StudetMgr {
	return &StudetMgr{
		AllStudents: make([]*Student, 0, 100),
	}

}

// AddStudent 添加一个学生
func (s *StudetMgr) AddStudent(stu *Student) {
	for _, v := range s.AllStudents {
		if v.Name == stu.Name {
			fmt.Printf("姓名为%s的学生已经存在\n", stu.Name)
			return
		}
	}
	s.AllStudents = append(s.AllStudents, stu)
}

// EditStudent 修改学生的方法
func (s *StudetMgr) EditStudent(stu *Student) {
	for index, v := range s.AllStudents {
		if v.Name == stu.Name {
			s.AllStudents[index] = stu
			return
		}
	}
	fmt.Printf("姓名为%s的学生不存在", stu.Name)
}

// ShowStudent 展示学生方法
func (s *StudetMgr) ShowStudent() {
	for _, v := range s.AllStudents {
		fmt.Printf("姓名:%s  age:%d 序号:%d 类型:%s\n ", v.Name, v.Age, v.ID, v.Class)
	}
}

// DeleteStudent  删除学生方法
func (s *StudetMgr) DeleteStudent(stu *Student) {
	for index, v := range s.AllStudents {
		if v.Name == stu.Name {
			// 从切片中安装索引删除指定的元素
			s.AllStudents = append(s.AllStudents[:index], s.AllStudents[index+1:]...)
			return
		}
	}
	fmt.Printf("姓名为%s的学生不存在", stu.Name)
}
