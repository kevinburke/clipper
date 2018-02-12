# Go html boilerplate

This is a starter pack for doing [web development with Go][post], with support
for some of the things you'll usually want to add to an HTML web server:

- Adding templates and rendering them
- Regex matching for routes
- Logging requests and responses
- Serving static content with caching/busting
- Watching/restarting the server after changes to CSS/templates
- Loading configuration from a config file
- Flash success and error messages

[Read more about the choices and the feature set found here][post].

Feel free to adapt the project as you see fit; that should be pretty easy to
do since no one component does too much on its own, and all of them operate on
standard library interfaces like `http.Handler`.

To get started, run `go get ./...` and then `make serve` to start a server on
port 7065. You may need to run `make generate_cert` to generate a self-signed
certificate for local use.

Templates go in the "templates" folder; you can see how they're loaded by
examining the `init` function in main.go. Run `make assets` to recompile them
into the binary.

Static files go in the "static" folder. Run `make assets` to recompile them into
the binary.

### Watching for changes

Run `make watch` to restart the server after you make changes to the assets
directory.

If you are on a Mac, be sure to **add this folder to the Spotlight privacy
list**, or file modify events will [fire a second time when Spotlight indexes
updates][fsnotify].

[fsnotify]: https://github.com/fsnotify/fsnotify/issues/15
[post]: https://kev.inburke.com/kevin/go-web-development/?github
