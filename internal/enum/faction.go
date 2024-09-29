package enum

type Faction string

const (
	FactionCounter Faction = "COUNTER"
	FactionSoldier Faction = "SOLDIER"
	FactionMech    Faction = "MECH"
)

func (f Faction) String() string {
	return string(f)
}

func (f Faction) isValid() bool {
	switch f {
	case FactionCounter, FactionSoldier, FactionMech:
		return true
	default:
		return false
	}
}
