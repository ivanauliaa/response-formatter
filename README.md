# Import

```go
import formatter "github.com/ivanauliaa/response-formatter"
```

# Usage

Functions
```go
- ResponseFormatter(status int32, message string, data interface{})
- BadRequestResponse(data interface{})
- NotFoundResponse(data interface{})
- UnauthorizedResponse(data interface{})
- InternalServerErrorResponse(data interface{})
- SuccessResponse(data interface{})
```

Params
- `status`: an `int32` which represent HTTP status code. You can pass with int literal, but I prefer using net/http HTTP status constants.
- `message`: a `string` which represent response message which is success or fail.
- `data`: an `interface{}` which represent requested data from client or detailed error messages. You can pass with either `map[string]interface{}` or `struct` with exported JSON tagged properties data.
# Example

**gin**

```go
func Redirect(c *gin.Context) {
	collection := database.Connect()
	urlCode := c.Param("code")

	result := &model.URLDoc{}
	err := collection.FindOne(utils.GLOBAL_CONTEXT, bson.M{"urlCode": urlCode}).Decode(result)
	if err != nil {
		if err == mongo.ErrNoDocuments {
			c.JSON(http.StatusBadRequest, formatter.BadRequestResponse({"error": fmt.Sprintf("No URL with code: %s", urlCode)}))
			return
		} else {
			c.JSON(http.StatusInternalServerError, formatter.BadRequestResponse({"error": err.Error()}))
			return
		}
	}

	longURL := result.LongURL
	c.Redirect(http.StatusPermanentRedirect, longURL)
}

```

**echo**

```go
func DeleteUserController(c echo.Context) error {
	id, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		return c.JSON(http.StatusBadRequest, formatter.BadRequestResponse(nil))
	}

	for index, user := range users {
		if user.Id == id {
			users = append(users[:index], users[index+1:]...)
			return c.JSON(http.StatusOK, formatter.SuccessResponse(user))
		}
	}

	return c.JSON(http.StatusNotFound, formatter.NotFoundResponse(nil))
}
```
