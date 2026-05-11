package entity

type Queue struct {
	tracks []Track
}

func NewQueue() Queue {
	return Queue{
		tracks: []Track{},
	}
}
