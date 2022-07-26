package player

type Player struct {
	x int
	y int
}

func (p *Player) GetPos() (int, int) {
	return p.x, p.y
}

func (p *Player) SetPos(x, y int) bool {

}
