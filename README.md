# Import

```go
import formatter "github.com/ivanauliaa/response-formatter"
```

# Usage

**gin Example**

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

**echo Example**

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
