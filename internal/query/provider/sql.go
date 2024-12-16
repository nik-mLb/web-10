package provider

import (
	"database/sql"
)


func (p *Provider) SelectQuery(msg string) (bool, error) {
	var rec string

	row := p.conn.QueryRow("SELECT record FROM query WHERE record = ($1)", msg)
	err := row.Scan(&rec)
	if err != nil {
		if err == sql.ErrNoRows {
            return false, nil
        }
        return false, err
	}

	return true, nil
}

func (p *Provider) InsertQuery(msg string) error {
	_, err := p.conn.Exec("INSERT INTO query (record) VALUES ($1)", msg)
	if err != nil {
		return err
	}

	return nil
}