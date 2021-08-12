package main

import (
	"flag"
	_ "go.mongodb.org/mongo-driver/bson"
	"log"
	"study/cmd"
)

func main1() {
	var name string
	flag.Parse()
	args := flag.Args()
	if len(args) <= 0 {
		return
	}

	switch args[0] {
	case "go":
		goCmd := flag.NewFlagSet("go", flag.ExitOnError)
		goCmd.StringVar(&name, "name", "go语言", "帮助信息")
		_ = goCmd.Parse(args[1:])
	case "php":
		phpCmd := flag.NewFlagSet("php", flag.ExitOnError)
		phpCmd.StringVar(&name, "name", "php语言", "帮助信息")
		_ = phpCmd.Parse(args[1:])
	}
	log.Printf("name:%s", name)
}

func main() {
	err := cmd.Execute()
	if err != nil {
		log.Fatal("cmd.Execute err: %v", err)
	}
}
