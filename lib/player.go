package lib

type Player struct {
	Name string
}

func (p Player) String() string {
	return p.Name
}
