package characters

type Scout struct {
	Health   int
	Armoured int
	Damage   int
	Speed    int
}

type Hunter struct {
	Health   int
	Armoured int
	Damage   int
	Speed    int
}

type Paladin struct {
	Health   int
	Armoured int
	Damage   int
	Speed    int
}

func (s *Scout) SetHealth() {
	s.Health = 100
	return
}

func (s *Scout) SetArmour() {
	s.Armoured = 25
	return
}

func (s *Scout) SetDamage() {
	s.Damage = 1000
	return
}

func (s *Scout) SetSpeed() {
	s.Speed = 10
	return
}

func (h *Hunter) SetHealth() {
	h.Health = 150
	return
}

func (h *Hunter) SetArmour() {
	h.Armoured = 50
	return
}

func (h *Hunter) SetDamage() {
	h.Damage = 1
	return
}

func (h *Hunter) SetSpeed() {
	h.Speed = 7
	return
}

func (p *Paladin) SetHealth() {
	p.Health = 200
	return
}

func (p *Paladin) SetArmour() {
	p.Armoured = 100
	return
}

func (p *Paladin) SetDamage() {
	p.Damage = 1
	return
}

func (p *Paladin) SetSpeed() {
	p.Speed = 5
	return
}
