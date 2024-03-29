// Code generated by goa v3.4.3, DO NOT EDIT.
//
// goa_starter HTTP client CLI support package
//
// Command:
// $ goa gen goa_starter/design

package cli

import (
	"flag"
	"fmt"
	goastarterc "goa_starter/gen/http/goa_starter/client"
	goastartercalcc "goa_starter/gen/http/goa_starter_calc/client"
	termlimitc "goa_starter/gen/http/term_limit/client"
	"net/http"
	"os"

	goahttp "goa.design/goa/v3/http"
	goa "goa.design/goa/v3/pkg"
)

// UsageCommands returns the set of commands and sub-commands using the format
//
//    command (subcommand1|subcommand2|...)
//
func UsageCommands() string {
	return `goa-starter add
goa-starter-calc add
term-limit get
`
}

// UsageExamples produces an example of a valid invocation of the CLI tool.
func UsageExamples() string {
	return os.Args[0] + ` goa-starter add --a 2783468530862518908 --b 5219281975954383774` + "\n" +
		os.Args[0] + ` goa-starter-calc add --c2 4786448540459506888 --d 7305426396677949776` + "\n" +
		os.Args[0] + ` term-limit get --body '{
      "id": 36983429280391996
   }' --muid "At vel maxime." --media-id 10477105092621605603` + "\n" +
		""
}

// ParseEndpoint returns the endpoint and payload as specified on the command
// line.
func ParseEndpoint(
	scheme, host string,
	doer goahttp.Doer,
	enc func(*http.Request) goahttp.Encoder,
	dec func(*http.Response) goahttp.Decoder,
	restore bool,
) (goa.Endpoint, interface{}, error) {
	var (
		goaStarterFlags = flag.NewFlagSet("goa-starter", flag.ContinueOnError)

		goaStarterAddFlags = flag.NewFlagSet("add", flag.ExitOnError)
		goaStarterAddAFlag = goaStarterAddFlags.String("a", "REQUIRED", "Left operand")
		goaStarterAddBFlag = goaStarterAddFlags.String("b", "REQUIRED", "Right operand")

		goaStarterCalcFlags = flag.NewFlagSet("goa-starter-calc", flag.ContinueOnError)

		goaStarterCalcAddFlags  = flag.NewFlagSet("add", flag.ExitOnError)
		goaStarterCalcAddC2Flag = goaStarterCalcAddFlags.String("c2", "REQUIRED", "Left operand")
		goaStarterCalcAddDFlag  = goaStarterCalcAddFlags.String("d", "REQUIRED", "Right operand")

		termLimitFlags = flag.NewFlagSet("term-limit", flag.ContinueOnError)

		termLimitGetFlags       = flag.NewFlagSet("get", flag.ExitOnError)
		termLimitGetBodyFlag    = termLimitGetFlags.String("body", "REQUIRED", "")
		termLimitGetMuidFlag    = termLimitGetFlags.String("muid", "REQUIRED", "")
		termLimitGetMediaIDFlag = termLimitGetFlags.String("media-id", "REQUIRED", "")
	)
	goaStarterFlags.Usage = goaStarterUsage
	goaStarterAddFlags.Usage = goaStarterAddUsage

	goaStarterCalcFlags.Usage = goaStarterCalcUsage
	goaStarterCalcAddFlags.Usage = goaStarterCalcAddUsage

	termLimitFlags.Usage = termLimitUsage
	termLimitGetFlags.Usage = termLimitGetUsage

	if err := flag.CommandLine.Parse(os.Args[1:]); err != nil {
		return nil, nil, err
	}

	if flag.NArg() < 2 { // two non flag args are required: SERVICE and ENDPOINT (aka COMMAND)
		return nil, nil, fmt.Errorf("not enough arguments")
	}

	var (
		svcn string
		svcf *flag.FlagSet
	)
	{
		svcn = flag.Arg(0)
		switch svcn {
		case "goa-starter":
			svcf = goaStarterFlags
		case "goa-starter-calc":
			svcf = goaStarterCalcFlags
		case "term-limit":
			svcf = termLimitFlags
		default:
			return nil, nil, fmt.Errorf("unknown service %q", svcn)
		}
	}
	if err := svcf.Parse(flag.Args()[1:]); err != nil {
		return nil, nil, err
	}

	var (
		epn string
		epf *flag.FlagSet
	)
	{
		epn = svcf.Arg(0)
		switch svcn {
		case "goa-starter":
			switch epn {
			case "add":
				epf = goaStarterAddFlags

			}

		case "goa-starter-calc":
			switch epn {
			case "add":
				epf = goaStarterCalcAddFlags

			}

		case "term-limit":
			switch epn {
			case "get":
				epf = termLimitGetFlags

			}

		}
	}
	if epf == nil {
		return nil, nil, fmt.Errorf("unknown %q endpoint %q", svcn, epn)
	}

	// Parse endpoint flags if any
	if svcf.NArg() > 1 {
		if err := epf.Parse(svcf.Args()[1:]); err != nil {
			return nil, nil, err
		}
	}

	var (
		data     interface{}
		endpoint goa.Endpoint
		err      error
	)
	{
		switch svcn {
		case "goa-starter":
			c := goastarterc.NewClient(scheme, host, doer, enc, dec, restore)
			switch epn {
			case "add":
				endpoint = c.Add()
				data, err = goastarterc.BuildAddPayload(*goaStarterAddAFlag, *goaStarterAddBFlag)
			}
		case "goa-starter-calc":
			c := goastartercalcc.NewClient(scheme, host, doer, enc, dec, restore)
			switch epn {
			case "add":
				endpoint = c.Add()
				data, err = goastartercalcc.BuildAddPayload(*goaStarterCalcAddC2Flag, *goaStarterCalcAddDFlag)
			}
		case "term-limit":
			c := termlimitc.NewClient(scheme, host, doer, enc, dec, restore)
			switch epn {
			case "get":
				endpoint = c.Get()
				data, err = termlimitc.BuildGetPayload(*termLimitGetBodyFlag, *termLimitGetMuidFlag, *termLimitGetMediaIDFlag)
			}
		}
	}
	if err != nil {
		return nil, nil, err
	}

	return endpoint, data, nil
}

// goa-starterUsage displays the usage of the goa-starter command and its
// subcommands.
func goaStarterUsage() {
	fmt.Fprintf(os.Stderr, `The goa_starter service performs operations on numbers.
Usage:
    %s [globalflags] goa-starter COMMAND [flags]

COMMAND:
    add: Add implements add.

Additional help:
    %s goa-starter COMMAND --help
`, os.Args[0], os.Args[0])
}
func goaStarterAddUsage() {
	fmt.Fprintf(os.Stderr, `%s [flags] goa-starter add -a INT -b INT

Add implements add.
    -a INT: Left operand
    -b INT: Right operand

Example:
    `+os.Args[0]+` goa-starter add --a 2783468530862518908 --b 5219281975954383774
`, os.Args[0])
}

// goa-starter-calcUsage displays the usage of the goa-starter-calc command and
// its subcommands.
func goaStarterCalcUsage() {
	fmt.Fprintf(os.Stderr, `The goa_starter-calc service performs operations on numbers.
Usage:
    %s [globalflags] goa-starter-calc COMMAND [flags]

COMMAND:
    add: Add implements add.

Additional help:
    %s goa-starter-calc COMMAND --help
`, os.Args[0], os.Args[0])
}
func goaStarterCalcAddUsage() {
	fmt.Fprintf(os.Stderr, `%s [flags] goa-starter-calc add -c2 INT -d INT

Add implements add.
    -c2 INT: Left operand
    -d INT: Right operand

Example:
    `+os.Args[0]+` goa-starter-calc add --c2 4786448540459506888 --d 7305426396677949776
`, os.Args[0])
}

// term-limitUsage displays the usage of the term-limit command and its
// subcommands.
func termLimitUsage() {
	fmt.Fprintf(os.Stderr, `OW表示可能一覧取得
Usage:
    %s [globalflags] term-limit COMMAND [flags]

COMMAND:
    get: Get implements get.

Additional help:
    %s term-limit COMMAND --help
`, os.Args[0], os.Args[0])
}
func termLimitGetUsage() {
	fmt.Fprintf(os.Stderr, `%s [flags] term-limit get -body JSON -muid STRING -media-id UINT64

Get implements get.
    -body JSON: 
    -muid STRING: 
    -media-id UINT64: 

Example:
    `+os.Args[0]+` term-limit get --body '{
      "id": 36983429280391996
   }' --muid "At vel maxime." --media-id 10477105092621605603
`, os.Args[0])
}
