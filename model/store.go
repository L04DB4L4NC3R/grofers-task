package model

func (s *Store) PutValue(c chan error) {
	_, err := con.Exec("INSERT INTO STORE VALUES(?, ?)", s.Key, s.Value)
	if err != nil {
		c <- err
		return
	}
	c <- nil
	return
}

func (s *Store) GetValue(c chan error) {
	row := con.QueryRow("SELECT v FROM STORE WHERE k = ?", s.Key)
	if err := row.Scan(&s.Value); err != nil {
		c <- err
		return
	}
	c <- nil
	return
}

func (s *Store) DelValue(c chan error) {

	// get key
	row := con.QueryRow("SELECT v FROM STORE WHERE k = ?", s.Key)
	if err := row.Scan(&s.Value); err != nil {
		c <- err
		return
	}

	// delete from DB
	_, err := con.Exec("DELETE FROM STORE WHERE k = ?", s.Key)
	if err != nil {
		c <- err
		return
	}
	c <- nil
	return
}

func (s *Store) UpdateValue(val string, c chan error) {

	// update from DB
	_, err := con.Exec("UPDATE STORE SET v = ? WHERE k = ?", val, s.Key)
	if err != nil {
		c <- err
		return
	}

	// get key
	row := con.QueryRow("SELECT v FROM STORE WHERE k = ?", s.Key)
	if err := row.Scan(&s.Value); err != nil {
		c <- err
		return
	}

	c <- nil
	return
}
