package service

import "fmt"

type SampleService interface {
	Hello(msg string)
}

/*
サービス
`service:"SampleService"``
*/
type ImplementsService struct{}

func (s *ImplementsService) Hello(msg string) {
	fmt.Printf("Hello %s.\n", msg)
}
