package weasel

import (
	"fmt"
	"net/http"
	nhpprof "net/http/pprof"
	"runtime/pprof"
	"strings"
	"text/template"
)

func handler() http.Handler {
	info := struct {
		Profiles []*pprof.Profile
		Token    string
	}{
		Profiles: pprof.Profiles(),
	}

	h := func(w http.ResponseWriter, r *http.Request) {
		// get the last path, allowing this to be mounted under any prefix
		split := strings.Split(r.URL.Path, "/")
		name := split[len(split)-1]

		switch name {
		case "":
			// Index page.
			if err := indexTmpl.Execute(w, info); err != nil {
				fmt.Fprintf(w, "something went wrong - %s", err)
				return
			}
		case "cmdline":
			nhpprof.Cmdline(w, r)
		case "profile":
			nhpprof.Profile(w, r)
		case "trace":
			nhpprof.Trace(w, r)
		case "symbol":
			nhpprof.Symbol(w, r)
		default:
			// Provides access to all profiles under runtime/pprof
			nhpprof.Handler(name).ServeHTTP(w, r)
		}
	}
	return http.HandlerFunc(h)
}

// Handler returns an http.Handler that provides access to the various
// profiler and debug tools in the /net/http/pprof and /runtime/pprof
// packages.
func Handler() http.Handler {
	return handler()
}

var indexTmpl = template.Must(template.New("index").Parse(`
<html>
  <head>
    <title>Weasel Debug Information</title>
  </head>
  <br>
  <body>
    profiles:<br>
    <table>
    {{range .Profiles}}
      <tr><td align=right>{{.Count}}<td><a href="{{.Name}}?debug=1">{{.Name}}</a>
    {{end}}
    <tr><td align=right><td><a href="profile">CPU</a>
    <tr><td align=right><td><a href="trace?seconds=5">5-second trace</a>
    <tr><td align=right><td><a href="trace?seconds=30">30-second trace</a>
    </table>
    <br>
    debug information:<br>
    <table>
      <tr><td align=right><td><a href="cmdline">cmdline</a>
      <tr><td align=right><td><a href="symbol">symbol</a>
    <tr><td align=right><td><a href="goroutine?debug=2">full goroutine stack dump</a><br>
    <table>
  </body>
</html>`))
