package main

import (
	"Game/characters"
	"bufio"
	"fmt"
	"math/rand"
	"os"
	"strconv"
	"strings"
	"time"
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


var BossHealth int
var BossDamage int
var BossSpeed int
func CreateBosses() {

	bee := characters.GiantBee{}
	bee.SetHealth()

	bee.SetDamage()

	bee.SetSpeed()


	rabbit := characters.CrazyRabbit{}
	rabbit.SetHealth()

	rabbit.SetDamage()

	rabbit.SetSpeed()


	bear := characters.BearRod{}
	bear.SetHealth()

	bear.SetDamage()

	bear.SetSpeed()


	fmt.Println("Выбирете противника:")
	fmt.Println("1. Королева пчел")
	fmt.Println("2. Бешеный кролик")
	fmt.Println("3. Медведь шатун")

	reader := bufio.NewReader(os.Stdin)
	choiceStr, _ := reader.ReadString('\n')
	choiceStr = strings.TrimSpace(choiceStr)

	choice, _ := strconv.ParseFloat(choiceStr, 32)

	switch choice {
	case 1:
		fmt.Println("Выбран противник - Королева пчел!")
		BossHealth = bee.BossHealth
		fmt.Println("Здоровье противника: ", BossHealth)
		BossDamage = bee.BossDamage
		fmt.Println("Урон противника: ", BossDamage)
		BossSpeed = bee.BossSpeed
		fmt.Println("Скорость противника: ", BossSpeed)
	case 2:
		fmt.Println("Выбран противник - Бешеный кролик!")
		BossHealth = rabbit.BossHealth
		fmt.Println("Здоровье противника: ", BossHealth)
		BossDamage = rabbit.BossDamage
		fmt.Println("Урон противника: ", BossDamage)
		BossSpeed = rabbit.BossSpeed
		fmt.Println("Скорость противника:", BossSpeed)

	case 3:
		fmt.Println("Выбран противник - Медведь шатун!")
		BossHealth = bear.BossHealth
		fmt.Println("Здоровье противника: ",BossHealth)
		BossDamage = bear.BossDamage
		fmt.Println("Урон противника: ",BossHealth )
		BossSpeed = bear.BossSpeed
		fmt.Println("Скорость противника: ", BossSpeed)



	}


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
		Health = Hero.Health
		fmt.Println("Здоровье персонажа :", Health )
		Hero.SetArmour()
		Armored = Hero.Armoured
		fmt.Println("Защита персонажа :", Armored )
		Hero.SetDamage()
		Damage = Hero.Damage
		fmt.Println("Урон персонажа :", Damage )
		Hero.SetSpeed()
		Speed = Hero.Speed
		fmt.Println("Скорость персонажа :", Speed )

	case 2:
		fmt.Println("Выбран персонаж Охотник")
		Hero := characters.Hunter{}
		Hero.SetHealth()
		Health = Hero.Health
		fmt.Println("Здоровье персонажа :", Health )
		Hero.SetArmour()
		Armored = Hero.Armoured
		fmt.Println("Защита персонажа :", Armored )
		Hero.SetDamage()
		Damage = Hero.Damage
		fmt.Println("Урон персонажа :", Damage )
		Hero.SetSpeed()
		Speed = Hero.Speed
		fmt.Println("Скорость персонажа :", Speed )

	case 3:
		fmt.Println("Выбран персонаж Паладин")
		Hero := characters.Paladin{}
		Hero.SetHealth()
		Health = Hero.Health
		fmt.Println("Здоровье персонажа :", Health )
		Hero.SetArmour()
		Armored = Hero.Armoured
		fmt.Println("Защита персонажа :", Armored )
		Hero.SetDamage()
		Damage = Hero.Damage
		fmt.Println("Урон персонажа :", Damage )
		Hero.SetSpeed()
		Speed = Hero.Speed
		fmt.Println("Скорость персонажа :", Speed )

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





func Fight(){

	fmt.Println("Битва начинается!!!")
	fmt.Println("Выбирете действие:")
	fmt.Println("1. Атаковать")
	fmt.Println("2. Использовать расходник")
	fmt.Println("3. Отступить")

	reader := bufio.NewReader(os.Stdin)
	moveStr, _ := reader.ReadString('\n')
	moveStr = strings.TrimSpace(moveStr)

	move, _ := strconv.ParseFloat(moveStr, 32)

	switch move {

	case 1:
		sec := time.Now().Unix()
		rand.Seed(sec)
		chance := rand.Intn(10) + 1

		if chance >= Speed{
			BossHealth =- Damage
		} else {
			fmt.Println("Ваш персонаж не успел за уворачивающимся противником...")
		}
		Health =- BossDamage

	case 2:
		fmt.Println("Выбирете расходник:")
		fmt.Println("1. Аптечка")
		fmt.Println("2. Заточка оружия")
		fmt.Println("3. Улучшить броню")
		fmt.Println("4. Зелье скорости (повышенный шанс нанести урон")

		reader := bufio.NewReader(os.Stdin)
		choiceStr, _ := reader.ReadString('\n')
		choiceStr = strings.TrimSpace(choiceStr)

		choice, _ := strconv.ParseFloat(choiceStr, 32)

		switch choice {
		case 1:
			Healing()

		case 2:
			PlusDamage()

		case 3:
			RainForce()

		case 4:
			SpeedUp()

		}

	case 3:
		break

	}


}


func main() {

	CreateItems()

	ChooseHero()

	CreateBosses()

	Fight()
}
