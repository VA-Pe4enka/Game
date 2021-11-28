package characters

type GiantBee struct {
	BossHealth int
	BossDamage int
	BossSpeed  int
}

type CrazyRabbit struct {
	BossHealth int
	BossDamage int
	BossSpeed  int
}

type BearRod struct {
	BossHealth int
	BossDamage int
	BossSpeed  int
}

func (bee *GiantBee) SetHealth() {
	bee.BossHealth = 200

}

func (bee *GiantBee) SetDamage() {
	bee.BossDamage = 13
}

func (bee *GiantBee) SetSpeed() {
	bee.BossSpeed = 7
}

func (r *CrazyRabbit) SetHealth() {
	r.BossHealth = 300
}

func (r *CrazyRabbit) SetDamage() {
	r.BossDamage = 26
}

func (r *CrazyRabbit) SetSpeed() {
	r.BossSpeed = 5
}

func (bear *BearRod) SetHealth() {
	bear.BossHealth = 400
}

func (bear *BearRod) SetDamage() {
	bear.BossDamage = 43
}

func (bear *BearRod) SetSpeed() {
	bear.BossSpeed = 3
}
