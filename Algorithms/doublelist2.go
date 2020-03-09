package main

import (
	"fmt"
)

func main(){
	l := DoubleList{}

	l.Append(4)
	l.Append(1)
	l.Append(3)
	l.Append(9)
	l.PrintList()
	l.Search(3)
	l.Insert(2, 5)
	fmt.Printf("插入前的链表\n")
	
	l.PrintList()

	fmt.Printf("排序后的链表------\n")
	//l.Reverse()
	l.PrintList()


	
}

type Node struct{
	value interface{}
	Next *Node
	Prev *Node 
}

type DoubleList struct{
	headNode *Node
	tailNode *Node
	length int
}

//初始化链表
func (list *DoubleList) InitList() *DoubleList{
//	list = *(DoubleList)
	list.length = 0
	list.headNode = nil
	list.tailNode = nil
	return list
}


//获取长度
func (list *DoubleList) Getlength() int {
	return (*list).length
}



//打印链表

func (list *DoubleList) PrintList(){
	cur := list.headNode
	len := list.length
	if len == 0{
		fmt.Print("the list is empty")
	}else{
		for cur!=nil{
	
			fmt.Printf("the list is %v\n", cur.value)
			cur = cur.Next
		}
		
	}
}

//向链表赋值
func (list *DoubleList) Append(data interface{}) {
	newNode := &Node{value: data}
	//h := (*list).headNode
	//t := (*list).tailNode
	//如果链表为空
	if list.Getlength() == 0 {
		list.headNode = newNode 
		list.tailNode = newNode
		newNode.Next = nil
		newNode.Prev = nil
		} else{
			
			newNode.Prev = list.tailNode
			newNode.Next = nil
	
			//将尾节点指向新的节点ti
			list.tailNode.Next = newNode
			//更新尾节点
			list.tailNode = newNode

		}
		(*list).length ++
	
}

// func (list *DoubleList) PrintList() {



// }


//查找给定值，并返回对应下标
func (list *DoubleList) Search(data interface{}) int {
	len := list.length
	//node := new(Node)
	node := list.headNode
	if len ==0 {
		fmt.Print("no such data")
		return -1
	}

	for i:=0;i<len-1; i++{
		if node.value != data{
			node = node.Next
		} else{
			fmt.Printf("data is %v, index is %v", node.value, i)
			return i
		}
	}
	return -1
}
/*
//向链表赋值
func (list *DoubleList) Append(data interface{}) {
	newNode := &Node{value: data}
	//h := (*list).headNode
	//t := (*list).tailNode
	//如果链表为空
	if (*list).Getlength() == 0 {
		(*list).headNode = newNode 
		(*list).tailNode = newNode
		(*newNode).Next = nil
		(*newNode).Prev = nil
		} else{
			
			(*newNode).Prev = (*list).tailNode
			(*newNode).Next = nil
	
			//将尾节点指向新的节点
			(*((*list).tailNode)).Next = newNode
			//更新尾节点
			(*list).tailNode = newNode

		}
		(*list).length ++
	
}
*/


//向指定位置插入

func (list *DoubleList) Insert(at int, data interface{}) {
	len := list.length
	curr := list.headNode
	for i:=0;i < at; i++{
		curr= curr.Next
	}

	if  at == 0{
		return
	}

	if at > len-1{
		return 
	}
	if at == len-1{
		list.Append(data)
	} else{
		//1.先将新节点指向下个节点
		newNode := new(Node)
		newNode.value = data
		newNode.Next = curr.Next
		//下个节点指向新节点
		curr.Next.Prev = newNode
		//上个节点指向新节点
		curr.Next = newNode
		//新结点指向上个节点
		newNode.Prev = curr
		len++
	}
	
	
}
	
	//链表反转,反转指针
	//替换值
	/*
	func (list *DoubleList) Reverse()  {
		
		v := list.headNode.value.(int)
		
		fmt.Println(v)

		/*
		if list.headNode == nil || list.headNode.Next == nil{
			return 
		}

	     var dum *Node //哨兵
		curr := list.headNode //头节点
		pre.Next = curr //哨兵指向头节点
		curr.Prev = pre
		for curr != nil{
			//定义第二个节点
			next := curr.Next
			next.Prev = curr
			//将头节点倒序指向哨兵
			// curr.Prev= next
			// next.Next = curr
			// dumpy.Prev = curr

			curr.Next= dummpy
			dummpy.Next =  next
			next.Prev = dummpy
			dummpy.Prev = curr
			
			//第二个节点的反转
	
			curr = next
		}mpy
		
		

	}
*/	

	// func  (list *DoubleList) Sort() {

	// }

	func (list *DoubleList) SortLists(l *DoubleList)  {
		if l.headNode  == nil || l.headNode.Next == nil {
			return
		}

		//找中点
		fast := l.headNode
		slow := l.headNode.Next
		for fast != nil && fast.Next != nil{
			fast = fast.Next.Next
			slow = slow.Next
		} 
			//新建两个双链表
		var left *DoubleList
		left.headNode = l.headNode
		var middle *DoubleList
		middle.headNode = slow.Next
		//切断链表
		slow.Next = nil
		return MergeList (list.SortLists(left), list.SortLists(middle))

	}

	func (list *DoubleList) MergeList(l1 *DoubleList, l2 *DoubleList) *DoubleList {
			var dummpy *DoubleList
			dummpy.headNode.Next = nil
			
			
			for l1 != nil && l2!= nil{
				if l1.headNode.value.(int) <= l2.headNode.value.(int){
						dummpy = l1
						l1.headNode = l1.headNode.Next
				} else{
					 dummpy = l2
					 l2.headNode = l2.headNode.Next 
				}

				 dummpy.headNode.Next = l1.headNode
				 dummpy =l1


				if l1 != nil{
					dummpy.headNode.Next=  l1.headNode
				}
				if l2 != nil{
					dummpy.headNode.Next=  l2.headNode
					
				}
			}

			return dummpy

	}


	
