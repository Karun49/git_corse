package main

import "fmt"

type Runner interface {
	Run() string
}

type Swimmer interface {
	Swim() string
}

type Flyer interface {
	Fly() string
}

type Ducker interface {
	Runner
	Swimmer
	Flyer
}

type Human struct {
	Name string
}

func (h Human) Run() string {
	return fmt.Sprintf("Человек %s бегает", h.Name)
}

func (h Human) writeCode() {
	fmt.Println("Человек пишет код")
}

type Duck struct {
	Name, Surname string
}

func (d Duck) Run() string {
	return "Утка бегает"
}

func (d Duck) Swim() string {
	return "Утка плавает"
}

func (d Duck) Fly() string {
	return "Утка летает"
}

func main() {
	//interfaceValues()
	typeAssertionAndPolymorphism()
}

func interfaceValues() {
	var runner Runner
	fmt.Printf("Type: %T Value: %#v\n", runner, runner)

	if runner == nil {
		fmt.Println("Runner is nil")
	}
	//runner = int64(1)
	//runner.Run()

	var unnamedRunner *Human
	fmt.Printf("Type: %T Value: %#v\n", unnamedRunner, unnamedRunner)

	runner = unnamedRunner
	fmt.Printf("Type: %T Value: %#v\n", runner, runner)
	if runner == nil {
		fmt.Println("Runner is nil")
	}

	namedRunner := &Human{Name: "Джек"}
	fmt.Printf("Type: %T Value: %#v\n", namedRunner, namedRunner)

	runner = namedRunner
	fmt.Printf("Type: %T Value: %#v\n", runner, runner)

	// empty interface{}

	var emptyInterface interface{} = unnamedRunner
	fmt.Printf("Type: %T Value: %#v\n", emptyInterface, emptyInterface)

	emptyInterface = runner
	fmt.Printf("Type: %T Value: %#v\n", emptyInterface, emptyInterface)

	emptyInterface = int64(1)
	fmt.Printf("Type: %T Value: %#v\n", emptyInterface, emptyInterface)

	emptyInterface = true
	fmt.Printf("Type: %T Value: %#v\n", emptyInterface, emptyInterface)
}

func typeAssertionAndPolymorphism() {
	var runner Runner
	fmt.Printf("Type: %T Value: %#v\n", runner, runner)

	john := &Human{"John"}
	runner = john
	polymorphism(john)
	typeAssertion(john)

	donald := &Duck{Name: "Donald", Surname: "Duck"}
	runner = donald
	polymorphism(donald)
	typeAssertion(donald)
}
func polymorphism(runner Runner) {
	fmt.Println(runner.Run())
}

func typeAssertion(runner Runner) {
	fmt.Printf("Type: %T Value: %#v\n", runner, runner)
	if human, ok := runner.(*Human); ok {
		fmt.Printf("Type: %T Value: %#v\n", human, human)
		human.writeCode()
	}
	if duck, ok := runner.(*Duck); ok {
		fmt.Printf("Type: %T Value: %#v\n", duck, duck)
		fmt.Println(duck.Fly())
	}

	switch v := runner.(type) {
	case *Human:
		fmt.Println(v.Run())
	case *Duck:
		fmt.Println(v.Swim())
	default:
		fmt.Printf("Type: %T Value: %#v\n", v, v)
	}
}

///////////////////////////////////////////////////////
/*
type Square struct {
	Side int
}

func main() {
	defenition()
}

func defenition() {
	//////МЕТОДЫ
	square := Square{Side: 8}
	psquare := &square

	square2 := Square{Side: 2}

	square.Perimater()
	square2.Perimater()
	psquare.Scale(2)

	fmt.Print(square.Side, "\n")

	psquare.Perimater()

	fmt.Print(square.Side, "\n")

	square.Scale(2)
	psquare.Perimater()

	fmt.Print(square.Side, "\n")
	square.Wrong_Scale(2)
	square.Perimater()
}

func (s Square) Perimater() {
	fmt.Printf("%T %#v \n", s, s)
	fmt.Printf("Периметр фигуры %d \n", s.Side*4)
}

func (s *Square) Scale(Multiplier int) {
	fmt.Printf("%T %#v \n", s, s)
	s.Side *= Multiplier
	fmt.Printf("%T %#v \n", s, s)
}
func (s Square) Wrong_Scale(multiplier int) {
	fmt.Printf("%T %#v \n", s, s)
	s.Side *= multiplier
	fmt.Printf("%T %#v \n", s, s)

}

/*
/////////Кастомные типы ,Структуры
type Our_strins string
type Our_int int

type Person struct {
	Name string
	Age  int
}

func main() {
	var Custom_Strings Our_strins
	fmt.Printf("%T %#v \n", Custom_Strings, Custom_Strings)
	Custom_Strings = "Hello dudes"
	fmt.Printf("%T %#v \n", Custom_Strings, Custom_Strings)

	Custom_Int := Our_int(5)
	fmt.Printf("%T %#v \n", Custom_Int, Custom_Int)

	var John Person
	fmt.Printf("%T %#v \n", John, John)
	//Fields accessing
	John.Name = "John"
	John.Age = 23
	fmt.Println(John)
	////create with named field
	Brad := Person{
		Name: "Brad",
		Age:  20,
	}
	fmt.Println(Brad)
	//create without field names
	Vladimir := Person{"Vladimir", 18}
	fmt.Println(Vladimir)

	//fiend accessing the pointer
	pVladimir := &Vladimir
	fmt.Println((*pVladimir).Age) //==//↓↓↓↓↓↓
	fmt.Println(pVladimir.Age)    //==//↑↑↑↑↑↑↑↑

	// create pointer in struct
	pIvan := &Person{"Ivan", 23}
	fmt.Println(pIvan)

	unnamed_Struct := struct {
		Name, LastName, BirthDate string
	}{
		Name:      "NoName",
		LastName:  "NoLastName",
		BirthDate: fmt.Sprintf("%s", time.Now()),
	}
	fmt.Println(unnamed_Struct)
}

/*{
	//////////////////////Pointers - Указатели ///////////////////
	// defalt value
	var intPointer *int
	fmt.Printf("%T %#v \n", intPointer, intPointer)

	var A float64 = 7
	fmt.Printf("%T %#v \n", A, A)

	var intPointerA = &A
	fmt.Printf("%T %#v %#v \n", intPointerA, intPointerA, *intPointerA)
	var newPointer = new(float64)
	fmt.Printf("%T %#v %#v \n", newPointer, newPointer, *newPointer)
	*newPointer = 4
	fmt.Printf("%T %#v %#v \n", newPointer, newPointer, *newPointer)
}
{	num := 3
	square(num)
	fmt.Println(num)
	square_Pointer(&num)
	fmt.Println(num)

	var wallet1 *int
	fmt.Println(hex_Wallet(wallet1))
	wallet2 := 0
	fmt.Println(hex_Wallet(&wallet2))
	wallet3 := 10
	fmt.Println(hex_Wallet(&wallet3))

}
func square(num int)             { num *= num }
func square_Pointer(num *int)    { *num *= *num }
func hex_Wallet(money *int) bool { return money != nil }

//////// swithc case base//////////////
/*
func main() {

	rand.Seed(time.Now().UnixNano())
	value := rand.Intn(max-min) + min

	// if example
	if value == 1{
		fmt.Println("One")
	}else if value == 2 || value == 3{
		fmt.Println("Two or three")
	}else if value == getFour() {
		fmt.Println("four")
	}else {
		fmt.Print("Defalt case is shown")
	}

	// base switch
	switch value {
	case 1:
		fmt.Println("1")
	case 2,3:
		fmt.Println("2 or 3")
	case getFour():
		fmt.Println("4")
	default:
		fmt.Println("5")
	}
// switch with local variable defenition
	switch num := rand.Intn(max-min)+ min; num{
	case 1:
		fmt.Println("1")
	case 2,3:
		fmt.Println("2 or 3")
	case getFour():
		fmt.Println("4")
		fallthrough
	case 10:
		fmt.Println("Strange things heppens")
	default:
		fmt.Println("Defalt")
	}

	switch{
	case value > 2:
		fmt.Println(("value > 2"))
	case value < 2:
		fmt.Println(("value < 2"))
	default:
		fmt.Println("Value equals 2")
	}
}

func getFour()int  {return 4}
*/

///////////////////////// Цыклы ///////////////////////////
// 	for j := 0; j <= 20; j++ {
// 		if j%2 == 1 {
// 			continue
// 		}
// 		//fmt.Println(j)
// 	}

// lebal1:
// 	for i := 1; i <= 20; i++ {
// 		for j := 1; j <= 10; j++ {
// 			//fmt.Println("I:", i, "J:", j)
// 			if i >= 10 {
// 				continue lebal1
// 			}
// 		}
// 	}

// 	for i := 0; i <= 20; i++ {
// 		if i >= 10 {
// 			break
// 		}
// 		//fmt.Println(i)
// 	}

// lebal2:
// 	for i := 1; i <= 20; i++ {
// 		for j := 1; j <= 10; j++ {
// 			fmt.Println("I:", i, "J:", j)
// 			if i >= 10 {
// 				break lebal2
// 			}
// 		}
// 	}

////////////IF ELSE ///////////////////////
// func main() {
// 	age := 15
// 	// if ... else
// 	if age > 17 {
// 		fmt.Println("Yoy are an adult")
// 	} else {
// 		fmt.Println("Yoy are too young")
// 	}

// 	// &&
// 	if age >= 6 && age <= 18 {
// 		fmt.Println("You are in school")
// 	}
// 	// ||
// 	if age >= 14 || age >= 20 || age >= 48 {
// 		fmt.Println("DSADSADSA")
// 	}

// }

//////////////////////Продвинутая работа с функциями/////////////////////
// func main() {
// 	SumCalc := func(x, y int)int{return x + y}
// 	first , second := 10 , 5
// 	fmt.Println(Calcul(first, second, SumCalc))
// 	dividerBy2 := createDivider(2)
// 	dividerBy10 := createDivider(10)

// 	fmt.Println(dividerBy2(100))
// 	fmt.Println(dividerBy10(100))

// }

// func Calcul(x, y int, action func(x, y int)int)int{
// 	return action(x, y)
// }
// func createDivider(divider int) func(x int) int {
// 	dividerFunc := func(x int) int{
// 		return x / divider
// 	}
// 	return dividerFunc
// }
