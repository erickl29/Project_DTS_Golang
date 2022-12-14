package models

import (
	"training/src/config"
	"training/src/entities"
)

type ProductModel struct {
}

func (*ProductModel) FindAll() ([]entities.Product, error) {

	db, err := config.GetDB()
	if err != nil {
		return nil, err
	} else {
		rows, err2 := db.Query("select * from pegawai")
		if err2 != nil {
			return nil, err2
		} else {
			var products []entities.Product
			for rows.Next() {
				var pegawai entities.Product
				rows.Scan(&pegawai.NIP, &pegawai.Nama, &pegawai.Alamat, &pegawai.Nohp)
				products = append(products, pegawai)
			}
			return products, nil
		}
	}
}

func (*ProductModel) Find(NIP int64) (entities.Product, error) {

	db, err := config.GetDB()
	if err != nil {
		return entities.Product{}, err
	} else {
		rows, err2 := db.Query("select * from pegawai where id =?", NIP)
		if err2 != nil {
			return entities.Product{}, err2
		} else {
			var pegawai entities.Product
			for rows.Next() {
				rows.Scan(&pegawai.NIP, &pegawai.Nama, &pegawai.Alamat, &pegawai.Nohp)

			}
			return pegawai, nil
		}
	}
}

func (*ProductModel) Create(pegawai *entities.Product) bool {

	db, err := config.GetDB()
	if err != nil {
		return false
	} else {
		result, err2 := db.Exec("insert into pegawai(NIP, Nama, Alamat, Nohp) values (?,?,?,?)", pegawai.NIP, pegawai.Nama, pegawai.Alamat, pegawai.Nohp)
		if err2 != nil {
			return false
		} else {
			rowAffected, _ := result.RowsAffected()
			return rowAffected > 0
		}

	}
}

func (*ProductModel) Update(pegawai entities.Product) bool {

	db, err := config.GetDB()
	if err != nil {
		return false
	} else {
		result, err2 := db.Exec("update pegawai set Nama = ?, Alamat = ?, Nohp =? , where NIP = ?", pegawai.Nama, pegawai.Alamat, pegawai.Nohp, pegawai.NIP)
		if err2 != nil {
			return false
		} else {
			rowAffected, _ := result.RowsAffected()
			return rowAffected > 0
		}

	}
}

func (*ProductModel) Delete(NIP int64) bool {

	db, err := config.GetDB()
	if err != nil {
		return false
	} else {
		result, err2 := db.Exec("delete from pegawai where NIP =?", NIP)
		if err2 != nil {
			return false
		} else {
			rowAffected, _ := result.RowsAffected()
			return rowAffected > 0
		}

	}
}
