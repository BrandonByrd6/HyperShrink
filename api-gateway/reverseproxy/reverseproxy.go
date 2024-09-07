package reverseproxy

import (
	"net/http"
	"net/http/httputil"
	"net/url"
	"strings"

	"github.com/gorilla/mux"
)

type ReverseProxy struct {
	proxy  *httputil.ReverseProxy
	Target []*Target
}

type Target struct {
	router   *mux.Router
	upstream *url.URL
}

func (r *ReverseProxy) AddTarget(upstream string, router *mux.Router) error {
	url, err := url.Parse(upstream)
	if err != nil {
		return err
	}

	if router == nil {
		router = mux.NewRouter()
		router.PathPrefix("/")
	}

	r.Target = append(r.Target, &Target{
		router:   router,
		upstream: url,
	})

	return nil
}

func (r *ReverseProxy) Start() error {
	r.proxy = &httputil.ReverseProxy{
		Director: r.Director(),
	}

	// Hard-coding port 80 for now
	srv := &http.Server{Addr: ":80", Handler: r.proxy}

	return srv.ListenAndServe()
}

func (r *ReverseProxy) Director() func(req *http.Request) {
	return func(req *http.Request) {

		for _, t := range r.Target {
			match := mux.RouteMatch{}
			if t.router.Match(req, &match) {
				targetQuery := t.upstream.RawQuery
				req.URL.Scheme = t.upstream.Scheme
				req.URL.Host = t.upstream.Host
				req.URL.Path, req.URL.RawPath = joinURLPath(t.upstream, req.URL)
				if targetQuery == "" || req.URL.RawQuery == "" {
					req.URL.RawQuery = targetQuery + req.URL.RawQuery
				} else {
					req.URL.RawQuery = targetQuery + "&" + req.URL.RawQuery
				}
				if _, ok := req.Header["User-Agent"]; !ok {
					// explicitly disable User-Agent so it's not set to default value
					req.Header.Set("User-Agent", "")
				}
			}
		}
	}
}

func singleJoiningSlash(a, b string) string {
	aslash := strings.HasSuffix(a, "/")
	bslash := strings.HasPrefix(b, "/")
	switch {
	case aslash && bslash:
		return a + b[1:]
	case !aslash && !bslash:
		return a + "/" + b
	}
	return a + b
}

func joinURLPath(a, b *url.URL) (path, rawpath string) {
	if a.RawPath == "" && b.RawPath == "" {
		return singleJoiningSlash(a.Path, b.Path), ""
	}
	// Same as singleJoiningSlash, but uses EscapedPath to determine
	// whether a slash should be added
	apath := a.EscapedPath()
	bpath := b.EscapedPath()

	aslash := strings.HasSuffix(apath, "/")
	bslash := strings.HasPrefix(bpath, "/")

	switch {
	case aslash && bslash:
		return a.Path + b.Path[1:], apath + bpath[1:]
	case !aslash && !bslash:
		return a.Path + "/" + b.Path, apath + "/" + bpath
	}
	return a.Path + b.Path, apath + bpath
}
