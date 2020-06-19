package models

type Books struct {
	Id          				int 				`orm:"column(id);pk;auto" description:"id" json:"id"`
	Title        				string				`orm:"column(title);size(255)" description:"title" json:"title"`
	Author						string				`orm:"column(author);size(255)" description:"author" json:"author"`
	Publisher					string				`orm:"column(publisher);size(255)" description:"publisher" json:"publisher"`
	Languages 					string				`orm:"column(languages);size(12)" description:"languages" json:"languages"`
	Published         			string         		`orm:"column(published)" description:"published" json:"published"`
	Identifiers					string				`orm:"column(identifiers)" description:"identifiers" json:"identifiers"`
}