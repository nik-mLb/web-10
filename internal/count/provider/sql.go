package provider

import (
)

func (p *Provider) SelectCounter() (int, error) {
	var msg int

	row := p.conn.QueryRow("SELECT number FROM counter WHERE id_number = 1")
	err := row.Scan(&msg)
	if err != nil {
		return -1, err
	}

	return msg, nil
}

func (p *Provider) UpdateCounter(msg int) error {
	_, err := p.conn.Exec("UPDATE counter SET number = number + $1 WHERE id_number = 1", msg)
	if err != nil {
		return err
	}

	return nil
}