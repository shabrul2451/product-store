package main

import "fmt"

type Song struct {
	ID   int
	Name string
}

type LRUCache struct {
	capacity int
	songs    map[int]*Song
	order    []int
}

func NewLRUCache(capacity int) *LRUCache {
	return &LRUCache{
		capacity: capacity,
		songs:    make(map[int]*Song),
		order:    make([]int, 0),
	}
}

func (lru *LRUCache) Get(songID int) (*Song, bool) {
	if song, ok := lru.songs[songID]; ok {
		lru.updateOrder(songID)
		return song, true
	}
	return nil, false
}

func (lru *LRUCache) Add(song *Song) {
	if _, ok := lru.songs[song.ID]; ok {
		lru.updateOrder(song.ID)
		return
	}

	if len(lru.songs) >= lru.capacity {
		oldestID := lru.order[0]
		delete(lru.songs, oldestID)
		lru.order = lru.order[1:]
	}

	lru.songs[song.ID] = song
	lru.order = append(lru.order, song.ID)
}

func (lru *LRUCache) updateOrder(songID int) {
	for i, id := range lru.order {
		if id == songID {
			copy(lru.order[i:], lru.order[i+1:])
			lru.order[len(lru.order)-1] = songID
			break
		}
	}
}

func main() {
	cache := NewLRUCache(3)

	cache.Add(&Song{ID: 1, Name: "Song A"})
	cache.Add(&Song{ID: 2, Name: "Song B"})
	cache.Add(&Song{ID: 3, Name: "Song C"})

	cache.Get(2)

	cache.Add(&Song{ID: 4, Name: "Song D"})

	cache.Get(3)
	cache.Add(&Song{ID: 5, Name: "Song E"})
	for _, songID := range cache.order {
		song := cache.songs[songID]
		fmt.Println(song.Name)
	}
}
