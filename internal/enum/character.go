package enum

type CharacterStatus string

const (
	CharacterStatusActive   CharacterStatus = "ACTIVE"
	CharacterStatusInactive CharacterStatus = "INACTIVE"
	CharacterStatusDeleted  CharacterStatus = "DELETED"
)

func (s CharacterStatus) String() string {
	return string(s)
}

func (s CharacterStatus) isValid() bool {
	switch s {
	case CharacterStatusActive, CharacterStatusInactive, CharacterStatusDeleted:
		return true
	default:
		return false
	}
}
