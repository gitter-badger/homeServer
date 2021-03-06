package day

import (
	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/sqlite"
	"testing"
	"time"
)

func TestBasic(t *testing.T) {
	db, err := gorm.Open("sqlite3", "/tmp/openboiler_test.sqlite")
	if err != nil {
		t.Fatal(err)
	}
	defer db.Close()

	db.CreateTable(&Day{})

	p := Day{Day: time.Now()}
	p.SetSince(12, 10)
	err = p.Save(db)
	if err != nil {
		t.Fatal(err)
	}
	if p.ID == 0 {
		t.Fatal("id null")
	}

	p.SetSince(13, 10)
	err = p.Update(db)
	if err != nil {
		t.Fatal(err)
	}

	p2, err := FindById(db, p.ID)
	if err != nil {
		t.Fatal(err)
	}
	if p2 == nil {
		t.Fatal("FindById Failed")
	}
	if p2.SinceMin != p.SinceMin {
		t.Fatal("Updated Failed")
	}

	if p.Delete(db) != nil {
		t.Fatalf("Delete Failed")
	}
}
