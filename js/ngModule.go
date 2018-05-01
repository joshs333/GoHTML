package js

//defines a n angular modular
type Module interface {
	Load()        //to be run before page is loaded
	ToJS() string //to be done during runtime
}
