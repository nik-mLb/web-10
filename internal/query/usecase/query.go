package usecase

func (u *Usecase) FetchQuery(name string) (bool, error) {
	msg, err := u.p.SelectQuery(name)
	if msg && err != nil {
		return msg, err
	}

	if name == "" {
		return true, nil
	}

	return msg, nil
}

func (u *Usecase) SetQuery(name string) error {
	err := u.p.InsertQuery(name)
	if err != nil {
		return err
	}

	return nil
}