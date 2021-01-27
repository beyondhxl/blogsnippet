package singleton

import "testing"

import "github.com/stretchr/testify/assert"

func TestGetInstance(t *testing.T) {
	assert.Equal(t, GetInstance(), GetInstance())
}

func BenchmarkGetInstanceParallel(b *testing.B) {
	b.RunParallel(func(pb *testing.PB) {
		for pb.Next() {
			if GetInstance() != GetInstance() {
				b.Errorf("BenchmarkGetInstanceParallel failed")
			}
		}
	})
}

func TestGetLazySingleInstance(t *testing.T) {
	assert.Equal(t, GetLazySingleInstance(), GetLazySingleInstance())
}

func BenchmarkGetLazySingleInstanceParallel(b *testing.B) {
	b.RunParallel(func(pb *testing.PB) {
		for pb.Next() {
			if GetLazySingleInstance() != GetLazySingleInstance() {
				b.Errorf("BenchmarkGetLazySingleInstanceParallel failed")
			}
		}
	})
}
