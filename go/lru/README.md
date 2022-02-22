# LRU Cache Implementation

A simple LRU cache implementation using a linked list and a map.  key is a string, value is a string

See tests for more detailed behavior

```go
l := NewLRU(3) // capacity is 3
l.Add("1", "a")
l.Add("2", "b")
l.Add("3", "c")
l.Add("4", "d") // evicted, now contains 2,3,4
l.Read("2") // returns "b", now next to evict is 3
l.Remove("4") //  you can remove an element as well
l.Dump() // this will dump what's in head and tail, as well as the whole linked list
```
