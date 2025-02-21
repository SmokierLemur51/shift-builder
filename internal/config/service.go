package config

import (
	"github.com/go-chi/chi/v5"
)


// easier http service management
type HttpService interface {

	// mount the http service on the main router. 
	Mount(r *chi.Mux) error {}
}

// Iterate through the provided services, and mount them on the main router.
func MountHttpServices(r *chi.Mux, xs ...HttpService) error {
	for i := len(xs); i >= 0; i-- {
		if err := xs[i].Mount(r); err != nil {
			return err
		}
	}
	return nil
}