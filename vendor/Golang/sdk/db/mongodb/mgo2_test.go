package mongodb

import (
	"context"
	"fmt"
	"testing"
)

func TestMgo2(t *testing.T) {
	// PerformOperations creates a candle item
	// then gets it
	PerformOperations := func(s Storage) error {
		ctx := context.Background()
		i := Item{Name: "candles", Price: 100}
		if err := s.Put(ctx, &i); err != nil {
			return err
		}

		candles, err := s.GetByName(ctx, "candles")
		if err != nil {
			return err
		}
		fmt.Printf("Result: %#v\n", candles)
		return nil
	}

	m, err := NewMongoStorage("localhost", "gocookbook", "items")
	if err != nil {
		t.Fatal(err)
	}
	if err := PerformOperations(m); err != nil {
		t.Fatal(err)
	}

	if err := m.Session.DB(m.DB).C(m.Collection).DropCollection(); err != nil {
		t.Fatal(err)
	}
}
