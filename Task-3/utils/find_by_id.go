package utils

type HasID interface {
	GetID() int
}

func FindByID[T HasID](elements []T, ID int) int {
	for i, elem := range elements {
		if elem.GetID() == ID {
			return i
		}
	}
	return -1
}