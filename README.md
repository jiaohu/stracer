# stracer
simple tracer for step used time

# Example
```go
func Demo() {
  ca := stracer.New("Demo", true)
  a := 1
  ca.Step("init")
  a += 1
  ca.Step("add")
  ca.Done()
}
```
