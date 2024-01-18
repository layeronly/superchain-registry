package superchain

import (
	"fmt"
	"reflect"
	"strings"

	"golang.org/x/mod/semver"
)

func validateConfigs(Superchains map[string]*Superchain,
	OPChains map[uint64]*ChainConfig,
	Addresses map[uint64]*AddressList,
	GenesisSystemConfigs map[uint64]*GenesisSystemConfig,
	Implementations map[uint64]ContractImplementations,
	SuperchainSemver ContractVersions,
) error {
	if err := validateUniqueChainIds(OPChains); err != nil {
		return fmt.Errorf("chain IDs not unique: %w", err)
	}
	if err := SuperchainSemver.Validate(); err != nil {
		return fmt.Errorf("semver.yaml is invalid: %w", err)
	}
	return nil
}

func validateUniqueChainIds(OPChains map[uint64]*ChainConfig) error {
	// Here we assume the code building the OPChains mapping
	// errored or panicked if it found a duplicate
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

// canonicalizeSemver will ensure that the version string has a "v" prefix.
// This is because the semver library being used requires the "v" prefix,
// even though
func canonicalizeSemver(version string) string {
	if !strings.HasPrefix(version, "v") {
		version = "v" + version
	}
	return version
}
