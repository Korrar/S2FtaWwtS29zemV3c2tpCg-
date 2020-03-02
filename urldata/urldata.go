package urldata


type Adder interface {
	Add(item Item)
}

type Getter interface {
	GetAll() []Item
}

type Deleter interface {
	Delete(id int)
}

type Exist interface {
	Exists(id int)
}

type Item struct {
	Id   int `json:"ID"`
	Url  string `json:"url"`
	Interval int `json:"interval"`
}

type ItemSlice struct {
	items []Item
}
func New() *ItemSlice {
	return &ItemSlice{
		items: []Item{},
	}
}


func (p *ItemSlice) Add(item Item) {
	p.items = append(p.items, item)
}

func (p *ItemSlice) GetAll() []Item {
	return p.items
}


func (p *ItemSlice) Delete(id int)  {
	for i := range p.items {
		if p.items[i].Id == id{
			p.items = append(p.items[:i], p.items[i+1:]...)
			break
		}
	}

}

func (p *ItemSlice) Exists(id int) bool {
	for i := range p.items {
		if p.items[i].Id == id{
			return true
		}
	}
	return false
}



