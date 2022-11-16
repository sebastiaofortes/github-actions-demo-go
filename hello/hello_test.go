package hello

import "testing"

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