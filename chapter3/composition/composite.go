package composition

import "fmt"

// ---- composicion directa

type Athlete struct{}

func (a *Athlete) Train() {
	fmt.Println("Training...")
}

func Swim() {
	fmt.Println("Swimming!")
}

type CompositeSwimmerA struct {
	MyAthlete Athlete
	MySwim    func()
}

func MainCompositeSwim() {
	localSwim := Swim
	swimmer := CompositeSwimmerA{
		MySwim: localSwim,
	}

	swimmer.MyAthlete.Train()
	swimmer.MySwim()
}

// ------------------------------------------------------------

// --- composicion embebida

type Animal struct{}

func (a *Animal) Eat() {
	fmt.Println("Eating")
}

type Shark struct {
	Animal
	Swim func()
}

func MainCompositeAnimal() {
	fish := Shark{
		Swim: Swim,
	}

	fish.Eat()
	fish.Swim()
}

// ------------------------------------------------------------

// --- composicion embebida con interface

type Swimmer interface {
	Swim()
}

type Trainer interface {
	Train()
}

type SwimmerImplementor struct{}

func (s *SwimmerImplementor) Swim() {
	fmt.Println("Swimming !")
}

type CompositeSwimmerB struct {
	Trainer
	Swimmer
}

// El Swimmer campo está incrustado,
// pero no se inicializará en cero ya que es un puntero a una interfaz.
func MainCompositeB() {
	swimmer := CompositeSwimmerB{
		// &Athlete{},
		new(Athlete),
		&SwimmerImplementor{},
	}

	swimmer.Train()
	swimmer.Swim()
}

// ------------------------------------------------------------

// ---- Composiciones de árboles binarios

type Tree struct {
	LeafValue int
	Right     *Tree
	Left      *Tree
}

func MainTree() {
	tree := Tree{
		LeafValue: 0,
		Right: &Tree{
			LeafValue: 5,
			Right:     &Tree{6, nil, nil},
			Left:      nil,
		},
		Left: &Tree{4, nil, nil},
	}

	fmt.Println(tree.Right.Right.LeafValue)
}

// ------------------------------------------------------------

// ---- Patrón compuesto versus herencia

type Parent struct {
	SomeField int
}

type Son struct {
	P Parent
}

func GetParentField(p *Parent) int {
	return p.SomeField
}

func MainParent() {
	son := Son{}
	res := GetParentField(&son.P)
	fmt.Println(res)
}

// ------------------------------------------------------------
