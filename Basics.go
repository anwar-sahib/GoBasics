package main

import (
	"bufio"
	"errors"
	"fmt"
	"math"
	"math/cmplx"
	"os"
	"strconv"
	"strings"
	"time"
)

func main() {

	fmt.Println("1. Variable Declaration")
	fmt.Println("2. Array and Slice")
	fmt.Println("3. Map")
	fmt.Println("4. For Loop")
	fmt.Println("5. Function Calling")
	fmt.Println("6. Struct")
	fmt.Println("7. Pointer")
	fmt.Println("8. Defer")
	fmt.Println("9. Function As Values")
	fmt.Println("10. Interfaces")
	fmt.Println("11. Go Routines and channels")
	fmt.Println("12. Panic and Recover")

	reader := bufio.NewReader(os.Stdin)
	fmt.Print("Enter option number: ")
	text, _ := reader.ReadString('\n')
	i, err := strconv.Atoi(strings.Replace(text, "\n", "", -1))
	if err != nil {
		fmt.Println("Error in coverion", err)
	}

	switch i {
	case 1:
		variableDec()
	case 2:
		arrayAndSlice()
	case 3:
		mapUsage()
	case 4:
		forLoop()
	case 5:
		functionCalling()
	case 6:
		structUsage()
	case 7:
		pointerUsage()
	case 8:
		deferUsage()
	case 9:
		functionAsValuesAndClosures()
	case 10:
		interfaceUsage()
	case 11:
		routinesAndChannelUsage()
	case 12:
		panicAndRecoverUsage()
	default:
		fmt.Println("Unknown Option")
	}

}

func variableDec() {
	//Un initialzed variables
	var i int
	var flag bool
	fmt.Println("Uninitialized variables:", i, flag)

	//Variable declaration
	var x int = 5
	y := 7 //:= short assignment statement for implicit type declaration
	sum := x + y
	fmt.Println("Sum of two interger:", sum)

	//multiple assignments
	var c, python, java = true, false, "no!" //c, python, java := true, false, "no!"
	fmt.Println("multiple assignments:", i, c, python, java)

	//Basic types in go bool, string, int (int8  int16  int32  int64), uint (uint8 uint16 uint32 uint64 uintptr),
	// byte (alias for uint8), rune (alias for int32, represents a Unicode code point), float32, float64, complex64, complex128
	var MaxInt uint64 = 1<<64 - 1
	var z complex128 = cmplx.Sqrt(-5 + 12i)

	fmt.Printf("Type: %T     Value: %v\n", MaxInt, MaxInt)
	fmt.Printf("Type: %T     Value: %v\n", z, z)

	const Pi = 3.14              //Constants are declared like variables, but with the const
	fmt.Println("Constant:", Pi) //Constants cannot be declared using the := syntax
}

func arrayAndSlice() {
	//Array
	var array [5]int
	array[2] = 4
	fmt.Println("Initialising single element of array:", array)

	a := [5]int{1, 2, 3, 4, 5}
	fmt.Println("Printing entire array:", a)

	//Slice
	s := []int{1, 2, 3, 4, 5} //Slice is an array which does not have size defined
	s = append(s, 6, 7)
	fmt.Println("Appending elements at end of slice:", s)

	//Creating a slice using make function
	m := make([]int, 5)
	printSlice(m)

	//Array is zero based index. Below slice includes the first element, but excludes the last one
	var slice1 []int = s[1:3]
	fmt.Println("Slice1 of array:", slice1)

	var slice2 []int = s[2:4]
	fmt.Println("Slice2 of array:", slice2)
	slice1[1] = 999
	fmt.Println("Original array:", s)      // Changing the elements of a slice modifies the corresponding elements of its underlying array/slice
	fmt.Println("Slices:", slice1, slice2) // Other slices that share the same underlying array will see those changes.

	// A slice has both a length and a capacity. The length of a slice is the number of elements it contains.
	// The capacity of a slice is the number of elements in the underlying array, counting from the first element in the slice.
	printSlice(s) //No idea why capacity is coming 10. It should be 7.
	// Answer - If the backing array of s is too small to fit all the given values a bigger array will be allocated. The returned slice will point to the newly allocated array.

	// Slice the slice to give it zero length.
	s = s[:0]
	printSlice(s)

	// Extend its length.
	s = s[:4]
	printSlice(s)

	// Drop its first two values.
	s = s[2:]
	printSlice(s)
	//If we extend the slice beyong it's capacity  we get "panic: runtime error: slice bounds out of range [:5] with capacity 4"
}

func mapUsage() {
	//Map : - syntax "map[ key ] value"
	vertices := make(map[string]int)
	vertices["a"] = 1
	vertices["b"] = 2
	vertices["c"] = 3
	fmt.Println(vertices)

	fmt.Println("Value of key b:", vertices["b"])

	//Looping over a map using range
	for k, v := range vertices {
		fmt.Println("key:", k, " value:", v)
	}

	delete(vertices, "b")
	fmt.Println("After deleting an entry:", vertices)

	value, ok := vertices["b"]                       //If key is in vertices, ok is true. If not, ok is false.
	fmt.Println("The value:", value, "Present?", ok) //If key is not in the map, then value is the zero for the map's element type.
}

func ifUsage() {
	//Simple if else
	i := 7
	if i <= 5 {
		fmt.Println("Less than 5")
	} else {
		fmt.Println("Greater than 5")
	}

	//Simple if - else if statement
	if i <= 5 {
		fmt.Println("Average")
	} else if i >= 7 {
		fmt.Println("Good")
	} else {
		fmt.Println("Excellent")
	}

	var x, n, limit float64 = 3, 2, 10
	if v := math.Pow(x, n); v < limit { //Variables declared by the statement are only in scope until the end of the if.
		fmt.Println("Calculated Value:", v)
	}
	fmt.Println("Limit :", limit)

}

func forLoop() {
	//for loop
	for i := 0; i < 5; i++ {
		fmt.Println(i)
	}

	//While loop equivalent
	sum := 1
	for sum < 10 {
		sum += sum
	}
	fmt.Println(sum)

	//Looping an array using range
	arr := []string{"a", "b", "c"}
	for index, value := range arr {
		fmt.Println("index:", index, " value:", value)
	}

	//Looping a map using range
	m := make(map[string]string)
	m["a"] = "alpha"
	m["b"] = "beta"
	for key, value := range m {
		fmt.Println("key:", key, " value:", value)
	}
}

func functionCalling() {
	//Calling functions
	result := sumInt(4, 5)
	fmt.Println("Result of sum func :", result)

	//Returning multiple values
	result64, err := sqrt(-16)
	if err != nil {
		fmt.Println(err)
	} else {
		fmt.Println(result64)
	}

}

func structUsage() {
	//Stuct - equivalent of Java class but without methods
	p := person{name: "John Doe", age: 34}
	fmt.Println("Printing Struct:", p)
	fmt.Println("Accessing name:", p.name)

	structPtr := &p
	//To access the field name of a struct when we have the struct pointer structPtr we could write (*structPtr).name.
	// However, that notation is cumbersome, so the language permits us instead to write just structPtr.name, without the explicit dereference.
	fmt.Println("Accessing name via pointer:", structPtr.name)

	//struct along with methods can form an equivalent to class, however this is less restrictive way of adding behaviour to struct (data representation)
	fmt.Println("First name via method:", p.getFirstName())

	fmt.Println("First name via function:", getfirstname(p))

	//You can declare a method on non-struct types, too. But you can only declare a method with a receiver whose type is defined
	//in the same package as the method. You cannot declare a method with a receiver whose type is defined in another package

	//Methods in Go are more general than in C++ or Java: they can be defined for any sort of data, even built-in types such as plain, “unboxed” integers.
	// They are not restricted to structs (classes).
	fmt.Println("Age of person:", p.age)
	p.incrementAge()
	fmt.Println("Age of person after increment:", p.age)

	var nullValue *person
	nullValue.incrementAge() //In other langauges like JAVA this give NPE, but this can be handled gracefully in go

	//NOTE : In general, all methods on a given type should have either value or pointer receivers, but not a mixture of both. Why??
}

// A method is a function with a special receiver argument.
// The receiver appears in its own argument list between the func keyword and the method name.
func (p person) getFirstName() string { //the printFirstName method has a receiver of type person named p.
	firstName := strings.Split(p.name, " ")
	return firstName[0]
}

//getFirstName() method transformed as function
func getfirstname(p person) string {
	firstName := strings.Split(p.name, " ")
	return firstName[0]
}

//Methods with pointer receivers can modify the value to which the receiver points.
// Since methods often need to modify their receiver, pointer receivers are more common than value receivers.
func (p *person) incrementAge() { //If we remove the point and pass just person type the age, then p.age is not changed
	if p == nil {
		fmt.Println("Called function with <nil> recevier")
		return
	}
	p.age = p.age + 1
}

func pointerUsage() {
	//Pointers
	i := 7
	fmt.Println("Value:", i)
	fmt.Println("Address of value:", &i)
	intVar(&i) //Get the addresss
	fmt.Println("Value after change via pointers derefrencing:", i)
	//struct pointers are explored in structUsage
}

func functionAsValuesAndClosures() {
	hypot := func(x, y float64) float64 { //hypot function takes two floats and returns a float
		return math.Sqrt(x*x + y*y)
	}

	fmt.Println("Passing functions as values")
	//Function values may be used as function arguments and return values
	fmt.Println(compute(hypot)) //Both hypot and math.pow are type of functions that compute accepts
	fmt.Println(compute(math.Pow))

	fmt.Println("Using functions as closures")
	//A closure is a function value that references variables from outside its body.
	// The function may access and assign to the referenced variables; in this sense the function is "bound" to the variables.

	pos, neg := adder(), adder()
	for i := 0; i < 3; i++ {
		fmt.Println(
			pos(i), //Each closure is bound to its own sum variable.
			neg(-2*i),
		)
	}
}

func deferUsage() {
	//A defer statement defers the execution of a function until the surrounding function returns.
	fileUsage()

	//Deferred function calls are pushed onto a stack. When a function returns, its deferred calls are executed in last-in-first-out order.
	fmt.Println("counting")

	for i := 0; i < 3; i++ {
		defer fmt.Println(i)
	}
	fmt.Println("done")

	defer fmt.Println("world")
	fmt.Println("hello") //The deferred call's arguments are evaluated immediately, but the function call is not executed until the surrounding function returns.

}

func fileUsage() {
	//Can be used to ensure that files are closed properly
	src, err := os.Open("srcName")
	if err != nil {
		fmt.Println("Error ", err)
		return
	}
	defer src.Close() //Differing closing so that file is closed when function is over

	// Other uses of defer include releasing a mutex:
	// mu.Lock()
	// defer mu.Unlock()
}

type iAbs interface {
	Abs() float64
}

type myFloat float64

//A type implements an interface by implementing its methods. There is no explicit declaration of intent, no "implements" keyword.
//Here MyFloat implements IAbs interface
func (f myFloat) Abs() float64 {
	if f < 0 {
		return float64(-f)
	}
	return float64(f)
}

type vertex struct {
	X, Y float64
}

func (v *vertex) Abs() float64 {
	return math.Sqrt(v.X*v.X + v.Y*v.Y)
}

func interfaceUsage() {
	var a iAbs
	f := myFloat(4)
	v := vertex{3, 4}

	a = f // a MyFloat implements Abser
	fmt.Println("Using abs implementation of float:", a.Abs())
	a = &v // a *Vertex implements Abser
	fmt.Println("Using abs implementation of Vertex:", a.Abs())
	// In the following line, v is a Vertex (not *Vertex) and does NOT implement IAbs.
	//a = v

	//Interface with zero methods is called an empty interface
	var i interface{}
	fmt.Printf("(%v, %T)\n", i, i)
	//An empty interface may hold values of any type
	i = 42
	fmt.Printf("(%v, %T)\n", i, i)
	//Empty interfaces are used by code that handles values of unknown type
	i = "hello"
	fmt.Printf("(%v, %T)\n", i, i)

	//Type assertion for interfaces
	s := i.(string)
	fmt.Println(s)

	s, ok := i.(string)
	fmt.Println(s, ok)

	fl, ok := i.(float64) //If we do not user ok and type check is wrong then we get "panic: interface conversion: interface {} is string, not float64"
	fmt.Println(fl, ok)
}

func routinesAndChannelUsage() {
	s := []int{7, 2, 8, -9, 4, 0}

	c := make(chan int) //Like maps and slices, channels must be created before use
	go sum(s[:len(s)/2], c)
	go sum(s[len(s)/2:], c)
	//By default, sends and receives block until the other side is ready.
	//This allows goroutines to synchronize without explicit locks or condition variables.
	x, y := <-c, <-c // receive from c
	fmt.Println("channel 1 sum:", x, " channel 2 sum:", y, " total:", x+y)

	ch := make(chan int, 5)
	go fibonacci(cap(ch), ch)
	//Channels aren't like files; you don't usually need to close them. Closing is only necessary when the receiver
	// must be told there are no more values coming, such as to terminate a range loop
	for i := range ch {
		fmt.Println("Output read from channel:", i)
	}

	//A sender can close a channel to indicate that no more values will be sent. Receivers can test whether a
	//channel has been closed by assigning a second parameter to the receive expression
	_, ok := <-ch
	fmt.Println("Channel Open:", ok)

	quit := make(chan int)
	go func() { // Anynomous go function to enable to run it inside a go routine
		for i := 0; i < 10; i++ {
			fmt.Println("Reading from channel with select option:", <-c)
			if i == 5 {
				quit <- 0
			}
		}

	}()
	fibonacci2(c, quit)

	tick := time.Tick(100 * time.Millisecond)
	boom := time.After(500 * time.Millisecond)
	for {
		select {
		case <-tick:
			fmt.Println("tick.")
		case <-boom:
			fmt.Println("BOOM!")
			return
		default: //Use a default case to try a send or receive without blocking
			fmt.Println("    .") //The default case in a select is run if no other case is ready.
			time.Sleep(50 * time.Millisecond)
		}
	}
}

func sum(s []int, c chan int) {
	sum := 0
	for _, v := range s {
		sum += v
	}
	c <- sum // send sum to c
}

func fibonacci(n int, c chan int) {
	x, y := 0, 1
	for i := 0; i < n; i++ {
		c <- x
		x, y = y, x+y
	}
	close(c) //Only the sender should close a channel, never the receiver. Sending on a closed channel will cause a panic.
}

func fibonacci2(c, quit chan int) {
	x, y := 0, 1
	for {
		select { //The select statement lets a goroutine wait on multiple communication operations.
		case c <- x:
			x, y = y, x+y
		case <-quit: //A select blocks until one of its cases can run, then it executes that case. It chooses one at random if multiple are ready.
			fmt.Println("quit")
			return
		}
	}
}

func intVar(x *int) { //Addess of a type is stored in pointer of that type
	*x++ //Dereferencing a pointer.( Getting the value stored at that pointer address)
}

func sumInt(x int, y int) int {
	return x + y
}

func sqrt(x float64) (float64, error) {
	if x < 0 {
		return 0, errors.New("Square root is undefined for negative numbers")
	} else {
		return math.Sqrt(x), nil
	}

}

func printSlice(s []int) {
	fmt.Printf("Printing of slice, len=%d cap=%d %v\n", len(s), cap(s), s)
}

//fn is func type which takes two float64 and returns a float64.
//Compute function takes a func (of defined type) and returns a float64
func compute(fn func(float64, float64) float64) float64 {
	return fn(3, 4)
}

//adder is returning function (which is used as closure)
func adder() func(int) int {
	sum := 0 //Each closure is bound to its own sum variable.
	return func(x int) int {
		sum += x
		return sum
	}
}

type person struct {
	name string
	age  int
}

func panicAndRecoverUsage() {
	catchRecover()
	fmt.Println("Returned normally from panicAndRecoverUsage.")
}

//Recover is a built-in function that regains control of a panicking goroutine.
func catchRecover() {
	defer func() {
		//Recover is only useful inside deferred functions. During normal execution, a call to recover will return nil and have no other effect.
		if r := recover(); r != nil {
			fmt.Println("Recovered in catchRecover:", r) //If the current goroutine is panicking, a call to recover will capture the value given to panic and resume normal execution.
		}
	}()
	fmt.Println("Calling g.")
	//To the caller, F then behaves like a call to panic. The process continues up the stack until all functions in the current
	// goroutine have returned, at which point the program crashes.
	createPanic(0)
	fmt.Println("Returned normally from g.")
}

//When the function F calls panic, execution of F stops, any deferred functions in F are executed normally, and then F returns to its caller.
func createPanic(i int) {
	if i > 3 {
		fmt.Println("Panicking!")
		panic(fmt.Sprintf("%v", i))
	}
	defer fmt.Println("Defer in createPanic:", i)
	fmt.Println("Printing in createPanic:", i)
	createPanic(i + 1)
}
