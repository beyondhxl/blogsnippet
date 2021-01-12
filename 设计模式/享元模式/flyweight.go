package flyweight

import "fmt"

func (f *ObjFlyweightFactory) Get(objName string) *ObjFlyweight {
	obj := f.maps[objName]
	if obj == nil {
		obj = NewObjFlyweight(objName)
		f.maps[objName] = obj
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

func NewObjFlyweight(objName string) *ObjFlyweight {
	// Load image file
	data := fmt.Sprintf("data %s", objName)

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

func NewObjDisplay(objName string) *ObjDisplay {
	obj := GetObjFlyweightFactory().Get(objName)
	return &ObjDisplay{
		ObjFlyweight: obj,
	}
}

func (odisp *ObjDisplay) Display() {
	fmt.Printf("Display: %s\n", odisp.Data())
}
