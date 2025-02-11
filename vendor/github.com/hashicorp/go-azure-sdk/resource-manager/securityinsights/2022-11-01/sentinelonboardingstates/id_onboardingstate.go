package sentinelonboardingstates

import (
	"fmt"
	"strings"

	"github.com/hashicorp/go-azure-helpers/resourcemanager/resourceids"
)

var _ resourceids.ResourceId = OnboardingStateId{}

// OnboardingStateId is a struct representing the Resource ID for a Onboarding State
type OnboardingStateId struct {
	SubscriptionId              string
	ResourceGroupName           string
	WorkspaceName               string
	SentinelOnboardingStateName string
}

// NewOnboardingStateID returns a new OnboardingStateId struct
func NewOnboardingStateID(subscriptionId string, resourceGroupName string, workspaceName string, sentinelOnboardingStateName string) OnboardingStateId {
	return OnboardingStateId{
		SubscriptionId:              subscriptionId,
		ResourceGroupName:           resourceGroupName,
		WorkspaceName:               workspaceName,
		SentinelOnboardingStateName: sentinelOnboardingStateName,
	}
}

// ParseOnboardingStateID parses 'input' into a OnboardingStateId
func ParseOnboardingStateID(input string) (*OnboardingStateId, error) {
	parser := resourceids.NewParserFromResourceIdType(OnboardingStateId{})
	parsed, err := parser.Parse(input, false)
	if err != nil {
		return nil, fmt.Errorf("parsing %q: %+v", input, err)
	}

	var ok bool
	id := OnboardingStateId{}

	if id.SubscriptionId, ok = parsed.Parsed["subscriptionId"]; !ok {
		return nil, fmt.Errorf("the segment 'subscriptionId' was not found in the resource id %q", input)
	}

	if id.ResourceGroupName, ok = parsed.Parsed["resourceGroupName"]; !ok {
		return nil, fmt.Errorf("the segment 'resourceGroupName' was not found in the resource id %q", input)
	}

	if id.WorkspaceName, ok = parsed.Parsed["workspaceName"]; !ok {
		return nil, fmt.Errorf("the segment 'workspaceName' was not found in the resource id %q", input)
	}

	if id.SentinelOnboardingStateName, ok = parsed.Parsed["sentinelOnboardingStateName"]; !ok {
		return nil, fmt.Errorf("the segment 'sentinelOnboardingStateName' was not found in the resource id %q", input)
	}

	return &id, nil
}

// ParseOnboardingStateIDInsensitively parses 'input' case-insensitively into a OnboardingStateId
// note: this method should only be used for API response data and not user input
func ParseOnboardingStateIDInsensitively(input string) (*OnboardingStateId, error) {
	parser := resourceids.NewParserFromResourceIdType(OnboardingStateId{})
	parsed, err := parser.Parse(input, true)
	if err != nil {
		return nil, fmt.Errorf("parsing %q: %+v", input, err)
	}

	var ok bool
	id := OnboardingStateId{}

	if id.SubscriptionId, ok = parsed.Parsed["subscriptionId"]; !ok {
		return nil, fmt.Errorf("the segment 'subscriptionId' was not found in the resource id %q", input)
	}

	if id.ResourceGroupName, ok = parsed.Parsed["resourceGroupName"]; !ok {
		return nil, fmt.Errorf("the segment 'resourceGroupName' was not found in the resource id %q", input)
	}

	if id.WorkspaceName, ok = parsed.Parsed["workspaceName"]; !ok {
		return nil, fmt.Errorf("the segment 'workspaceName' was not found in the resource id %q", input)
	}

	if id.SentinelOnboardingStateName, ok = parsed.Parsed["sentinelOnboardingStateName"]; !ok {
		return nil, fmt.Errorf("the segment 'sentinelOnboardingStateName' was not found in the resource id %q", input)
	}

	return &id, nil
}

// ValidateOnboardingStateID checks that 'input' can be parsed as a Onboarding State ID
func ValidateOnboardingStateID(input interface{}, key string) (warnings []string, errors []error) {
	v, ok := input.(string)
	if !ok {
		errors = append(errors, fmt.Errorf("expected %q to be a string", key))
		return
	}

	if _, err := ParseOnboardingStateID(v); err != nil {
		errors = append(errors, err)
	}

	return
}

// ID returns the formatted Onboarding State ID
func (id OnboardingStateId) ID() string {
	fmtString := "/subscriptions/%s/resourceGroups/%s/providers/Microsoft.OperationalInsights/workspaces/%s/providers/Microsoft.SecurityInsights/onboardingStates/%s"
	return fmt.Sprintf(fmtString, id.SubscriptionId, id.ResourceGroupName, id.WorkspaceName, id.SentinelOnboardingStateName)
}

// Segments returns a slice of Resource ID Segments which comprise this Onboarding State ID
func (id OnboardingStateId) Segments() []resourceids.Segment {
	return []resourceids.Segment{
		resourceids.StaticSegment("staticSubscriptions", "subscriptions", "subscriptions"),
		resourceids.SubscriptionIdSegment("subscriptionId", "12345678-1234-9876-4563-123456789012"),
		resourceids.StaticSegment("staticResourceGroups", "resourceGroups", "resourceGroups"),
		resourceids.ResourceGroupSegment("resourceGroupName", "example-resource-group"),
		resourceids.StaticSegment("staticProviders", "providers", "providers"),
		resourceids.ResourceProviderSegment("staticMicrosoftOperationalInsights", "Microsoft.OperationalInsights", "Microsoft.OperationalInsights"),
		resourceids.StaticSegment("staticWorkspaces", "workspaces", "workspaces"),
		resourceids.UserSpecifiedSegment("workspaceName", "workspaceValue"),
		resourceids.StaticSegment("staticProviders2", "providers", "providers"),
		resourceids.ResourceProviderSegment("staticMicrosoftSecurityInsights", "Microsoft.SecurityInsights", "Microsoft.SecurityInsights"),
		resourceids.StaticSegment("staticOnboardingStates", "onboardingStates", "onboardingStates"),
		resourceids.UserSpecifiedSegment("sentinelOnboardingStateName", "sentinelOnboardingStateValue"),
	}
}

// String returns a human-readable description of this Onboarding State ID
func (id OnboardingStateId) String() string {
	components := []string{
		fmt.Sprintf("Subscription: %q", id.SubscriptionId),
		fmt.Sprintf("Resource Group Name: %q", id.ResourceGroupName),
		fmt.Sprintf("Workspace Name: %q", id.WorkspaceName),
		fmt.Sprintf("Sentinel Onboarding State Name: %q", id.SentinelOnboardingStateName),
	}
	return fmt.Sprintf("Onboarding State (%s)", strings.Join(components, "\n"))
}
