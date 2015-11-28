package cart

import (
	"github.com/gin-gonic/contrib/sessions"
	"encoding/json"
)

type GinGonicSession struct {
	Session sessions.Session
}

func (gcs GinGonicSession) Restore() (map[string]*CartItem, error) {
	
	var list map[string]*CartItem

	session := gcs.Session
	data := session.Get("__meta_gin_cart")

	if data == nil {

		list = make(map[string]*CartItem)
		
		return list, nil
	} else {

		encoded := data.(string)

		if err := json.Unmarshal([]byte(encoded), &list); err != nil {
	        
	        return list, err
	    }

	    return list, nil
	}
} 

func (gcs GinGonicSession) Save(data map[string]*CartItem) error {
	
	encoded, err := json.Marshal(data)

	if err != nil {
		return err
	}

	session := gcs.Session

	session.Set("__meta_gin_cart", string(encoded))
	session.Save()

	return nil
} 