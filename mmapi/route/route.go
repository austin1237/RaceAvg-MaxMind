package route

import (
	"github.com/labstack/echo"
	"github.com/user/mmapi/handler"
)

func SetUpRoutes(e *echo.Echo) {
	/**
	* @api {get} /v1/location GET Location
	* @apiVersion 1.0.0
	* @apiDescription get the location of an ip address
	* @apiName GetLocation
	* @apiGroup Location
	* @apiParam {string} ip ipv4 address
	* @apiExample {curl} Example usage:
	*     curl -i http://localhost:[port]/v1/location?ip="174.29.5.118"
	* @apiSuccessExample {json} Success-Response:
	*     HTTP/1.1 200 OK
	*     {
	* 	   		"locationId": 14288,
	* 			"country": "US",
	* 			"region": "CO",
	* 			"city": "Arvada",
	* 			"postalCode": "",
	* 			"latitude": 39.8131,
	* 			"longitude": -105.1257,
	* 			"metroCode": "751",
	* 			"areaCode": ""
	*	  }
	*
	* @apiErrorExample {json} Error-Response:
	*     HTTP/1.1 400 Bad Request
	*     {
	*       "message": "error text"
	*     }
	 */
	e.GET("v1/location", handler.GetLocations)

	/**
	 * @api {get} /v1/health GET Health
	 * @apiVersion 1.0.0
	 * @apiDescription Is the service healthy
	 * @apiName GetHealth
	 * @apiGroup Health
	 * @apiSuccessExample {json} Success-Response:
	 *     HTTP/1.1 200 OK
	 */
	e.GET("v1/health", handler.GetHealth)
}
