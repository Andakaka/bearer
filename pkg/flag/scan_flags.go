package flag

import (
	"github.com/bearer/curio/pkg/types"
)

var (
	SkipDirsFlag = Flag{
		Name:       "skip-dirs",
		ConfigName: "scan.skip-dirs",
		Value:      []string{},
		Usage:      "specify the directories where the traversal is skipped",
	}
	SkipFilesFlag = Flag{
		Name:       "skip-files",
		ConfigName: "scan.skip-files",
		Value:      []string{},
		Usage:      "specify the file paths to skip traversal",
	}
	OfflineScanFlag = Flag{
		Name:       "offline-scan",
		ConfigName: "scan.offline",
		Value:      false,
		Usage:      "do not issue API requests to identify dependencies",
	}
	SecurityChecksFlag = Flag{
		Name:       "security-checks",
		ConfigName: "scan.security-checks",
		Value:      []string{types.SecurityCheckConfig},
		Usage:      "comma-separated list of what security issues to detect (vuln,config)",
	}
	FilePatternsFlag = Flag{
		Name:       "file-patterns",
		ConfigName: "scan.file-patterns",
		Value:      []string{},
		Usage:      "specify config file patterns",
	}
)

type ScanFlagGroup struct {
	SkipDirs  *Flag
	SkipFiles *Flag
	// SecurityChecks *Flag
	FilePatterns *Flag
}

type ScanOptions struct {
	Target         string
	SkipDirs       []string
	SkipFiles      []string
	SecurityChecks []string
	FilePatterns   []string
}

func NewScanFlagGroup() *ScanFlagGroup {
	return &ScanFlagGroup{
		SkipDirs:  &SkipDirsFlag,
		SkipFiles: &SkipFilesFlag,
		// SecurityChecks: &SecurityChecksFlag,
		FilePatterns: &FilePatternsFlag,
	}
}

func (f *ScanFlagGroup) Name() string {
	return "Scan"
}

func (f *ScanFlagGroup) Flags() []*Flag {
	return []*Flag{f.SkipDirs, f.SkipFiles, f.FilePatterns}
}

func (f *ScanFlagGroup) ToOptions(args []string) (ScanOptions, error) {
	var target string
	if len(args) == 1 {
		target = args[0]
	}
	// securityChecks, err := parseSecurityCheck(getStringSlice(f.SecurityChecks))
	// if err != nil {
	// 	return ScanOptions{}, xerrors.Errorf("unable to parse security checks: %w", err)
	// }

	return ScanOptions{
		Target:    target,
		SkipDirs:  getStringSlice(f.SkipDirs),
		SkipFiles: getStringSlice(f.SkipFiles),
		// SecurityChecks: securityChecks,
		FilePatterns: getStringSlice(f.FilePatterns),
	}, nil
}

// func parseSecurityCheck(securityCheck []string) ([]string, error) {
// 	var securityChecks []string
// 	for _, v := range securityCheck {
// 		if !slices.Contains(types.SecurityChecks, v) {
// 			return nil, xerrors.Errorf("unknown security check: %s", v)
// 		}
// 		securityChecks = append(securityChecks, v)
// 	}
// 	return securityChecks, nil
// }