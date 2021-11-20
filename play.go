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

var Heal int
var Armour int
var Weapon int
var Boost int

func CreateItems() {

	heal := characters.HealthKit{}
	heal.SetHealing()
	fmt.Println("Бонус расходника: ", heal, "хп" )
	Heal = heal.Heal

	armour := characters.Armour{}
	armour.SetRainForce()
	fmt.Println("Бонус расходника: ", armour, "брони")
	Armour = armour.RainForce

	weapon := characters.Weapon{}
	weapon.SetDamageBoost()
	fmt.Println("Бонус расходника: ", weapon, "урона")
	Weapon = weapon.DamageBoost

	boost := characters.Poison{}
	boost.SetSpeedBoost()
	fmt.Println("Бонус расходника: ", boost, "скорости" )
	Boost = boost.SpeedBoost

}

var BossHealth int
var BossDamage int
var BossSpeed int


var Health int
var Armored int
var Damage int
var Speed int

func ChooseHero() {
pers:
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
		fmt.Println("Здоровье персонажа :", Health)
		Hero.SetArmour()
		Armored = Hero.Armoured
		fmt.Println("Защита персонажа :", Armored)
		Hero.SetDamage()
		Damage = Hero.Damage
		fmt.Println("Урон персонажа :", Damage)
		Hero.SetSpeed()
		Speed = Hero.Speed
		fmt.Println("Скорость персонажа :", Speed)

	case 2:
		fmt.Println("Выбран персонаж Охотник")
		Hero := characters.Hunter{}
		Hero.SetHealth()
		Health = Hero.Health
		fmt.Println("Здоровье персонажа :", Health)
		Hero.SetArmour()
		Armored = Hero.Armoured
		fmt.Println("Защита персонажа :", Armored)
		Hero.SetDamage()
		Damage = Hero.Damage
		fmt.Println("Урон персонажа :", Damage)
		Hero.SetSpeed()
		Speed = Hero.Speed
		fmt.Println("Скорость персонажа :", Speed)

	case 3:
		fmt.Println("Выбран персонаж Паладин")
		Hero := characters.Paladin{}
		Hero.SetHealth()
		Health = Hero.Health
		fmt.Println("Здоровье персонажа :", Health)
		Hero.SetArmour()
		Armored = Hero.Armoured
		fmt.Println("Защита персонажа :", Armored)
		Hero.SetDamage()
		Damage = Hero.Damage
		fmt.Println("Урон персонажа :", Damage)
		Hero.SetSpeed()
		Speed = Hero.Speed
		fmt.Println("Скорость персонажа :", Speed)

	default:
		fmt.Println("Такого персонажа пока нет :)")
		goto pers
	}

}

func Healing() {
	Health = Health + Heal
	fmt.Println("Добавлено ", Heal, "хп")
}

func PlusDamage() {
	Damage = Damage + Weapon
	fmt.Println("Добавлено ", Weapon, "урона" )
}

func RainForce() {
	Armored = Armored + Armour
	fmt.Println("Добавлено ", Armour, "брони")
}

func SpeedUp() {
	Speed = Speed + Boost
	fmt.Println("Добавлено ", Boost, "скорости")
}

func Fight() {

	bee := characters.GiantBee{}
	bee.SetHealth()
	bee.SetDamage()
	bee.SetSpeed()

	fmt.Println("Первый противник - Королева пчел!")
	BossHealth = bee.BossHealth
	fmt.Println("Здоровье противника: ", BossHealth)
	BossDamage = bee.BossDamage
	fmt.Println("Урон противника: ", BossDamage)
	BossSpeed = bee.BossSpeed
	fmt.Println("Скорость противника: ", BossSpeed)

	Battle()

	fmt.Println("Поебда... Впереди еще есть враги...")

	rabbit := characters.CrazyRabbit{}
	rabbit.SetHealth()
	rabbit.SetDamage()
	rabbit.SetSpeed()

	fmt.Println("Второй противник - Бешеный кролик!")
	BossHealth = rabbit.BossHealth
	fmt.Println("Здоровье противника: ", BossHealth)
	BossDamage = rabbit.BossDamage
	fmt.Println("Урон противника: ", BossDamage)
	BossSpeed = rabbit.BossSpeed
	fmt.Println("Скорость противника:", BossSpeed)

	Battle()

	fmt.Println("Очередная битва позади... Остался послсдений противник!")

	bear := characters.BearRod{}
	bear.SetHealth()
	bear.SetDamage()
	bear.SetSpeed()

	fmt.Println("Выбран противник - Медведь шатун!")
	BossHealth = bear.BossHealth
	fmt.Println("Здоровье противника: ", BossHealth)
	BossDamage = bear.BossDamage
	fmt.Println("Урон противника: ", BossHealth)
	BossSpeed = bear.BossSpeed
	fmt.Println("Скорость противника: ", BossSpeed)

	Battle()

	fmt.Println("ПОБЕДА!")
}

func Battle(){

	fmt.Println("Битва начинается!!!")
	fmt.Println("Выбирете действие:")

	var Action []string
	Action = append(Action, "Атаковать")
	Action = append(Action, "Использовать расходник")
	Action = append(Action, "Отступить")

Act:
	fmt.Println("1. ", Action[0])
	fmt.Println("2. ", Action[1])
	fmt.Println("3. ", Action[2])

	reader := bufio.NewReader(os.Stdin)
	moveStr, _ := reader.ReadString('\n')
	moveStr = strings.TrimSpace(moveStr)

	move, _ := strconv.ParseFloat(moveStr, 32)

	switch move {

	case 1:
		sec := time.Now().Unix()
		rand.Seed(sec)
		chance := rand.Intn(10) + 1
		enemyChance := rand.Intn(10) + 1

		if chance <= Speed {
			BossHealth = BossHealth - Damage
		} else {
			fmt.Println("Ваш персонаж не успел за уворачивающимся противником...")
		}
		if enemyChance < BossSpeed{
			Health = Health - BossDamage
		} else {
			fmt.Println("Вы смогли увернуться от атаки противника!")
		}


		fmt.Println("Ваше здоровье: ", Health)
		fmt.Println("Здоровье противника: ", BossHealth)

		if Health < 1{
			fmt.Println("Вы погибли!")
			os.Exit(0)
		} else if BossHealth < 1 {
			fmt.Println("Босс повержен!")
		} else {
			goto Act
		}

	case 2:
		if strings.Contains(Action[1], "использовано") {
			fmt.Println("Вы уже использовали расходник!")
			goto Act
		} else {
			Action[1] = Action[1] + "(использовано)"
			fmt.Println("Выбирете расходник:")
			fmt.Println("1. Аптечка")
			fmt.Println("2. Заточка оружия")
			fmt.Println("3. Улучшить броню")
			fmt.Println("4. Зелье скорости (повышенный шанс нанести урон)")

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

			fmt.Println("Здоровье персонажа: ", Health)
			fmt.Println("Урон персонажа: ", Damage)
			fmt.Println("Броня персонажа: ", Armored)
			fmt.Println("Скорость персонажа: ", Speed)

			goto Act
		}

	case 3:
		break


	default:
		fmt.Println("Такого дейстия пока нет :)")

	}

}


func main() {

	CreateItems()

	ChooseHero()

	Fight()
}
