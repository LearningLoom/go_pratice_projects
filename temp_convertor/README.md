# Temperature convertor 

Steps to install a third party liberary 
- before doing anything first initilize your project with ``` go mod init <project name>``` 
- then install the required package (which in our case is github.com/syscll/tempconv) by ``` go get github.com/syscll/tempconv```
- then import that package and use it in your code 
```go
import (
	"fmt"
	"strings"

	"github.com/syscll/tempconv"
)

```