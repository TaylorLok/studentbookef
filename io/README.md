# ostmfe
This package is responsible for input and output communication
from front end and backend. 

example:
_this method takes an object of type Book that is located in domain package and returns same object and an error_
In GO a method can have multiple returns like in this case it returns an object and a error

func CreateBook(book domain.Book)(domain.Book,error){ 
_we first declare an empty object of type Book (in JAVA you could have wrote: private Book book:)_

entity:=domain.Book{}

_now we sending the object to the back end through the API address. we expecting a response that we will store in resp_
_by sending the object to the backend we specify the type of HTTP method that the backend is expecting_


resp, _ := api.Rest().SetBody(mybook).Post(bookURL + "create")

_Now we checking if resp received an error when sending the object. If yes we will return an empty book object(entity)_


if resp.IsError() {
		return entity, errors.New(resp.Status())
	}
	
_In case if resp didn't meet any interruption. we need to format the repose that is coming from the backend in JSON format _


err := api.JSON.Unmarshal(resp.Body(), &entity)


_Now we checking if the Unmarshaling has successfully passed without an obstacle_
_if an error has occurred we will return an empty object and error status_


if err != nil {
		return entity, errors.New(resp.Status())
	}


_If all went well we will return the result that came from the frontend_

}

