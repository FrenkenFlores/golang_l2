package main

type AClass struct {
}

func (a *AClass) DoJobA() {

}

type BClass struct {
}

func (b *BClass) DoJobB() {

}

type Facade struct {
	aClass *AClass
	bClass *BClass
}

func (f *Facade) DoJob() {
	f.aClass.DoJobA()
	f.bClass.DoJobB()

}

func main() {
	facade := Facade{
		aClass: &AClass{},
		bClass: &BClass{},
	}
	facade.DoJob()
}
