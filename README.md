# BERROR
Bearaujus Error (berror) is an error tools that can help you to handle errors even more easier!

# Example: Error Merger
```go
package main

import (
	"fmt"

	"github.com/Bearaujus/berror"
)

func main() {
	var errMerger = berror.NewErrorMerger()

	for i := 0; i < 10; i++ {
		errMerger.AddIfNotNil(validateIsOddNum(i))
	}

	if errMerger.HasError() {
		fmt.Println(errMerger.GetError())
		return
	}

	fmt.Println("hi")
}

func validateIsOddNum(num int) error {
	if num%2 == 0 {
		return fmt.Errorf("validate: '%v' is not odd number", num)
	}

	return nil
}

```

# Example: Error Wrapper
```go
package main

import (
	"fmt"

	"github.com/Bearaujus/berror"
)

var (
	ErrValidateNumIsZero = berror.NewErrorWrapper("num cannot zero")
	ErrValidateIsOddNum  = berror.NewErrorWrapper("'%v' is not odd number")
)

func main() {
	for i := 0; i < 10; i++ {
		if err := validateIsOddNum(i); err != nil {
			fmt.Println(err.Error())
		}
	}
}

func validateIsOddNum(num int) error {
	if num == 0 {
		return ErrValidateNumIsZero.New()
	}
	
	if num%2 == 0 {
		return ErrValidateIsOddNum.New(num)
	}

	return nil
}

```

# Example: Exit If Error Handler
```go
package main

import (
	"errors"
	"fmt"

	"github.com/Bearaujus/berror"
)

func main() {
	var eiErrHandler = berror.NewexitIfErrorHandler(exampleFuncWhenError1, exampleFuncWhenError2)

	var err error
	eiErrHandler.ExitIfError(err)
	fmt.Println("hi")

	err = errors.New("an error")
	eiErrHandler.ExitIfError(err)
	fmt.Println("hi2")
}

func exampleFuncWhenError1() error {
	fmt.Println("you just got an error, this func 1 is triggered")
	return nil
}

func exampleFuncWhenError2() error {
	fmt.Println("you just got an error, this func 2 is triggered")
	return nil
}

```

# Example: Basic
```go
package main

import (
	"errors"
	"fmt"

	"github.com/Bearaujus/berror"
)

func main() {
	var err = errors.New("an error")
	if berror.IsError(err) {
		fmt.Println(err.Error())
	}

	berror.ExitIfError(err)
	fmt.Println("hi")
}

```