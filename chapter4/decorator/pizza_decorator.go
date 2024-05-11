package decorator

import (
	"errors"
	"fmt"
)

type IngredientAdder interface {
	AddIngredient() (string, error)
}

// --------------------------------------------------------

// PizzaDecorator tipo base
type PizzaDecorator struct {
	Ingredient IngredientAdder
}

var ResultPizzaDecorator = "Pizza with the following ingredients:"

func (p *PizzaDecorator) AddIngredient() (string, error) {
	return ResultPizzaDecorator, nil
}

// --------------------------------------------------------

// -- first decorator

type Onion struct {
	Ingredient IngredientAdder
}

func (o *Onion) AddIngredient() (string, error) {
	if o.Ingredient == nil {
		return "", errors.New("An IngredientAdd is needed in the Ingredient field of the Onion")
	}

	s, err := o.Ingredient.AddIngredient()
	if err != nil {
		return "", err
	}

	return fmt.Sprintf("%s %s,", s, "onion"), nil
}

// -- second decorator

type Meat struct {
	Ingredient IngredientAdder
}

func (m *Meat) AddIngredient() (string, error) {
	if m.Ingredient == nil {
		return "", errors.New("An IngredientAdd is needed in the Ingredient field of the Meat")
	}

	s, err := m.Ingredient.AddIngredient()
	if err != nil {
		return "", nil
	}

	return fmt.Sprintf("%s %s,", s, "meat"), nil
}

// --------------------------------------------------------
