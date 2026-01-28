package store

import (
	"database/sql"
	"time"
)

type Part struct {
	ID        int
	Name      string
	Type      string
	Brand     string
	Specs     string
	ImageURL  string
	CreatedAt time.Time
}

type Price struct {
	ID           int
	PartID       int
	DealerID     int
	Price        float64
	Currency     string
	InStock      bool
	LastUpdated  time.Time
	DealerName   string
	DealerURL    string
	DealerRating float64
}

type Dealer struct {
	ID                 int
	Name               string
	URL                string
	AuthenticityRating float64
	IsVerified         bool
}

type PartStore struct {
	db *sql.DB
}

func NewPartStore(db *sql.DB) *PartStore {
	return &PartStore{db: db}
}

func (s *PartStore) GetAllParts() ([]Part, error) {
	query := `SELECT id, name, type, brand, specs, image_url, created_at 
	          FROM parts ORDER BY created_at DESC LIMIT 100`

	rows, err := s.db.Query(query)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	var parts []Part
	for rows.Next() {
		var p Part
		err := rows.Scan(&p.ID, &p.Name, &p.Type, &p.Brand, &p.Specs, &p.ImageURL, &p.CreatedAt)
		if err != nil {
			return nil, err
		}
		parts = append(parts, p)
	}
	return parts, nil
}

func (s *PartStore) SearchParts(query string) ([]Part, error) {
	sqlQuery := `SELECT id, name, type, brand, specs, image_url, created_at 
	             FROM parts 
	             WHERE name ILIKE $1 OR brand ILIKE $1 OR specs ILIKE $1
	             ORDER BY created_at DESC LIMIT 50`

	rows, err := s.db.Query(sqlQuery, "%"+query+"%")
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	var parts []Part
	for rows.Next() {
		var p Part
		err := rows.Scan(&p.ID, &p.Name, &p.Type, &p.Brand, &p.Specs, &p.ImageURL, &p.CreatedAt)
		if err != nil {
			return nil, err
		}
		parts = append(parts, p)
	}
	return parts, nil
}

func (s *PartStore) GetPartsByType(partType string) ([]Part, error) {
	query := `SELECT id, name, type, brand, specs, image_url, created_at 
	          FROM parts WHERE type = $1 ORDER BY name`

	rows, err := s.db.Query(query, partType)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	var parts []Part
	for rows.Next() {
		var p Part
		err := rows.Scan(&p.ID, &p.Name, &p.Type, &p.Brand, &p.Specs, &p.ImageURL, &p.CreatedAt)
		if err != nil {
			return nil, err
		}
		parts = append(parts, p)
	}
	return parts, nil
}

func (s *PartStore) GetPartWithPrices(partID int) (*Part, []Price, error) {
	// Get part details
	var part Part
	err := s.db.QueryRow(`
		SELECT id, name, type, brand, specs, image_url, created_at 
		FROM parts WHERE id = $1
	`, partID).Scan(&part.ID, &part.Name, &part.Type, &part.Brand, &part.Specs, &part.ImageURL, &part.CreatedAt)

	if err != nil {
		return nil, nil, err
	}

	// Get prices from all dealers
	rows, err := s.db.Query(`
		SELECT p.id, p.part_id, p.dealer_id, p.price, p.currency, p.in_stock, p.last_updated,
		       d.name, d.url, d.authenticity_rating
		FROM prices p
		JOIN dealers d ON p.dealer_id = d.id
		WHERE p.part_id = $1
		ORDER BY p.price ASC
	`, partID)

	if err != nil {
		return &part, nil, err
	}
	defer rows.Close()

	var prices []Price
	for rows.Next() {
		var price Price
		err := rows.Scan(
			&price.ID, &price.PartID, &price.DealerID, &price.Price, &price.Currency,
			&price.InStock, &price.LastUpdated, &price.DealerName, &price.DealerURL,
			&price.DealerRating,
		)
		if err != nil {
			return &part, nil, err
		}
		prices = append(prices, price)
	}

	return &part, prices, nil
}
