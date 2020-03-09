package doubleList

import(
	"fmt"
)

//节点结构
type Object interface{}
type DNode struct{
	Data Object //定义数据域
	Prev *DNode //前驱节点
	Next *DNode //后继节点
}

//链表结构
type DList struct{
	length int
	head *DNode
	tail *DNode
}

//初始化
func (list *DList) Init() {
	//list = *(dList)
	list.length = 0
	list.head = nil
	list.tail = nil
}

//判断链表是否为空
func (list *DList) IsEmpty() bool{
	if list.head == nil{ //头节点为空=>链表空
		return true
	}else{
		return false
	}
}

//头部插入元素
func (list *DList) Add(data Object){
	newNode := &DNode{Data:data}
	newNode.Next = list.head
	newNode.Prev = nil
	list.head = newNode
	list.length++
}

//尾部添加元素
func (list *DList) Append(data Object){
	newNode := &DNode{Data:data}
	//空链表则头节点是该元素
	if list.head == nil{
		list.head= newNode
		list.tail= newNode
		newNode.Prev = nil
		newNode.Next = nil
	}else{
		newNode.Prev = list.tail
		newNode.Next = nil
		list.tail.Next = newNode
		list.tail = newNode
	}
	list.length++
}

//中间插入：在下标是index后插入值data的节点
func (list *DList) Insert(index int, data Object) bool{
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
	for i := 1; i < list.length-1; i++{
		if i == index {
			newNode := &DNode{Data:data}
			newNode.Prev = cur
			newNode.Next = cur.Next
			//改变cur指针
			cur.Next = newNode
			newNode.Next.Prev = newNode
			list.length++
			return true
		}
		cur = cur.Next		
	}
	return false
}

//删除下标index的节点
func (list *DList) DeleteIndex(index int) interface{} {
	if index < 0 || index >= list.length {
		fmt.Println("该节点不存在！")
		return nil
	}
	cur := list.head
	if cur.Next == nil {
		list.head = cur.Next
		cur.Next.Prev = nil
	}
	for i := 0; i < list.length; i++ {
		if i == index{
			//(cur.Next.Next).Prev= cur
			cur.Prev.Next = cur.Next
			cur.Next.Prev = cur.Prev
			break
		}
		cur=cur.Next	
	}
	list.length--
	return cur.Data
}

//查找值data
//???是返回查找成功与否还是返回下标index
func (list *DList) Search(data Object) bool{
	if list.length == 0 {
		return false
	}
	cur := list.head
	//count :=0
	//index :=[]int{}
	//isSearch := false
	// if cur.Data == data {
	// 	return count
	// }
	for cur != nil{
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
func (this *DList) ShowList(){
	fmt.Println("该双向链表是：")
	if this.IsEmpty(){
		fmt.Println("空链表！")
	}
	cur := this.head
	for cur != nil{
		fmt.Printf("%v ", cur.Data)
		cur = cur.Next
	}
	fmt.Println("")
}


//草稿

/*
//插入排序
func (list *DList) Sort() {
	if list.length <= 1 {
		return
	}
	list.Insert(0, 0) //哨兵头节点
	fmt.Printf(" 哨兵： ")
	list.ShowList()
	cur := list.head.Next.Next
	//i := 1
	for i := 2; i < list.length; i++ {
		j := i - 1
		list.head.Data = cur.Data
		var value int
		//value = cur.Data.(int) //类型断言value
		value = list.head.Data.(int)

		//fmt.Printf("value before=%v", value)
		temp := cur.Prev //定义在左边已经排序部分的指针
		var t int
		t = temp.Data.(int)
		for temp.Prev != nil {
		//for j > 0 {
			 //类型断言temp
			fmt.Printf(" t before=%v ", t)
			if t > value {
				temp = temp.Prev
				j--
			}else {
				break
			}
		}
		list.ShowList()
		list.DeleteIndex(i)
		fmt.Printf(" value=%v ", value)
		list.Insert(j, value)

		cur = cur.Next
		//i++
	}
	return
}*/
