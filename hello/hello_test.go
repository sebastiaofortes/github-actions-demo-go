package hello

import (
	"fmt"
	"testing"
)

func TestGreetsGitHub(t *testing.T) {
	result := Greet()
	if result != "Hello GitHub Actions. Dev.to is awesome" {
		t.Errorf("Greet() = %s; want Hello GitHub Actions. Dev.to is awesome", result)
	}
}

func TestGreetsGitHub2(t *testing.T) {
	result := Greet()
	if result != "Hello GitHub Actions. Dev.to is awesome" {
		t.Errorf("Greet() = %s; want Hello GitHub Actions. Dev.to is awesome", result)
	}
}

func TestGreetsGitHub3(t *testing.T) {
	result := Greet()
	if result != "Hello GitHub Actions. Dev.to is awesome" {
		t.Errorf("Greet() = %s; want Hello GitHub Actions. Dev.to is awesome", result)
	}
}
func TestGreetsGitHub4(t *testing.T) {
	result := Greet()
	if result != "Hello GitHub Actions. Dev.to is awesome" {
		t.Errorf("Greet() = %s; want Hello GitHub Actions. Dev.to is awesome", result)
	}
}
func TestSoma(t *testing.T) {
	inteiros := []int64{1,2,3,4,5}
	fmt.Println(Soma(inteiros))
	flutuantes := []float32{1.5,2,3,4}
	fmt.Println(Soma(flutuantes))
	
}