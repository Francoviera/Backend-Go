package event

import (
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
)

// HTTPService ...
type HTTPService interface {
	Register(*gin.Engine)
}

// endpoint ...
type endpoint struct {
	method   string
	path     string
	function gin.HandlerFunc
}

// httpService ...
type httpService struct {
	endpoints []*endpoint
}

// NewHTTPTransport ...
func NewHTTPTransport(s Service) HTTPService {
	endpoints := makeEndpoints(s)
	return httpService{endpoints}
}

// makeEndpoints ...
func makeEndpoints(s Service) []*endpoint {
	list := []*endpoint{}

	list = append(list, &endpoint{
		method:   "GET",
		path:     "/events",
		function: getAll(s),
	})

	list = append(list, &endpoint{
		method:   "GET",
		path:     "/event/:id",
		function: getByID(s),
	})
	list = append(list, &endpoint{
		method:   "POST",
		path:     "/event",
		function: add(s),
	})
	list = append(list, &endpoint{
		method:   "DELETE",
		path:     "/event/:id",
		function: delete(s),
	})
	list = append(list, &endpoint{
		method:   "PUT",
		path:     "/event/:id",
		function: put(s),
	})

	return list
}

// getAll ...
func getAll(s Service) gin.HandlerFunc {
	return func(c *gin.Context) {
		c.JSON(http.StatusOK, gin.H{
			"events": s.FindAll(),
		})
	}
}

// getByID ...
func getByID(s Service) gin.HandlerFunc {
	// var httpErrorMsg *ErrorStruct

	return func(c *gin.Context) {
		ID, err := strconv.Atoi(c.Param("id"))
		if err != nil {
			c.JSON(http.StatusBadRequest, gin.H{
				"Message": "No se Encontro ningun id de Eventos por parametro",
			})
		}

		result, err := s.FindByID(ID)

		if err != nil {
			c.JSON(http.StatusBadRequest, gin.H{
				"Message": "No se Encontro ningun Evento en la Base de Datos",
			})
		}

		c.JSON(http.StatusOK, gin.H{
			"Event": *result,
		})
	}
}

// add ...
func add(s Service) gin.HandlerFunc {
	return func(c *gin.Context) {
		var data Event
		c.BindJSON(&data)
		result, err := s.AddEvent(data)
		if err != nil {
			c.JSON(http.StatusBadRequest, gin.H{
				"error": err,
			})
		} else {
			data.ID = result
			c.JSON(http.StatusOK, gin.H{
				"Message":  "Se agrego exitosamente",
				"new-data": data,
			})
		}

	}
}

// delete ...
func delete(s Service) gin.HandlerFunc {
	return func(c *gin.Context) {
		ID, _ := strconv.Atoi(c.Param("id"))

		err := s.Delete(ID)

		if err != nil {
			c.JSON(http.StatusBadRequest, gin.H{
				"error": err,
			})
		} else {
			c.JSON(http.StatusBadRequest, gin.H{
				"Message": "Se elimino exitosamente",
			})
		}

	}
}

// put ...
func put(s Service) gin.HandlerFunc {
	return func(c *gin.Context) {
		var data Event
		c.BindJSON(&data)
		ID, _ := strconv.Atoi(c.Param("id"))
		err := s.Put(data, ID)
		if err != nil {
			c.JSON(http.StatusOK, gin.H{
				"error": err,
			})
		} else {
			c.JSON(http.StatusBadRequest, gin.H{
				"Message":  "Se mofico exitosamente",
				"new-data": data,
			})
		}

	}
}

// Register ...
func (s httpService) Register(r *gin.Engine) {
	for _, e := range s.endpoints {
		r.Handle(e.method, e.path, e.function)
	}
}
