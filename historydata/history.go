package historydata

import "time"

type Adder interface {
	Add(item History)
}

type Getter interface {
	GetAll() []History
}


type History struct {
	Response  string  `json:"response"`
	Duration  time.Duration`json:"duration"`
	CreatedAt time.Time `json:"created_at"`
	Id int `json:"-"`
}

type HistorySlice struct {
	history []History
}

func New() *HistorySlice {
	return &HistorySlice{
		history: []History{},
	}
}


func (p *HistorySlice) Add(his History) {
	p.history = append(p.history, his)
}

func (p *HistorySlice) GetAll() []History {
	return p.history
}
