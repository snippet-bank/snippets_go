package main

import (
	"fmt"
	"strconv"
	"unsafe"
)

// TeamMember is a sample struct demonstrate Go's padding
// to align data along word boundaries for optimal fetching.
type TeamMember struct {
	isActive   bool // 1 byte
	isFulltime bool // 1 byte
	age        int  // 8 bytes
}

func main() {
	member1 := &TeamMember{
		isActive:   true,
		isFulltime: false,
		age:        28,
	}

	member2 := &TeamMember{
		isActive:   true,
		isFulltime: true,
		age:        32,
	}

	alignmentBoundary := unsafe.Alignof(member1)

	fmt.Printf("Alignment Boundary: %d\n", alignmentBoundary)
	fmt.Println("---------------------------")
	fmt.Println("Teammate struct #1")
	fmt.Printf("isActive = Size: %d Offset: %d Addr: %v\n",
		unsafe.Sizeof(member1.isActive), unsafe.Offsetof(member1.isActive), &member1.isActive)
	fmt.Printf("isFulltime = Size: %d Offset: %d Addr: %v\n",
		unsafe.Sizeof(member1.isFulltime), unsafe.Offsetof(member1.isFulltime), &member1.isFulltime)
	fmt.Printf("age = Size: %d Offset: %d Addr: %v\n",
		unsafe.Sizeof(member1.age), unsafe.Offsetof(member1.age), &member1.age)

	fmt.Println("---------------------------")
	fmt.Println("Teammate struct #2")
	fmt.Printf("isActive = Size: %d Offset: %d Addr: %v\n",
		unsafe.Sizeof(member2.isActive), unsafe.Offsetof(member2.isActive), &member2.isActive)
	fmt.Printf("isFulltime = Size: %d Offset: %d Addr: %v\n",
		unsafe.Sizeof(member2.isFulltime), unsafe.Offsetof(member2.isFulltime), &member2.isFulltime)
	fmt.Printf("age = Size: %d Offset: %d Addr: %v\n",
		unsafe.Sizeof(member2.age), unsafe.Offsetof(member2.age), &member2.age)

	fmt.Println("")
	fmt.Println("Let's focus on the memory addresses... ")
	fmt.Println("We'll convert them to decimal and just look at the last two digits, ")
	fmt.Println("so we can clearly see the byte-alignment happening.")
	fmt.Println("---------------------------")
	fmt.Println(addrLastDigits(&member1.isActive))
	fmt.Println(addrLastDigits(&member1.isFulltime))
	fmt.Println(addrLastDigits(&member1.age))
	fmt.Println("---------------------------")
	fmt.Println(addrLastDigits(&member2.isActive))
	fmt.Println(addrLastDigits(&member2.isFulltime))
	fmt.Println(addrLastDigits(&member2.age))
	fmt.Println("---------------------------")
	fmt.Println("")

	fmt.Println("See how the 8-byte int field lines up on 8-byte word boundaries?")
	fmt.Println("The 2 boolean fields are next to each other, 1 byte apart, ")
	fmt.Println("but the int field jumps ahead 6 bytes in order to start on an 8-byte boundary.")

}

func addrLastDigits(ptr interface{}) string {
	hex := fmt.Sprintf("%v", ptr)
	i, _ := strconv.ParseInt(hex, 0, 64)
	s := strconv.Itoa(int(i))
	return "..." + s[len(s)-2:]

}
