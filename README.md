This is a simple Go program, which extracts Go routine ID (an uint64 number) from the Go runtime internals.

Such obtaining of Go routine ID is highly discouraged by the Go authors and maintainers. It's intentionally
hidden since in most cases building logic, based on knowing the current Go routine ID, is a bad idea. It
often leads to a tangled, hard to understand and debug, code.

However, in some cases it's still needed. For experienced developers, who know what they do, it may be worth
knowing the current Go routine ID. That's why this project is here.

I won't pretend I've invented something brand new here. All the code here is a result of my reading of the
Go 1.23 runtime sources and googling around. There are several other projects on GitHub, where obtaining of
the Go routine ID is implemented in a similar way. In all those projects, they provide some additional
functionality around it.

But I don't need all those fancy features built around knowing the current Go routine ID. All I need is a
fast and reliable way to obtain the ID as a number. I'll build all the logic around that on myself, if it's
needed.

So I dived deep inside the Go 1.23 runtime sources and learned how it works with Go routines. Then, I found
implementations of `runtime·setg` and implemented corresponding `·getg` functions in Go Assembler for amd64
and arm64. For now, those are all architectures I need. In the future I might add more, if needed.

And then, knowing the pointer of the `g` object, I found a way to automatically find an offset of the `goid`
field, using reflection and a few more functions from the Go runtime, which give me information about all
the types in the program, and a way to resolve them into the `reflect.Type` objects.

That's it. Feel free to use this code, but do it only if you really know what you're doing. You've been warned.
