package repo

import (
	"fmt"
	"strconv"
	"strings"
)

type VersionComparator struct{}

func NewVersionComparator() *VersionComparator { return &VersionComparator{} }

type VersionDiff struct {
	IsNewer        bool
	IsSame         bool
	IsOlder        bool
	CurrentVersion string
	TargetVersion  string
}

func (c *VersionComparator) Compare(current, target string) (*VersionDiff, error) {
	if current == "" || target == "" {
		return nil, fmt.Errorf("versions cannot be empty")
	}

	currentParts, err := parseVersion(current)
	if err != nil {
		return nil, fmt.Errorf("invalid current version %s: %w", current, err)
	}

	targetParts, err := parseVersion(target)
	if err != nil {
		return nil, fmt.Errorf("invalid target version %s: %w", target, err)
	}

	comparison := compareVersionParts(currentParts, targetParts)

	return &VersionDiff{
		IsNewer:        comparison > 0,
		IsSame:         comparison == 0,
		IsOlder:        comparison < 0,
		CurrentVersion: current,
		TargetVersion:  target,
	}, nil
}

func (c *VersionComparator) IsUpdateAvailable(current, target string) (bool, error) {
	diff, err := c.Compare(current, target)
	if err != nil {
		return false, err
	}
	return diff.IsOlder, nil
}

func parseVersion(version string) ([]int, error) {
	version = strings.TrimPrefix(version, "v")
	parts := strings.Split(version, ".")
	if len(parts) != 3 {
		return nil, fmt.Errorf("invalid semver format (expected major.minor.patch)")
	}
	result := make([]int, 3)
	for i, part := range parts {
		num, err := strconv.Atoi(part)
		if err != nil {
			return nil, fmt.Errorf("invalid version number: %s", part)
		}
		result[i] = num
	}
	return result, nil
}

func compareVersionParts(v1, v2 []int) int {
	for i := 0; i < 3; i++ {
		if v1[i] > v2[i] {
			return 1
		}
		if v1[i] < v2[i] {
			return -1
		}
	}
	return 0
}
