A simple web app to provide `go get` [the manifest](http://golang.org/cmd/go/#Remote_import_path_syntax) needed to clone internal Go repos.

# Example

To allow a set of internal git repos to be namespaced at "code.example.com/..." so that the following may work:

    $ go get code.example.com/mypkg

.. run this app as:

    $ gogettable -import code.example.com -repo ssh://git@git.example.com/%s.git
    
and serve it from the domain "code.example.com".

`go get` will then retrieve the manifest as:

    $ curl http://code.example.com/mypkg/subpkg
    <html><head>
    <meta name="go-import" content="code.example.com/mypkg git ssh://git@git.example.com.com/mypkg.git">
    </head></html>

If the repo is public, you could also submit it to services like [godoc.org](http://godoc.org/)

# Deploying

After changing the `Procfile` to suit your needs, you can deploy this app directly to Stackato or Heroku.
