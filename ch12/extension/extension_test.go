package extension

import "fmt"
import "testing"

type Pet struct {

}

func (p *Pet) Speak() {
	fmt.Println("...")
}

func (p *Pet) SpeakTo(host string) {
	p.Speak()
	fmt.Println(" ", host) 
}

/*
type Dog struct {
	p *Pet 
}

func (d *Dog) Speak() {
	//d.p.Speak()
	fmt.Println("Wang!")
}

func (d *Dog) SpeakTo(host string) {
	//d.p.SpeakTo(host)
	d.Speak()
	fmt.Println(" ", host) 
}
*/

type Dog struct {
	Pet 
}

func (d *Dog) Speak() {
	fmt.Println("Wang!")
}

func TestDog(t *testing.T) {
	dog := new(Dog)
	dog.SpeakTo("Chao")
	//dog.Speak()
}