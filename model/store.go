package model

import "log"

func (s *Store) PutValue(c chan error) {
	res, err := con.Exec("INSERT INTO STORE VALUES(?, ?)", s.Key, s.Value)
	if err != nil {
		c <- err
		return
	}
	log.Println(res.RowsAffected())
	c <- nil
	return
}

func (s Store) GetValue(c chan error) {
	row := con.QueryRow("SELECT value FROM STORE WHERE key = ?", s.Key)
	if err := row.Scan(&s.Value); err != nil {
		c <- err
		return
	}
	c <- nil
	return
}

func GetAllValues(c chan StoreReturn) {

	var storeArr []Store
	rows, err := con.Query("SELECT key, value FROM STORE WHERE")
	if err != nil {
		c <- StoreReturn{nil, err}
		return
	}
	for rows.Next() {
		var s Store
		err = rows.Scan(&s.Key, &s.Value)
		if err != nil {
			c <- StoreReturn{nil, err}
			return
		}
		storeArr = append(storeArr, s)
	}

	c <- StoreReturn{storeArr, nil}
	return
}
