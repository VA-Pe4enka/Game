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
	bee.BossDamage = 1
}

func (bee *GiantBee) SetSpeed() {
	bee.BossSpeed = 1
}

func (r *CrazyRabbit) SetHealth() {
	r.BossHealth = 150
}

func (r *CrazyRabbit) SetDamage() {
	r.BossDamage = 1
}

func (r *CrazyRabbit) SetSpeed() {
	r.BossSpeed = 1
}

func (bear *BearRod) SetHealth() {
	bear.BossHealth = 200
}

func (bear *BearRod) SetDamage() {
	bear.BossDamage = 1
}

func (bear *BearRod) SetSpeed() {
	bear.BossSpeed = 1
}
