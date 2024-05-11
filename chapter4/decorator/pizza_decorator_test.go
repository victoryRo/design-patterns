package decorator

import (
	"strings"
	"testing"
)

func TestPizzaDecorator_AddIngredient(t *testing.T) {
	pizza := &PizzaDecorator{}

	exp := ResultPizzaDecorator
	actual, _ := pizza.AddIngredient()

	if !strings.Contains(actual, exp) {
		t.Errorf("When calling the add ingredient of the pizza decorator"+
			"it must return the text %s not '%s'", exp, actual)
	}
}

func TestOnion_AddIngredient(t *testing.T) {
	t.Run("Check the error correctly", func(t *testing.T) {
		onion := &Onion{}

		actual, err := onion.AddIngredient()
		if err == nil {
			t.Errorf("When calling AddIngredient on the onion decorator without"+
				"an IngredientAdd on its Ingredient field must return an error,"+
				" not a string with '%s'", actual)
		}
	})

	t.Run("Step of the expected type", func(t *testing.T) {
		onion := &Onion{&PizzaDecorator{}}

		actual, err := onion.AddIngredient()
		if err != nil {
			t.Error(err)
		}

		if !strings.Contains(actual, "onion") {
			t.Errorf("When calling the add ingredient of the onion decorator it"+
				"must return a text with the word 'onion', not '%s'", actual)
		}
	})
}

func TestMeat_AddIngredient(t *testing.T) {
	t.Run("Check the error correctly", func(t *testing.T) {
		meat := new(Meat)

		actual, err := meat.AddIngredient()
		if err == nil {
			t.Errorf("When calling AddIngredient on the meat decorator without"+
				"an IngredientAdd in its Ingredient field must return an error,"+
				"not a string with '%s'", actual)
		}
	})

	t.Run("Step of the expected type", func(t *testing.T) {
		meat := &Meat{&PizzaDecorator{}}

		actual, err := meat.AddIngredient()
		if err != nil {
			t.Error(err)
		}

		if !strings.Contains(actual, "meat") {
			t.Errorf("When calling the add ingredient of the meat decorator it"+
				"must return a text with the word 'meat', not '%s'", actual)
		}
	})
}

func TestPizzaDecorator_FullStack(t *testing.T) {
	pizza := &Onion{&Meat{&PizzaDecorator{}}}

	actual, err := pizza.AddIngredient()
	if err != nil {
		t.Error(err)
	}

	exp := "Pizza with the following ingredients: meat, onion"
	if !strings.Contains(actual, exp) {
		t.Errorf("When asking for a pizza with onion and meat the returned "+
			"string must contain the text '%s' but '%s' didn't have it", exp, actual)
	}

	t.Log(actual)
}
