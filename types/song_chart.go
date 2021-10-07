package types

import (
	"encoding/json"
	"fmt"
)

type SongChart struct {
	Count    int    `json:"count"`
	SongName string `json:"songName"`
	SongId   int    `json:"-"`
}

func (s *SongChart) SetCount(count int) *SongChart {
	s.Count = count
	return s
}

func (s *SongChart) SetSongName(SongName string) *SongChart {
	s.SongName = SongName
	return s
}

func (s *SongChart) SetSongId(SongId int) *SongChart {
	s.SongId = SongId
	return s
}

func NewSongChart() *SongChart {
	return &SongChart{}
}

func (s *SongChart) Json() string {
	data, _ := json.Marshal(s)
	return string(data)
}

func (s *SongChart) InsertSql(table string) string {
	query := fmt.Sprintf(`INSERT INTO %s (count,songname) values (%d,'%s')`, table, s.Count, s.SongName)
	return query
}
