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
	Heal = heal.Heal

	armour := characters.Armour{}
	armour.SetRainForce()
	Armour = armour.RainForce

	weapon := characters.Weapon{}
	weapon.SetDamageBoost()
	Weapon = weapon.DamageBoost

	boost := characters.Poison{}
	boost.SetSpeedBoost()
	Boost = boost.SpeedBoost

}

var BossHealth int
var BossDamage int
var BossSpeed int

var Health int
var OriginalHealth int
var Armored int
var Damage int
var Speed int

var Heroes[]string

func ChooseHero() {
	Heroes =append(Heroes, "Разведчик")
	Heroes =append(Heroes,  "Охотник")
	Heroes =append(Heroes,  "Паладин")
	Heroes =append(Heroes,  "??????")

Pers:
	fmt.Println("Выберете персонажа:")
	for i := range Heroes{
		fmt.Println(i+1,". ", Heroes[i])
	}



	reader := bufio.NewReader(os.Stdin)
	choiceStr, _ := reader.ReadString('\n')
	choiceStr = strings.TrimSpace(choiceStr)

	choice, _ := strconv.ParseFloat(choiceStr, 32)

	switch choice {

	case 1:
		fmt.Println("Выбран персонаж Разведчик")
		fmt.Println()
		fmt.Println("_______________________________")
		fmt.Println()
		Hero := characters.Scout{}
		Hero.SetHealth()
		Health = Hero.Health
		OriginalHealth = Health
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
		fmt.Println()
		fmt.Println("_______________________________")
		fmt.Println()


	case 2:
		fmt.Println("Выбран персонаж Охотник")
		fmt.Println()
		fmt.Println("_______________________________")
		fmt.Println()
		Hero := characters.Hunter{}
		Hero.SetHealth()
		Health = Hero.Health
		OriginalHealth = Health
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
		fmt.Println()
		fmt.Println("_______________________________")
		fmt.Println()

	case 3:
		fmt.Println("Выбран персонаж Паладин")
		fmt.Println()
		fmt.Println("_______________________________")
		fmt.Println()
		Hero := characters.Paladin{}
		Hero.SetHealth()
		Health = Hero.Health
		OriginalHealth = Health
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
		fmt.Println()
		fmt.Println("_______________________________")
		fmt.Println()

	case 4:
		if strings.Contains(Heroes[3], "больше не доступен"){
			fmt.Println("Этот персонаж больше вам не доступен! Выбирете другого!")
			fmt.Println()
			fmt.Println("_______________________________")
			fmt.Println()

			goto Pers
		} else{
			fmt.Println("Только самый достойный сможет выбрать этого персонажа!!!")
			fmt.Println()
			fmt.Println("Вам предстоит испытание...")
			fmt.Println()
			fmt.Println("Дайте ответ на вопрос: Какой у вас любимый цвет?")
			reader := bufio.NewReader(os.Stdin)
			choiceStr, _ := reader.ReadString('\n')
			choiceStr = strings.ToLower(choiceStr)
			if strings.Contains(choiceStr, "фиолетовый") {
				fmt.Println()
				fmt.Println("_______________________________")
				fmt.Println()
				fmt.Println("Вы оказались достойны")
				fmt.Println()
				fmt.Println("Выбран персонаж Ёж Русак!")
				fmt.Println()
				fmt.Println("_______________________________")
				fmt.Println()
				Hero := characters.Hedgehog{}
				Hero.SetHealth()
				Health = Hero.Health
				OriginalHealth = Health
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
				fmt.Println()
				fmt.Println("_______________________________")
				fmt.Println()
			} else {
				fmt.Println("Вы оказались не дойстойны!!!")
				fmt.Println()
				fmt.Println("_______________________________")
				fmt.Println()
				Heroes[3] = Heroes[3] + " (больше не доступен)"
				goto Pers
			}
		}




	default:
		fmt.Println("Такого персонажа пока нет :)")
		goto Pers
	}

}

func Healing() {
	Health = Health + Heal
	fmt.Println("Добавлено ", Heal, "хп")
	fmt.Println()
	fmt.Println("_______________________________")
	fmt.Println()
}

func PlusDamage() {
	Damage = Damage + Weapon
	fmt.Println("Добавлено ", Weapon, "урона")
	fmt.Println()
	fmt.Println("_______________________________")
	fmt.Println()
}

func RainForce() {
	Armored = Armored + Armour
	fmt.Println("Добавлено ", Armour, "брони")
	fmt.Println()
	fmt.Println("_______________________________")
	fmt.Println()
}

func SpeedUp() {
	Speed = Speed + Boost
	fmt.Println("Добавлено ", Boost, "скорости")
	fmt.Println()
	fmt.Println("_______________________________")
	fmt.Println()
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
	fmt.Println()
	fmt.Println("_______________________________")
	fmt.Println()

	Battle()

	fmt.Println("Поебда... Впереди еще есть враги...")
	fmt.Println()
	fmt.Println("_______________________________")
	fmt.Println()

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
	fmt.Println()
	fmt.Println("Здоровье персонажа :", Health)
	fmt.Println("Защита персонажа :", Armored)
	fmt.Println("Урон персонажа :", Damage)
	fmt.Println("Скорость персонажа :", Speed)
	fmt.Println()
	fmt.Println("_______________________________")
	fmt.Println()

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
	fmt.Println("Урон противника: ", BossDamage)
	BossSpeed = bear.BossSpeed
	fmt.Println("Скорость противника: ", BossSpeed)
	fmt.Println()
	fmt.Println("Здоровье персонажа :", Health)
	fmt.Println("Защита персонажа :", Armored)
	fmt.Println("Урон персонажа :", Damage)
	fmt.Println("Скорость персонажа :", Speed)
	fmt.Println()
	fmt.Println("_______________________________")
	fmt.Println()

	Battle()

	fmt.Println("ПОБЕДА!")
}

func Battle() {


	fmt.Println("Битва начинается!!!")
	fmt.Println()
	fmt.Println("_______________________________")
	fmt.Println()
	fmt.Println("Выбирете действие:")

	var Action []string
	Action = append(Action, "Атаковать")
	Action = append(Action, "Использовать расходник")
	Action = append(Action, "Отступить")

countMedKit := 3
countSpeed := 1
countArmoury := 2
countWeaponry := 2

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
		if enemyChance < BossSpeed {
			Health = Health - (BossDamage - int(float64(BossHealth * (Armour/100))))
		} else {
			fmt.Println("Вы смогли увернуться от атаки противника!")
		}
		fmt.Println()
		fmt.Println("_______________________________")
		fmt.Println()
		fmt.Println("Ваше здоровье: ", Health)
		fmt.Println("Здоровье противника: ", BossHealth)
		fmt.Println()
		fmt.Println("_______________________________")
		fmt.Println()

		if Health < 1 {
			fmt.Println("Вы погибли!")
			os.Exit(0)
		} else if BossHealth < 1 {
			fmt.Println("Босс повержен!")
		} else {
			Action[1] = "Использовать расходник"
			goto Act
		}

	case 2:
		if strings.Contains(Action[1], "использовано") {
			fmt.Println("Вы уже использовали расходник!")
			fmt.Println()
			fmt.Println("_______________________________")
			fmt.Println()
			goto Act
		} else {
			Action[1] = Action[1] + "(использовано)"

		B:
			fmt.Println()
			fmt.Println("_______________________________")
			fmt.Println()
			fmt.Println("Выбирете расходник:")

			fmt.Println("1. Аптечка. Осталось: ", countMedKit)
			fmt.Println("2. Заточка оружия. Осталось:", countWeaponry)
			fmt.Println("3. Улучшить броню. Осталось:", countArmoury)
			fmt.Println("4. Зелье скорости (повышенный шанс нанести урон). Осталось:", countSpeed)
			fmt.Println("5. Отмена")

			reader := bufio.NewReader(os.Stdin)
			choiceStr, _ := reader.ReadString('\n')
			choiceStr = strings.TrimSpace(choiceStr)

			choice, _ := strconv.ParseFloat(choiceStr, 32)

			switch choice {
			case 1:
				if countMedKit > 0 {
					Healing()
					if Health > OriginalHealth{
						Health = OriginalHealth
					}
					countMedKit--
				} else {
					fmt.Println("У вас не осталось аптечек!")
					goto B
				}

			case 2:
				if countWeaponry > 0{
					PlusDamage()
					countWeaponry--
				} else {
					fmt.Println("У вас не осталось заточек!")
					goto B
				}


			case 3:
				if countArmoury > 0 {
					RainForce()
					countArmoury--
				} else{
					fmt.Println("У вас не осталось укреплений для брони!")
					goto B
				}


			case 4:
				if countSpeed > 0 {
					SpeedUp()
					if Speed > 10{
						Speed = 10
					}
					countSpeed--
				} else {
					fmt.Println("У вас не осталось зелий!")
					goto B
				}
			case 5:
				fmt.Println()
				fmt.Println("_______________________________")
				fmt.Println()
				Action[1] = "Использовать расходник"
				goto Act

			default:
				goto B


			}
			fmt.Println()
			fmt.Println("_______________________________")
			fmt.Println()
			fmt.Println("Здоровье персонажа: ", Health)
			fmt.Println("Урон персонажа: ", Damage)
			fmt.Println("Броня персонажа: ", Armored)
			fmt.Println("Скорость персонажа: ", Speed)
			fmt.Println()
			fmt.Println("_______________________________")
			fmt.Println()

			goto Act
		}

	case 3:
		fmt.Println("Вы сбежали с поля боя!")
		os.Exit(0)

	default:
		fmt.Println("Такого дейстия пока нет :)")
		goto Act

	}

}

func main() {

	CreateItems()

	ChooseHero()

	Fight()
}
