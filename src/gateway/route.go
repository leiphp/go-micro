package gateway

import (
	"net/http"
)

type Predicates struct {
	Header HeaderPredicate
	Method string
	Host string
	Path PathPredicate
}

type Filter struct {

}

type Route struct {
	Id string
	Url string
	Predicates Predicates
	Filters []*Filter
}

type Routes []*Route

func(this Routes) Match(request *http.Request) *Route{
	for _,r := range this{
		if this.isMatch(r, request){
			return r
		}
	}
	return nil
}

func (this Routes) isMatch(r *Route,request *http.Request) bool {
	if !r.Predicates.Path.Match(request.URL.Path){
		return false
	}
	if !r.Predicates.Header.Match(request.Header){
		return false
	}
	return true
}