package internal

import (
	"errors"
	util "github.com/aldelo/common"
	data "github.com/aldelo/common/wrapper/mysql"
	"strings"
	"time"
)

// ---------------------------------------------------------------------------------------------------------------------
// DiscordBotChild Actions
// ---------------------------------------------------------------------------------------------------------------------

// indicates if DiscordBotChild struct will prefer dbReader over dbWriter where available
var _DiscordBotChildPrefersDBReader bool

// UseDBReaderPreferred sets indicator that this struct prefers to use db reader where applicable and if available
func (r DiscordBotChild) UseDBReaderPreferred() {
	// set db reader usage preference to true, so if db reader is available, we will use it if code is applicable for it
	_DiscordBotChildPrefersDBReader = true
}

// UseDBWriterPreferred sets indicator that this struct prefers to use db writer (this is the default preference)
func (r DiscordBotChild) UseDBWriterPreferred() {
	// set db writer usage preference to true, so this struct will use the db writer for operations
	_DiscordBotChildPrefersDBReader = false
}

// Fill is a helper function to fill in required or important fields of this struct,
// this function consolidates all required or important fields into one input path to simply data input,
// note: values filled into this function is not yet persisted to database, to persist, must call Set() function
// note: values filled into this function is not validated at this point yet, call Validate() function to validate if desired
// note: this function will have struct tag ggAttr:"8 = Fill" as fields to include as parameters
// [ Parameters ]
//		DiscordBotID int64 = REQUIRED
//		ImageURL string = REQUIRED
//		TextContent string = REQUIRED
//		Datetime time.Time = REQUIRED
//		AuthorDiscordID string = REQUIRED
//		AuthorDiscordTag string = REQUIRED
func (r *DiscordBotChild) Fill(DiscordBotID int64, ImageURL string, TextContent string, Datetime time.Time, AuthorDiscordID string, AuthorDiscordTag string) {
	r.DiscordBotID = DiscordBotID
	r.ImageURL = ImageURL
	r.TextContent = TextContent
	r.Datetime = Datetime
	r.AuthorDiscordID = AuthorDiscordID
	r.AuthorDiscordTag = AuthorDiscordTag
}

// GetDiscordBotID retrieves struct field value,
// if any field data type is int representing enum, then the corresponding enum is returned,
// if any field data type is sql null type or pointer, then the null or pointer is converted to base type
func (r *DiscordBotChild) GetDiscordBotID() int64 {
	return r.DiscordBotID
}

// GetImageURL retrieves struct field value,
// if any field data type is int representing enum, then the corresponding enum is returned,
// if any field data type is sql null type or pointer, then the null or pointer is converted to base type
func (r *DiscordBotChild) GetImageURL() string {
	return r.ImageURL
}

// GetTextContent retrieves struct field value,
// if any field data type is int representing enum, then the corresponding enum is returned,
// if any field data type is sql null type or pointer, then the null or pointer is converted to base type
func (r *DiscordBotChild) GetTextContent() string {
	return r.TextContent
}

// GetDatetime retrieves struct field value,
// if any field data type is int representing enum, then the corresponding enum is returned,
// if any field data type is sql null type or pointer, then the null or pointer is converted to base type
func (r *DiscordBotChild) GetDatetime() time.Time {
	return r.Datetime
}

// GetAuthorDiscordID retrieves struct field value,
// if any field data type is int representing enum, then the corresponding enum is returned,
// if any field data type is sql null type or pointer, then the null or pointer is converted to base type
func (r *DiscordBotChild) GetAuthorDiscordID() string {
	return r.AuthorDiscordID
}

// GetAuthorDiscordTag retrieves struct field value,
// if any field data type is int representing enum, then the corresponding enum is returned,
// if any field data type is sql null type or pointer, then the null or pointer is converted to base type
func (r *DiscordBotChild) GetAuthorDiscordTag() string {
	return r.AuthorDiscordTag
}

// SetDiscordBotID sets value into struct field value and checks for data input initial validation rules,
// if any field data type is int representing enum, then the corresponding enum is used as input parameter,
// if any field data type is sql null type or pointer, then the base type is used as parameter and converted into null or pointer
//		[ Parameters ]
//		v int64 = Struct internal data type: int64; REQUIRED
func (r *DiscordBotChild) SetDiscordBotID(v int64) error {
	r.DiscordBotID = v

	return nil
}

// SetImageURL sets value into struct field value and checks for data input initial validation rules,
// if any field data type is int representing enum, then the corresponding enum is used as input parameter,
// if any field data type is sql null type or pointer, then the base type is used as parameter and converted into null or pointer
//		[ Parameters ]
//		v string = Struct internal data type: string; REQUIRED
func (r *DiscordBotChild) SetImageURL(v string) error {
	if util.LenTrim(v) == 0 {
		return errors.New("SetImageURL Failed, A Text Length Greater Than Zero is Required")
	}

	r.ImageURL = v

	return nil
}

// SetTextContent sets value into struct field value and checks for data input initial validation rules,
// if any field data type is int representing enum, then the corresponding enum is used as input parameter,
// if any field data type is sql null type or pointer, then the base type is used as parameter and converted into null or pointer
//		[ Parameters ]
//		v string = Struct internal data type: string; REQUIRED
func (r *DiscordBotChild) SetTextContent(v string) error {
	if util.LenTrim(v) == 0 {
		return errors.New("SetTextContent Failed, A Text Length Greater Than Zero is Required")
	}

	r.TextContent = v

	return nil
}

// SetDatetime sets value into struct field value and checks for data input initial validation rules,
// if any field data type is int representing enum, then the corresponding enum is used as input parameter,
// if any field data type is sql null type or pointer, then the base type is used as parameter and converted into null or pointer
//		[ Parameters ]
//		v time.Time = Struct internal data type: time.Time; REQUIRED
func (r *DiscordBotChild) SetDatetime(v time.Time) error {
	if v.IsZero() {
		return errors.New("SetDatetime Failed, A Valid Date or Time is Required")
	}

	r.Datetime = v

	return nil
}

// SetAuthorDiscordID sets value into struct field value and checks for data input initial validation rules,
// if any field data type is int representing enum, then the corresponding enum is used as input parameter,
// if any field data type is sql null type or pointer, then the base type is used as parameter and converted into null or pointer
//		[ Parameters ]
//		v string = Struct internal data type: string; REQUIRED
func (r *DiscordBotChild) SetAuthorDiscordID(v string) error {
	if util.LenTrim(v) == 0 {
		return errors.New("SetAuthorDiscordID Failed, A Text Length Greater Than Zero is Required")
	}

	r.AuthorDiscordID = v

	return nil
}

// SetAuthorDiscordTag sets value into struct field value and checks for data input initial validation rules,
// if any field data type is int representing enum, then the corresponding enum is used as input parameter,
// if any field data type is sql null type or pointer, then the base type is used as parameter and converted into null or pointer
//		[ Parameters ]
//		v string = Struct internal data type: string; REQUIRED
func (r *DiscordBotChild) SetAuthorDiscordTag(v string) error {
	if util.LenTrim(v) == 0 {
		return errors.New("SetAuthorDiscordTag Failed, A Text Length Greater Than Zero is Required")
	}

	r.AuthorDiscordTag = v

	return nil
}

// GetByID retrieves a row from database using the ID field value,
// and marshals found result fields into struct,
// an error of nil indicates success
// note: this function will have struct tag ggAttr:"9 = GetBy" for columns acting as its source
func (r *DiscordBotChild) GetByID(ID int64) (notFound bool, err error) {
	// clean up
	r._originalValue = nil

	// compose query
	q := new(data.QueryBuilder)

	q.Set("SELECT * FROM DiscordBotChild WHERE ID=? LIMIT 1;")
	q.Ordinal(ID)

	// execute query
	var dbCurrent *data.MySql

	if !_DiscordBotChildPrefersDBReader {
		dbCurrent = db
	} else {
		dbCurrent = getReaderDB()
	}

	notFound, err = dbCurrent.GetStruct(r, q.SQL(), q.ParamsSlice()...)
	dbCurrent = nil

	if err != nil {
		// error detected
		return false, err
	}

	if notFound {
		// not found but not error
		return true, nil
	}

	// store into original value
	r._originalValue = *r

	// success
	return false, nil
}

// GetByDiscordBotID retrieves a row from database using the DiscordBotID field value,
// and marshals found result fields into struct,
// an error of nil indicates success
// note: this function will have struct tag ggAttr:"9 = GetBy" for columns acting as its source
func (r *DiscordBotChild) GetByDiscordBotID(DiscordBotID int64) (notFound bool, err error) {
	// clean up
	r._originalValue = nil

	// compose query
	q := new(data.QueryBuilder)

	q.Set("SELECT * FROM DiscordBotChild WHERE DiscordBotID=? LIMIT 1;")
	q.Ordinal(DiscordBotID)

	// execute query
	var dbCurrent *data.MySql

	if !_DiscordBotChildPrefersDBReader {
		dbCurrent = db
	} else {
		dbCurrent = getReaderDB()
	}

	notFound, err = dbCurrent.GetStruct(r, q.SQL(), q.ParamsSlice()...)
	dbCurrent = nil

	if err != nil {
		// error detected
		return false, err
	}

	if notFound {
		// not found but not error
		return true, nil
	}

	// store into original value
	r._originalValue = *r

	// success
	return false, nil
}

// GetScalar supports flexible get query to retrieve a scalar value from database
// [ Parameters ]
//		scalarColumnName = the scalar value column name to retrieve
//		filter = optional sql filter portion to be injected into query (if ignored, then all records are selected and top result limit 1 is applied)
//					if query needs to override FROM ... WHERE ..., then filter value begin with FROM to indicate override required (FROM AnotherTable WHERE ... or FROM AnotherTable INNER JOIN AnotherTable2 WHERE ...)
//		sort = optional sql sort portion to be injected into query
//		args = optional ordinal parameter arguments to pass into the filter query portion (must appear in ordinal position of params defined within filter)
func (r *DiscordBotChild) GetScalar(scalarColumnName string, filter string, sort string, args ...interface{}) (result string, notFound bool, err error) {
	// validate
	if util.LenTrim(scalarColumnName) == 0 {
		return "", false, errors.New("ScalarColumnName is Required")
	}

	// is this FROM override?
	isFromOverride := false

	if util.Left(strings.ToUpper(util.Trim(filter)), 5) == "FROM " {
		isFromOverride = true
	}

	// pre-process where clause
	filter = util.Trim(filter)

	if !isFromOverride {
		if util.LenTrim(filter) > 0 {
			if util.Left(strings.ToUpper(filter), 5) == "WHERE" {
				filter += " "
			} else {
				filter = "WHERE " + filter + " "
			}
		}
	}

	// pre-process sort clause
	sort = util.Trim(sort)
	if util.LenTrim(sort) > 0 {
		// has custom sort
		if util.Left(strings.ToUpper(sort), 8) != "ORDER BY" {
			sort = "ORDER BY " + sort
		}

		// remove; if exists
		if util.Right(sort, 1) == ";" {
			sort = util.Left(sort, len(sort)-1)
		}
	} else {
		// no custom sort
		sort = "ORDER BY ID ASC"
	}

	// compose query
	q := new(data.QueryBuilder)

	if !isFromOverride {
		q.Set("SELECT " + scalarColumnName + " FROM DiscordBotChild " + filter + sort + " LIMIT 1;")
	} else {
		q.Set("SELECT " + scalarColumnName + " " + filter + sort + " LIMIT 1;")
	}

	if args != nil {
		if len(args) > 0 {
			for _, v := range args {
				q.Ordinal(v)
			}
		}
	}

	// execute query
	var dbCurrent *data.MySql

	if !_DiscordBotChildPrefersDBReader {
		dbCurrent = db
	} else {
		dbCurrent = getReaderDB()
	}

	if len(q.ParamsSlice()) > 0 {
		// has parameters
		result, notFound, err = dbCurrent.GetScalarString(q.SQL(), q.ParamsSlice()...)
	} else {
		// no parameters
		result, notFound, err = dbCurrent.GetScalarString(q.SQL())
	}
	dbCurrent = nil

	// evaluate execute response
	if err != nil {
		return "", false, errors.New("GetScalar Value Failed: " + err.Error())
	}

	if notFound {
		// if nothing is found, do not return error
		return "", true, nil
	}

	// success
	return result, false, nil
}

// IsDataChanged checks current struct's participating fields if its current data vs original data has changed
// note: this function will have struct tag ggAttr:"7 = Check" for participating columns
func (r *DiscordBotChild) IsDataChanged() bool {
	if r._originalValue == nil {
		// default to data changed status
		return true
	}

	// if primary key is 0 then this is new record, assume data changed
	if r.ID <= 0 {
		return true
	}

	// assertion of original value to struct for old value
	old := r._originalValue.(DiscordBotChild)

	// check if old value vs current value was changed
	if r.DiscordBotID != old.DiscordBotID {
		return true
	}

	if r.ImageURL != old.ImageURL {
		return true
	}

	if r.TextContent != old.TextContent {
		return true
	}

	if r.Datetime != old.Datetime {
		return true
	}

	if r.AuthorDiscordID != old.AuthorDiscordID {
		return true
	}

	if r.AuthorDiscordTag != old.AuthorDiscordTag {
		return true
	}

	// no changes
	return false
}

// Validate checks if the struct fields' data conforms to the expected data integrity rules,
// such as required field is met, string fields size min max is met, numeric range is met,
// returns nil is validate is successful, otherwise returns error containing the validation failure reason
func (r *DiscordBotChild) Validate() error {
	// check for required fields
	// based on ggAttr:"3 = - (required)" and ggAttr:"6 = Set-Upd, Set-InsUpd, Set-Ins"
	if r.DiscordBotID == 0 {
		return errors.New("DiscordBotID is Required")
	}

	if util.LenTrim(r.ImageURL) == 0 {
		return errors.New("ImageURL is Required")
	}

	if util.LenTrim(r.TextContent) == 0 {
		return errors.New("TextContent is Required")
	}

	if r.Datetime.IsZero() {
		return errors.New("Datetime is Required")
	}

	if util.LenTrim(r.AuthorDiscordID) == 0 {
		return errors.New("AuthorDiscordID is Required")
	}

	if util.LenTrim(r.AuthorDiscordTag) == 0 {
		return errors.New("AuthorDiscordTag is Required")
	}

	// success
	return nil
}

// Set persists the current struct fields into the underlying database table row,
// a new database row will be inserted if such was not found,
// otherwise existing database row will be updated
// note: this function will use struct tag ggAttr:"1 = PK-Custom"
//       and ggAttr:"6 = Set-Ins, Set-Upd, Set-InsUpd" as participating columns
func (r *DiscordBotChild) Set() error {
	// check if need to persist data to database via set
	if !r.IsDataChanged() {
		// data not changed, no need to set
		return nil
	}

	// declare error variable
	var err error

	// validate values
	if err = r.Validate(); err != nil {
		// validation error
		return err
	}

	// ready to persist data to database
	// create query to insert or update into database
	// note: if ID is not auto generated, then for new row ID will still be set (need to evaluate _originalValue = nil for new row condition)
	isNewRow := false
	q := new(data.QueryBuilder)

	if r.ID <= 0 {
		// insert
		isNewRow = true

		// compose insert action query
		q.Set("INSERT INTO DiscordBotChild ")
		q.Set("(DiscordBotID, ImageURL, TextContent, Datetime, AuthorDiscordID, AuthorDiscordTag) ")
		q.Set("VALUES ")
		q.Set("(:DiscordBotID, :ImageURL, :TextContent, :Datetime, :AuthorDiscordID, :AuthorDiscordTag);")

		q.Named("DiscordBotID", r.DiscordBotID)
		q.Named("ImageURL", r.ImageURL)
		q.Named("TextContent", r.TextContent)
		q.Named("Datetime", r.Datetime)
		q.Named("AuthorDiscordID", r.AuthorDiscordID)
		q.Named("AuthorDiscordTag", r.AuthorDiscordTag)
	} else {
		// update
		isNewRow = false

		// compose update action query
		q.Set("UPDATE DiscordBotChild ")
		q.Set("SET DiscordBotID=:DiscordBotID, ImageURL=:ImageURL, TextContent=:TextContent, Datetime=:Datetime, AuthorDiscordID=:AuthorDiscordID, AuthorDiscordTag=:AuthorDiscordTag ")
		q.Set("WHERE ID=:ID;")

		q.Named("ID", r.ID)
		q.Named("DiscordBotID", r.DiscordBotID)
		q.Named("ImageURL", r.ImageURL)
		q.Named("TextContent", r.TextContent)
		q.Named("Datetime", r.Datetime)
		q.Named("AuthorDiscordID", r.AuthorDiscordID)
		q.Named("AuthorDiscordTag", r.AuthorDiscordTag)
	}

	// begin transaction
	var tx *data.MySqlTransaction
	var e error

	if tx, e = db.Begin(); e != nil {
		return errors.New("DiscordBotChild Persist Data Failed at BeginTran: " + e.Error())
	}

	// execute query
	result := tx.ExecByNamedMapParam(q.SQL(), q.ParamsMap())

	if result.Err != nil {
		if err1 := tx.Rollback(); err1 != nil {
			return errors.New("DiscordBotChild Persist Data Failed at RollBack 1: " + err1.Error() + " (Orig Error: " + result.Err.Error() + ")")
		} else {
			return errors.New("DiscordBotChild Persist Data Failed: " + result.Err.Error())
		}
	}

	if !isNewRow {
		// was update
		if result.RowsAffected <= 0 {
			// update failure
			if err1 := tx.Rollback(); err1 != nil {
				return errors.New("DiscordBotChild Persist Data Failed at RollBack 2: " + err1.Error() + " (Orig Error: No Rows Affected)")
			} else {
				return errors.New("DiscordBotChild Persist Updated Data Failed: No Rows Affected")
			}
		}
	} else {
		// was insert
		// if ID was Auto Generated, we will need to assign into this struct ID field
		// if ID was NOT Auto Generated, then we don't evaluate newly generated id since ID already contain the PK value
		if result.NewlyInsertedID > 0 {
			r.ID = result.NewlyInsertedID
		} else if r.ID <= 0 {
			// if newly inserted id is zero then this is insert failure
			if err1 := tx.Rollback(); err1 != nil {
				return errors.New("DiscordBotChild Persist Data Failed at RollBack 3: " + err1.Error() + " (Orig Error: No New Primary ID Generated)")
			} else {
				return errors.New("DiscordBotChild Persist Inserted Data Failed: No New Primary ID Generated")
			}
		}
	}

	r._originalValue = nil
	r._originalValue = *r

	// commit persist data action
	if err1 := tx.Commit(); err1 != nil {
		return errors.New("DiscordBotChild Persist Data Failed at Commit: " + err1.Error())
	}

	// return nil as success
	return nil
}

// physicalDelete internal helper
func (r *DiscordBotChild) physicalDelete() error {
	// compose query
	q := new(data.QueryBuilder)

	q.Set("DELETE FROM DiscordBotChild WHERE ID=?;")
	q.Ordinal(r.ID)

	// execute query
	result := db.ExecByOrdinalParams(q.SQL(), q.ParamsSlice()...)

	if result.Err != nil {
		return errors.New("DiscordBotChild Physical Delete Failed: " + result.Err.Error())
	}

	if result.RowsAffected <= 0 {
		return errors.New("DiscordBotChild Physical Delete Failed: No Rows Affected")
	}

	// physical delete successful
	return nil
}

// Delete removes the current struct record from the underlying database table row
func (r *DiscordBotChild) Delete() error {
	// validate
	if r.ID <= 0 {
		return errors.New("DiscordBotChild Delete Data Failed: Row Primary ID is Required")
	}

	if r._originalValue == nil {
		return errors.New("DiscordBotChild Delete Data Failed: Row Must Be Loaded First")
	}

	// perform delete action
	// physical delete
	return r.physicalDelete()
}

// ---------------------------------------------------------------------------------------------------------------------
// DiscordBotChildList Actions
// ---------------------------------------------------------------------------------------------------------------------

// DiscordBotChildList collection level dal model
type DiscordBotChildList struct {
	List  *[]DiscordBotChild
	Count int
}

// indicates if DiscordBotChild list will prefer dbReader over dbWriter where available
var _DiscordBotChildListPrefersDBReader bool

// UseDBReaderPreferred sets indicator that this struct prefers to use db reader where applicable and if available
func (l DiscordBotChildList) UseDBReaderPreferred() {
	// set db reader usage preference to true, so if db reader is available, we will use it if code is applicable for it
	_DiscordBotChildListPrefersDBReader = true
}

// UseDBWriterPreferred sets indicator that this struct prefers to use db writer (this is the default preference)
func (l DiscordBotChildList) UseDBWriterPreferred() {
	// set db writer usage preference to true, so this struct list will use the db writer for operations
	_DiscordBotChildListPrefersDBReader = false
}

// Element gets one of the slice elements in list, identified by ordinal index position
func (l *DiscordBotChildList) Element(position int) (*DiscordBotChild, error) {
	// validate
	if l == nil {
		return nil, errors.New("List Object Nil")
	}

	// position cannot be less than 0
	if position < 0 {
		return nil, errors.New("Element Position to Retrieve Cannot Be Less Than 0")
	}

	// ensure position is valid
	if position >= l.Count {
		// not valid, because position cannot be same or greater than count
		return nil, errors.New("Element Position Must Be Less Than Count in List Object")
	}

	// ensure list is valid
	if l.List == nil {
		// list must not be nil
		return nil, errors.New("List Must Not Be Nil")
	}

	// ensure list is valid
	if len(*l.List) == 0 {
		// list must not be 0
		return nil, errors.New("List Count Must Not Be Zero")
	}

	// ensure list count is valid
	if position >= len(*l.List) {
		// list count must not be same or greater than position
		return nil, errors.New("Element Position Must Be Less Than Actual List Count")
	}

	// get element from position within slice
	return &(*l.List)[position], nil
}

// GetAll loads all active records into list struct
func (l *DiscordBotChildList) GetAll(limitCount int, offsetCount ...int) error {
	offset := util.GetFirstIntOrDefault(0, offsetCount...)
	return l.getInternal("", "", offset, limitCount)
}

// GetByID will load struct slice for one or more matching values of same parameter from database
// [ Parameters ]
//		IDToLoad = variadic, one or more table ID to load from database, for example, enter 2, 3, 6, 9, will load records with ID 2, 3, 6, and 9 from database to struct list
// [ Notes ]
// 		this function is based on the database column dbAttr:"GetBy" for suffix naming
func (l *DiscordBotChildList) GetByID(IDToLoad ...int64) error {
	// validate
	if l == nil {
		return errors.New("List Object Nil")
	}

	if IDToLoad == nil {
		return errors.New("Get Requires One or More Filter Values (Input Nil)")
	}

	if len(IDToLoad) == 0 {
		return errors.New("Get Requires One or More Filter Values (Input Count 0)")
	}

	// compose filter
	filter := ""

	for _, v := range IDToLoad {
		if util.LenTrim(filter) > 0 {
			filter += ", "
		}

		filter += util.Int64ToString(v)
	}

	if strings.Contains(filter, ",") {
		filter = "WHERE ID IN (" + filter + ")"
	} else {
		filter = "WHERE ID = " + filter
	}

	// perform action
	return l.getInternal(filter, "", 0, 0)
}

// GetByDiscordBotID will load struct slice for one or more matching values of same parameter from database
// [ Parameters ]
//		DiscordBotIDToLoad = variadic, one or more table DiscordBotID to load from database, for example, enter 2, 3, 6, 9, will load records with DiscordBotID 2, 3, 6, and 9 from database to struct list
// [ Notes ]
// 		this function is based on the database column dbAttr:"GetBy" for suffix naming
func (l *DiscordBotChildList) GetByDiscordBotID(DiscordBotIDToLoad ...int64) error {
	// validate
	if l == nil {
		return errors.New("List Object Nil")
	}

	if DiscordBotIDToLoad == nil {
		return errors.New("Get Requires One or More Filter Values (Input Nil)")
	}

	if len(DiscordBotIDToLoad) == 0 {
		return errors.New("Get Requires One or More Filter Values (Input Count 0)")
	}

	// compose filter
	filter := ""

	for _, v := range DiscordBotIDToLoad {
		if util.LenTrim(filter) > 0 {
			filter += ", "
		}

		filter += util.Int64ToString(v)
	}

	if strings.Contains(filter, ",") {
		filter = "WHERE DiscordBotID IN (" + filter + ")"
	} else {
		filter = "WHERE DiscordBotID = " + filter
	}

	// perform action
	return l.getInternal(filter, "", 0, 0)
}

// GetCustom will load struct slice with custom filter, sort, offsetCount, limitCount from database records
//	filter = (optional) SQL WHERE clause portion of filter conditions (may start with WHERE or excluded)
//							if query needs to override FROM ... WHERE ..., then filter value begin with FROM to indicate override required (FROM AnotherTable WHERE ... or FROM AnotherTable INNER JOIN AnotherTable2 WHERE ...)
//	sort = (optional) SQL ORDER BY clause portion of sorting conditions (may start with ORDER BY or excluded)
//	offsetCount = Number of records left out before records returns
//	limitCount = limit number of records to return, 0 turns off limit
//	args = varidic parameters to match ordinal position of parameters denoted as ? within filter
func (l *DiscordBotChildList) GetCustom(filter string, sort string, offsetCount int, limitCount int, args ...interface{}) error {
	// validate
	if l == nil {
		return errors.New("List Object Nil")
	}

	// perform action
	return l.getInternal(filter, sort, offsetCount, limitCount, args...)
}

// getInternal is the internal helper that supports flexible get query and load data into list objects
// [ Parameters ]
//		filter = optional sql filter portion to be injected into query
//					if query needs to override FROM ... WHERE ..., then filter value begin with FROM to indicate override required (FROM AnotherTable WHERE ... or FROM AnotherTable INNER JOIN AnotherTable2 WHERE ...)
//		sort = optional sql sort portion to be injected into query
//		offsetCount = optional integer indicating how many rows to skip before applying limit count, 3 indicates first 3 rows skipped and then start applying limit count (0 = no offset)
//		limitCount = optional integer indicating the number of rows limited to for this action (0 = no limit)
//		args = optional ordinal parameter arguments to pass into the filter query portion (must appear in ordinal position of params defined within filter)
func (l *DiscordBotChildList) getInternal(filter string, sort string, offsetCount int, limitCount int, args ...interface{}) error {
	// validate
	if l == nil {
		return errors.New("List Object Nil")
	}

	// is this FROM override?
	isFromOverride := false

	if util.Left(strings.ToUpper(util.Trim(filter)), 5) == "FROM " {
		isFromOverride = true
	}

	// pre-process where clause
	filter = util.Trim(filter)

	if !isFromOverride {
		if util.LenTrim(filter) > 0 {
			if util.Left(strings.ToUpper(filter), 5) == "WHERE" {
				filter += " "
			} else {
				filter = "WHERE " + filter + " "
			}
		}
	}

	// pre-process sort clause
	sort = util.Trim(sort)
	if util.LenTrim(sort) > 0 {
		// has custom sort
		if util.Left(strings.ToUpper(sort), 8) != "ORDER BY" {
			sort = "ORDER BY " + sort
		}

		// remove; if exists
		if util.Right(sort, 1) == ";" {
			sort = util.Left(sort, len(sort)-1)
		}
	} else {
		// no custom sort
		sort = "ORDER BY ID ASC"
	}

	// pre-process limit clause
	limit := ""

	if limitCount > 0 {
		if offsetCount > 0 {
			limit = " LIMIT " + util.Itoa(offsetCount) + ", " + util.Itoa(limitCount)
		} else {
			limit = " LIMIT " + util.Itoa(limitCount)
		}
	}

	// first we will clear prior list objects and count
	l.List = nil
	l.Count = 0

	// compose query
	q := new(data.QueryBuilder)

	if !isFromOverride {
		q.Set("SELECT * FROM DiscordBotChild " + filter + sort + limit + ";")
	} else {
		q.Set("SELECT * " + filter + sort + limit + ";")
	}

	if args != nil {
		if len(args) > 0 {
			for _, v := range args {
				q.Ordinal(v)
			}
		}
	}

	// declare result
	var notFound bool
	var err error

	output := &[]DiscordBotChild{}

	// execute query
	var dbCurrent *data.MySql

	if !_DiscordBotChildListPrefersDBReader {
		dbCurrent = db
	} else {
		dbCurrent = getReaderDB()
	}

	if len(q.ParamsSlice()) > 0 {
		// has parameters
		notFound, err = dbCurrent.GetStructSlice(output, q.SQL(), q.ParamsSlice()...)
	} else {
		// no parameters
		notFound, err = dbCurrent.GetStructSlice(output, q.SQL())
	}
	dbCurrent = nil

	// evaluate execute response
	if err != nil {
		return errors.New("Get List Query Failed: " + err.Error())
	}

	if notFound {
		// if nothing is found, do not return error
		return nil
	}

	if len(*output) == 0 {
		// nothing found
		return nil
	}

	// at this point output contains loaded objects, assign into struct
	l.List = output
	l.Count = len(*output)

	// success
	return nil
}

// IsDataChanged checks if each struct's participating data in the struct slice has changed
// against its original value in the corresponding database row,
// returns pointer to struct slice that are changed,
// if no structs changed in the struct slice, then nil is returned
// [ Return Values ]
//		changedList = pointer to struct slice that have participating data changed vs database row's original values
func (l *DiscordBotChildList) IsDataChanged() (changedList *[]DiscordBotChild) {
	// validate
	if l == nil {
		// if not valid, then treat as no data changed
		return nil
	}

	// if count of list is zero, then treat as no change
	if l.Count <= 0 {
		// treat as no data changed
		return nil
	}

	// loop thru each object in list and check if data has changed
	// if data was changed, append that object into changed return list
	diffList := new([]DiscordBotChild)

	for _, v := range *l.List {
		if v.IsDataChanged() {
			// data was changed
			*diffList = append(*diffList, v)
		}
	}

	// evaluate result
	if len(*diffList) == 0 {
		// no changes
		return nil
	}

	// return changed list
	return diffList
}

// Validate checks if each struct fields' data in the struct slice conforms to the expected data integrity rules,
// such as required field is met, string fields size min max is met, numeric range is met,
// returns nil if validate is successful,
// otherwise returns first invalid struct reference, and error containing the validation failure reason
// [ Return Values ]
//		invalid = pointer to the first invalid struct that failed the Validate action, if all valid, then nil is returned
//		invalidInfo = string of invalid reason
//		err = if Validate failed, the validation failure reason
func (l *DiscordBotChildList) Validate() (invalid *DiscordBotChild, invalidInfo string, err error) {
	// validate
	if l == nil {
		// if not valid, then return error
		return nil, "", errors.New("List Object Nil")
	}

	// if count of list is zero, then treat as success
	if l.Count <= 0 {
		// treat as success
		return nil, "", nil
	}

	// loop thru each object in list and validate
	// return the first invalid object
	for _, v := range *l.List {
		if err = v.Validate(); err != nil {
			// validate failed
			return &v, err.Error(), nil
		}
	}

	// at this point, validate success
	return nil, "", nil
}

// Set persists the current struct slice fields into the underlying database table rows,
// a new database row will be inserted if such was not found,
// otherwise existing database row will be updated
func (l *DiscordBotChildList) Set() (failed *DiscordBotChild, err error) {
	// validate
	if l == nil {
		// if not valid, then return error
		return nil, errors.New("List Object Nil")
	}

	// if there are any to set
	if l.Count <= 0 {
		// return as success
		return nil, nil
	}

	// loop thru each object in list and perform set action
	for _, v := range *l.List {
		if err = v.Set(); err != nil {
			// failed
			return &v, err
		}
	}

	// all set success
	return nil, nil
}

// Delete removes the current struct slice records from the underlying database table rows
func (l *DiscordBotChildList) Delete() (failed *DiscordBotChild, err error) {
	// validate
	if l == nil {
		// if not valid, then return error
		return nil, errors.New("List Object Nil")
	}

	// if there are any to delete
	if l.Count <= 0 {
		// return as success
		return nil, nil
	}

	// loop thru each object in list and perform delete action
	for _, v := range *l.List {
		if err = v.Delete(); err != nil {
			// failed
			return &v, err
		}
	}

	// all delete success
	return nil, nil
}

// FindByID will search existing struct slice for one or more matching values of same parameter,
// it will return a list of found objects, or error if encountered
// [ Parameters ]
//		IDToFind = variadic, one or more table ID to find, for example, enter 2, 3, 6, 9, will find records with ID 2, 3, 6, and 9 in list
// [ Notes ]
//		this function is based on the database column dbAttr:"GetBy" for suffix naming
func (l *DiscordBotChildList) FindByID(IDToFind ...int64) (foundList *[]DiscordBotChild, err error) {
	// validate
	if l == nil {
		// if not valid, then return error
		return nil, errors.New("List Object Nil")
	}

	// check if there are any objects
	if l.Count <= 0 {
		// return as nothing found, and its not an error
		return nil, nil
	}

	// check parameters
	if len(IDToFind) <= 0 {
		// return as nothing found, and its not an error
		return nil, nil
	}

	// initialize foundList
	foundList = new([]DiscordBotChild)

	// loop thru list to match
	for _, seek := range IDToFind {
		if seek > 0 {
			for _, v := range *l.List {
				if v.ID != 0 {
					if v.ID == seek {
						// found match
						*foundList = append(*foundList, v)
					}
				}
			}
		}
	}

	// find is complete
	if len(*foundList) == 0 {
		// nothing is found
		return nil, nil
	} else {
		// one or more results found
		return foundList, nil
	}
}

// FindByDiscordBotID will search existing struct slice for one or more matching values of same parameter,
// it will return a list of found objects, or error if encountered
// [ Parameters ]
//		DiscordBotIDToFind = variadic, one or more table DiscordBotID to find, for example, enter 2, 3, 6, 9, will find records with DiscordBotID 2, 3, 6, and 9 in list
// [ Notes ]
//		this function is based on the database column dbAttr:"GetBy" for suffix naming
func (l *DiscordBotChildList) FindByDiscordBotID(DiscordBotIDToFind ...int64) (foundList *[]DiscordBotChild, err error) {
	// validate
	if l == nil {
		// if not valid, then return error
		return nil, errors.New("List Object Nil")
	}

	// check if there are any objects
	if l.Count <= 0 {
		// return as nothing found, and its not an error
		return nil, nil
	}

	// check parameters
	if len(DiscordBotIDToFind) <= 0 {
		// return as nothing found, and its not an error
		return nil, nil
	}

	// initialize foundList
	foundList = new([]DiscordBotChild)

	// loop thru list to match
	for _, seek := range DiscordBotIDToFind {
		if seek > 0 {
			for _, v := range *l.List {
				if v.DiscordBotID != 0 {
					if v.DiscordBotID == seek {
						// found match
						*foundList = append(*foundList, v)
					}
				}
			}
		}
	}

	// find is complete
	if len(*foundList) == 0 {
		// nothing is found
		return nil, nil
	} else {
		// one or more results found
		return foundList, nil
	}
}
