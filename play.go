package main

import (
	"Game/characters"
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

var Heal float64
var Armour int
var Weapon int
var Boost int

func CreateItems() {

	Heal := characters.HealthKit{}
	Heal.SetHealing()

	Armour := characters.Armour{}
	Armour.SetRainForce()

	Weapon := characters.Weapon{}
	Weapon.SetDamageBoost()

	Boost := characters.Poison{}
	Boost.SetSpeedBoost()

}

var BeeHealth int
var BeeDamage int
var BeeSpeed int
var RabbitHealth int
var RabbitDamage int
var RabbitSpeed int
var BearHealth int
var BearDamage int
var BearSpeed int

func CreateBosses() {

	bee := characters.GiantBee{}
	bee.SetHealth()
	BeeHealth = bee.BossHealth
	bee.SetDamage()
	BeeDamage = bee.BossDamage
	bee.SetSpeed()
	BeeSpeed = bee.BossSpeed

	rabbit := characters.CrazyRabbit{}
	rabbit.SetHealth()
	RabbitHealth = rabbit.BossHealth
	rabbit.SetDamage()
	RabbitDamage = rabbit.BossDamage
	rabbit.SetSpeed()
	RabbitSpeed = rabbit.BossSpeed

	bear := characters.BearRod{}
	bear.SetHealth()
	BearHealth = bear.BossHealth
	bear.SetDamage()
	BearDamage = bear.BossDamage
	bear.SetSpeed()
	BearSpeed = bear.BossSpeed

}

var Health int
var Armored int
var Damage int
var Speed int

func ChooseHero() {

	fmt.Println("Выберете персонажа:")
	fmt.Println("1. Разведчик")
	fmt.Println("2. Охотник")
	fmt.Println("3. Паладин")

	reader := bufio.NewReader(os.Stdin)
	choiceStr, _ := reader.ReadString('\n')
	choiceStr = strings.TrimSpace(choiceStr)

	choice, _ := strconv.ParseFloat(choiceStr, 32)

	switch choice {

	case 1:
		fmt.Println("Выбран персонаж Разведчик")
		Hero := characters.Scout{}
		Hero.SetHealth()
		Hero.SetArmour()
		Hero.SetDamage()
		Hero.SetSpeed()

	case 2:
		fmt.Println("Выбран персонаж Охотник")
		Hero := characters.Hunter{}
		Hero.SetHealth()
		Hero.SetArmour()
		Hero.SetDamage()
		Hero.SetSpeed()

	case 3:
		fmt.Println("Выбран персонаж Паладин")
		Hero := characters.Paladin{}
		Hero.SetHealth()
		Hero.SetArmour()
		Hero.SetDamage()
		Hero.SetSpeed()

	default:
		fmt.Println("Такого персонажа пока нет :)")
	}

}

func Healing() {
	Health = int(float64(Health) * Heal)
}

func PlusDamage() {
	Damage = +Weapon
}

func RainForce() {
	Armored = +Armour
}

func SpeedUp() {
	Speed = +Boost
}

func main() {

	CreateItems()

	CreateBosses()

	ChooseHero()

}
