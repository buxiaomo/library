package models

type Metadata struct {
	Id          				int 				`orm:"column(id);pk;auto" description:"id" json:"id"`
	Isbn						string				`orm:"column(isbn);size(255)" description:"ISBN" json:"isbn"`
	Title        				string				`orm:"column(title);size(255)" description:"书名" json:"title"`
	Author						string				`orm:"column(author);size(255)" description:"作者" json:"author"`
	Publisher					string				`orm:"column(publisher);size(255)" description:"出版社" json:"publisher"`
	Languages 					string				`orm:"column(languages);size(12)" description:"语言" json:"languages"`
	Published         			string         		`orm:"column(published)" description:"发布时间" json:"published"`
	Comments					string				`orm:"column(comments)" description:"简介" json:"comments"`
	Picture						string				`orm:"column(picture)" description:"封面" json:"picture"`
	Score						string				`orm:"column(score)" description:"评分" json:"score"`
	Books      					[]*Book 			`json:"books" orm:"reverse(many)"`

	//Mobi						bool				`orm:"column(mobi)" description:"mobi" json:"mobi"`
	//Epub						bool				`orm:"column(epub)" description:"epub" json:"epub"`
	//Azw3						bool				`orm:"column(azw3)" description:"azw3" json:"azw3"`
	//Txt							bool				`orm:"column(txt)" description:"txt" json:"txt"`
	//Pdf							bool				`orm:"column(pdf)" description:"pdf" json:"pdf"`
}