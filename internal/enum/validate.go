package enum

type validatable interface {
	isValid() bool
}

func IsValid[T validatable](t T) bool {
	return t.isValid()
}
