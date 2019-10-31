package main

import (
	"bytes"
	"fmt"
	"reflect"
	"runtime"
	"sync"
	"time"
)

var wg = sync.WaitGroup{}
var counter = 0
var m = sync.RWMutex{}

type Animal struct {
	name    string `required max:"100"`
	species string
}

type Bird struct {
	Animal
	speedKPH float32
	canFly   bool
}

var i = 14

var (
	// block variables
	firstName string = "Mayur"
	lastName  string = "Patel"
)

// add examples for iota

func main() {
	basics(i)
	bitShifting()
	arrayAndSlices()
	structs()
	maps()
	switchCases()
	looping()
	loopingOverCollections()
	pointers()
	functions()
	interfaces()
	goRoutines()
	channelBasics()
	loggerExample()
}

func basics(i int) {
	fmt.Println("\n#1: Learning about the basics of Go:-")
	fmt.Println(i)
	i = 24
	var j float32 = 30 // can control type
	k := 27.
	fmt.Printf("Hello, Go !! \n")
	fmt.Printf("%v, %T \n", i, i)
	fmt.Printf("%v, %T \n", j, j)
	fmt.Printf("%v, %T \n", k, k)
	fmt.Println(firstName, lastName)
}

func bitShifting() {
	fmt.Println("\n#2: Implementing Bit Shifting:-")
	l := 8
	fmt.Println(l << 2)
	fmt.Println(l >> 1)
}

func arrayAndSlices() {
	fmt.Println("\n#3: Learning about Array and Slices:-")
	ar1 := [...]int{1, 2, 3}
	fmt.Println(ar1)

	var myMixedArray [3][3]int
	myMixedArray[0] = [3]int{1, 2, 3}
	fmt.Println(myMixedArray)

	slice1 := []string{"Mayur", "Patel"}
	slice2 := slice1[:]
	slice2[1] = "Kadiwar"
	fmt.Println(slice1)

	myArray := [...]int{1, 2, 3, 4, 5, 6, 7, 8, 9, 10}
	fmt.Println(myArray)
	myArrayAsSlice := myArray[:]
	fmt.Println(myArrayAsSlice)
	myArrayAsSlice1 := myArray[3:]
	myArrayAsSlice2 := myArray[:6]
	myArrayAsSlice3 := myArray[3:6]
	fmt.Println(myArrayAsSlice1)
	fmt.Println(myArrayAsSlice2)
	fmt.Println(myArrayAsSlice3)

	myArrayAsSlice1[1] = 100
	fmt.Printf("%v %T\n", myArrayAsSlice, myArrayAsSlice)
	fmt.Printf("%v %T\n", myArray, myArray)
}

func structs() {
	fmt.Println("\n#4: Learning about Structs:-")
	peacok := Bird{
		Animal:   Animal{name: "Peacok", species: "Bird"},
		speedKPH: 20.5,
		canFly:   true,
	}
	fmt.Println(peacok.name)

	t := reflect.TypeOf(Animal{})
	field, ok := t.FieldByName("name")
	fmt.Println(field.Tag)
	fmt.Println(ok)
}

func maps() {
	fmt.Println("\n#5: Learning about Maps:-")
	myMap := map[string]int{
		"mayur":  20,
		"nisha":  25,
		"dhaval": 50,
	}
	fmt.Println(myMap)
	myMap["nilam"] = 40
	fmt.Println(myMap)
	delete(myMap, "nilam")
	fmt.Println(myMap)

	nameToLookUp := "nisha1"
	nameValue, ok := myMap[nameToLookUp]
	if ok {
		fmt.Println(nameValue)
	} else {
		fmt.Println("No Key Found")
	}
}

func switchCases() {
	fmt.Println("\n#6: Learning about Switch Cases:-")
	i := 10
	switch i {
	case 1, 5, 10:
		fmt.Println("first case")
	case 2, 6, 8:
		fmt.Println("Second case")
	default:
		fmt.Println("Default case")
	}
	var j interface{} = [3]int{1, 2, 3}
	switch j.(type) {
	case int:
		fmt.Println("j is int")
	case string:
		fmt.Println("j is string")
	case [3]int:
		fmt.Println("j is [3]int")
	case []int:
		fmt.Println("j is slice")
	default:
		fmt.Println("j is another typed")
	}
}

func looping() {
	fmt.Println("\n#7: Learning about Looping:-")
	for i := 0; i < 1; i++ {
		fmt.Println(i)
	}

	j := 0
	for ; j < 6; j++ {
		fmt.Println(j)
	}

	fmt.Println("Infinite Loop")
	k := 0
	for {
		fmt.Println(k)
		k++
		if k > 3 {
			break
		}
	}
}

func loopingOverCollections() {
	fmt.Println("\n#8: Learning about Looping over Collections:-")
	k := [3]int{1, 2, 3}
	for k, v := range k {
		fmt.Println(k, v)
	}

	fmt.Println("Only keys")
	j := []int{1, 2, 3, 4}
	for k := range j {
		fmt.Println(k)
	}

	fmt.Println("Only Values")
	m := []int{1, 2, 3, 4}
	for _, v := range m {
		fmt.Println(v)
	}
}

func pointers() {
	fmt.Println("\n#9: Learning about Pointers:-")
	var mps *myPoniterStruct
	mps = new(myPoniterStruct)
	mps.name = "Shawn"    // (*mps).name = "Shawn"
	fmt.Println(mps.name) // fmt.Println((*mps).name)

	fmt.Println("Other syntaxt for Struct:")
	mps2 := &myPoniterStruct{name: "Patrict"}
	fmt.Println(mps2.name)

	k := 10
	l := &k
	fmt.Println("Dereferencing by *")
	fmt.Println(k, l, *l)
	*l = 20
	fmt.Println("Changing value from pointer")
	fmt.Println(k, l, *l)
}

type myPoniterStruct struct {
	name string
}

func functions() {
	fmt.Println("\n#10: Learning about Functions:-")
	s := sum("I am doing sum for: ", 1, 2, 3, 4, 5)
	fmt.Println("The result is: ", s)

	fmt.Println("Annonyms function")
	// i is coming from package level variable
	func(i int) {
		fmt.Println("Calling Annonyms Function", i)
	}(i)

	fmt.Println("Divide function")
	result, err := divide(5.0, 0.0)
	if err != nil {
		fmt.Println(err)
	}
	fmt.Println("Division is: ", result)

	fmt.Println("Methods:-")
	msp := myPoniterStruct{name: "Taylor"}
	msp.fullName()
	fmt.Println("After using pointers: ", msp.name)
}

func sum(msg string, values ...int) int {
	fmt.Println(msg, values)
	result := 0
	for _, v := range values {
		result += v
	}
	return result
}

func divide(a, b float32) (float32, error) {
	if b == 0.0 {
		return 0.0, fmt.Errorf("Can't divide by Zero")
	}
	return a / b, nil
}

func (h *myPoniterStruct) fullName() {
	fmt.Println(h.name, "Swift")
	h.name = "Lady"
}

func interfaces() {
	fmt.Println("\n#11: Learning about Interfaces:-")
	var w Writer = ConsoleWriter{} // w := ConsoleWriter{}
	w.Write([]byte("Hello World"))

	fmt.Println("Type Casting and Compositions:")
	var wc WriterCloser = NewBufferedWriterCloser()
	wc.Write([]byte("Hello Friends, I am learning GoLang!!!"))
	wc.Close()

	r, ok := wc.(*BufferedWriterCloser)
	if ok {
		fmt.Println(r)
	} else {
		fmt.Println("Conversion Failed!!")
	}

	fmt.Println("Using Int primitive type:")
	myInt := IntCounter(0)
	int := &myInt
	for n := 0; n < 5; n++ {
		fmt.Println(int.Increment())
	}
}

// "This is Writer interface"
type Writer interface {
	Write([]byte) (int, error)
}

type Closer interface {
	Close() error
}

type WriterCloser interface {
	// composition as we can do same for structs
	Writer
	Closer
}

type BufferedWriterCloser struct {
	buffer *bytes.Buffer
}

func (bwc *BufferedWriterCloser) Write(data []byte) (int, error) {
	n, err := bwc.buffer.Write(data)
	if err != nil {
		return 0, err
	}

	v := make([]byte, 8)
	for bwc.buffer.Len() > 8 {
		_, err := bwc.buffer.Read(v)
		if err != nil {
			return 0, err
		}
		_, err = fmt.Println(string(v))
		if err != nil {
			return 0, err
		}
	}
	return n, nil
}

func (bwc *BufferedWriterCloser) Close() error {
	for bwc.buffer.Len() > 0 {
		data := bwc.buffer.Next(8)
		_, err := fmt.Println(string(data))
		if err != nil {
			return err
		}
	}
	return nil
}

func NewBufferedWriterCloser() *BufferedWriterCloser {
	return &BufferedWriterCloser{
		buffer: bytes.NewBuffer([]byte{}),
	}
}

type ConsoleWriter struct{}

func (cw ConsoleWriter) Write(data []byte) (int, error) {
	n, err := fmt.Println(string(data))
	return n, err
}

type Incrementer interface {
	Increment() int
}

type IntCounter int

func (ic *IntCounter) Increment() int {
	*ic++
	return int(*ic)
}

func goRoutines() {
	fmt.Println("\n#12: Learning about goroutines:-")
	fmt.Printf("No of Threads: %v\n", runtime.GOMAXPROCS(-1))

	for k := 0; k < 5; k++ {
		wg.Add(2)
		m.RLock()
		go sayHello()
		m.Lock()
		go increament()
	}
}

func sayHello() {
	fmt.Printf("Hello #%v\n", counter)
	m.RUnlock()
	wg.Done()
}

func increament() {
	counter++
	m.Unlock()
	wg.Done()
}

func channelBasics() {
	fmt.Println("\n#13: Learning about Channel Basics:-")
	var ch = make(chan int)
	wg.Add(2)
	go func() {
		k := <-ch
		ch <- 11
		fmt.Println(k)
		wg.Done()
	}()
	go func() {
		ch <- 10
		l := <-ch
		fmt.Println(l)
		wg.Done()
	}()
	wg.Wait()
}

const (
	logInfo    = "INFO"
	logWarning = "Warning"
	logError   = "Error"
)

type logEntry struct {
	time     time.Time
	severity string
	message  string
}

var logCh = make(chan logEntry, 50)
var doneCh = make(chan struct{})

func loggerExample() {
	fmt.Println("\n#14: Implementing Logger Example using channels:-")
	go logger()
	logCh <- logEntry{time.Now(), logInfo, "App is starting"}
	logCh <- logEntry{time.Now(), logInfo, "App is processing1"}
	logCh <- logEntry{time.Now(), logInfo, "App is processing2"}
	logCh <- logEntry{time.Now(), logInfo, "App is processing3"}
	logCh <- logEntry{time.Now(), logInfo, "App is processing4"}
	logCh <- logEntry{time.Now(), logInfo, "App is terminating"}
	time.Sleep(500 * time.Microsecond)
}

func logger() {
	for {
		select {
		case entry := <-logCh:
			fmt.Printf("%v - [%v]%v\n", entry.time.Format("2006-01-02T15:04:05"), entry.severity, entry.message)
		case <-doneCh:
			break
		}
	}
}
