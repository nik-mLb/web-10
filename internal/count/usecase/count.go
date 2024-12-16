package usecase

func (u *Usecase) FetchCount() (int, error) {
	msg, err := u.p.SelectCounter()
	if err != nil {
		return 0, err
	}

	if msg == 0 {
		return u.defaultMsg, nil
	}

	return msg, nil
}

func (u *Usecase) SetCounter(msg int) error {
	err := u.p.UpdateCounter(msg)
	if err != nil {
		return err
	}

	return nil
}