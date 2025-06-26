package server

import "fmt"

// FetchConfigValues gets the required configuration values from
// environment variables.
// If a required environment variable is missing it returns an error.
func (app *Application) FetchConfigValues(requiredEnvVariables []string) error {
	for _, key := range requiredEnvVariables {
		app.config.BindEnv(key)
	}

	missingEnvVariables := make([]string, 0)
	for _, key := range requiredEnvVariables {
		if !app.config.IsSet(key) {
			missingEnvVariables = append(missingEnvVariables, key)
		}
	}

	if len(missingEnvVariables) != 0 {
		return fmt.Errorf("some required env. variable(s) were not properly set: %v. Check your .envrc file", missingEnvVariables)
	}

	return nil
}

// GetConfigValueString checks if the key has been set. If not (either the key was misspelled
// or the env. variable is unset), it returns an error.
// Use this function to safely get config values.
func (app *Application) GetConfigValueString(key string) (string, error) {
	if !app.config.IsSet(key) {
		return "", fmt.Errorf("env. configuration value (`%s`) is not set", key)
	}
	return app.config.GetString(key), nil
}

func (app *Application) GetConfigValueBool(key string) (bool, error) {
	if !app.config.IsSet(key) {
		return false, fmt.Errorf("env. configuration value (`%s`) is not set", key)
	}
	return app.config.GetBool(key), nil
}
