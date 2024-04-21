package common

import (
	"github.com/crossplane/crossplane-runtime/pkg/fieldpath"
	"github.com/crossplane/crossplane-runtime/pkg/reference"
	xpresource "github.com/crossplane/crossplane-runtime/pkg/resource"
)

const (
	// SelfPackagePath is the golang path for this package.
	SelfPackagePath = "github.com/viletay/provider-keycloak/config/common"

	// PathServiceAccountRoleIDExtractor is the golang path to ARNExtractor function
	// in this package.
	PathServiceAccountRoleIDExtractor = SelfPackagePath + ".ServiceAccountRoleIDExtractor()"
)

func ServiceAccountRoleIDExtractor() reference.ExtractValueFn {
	return func(mg xpresource.Managed) string {
		paved, err := fieldpath.PaveObject(mg)
		if err != nil {
			return ""
		}
		r, err := paved.GetString("status.atProvider.serviceAccountUserId")
		if err != nil {
			return ""
		}
		return r
	}
}
