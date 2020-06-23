package models

type Book struct {
	Id          				int 				`orm:"column(id);pk;auto" description:"id" json:"id"`
	Type						string				`orm:"column(type);size(255)" description:"类型" json:"type"`
	//Path						string				`orm:"column(path);size(255)" description:"路径" json:"path"`
	Metadata 					*Metadata 			`json:"metadata" orm:"rel(fk);index"`
}