package mlib

import "errors"

type MusicEntry struct {
	Id     string
	Name   string
	Artist string
	Source string
	Type   string
}

type MusicManager struct {
	musics []MusicEntry
}

func NewMusicManager() *MusicManager {
	return &MusicManager{make([]MusicEntry, 0)}
}

func (m *MusicManager) Len() int {
	return len(m.musics)
}

func (m *MusicManager) Get(index int) (music *MusicEntry, err error) {
	if index < 0 || index >= len(m.musics) {
		return nil, errors.New("Index out of range.")
	}
	return &m.musics[index], nil
}

func (m *MusicManager) Find(name string) *MusicEntry {
	if len(m.musics) == 0 {
		return nil
	}

	for _, m := range m.musics {
		if m.Name == name {
			return &m
		}
	}
	return nil
}

func (m *MusicManager) Add(music *MusicEntry) {
	m.musics = append(m.musics, *music)
}

func (m *MusicManager) Remove(index int) *MusicEntry {
	if index < 0 || index >= len(m.musics) {
		return nil
	}

	removedMusic := &m.musics[index]

	m.musics = append(m.musics[:index], m.musics[index+1:]...)
	return removedMusic
}

func (m *MusicManager) RemoveByName(name string) *MusicEntry {
	if name == "" {
		return nil
	}

	//获取需要删的值的索引
	index := -1
	for i, m := range m.musics {
		if m.Name == name {
			index = i
			break
		}
	}

	if index != -1 {
		return m.Remove(index)
	}
	return nil
}
