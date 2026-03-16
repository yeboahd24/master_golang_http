package main



// I — Interface Segregation Principle
// Don't force clients to depend on methods they don't use.
// In Go, keep interfaces small — ideally one or two methods. Clients define the interface they need, not the implementor.



// BAD — callers that only need Read() must still "know about" Write/Delete/List
type Storage interface {
    Read(key string) ([]byte, error)
    Write(key string, val []byte) error
    Delete(key string) error
    List(prefix string) ([]string, error)
}

// GOOD — small, composable interfaces
type Reader interface { Read(key string) ([]byte, error) }
type Writer interface { Write(key string, val []byte) error }
type ReadWriter interface {
    Reader
    Writer
}

// A cache service only depends on what it actually uses
func NewCache(store Reader) *Cache { ... }

type Cache struct {
    store Reader
}
func (c *Cache) Get(key string) ([]byte, error) {
    return c.store.Read(key)
}
// func (c *Cache) Set(key string, val []byte) error {
//     return c.store.Write(key, val)
// }
//
type Cache struct {
    store ReadWriter
}
func (c *Cache) Get(key string) ([]byte, error) {
    return c.store.Read(key)
}
func (c *Cache) Set(key string, val []byte) error {
    return c.store.Write(key, val)
}
