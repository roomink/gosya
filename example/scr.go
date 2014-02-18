package main

import (
	"github.com/roomink/gosya"
	"log"
)

type DB struct {
	Host string `yaml:"host"`
	User string `yaml:"user"`
	Pswd string `yaml:"pswd"`
}

type Settings struct {
	Database DB `yaml:"database"`
}

func main() {
	log.Println("Begin")
	path := "./"
	env := "development"
	s := Settings{}
	err := gosya.Merge(&s, path, env)
	if err != nil {
		log.Printf("Error fetch settings: %v", err)
	}
	log.Printf("Settings: %#v", s)
}
