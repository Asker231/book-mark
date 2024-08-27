package main

import (
	"fmt"
	"time"

	"github.com/fatih/color"
)

	type BookMark = map[string]string

    func main(){
		books := BookMark{}
		loop:
		for{
			result := Menu()
		 switch result{
			case 1:
				WatchMark(books)
				break
			case 2:
			    AddMark(books)
				break
			case 3:
				DeleteMark(books)
				break
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
			time.Sleep(500 * time.Millisecond)
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

	func WatchMark(books BookMark){
		if len(books) == 0{
			fmt.Println("Нет закладок")
		}
		for key,value := range books{
			fmt.Println(key + " : " + value)
		}
	}

	func AddMark(books BookMark){
		var key,value string
		fmt.Print("Введите название: ")
		fmt.Scan(&key)
		fmt.Print("Введите ссылку: ")
		fmt.Scan(&value)
		books[key] = value

	}
	func DeleteMark(books BookMark){
		var key string
		fmt.Print("Введите ключ для удаления: ")
		fmt.Scan(&key)
		for k := range books{
			if k == key{
				delete(books,k)
			}
		}	
	}

	func Out(){
		fmt.Println("Вы вышли")
	}


	func Load() *color.Color{
		c := color.New(color.FgCyan)
		return c
	}

