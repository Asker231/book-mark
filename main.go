package main

import (
	"fmt"
	"time"

	"github.com/fatih/color"
)

	type BookMark = map[string]string

    func main(){
		books := BookMark{}
		result := Menu()
		loop:
		for{
		 switch result{
			case 1:
				WatchMark(&books)
				break loop
			case 2:
			    AddMark(&books)
			break loop

			case 3:
				DeleteMark(&books)
				break loop

			case 4:
				Out()
				break loop

			default:
				fmt.Println("Введите корректные данные")
				break loop
		 }
		}
    }

	func Menu()int{
		var choice int
		c := Load()
		color.Blue("Меню")
		for range 3{
			time.Sleep(1 * time.Second)
			c.Print(" _ ")
		}
		fmt.Println("")
		color.Green("Выберите действие:")
		color.Green("1) Посмотреть закладки")
		color.Green("2) Добавить закладку")
		color.Green("3) Удалить закладку")
		color.Red("4) Выход")
		fmt.Scan(&choice)
		return choice
	}

	func WatchMark(books *BookMark){
		fmt.Println("wathc")
	}

	func AddMark(books *BookMark){
		fmt.Println("Add")
	}

	func DeleteMark(books *BookMark){
		fmt.Println("Delete")
	}

	func Out(){
		fmt.Println("Out")
	}


	func Load() *color.Color{
		c := color.New(color.FgCyan).Add(color.Underline)
		return c
	}

