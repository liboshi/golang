package main

import (
	"context"
	"flag"
	"fmt"
	"log"
	"os"
	"strconv"
	"strings"

	"github.com/google/subcommands"
)

type genSessionConfig struct {
	cfgFile     string
	topoCfgFile string
}

type genSessionData struct {
	isSession   bool
	daasVMId    string
	tenantId    string
	topoCfgFile string
}

type genTopo struct {
	cfgFile     string
	spId        string
	tenantId    string
	cmsTenantId string
	product     string
	poolStart   int
	poolNum     int
	poolSize    int
}

type runTest struct {
}

// Functions for generating configure files.
// Those functions are used to implement sub-commands.
func (*genSessionConfig) Name() string     { return "gensessionconfig" }
func (*genSessionConfig) Synopsis() string { return "Generate configure files." }
func (*genSessionConfig) Usage() string {
	return "gensessionconfig [-cfgfile=] [-topocfgfile=]\n"
}

func (c *genSessionConfig) SetFlags(f *flag.FlagSet) {
	f.StringVar(&c.cfgFile, "cfgfile", "test.csv", "Session configuration definition file")
	f.StringVar(&c.topoCfgFile, "topocfgfile", "test.csv", "Topo configuration definition file")
}

func (c *genSessionConfig) Execute(_ context.Context, f *flag.FlagSet, _ ...interface{}) subcommands.ExitStatus {
	topoCfgDef := TopoCSVFileToMap(c.topoCfgFile)
	sessionConfig := SessionCfgCSVFileToMap(c.cfgFile)
	for _, v := range topoCfgDef {
		pSize := strings.Split(v.Psize, ",")
		GenerateConfigJson(sessionConfig, v.Object, pSize)
	}
	return subcommands.ExitSuccess
}

// Functions for generating data files.
// Those functions are used to implement sub-commands.
func (*genSessionData) Name() string     { return "gensessiondata" }
func (*genSessionData) Synopsis() string { return "Genenrate data files." }
func (*genSessionData) Usage() string    { return "gensessiondata [-s=] [-topocfgfile=] [-did=] [-tid=]\n" }

func (d *genSessionData) SetFlags(f *flag.FlagSet) {
	f.BoolVar(&d.isSession, "s", false, "Generate data file for session or No session.")
	f.StringVar(&d.daasVMId, "did", "daas_vm_id", "Daas VM ID")
	f.StringVar(&d.tenantId, "tid", "tenantid", "Tenant ID")
	f.StringVar(&d.topoCfgFile, "topocfgfile", "test.csv", "Topo configuration definition file")
}

func (d *genSessionData) Execute(_ context.Context, f *flag.FlagSet, _ ...interface{}) subcommands.ExitStatus {
	topoCfgDef := TopoCSVFileToMap(d.topoCfgFile)
	for _, v := range topoCfgDef {
		GenerateDataJson(d.isSession, d.daasVMId, v.Id, v.Object)
	}
	return subcommands.ExitSuccess
}

// Functions for generating topo configure and data files.
// Those functions are used to implement sub-commands.
func (*genTopo) Name() string     { return "gentopo" }
func (*genTopo) Synopsis() string { return "Generate topo configure and data files." }
func (*genTopo) Usage() string {
	return "gentopo [-cfgfile=] [-sid=] [-tid=] [-ctid=] [-product=] [-pstart=] [-pnum=] [-psize=]\n"
}

func (t *genTopo) SetFlags(f *flag.FlagSet) {
	f.StringVar(&t.cfgFile, "cfgfile", "", "Topo configuration definition file")
	f.StringVar(&t.spId, "sid", "00", "SP id")
	f.StringVar(&t.tenantId, "tid", "tenantid", "Tenant id")
	f.StringVar(&t.cmsTenantId, "ctid", "cmstenantid", "CMS Tenant ID")
	f.StringVar(&t.product, "p", "product", "Product name or namespace")
	f.IntVar(&t.poolStart, "pstart", 0, "Pool start with")
	f.IntVar(&t.poolNum, "pnum", 1, "Number of pool")
	f.IntVar(&t.poolSize, "psize", 1, "The size of pool")
}

func (t *genTopo) Execute(_ context.Context, f *flag.FlagSet, _ ...interface{}) subcommands.ExitStatus {
	if t.cfgFile == "" {
		GenerateTopoJson(t.spId, t.tenantId, t.cmsTenantId, t.product, t.poolStart, t.poolNum)
	} else {
		if _, err := os.Stat(t.cfgFile); os.IsNotExist(err) {
			log.Fatal("%s is not exists", t.cfgFile)
		} else {
			topoCfgDef := TopoCSVFileToMap(t.cfgFile)
			for _, v := range topoCfgDef {
				newStructValue(v, t)
				PoolCapacities = strings.Split(v.Psize, ",")
				GenerateTopoJson(t.spId, t.tenantId, t.cmsTenantId, t.product, t.poolStart, t.poolNum)
			}
		}
	}
	return subcommands.ExitSuccess
}

// struct value
func newStructValue(val TopoCSVStruct, p *genTopo) *genTopo {
	p.spId = val.Type
	p.tenantId = val.Object
	p.cmsTenantId = val.Id
	p.product = val.Namespace
	p.poolNum, _ = strconv.Atoi(val.Pnum)
	p.poolSize, _ = strconv.Atoi(val.Psize)
	return (*genTopo)(p)
}

// run test
// Functions for testing.
// Those functions are used to implement sub-commands.
func (*runTest) Name() string     { return "test" }
func (*runTest) Synopsis() string { return "Run test functions." }
func (*runTest) Usage() string    { return "test\n" }

func (t *runTest) SetFlags(f *flag.FlagSet) {
}

func (t *runTest) Execute(_ context.Context, f *flag.FlagSet, _ ...interface{}) subcommands.ExitStatus {
	fmt.Println("******")
	sessionConfig := SessionCfgCSVFileToMap("test.csv")
	fmt.Println(sessionConfig)
	//GenerateConfigJson(sessionConfig)
	return subcommands.ExitSuccess
}

func main() {
	subcommands.Register(subcommands.HelpCommand(), "")
	subcommands.Register(subcommands.FlagsCommand(), "")
	subcommands.Register(subcommands.CommandsCommand(), "")
	subcommands.Register(&genSessionConfig{}, "")
	subcommands.Register(&genSessionData{}, "")
	subcommands.Register(&genTopo{}, "")
	subcommands.Register(&runTest{}, "")

	// Parse commandline flags.
	flag.Parse()
	ctx := context.Background()
	os.Exit(int(subcommands.Execute(ctx)))
}
