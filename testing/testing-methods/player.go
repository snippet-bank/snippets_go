package snippets

// Player is a character in the game.
type Player struct {
	HealthPoints  int
	DefensePoints int
	AttackPoints  int
}

func (player *Player) defend(attackPoints int) {
	attackPoints -= player.DefensePoints
	player.HealthPoints -= attackPoints
}

func (player *Player) isAlive() bool {
	return player.HealthPoints > 0
}

func attack(attacker Player, defender Player) {
	defender.defend(attacker.AttackPoints)
}
