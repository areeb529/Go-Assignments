package main

import (
	"sync"
	"testing"
)

func Test_increment(t *testing.T) {
	type fields struct {
		id    string
		count int
		m     sync.Mutex
	}
	tests := []struct {
		name   string
		field fields
	}{
		// TODO: Add test cases.
		{"test when view count is 0",fields{"hgds62",0,sync.Mutex{}}},
		{"test when view count is 1",fields{"hvhjs62",0,sync.Mutex{}}},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			v := &video{
				id:    tt.field.id,
				count: tt.field.count,
				m:     tt.field.m,
			}
			var wg sync.WaitGroup
			wg.Add(1)
			oldViewCount := v.count
			v.increment(&wg)
			newViewCount :=v.count
			expectedViewCount := oldViewCount + 1
			wg.Wait()
			if newViewCount != expectedViewCount{ 
				t.Errorf("Expected view count to be %d but got %d",expectedViewCount,newViewCount)
			}
		})
	}
}

func Test_viewCount(t *testing.T) {
	type fields struct {
		id    string
		count int
		m     sync.Mutex
	}
	tests := []struct {
		name   string
		field fields
	}{
		// TODO: Add test cases.
		{"test when view count is 0",fields{"hgds62",0,sync.Mutex{}}},
		{"test when view count is 1",fields{"hvhjs62",0,sync.Mutex{}}},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			v := &video{
				id:    tt.field.id,
				count: tt.field.count,
				m:     tt.field.m,
			}
			var wg sync.WaitGroup
			wg.Add(1)
			v.viewCount(&wg)
			wg.Wait()
		})
	}
}
