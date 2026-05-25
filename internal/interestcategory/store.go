package interestcategory

import (
	"context"
	"database/sql"
)

type mysqlStore struct {
	db *sql.DB
}

func NewStore(db *sql.DB) Repository {
	return &mysqlStore{db: db}
}

const selectCols = "SELECT category_id, parent_id, name, slug, category_group, is_sensitive FROM interest_categories"

func (s *mysqlStore) List(ctx context.Context) ([]InterestCategory, error) {
	rows, err := s.db.QueryContext(ctx, selectCols)
	if err != nil {
		return nil, err
	}
	defer rows.Close()
	return scanAll(rows)
}

func (s *mysqlStore) GetByID(ctx context.Context, id string) (*InterestCategory, error) {
	rows, err := s.db.QueryContext(ctx, selectCols+" WHERE category_id = ?", id)
	if err != nil {
		return nil, err
	}
	defer rows.Close()
	if !rows.Next() {
		return nil, nil
	}
	return scanCategory(rows)
}

func (s *mysqlStore) Create(ctx context.Context, c *InterestCategory) error {
	_, err := s.db.ExecContext(ctx,
		"INSERT INTO interest_categories (category_id, parent_id, name, slug, category_group, is_sensitive) VALUES (?, ?, ?, ?, ?, ?)",
		c.CategoryID, c.ParentID, c.Name, c.Slug, c.CategoryGroup, c.IsSensitive,
	)
	return err
}

func (s *mysqlStore) Update(ctx context.Context, c *InterestCategory) error {
	_, err := s.db.ExecContext(ctx,
		"UPDATE interest_categories SET parent_id = ?, name = ?, slug = ?, category_group = ?, is_sensitive = ? WHERE category_id = ?",
		c.ParentID, c.Name, c.Slug, c.CategoryGroup, c.IsSensitive, c.CategoryID,
	)
	return err
}

func (s *mysqlStore) Delete(ctx context.Context, id string) error {
	_, err := s.db.ExecContext(ctx, "DELETE FROM interest_categories WHERE category_id = ?", id)
	return err
}

type scanner interface {
	Scan(dest ...any) error
}

func scanAll(rows *sql.Rows) ([]InterestCategory, error) {
	var items []InterestCategory
	for rows.Next() {
		c, err := scanCategory(rows)
		if err != nil {
			return nil, err
		}
		items = append(items, *c)
	}
	return items, rows.Err()
}

func scanCategory(r scanner) (*InterestCategory, error) {
	var c InterestCategory
	var parentID, slug, categoryGroup sql.NullString
	var isSensitive sql.NullBool
	if err := r.Scan(&c.CategoryID, &parentID, &c.Name, &slug, &categoryGroup, &isSensitive); err != nil {
		return nil, err
	}
	if parentID.Valid {
		c.ParentID = &parentID.String
	}
	if slug.Valid {
		c.Slug = &slug.String
	}
	if categoryGroup.Valid {
		c.CategoryGroup = &categoryGroup.String
	}
	if isSensitive.Valid {
		c.IsSensitive = &isSensitive.Bool
	}
	return &c, nil
}
