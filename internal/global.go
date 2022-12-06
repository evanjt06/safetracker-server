package internal

import (
	"errors"
	"fmt"
	util "github.com/aldelo/common"
	data "github.com/aldelo/common/wrapper/mysql"
)

// ---------------------------------------------------------------------------------------------------------------------
// Public Database Connection Functions
// ---------------------------------------------------------------------------------------------------------------------
var _writerDBHost string
var _writerDBPort int
var _writerDBName string
var _writerDBUserName string
var _writerDBPassword string

var _readerDBHost string
var _readerDBPort int
var _readerDBName string
var _readerDBUserName string
var _readerDBPassword string

// SetWriterDBInfo sets the writer db connection parameter info
// if port is <= 0 then default port 3306 for mysql is used
func SetWriterDBInfo(host string, port int, dbname string, username string, password string) error {
	// validate
	if util.LenTrim(host) == 0 {
		return errors.New("MySql Writer DB Host is Required")
	}

	if util.LenTrim(dbname) == 0 {
		return errors.New("MySql Writer Database Name is Required")
	}

	if util.LenTrim(username) == 0 {
		return errors.New("MySql Writer DB User Name is Required")
	}

	if util.LenTrim(password) == 0 {
		return errors.New("MySql Writer DB User Password is Required")
	}

	// set default
	if port <= 0 {
		port = 3306
	}

	// set values
	_writerDBHost = host
	_writerDBPort = port
	_writerDBName = dbname
	_writerDBUserName = username
	_writerDBPassword = password

	// success
	return nil
}

// SetReaderDBInfo sets the reader db connection parameter info
// if port is <= 0 then default port 3306 for mysql is used
func SetReaderDBInfo(host string, port int, dbname string, username string, password string) error {
	// validate
	if util.LenTrim(host) == 0 {
		return errors.New("MySql Reader DB Host is Required")
	}

	if util.LenTrim(dbname) == 0 {
		return errors.New("MySql Reader Database Name is Required")
	}

	if util.LenTrim(username) == 0 {
		return errors.New("MySql Reader DB User Name is Required")
	}

	if util.LenTrim(password) == 0 {
		return errors.New("MySql Reader DB User Password is Required")
	}

	// set default
	if port <= 0 {
		port = 3306
	}

	// set values
	_readerDBHost = host
	_readerDBPort = port
	_readerDBName = dbname
	_readerDBUserName = username
	_readerDBPassword = password

	// success
	return nil
}

// ConnectToWriterDB will establish connection to writer db as configured by SetWriterDBInfo,
// it will also cache the writer db object into global cache
func ConnectToWriterDB() error {
	// -----------------------------------------------------------------------------------------------------------------
	// prepare
	// -----------------------------------------------------------------------------------------------------------------

	// close writer db first
	closeWriterDB()

	// validate
	if util.LenTrim(_writerDBHost) == 0 {
		return errors.New("Writer DB Host is Required")
	}

	if util.LenTrim(_writerDBName) == 0 {
		return errors.New("Writer Database Name is Required")
	}

	if util.LenTrim(_writerDBUserName) == 0 {
		return errors.New("Writer DB User Name is Required")
	}

	if util.LenTrim(_writerDBPassword) == 0 {
		return errors.New("Writer DB User Password is Required")
	}

	if _writerDBPort <= 0 {
		_writerDBPort = 3306
	}

	// -----------------------------------------------------------------------------------------------------------------
	// mysql connection setup
	// -----------------------------------------------------------------------------------------------------------------
	mysCn := new(data.MySql)

	mysCn.Host = _writerDBHost
	mysCn.Port = _writerDBPort
	mysCn.Database = _writerDBName
	mysCn.UserName = _writerDBUserName
	mysCn.Password = _writerDBPassword

	if err := mysCn.Open(); err != nil {
		closeWriterDB()
		return errors.New("Open MySQL Writer DB Failed: " + err.Error())
	} else {
		setWriterDB(mysCn)
		return nil
	}
}

// DisconnectFromWriterDB will close connection from the writer db
func DisconnectFromWriterDB() {
	// -----------------------------------------------------------------------------------------------------------------
	// close mysql db connection
	// -----------------------------------------------------------------------------------------------------------------
	closeWriterDB()
}

// ConnectToReaderDB will establish connection to writer db as configured by SetReaderDBInfo,
// it will also cache the reader db object into global cache
func ConnectToReaderDB() error {
	// -----------------------------------------------------------------------------------------------------------------
	// prepare
	// -----------------------------------------------------------------------------------------------------------------

	// close reader db first
	closeReaderDB()

	// validate
	if util.LenTrim(_readerDBHost) == 0 {
		return errors.New("Reader DB Host is Required")
	}

	if util.LenTrim(_readerDBName) == 0 {
		return errors.New("Reader Database Name is Required")
	}

	if util.LenTrim(_readerDBUserName) == 0 {
		return errors.New("Reader DB User Name is Required")
	}

	if util.LenTrim(_readerDBPassword) == 0 {
		return errors.New("Reader DB User Password is Required")
	}

	if _readerDBPort <= 0 {
		_readerDBPort = 3306
	}

	// -----------------------------------------------------------------------------------------------------------------
	// mysql connection setup
	// -----------------------------------------------------------------------------------------------------------------
	mysCn := new(data.MySql)

	mysCn.Host = _readerDBHost
	mysCn.Port = _readerDBPort
	mysCn.Database = _readerDBName
	mysCn.UserName = _readerDBUserName
	mysCn.Password = _readerDBPassword

	if err := mysCn.Open(); err != nil {
		closeReaderDB()
		return errors.New("Open MySQL Reader DB Failed: " + err.Error())
	} else {
		setReaderDB(mysCn)
		return nil
	}
}

// DisconnectFromReaderDB will close connection from the reader db
func DisconnectFromReaderDB() {
	// -----------------------------------------------------------------------------------------------------------------
	// close mysql db connection
	// -----------------------------------------------------------------------------------------------------------------
	closeReaderDB()
}

// IsWriterDBReady checks if the writer db is ready for operations
func IsWriterDBReady() bool {
	if db == nil {
		return false
	}

	if err := db.Ping(); err != nil {
		return false
	}

	return true
}

// IsReaderDBReady checks if the reader db is ready for operations
func IsReaderDBReady() bool {
	if dbReader == nil {
		return false
	}

	if err := dbReader.Ping(); err != nil {
		return false
	}

	return true
}

// ---------------------------------------------------------------------------------------------------------------------
// Private Database connection functions
// ---------------------------------------------------------------------------------------------------------------------

// package level accessible to the mysql server database object
var db *data.MySql

// package level accessible to the mysql server reader database object
var dbReader *data.MySql

// setWriterDB allows mysql writer database reference to be set
func setWriterDB(dbx *data.MySql) {
	db = dbx
}

// closeWriterDB closes writer db connection and cleans up object reference
func closeWriterDB() {
	if db != nil {
		_ = db.Close()
	}
}

// setReaderDB allows mysql reader database reference to be set
func setReaderDB(dbx *data.MySql) {
	dbReader = dbx
}

// closeReaderDB closes reader db connection and cleans up object reference
func closeReaderDB() {
	if dbReader != nil {
		_ = dbReader.Close()
	}
}

// getReaderDB will attempt to return the reader db and if reader db is nil, the writer db will be used instead
func getReaderDB() *data.MySql {
	if dbReader == nil {
		return db
	} else {
		return dbReader
	}
}

// ---------------------------------------------------------------------------------------------------------------------
// Transaction Functions
// ---------------------------------------------------------------------------------------------------------------------

// BeginTran starts writer db transaction
func BeginTran() (*data.MySqlTransaction, error) {
	if db != nil {
		return db.Begin()
	}

	// fail when no db writer object
	return nil, fmt.Errorf("BeginTran Failed, db Writer Object Nil")
}
