package customError

import (
	"fmt"
	"strings"

	"github.com/friendsofgo/errors"
)

//WrapFlag custom error wrap with flag
// example result : [<flag>] <args> : <err>
func WrapFlag(err error, flag string, args ...string) error {
	if flag = strings.TrimSpace(flag); flag != "" {
		flag = fmt.Sprintf("[%s] ", strings.ToUpper(flag))
	}

	msg := strings.Join(args, " | ")
	msg = flag + msg

	if err == nil {
		err = fmt.Errorf("[FOR DEVELOPER] forget to set error in the: %s", flag)
	}

	return errors.Wrap(err, msg)
}

//Wrap custom error wrap without flag
// example result : <args> : <err>
func Wrap(err error, args ...string) error {
	msg := strings.Join(args, " | ")
	if err == nil {
		err = fmt.Errorf("[FOR DEVELOPER] forget to set error in the: %s", msg)
	}

	return errors.Wrap(err, msg)
}

//Cause get cause error
func Cause(err error) error {
	if err == nil {
		file, line, funName := trace()
		return fmt.Errorf("[FILE]: %s | [FUNCTION NAME] %s | [LINE] %d", file, funName, line)
	}

	return errors.Cause(err)
}
