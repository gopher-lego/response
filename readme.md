# response

## example

```
import (
	"github.com/gopher-lego/response"
	"github.com/gin-gonic/gin"
)

func main() {

	r := gin.Default()

	r.GET("/ping", func(c *gin.Context) {

		response.Success(c, "any")

	})

	r.Run()
}
```
