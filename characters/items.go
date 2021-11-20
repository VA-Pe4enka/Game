package characters

type HealthKit struct {
	Heal int
}

type Armour struct {
	RainForce int
}

type Weapon struct {
	DamageBoost int
}

type Poison struct {
	SpeedBoost int
}

func (h *HealthKit) SetHealing() {
	h.Heal = 75
}

func (a *Armour) SetRainForce() {
	a.RainForce = 20
}

func (w *Weapon) SetDamageBoost() {
	w.DamageBoost = 2
}

func (p *Poison) SetSpeedBoost() {
	p.SpeedBoost = 2
}
