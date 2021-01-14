package gateway

import (
	"log"
	"net/http"
	"path/filepath"
	"regexp"
	"strings"
)
type HeaderPredicate string

func (this HeaderPredicate) Match(param http.Header) bool {
	s := string(this)
	if strings.Trim(s,"") == "" {
		return true
	}
	slist := strings.Split(s,",")
	if len(slist)<2 || len(slist)%2!=0 {
		return true
	}
	for i:=0;i<len(slist);i=i+2 {
		key := slist[i]
		pattern := slist[i+1]
		if value,ok := param[key];!ok {
			return false
		}else{
			reg,err := regexp.Compile(pattern)
			if err != nil {
				log.Println(err)
				return false
			}
			if !reg.MatchString(value[0]){
				return false
			}
		}
	}
	return true
}


type PathPredicate string

func (this PathPredicate) Match(param string) bool {
	if strings.Trim(string(this),"") == "" {
		return true
	}
	m,err := filepath.Match(string(this),param)
	if err != nil || !m {
		return false
	}
	return true
}
