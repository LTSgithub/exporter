package event

import "context"

type Event struct {
	Id int `json:"id"`
	Type string `json:"type"`
	Tittle string `json:"tittle"`
	Level  int `json:"level"`
	Content string `json:"content"`
}

func Create(ctx context.Context,event *Event) error  {



	return nil
}

func Delete(ctx context.Context,id int) error  {

	return nil
}

func Get(ctx context.Context,id int) error  {


	return nil
}

func List(ctx context.Context, pageIndex int, pageSize int) ([]Event,error)  {
	var resp  []Event


	return resp,nil
}

