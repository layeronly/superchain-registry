package superchain

import (
	"fmt"
	"reflect"

	"golang.org/x/mod/semver"
)

func validateConfigs(Superchains map[string]*Superchain,
	OPChains map[uint64]*ChainConfig,
	Addresses map[uint64]*AddressList,
	GenesisSystemConfigs map[uint64]*GenesisSystemConfig,
	Implementations map[uint64]ContractImplementations,
	SuperchainSemver ContractVersions,
) error {
	if err := SuperchainSemver.Validate(); err != nil {
		return fmt.Errorf("semver.yaml is invalid: %w", err)
	}
	return nil
}

// Validate will sanity check the validity of the semantic version strings
// in the ContractVersions struct.
func (c ContractVersions) Validate() error {
	val := reflect.ValueOf(c)
	for i := 0; i < val.NumField(); i++ {
		field := val.Field(i)
		str, ok := field.Interface().(string)
		if !ok {
			return fmt.Errorf("invalid type for field %s", val.Type().Field(i).Name)
		}
		if str == "" {
			return fmt.Errorf("empty version for field %s", val.Type().Field(i).Name)
		}
		str = canonicalizeSemver(str)
		if !semver.IsValid(str) {
			return fmt.Errorf("invalid semver %s for field %s", str, val.Type().Field(i).Name)
		}
	}
	return nil
}
