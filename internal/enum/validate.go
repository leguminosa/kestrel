package enum

type validateable interface {
	isValid() bool
}

func IsValid[T validateable](t T) bool {
	return t.isValid()
}
