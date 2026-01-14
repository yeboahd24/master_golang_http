package main

import (
	"fmt"
	"testing"
)

// ISP --> Interface Segregation Principle
// Code shouldn't accept anything it doesn't use.

// None ISP
type FileStorage struct{}

func (FileStorage) Save(data []byte) error {
	fmt.Println("Saving data to disk...")
	return nil
}

func (FileStorage) Load(id string) ([]byte, error) {
	fmt.Println("Loading data from disk...")
	return []byte("data"), nil
}

// FileStorage has two methods:
// Save and Load. Now suppose you write a function that only needs to save data:

// func Backup(fs FileStorage, data []byte) error {
// 	return fs.Save(data)
// }

// This works, but there are a few problems hiding here.
// Backup takes a FileStorage directly, so it only works with that type
// If you later want to back up to memory, a network location,
// or an encrypted store, you’ll need to rewrite the function. Because it depends on a concrete type,
// your tests have to use FileStorage

// Instead of depending on a specific type, we can depend on an abstraction.
// In Go, you can achieve that through an interface.

type Storage interface {
	Save(data []byte) error
	Load(id string) ([]byte, error)
}

// Now Backup can take a Storage instead:
// This fix coupling issue
// func Backup(store Storage, data []byte) error {
// 	return store.Save(data)
// }
//
// This also makes testing easier and free from side effect
// type FakeStorage struct{}
//
// func (FakeStorage) Save(data []byte) error         { return nil }
// func (FakeStorage) Load(id string) ([]byte, error) { return nil, nil }
//
// func TestBackup(t *testing.T) {
// 	fake := FakeStorage{}
// 	err := Backup(fake, []byte("test-data"))
// 	if err != nil {
// 		t.Fatal(err)
// 	}
// }

// However, there’s still one issue:
// Backup only calls Save, yet the Storage interface includes both Save and Load.
// If Storage later gains more methods, every fake must grow too, even if those methods aren’t used.
// That’s exactly what the ISP warns against.

// To fix that too we need to narrow down
type Saver interface {
	Save(data []byte) error
}

func Backup(s Saver, data []byte) error {
	return s.Save(data)
}

// Now the intent is clear. Backup only depends on Save.
// A test double can just implement that one method:

type FakeSaver struct{}

func (FakeSaver) Save(data []byte) error { return nil }

func TestBackup(t *testing.T) {
	fake := FakeSaver{}
	err := Backup(fake, []byte("test-data"))
	if err != nil {
		t.Fatal(err)
	}
}
