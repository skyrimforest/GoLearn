package main

// 0.5-1s接受一次信号 加速模式直接很少的时间

import (
	"fmt"
	"math/rand"
	"os"
	"time"
)

var ALL_SKILL_AVA [1]Skill = [1]Skill{
	NewHit(),
}

type Character struct {
	cname    string
	atk      int
	hth      int
	blu      int
	spd      int
	skl_pool []Skill
}

func (c *Character) useSkill(target *Character) {
	// fmt.Printf("%s starts using skill!\n", c.cname)
	choice := rand.Intn(len(c.skl_pool))
	// fmt.Println(c.skl_pool[choice].getName())
	if c.skl_pool[choice].getType() == "kill" {
		dmg := c.atk * c.skl_pool[choice].getDmg()
		fmt.Printf("%s did %d damage to %s,", c.cname, dmg, target.cname)
		target.hth -= dmg
		fmt.Printf("%d health left.\n", target.hth)
	} else if c.skl_pool[choice].getName() == "effect" {

	}
}

func NewCharacter(name string, at int, ht int, bl int, sp int) *Character {
	return &Character{
		cname:    name,
		atk:      at,
		hth:      ht,
		blu:      bl,
		spd:      sp,
		skl_pool: ALL_SKILL_AVA[:],
	}
}

type Skill interface {
	doSth(*Character)
	getName() string
	getType() string
	getDmg() int
}

// sk_type有2种,有kill,effect
type Hit struct {
	sname   string
	dmg     int
	sk_type string
}

func NewHit() *Hit {
	return &Hit{
		sname:   "Hit",
		dmg:     5,
		sk_type: "kill",
	}
}

// 设定Hit不需要doSth
func (h *Hit) doSth(target *Character) {
}
func (h *Hit) getName() string {
	return h.sname
}
func (h *Hit) getDmg() int {
	return h.dmg
}
func (h *Hit) getType() string {
	return h.sk_type
}

// spd 在40-50之间,spd_pool大小为50
type Ground struct {
	Alice      *Character
	Bob        *Character
	max_pool   int
	spd_pool_A int
	spd_pool_B int
	game_cond  bool
	round      int
	fir_time   map[string]int
}

func initGround() *Ground {
	ch1 := NewCharacter("Alice", 10, 900, 30, 100)
	ch2 := NewCharacter("Bob", 15, 1000, 20, 45)

	fmt.Println(ch1.spd)
	fmt.Println(ch2.spd)
	return &Ground{
		Alice:      ch1,
		Bob:        ch2,
		max_pool:   500,
		spd_pool_A: 0,
		spd_pool_B: 0,
		game_cond:  true,
		round:      0,
		fir_time:   make(map[string]int),
	}
}

func (g *Ground) spdGround(c1, c2 Character) []int {
	fmt.Println("Start spd judge!")
	res := []int{}
	for true {

		g.spd_pool_A = (g.spd_pool_A + c1.spd)
		g.spd_pool_B = (g.spd_pool_B + c2.spd)

		// fmt.Printf("pool_A is %d\n", g.spd_pool_A)
		// fmt.Printf("pool_B is %d\n", g.spd_pool_B)

		if g.spd_pool_A >= g.max_pool && g.spd_pool_B < g.max_pool {
			fmt.Println("Alice is the first this round!0")
			g.fir_time[g.Alice.cname]++
			g.spd_pool_A %= g.max_pool
			res = append(res, 0)
			res = append(res, 1)
			return res

		} else if g.spd_pool_B >= g.max_pool && g.spd_pool_A < g.max_pool {
			fmt.Println("Bob is the first this round!0")
			g.fir_time[g.Bob.cname]++

			g.spd_pool_B %= g.max_pool
			res = append(res, 1)
			res = append(res, 0)
			return res
		} else if g.spd_pool_B >= g.max_pool && g.spd_pool_A >= g.max_pool {
			g.spd_pool_A %= g.max_pool
			g.spd_pool_B %= g.max_pool

			if g.spd_pool_A > g.spd_pool_B {
				fmt.Println("Alice is the first this round!1")
				g.fir_time[g.Alice.cname]++

				res = append(res, 0)
				res = append(res, 1)
				return res
			} else if g.spd_pool_A < g.spd_pool_B {
				fmt.Println("Bob is the first this round!1")
				g.fir_time[g.Bob.cname]++

				res = append(res, 1)
				res = append(res, 0)
				return res
			} else if g.spd_pool_A == g.spd_pool_B {
				for rand.Intn(50) != rand.Intn(50) {
					if rand.Intn(50) > rand.Intn(50) {
						fmt.Println("Alice is the first this round!2")
						g.fir_time[g.Alice.cname]++

						res = append(res, 0)
						res = append(res, 1)
						return res
					} else if rand.Intn(50) < rand.Intn(50) {
						fmt.Println("Bob is the first this round!2")
						g.fir_time[g.Bob.cname]++

						res = append(res, 1)
						res = append(res, 0)
						return res
					} else {
						continue
					}
				}
			}

		}
	}
	return []int{0, 1}
}

func (g *Ground) judgeGround() {
	if g.Alice.hth <= 0 {
		fmt.Println("Alice Lose,Bob Win!")
		g.game_cond = false
	} else if g.Bob.hth <= 0 {
		fmt.Println("Alice Win,Bob Lose!")
		g.game_cond = false
	}
	if g.game_cond == false {
		fmt.Println("Game End!!!")
		g.showRes()
		os.Exit(0)
	}
}

func (g *Ground) showRes() {
	fmt.Printf("At last,%s has %d health\n", g.Alice.cname, g.Alice.hth)
	fmt.Printf("At last,%s has %d health\n", g.Bob.cname, g.Bob.hth)
	fmt.Println(g.fir_time)
}
func (g *Ground) runGround() {
	fmt.Println("Game Start!")
	for true {
		time.Sleep(time.Duration(1) * time.Second)
		g.round++
		fmt.Printf("Round %d started!\n", g.round)

		sequence := g.spdGround(*g.Alice, *g.Bob)
		if sequence[0] == 0 {
			g.Alice.useSkill(g.Bob)
			g.judgeGround()
			// g.Bob.useSkill(g.Alice)
			// g.judgeGround()
		} else {
			g.Bob.useSkill(g.Alice)
			g.judgeGround()
			// g.Alice.useSkill(g.Bob)
			// g.judgeGround()
		}
	}
}

// func main() {
// 	ground := initGround()
// 	ground.runGround()
// }
