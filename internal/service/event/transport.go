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
				"Events": "No se Econtro ningun Evento por parametro",
			})
		}

		result := s.FindByID(ID)

		if result != nil {
			c.JSON(http.StatusBadRequest, gin.H{
				"Events": "No se Econtro ningun Evento en la Base de Datos",
			})
		}

		c.JSON(http.StatusOK, gin.H{
			"Events": *result,
		})
	}
}

// add ...
func add(s Service) gin.HandlerFunc {
	return func(c *gin.Context) {
		var data Event
		c.BindJSON(&data)
		result, err := s.AddEvent(data)
		if result >= 0 {
			c.JSON(http.StatusOK, gin.H{
				"state": result,
				// "result": event,
				"data": data,
			})
		} else {
			c.JSON(http.StatusBadRequest, gin.H{
				"error": err,
			})
		}

	}
}

// delete ...
func delete(s Service) gin.HandlerFunc {
	return nil
}

// Register ...
func (s httpService) Register(r *gin.Engine) {
	for _, e := range s.endpoints {
		r.Handle(e.method, e.path, e.function)
	}
}
