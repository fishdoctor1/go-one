package main

import "fmt"

type Employee struct {
	Id     string
	Office string
	Person
}

func (e *Employee) Detail() string {
	return "ID:" + e.Id + ". Office :" + e.Office + ". Fullname" + e.FullName()
}

type Person struct {
	Name    string
	Surname string
}

func (p *Person) FullName() string {
	return p.Name + " " + p.Surname
}

type Programmer struct {
	Employee
	Language []string
}

func (p *Programmer) Detail() string {
	return "Programmer:" + p.Detail()
}

func (e *Employee) IsSameOffice(other *Employee) bool {
	return e.Office == other.Office
}

type Tester struct {
	Employee
	Tools []string
}

func main() {

	david := Person{
		Name:    "david",
		Surname: "surname",
	}
	empDavid := Employee{
		Id:     "456",
		Office: "thai",
		Person: david,
	}

	progDavid := Programmer{
		Employee: empDavid,
		Language: []string{"GO"},
	}
	oliver := Person{
		Name:    "oliver",
		Surname: "surname",
	}
	empOliver := Employee{
		Id:     "123",
		Office: "thai",
		Person: oliver,
	}
	testerOliver := Tester{
		Employee: empOliver,
		Tools:    []string{"robot"},
	}

	// fmt.Printf("%+v", tester)
	fmt.Println(empDavid.FullName())
	fmt.Println(progDavid.FullName())
	fmt.Println(empDavid.IsSameOffice(&empOliver))
	fmt.Println(progDavid.IsSameOffice(&(testerOliver.Employee)))
	// fmt.Println(testerOliver)
	fmt.Println((*Employee).IsSameOffice(&empDavid, &empOliver))
	// or
	sameOff := (*Employee).IsSameOffice
	fmt.Println(sameOff(&empDavid, &empOliver))
}
