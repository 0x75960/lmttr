Lmttr
=====

limitter for concurrent processing

What to do
----------

* provides semaphore
* wait processes

Usage
------

```sh
go get -u github.com/0x75960/lmttr
```

```go
	import "github.com/0x75960/lmttr"

	// ...

	func main() {

		l, err := lmttr.NewLimitter(4) // 4 processes at the same time

		for _, s := range []string{"a", "b", "c", "d"} {

			l.Start()

			go func() {
				defer l.End()
				time.Sleep(10)
				fmt.Println(s)
			}()

			l.Wait()
		}

	}
```
