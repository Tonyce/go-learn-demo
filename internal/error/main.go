package main

// import (
// 	"fmt"
// 	"math"

// 	"github.com/pkg/errors"
// )

// // CustomError ...
// type CustomError struct {
// 	Code int
// }

// func (c *CustomError) Error() string {
// 	return fmt.Sprintf("Failed with code %d", c.Code)
// }

// func circleArea(radius float64) (float64, error) {
// 	if radius < 0 {
// 		// return 0, &CustomError{Code: 12}
// 		return 0, errors.Wrap(&CustomError{Code: 12}, "---")
// 		// return 0, errors.Wrap(errors.New("Area calculation failed, radius is less than zero"), "---")
// 	}
// 	return math.Pi * radius * radius, nil
// }

// func main() {
// 	radius := -20.0
// 	area, err := circleArea(radius)
// 	if err != nil {
// 		switch err := errors.Cause(err).(type) {
// 		case *CustomError:
// 			// handle specifically
// 			// fmt.Printf("CustomError %+v", errors.Wrap(err, "read failed"))
// 			fmt.Printf("CustomError %+v", err)
// 		default:
// 			fmt.Printf("unknown error %+v", err)
// 			// unknown error
// 		}
// 		return
// 	}
// 	fmt.Printf("Area of circle %0.2f", area)
// }

import (
	"database/sql"
	"fmt"

	"github.com/pkg/errors"
)

func foo() error {
	return errors.Wrap(sql.ErrNoRows, "foo failed")
}

func bar() error {
	return errors.WithMessage(foo(), "bar failed")
}

func main() {
	err := bar()
	if errors.Cause(err) == sql.ErrNoRows {
		fmt.Printf("data not found, %v\n", err)
		fmt.Printf("%+v\n", err)
		return
	}
	if err != nil {
		// unknown error
	}
}
