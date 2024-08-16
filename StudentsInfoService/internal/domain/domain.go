package domain

import "fmt"

var LastId = uint(0)

type Student struct {
	FirstName  string `json:"firstName"`
	SecondName string `json:"secondName"`
	Age        uint   `json:"age"`
	Faculty    string `json:"faculty"`
	Hostel     uint   `json:"hostel"`
	Room       uint   `json:"room"`

	Id uint `json:"id"`
}

func New(firstName, secondName string, age uint, faculty string, hostel, room uint) (*Student, error) {
	user := Student{}

	if err := user.SetFirstName(firstName); err != nil {
		return nil, err
	}
	if err := user.SetSecondName(secondName); err != nil {
		return nil, err
	}
	if err := user.SetAge(age); err != nil {
		return nil, err
	}
	if err := user.SetFaculty(faculty); err != nil {
		return nil, err
	}
	if err := user.SetHostel(hostel); err != nil {
		return nil, err
	}
	if err := user.SetRoom(room); err != nil {
		return nil, err
	}

	user.Id = LastId
	LastId++

	return &user, nil
}

func (user *Student) SetFirstName(firstName string) error {
	size := len(firstName)
	if size == 0 {
		return fmt.Errorf("bad first name <%v>", firstName)
	}

	user.FirstName = firstName

	return nil
}

func (user *Student) SetSecondName(secondName string) error {
	size := len(secondName)
	if size == 0 {
		return fmt.Errorf("bad second name <%v>", secondName)
	}

	user.SecondName = secondName

	return nil
}

func (user *Student) SetAge(age uint) error {
	if age <= 0 {
		return fmt.Errorf("bad age <%v>", age)
	}

	user.Age = age

	return nil
}

func (user *Student) SetFaculty(faculty string) error {
	size := len(faculty)
	if size == 0 {
		return fmt.Errorf("bad faculty <%v>", faculty)
	}

	user.Faculty = faculty

	return nil
}

func (user *Student) SetHostel(hostel uint) error {
	if hostel <= 0 {
		return fmt.Errorf("bad hostel <%v>", hostel)
	}

	user.Hostel = hostel

	return nil
}

func (user *Student) SetRoom(room uint) error {
	if room <= 0 {
		return fmt.Errorf("bad room <%v>", room)
	}

	user.Room = room

	return nil
}

//------------------------------

func (user *Student) GetFirstName() string {
	return user.FirstName
}

func (user *Student) GetSecondName() string {
	return user.SecondName
}

func (user *Student) GetAge() uint {
	return user.Age
}

func (user *Student) GetFaculty() string {
	return user.Faculty
}

func (user *Student) GetHostel() uint {
	return user.Hostel
}

func (user *Student) GetRoom() uint {
	return user.Room
}

func (user *Student) GetId() uint {
	return user.Id
}
