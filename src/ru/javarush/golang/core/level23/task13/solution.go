package main

/*
Вы собираете «коробку» для отправки деталей: к ней прикреплены сразу два ярлыка от разных отделов, и у каждого ярлыка свой ID.
Вы решили встроить оба ярлыка как embedded-части, чтобы поля были рядом, но тут появляется реальная проблема: оба ярлыка
содержат поле с одинаковым именем ID, и короткий доступ становится неоднозначным. В такой ситуации Go требует явного пути.

Описаны типы PartA и PartB, у каждого есть поле ID типа int. Создайте тип Box, который встраивает (embedding) оба
типа: PartA и PartB.

Программа читает из stdin два целых числа: aID и bID. Затем она создаёт значение box так, чтобы PartA.ID был
равен aID, а PartB.ID был равен bID.
Программа должна вывести одну строку: A=<aID> B=<bID>. В решении нельзя обращаться к полю ID как box.ID, нужен явный
доступ box.PartA.ID и box.PartB.ID.
*/

import "fmt"

type PartA struct {
	ID int
}

type PartB struct {
	ID int
}

// TODO: Объявите структуру Box так, чтобы она встраивала (embedding) и PartA, и PartB.
type Box struct {
	PartA
	PartB
}

func main() {
	var aID, bID int
	fmt.Scan(&aID, &bID)

	var box Box

	// TODO: Установите ID для обеих частей через явный путь: box.PartA.ID и box.PartB.ID.
	box.PartA.ID = aID
	box.PartB.ID = bID
	// TODO: Выведите ровно одну строку в формате A=<aID> B=<bID> (через явный доступ к полям).
	fmt.Printf("A=%d B=%d\n", box.PartA.ID, box.PartB.ID)
}
