package openapi

import (
	"github.com/aliyun/aliyun-cli/cli"
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestAFlags(t *testing.T) {
	flagset := cli.NewFlagSet()
	AddFlags(flagset)
	secureflag := SecureFlag(flagset)
	assert.Equal(t, "secure", secureflag.Name)
	assert.Equal(t, "use `--secure` to force https", secureflag.Short.Text())

	forceflag := ForceFlag(flagset)
	assert.Equal(t, "force", forceflag.Name)
	assert.Equal(t, "use `--force` to skip api and parameters check", forceflag.Short.Text())

	endpointflag := EndpointFlag(flagset)
	assert.Equal(t, "endpoint", endpointflag.Name)
	assert.Equal(t, "use `--endpoint <endpoint>` to assign endpoint", endpointflag.Short.Text())

	versionflag := VersionFlag(flagset)
	assert.Equal(t, "version", versionflag.Name)
	assert.Equal(t, "use `--version <YYYY-MM-DD>` to assign product api version", versionflag.Short.Text())

	headerflag := HeaderFlag(flagset)
	assert.Equal(t, "header", headerflag.Name)
	assert.Equal(t, "use `--header X-foo=bar` to add custom HTTP header, repeatable", headerflag.Short.Text())

	bodyflag := BodyFlag(flagset)
	assert.Equal(t, "body", bodyflag.Name)
	assert.Equal(t, "use `--body $(cat foo.json)` to assign http body in RESTful call", bodyflag.Short.Text())

	bodyfileflag := BodyFileFlag(flagset)
	assert.Equal(t, "body-file", bodyfileflag.Name)
	assert.Equal(t, "assign http body in Restful call with local file", bodyfileflag.Short.Text())

	acceptflag := AcceptFlag(flagset)
	assert.Equal(t, "accept", acceptflag.Name)
	assert.Equal(t, "add `--accept {json|xml}` to add Accept header", acceptflag.Short.Text())

	roaflag := RoaFlag(flagset)
	assert.Equal(t, "roa", roaflag.Name)
	assert.Equal(t, "use `--roa {GET|PUT|POST|DELETE}` to assign restful call.[DEPRECATED]", roaflag.Short.Text())

	dryrunflag := DryRunFlag(flagset)
	assert.Equal(t, "dryrun", dryrunflag.Name)
	assert.Equal(t, "add `--dryrun` to validate and print request without running.", dryrunflag.Short.Text())

	quietflag := QuietFlag(flagset)
	assert.Equal(t, "quiet", quietflag.Name)
	assert.Equal(t, "add `--quiet` to hide normal output", quietflag.Short.Text())

	outputflag := OutputFlag(flagset)
	assert.Equal(t, "output", outputflag.Name)
	assert.Equal(t, "use `--output cols=Field1,Field2 [rows=jmesPath]` to print output as table", outputflag.Short.Text())
}
