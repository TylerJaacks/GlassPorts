package main

import (
	"database/sql"
	"fmt"

	_ "github.com/go-sql-driver/mysql"
)

func openDB() *sql.DB {
	LogMsg("Opening database")
	data, err := sql.Open("mysql", fmt.Sprintf("%+v:%+v@%+v",
		config.DBUsername, config.DBPassword, config.DatabaseInfo))
	if err != nil {
		LogMsg("%+v:", err)
	}
	return data
}

func (u *User) Insert() {
	db := openDB()
	defer db.Close()
	LogMsg("Inserting %+v", u)
	insert, err := db.Query(fmt.Sprintf("INSERT INTO users VALUES ('%+v','%+v','%+v','%+v','%+v','%+v','%+v','%+v','%+v','%+v','%+v','%+v','%+v','%+v')",
		u.Identifier, u.Upload, u.Username, u.Email, u.Bio, u.Location, u.Twitter, u.Github, u.Keybase, u.Discord, u.LinkedIn, u.Website, u.Reddit, u.ProfilePic))
	if err != nil {
		LogMsg("%+v", err)
	}
	LogMsg("Insert complete. Closing!")
	insert.Close()

}

func (g *GlassApp) Insert() {
	db := openDB()
	defer db.Close()
	LogMsg("Inserting %+v", g)
	insert, err := db.Query(fmt.Sprintf("INSERT INTO glassapps VALUES ('%+v','%+v','%+v','%+v','%+v','%+v','%+v','%+v')",
		g.AppName, g.ShortDesc, g.LongDesc, g.APKLink, g.AppID, g.Screenshots, g.Maintainer, g.Icon))
	if err != nil {
		LogMsg("%+v", err)
	}
	LogMsg("Insert complete. Closing!")
	insert.Close()
}

func (c *CompanionApp) Insert() {
	db := openDB()
	defer db.Close()
	LogMsg("Inserting %+v", c)
	insert, err := db.Query(fmt.Sprintf("INSERT INTO companionapps VALUES ('%+v','%+v','%+v','%+v','%+v')",
		c.AppName, c.APKLink, c.GlassAppID, c.AppID, c.Icon))
	if err != nil {
		LogMsg("%+v", err)
	}
	LogMsg("Insert complete. Closing!")
	insert.Close()
}

func (d *DownloadStats) Insert() {
	db := openDB()
	defer db.Close()
	LogMsg("Inserting %+v", d)
	insert, err := db.Query(fmt.Sprintf("INSERT INTO companionapps VALUES ('%+v','%+v','%+v','%+v')",
		d.Identifier, d.Rating, d.Review, d.GlassAppID))
	if err != nil {
		LogMsg("%+v", err)
	}
	LogMsg("Insert complete. Closing!")
	insert.Close()
}

func (a *AuthToken) Insert() {
	db := openDB()
	defer db.Close()
	LogMsg("Inserting %+v", a)
	insert, err := db.Query(fmt.Sprintf("INSERT INTO authtoken VALUES ('%+v','%+v','%+v','%+v','%+v','%+v','%+v','%+v')",
		a.Identifier, a.Email, a.Username, a.PasswordHash, a.LastIP, a.AuthToken, a.PreviousHash, a.Current))
	if err != nil {
		LogMsg("%+v", err)
	}
	LogMsg("Insert complete. Closing!")
	insert.Close()
}

func (u *User) Retrieve() error {
	db := openDB()
	defer db.Close()
	var parms []interface{}
	statement := "SELECT * FROM users WHERE"
	first := true
	if u.Identifier != "" {
		parms = append(parms, u.Identifier)
		statement += " Identifier = ? "
		first = false
	}
	if u.Upload != "" {
		parms = append(parms, u.Upload)
		if !first {
			statement += " AND "
		} else {
			first = false
		}
		statement += " Upload = ? "

	}
	if u.Username != "" {
		parms = append(parms, u.Username)
		if !first {
			statement += " AND "
		} else {
			first = false
		}
		statement += " Username = ? "

	}
	if u.Email != "" {
		parms = append(parms, u.Email)
		if !first {
			statement += " AND "
		} else {
			first = false
		}
		statement += " Email = ? "

	}
	if u.Bio != "" {
		parms = append(parms, u.Bio)
		if !first {
			statement += " AND "
		} else {
			first = false
		}
		statement += " Bio = ? "

	}
	if u.Location != "" {
		parms = append(parms, u.Location)
		if !first {
			statement += " AND "
		} else {
			first = false
		}
		statement += " Location = ? "

	}
	if u.Twitter != "" {
		parms = append(parms, u.Twitter)
		if !first {
			statement += " AND "
		} else {
			first = false
		}
		statement += " Twitter = ? "

	}
	if u.Github != "" {
		parms = append(parms, u.Github)
		if !first {
			statement += " AND "
		} else {
			first = false
		}
		statement += " Github = ? "

	}
	if u.Keybase != "" {
		parms = append(parms, u.Keybase)
		if !first {
			statement += " AND "
		} else {
			first = false
		}
		statement += " Keybase = ? "

	}
	if u.Discord != "" {
		parms = append(parms, u.Discord)
		statement += " Discord = ? "
		first = false
	}
	if u.LinkedIn != "" {
		parms = append(parms, u.LinkedIn)
		if !first {
			statement += " AND "
		} else {
			first = false
		}
		statement += " LinkedIn = ? "

	}
	if u.Website != "" {
		parms = append(parms, u.Website)
		if !first {
			statement += " AND "
		} else {
			first = false
		}
		statement += " Website = ? "

	}
	if u.Reddit != "" {
		parms = append(parms, u.Reddit)
		if !first {
			statement += " AND "
		} else {
			first = false
		}
		statement += " Reddit = ? "

	}
	if u.ProfilePic != "" {
		parms = append(parms, u.ProfilePic)
		if !first {
			statement += " AND "
		} else {
			first = false
		}
		statement += " ProfilePic = ? "

	}
	LogMsg("Attempting to query using the following statement: %+v\nParams: %+v", statement, parms)
	rows, err := db.Query(statement, parms...)
	if err != nil {
		return err
	}
	for rows.Next() {
		rows.Scan(&u.Identifier, &u.Upload, &u.Username, &u.Email, &u.Bio, &u.Location,
			&u.Twitter, &u.Github, &u.Keybase, &u.Discord, &u.LinkedIn, &u.Website,
			&u.Reddit, &u.ProfilePic)
	}
	LogMsg("DB Returned: %+v", u)
	return nil

}

func (g *GlassApp) Retrieve() error {
	db := openDB()
	defer db.Close()
	var parms []interface{}
	statement := "SELECT * FROM glassapps WHERE"
	first := true

	if g.AppName != "" {
		parms = append(parms, g.AppName)
		if !first {
			statement += " AND "
		} else {
			first = false
		}
		statement += " AppName = ? "

	}
	if g.ShortDesc != "" {
		parms = append(parms, g.ShortDesc)
		if !first {
			statement += " AND "
		} else {
			first = false
		}
		statement += " ShortDesc = ? "

	}
	if g.LongDesc != "" {
		parms = append(parms, g.LongDesc)
		if !first {
			statement += " AND "
		} else {
			first = false
		}
		statement += " LongDesc = ? "
	}
	if g.APKLink != "" {
		parms = append(parms, g.APKLink)
		if !first {
			statement += " AND "
		} else {
			first = false
		}
		statement += " APKLink = ? "
	}
	if g.AppID != "" {
		parms = append(parms, g.AppID)
		if !first {
			statement += " AND "
		} else {
			first = false
		}
		statement += " AppID = ? "
	}
	if g.Screenshots != "" {
		parms = append(parms, g.Screenshots)
		if !first {
			statement += " AND "
		} else {
			first = false
		}
		statement += " Screenshots = ? "
	}
	if g.Maintainer != "" {
		parms = append(parms, g.Maintainer)
		if !first {
			statement += " AND "
		} else {
			first = false
		}
		statement += " Maintainer = ? "
	}
	if g.Icon != "" {
		parms = append(parms, g.Icon)
		if !first {
			statement += " AND "
		} else {
			first = false
		}
		statement += " Icon = ? "
	}

	LogMsg("Attempting to query using the following statement: %+v\nParams: %+v", statement, parms)
	rows, err := db.Query(statement, parms...)
	if err != nil {
		return err
	}

	for rows.Next() {
		rows.Scan(&g.AppName, &g.ShortDesc, &g.LongDesc, &g.APKLink, &g.AppID, &g.Screenshots, &g.Maintainer, &g.Icon)
	}
	LogMsg("DB Returned: %+v", g)
	return nil
}

func (g *GlassApp) RetrievePage(pageSize int32, page int32) ([]GlassApp, error) {
	db := openDB()
	defer db.Close()
	var parms []interface{}
	statement := "SELECT * FROM glassapps WHERE"
	first := true

	if g.AppName != "" {
		parms = append(parms, g.AppName)
		if !first {
			statement += " AND "
		} else {
			first = false
		}
		statement += " AppName = ? "

	}
	if g.ShortDesc != "" {
		parms = append(parms, g.ShortDesc)
		if !first {
			statement += " AND "
		} else {
			first = false
		}
		statement += " ShortDesc = ? "

	}
	if g.LongDesc != "" {
		parms = append(parms, g.LongDesc)
		if !first {
			statement += " AND "
		} else {
			first = false
		}
		statement += " LongDesc = ? "
	}
	if g.APKLink != "" {
		parms = append(parms, g.APKLink)
		if !first {
			statement += " AND "
		} else {
			first = false
		}
		statement += " APKLink = ? "
	}
	if g.AppID != "" {
		parms = append(parms, g.AppID)
		if !first {
			statement += " AND "
		} else {
			first = false
		}
		statement += " AppID = ? "
	}
	if g.Screenshots != "" {
		parms = append(parms, g.Screenshots)
		if !first {
			statement += " AND "
		} else {
			first = false
		}
		statement += " Screenshots = ? "
	}
	if g.Maintainer != "" {
		parms = append(parms, g.Maintainer)
		if !first {
			statement += " AND "
		} else {
			first = false
		}
		statement += " Maintainer = ? "
	}
	if g.Icon != "" {
		parms = append(parms, g.Icon)
		if !first {
			statement += " AND "
		} else {
			first = false
		}
		statement += " Icon = ? "
	}
	statement += fmt.Sprintf(" LIMIT %+v OFFSET %+v", pageSize, pageSize * page)
	LogMsg("Attempting to query using the following statement: %+v\nParams: %+v", statement, parms)
	rows, err := db.Query(statement, parms...)
	if err != nil {
		return nil, err
	}
	var ret []GlassApp
	for rows.Next() {
		var g2 GlassApp
		rows.Scan(&g2.AppName, &g2.ShortDesc, &g2.LongDesc, &g2.APKLink, &g2.AppID, &g2.Screenshots, &g2.Maintainer, &g2.Icon)
		ret = append(ret, g2)
	}
	LogMsg("DB Returned: %+v", g)
	return ret, nil
}

func (c *CompanionApp) Retrieve() error {
	db := openDB()
	defer db.Close()
	var parms []interface{}
	statement := "SELECT * FROM companionapp WHERE"
	first := true

	if c.AppName != "" {
		parms = append(parms, c.AppName)
		if !first {
			statement += " AND "
		} else {
			first = false
		}
		statement += " AppName = ? "

	}
	if c.APKLink != "" {
		parms = append(parms, c.APKLink)
		if !first {
			statement += " AND "
		} else {
			first = false
		}
		statement += " APKLink = ? "

	}
	if c.GlassAppID != "" {
		parms = append(parms, c.GlassAppID)
		if !first {
			statement += " AND "
		} else {
			first = false
		}
		statement += " GlassAppID = ? "

	}
	if c.AppID != "" {
		parms = append(parms, c.AppID)
		if !first {
			statement += " AND "
		} else {
			first = false
		}
		statement += " AppID = ? "

	}
	LogMsg("Attempting to query using the following statement: %+v\nParams: %+v", statement, parms)
	rows, err := db.Query(statement, parms...)
	if err != nil {
		return err
	}
	for rows.Next() {
		rows.Scan(&c.AppName, &c.APKLink, &c.GlassAppID, &c.AppID)
	}
	LogMsg("DB Returned: %+v", c)
	return nil
}

func (c *CompanionApp) RetrievePage(pageSize int32, page int32) ([]CompanionApp, error) {
	db := openDB()
	defer db.Close()
	var parms []interface{}
	statement := "SELECT * FROM companionapp WHERE"
	first := true

	if c.AppName != "" {
		parms = append(parms, c.AppName)
		if !first {
			statement += " AND "
		} else {
			first = false
		}
		statement += " AppName = ? "

	}
	if c.APKLink != "" {
		parms = append(parms, c.APKLink)
		if !first {
			statement += " AND "
		} else {
			first = false
		}
		statement += " APKLink = ? "

	}
	if c.GlassAppID != "" {
		parms = append(parms, c.GlassAppID)
		if !first {
			statement += " AND "
		} else {
			first = false
		}
		statement += " GlassAppID = ? "

	}
	if c.AppID != "" {
		parms = append(parms, c.AppID)
		if !first {
			statement += " AND "
		} else {
			first = false
		}
		statement += " AppID = ? "

	}
	statement += fmt.Sprintf(" LIMIT %+v OFFSET %+v", pageSize, pageSize * page)
	LogMsg("Attempting to query using the following statement: %+v\nParams: %+v", statement, parms)
	rows, err := db.Query(statement, parms...)
	if err != nil {
		return nil, err
	}

	var ret []CompanionApp
	for rows.Next() {
		var c2 CompanionApp
		rows.Scan(&c2.AppName, &c2.APKLink, &c2.GlassAppID, &c2.AppID)
	ret = append(ret, c2)
	}
	LogMsg("DB Returned: %+v", c)
	return ret, nil
}


func (d *DownloadStats) Retrieve() error {
	db := openDB()
	defer db.Close()
	var parms []interface{}
	statement := "SELECT * FROM downloadstats WHERE"
	first := true

	if d.Identifier != "" {
		parms = append(parms, d.Identifier)
		if !first {
			statement += " AND "
		} else {
			first = false
		}
		statement += " Identifier = ? "

	}
	if d.Rating != "" {
		parms = append(parms, d.Rating)
		if !first {
			statement += " AND "
		} else {
			first = false
		}
		statement += " Rating = ? "

	}
	if d.Review != "" {
		parms = append(parms, d.Review)
		if !first {
			statement += " AND "
		} else {
			first = false
		}
		statement += " Review = ? "

	}
	if d.GlassAppID != "" {
		parms = append(parms, d.GlassAppID)
		if !first {
			statement += " AND "
		} else {
			first = false
		}
		statement += " GlassAppID = ? "

	}
	LogMsg("Attempting to query using the following statement: %+v\nParams: %+v", statement, parms)
	rows, err := db.Query(statement, parms...)
	if err != nil {
		return err
	}
	for rows.Next() {
		rows.Scan(&d.Identifier, &d.Rating, &d.Review, &d.GlassAppID)
	}
	LogMsg("DB Returned: %+v", d)
	return nil
}

func (d *DownloadStats) RetrievePage(pageSize int32, page int32) ([]DownloadStats, error) {
	db := openDB()
	defer db.Close()
	var parms []interface{}
	statement := "SELECT * FROM downloadstats WHERE"
	first := true

	if d.Identifier != "" {
		parms = append(parms, d.Identifier)
		if !first {
			statement += " AND "
		} else {
			first = false
		}
		statement += " Identifier = ? "

	}
	if d.Rating != "" {
		parms = append(parms, d.Rating)
		if !first {
			statement += " AND "
		} else {
			first = false
		}
		statement += " Rating = ? "

	}
	if d.Review != "" {
		parms = append(parms, d.Review)
		if !first {
			statement += " AND "
		} else {
			first = false
		}
		statement += " Review = ? "

	}
	if d.GlassAppID != "" {
		parms = append(parms, d.GlassAppID)
		if !first {
			statement += " AND "
		} else {
			first = false
		}
		statement += " GlassAppID = ? "

	}
	statement += fmt.Sprintf(" LIMIT %+v OFFSET %+v", pageSize, pageSize * page)
	LogMsg("Attempting to query using the following statement: %+v\nParams: %+v", statement, parms)
	rows, err := db.Query(statement, parms...)
	if err != nil {
		return nil, err
	}
	var ret []DownloadStats
	for rows.Next() {
		var d2 DownloadStats
		rows.Scan(&d2.Identifier, &d2.Rating, &d2.Review, &d2.GlassAppID)
	ret = append(ret, d2)
	}
	LogMsg("DB Returned: %+v", d)
	return ret, nil
}

func (a *AuthToken) Retrieve() error {
	db := openDB()
	defer db.Close()
	var parms []interface{}
	statement := "SELECT * FROM authtoken WHERE"
	first := true
	if a.Identifier != "" {
		parms = append(parms, a.Identifier)
		statement += " Identifier = ? "
		first = false
	}
	if a.Email != "" {
		parms = append(parms, a.Email)
		if !first {
			statement += " AND "
		} else {
			first = false
		}
		statement += " Email = ? "

	}
	if a.Username != "" {
		parms = append(parms, a.Username)
		if !first {
			statement += " AND "
		} else {
			first = false
		}
		statement += " Username = ? "

	}
	if a.PasswordHash != "" {
		parms = append(parms, a.PasswordHash)
		if !first {
			statement += " AND "
		} else {
			first = false
		}
		statement += " PasswordHash = ? "

	}
	if a.LastIP != "" {
		parms = append(parms, a.LastIP)
		if !first {
			statement += " AND "
		} else {
			first = false
		}
		statement += " LastIP = ? "

	}
	if a.AuthToken != "" {
		parms = append(parms, a.AuthToken)
		if !first {
			statement += " AND "
		} else {
			first = false
		}
		statement += " AuthToken = ? "

	}
	if a.PreviousHash != "" {
		parms = append(parms, a.PreviousHash)
		if !first {
			statement += " AND "
		} else {
			first = false
		}
		statement += " PreviousHash = ? "

	}
	if a.Current != "" {
		parms = append(parms, a.Current)
		if !first {
			statement += " AND "
		} else {
			first = false
		}
		statement += " Current = ? "

	}

	LogMsg("Attempting to query using the following statement: %+v\nParams: %+v", statement, parms)
	rows, err := db.Query(statement, parms...)
	if err != nil {
		return err
	}
	for rows.Next() {
		rows.Scan(&a.Identifier, &a.Email, &a.Username, &a.PasswordHash, &a.LastIP, &a.AuthToken,
			&a.PreviousHash, &a.Current)
	}
	LogMsg("DB Returned: %+v", a)
	return nil

}

func (a *AuthToken) RetrievePage(pageSize int32, page int32) ([]AuthToken, error) {
	db := openDB()
	defer db.Close()
	var parms []interface{}
	statement := "SELECT * FROM authtoken WHERE"
	first := true
	if a.Identifier != "" {
		parms = append(parms, a.Identifier)
		statement += " Identifier = ? "
		first = false
	}
	if a.Email != "" {
		parms = append(parms, a.Email)
		if !first {
			statement += " AND "
		} else {
			first = false
		}
		statement += " Email = ? "

	}
	if a.Username != "" {
		parms = append(parms, a.Username)
		if !first {
			statement += " AND "
		} else {
			first = false
		}
		statement += " Username = ? "

	}
	if a.PasswordHash != "" {
		parms = append(parms, a.PasswordHash)
		if !first {
			statement += " AND "
		} else {
			first = false
		}
		statement += " PasswordHash = ? "

	}
	if a.LastIP != "" {
		parms = append(parms, a.LastIP)
		if !first {
			statement += " AND "
		} else {
			first = false
		}
		statement += " LastIP = ? "

	}
	if a.AuthToken != "" {
		parms = append(parms, a.AuthToken)
		if !first {
			statement += " AND "
		} else {
			first = false
		}
		statement += " AuthToken = ? "

	}
	if a.PreviousHash != "" {
		parms = append(parms, a.PreviousHash)
		if !first {
			statement += " AND "
		} else {
			first = false
		}
		statement += " PreviousHash = ? "

	}
	if a.Current != "" {
		parms = append(parms, a.Current)
		if !first {
			statement += " AND "
		} else {
			first = false
		}
		statement += " Current = ? "

	}
	statement += fmt.Sprintf(" LIMIT %+v OFFSET %+v", pageSize, pageSize * page)
	LogMsg("Attempting to query using the following statement: %+v\nParams: %+v", statement, parms)
	rows, err := db.Query(statement, parms...)
	if err != nil {
		return nil, err
	}
	var ret []AuthToken
	for rows.Next() {
		var a2 AuthToken
		rows.Scan(&a2.Identifier, &a2.Email, &a2.Username, &a2.PasswordHash, &a2.LastIP, &a2.AuthToken,
			&a2.PreviousHash, &a2.Current)
		ret = append(ret, a2)
	}
	
	LogMsg("DB Returned: %+v", a)
	return ret, nil

}

func (u *User) Delete() {
	db := openDB()
	defer db.Close()
	LogMsg("Deleting %+v", u)
	delete, err := db.Query(fmt.Sprintf("DELETE FROM users WHERE Identifier='%+v' AND Upload='%+v' AND Username='%+v' AND Email='%+v' AND Bio='%+v' AND Location='%+v' AND Twitter='%+v' AND Github='%+v' AND Keybase='%+v' AND Discord='%+v' AND LinkedIn='%+v' AND Website='%+v' AND Reddit='%+v' AND ProfilePic='%+v'",
		u.Identifier, u.Upload, u.Username, u.Email, u.Bio, u.Location, u.Twitter, u.Github, u.Keybase, u.Discord, u.LinkedIn, u.Website, u.Reddit, u.ProfilePic))
	if err != nil {
		LogMsg("%+v", err)
	}
	LogMsg("Delete complete. Closing!")
	delete.Close()
}
func (g *GlassApp) Delete() {
	db := openDB()
	defer db.Close()
	LogMsg("Deleting %+v", g)
	delete, err := db.Query(fmt.Sprintf("DELETE FROM glassapps WHERE AppName='%+v' AND ShortDesc='%+v' AND LongDesc='%+v' AND APKLink='%+v' AND AppID='%+v' AND Screenshots='%+v' AND Maintainer='%+v' AND Icon='%+v'",
		g.AppName, g.ShortDesc, g.LongDesc, g.APKLink, g.AppID, g.Screenshots, g.Maintainer, g.Icon))
	if err != nil {
		LogMsg("%+v", err)
	}
	LogMsg("Delete complete. Closing!")
	delete.Close()
}
func (c *CompanionApp) Delete() {
	db := openDB()
	defer db.Close()
	LogMsg("Deleting %+v", c)
	delete, err := db.Query(fmt.Sprintf("DELETE FROM companionapps WHERE AppName='%+v' AND APKLink='%+v' AND GlassAppID='%+v' AND AppID='%+v' AND Icon='%+v'",
		c.AppName, c.APKLink, c.GlassAppID, c.AppID, c.Icon))
	if err != nil {
		LogMsg("%+v", err)
	}
	LogMsg("Delete complete. Closing!")
	delete.Close()
}

func (d *DownloadStats) Delete() {
	db := openDB()
	defer db.Close()
	LogMsg("Deleting %+v", d)
	delete, err := db.Query(fmt.Sprintf("DELETE FROM companionapps WHERE Identifier='%+v' AND Rating='%+v' AND Review='%+v' AND GlassAppID='%+v'",
		d.Identifier, d.Rating, d.Review, d.GlassAppID))
	if err != nil {
		LogMsg("%+v", err)
	}
	LogMsg("Delete complete. Closing!")
	delete.Close()
}

func (a *AuthToken) Delete() {
	db := openDB()
	defer db.Close()
	LogMsg("Deleting %+v", a)
	delete, err := db.Query(fmt.Sprintf("DELETE FROM authtoken WHERE Identifier='%+v' AND Email='%+v' AND Username='%+v' AND PasswordHash='%+v' AND LastIP='%+v' AND AuthToken='%+v' AND PreviousHash='%+v' AND Current='%+v'",
		a.Identifier, a.Email, a.Username, a.PasswordHash, a.LastIP, a.AuthToken, a.PreviousHash, a.Current))
	if err != nil {
		LogMsg("%+v", err)
	}
	LogMsg("Delete complete. Closing!")
	delete.Close()
}
