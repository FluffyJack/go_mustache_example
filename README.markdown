# Go Mustache Example

This is basically the first somewhat non-hello world go program I've written.

## Running it yourself

Create your go folder and set up your paths.

```
mkdir fluffyjacks_example
cd fluffyjacks_example
mkdir bin pkg src
export GOPATH=$HOME/fluffyjacks_example
export PATH=$GOPATH/bin:$PATH
```

This seems to be the standard way of starting a project in go... I don't really know to be honest, it works though.

Now you just need to pull in my code.

```
go get github.com/FluffyJack/go_mustache_example
```

What I reckon is awesome is that it pulls in the dependancies too, which is the mustache repo.

Now install.

```
go install github.com/FluffyJack/go_mustache_example
```

That compiled everything, and installed it in `$GOPATH/bin`. Seeming we added that to our path, we can simply run `go_mustache_example` now because it's in our path.

If you visit `http://localhost:8080/` now, you'll see this in your browser.

```
open /Users/jackwatson-hamblin/GoApps/mw/views/home.mustache: no such file or directory
```

Alright, so we need `views/home.mustache` in the working directory that we ran our go program from, remember to exit the server first.

```
mkdir views
echo "Hello {{name}}" > views/home.mustache
```

And run the server again with `go_mustache_example`.

Visit `http://localhost:8080/` and you will see `Hello`. Where is the name?

If you look in my code, you'll see I'm grabbing the `name` query parameter. So if you go to `http://localhost:8080/?name=Jack` instead, you'll see `Hello Jack`.

The interesting thing about Go's net/http library for me was that it has an interesting way of handling the URL handlers. It will match the most specific and cloest registered URL handler. Of course I only have `/` in this example, but this means we can actually go to any URL (try it: `http://localhost:8080/test/foo/bar?name=Jack`) and it will still not 404 and will show `Hello Jack`.

## Send me help!

I'm finding a lot of interest in Go all of a sudden, if you know of any really great learning resources, maybe like [https://gobyexample.com/](https://gobyexample.com/) but video content, that would be awesome, please let me know!
