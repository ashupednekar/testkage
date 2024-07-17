package server

import (
	"fmt"
	"log"
	"net/http"
	"net/url"

	"github.com/ashupednekar/testkage/internal/conf"
)

func CallService(targetUrl *url.URL, r *http.Request) (*http.Response, error) {
  log.Printf("Forwarding request to %s\n", targetUrl)
	req, err := http.NewRequest(r.Method, targetUrl.ResolveReference(r.URL).String(), r.Body)
	if err != nil {
		return &http.Response{}, err
	}
	for key, values := range r.Header {
		for _, value := range values {
			req.Header.Add(key, value)
		}
	}
	client := &http.Client{}
	resp, err := client.Do(req)
	return resp, err
}


func ForwardRequests(l conf.Location , r *http.Request) error {
  target1, err := url.Parse(fmt.Sprintf("http://localhost:%d", l.Port1))
  if err != nil {
    log.Fatal(err)
    return err
  }
  resp1, err1 := CallService(target1, r)
  if err1 != nil{
    return err1
  }
  log.Println(resp1)
  target2, err := url.Parse(fmt.Sprintf("http://localhost:%d", l.Port1))
  if err != nil {
    return err
  }
  resp2, err2 := CallService(target2, r)
  if err2 != nil{
    return err2
  }
  log.Println(resp2)
  return nil
}

func (s *Server) Handle(w http.ResponseWriter, r *http.Request) {
	path := r.URL.Path
  ForwardRequests(s.Conf[path], r)
}

func (s *Server) BuildRoutes() error {
	for path, loc := range s.Conf {
		log.Printf("%v", loc)
		s.Router.HandleFunc(path, s.Handle)
	}
	return nil
}
