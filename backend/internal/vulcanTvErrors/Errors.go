package vulcanTvErrors

import "errors"

var ErrAlreadyExists = errors.New("already exists")
var NoSuchPageContent = errors.New("no such pagecontent")