package main

import (
	"fmt"
	"log"
)

func main() {
	const LOGO = "\n    _   __      __   _ ____      ______    \n   / | / /___  / /__(_) / /     / ____/___ \n  /  |/ / __ \\/ //_/ / / /_____/ / __/ __ \\\n / /|  / /_/ / ,< / / / /_____/ /_/ / /_/ /\n/_/ |_/ .___/_/|_/_/_/_/      \\____/\\____/ \n     /_/                                   \n"
	fmt.Println(LOGO)
	dirList, err := GetDirList()
	if err != nil {
		log.Fatalln(err)
	}
	interactive(dirList)

}
