package flyweight

import "fmt"

func (f *ObjFlyweightFactory) Get(objname string) *ObjFlyweight {
	obj := f.maps[objname]
	if obj == nil {
		obj = NewObjFlyweight(objname)
		f.maps[objname] = obj
	}
	return obj
}

type ObjFlyweightFactory struct {
	maps map[string]*ObjFlyweight
}

var g_ObjFactory *ObjFlyweightFactory

func GetObjFlyweightFactory() *ObjFlyweightFactory {
	if g_ObjFactory == nil {
		g_ObjFactory = &ObjFlyweightFactory{
			maps: make(map[string]*ObjFlyweight),
		}
	}
	return g_ObjFactory
}

// 享元对象
type ObjFlyweight struct {
	data string
}

func NewObjFlyweight(objname string) *ObjFlyweight {
	// Load image file
	data := fmt.Sprintf("data %s", objname)
	return &ObjFlyweight{
		data: data,
	}
}

func (o *ObjFlyweight) Data() string {
	return o.data
}

type ObjDisplay struct {
	*ObjFlyweight
}

func NewObjDisplay(objname string) *ObjDisplay {
	obj := GetObjFlyweightFactory().Get(objname)
	return &ObjDisplay{
		ObjFlyweight: obj,
	}
}

func (odisp *ObjDisplay) Display() {
	fmt.Printf("Display: %s\n", odisp.Data())
}
