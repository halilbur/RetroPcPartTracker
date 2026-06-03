package main

import "testing"

func TestDatabaseURLFromEnvReturnsConfiguredValue(t *testing.T) {
	t.Setenv("DATABASE_URL", "postgres://app:secret@db:5432/pcparts?sslmode=disable")

	got, err := databaseURLFromEnv()
	if err != nil {
		t.Fatalf("databaseURLFromEnv() returned error: %v", err)
	}

	want := "postgres://app:secret@db:5432/pcparts?sslmode=disable"
	if got != want {
		t.Fatalf("databaseURLFromEnv() = %q, want %q", got, want)
	}
}

func TestDatabaseURLFromEnvRequiresValue(t *testing.T) {
	t.Setenv("DATABASE_URL", "")

	_, err := databaseURLFromEnv()
	if err == nil {
		t.Fatal("databaseURLFromEnv() error = nil, want error")
	}
}
