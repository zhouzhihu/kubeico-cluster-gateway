package config

import (
	"fmt"

	"github.com/spf13/pflag"
)

var SecretNamespace string

func ValidateSecret() error {
	if len(SecretNamespace) == 0 {
		return fmt.Errorf("must specify --secret-namespace")
	}
	return nil
}

func AddSecretFlags(set *pflag.FlagSet) {
	set.StringVarP(&SecretNamespace, "secret-namespace", "", "",
		"the namespace to reading secrets")
}
