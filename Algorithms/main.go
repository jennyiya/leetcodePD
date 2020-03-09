package main

import(
	"fmt"
	_"go_code/algorithm/linkedList/singleList"
	_"go_code/algorithm/linkedList/doubleList"
	"go_code/algorithm/linkedList/circleList"
)

func main(){
	/*
	list:=singleList.List{}
	list.Append(1)
	list.Append(2)
	list.Append(3)
	list.Append(4)
	list.Append(4)
	list.Append(5)

	//list.Delete(3)
	//s:=list.Delete(4)
	//fmt.Printf("删除值3的index是：%v\n",s)
	list.ShowList()
	//list.PrintList()
	

	//2
	//dlist:=doubleList.DList{}
	dlist := new(doubleList.DList)
	dlist.Init()
	dlist.Append(8)
	dlist.Append(7)
	dlist.Append(9)
	dlist.Append(8)
	dlist.Append(3)
	dlist.Append(6)
	dlist.Append(1)
	//dlist.Append("app")
	dlist.ShowList()
	//dlist.Sort()
	head := doubleList.DNode{}
	dlist.SortMerge(&head)
	// i:=dlist.Insert(0,"4")
	// fmt.Printf("是否成功插入？%v\n", i)
	//dlist.Search(3)
	//d1:=dlist.DeleteIndex(2)
	//fmt.Printf("\n双向链表删除index=2的值是：%v\n",d1)
	// d2:=dlist.DeleteValue(8)
	// fmt.Printf("\n双向链表删除值8的index是：%v\n",d2)
	dlist.ShowList()
*/

	//3
	clist := new(circleList.CList)
	clist.Init()
	clist.Append(11)
	clist.Append(22)
	clist.Append(33)
	clist.Append(44)

	clist.ShowList()

	i:=clist.Insert(0,"7")
	fmt.Printf("是否成功插入？%v\n", i)
	//i: = clist.Search(3)
	//fmt.Printf("是否成功查找？%v\n", i)
	//d1:=clist.DeleteIndex(1)
	//fmt.Printf("\n双向循环链表删除index=2的值是：%v\n",d1)
	// d2:=dlist.DeleteValue(8)
	// fmt.Printf("\n双向链表删除值8的index是：%v\n",d2)
	clist.ShowList()
}
