package types

var Songs = map[int]string{
	1: "The Good The Bad And The Ugly played",
	2: "Believe",
	3: "Still Loving You",
	4: "Perfect",
	5: "Bohemian Rhapsody",
	6: "Sometimes",
	7: "Into The Unknown",
}

type Player struct {
	Log map[int]*SongChart
}

func NewPlayer() *Player {
	return &Player{
		Log: map[int]*SongChart{},
	}
}

func (p *Player) PlayRandomSong() *SongChart {
	for id, song := range Songs {
		chart := p.Log[id]
		if chart == nil {
			chart = NewSongChart().
				SetSongId(id).
				SetSongName(song).
				SetCount(1)
		} else {
			chart.SetCount(chart.Count + 1)
		}
		p.Log[id] = chart
		return chart
	}
	return nil
}
