package internal

import (
	"errors"
	util "github.com/aldelo/common"
	data "github.com/aldelo/common/wrapper/mysql"
	"strings"
)

// ---------------------------------------------------------------------------------------------------------------------
// LiveFeed Actions
// ---------------------------------------------------------------------------------------------------------------------

// indicates if LiveFeed struct will prefer dbReader over dbWriter where available
var _LiveFeedPrefersDBReader bool

// UseDBReaderPreferred sets indicator that this struct prefers to use db reader where applicable and if available
func (r LiveFeed) UseDBReaderPreferred() {
	// set db reader usage preference to true, so if db reader is available, we will use it if code is applicable for it
	_LiveFeedPrefersDBReader = true
}

// UseDBWriterPreferred sets indicator that this struct prefers to use db writer (this is the default preference)
func (r LiveFeed) UseDBWriterPreferred() {
	// set db writer usage preference to true, so this struct will use the db writer for operations
	_LiveFeedPrefersDBReader = false
}

// Fill is a helper function to fill in required or important fields of this struct,
// this function consolidates all required or important fields into one input path to simply data input,
// note: values filled into this function is not yet persisted to database, to persist, must call Set() function
// note: values filled into this function is not validated at this point yet, call Validate() function to validate if desired
// note: this function will have struct tag ggAttr:"8 = Fill" as fields to include as parameters
// [ Parameters ]
//		UserID int64 = REQUIRED
//		YouTubeLiveLink string = REQUIRED
//		Location string = REQUIRED
func (r *LiveFeed) Fill(UserID int64, YouTubeLiveLink string, Location string) {
	r.UserID = UserID
	r.YouTubeLiveLink = YouTubeLiveLink
	r.Location = Location
}

// GetUserID retrieves struct field value,
// if any field data type is int representing enum, then the corresponding enum is returned,
// if any field data type is sql null type or pointer, then the null or pointer is converted to base type
func (r *LiveFeed) GetUserID() int64 {
	return r.UserID
}

// GetYouTubeLiveLink retrieves struct field value,
// if any field data type is int representing enum, then the corresponding enum is returned,
// if any field data type is sql null type or pointer, then the null or pointer is converted to base type
func (r *LiveFeed) GetYouTubeLiveLink() string {
	return r.YouTubeLiveLink
}

// GetLocation retrieves struct field value,
// if any field data type is int representing enum, then the corresponding enum is returned,
// if any field data type is sql null type or pointer, then the null or pointer is converted to base type
func (r *LiveFeed) GetLocation() string {
	return r.Location
}

// SetUserID sets value into struct field value and checks for data input initial validation rules,
// if any field data type is int representing enum, then the corresponding enum is used as input parameter,
// if any field data type is sql null type or pointer, then the base type is used as parameter and converted into null or pointer
//		[ Parameters ]
//		v int64 = Struct internal data type: int64; REQUIRED
func (r *LiveFeed) SetUserID(v int64) error {
	r.UserID = v

	return nil
}

// SetYouTubeLiveLink sets value into struct field value and checks for data input initial validation rules,
// if any field data type is int representing enum, then the corresponding enum is used as input parameter,
// if any field data type is sql null type or pointer, then the base type is used as parameter and converted into null or pointer
//		[ Parameters ]
//		v string = Struct internal data type: string; REQUIRED
func (r *LiveFeed) SetYouTubeLiveLink(v string) error {
	if util.LenTrim(v) == 0 {
		return errors.New("SetYouTubeLiveLink Failed, A Text Length Greater Than Zero is Required")
	}

	r.YouTubeLiveLink = v

	return nil
}

// SetLocation sets value into struct field value and checks for data input initial validation rules,
// if any field data type is int representing enum, then the corresponding enum is used as input parameter,
// if any field data type is sql null type or pointer, then the base type is used as parameter and converted into null or pointer
//		[ Parameters ]
//		v string = Struct internal data type: string; REQUIRED
func (r *LiveFeed) SetLocation(v string) error {
	if util.LenTrim(v) == 0 {
		return errors.New("SetLocation Failed, A Text Length Greater Than Zero is Required")
	}

	r.Location = v

	return nil
}

// GetByID retrieves a row from database using the ID field value,
// and marshals found result fields into struct,
// an error of nil indicates success
// note: this function will have struct tag ggAttr:"9 = GetBy" for columns acting as its source
func (r *LiveFeed) GetByID(ID int64) (notFound bool, err error) {
	// clean up
	r._originalValue = nil

	// compose query
	q := new(data.QueryBuilder)

	q.Set("SELECT * FROM LiveFeed WHERE ID=? LIMIT 1;")
	q.Ordinal(ID)

	// execute query
	var dbCurrent *data.MySql

	if !_LiveFeedPrefersDBReader {
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

// GetByUserID retrieves a row from database using the UserID field value,
// and marshals found result fields into struct,
// an error of nil indicates success
// note: this function will have struct tag ggAttr:"9 = GetBy" for columns acting as its source
func (r *LiveFeed) GetByUserID(UserID int64) (notFound bool, err error) {
	// clean up
	r._originalValue = nil

	// compose query
	q := new(data.QueryBuilder)

	q.Set("SELECT * FROM LiveFeed WHERE UserID=? LIMIT 1;")
	q.Ordinal(UserID)

	// execute query
	var dbCurrent *data.MySql

	if !_LiveFeedPrefersDBReader {
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
func (r *LiveFeed) GetScalar(scalarColumnName string, filter string, sort string, args ...interface{}) (result string, notFound bool, err error) {
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
		q.Set("SELECT " + scalarColumnName + " FROM LiveFeed " + filter + sort + " LIMIT 1;")
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

	if !_LiveFeedPrefersDBReader {
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
func (r *LiveFeed) IsDataChanged() bool {
	if r._originalValue == nil {
		// default to data changed status
		return true
	}

	// if primary key is 0 then this is new record, assume data changed
	if r.ID <= 0 {
		return true
	}

	// assertion of original value to struct for old value
	old := r._originalValue.(LiveFeed)

	// check if old value vs current value was changed
	if r.UserID != old.UserID {
		return true
	}

	if r.YouTubeLiveLink != old.YouTubeLiveLink {
		return true
	}

	if r.Location != old.Location {
		return true
	}

	// no changes
	return false
}

// Validate checks if the struct fields' data conforms to the expected data integrity rules,
// such as required field is met, string fields size min max is met, numeric range is met,
// returns nil is validate is successful, otherwise returns error containing the validation failure reason
func (r *LiveFeed) Validate() error {
	// check for required fields
	// based on ggAttr:"3 = - (required)" and ggAttr:"6 = Set-Upd, Set-InsUpd, Set-Ins"
	if r.UserID == 0 {
		return errors.New("UserID is Required")
	}

	if util.LenTrim(r.YouTubeLiveLink) == 0 {
		return errors.New("YouTubeLiveLink is Required")
	}

	if util.LenTrim(r.Location) == 0 {
		return errors.New("Location is Required")
	}

	// success
	return nil
}

// Set persists the current struct fields into the underlying database table row,
// a new database row will be inserted if such was not found,
// otherwise existing database row will be updated
// note: this function will use struct tag ggAttr:"1 = PK-Custom"
//       and ggAttr:"6 = Set-Ins, Set-Upd, Set-InsUpd" as participating columns
func (r *LiveFeed) Set() error {
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
		q.Set("INSERT INTO LiveFeed ")
		q.Set("(UserID, YouTubeLiveLink, Location) ")
		q.Set("VALUES ")
		q.Set("(:UserID, :YouTubeLiveLink, :Location);")

		q.Named("UserID", r.UserID)
		q.Named("YouTubeLiveLink", r.YouTubeLiveLink)
		q.Named("Location", r.Location)
	} else {
		// update
		isNewRow = false

		// compose update action query
		q.Set("UPDATE LiveFeed ")
		q.Set("SET UserID=:UserID, YouTubeLiveLink=:YouTubeLiveLink, Location=:Location ")
		q.Set("WHERE ID=:ID;")

		q.Named("ID", r.ID)
		q.Named("UserID", r.UserID)
		q.Named("YouTubeLiveLink", r.YouTubeLiveLink)
		q.Named("Location", r.Location)
	}

	// begin transaction
	var tx *data.MySqlTransaction
	var e error

	if tx, e = db.Begin(); e != nil {
		return errors.New("LiveFeed Persist Data Failed at BeginTran: " + e.Error())
	}

	// execute query
	result := tx.ExecByNamedMapParam(q.SQL(), q.ParamsMap())

	if result.Err != nil {
		if err1 := tx.Rollback(); err1 != nil {
			return errors.New("LiveFeed Persist Data Failed at RollBack 1: " + err1.Error() + " (Orig Error: " + result.Err.Error() + ")")
		} else {
			return errors.New("LiveFeed Persist Data Failed: " + result.Err.Error())
		}
	}

	if !isNewRow {
		// was update
		if result.RowsAffected <= 0 {
			// update failure
			if err1 := tx.Rollback(); err1 != nil {
				return errors.New("LiveFeed Persist Data Failed at RollBack 2: " + err1.Error() + " (Orig Error: No Rows Affected)")
			} else {
				return errors.New("LiveFeed Persist Updated Data Failed: No Rows Affected")
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
				return errors.New("LiveFeed Persist Data Failed at RollBack 3: " + err1.Error() + " (Orig Error: No New Primary ID Generated)")
			} else {
				return errors.New("LiveFeed Persist Inserted Data Failed: No New Primary ID Generated")
			}
		}
	}

	r._originalValue = nil
	r._originalValue = *r

	// commit persist data action
	if err1 := tx.Commit(); err1 != nil {
		return errors.New("LiveFeed Persist Data Failed at Commit: " + err1.Error())
	}

	// return nil as success
	return nil
}

// physicalDelete internal helper
func (r *LiveFeed) physicalDelete() error {
	// compose query
	q := new(data.QueryBuilder)

	q.Set("DELETE FROM LiveFeed WHERE ID=?;")
	q.Ordinal(r.ID)

	// execute query
	result := db.ExecByOrdinalParams(q.SQL(), q.ParamsSlice()...)

	if result.Err != nil {
		return errors.New("LiveFeed Physical Delete Failed: " + result.Err.Error())
	}

	if result.RowsAffected <= 0 {
		return errors.New("LiveFeed Physical Delete Failed: No Rows Affected")
	}

	// physical delete successful
	return nil
}

// Delete removes the current struct record from the underlying database table row
func (r *LiveFeed) Delete() error {
	// validate
	if r.ID <= 0 {
		return errors.New("LiveFeed Delete Data Failed: Row Primary ID is Required")
	}

	if r._originalValue == nil {
		return errors.New("LiveFeed Delete Data Failed: Row Must Be Loaded First")
	}

	// perform delete action
	// physical delete
	return r.physicalDelete()
}

// ---------------------------------------------------------------------------------------------------------------------
// LiveFeedList Actions
// ---------------------------------------------------------------------------------------------------------------------

// LiveFeedList collection level dal model
type LiveFeedList struct {
	List  *[]LiveFeed
	Count int
}

// indicates if LiveFeed list will prefer dbReader over dbWriter where available
var _LiveFeedListPrefersDBReader bool

// UseDBReaderPreferred sets indicator that this struct prefers to use db reader where applicable and if available
func (l LiveFeedList) UseDBReaderPreferred() {
	// set db reader usage preference to true, so if db reader is available, we will use it if code is applicable for it
	_LiveFeedListPrefersDBReader = true
}

// UseDBWriterPreferred sets indicator that this struct prefers to use db writer (this is the default preference)
func (l LiveFeedList) UseDBWriterPreferred() {
	// set db writer usage preference to true, so this struct list will use the db writer for operations
	_LiveFeedListPrefersDBReader = false
}

// Element gets one of the slice elements in list, identified by ordinal index position
func (l *LiveFeedList) Element(position int) (*LiveFeed, error) {
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
func (l *LiveFeedList) GetAll(limitCount int, offsetCount ...int) error {
	offset := util.GetFirstIntOrDefault(0, offsetCount...)
	return l.getInternal("", "", offset, limitCount)
}

// GetByID will load struct slice for one or more matching values of same parameter from database
// [ Parameters ]
//		IDToLoad = variadic, one or more table ID to load from database, for example, enter 2, 3, 6, 9, will load records with ID 2, 3, 6, and 9 from database to struct list
// [ Notes ]
// 		this function is based on the database column dbAttr:"GetBy" for suffix naming
func (l *LiveFeedList) GetByID(IDToLoad ...int64) error {
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

// GetByUserID will load struct slice for one or more matching values of same parameter from database
// [ Parameters ]
//		UserIDToLoad = variadic, one or more table UserID to load from database, for example, enter 2, 3, 6, 9, will load records with UserID 2, 3, 6, and 9 from database to struct list
// [ Notes ]
// 		this function is based on the database column dbAttr:"GetBy" for suffix naming
func (l *LiveFeedList) GetByUserID(UserIDToLoad ...int64) error {
	// validate
	if l == nil {
		return errors.New("List Object Nil")
	}

	if UserIDToLoad == nil {
		return errors.New("Get Requires One or More Filter Values (Input Nil)")
	}

	if len(UserIDToLoad) == 0 {
		return errors.New("Get Requires One or More Filter Values (Input Count 0)")
	}

	// compose filter
	filter := ""

	for _, v := range UserIDToLoad {
		if util.LenTrim(filter) > 0 {
			filter += ", "
		}

		filter += util.Int64ToString(v)
	}

	if strings.Contains(filter, ",") {
		filter = "WHERE UserID IN (" + filter + ")"
	} else {
		filter = "WHERE UserID = " + filter
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
func (l *LiveFeedList) GetCustom(filter string, sort string, offsetCount int, limitCount int, args ...interface{}) error {
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
func (l *LiveFeedList) getInternal(filter string, sort string, offsetCount int, limitCount int, args ...interface{}) error {
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
		q.Set("SELECT * FROM LiveFeed " + filter + sort + limit + ";")
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

	output := &[]LiveFeed{}

	// execute query
	var dbCurrent *data.MySql

	if !_LiveFeedListPrefersDBReader {
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
func (l *LiveFeedList) IsDataChanged() (changedList *[]LiveFeed) {
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
	diffList := new([]LiveFeed)

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
func (l *LiveFeedList) Validate() (invalid *LiveFeed, invalidInfo string, err error) {
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
func (l *LiveFeedList) Set() (failed *LiveFeed, err error) {
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
func (l *LiveFeedList) Delete() (failed *LiveFeed, err error) {
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
func (l *LiveFeedList) FindByID(IDToFind ...int64) (foundList *[]LiveFeed, err error) {
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
	foundList = new([]LiveFeed)

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

// FindByUserID will search existing struct slice for one or more matching values of same parameter,
// it will return a list of found objects, or error if encountered
// [ Parameters ]
//		UserIDToFind = variadic, one or more table UserID to find, for example, enter 2, 3, 6, 9, will find records with UserID 2, 3, 6, and 9 in list
// [ Notes ]
//		this function is based on the database column dbAttr:"GetBy" for suffix naming
func (l *LiveFeedList) FindByUserID(UserIDToFind ...int64) (foundList *[]LiveFeed, err error) {
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
	if len(UserIDToFind) <= 0 {
		// return as nothing found, and its not an error
		return nil, nil
	}

	// initialize foundList
	foundList = new([]LiveFeed)

	// loop thru list to match
	for _, seek := range UserIDToFind {
		if seek > 0 {
			for _, v := range *l.List {
				if v.UserID != 0 {
					if v.UserID == seek {
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
