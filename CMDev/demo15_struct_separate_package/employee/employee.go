package employee

import "fmt"

type Employee struct {
	// Em..ตั้งตัวใหญ่โลกภายนอกจะเห็น
	FirstName   string
	LastName    string
	TotalLeaves int
	LeavesTaken int
}

func (e Employee) LeavesRemaining() {
	fmt.Printf("%s %s has %d leaves remaining\n", e.FirstName, e.LastName, (e.TotalLeaves - e.LeavesTaken))
}
