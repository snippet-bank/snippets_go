package snippets

import "testing"

/*

There are a couple of test-naming styles to choose from.

One style is the usual Test{FunctionName}.
E.g., "TestDefend()".
In this style, the fact that the function is a method is disregarded.

Another style is Test{StructName}_{MethodName}
E.g., "TestPlayer_Defend()"
In this style, the function being a method on a struct
is made very explicit.

The two tests below do exactly the same thing, just illustrate
the two naming styles.
*/

func TestDefend(t *testing.T) {
	player1 := Player{
		HealthPoints:  10,
		DefensePoints: 1,
		AttackPoints:  3,
	}
	player1.defend(5)

	// attack of 5, defends 1, loses 4 points
	got := player1.HealthPoints
	want := 6
	if got != want {
		t.Errorf("Got %v health points but expected %d", got, want)
	}
}

func TestPlayer_Defend(t *testing.T) {
	player1 := Player{
		HealthPoints:  10,
		DefensePoints: 1,
		AttackPoints:  3,
	}
	player1.defend(5)

	// attack of 5, defends 1, loses 4 points
	got := player1.HealthPoints
	want := 6
	if got != want {
		t.Errorf("Got %v health points but expected %d", got, want)
	}
}
