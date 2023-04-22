package main

import "github.com/guioliunb/Ledger-API/back-end/server"


func main(){

	s := server.NewServer()
	if err := s.Init(6000); err != nil{
		panic(err)
	}

	s.Start()
}