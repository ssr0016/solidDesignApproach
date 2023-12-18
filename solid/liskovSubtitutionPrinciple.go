package solid

// The Liskov Substitution Principle states that objects of a derived class should be able to replace objects of the base class without affecting
// the correctness of the program. In Golang, this principle applies to interfaces and their implementations, ensuring that the code remains consistent and reliable.

type Bird interface {
	Fly() string
}

type pigeon struct{}

func (p *Pigeon) Fly() string {
	return "Pigeon is flying"
}

type Penguin struct{}

func (p *Penguin) Fly() string {
	return "Penguin is flying"
}

// In this case, both Pigeon and Penguin implement the Bird interface.
//  However, penguins cannot fly, so the Penguin implementation violates the Liskov Substitution Principle.
//  To fix this, we can refactor our code to separate the concerns:

type Birds interface {
	Makesound() string
}

type FlyingBird interface {
	Bird
	Fly() string
}

type Pigeon struct{}

func (p *Pigeon) Makesound() string {
	return "Coo"
}

func (p *Pigeon) Fly1() string {
	return "Pigeon is flying"
}

type Penguin1 struct{}

func (p *Penguin) Makesound() string {
	return "Squawk"
}
