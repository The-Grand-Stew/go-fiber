package models

type {{ .DomainName }} struct {

}

type {{ .DomainName }}Service interface{
	Create()(any,error)
	Get()(any,error)
	GetAll()(any,error)
	Update()(any,error)
	Delete()(any,error)
}
