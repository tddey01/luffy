//package rpc
//
//import (
//	"net"
//	"reflect"
//)
//
//// 声明客户端
//type Client struct {
//	conn net.Conn
//}
//
//// 创建客户端对象
//func NewClient(conn net.Conn) *Client {
//	return &Client{
//		conn: conn,
//	}
//}
//
//// 实现通用的RPC客户端
////绑定RPC访问的方法
//// 传入访问的函数名
//// 函数具体实现在Server端， Client只有函数原型
////使用MakeFunc（） 完成原型的函数的调用
////fPtr指向函数原型
//// xxx.callRPC (queryUser &query)
//func (c *Client) callRPC(rpcName string, fPtr interface{}) {
//	//通过反射， 获取fPtr为初始化的函数原型
//	fn := reflect.ValueOf(fPtr).Elem()
//	// 另一个函数，作用是对第一个函数参数操作
//	// 完成与Server交互
//	f := func(args []reflect.Value) []reflect.Value {
//		// 处理传入的参数
//		inArgs := make([]interface{}, 0, len(args))
//		for _, arg := range args {
//			inArgs = append(inArgs, arg.Interface())
//		}
//		//	 创建连接
//		cliSession := NewSession(c.conn)
//		// 编码数据
//		reqRPC := RPCData{Name: rpcName, Args: inArgs}
//		b, err := encode(reqRPC)
//		if err != nil {
//			panic(err)
//		}
//		// 写出数据
//		err = cliSession.Write(b)
//		if err != nil {
//			panic(err)
//		}
//		//	读取相应数据
//		respBytes, err := cliSession.Read()
//		if err != nil {
//			panic(err)
//		}
//		// 解码数据
//		respRPC, err := decode(respBytes)
//		if err != nil {
//			panic(err)
//		}
//		//	处理服务端返回的数据
//		outArgs := make([]reflect.Value, 0, len(respRPC.Args))
//		for i, arg := range respRPC.Args {
//			// 必须进行 nil转换
//			if arg != nil {
//				// 必须填充一个真正类型， 不能是nil
//				outArgs = append(outArgs, reflect.Zero(fn.Type().Out(i))) // 该类型的0位置
//				continue
//			}
//			outArgs = append(outArgs, reflect.ValueOf(arg))
//		}
//		return outArgs
//	}
//	// 参数： 一个未初始化函数的方法值， 类型是reflect.Type
//	// 参数: 另一个函数 作用对应一个函数参数操作
//	//	返回reflect.Value类型
//	//	 MakeFunc 使用传入的函数原型，创建一个绑定 参数2 的新函数
//	v := reflect.MakeFunc(fn.Type(), f)
//	//  为函数fPtr赋值
//	fn.Set(v)
//}
package rpc

import (
	"net"
	"reflect"
)

// 声明服务端
type Client struct {
	conn net.Conn
}

// 构造方法
func NewClient(conn net.Conn) *Client {
	return &Client{conn: conn}
}

// 实现通用的RPC客户端
// 传入访问的函数名
// fPtr指向的是函数原型
//var select fun xx(User)
//cli.callRPC("selectUser",&select)
func (c *Client) callRPC(rpcName string, fPtr interface{}) {
	// 通过反射，获取fPtr未初始化的函数原型
	fn := reflect.ValueOf(fPtr).Elem()
	// 需要另一个函数，作用是对第一个函数参数操作
	f := func(args []reflect.Value) []reflect.Value {
		// 处理参数
		inArgs := make([]interface{}, 0, len(args))
		for _, arg := range args {
			inArgs = append(inArgs, arg.Interface())
		}
		// 连接
		cliSession := NewSession(c.conn)
		// 编码数据
		reqRPC := RPCData{Name: rpcName, Args: inArgs}
		b, err := encode(reqRPC)
		if err != nil {
			panic(err)
		}
		// 写数据
		err = cliSession.Write(b)
		if err != nil {
			panic(err)
		}
		// 服务端发过来返回值，此时应该读取和解析
		respBytes, err := cliSession.Read()
		if err != nil {
			panic(err)
		}
		// 解码
		respRPC, err := decode(respBytes)
		if err != nil {
			panic(err)
		}
		// 处理服务端返回的数据
		outArgs := make([]reflect.Value, 0, len(respRPC.Args))
		for i, arg := range respRPC.Args {
			// 必须进行nil转换
			if arg == nil {
				// reflect.Zero()会返回类型的零值的value
				// .out()会返回函数输出的参数类型
				outArgs = append(outArgs, reflect.Zero(fn.Type().Out(i)))
				continue
			}
			outArgs = append(outArgs, reflect.ValueOf(arg))
		}
		return outArgs
	}
	// 完成原型到函数调用的内部转换
	// 参数1是reflect.Type
	// 参数2 f是函数类型，是对于参数1 fn函数的操作
	// fn是定义，f是具体操作
	v := reflect.MakeFunc(fn.Type(), f)
	// 为函数fPtr赋值，过程
	fn.Set(v)
}
