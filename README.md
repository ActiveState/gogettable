A simple web app to provide `go get` [the manifest](http://golang.org/cmd/go/#Remote_import_path_syntax) needed to clone internal Go repos.

# Example

To allow a set of internal git repos to be namespaced at "go.stackato.to/..." so that the following may work:

    $ go get go.stacka.to/stackato-go

.. run this app as:

    $ gogettable -import go.stacka.to -repo ssh://gitolite@gitolite.activestate.com/%s.git

`go get` will then retrieve the manifest as:

    $ curl http://go.stacka.to/foo/bar
    <html><head><meta name="go-import" content="go.stacka.to/foo git ssh://gitolite@gitolite.activestate.com/foo.git"></head></html>

If the repo is public, you could also submit it to services like [godoc.org](http://godoc.org/)

# Deploying

After changing the `Procfile` to suit your needs, you can deploy this app directly to Stackato or Heroku.
