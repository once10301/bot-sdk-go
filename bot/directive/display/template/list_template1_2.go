package template

type ListTemplate1 struct {
	ListTemplate
}

func NewListTemplate1() *ListTemplate1 {
	listTemplate := &ListTemplate1{}
	listTemplate.Type = "ListTemplate1"
	listTemplate.Token = listTemplate.GenToken()
	return listTemplate
}

type ListTemplate2 struct {
	ListTemplate
}

func NewListTemplate2() *ListTemplate2 {
	listTemplate := &ListTemplate2{}
	listTemplate.Type = "ListTemplate2"
	listTemplate.Token = listTemplate.GenToken()
	return listTemplate
}

type ListTemplate3 struct {
	ListTemplate
}

func NewListTemplate3() *ListTemplate3 {
	listTemplate := &ListTemplate3{}
	listTemplate.Type = "ListTemplate3"
	listTemplate.Token = listTemplate.GenToken()
	return listTemplate
}

type ListTemplate4 struct {
	ListTemplate
}

func NewListTemplate4() *ListTemplate4 {
	listTemplate := &ListTemplate4{}
	listTemplate.Type = "ListTemplate4"
	listTemplate.Token = listTemplate.GenToken()
	return listTemplate
}
