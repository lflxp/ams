package config

import (
	"fmt"
)

type Html string

var Htmls Html

func init() {
	Htmls = "html生成器"
}
//参数 
//id,href,name,info
func (this *Html) Heading(info string) string {
	return fmt.Sprintf(`<li class="list-group-item">%s</li>`,info)
}

//href = #id
func (this *Html) A(href,name string,more bool) string {
	var result string
	if more == true {
		result = fmt.Sprintf(`<span class="glyphicon glyphicon-star" aria-hidden="true"></span><a href="%s" data-toggle="tab" > %s </a><span class="badge">42</span>`,href,name)
	} else {
		result = fmt.Sprintf(`<a href="%s" data-toggle="tab" role="tab">%s</a>`,href,name) 
	}
	return result
}

func (this *Html) A2(href,name string,more bool) string {
	var result string
	if more == true {
		result = fmt.Sprintf(`<span class="glyphicon glyphicon-star" aria-hidden="true"></span><a href="%s" class="list-group-item" data-toggle="tab" > %s </a><span class="badge">42</span>`,href,name)
	} else {
		result = fmt.Sprintf(`<a href="%s" data-toggle="tab" class="list-group-item" role="tab">%s</a>`,href,name) 
	}
	return result
}

func (this *Html) Tab(id,info string) string {
	return fmt.Sprintf(`<div class="tab-pane fade" id="%s" role="tabpanel">%s</div>`,id,info)
}

func (this *Html) Create(id,href,name,info string,more bool) (string,string) {
	return this.Heading(this.A(href,name,more)),this.Tab(id,info) 
}

func (this *Html) Create2(id,href,name,info string,more bool) (string,string) {
	return this.A2(href,name,more),this.Tab(id,info) 
}