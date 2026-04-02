# Repository

Repository is for DB/API/external calls.

## Guidelines

- Keep functions thin.
- No business formula here.
- Return data/errors to usecase or logic.

## Example

```go
package transaction

import (
	"database/sql"

	"github.com/jekiapp/hi-mod-arch/internal/model"
)

func SelectCartByUserID(db *sql.DB, userID int64) (model.CartData, error) {
	data := model.CartData{}
	query := "SELECT * from cart WHERE user_id=$1"
	rows, err := db.Query(query, userID)
	if err != nil {
		return data, err
	}
	_ = rows
	return data, nil
}
```
