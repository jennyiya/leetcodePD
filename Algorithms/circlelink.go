package circleList

import(
	"fmt"
)

//节点结构
type Object interface{}
type CNode struct{
	Data Object //定义数据域
	Prev *CNode //前驱节点
	Next *CNode //后继节点
}

//链表结构
type CList struct{
	length int
	head *CNode
	tail *CNode
}

//初始化
func (list *CList) Init() {
	//list = *(dList)
	list.length = 1
	// list.head = nil
	// list.tail = nil
	s := new(CNode)
	s.Next, s.Prev = s, s
    //return &CList{s, s, 0}
}

//判断链表是否为空
func (list *CList) IsEmpty() bool{
	if list.head == nil{ //头节点为空=>链表空
		return true
	}else{
		return false
	}
}

//头部插入元素
func (list *CList) Add(data Object){
	newNode := &CNode{Data:data}
	newNode.Next = list.head
	newNode.Prev = nil
	list.head.Prev = newNode
	//list.head = newNode
	list.length++
}

//尾部添加元素
func (list *CList) Append(data Object){
	newNode := &CNode{Data:data}
	//空链表则头节点是该元素
	if list.head == nil{
		list.head= newNode
		list.tail= newNode
		newNode.Prev = nil
		newNode.Next = newNode
	}else{
		newNode.Prev = list.tail
		newNode.Next = list.head
		list.tail.Next = newNode
		list.tail = newNode
	}
	list.length++
}

//中间插入：在下标是index后插入值data的节点
func (list *CList) Insert(index int, data Object) bool{
	cur := list.head
	if index < 0 {
		fmt.Println("无法插入！")
		return false
	}
	//插入开头
	if index == 0 {
		list.Add(data)
		return true
	}
	//插入在末尾
	if index == list.length-1 {
		list.Append(data)
		return true
	}
	i := 1
	//cur = cur.Next
	for cur.Next != list.head{
		//for i := 1; i < list.length-1; i++{
		if i == index {
			newNode := &CNode{Data:data}
			newNode.Prev = cur
			newNode.Next = cur.Next
			//改变cur指针
			cur.Next = newNode
			newNode.Next.Prev = newNode
			list.length++
			return true
		}
		cur = cur.Next
		i++		
	}
	return false
}

//删除下标index的节点
func (list *CList) DeleteIndex(index int) interface{} {
	if index < 0 || index >= list.length {
		fmt.Println("该节点不存在！")
		return nil
	}
	cur := list.head
	//只有一个节点删除
	if cur.Next == nil {
		list.head = nil
		cur.Next.Prev = nil
	}else{
		for i := 0; i < list.length; i++ {
			if i == index{
				//(cur.Next.Next).Prev= cur
				cur.Prev.Next = cur.Next
				cur.Next.Prev = cur.Prev
				break
			}
			cur=cur.Next	
		}
	}
	list.length--
	return cur.Data
}

//查找值data
//???是返回查找成功与否还是返回下标index
func (list *CList) Search(data Object) bool{
	if list.length == 0 {
		return false
	}
	cur := list.head.Next
	//count :=0
	//index :=[]int{}
	//isSearch := false
	// if cur.Data == data {
	// 	return count
	// }
	for cur != list.head{
		if cur.Data == data{
			//isSearch = true
			return true
			//index=append(index,count)
		}
		cur=cur.Next
		//count++	
	}
	// if isSearch == true {
	// 	fmt.Printf("双向链表值查找的index是：%v\n",index)
	// }
	return false
}

//打印链表元素
func (this *CList) ShowList(){
	fmt.Println("该双向循环链表是：")
	if this.IsEmpty(){
		fmt.Println("空链表！")
	}
	cur := this.head
	fmt.Printf("%v ", cur.Data)
	cur = cur.Next
	for cur!= this.head{
		fmt.Printf("%v ", cur.Data)
		cur = cur.Next
	}
	fmt.Println("")
}
