// Package testing provides access to the Cloud Testing API.
//
// See https://developers.google.com/cloud-test-lab/
//
// Usage example:
//
//   import "google.golang.org/api/testing/v1"
//   ...
//   testingService, err := testing.New(oauthHttpClient)
package testing // import "google.golang.org/api/testing/v1"

import (
	"bytes"
	"encoding/json"
	"errors"
	"fmt"
	context "golang.org/x/net/context"
	ctxhttp "golang.org/x/net/context/ctxhttp"
	gensupport "google.golang.org/api/gensupport"
	googleapi "google.golang.org/api/googleapi"
	"io"
	"net/http"
	"net/url"
	"strconv"
	"strings"
)

// Always reference these packages, just in case the auto-generated code
// below doesn't.
var _ = bytes.NewBuffer
var _ = strconv.Itoa
var _ = fmt.Sprintf
var _ = json.NewDecoder
var _ = io.Copy
var _ = url.Parse
var _ = gensupport.MarshalJSON
var _ = googleapi.Version
var _ = errors.New
var _ = strings.Replace
var _ = context.Canceled
var _ = ctxhttp.Do

const apiId = "testing:v1"
const apiName = "testing"
const apiVersion = "v1"
const basePath = "https://testing.googleapis.com/"

// OAuth2 scopes used by this API.
const (
	// View and manage your data across Google Cloud Platform services
	CloudPlatformScope = "https://www.googleapis.com/auth/cloud-platform"

	// View your data across Google Cloud Platform services
	CloudPlatformReadOnlyScope = "https://www.googleapis.com/auth/cloud-platform.read-only"
)

func New(client *http.Client) (*Service, error) {
	if client == nil {
		return nil, errors.New("client is nil")
	}
	s := &Service{client: client, BasePath: basePath}
	s.ApplicationDetailService = NewApplicationDetailServiceService(s)
	s.Projects = NewProjectsService(s)
	s.TestEnvironmentCatalog = NewTestEnvironmentCatalogService(s)
	return s, nil
}

type Service struct {
	client    *http.Client
	BasePath  string // API endpoint base URL
	UserAgent string // optional additional User-Agent fragment

	ApplicationDetailService *ApplicationDetailServiceService

	Projects *ProjectsService

	TestEnvironmentCatalog *TestEnvironmentCatalogService
}

func (s *Service) userAgent() string {
	if s.UserAgent == "" {
		return googleapi.UserAgent
	}
	return googleapi.UserAgent + " " + s.UserAgent
}

func NewApplicationDetailServiceService(s *Service) *ApplicationDetailServiceService {
	rs := &ApplicationDetailServiceService{s: s}
	return rs
}

type ApplicationDetailServiceService struct {
	s *Service
}

func NewProjectsService(s *Service) *ProjectsService {
	rs := &ProjectsService{s: s}
	rs.TestMatrices = NewProjectsTestMatricesService(s)
	return rs
}

type ProjectsService struct {
	s *Service

	TestMatrices *ProjectsTestMatricesService
}

func NewProjectsTestMatricesService(s *Service) *ProjectsTestMatricesService {
	rs := &ProjectsTestMatricesService{s: s}
	return rs
}

type ProjectsTestMatricesService struct {
	s *Service
}

func NewTestEnvironmentCatalogService(s *Service) *TestEnvironmentCatalogService {
	rs := &TestEnvironmentCatalogService{s: s}
	return rs
}

type TestEnvironmentCatalogService struct {
	s *Service
}

// Account: Identifies an account and how to log into it
type Account struct {
	// GoogleAuto: An automatic google login account
	GoogleAuto *GoogleAuto `json:"googleAuto,omitempty"`

	// ForceSendFields is a list of field names (e.g. "GoogleAuto") to
	// unconditionally include in API requests. By default, fields with
	// empty values are omitted from API requests. However, any non-pointer,
	// non-interface field appearing in ForceSendFields will be sent to the
	// server regardless of whether the field is empty or not. This may be
	// used to include empty fields in Patch requests.
	ForceSendFields []string `json:"-"`

	// NullFields is a list of field names (e.g. "GoogleAuto") to include in
	// API requests with the JSON null value. By default, fields with empty
	// values are omitted from API requests. However, any field with an
	// empty value appearing in NullFields will be sent to the server as
	// null. It is an error if a field in this list has a non-empty value.
	// This may be used to include null fields in Patch requests.
	NullFields []string `json:"-"`
}

func (s *Account) MarshalJSON() ([]byte, error) {
	type NoMethod Account
	raw := NoMethod(*s)
	return gensupport.MarshalJSON(raw, s.ForceSendFields, s.NullFields)
}

// AndroidDevice: A single Android device.
type AndroidDevice struct {
	// AndroidModelId: The id of the Android device to be used.
	// Use the EnvironmentDiscoveryService to get supported
	// options.
	// Required
	AndroidModelId string `json:"androidModelId,omitempty"`

	// AndroidVersionId: The id of the Android OS version to be used.
	// Use the EnvironmentDiscoveryService to get supported
	// options.
	// Required
	AndroidVersionId string `json:"androidVersionId,omitempty"`

	// Locale: The locale the test device used for testing.
	// Use the EnvironmentDiscoveryService to get supported
	// options.
	// Required
	Locale string `json:"locale,omitempty"`

	// Orientation: How the device is oriented during the test.
	// Use the EnvironmentDiscoveryService to get supported
	// options.
	// Required
	Orientation string `json:"orientation,omitempty"`

	// ForceSendFields is a list of field names (e.g. "AndroidModelId") to
	// unconditionally include in API requests. By default, fields with
	// empty values are omitted from API requests. However, any non-pointer,
	// non-interface field appearing in ForceSendFields will be sent to the
	// server regardless of whether the field is empty or not. This may be
	// used to include empty fields in Patch requests.
	ForceSendFields []string `json:"-"`

	// NullFields is a list of field names (e.g. "AndroidModelId") to
	// include in API requests with the JSON null value. By default, fields
	// with empty values are omitted from API requests. However, any field
	// with an empty value appearing in NullFields will be sent to the
	// server as null. It is an error if a field in this list has a
	// non-empty value. This may be used to include null fields in Patch
	// requests.
	NullFields []string `json:"-"`
}

func (s *AndroidDevice) MarshalJSON() ([]byte, error) {
	type NoMethod AndroidDevice
	raw := NoMethod(*s)
	return gensupport.MarshalJSON(raw, s.ForceSendFields, s.NullFields)
}

// AndroidDeviceCatalog: The currently supported Android devices.
type AndroidDeviceCatalog struct {
	// Models: The set of supported Android device models.
	// @OutputOnly
	Models []*AndroidModel `json:"models,omitempty"`

	// RuntimeConfiguration: The set of supported runtime
	// configurations.
	// @OutputOnly
	RuntimeConfiguration *AndroidRuntimeConfiguration `json:"runtimeConfiguration,omitempty"`

	// Versions: The set of supported Android OS versions.
	// @OutputOnly
	Versions []*AndroidVersion `json:"versions,omitempty"`

	// ForceSendFields is a list of field names (e.g. "Models") to
	// unconditionally include in API requests. By default, fields with
	// empty values are omitted from API requests. However, any non-pointer,
	// non-interface field appearing in ForceSendFields will be sent to the
	// server regardless of whether the field is empty or not. This may be
	// used to include empty fields in Patch requests.
	ForceSendFields []string `json:"-"`

	// NullFields is a list of field names (e.g. "Models") to include in API
	// requests with the JSON null value. By default, fields with empty
	// values are omitted from API requests. However, any field with an
	// empty value appearing in NullFields will be sent to the server as
	// null. It is an error if a field in this list has a non-empty value.
	// This may be used to include null fields in Patch requests.
	NullFields []string `json:"-"`
}

func (s *AndroidDeviceCatalog) MarshalJSON() ([]byte, error) {
	type NoMethod AndroidDeviceCatalog
	raw := NoMethod(*s)
	return gensupport.MarshalJSON(raw, s.ForceSendFields, s.NullFields)
}

// AndroidDeviceList: A list of Android device configurations in which
// the test is to be executed.
type AndroidDeviceList struct {
	// AndroidDevices: A list of Android devices
	// Required
	AndroidDevices []*AndroidDevice `json:"androidDevices,omitempty"`

	// ForceSendFields is a list of field names (e.g. "AndroidDevices") to
	// unconditionally include in API requests. By default, fields with
	// empty values are omitted from API requests. However, any non-pointer,
	// non-interface field appearing in ForceSendFields will be sent to the
	// server regardless of whether the field is empty or not. This may be
	// used to include empty fields in Patch requests.
	ForceSendFields []string `json:"-"`

	// NullFields is a list of field names (e.g. "AndroidDevices") to
	// include in API requests with the JSON null value. By default, fields
	// with empty values are omitted from API requests. However, any field
	// with an empty value appearing in NullFields will be sent to the
	// server as null. It is an error if a field in this list has a
	// non-empty value. This may be used to include null fields in Patch
	// requests.
	NullFields []string `json:"-"`
}

func (s *AndroidDeviceList) MarshalJSON() ([]byte, error) {
	type NoMethod AndroidDeviceList
	raw := NoMethod(*s)
	return gensupport.MarshalJSON(raw, s.ForceSendFields, s.NullFields)
}

// AndroidInstrumentationTest: A test of an Android application that can
// control an Android component
// independently of its normal lifecycle.
// Android instrumentation tests run an application APK and test APK
// inside the
// same process on a virtual or physical AndroidDevice.  They also
// specify
// a test runner class, such as com.google.GoogleTestRunner, which can
// vary
// on the specific instrumentation framework chosen.
//
// See <http://developer.android.com/tools/testing/testing_android.html>
// for
// more information on types of Android tests.
type AndroidInstrumentationTest struct {
	// AppApk: The APK for the application under test.
	AppApk *FileReference `json:"appApk,omitempty"`

	// AppPackageId: The java package for the application under
	// test.
	// Optional, default is determined by examining the application's
	// manifest.
	AppPackageId string `json:"appPackageId,omitempty"`

	// OrchestratorOption: The option of whether running each test within
	// its own invocation of
	// instrumentation with Android Test Orchestrator or not.
	// ** Orchestrator is only compatible with AndroidJUnitRunner version
	// 1.0 or
	// higher! **
	// Orchestrator offers the following benefits:
	//  - No shared state
	//  - Crashes are isolated
	//  - Logs are scoped per
	// test
	//
	// See
	// <https://developer.android.com/training/testing/junit-runner
	// .html#using-android-test-orchestrator>
	// for more information about Android Test Orchestrator.
	//
	// Optional. If not set, the test will be run without the orchestrator.
	//
	// Possible values:
	//   "ORCHESTRATOR_OPTION_UNSPECIFIED" - Default value: the server will
	// choose the mode. Currently implies that
	// the test will run without the orchestrator. In the future,
	// all instrumentation tests will be run with the orchestrator.
	// Using the orchestrator is highly encouraged because of all the
	// benefits it
	// offers.
	//   "USE_ORCHESTRATOR" - Run test using orchestrator.
	// ** Only compatible with AndroidJUnitRunner version 1.0 or higher!
	// **
	// Recommended.
	//   "DO_NOT_USE_ORCHESTRATOR" - Run test without using orchestrator.
	OrchestratorOption string `json:"orchestratorOption,omitempty"`

	// TestApk: The APK containing the test code to be executed.
	// Required
	TestApk *FileReference `json:"testApk,omitempty"`

	// TestPackageId: The java package for the test to be
	// executed.
	// Optional, default is determined by examining the application's
	// manifest.
	TestPackageId string `json:"testPackageId,omitempty"`

	// TestRunnerClass: The InstrumentationTestRunner class.
	// Optional, default is determined by examining the application's
	// manifest.
	TestRunnerClass string `json:"testRunnerClass,omitempty"`

	// TestTargets: Each target must be fully qualified with the package
	// name or class name,
	// in one of these formats:
	//  - "package package_name"
	//  - "class package_name.class_name"
	//  - "class package_name.class_name#method_name"
	//
	// Optional, if empty, all targets in the module will be run.
	TestTargets []string `json:"testTargets,omitempty"`

	// ForceSendFields is a list of field names (e.g. "AppApk") to
	// unconditionally include in API requests. By default, fields with
	// empty values are omitted from API requests. However, any non-pointer,
	// non-interface field appearing in ForceSendFields will be sent to the
	// server regardless of whether the field is empty or not. This may be
	// used to include empty fields in Patch requests.
	ForceSendFields []string `json:"-"`

	// NullFields is a list of field names (e.g. "AppApk") to include in API
	// requests with the JSON null value. By default, fields with empty
	// values are omitted from API requests. However, any field with an
	// empty value appearing in NullFields will be sent to the server as
	// null. It is an error if a field in this list has a non-empty value.
	// This may be used to include null fields in Patch requests.
	NullFields []string `json:"-"`
}

func (s *AndroidInstrumentationTest) MarshalJSON() ([]byte, error) {
	type NoMethod AndroidInstrumentationTest
	raw := NoMethod(*s)
	return gensupport.MarshalJSON(raw, s.ForceSendFields, s.NullFields)
}

// AndroidMatrix: A set of Android device configuration permutations is
// defined by the
// the cross-product of the given axes.  Internally, the given
// AndroidMatrix
// will be expanded into a set of AndroidDevices.
//
// Only supported permutations will be instantiated.  Invalid
// permutations
// (e.g., incompatible models/versions) are ignored.
type AndroidMatrix struct {
	// AndroidModelIds: The ids of the set of Android device to be used.
	// Use the EnvironmentDiscoveryService to get supported
	// options.
	// Required
	AndroidModelIds []string `json:"androidModelIds,omitempty"`

	// AndroidVersionIds: The ids of the set of Android OS version to be
	// used.
	// Use the EnvironmentDiscoveryService to get supported
	// options.
	// Required
	AndroidVersionIds []string `json:"androidVersionIds,omitempty"`

	// Locales: The set of locales the test device will enable for
	// testing.
	// Use the EnvironmentDiscoveryService to get supported
	// options.
	// Required
	Locales []string `json:"locales,omitempty"`

	// Orientations: The set of orientations to test with.
	// Use the EnvironmentDiscoveryService to get supported
	// options.
	// Required
	Orientations []string `json:"orientations,omitempty"`

	// ForceSendFields is a list of field names (e.g. "AndroidModelIds") to
	// unconditionally include in API requests. By default, fields with
	// empty values are omitted from API requests. However, any non-pointer,
	// non-interface field appearing in ForceSendFields will be sent to the
	// server regardless of whether the field is empty or not. This may be
	// used to include empty fields in Patch requests.
	ForceSendFields []string `json:"-"`

	// NullFields is a list of field names (e.g. "AndroidModelIds") to
	// include in API requests with the JSON null value. By default, fields
	// with empty values are omitted from API requests. However, any field
	// with an empty value appearing in NullFields will be sent to the
	// server as null. It is an error if a field in this list has a
	// non-empty value. This may be used to include null fields in Patch
	// requests.
	NullFields []string `json:"-"`
}

func (s *AndroidMatrix) MarshalJSON() ([]byte, error) {
	type NoMethod AndroidMatrix
	raw := NoMethod(*s)
	return gensupport.MarshalJSON(raw, s.ForceSendFields, s.NullFields)
}

// AndroidModel: A description of an Android device tests may be run on.
type AndroidModel struct {
	// Brand: The company that this device is branded with.
	// Example: "Google", "Samsung"
	// @OutputOnly
	Brand string `json:"brand,omitempty"`

	// Codename: The name of the industrial design.
	// This corresponds to android.os.Build.DEVICE
	// @OutputOnly
	Codename string `json:"codename,omitempty"`

	// Form: Whether this device is virtual or physical.
	// @OutputOnly
	//
	// Possible values:
	//   "DEVICE_FORM_UNSPECIFIED" - Do not use.  For proto versioning only.
	//   "VIRTUAL" - A software stack that simulates the device
	//   "PHYSICAL" - Actual hardware
	Form string `json:"form,omitempty"`

	// FormFactor: Whther this device is a phone, tablet, wearable,
	// etc.
	// @OutputOnly
	//
	// Possible values:
	//   "DEVICE_FORM_FACTOR_UNSPECIFIED" - Do not use. For proto versioning
	// only.
	//   "PHONE" - This device has the shape of a phone
	//   "TABLET" - This device has the shape of a tablet
	//   "WEARABLE" - This device has the shape of a watch or other wearable
	FormFactor string `json:"formFactor,omitempty"`

	// Id: The unique opaque id for this model.
	// Use this for invoking the TestExecutionService.
	// @OutputOnly
	Id string `json:"id,omitempty"`

	// Manufacturer: The manufacturer of this device.
	// @OutputOnly
	Manufacturer string `json:"manufacturer,omitempty"`

	// Name: The human-readable marketing name for this device
	// model.
	// Examples: "Nexus 5", "Galaxy S5"
	// @OutputOnly
	Name string `json:"name,omitempty"`

	// ScreenDensity: Screen density in DPI.
	// This corresponds to ro.sf.lcd_density
	// @OutputOnly
	ScreenDensity int64 `json:"screenDensity,omitempty"`

	// ScreenX: Screen size in the horizontal (X) dimension measured in
	// pixels.
	// @OutputOnly
	ScreenX int64 `json:"screenX,omitempty"`

	// ScreenY: Screen size in the vertical (Y) dimension measured in
	// pixels.
	// @OutputOnly
	ScreenY int64 `json:"screenY,omitempty"`

	// SupportedAbis: The list of supported ABIs for this device.
	// This corresponds to either android.os.Build.SUPPORTED_ABIS (for API
	// level
	// 21 and above) or android.os.Build.CPU_ABI/CPU_ABI2.
	// The most preferred ABI is the first element in the list.
	//
	// Elements are optionally prefixed by "version_id:" (where version_id
	// is
	// the id of an AndroidVersion), denoting an ABI that is supported only
	// on
	// a particular version.
	// @OutputOnly
	SupportedAbis []string `json:"supportedAbis,omitempty"`

	// SupportedVersionIds: The set of Android versions this device
	// supports.
	// @OutputOnly
	SupportedVersionIds []string `json:"supportedVersionIds,omitempty"`

	// Tags: Tags for this dimension.
	// Examples: "default", "preview", "deprecated"
	Tags []string `json:"tags,omitempty"`

	// VideoRecordingNotSupported: True if and only if tests with this model
	// DO NOT have video output.
	// See also TestSpecification.disable_video_recording
	// @OutputOnly
	VideoRecordingNotSupported bool `json:"videoRecordingNotSupported,omitempty"`

	// ForceSendFields is a list of field names (e.g. "Brand") to
	// unconditionally include in API requests. By default, fields with
	// empty values are omitted from API requests. However, any non-pointer,
	// non-interface field appearing in ForceSendFields will be sent to the
	// server regardless of whether the field is empty or not. This may be
	// used to include empty fields in Patch requests.
	ForceSendFields []string `json:"-"`

	// NullFields is a list of field names (e.g. "Brand") to include in API
	// requests with the JSON null value. By default, fields with empty
	// values are omitted from API requests. However, any field with an
	// empty value appearing in NullFields will be sent to the server as
	// null. It is an error if a field in this list has a non-empty value.
	// This may be used to include null fields in Patch requests.
	NullFields []string `json:"-"`
}

func (s *AndroidModel) MarshalJSON() ([]byte, error) {
	type NoMethod AndroidModel
	raw := NoMethod(*s)
	return gensupport.MarshalJSON(raw, s.ForceSendFields, s.NullFields)
}

// AndroidRoboTest: A test of an android application that explores the
// application on a virtual
// or physical Android Device, finding culprits and crashes as it goes.
type AndroidRoboTest struct {
	// AppApk: The APK for the application under test.
	AppApk *FileReference `json:"appApk,omitempty"`

	// AppInitialActivity: The initial activity that should be used to start
	// the app.
	// Optional
	AppInitialActivity string `json:"appInitialActivity,omitempty"`

	// AppPackageId: The java package for the application under
	// test.
	// Optional, default is determined by examining the application's
	// manifest.
	AppPackageId string `json:"appPackageId,omitempty"`

	// MaxDepth: The max depth of the traversal stack Robo can explore.
	// Needs to be at least
	// 2 to make Robo explore the app beyond the first activity.
	// Default is 50.
	// Optional
	MaxDepth int64 `json:"maxDepth,omitempty"`

	// MaxSteps: The max number of steps Robo can execute.
	// Default is no limit.
	// Optional
	MaxSteps int64 `json:"maxSteps,omitempty"`

	// RoboDirectives: A set of directives Robo should apply during the
	// crawl.
	// This allows users to customize the crawl. For example, the username
	// and
	// password for a test account can be provided.
	// Optional
	RoboDirectives []*RoboDirective `json:"roboDirectives,omitempty"`

	// RoboScript: A JSON file with a sequence of actions Robo should
	// perform as a prologue
	// for the crawl.
	// Optional
	RoboScript *FileReference `json:"roboScript,omitempty"`

	// StartingIntents: The intents used to launch the app for the crawl.
	// If none are provided, then the main launcher activity is launched.
	// If some are provided, then only those provided are launched (the
	// main
	// launcher activity must be provided explicitly).
	StartingIntents []*RoboStartingIntent `json:"startingIntents,omitempty"`

	// ForceSendFields is a list of field names (e.g. "AppApk") to
	// unconditionally include in API requests. By default, fields with
	// empty values are omitted from API requests. However, any non-pointer,
	// non-interface field appearing in ForceSendFields will be sent to the
	// server regardless of whether the field is empty or not. This may be
	// used to include empty fields in Patch requests.
	ForceSendFields []string `json:"-"`

	// NullFields is a list of field names (e.g. "AppApk") to include in API
	// requests with the JSON null value. By default, fields with empty
	// values are omitted from API requests. However, any field with an
	// empty value appearing in NullFields will be sent to the server as
	// null. It is an error if a field in this list has a non-empty value.
	// This may be used to include null fields in Patch requests.
	NullFields []string `json:"-"`
}

func (s *AndroidRoboTest) MarshalJSON() ([]byte, error) {
	type NoMethod AndroidRoboTest
	raw := NoMethod(*s)
	return gensupport.MarshalJSON(raw, s.ForceSendFields, s.NullFields)
}

// AndroidRuntimeConfiguration: Android configuration that can be
// selected at the time a test is run.
type AndroidRuntimeConfiguration struct {
	// Locales: The set of available locales.
	// @OutputOnly
	Locales []*Locale `json:"locales,omitempty"`

	// Orientations: The set of available orientations.
	// @OutputOnly
	Orientations []*Orientation `json:"orientations,omitempty"`

	// ForceSendFields is a list of field names (e.g. "Locales") to
	// unconditionally include in API requests. By default, fields with
	// empty values are omitted from API requests. However, any non-pointer,
	// non-interface field appearing in ForceSendFields will be sent to the
	// server regardless of whether the field is empty or not. This may be
	// used to include empty fields in Patch requests.
	ForceSendFields []string `json:"-"`

	// NullFields is a list of field names (e.g. "Locales") to include in
	// API requests with the JSON null value. By default, fields with empty
	// values are omitted from API requests. However, any field with an
	// empty value appearing in NullFields will be sent to the server as
	// null. It is an error if a field in this list has a non-empty value.
	// This may be used to include null fields in Patch requests.
	NullFields []string `json:"-"`
}

func (s *AndroidRuntimeConfiguration) MarshalJSON() ([]byte, error) {
	type NoMethod AndroidRuntimeConfiguration
	raw := NoMethod(*s)
	return gensupport.MarshalJSON(raw, s.ForceSendFields, s.NullFields)
}

// AndroidTestLoop: A test of an Android Application with a Test
// Loop.
// The intent <intent-name> will be implicitly added, since Games is the
// only
// user of this api, for the time being.
type AndroidTestLoop struct {
	// AppApk: The APK for the application under test.
	AppApk *FileReference `json:"appApk,omitempty"`

	// AppPackageId: The java package for the application under
	// test.
	// Optional, default is determined by examining the application's
	// manifest.
	AppPackageId string `json:"appPackageId,omitempty"`

	// ScenarioLabels: The list of scenario labels that should be run during
	// the test.
	// The scenario labels should map to labels defined in the
	// application's
	// manifest. For example, player_experience
	// and
	// com.google.test.loops.player_experience add all of the loops labeled
	// in the
	// manifest with the com.google.test.loops.player_experience name to
	// the
	// execution.
	// Optional. Scenarios can also be specified in the scenarios field.
	ScenarioLabels []string `json:"scenarioLabels,omitempty"`

	// Scenarios: The list of scenarios that should be run during the
	// test.
	// Optional, default is all test loops, derived from the
	// application's
	// manifest.
	Scenarios []int64 `json:"scenarios,omitempty"`

	// ForceSendFields is a list of field names (e.g. "AppApk") to
	// unconditionally include in API requests. By default, fields with
	// empty values are omitted from API requests. However, any non-pointer,
	// non-interface field appearing in ForceSendFields will be sent to the
	// server regardless of whether the field is empty or not. This may be
	// used to include empty fields in Patch requests.
	ForceSendFields []string `json:"-"`

	// NullFields is a list of field names (e.g. "AppApk") to include in API
	// requests with the JSON null value. By default, fields with empty
	// values are omitted from API requests. However, any field with an
	// empty value appearing in NullFields will be sent to the server as
	// null. It is an error if a field in this list has a non-empty value.
	// This may be used to include null fields in Patch requests.
	NullFields []string `json:"-"`
}

func (s *AndroidTestLoop) MarshalJSON() ([]byte, error) {
	type NoMethod AndroidTestLoop
	raw := NoMethod(*s)
	return gensupport.MarshalJSON(raw, s.ForceSendFields, s.NullFields)
}

// AndroidVersion: A version of the Android OS
type AndroidVersion struct {
	// ApiLevel: The API level for this Android version.
	// Examples: 18, 19
	// @OutputOnly
	ApiLevel int64 `json:"apiLevel,omitempty"`

	// CodeName: The code name for this Android version.
	// Examples: "JellyBean", "KitKat"
	// @OutputOnly
	CodeName string `json:"codeName,omitempty"`

	// Distribution: Market share for this version.
	// @OutputOnly
	Distribution *Distribution `json:"distribution,omitempty"`

	// Id: An opaque id for this Android version.
	// Use this id to invoke the TestExecutionService.
	// @OutputOnly
	Id string `json:"id,omitempty"`

	// ReleaseDate: The date this Android version became available in the
	// market.
	// @OutputOnly
	ReleaseDate *Date `json:"releaseDate,omitempty"`

	// Tags: Tags for this dimension.
	// Examples: "default", "preview", "deprecated"
	Tags []string `json:"tags,omitempty"`

	// VersionString: A string representing this version of the Android
	// OS.
	// Examples: "4.3", "4.4"
	// @OutputOnly
	VersionString string `json:"versionString,omitempty"`

	// ForceSendFields is a list of field names (e.g. "ApiLevel") to
	// unconditionally include in API requests. By default, fields with
	// empty values are omitted from API requests. However, any non-pointer,
	// non-interface field appearing in ForceSendFields will be sent to the
	// server regardless of whether the field is empty or not. This may be
	// used to include empty fields in Patch requests.
	ForceSendFields []string `json:"-"`

	// NullFields is a list of field names (e.g. "ApiLevel") to include in
	// API requests with the JSON null value. By default, fields with empty
	// values are omitted from API requests. However, any field with an
	// empty value appearing in NullFields will be sent to the server as
	// null. It is an error if a field in this list has a non-empty value.
	// This may be used to include null fields in Patch requests.
	NullFields []string `json:"-"`
}

func (s *AndroidVersion) MarshalJSON() ([]byte, error) {
	type NoMethod AndroidVersion
	raw := NoMethod(*s)
	return gensupport.MarshalJSON(raw, s.ForceSendFields, s.NullFields)
}

// Apk: An Android package file to install.
type Apk struct {
	// Location: The path to an APK to be installed on the device before the
	// test begins.
	// Optional
	Location *FileReference `json:"location,omitempty"`

	// PackageName: The java package for the APK to be installed.
	// Optional, value is determined by examining the application's
	// manifest.
	PackageName string `json:"packageName,omitempty"`

	// ForceSendFields is a list of field names (e.g. "Location") to
	// unconditionally include in API requests. By default, fields with
	// empty values are omitted from API requests. However, any non-pointer,
	// non-interface field appearing in ForceSendFields will be sent to the
	// server regardless of whether the field is empty or not. This may be
	// used to include empty fields in Patch requests.
	ForceSendFields []string `json:"-"`

	// NullFields is a list of field names (e.g. "Location") to include in
	// API requests with the JSON null value. By default, fields with empty
	// values are omitted from API requests. However, any field with an
	// empty value appearing in NullFields will be sent to the server as
	// null. It is an error if a field in this list has a non-empty value.
	// This may be used to include null fields in Patch requests.
	NullFields []string `json:"-"`
}

func (s *Apk) MarshalJSON() ([]byte, error) {
	type NoMethod Apk
	raw := NoMethod(*s)
	return gensupport.MarshalJSON(raw, s.ForceSendFields, s.NullFields)
}

// ApkDetail: Android application details based on application manifest
// and apk archive
// contents
type ApkDetail struct {
	ApkManifest *ApkManifest `json:"apkManifest,omitempty"`

	// ForceSendFields is a list of field names (e.g. "ApkManifest") to
	// unconditionally include in API requests. By default, fields with
	// empty values are omitted from API requests. However, any non-pointer,
	// non-interface field appearing in ForceSendFields will be sent to the
	// server regardless of whether the field is empty or not. This may be
	// used to include empty fields in Patch requests.
	ForceSendFields []string `json:"-"`

	// NullFields is a list of field names (e.g. "ApkManifest") to include
	// in API requests with the JSON null value. By default, fields with
	// empty values are omitted from API requests. However, any field with
	// an empty value appearing in NullFields will be sent to the server as
	// null. It is an error if a field in this list has a non-empty value.
	// This may be used to include null fields in Patch requests.
	NullFields []string `json:"-"`
}

func (s *ApkDetail) MarshalJSON() ([]byte, error) {
	type NoMethod ApkDetail
	raw := NoMethod(*s)
	return gensupport.MarshalJSON(raw, s.ForceSendFields, s.NullFields)
}

// ApkManifest: An Android app manifest.
// See
// http://developer.android.com/guide/topics/manifest/manifest-intro.
// html
type ApkManifest struct {
	// ApplicationLabel: User-readable name for the application.
	ApplicationLabel string `json:"applicationLabel,omitempty"`

	IntentFilters []*IntentFilter `json:"intentFilters,omitempty"`

	// MaxSdkVersion: Maximum API level on which the application is designed
	// to run.
	MaxSdkVersion int64 `json:"maxSdkVersion,omitempty"`

	// MinSdkVersion: Minimum API level required for the application to run.
	MinSdkVersion int64 `json:"minSdkVersion,omitempty"`

	// PackageName: Full Java-style package name for this application,
	// e.g.
	// "com.example.foo".
	PackageName string `json:"packageName,omitempty"`

	// ForceSendFields is a list of field names (e.g. "ApplicationLabel") to
	// unconditionally include in API requests. By default, fields with
	// empty values are omitted from API requests. However, any non-pointer,
	// non-interface field appearing in ForceSendFields will be sent to the
	// server regardless of whether the field is empty or not. This may be
	// used to include empty fields in Patch requests.
	ForceSendFields []string `json:"-"`

	// NullFields is a list of field names (e.g. "ApplicationLabel") to
	// include in API requests with the JSON null value. By default, fields
	// with empty values are omitted from API requests. However, any field
	// with an empty value appearing in NullFields will be sent to the
	// server as null. It is an error if a field in this list has a
	// non-empty value. This may be used to include null fields in Patch
	// requests.
	NullFields []string `json:"-"`
}

func (s *ApkManifest) MarshalJSON() ([]byte, error) {
	type NoMethod ApkManifest
	raw := NoMethod(*s)
	return gensupport.MarshalJSON(raw, s.ForceSendFields, s.NullFields)
}

// CancelTestMatrixResponse: Response containing the current state of
// the specified test matrix.
type CancelTestMatrixResponse struct {
	// TestState: The current rolled-up state of the test matrix.
	// If this state is already final, then the cancelation request
	// will
	// have no effect.
	//
	// Possible values:
	//   "TEST_STATE_UNSPECIFIED" - Do not use.  For proto versioning only.
	//   "VALIDATING" - The execution or matrix is being validated.
	//   "PENDING" - The execution or matrix is waiting for resources to
	// become available.
	//   "RUNNING" - The execution is currently being processed.
	//
	// Can only be set on an execution.
	//   "FINISHED" - The execution or matrix has terminated normally.
	//
	// On a matrix this means that the matrix level processing completed
	// normally,
	// but individual executions may be in an ERROR state.
	//   "ERROR" - The execution or matrix has stopped because it
	// encountered an
	// infrastructure failure.
	//   "UNSUPPORTED_ENVIRONMENT" - The execution was not run because it
	// corresponds to a unsupported
	// environment.
	//
	// Can only be set on an execution.
	//   "INCOMPATIBLE_ENVIRONMENT" - The execution was not run because the
	// provided inputs are incompatible with
	// the requested environment.
	//
	// Example: requested AndroidVersion is lower than APK's
	// minSdkVersion
	//
	// Can only be set on an execution.
	//   "INCOMPATIBLE_ARCHITECTURE" - The execution was not run because the
	// provided inputs are incompatible with
	// the requested architecture.
	//
	// Example: requested device does not support running the native code
	// in
	// the supplied APK
	//
	// Can only be set on an execution.
	//   "CANCELLED" - The user cancelled the execution.
	//
	// Can only be set on an execution.
	//   "INVALID" - The execution or matrix was not run because the
	// provided inputs are not
	// valid.
	//
	// Examples: input file is not of the expected type, is
	// malformed/corrupt, or
	// was flagged as malware
	TestState string `json:"testState,omitempty"`

	// ServerResponse contains the HTTP response code and headers from the
	// server.
	googleapi.ServerResponse `json:"-"`

	// ForceSendFields is a list of field names (e.g. "TestState") to
	// unconditionally include in API requests. By default, fields with
	// empty values are omitted from API requests. However, any non-pointer,
	// non-interface field appearing in ForceSendFields will be sent to the
	// server regardless of whether the field is empty or not. This may be
	// used to include empty fields in Patch requests.
	ForceSendFields []string `json:"-"`

	// NullFields is a list of field names (e.g. "TestState") to include in
	// API requests with the JSON null value. By default, fields with empty
	// values are omitted from API requests. However, any field with an
	// empty value appearing in NullFields will be sent to the server as
	// null. It is an error if a field in this list has a non-empty value.
	// This may be used to include null fields in Patch requests.
	NullFields []string `json:"-"`
}

func (s *CancelTestMatrixResponse) MarshalJSON() ([]byte, error) {
	type NoMethod CancelTestMatrixResponse
	raw := NoMethod(*s)
	return gensupport.MarshalJSON(raw, s.ForceSendFields, s.NullFields)
}

// ClientInfo: Information about the client which invoked the test.
type ClientInfo struct {
	// ClientInfoDetails: The list of detailed information about client.
	ClientInfoDetails []*ClientInfoDetail `json:"clientInfoDetails,omitempty"`

	// Name: Client name, such as gcloud.
	// Required
	Name string `json:"name,omitempty"`

	// ForceSendFields is a list of field names (e.g. "ClientInfoDetails")
	// to unconditionally include in API requests. By default, fields with
	// empty values are omitted from API requests. However, any non-pointer,
	// non-interface field appearing in ForceSendFields will be sent to the
	// server regardless of whether the field is empty or not. This may be
	// used to include empty fields in Patch requests.
	ForceSendFields []string `json:"-"`

	// NullFields is a list of field names (e.g. "ClientInfoDetails") to
	// include in API requests with the JSON null value. By default, fields
	// with empty values are omitted from API requests. However, any field
	// with an empty value appearing in NullFields will be sent to the
	// server as null. It is an error if a field in this list has a
	// non-empty value. This may be used to include null fields in Patch
	// requests.
	NullFields []string `json:"-"`
}

func (s *ClientInfo) MarshalJSON() ([]byte, error) {
	type NoMethod ClientInfo
	raw := NoMethod(*s)
	return gensupport.MarshalJSON(raw, s.ForceSendFields, s.NullFields)
}

// ClientInfoDetail: Key-value pair of detailed information about the
// client which invoked the
// test. For example {'Version', '1.0'}, {'Release Track', 'BETA'}
type ClientInfoDetail struct {
	// Key: The key of detailed client information.
	// Required
	Key string `json:"key,omitempty"`

	// Value: The value of detailed client information.
	// Required
	Value string `json:"value,omitempty"`

	// ForceSendFields is a list of field names (e.g. "Key") to
	// unconditionally include in API requests. By default, fields with
	// empty values are omitted from API requests. However, any non-pointer,
	// non-interface field appearing in ForceSendFields will be sent to the
	// server regardless of whether the field is empty or not. This may be
	// used to include empty fields in Patch requests.
	ForceSendFields []string `json:"-"`

	// NullFields is a list of field names (e.g. "Key") to include in API
	// requests with the JSON null value. By default, fields with empty
	// values are omitted from API requests. However, any field with an
	// empty value appearing in NullFields will be sent to the server as
	// null. It is an error if a field in this list has a non-empty value.
	// This may be used to include null fields in Patch requests.
	NullFields []string `json:"-"`
}

func (s *ClientInfoDetail) MarshalJSON() ([]byte, error) {
	type NoMethod ClientInfoDetail
	raw := NoMethod(*s)
	return gensupport.MarshalJSON(raw, s.ForceSendFields, s.NullFields)
}

// Date: Represents a whole calendar date, for example date of birth.
// The time of day
// and time zone are either specified elsewhere or are not significant.
// The date
// is relative to the Proleptic Gregorian Calendar. The day can be 0
// to
// represent a year and month where the day is not significant, for
// example
// credit card expiration date. The year can be 0 to represent a month
// and day
// independent of year, for example anniversary date. Related types
// are
// google.type.TimeOfDay and `google.protobuf.Timestamp`.
type Date struct {
	// Day: Day of month. Must be from 1 to 31 and valid for the year and
	// month, or 0
	// if specifying a year/month where the day is not significant.
	Day int64 `json:"day,omitempty"`

	// Month: Month of year. Must be from 1 to 12, or 0 if specifying a date
	// without a
	// month.
	Month int64 `json:"month,omitempty"`

	// Year: Year of date. Must be from 1 to 9999, or 0 if specifying a date
	// without
	// a year.
	Year int64 `json:"year,omitempty"`

	// ForceSendFields is a list of field names (e.g. "Day") to
	// unconditionally include in API requests. By default, fields with
	// empty values are omitted from API requests. However, any non-pointer,
	// non-interface field appearing in ForceSendFields will be sent to the
	// server regardless of whether the field is empty or not. This may be
	// used to include empty fields in Patch requests.
	ForceSendFields []string `json:"-"`

	// NullFields is a list of field names (e.g. "Day") to include in API
	// requests with the JSON null value. By default, fields with empty
	// values are omitted from API requests. However, any field with an
	// empty value appearing in NullFields will be sent to the server as
	// null. It is an error if a field in this list has a non-empty value.
	// This may be used to include null fields in Patch requests.
	NullFields []string `json:"-"`
}

func (s *Date) MarshalJSON() ([]byte, error) {
	type NoMethod Date
	raw := NoMethod(*s)
	return gensupport.MarshalJSON(raw, s.ForceSendFields, s.NullFields)
}

// DeviceFile: A single device file description.
type DeviceFile struct {
	// ObbFile: A reference to an opaque binary blob file
	ObbFile *ObbFile `json:"obbFile,omitempty"`

	// RegularFile: A reference to a regular file
	RegularFile *RegularFile `json:"regularFile,omitempty"`

	// ForceSendFields is a list of field names (e.g. "ObbFile") to
	// unconditionally include in API requests. By default, fields with
	// empty values are omitted from API requests. However, any non-pointer,
	// non-interface field appearing in ForceSendFields will be sent to the
	// server regardless of whether the field is empty or not. This may be
	// used to include empty fields in Patch requests.
	ForceSendFields []string `json:"-"`

	// NullFields is a list of field names (e.g. "ObbFile") to include in
	// API requests with the JSON null value. By default, fields with empty
	// values are omitted from API requests. However, any field with an
	// empty value appearing in NullFields will be sent to the server as
	// null. It is an error if a field in this list has a non-empty value.
	// This may be used to include null fields in Patch requests.
	NullFields []string `json:"-"`
}

func (s *DeviceFile) MarshalJSON() ([]byte, error) {
	type NoMethod DeviceFile
	raw := NoMethod(*s)
	return gensupport.MarshalJSON(raw, s.ForceSendFields, s.NullFields)
}

// Distribution: Data about the relative number of devices running
// a
// given configuration of the Android platform.
type Distribution struct {
	// MarketShare: The estimated fraction (0-1) of the total market with
	// this configuration.
	// @OutputOnly
	MarketShare float64 `json:"marketShare,omitempty"`

	// MeasurementTime: The time this distribution was measured.
	// @OutputOnly
	MeasurementTime string `json:"measurementTime,omitempty"`

	// ForceSendFields is a list of field names (e.g. "MarketShare") to
	// unconditionally include in API requests. By default, fields with
	// empty values are omitted from API requests. However, any non-pointer,
	// non-interface field appearing in ForceSendFields will be sent to the
	// server regardless of whether the field is empty or not. This may be
	// used to include empty fields in Patch requests.
	ForceSendFields []string `json:"-"`

	// NullFields is a list of field names (e.g. "MarketShare") to include
	// in API requests with the JSON null value. By default, fields with
	// empty values are omitted from API requests. However, any field with
	// an empty value appearing in NullFields will be sent to the server as
	// null. It is an error if a field in this list has a non-empty value.
	// This may be used to include null fields in Patch requests.
	NullFields []string `json:"-"`
}

func (s *Distribution) MarshalJSON() ([]byte, error) {
	type NoMethod Distribution
	raw := NoMethod(*s)
	return gensupport.MarshalJSON(raw, s.ForceSendFields, s.NullFields)
}

func (s *Distribution) UnmarshalJSON(data []byte) error {
	type NoMethod Distribution
	var s1 struct {
		MarketShare gensupport.JSONFloat64 `json:"marketShare"`
		*NoMethod
	}
	s1.NoMethod = (*NoMethod)(s)
	if err := json.Unmarshal(data, &s1); err != nil {
		return err
	}
	s.MarketShare = float64(s1.MarketShare)
	return nil
}

// Environment: The environment in which the test is run.
type Environment struct {
	// AndroidDevice: An Android device which must be used with an Android
	// test.
	AndroidDevice *AndroidDevice `json:"androidDevice,omitempty"`

	// IosDevice: An iOS device which must be used with an iOS test.
	IosDevice *IosDevice `json:"iosDevice,omitempty"`

	// ForceSendFields is a list of field names (e.g. "AndroidDevice") to
	// unconditionally include in API requests. By default, fields with
	// empty values are omitted from API requests. However, any non-pointer,
	// non-interface field appearing in ForceSendFields will be sent to the
	// server regardless of whether the field is empty or not. This may be
	// used to include empty fields in Patch requests.
	ForceSendFields []string `json:"-"`

	// NullFields is a list of field names (e.g. "AndroidDevice") to include
	// in API requests with the JSON null value. By default, fields with
	// empty values are omitted from API requests. However, any field with
	// an empty value appearing in NullFields will be sent to the server as
	// null. It is an error if a field in this list has a non-empty value.
	// This may be used to include null fields in Patch requests.
	NullFields []string `json:"-"`
}

func (s *Environment) MarshalJSON() ([]byte, error) {
	type NoMethod Environment
	raw := NoMethod(*s)
	return gensupport.MarshalJSON(raw, s.ForceSendFields, s.NullFields)
}

// EnvironmentMatrix: The matrix of environments in which the test is to
// be executed.
type EnvironmentMatrix struct {
	// AndroidDeviceList: A list of Android devices; the test will be run
	// only on the specified
	// devices.
	AndroidDeviceList *AndroidDeviceList `json:"androidDeviceList,omitempty"`

	// AndroidMatrix: A matrix of Android devices.
	AndroidMatrix *AndroidMatrix `json:"androidMatrix,omitempty"`

	// IosDeviceList: A list of iOS devices.
	IosDeviceList *IosDeviceList `json:"iosDeviceList,omitempty"`

	// ForceSendFields is a list of field names (e.g. "AndroidDeviceList")
	// to unconditionally include in API requests. By default, fields with
	// empty values are omitted from API requests. However, any non-pointer,
	// non-interface field appearing in ForceSendFields will be sent to the
	// server regardless of whether the field is empty or not. This may be
	// used to include empty fields in Patch requests.
	ForceSendFields []string `json:"-"`

	// NullFields is a list of field names (e.g. "AndroidDeviceList") to
	// include in API requests with the JSON null value. By default, fields
	// with empty values are omitted from API requests. However, any field
	// with an empty value appearing in NullFields will be sent to the
	// server as null. It is an error if a field in this list has a
	// non-empty value. This may be used to include null fields in Patch
	// requests.
	NullFields []string `json:"-"`
}

func (s *EnvironmentMatrix) MarshalJSON() ([]byte, error) {
	type NoMethod EnvironmentMatrix
	raw := NoMethod(*s)
	return gensupport.MarshalJSON(raw, s.ForceSendFields, s.NullFields)
}

// EnvironmentVariable: A key-value pair passed as an environment
// variable to the test
type EnvironmentVariable struct {
	// Key: Key for the environment variable
	Key string `json:"key,omitempty"`

	// Value: Value for the environment variable
	Value string `json:"value,omitempty"`

	// ForceSendFields is a list of field names (e.g. "Key") to
	// unconditionally include in API requests. By default, fields with
	// empty values are omitted from API requests. However, any non-pointer,
	// non-interface field appearing in ForceSendFields will be sent to the
	// server regardless of whether the field is empty or not. This may be
	// used to include empty fields in Patch requests.
	ForceSendFields []string `json:"-"`

	// NullFields is a list of field names (e.g. "Key") to include in API
	// requests with the JSON null value. By default, fields with empty
	// values are omitted from API requests. However, any field with an
	// empty value appearing in NullFields will be sent to the server as
	// null. It is an error if a field in this list has a non-empty value.
	// This may be used to include null fields in Patch requests.
	NullFields []string `json:"-"`
}

func (s *EnvironmentVariable) MarshalJSON() ([]byte, error) {
	type NoMethod EnvironmentVariable
	raw := NoMethod(*s)
	return gensupport.MarshalJSON(raw, s.ForceSendFields, s.NullFields)
}

// FileReference: A reference to a file, used for user inputs.
type FileReference struct {
	// GcsPath: A path to a file in Google Cloud Storage.
	// Example: gs://build-app-1414623860166/app-debug-unaligned.apk
	GcsPath string `json:"gcsPath,omitempty"`

	// ForceSendFields is a list of field names (e.g. "GcsPath") to
	// unconditionally include in API requests. By default, fields with
	// empty values are omitted from API requests. However, any non-pointer,
	// non-interface field appearing in ForceSendFields will be sent to the
	// server regardless of whether the field is empty or not. This may be
	// used to include empty fields in Patch requests.
	ForceSendFields []string `json:"-"`

	// NullFields is a list of field names (e.g. "GcsPath") to include in
	// API requests with the JSON null value. By default, fields with empty
	// values are omitted from API requests. However, any field with an
	// empty value appearing in NullFields will be sent to the server as
	// null. It is an error if a field in this list has a non-empty value.
	// This may be used to include null fields in Patch requests.
	NullFields []string `json:"-"`
}

func (s *FileReference) MarshalJSON() ([]byte, error) {
	type NoMethod FileReference
	raw := NoMethod(*s)
	return gensupport.MarshalJSON(raw, s.ForceSendFields, s.NullFields)
}

// GetApkDetailsResponse: Response containing the details of the
// specified Android application APK.
type GetApkDetailsResponse struct {
	// ApkDetail: Details of the Android APK.
	ApkDetail *ApkDetail `json:"apkDetail,omitempty"`

	// ServerResponse contains the HTTP response code and headers from the
	// server.
	googleapi.ServerResponse `json:"-"`

	// ForceSendFields is a list of field names (e.g. "ApkDetail") to
	// unconditionally include in API requests. By default, fields with
	// empty values are omitted from API requests. However, any non-pointer,
	// non-interface field appearing in ForceSendFields will be sent to the
	// server regardless of whether the field is empty or not. This may be
	// used to include empty fields in Patch requests.
	ForceSendFields []string `json:"-"`

	// NullFields is a list of field names (e.g. "ApkDetail") to include in
	// API requests with the JSON null value. By default, fields with empty
	// values are omitted from API requests. However, any field with an
	// empty value appearing in NullFields will be sent to the server as
	// null. It is an error if a field in this list has a non-empty value.
	// This may be used to include null fields in Patch requests.
	NullFields []string `json:"-"`
}

func (s *GetApkDetailsResponse) MarshalJSON() ([]byte, error) {
	type NoMethod GetApkDetailsResponse
	raw := NoMethod(*s)
	return gensupport.MarshalJSON(raw, s.ForceSendFields, s.NullFields)
}

// GoogleAuto: Enables automatic Google account login.
// If set, the service will automatically generate a Google test account
// and add
// it to the device, before executing the test. Note that test accounts
// might be
// reused.
// Many applications show their full set of functionalities when an
// account is
// present on the device. Logging into the device with these generated
// accounts
// allows testing more functionalities.
type GoogleAuto struct {
}

// GoogleCloudStorage: A storage location within Google cloud storage
// (GCS).
type GoogleCloudStorage struct {
	// GcsPath: The path to a directory in GCS that will
	// eventually contain the results for this test.
	// The requesting user must have write access on the bucket in the
	// supplied
	// path.
	// Required
	GcsPath string `json:"gcsPath,omitempty"`

	// ForceSendFields is a list of field names (e.g. "GcsPath") to
	// unconditionally include in API requests. By default, fields with
	// empty values are omitted from API requests. However, any non-pointer,
	// non-interface field appearing in ForceSendFields will be sent to the
	// server regardless of whether the field is empty or not. This may be
	// used to include empty fields in Patch requests.
	ForceSendFields []string `json:"-"`

	// NullFields is a list of field names (e.g. "GcsPath") to include in
	// API requests with the JSON null value. By default, fields with empty
	// values are omitted from API requests. However, any field with an
	// empty value appearing in NullFields will be sent to the server as
	// null. It is an error if a field in this list has a non-empty value.
	// This may be used to include null fields in Patch requests.
	NullFields []string `json:"-"`
}

func (s *GoogleCloudStorage) MarshalJSON() ([]byte, error) {
	type NoMethod GoogleCloudStorage
	raw := NoMethod(*s)
	return gensupport.MarshalJSON(raw, s.ForceSendFields, s.NullFields)
}

// IntentFilter: The <intent-filter> section of an <activity>
// tag.
// https://developer.android.com/guide/topics/manifest/intent-filter
// -element.html
type IntentFilter struct {
	// ActionNames: The android:name value of the <action> tag
	ActionNames []string `json:"actionNames,omitempty"`

	// CategoryNames: The android:name value of the <category> tag
	CategoryNames []string `json:"categoryNames,omitempty"`

	// MimeType: The android:mimeType value of the <data> tag
	MimeType string `json:"mimeType,omitempty"`

	// ForceSendFields is a list of field names (e.g. "ActionNames") to
	// unconditionally include in API requests. By default, fields with
	// empty values are omitted from API requests. However, any non-pointer,
	// non-interface field appearing in ForceSendFields will be sent to the
	// server regardless of whether the field is empty or not. This may be
	// used to include empty fields in Patch requests.
	ForceSendFields []string `json:"-"`

	// NullFields is a list of field names (e.g. "ActionNames") to include
	// in API requests with the JSON null value. By default, fields with
	// empty values are omitted from API requests. However, any field with
	// an empty value appearing in NullFields will be sent to the server as
	// null. It is an error if a field in this list has a non-empty value.
	// This may be used to include null fields in Patch requests.
	NullFields []string `json:"-"`
}

func (s *IntentFilter) MarshalJSON() ([]byte, error) {
	type NoMethod IntentFilter
	raw := NoMethod(*s)
	return gensupport.MarshalJSON(raw, s.ForceSendFields, s.NullFields)
}

// IosDevice: A single iOS device.
type IosDevice struct {
	// IosModelId: Required. The id of the iOS device to be used.
	// Use the EnvironmentDiscoveryService to get supported options.
	IosModelId string `json:"iosModelId,omitempty"`

	// IosVersionId: Required. The id of the iOS major software version to
	// be used.
	// Use the EnvironmentDiscoveryService to get supported options.
	IosVersionId string `json:"iosVersionId,omitempty"`

	// Locale: Required. The locale the test device used for testing.
	// Use the EnvironmentDiscoveryService to get supported options.
	Locale string `json:"locale,omitempty"`

	// Orientation: Required. How the device is oriented during the
	// test.
	// Use the EnvironmentDiscoveryService to get supported options.
	Orientation string `json:"orientation,omitempty"`

	// ForceSendFields is a list of field names (e.g. "IosModelId") to
	// unconditionally include in API requests. By default, fields with
	// empty values are omitted from API requests. However, any non-pointer,
	// non-interface field appearing in ForceSendFields will be sent to the
	// server regardless of whether the field is empty or not. This may be
	// used to include empty fields in Patch requests.
	ForceSendFields []string `json:"-"`

	// NullFields is a list of field names (e.g. "IosModelId") to include in
	// API requests with the JSON null value. By default, fields with empty
	// values are omitted from API requests. However, any field with an
	// empty value appearing in NullFields will be sent to the server as
	// null. It is an error if a field in this list has a non-empty value.
	// This may be used to include null fields in Patch requests.
	NullFields []string `json:"-"`
}

func (s *IosDevice) MarshalJSON() ([]byte, error) {
	type NoMethod IosDevice
	raw := NoMethod(*s)
	return gensupport.MarshalJSON(raw, s.ForceSendFields, s.NullFields)
}

// IosDeviceCatalog: The currently supported iOS devices.
type IosDeviceCatalog struct {
	// Models: Output only. The set of supported iOS device models.
	Models []*IosModel `json:"models,omitempty"`

	// RuntimeConfiguration: Output only. The set of supported runtime
	// configurations.
	RuntimeConfiguration *IosRuntimeConfiguration `json:"runtimeConfiguration,omitempty"`

	// Versions: Output only. The set of supported iOS software versions.
	Versions []*IosVersion `json:"versions,omitempty"`

	// ForceSendFields is a list of field names (e.g. "Models") to
	// unconditionally include in API requests. By default, fields with
	// empty values are omitted from API requests. However, any non-pointer,
	// non-interface field appearing in ForceSendFields will be sent to the
	// server regardless of whether the field is empty or not. This may be
	// used to include empty fields in Patch requests.
	ForceSendFields []string `json:"-"`

	// NullFields is a list of field names (e.g. "Models") to include in API
	// requests with the JSON null value. By default, fields with empty
	// values are omitted from API requests. However, any field with an
	// empty value appearing in NullFields will be sent to the server as
	// null. It is an error if a field in this list has a non-empty value.
	// This may be used to include null fields in Patch requests.
	NullFields []string `json:"-"`
}

func (s *IosDeviceCatalog) MarshalJSON() ([]byte, error) {
	type NoMethod IosDeviceCatalog
	raw := NoMethod(*s)
	return gensupport.MarshalJSON(raw, s.ForceSendFields, s.NullFields)
}

// IosDeviceList: A list of iOS device configurations in which the test
// is to be executed.
type IosDeviceList struct {
	// IosDevices: Required. A list of iOS devices
	IosDevices []*IosDevice `json:"iosDevices,omitempty"`

	// ForceSendFields is a list of field names (e.g. "IosDevices") to
	// unconditionally include in API requests. By default, fields with
	// empty values are omitted from API requests. However, any non-pointer,
	// non-interface field appearing in ForceSendFields will be sent to the
	// server regardless of whether the field is empty or not. This may be
	// used to include empty fields in Patch requests.
	ForceSendFields []string `json:"-"`

	// NullFields is a list of field names (e.g. "IosDevices") to include in
	// API requests with the JSON null value. By default, fields with empty
	// values are omitted from API requests. However, any field with an
	// empty value appearing in NullFields will be sent to the server as
	// null. It is an error if a field in this list has a non-empty value.
	// This may be used to include null fields in Patch requests.
	NullFields []string `json:"-"`
}

func (s *IosDeviceList) MarshalJSON() ([]byte, error) {
	type NoMethod IosDeviceList
	raw := NoMethod(*s)
	return gensupport.MarshalJSON(raw, s.ForceSendFields, s.NullFields)
}

// IosModel: A description of an iOS device tests may be run on.
type IosModel struct {
	// DeviceCapabilities: Output only. Device capabilities.
	// Copied
	// from
	// https://developer.apple.com/library/archive/documentation/DeviceI
	// nformation/Reference/iOSDeviceCompatibility/DeviceCompatibilityMatrix/
	// DeviceCompatibilityMatrix.html
	DeviceCapabilities []string `json:"deviceCapabilities,omitempty"`

	// Id: Output only. The unique opaque id for this model.
	// Use this for invoking the TestExecutionService.
	Id string `json:"id,omitempty"`

	// Name: Output only. The human-readable name for this device
	// model.
	// Examples: "iPhone 4s", "iPad Mini 2"
	Name string `json:"name,omitempty"`

	// SupportedVersionIds: Output only. The set of iOS major software
	// versions this device supports.
	SupportedVersionIds []string `json:"supportedVersionIds,omitempty"`

	// Tags: Output only. Tags for this dimension.
	// Examples: "default", "preview", "deprecated"
	Tags []string `json:"tags,omitempty"`

	// ForceSendFields is a list of field names (e.g. "DeviceCapabilities")
	// to unconditionally include in API requests. By default, fields with
	// empty values are omitted from API requests. However, any non-pointer,
	// non-interface field appearing in ForceSendFields will be sent to the
	// server regardless of whether the field is empty or not. This may be
	// used to include empty fields in Patch requests.
	ForceSendFields []string `json:"-"`

	// NullFields is a list of field names (e.g. "DeviceCapabilities") to
	// include in API requests with the JSON null value. By default, fields
	// with empty values are omitted from API requests. However, any field
	// with an empty value appearing in NullFields will be sent to the
	// server as null. It is an error if a field in this list has a
	// non-empty value. This may be used to include null fields in Patch
	// requests.
	NullFields []string `json:"-"`
}

func (s *IosModel) MarshalJSON() ([]byte, error) {
	type NoMethod IosModel
	raw := NoMethod(*s)
	return gensupport.MarshalJSON(raw, s.ForceSendFields, s.NullFields)
}

// IosRuntimeConfiguration: iOS configuration that can be selected at
// the time a test is run.
type IosRuntimeConfiguration struct {
	// Locales: Output only. The set of available locales.
	Locales []*Locale `json:"locales,omitempty"`

	// Orientations: Output only. The set of available orientations.
	Orientations []*Orientation `json:"orientations,omitempty"`

	// ForceSendFields is a list of field names (e.g. "Locales") to
	// unconditionally include in API requests. By default, fields with
	// empty values are omitted from API requests. However, any non-pointer,
	// non-interface field appearing in ForceSendFields will be sent to the
	// server regardless of whether the field is empty or not. This may be
	// used to include empty fields in Patch requests.
	ForceSendFields []string `json:"-"`

	// NullFields is a list of field names (e.g. "Locales") to include in
	// API requests with the JSON null value. By default, fields with empty
	// values are omitted from API requests. However, any field with an
	// empty value appearing in NullFields will be sent to the server as
	// null. It is an error if a field in this list has a non-empty value.
	// This may be used to include null fields in Patch requests.
	NullFields []string `json:"-"`
}

func (s *IosRuntimeConfiguration) MarshalJSON() ([]byte, error) {
	type NoMethod IosRuntimeConfiguration
	raw := NoMethod(*s)
	return gensupport.MarshalJSON(raw, s.ForceSendFields, s.NullFields)
}

// IosTestSetup: A description of how to set up an iOS device prior to a
// test.
type IosTestSetup struct {
	// NetworkProfile: Optional. The network traffic profile used for
	// running the test.
	// Available network profiles can be queried by using
	// the
	// NETWORK_CONFIGURATION environment type when
	// calling
	// TestEnvironmentDiscoveryService.GetTestEnvironmentCatalog.
	NetworkProfile string `json:"networkProfile,omitempty"`

	// ForceSendFields is a list of field names (e.g. "NetworkProfile") to
	// unconditionally include in API requests. By default, fields with
	// empty values are omitted from API requests. However, any non-pointer,
	// non-interface field appearing in ForceSendFields will be sent to the
	// server regardless of whether the field is empty or not. This may be
	// used to include empty fields in Patch requests.
	ForceSendFields []string `json:"-"`

	// NullFields is a list of field names (e.g. "NetworkProfile") to
	// include in API requests with the JSON null value. By default, fields
	// with empty values are omitted from API requests. However, any field
	// with an empty value appearing in NullFields will be sent to the
	// server as null. It is an error if a field in this list has a
	// non-empty value. This may be used to include null fields in Patch
	// requests.
	NullFields []string `json:"-"`
}

func (s *IosTestSetup) MarshalJSON() ([]byte, error) {
	type NoMethod IosTestSetup
	raw := NoMethod(*s)
	return gensupport.MarshalJSON(raw, s.ForceSendFields, s.NullFields)
}

// IosVersion: An iOS version
type IosVersion struct {
	// Id: Output only. An opaque id for this iOS version.
	// Use this id to invoke the TestExecutionService.
	Id string `json:"id,omitempty"`

	// MajorVersion: Output only. An integer representing the major iOS
	// version.
	// Examples: "8", "9"
	MajorVersion int64 `json:"majorVersion,omitempty"`

	// MinorVersion: Output only. An integer representing the minor iOS
	// version.
	// Examples: "1", "2"
	MinorVersion int64 `json:"minorVersion,omitempty"`

	// Tags: Output only. Tags for this dimension.
	// Examples: "default", "preview", "deprecated"
	Tags []string `json:"tags,omitempty"`

	// ForceSendFields is a list of field names (e.g. "Id") to
	// unconditionally include in API requests. By default, fields with
	// empty values are omitted from API requests. However, any non-pointer,
	// non-interface field appearing in ForceSendFields will be sent to the
	// server regardless of whether the field is empty or not. This may be
	// used to include empty fields in Patch requests.
	ForceSendFields []string `json:"-"`

	// NullFields is a list of field names (e.g. "Id") to include in API
	// requests with the JSON null value. By default, fields with empty
	// values are omitted from API requests. However, any field with an
	// empty value appearing in NullFields will be sent to the server as
	// null. It is an error if a field in this list has a non-empty value.
	// This may be used to include null fields in Patch requests.
	NullFields []string `json:"-"`
}

func (s *IosVersion) MarshalJSON() ([]byte, error) {
	type NoMethod IosVersion
	raw := NoMethod(*s)
	return gensupport.MarshalJSON(raw, s.ForceSendFields, s.NullFields)
}

// IosXcTest: A test of an iOS application that uses the XCTest
// framework.
// Xcode supports the option to "build for testing", which generates
// an
// .xctestrun file that contains a test specification (arguments, test
// methods,
// etc). This test type accepts a zip file containing the .xctestrun
// file and
// the corresponding contents of the Build/Products directory that
// contains all
// the binaries needed to run the tests.
type IosXcTest struct {
	// TestsZip: Required. The .zip containing the .xctestrun file and the
	// contents of the
	// DerivedData/Build/Products directory.
	// The .xctestrun file in this zip is ignored if the xctestrun field
	// is
	// specified.
	TestsZip *FileReference `json:"testsZip,omitempty"`

	// Xctestrun: Optional. An .xctestrun file that will override the
	// .xctestrun file in the
	// tests zip. Because the .xctestrun file contains environment variables
	// along
	// with test methods to run and/or ignore, this can be useful for
	// sharding
	// tests. Default is taken from the tests zip.
	Xctestrun *FileReference `json:"xctestrun,omitempty"`

	// ForceSendFields is a list of field names (e.g. "TestsZip") to
	// unconditionally include in API requests. By default, fields with
	// empty values are omitted from API requests. However, any non-pointer,
	// non-interface field appearing in ForceSendFields will be sent to the
	// server regardless of whether the field is empty or not. This may be
	// used to include empty fields in Patch requests.
	ForceSendFields []string `json:"-"`

	// NullFields is a list of field names (e.g. "TestsZip") to include in
	// API requests with the JSON null value. By default, fields with empty
	// values are omitted from API requests. However, any field with an
	// empty value appearing in NullFields will be sent to the server as
	// null. It is an error if a field in this list has a non-empty value.
	// This may be used to include null fields in Patch requests.
	NullFields []string `json:"-"`
}

func (s *IosXcTest) MarshalJSON() ([]byte, error) {
	type NoMethod IosXcTest
	raw := NoMethod(*s)
	return gensupport.MarshalJSON(raw, s.ForceSendFields, s.NullFields)
}

// LauncherActivityIntent: Specifies an intent that starts the main
// launcher activity.
type LauncherActivityIntent struct {
}

// Locale: A location/region designation for language.
type Locale struct {
	// Id: The id for this locale.
	// Example: "en_US"
	// @OutputOnly
	Id string `json:"id,omitempty"`

	// Name: A human-friendly name for this language/locale.
	// Example: "English"
	// @OutputOnly
	Name string `json:"name,omitempty"`

	// Region: A human-friendly string representing the region for this
	// locale.
	// Example: "United States"
	// Not present for every locale.
	// @OutputOnly
	Region string `json:"region,omitempty"`

	// Tags: Tags for this dimension.
	// Examples: "default"
	Tags []string `json:"tags,omitempty"`

	// ForceSendFields is a list of field names (e.g. "Id") to
	// unconditionally include in API requests. By default, fields with
	// empty values are omitted from API requests. However, any non-pointer,
	// non-interface field appearing in ForceSendFields will be sent to the
	// server regardless of whether the field is empty or not. This may be
	// used to include empty fields in Patch requests.
	ForceSendFields []string `json:"-"`

	// NullFields is a list of field names (e.g. "Id") to include in API
	// requests with the JSON null value. By default, fields with empty
	// values are omitted from API requests. However, any field with an
	// empty value appearing in NullFields will be sent to the server as
	// null. It is an error if a field in this list has a non-empty value.
	// This may be used to include null fields in Patch requests.
	NullFields []string `json:"-"`
}

func (s *Locale) MarshalJSON() ([]byte, error) {
	type NoMethod Locale
	raw := NoMethod(*s)
	return gensupport.MarshalJSON(raw, s.ForceSendFields, s.NullFields)
}

type NetworkConfiguration struct {
	// DownRule: The emulation rule applying to the download traffic
	DownRule *TrafficRule `json:"downRule,omitempty"`

	// Id: The unique opaque id for this network traffic
	// configuration
	// @OutputOnly
	Id string `json:"id,omitempty"`

	// UpRule: The emulation rule applying to the upload traffic
	UpRule *TrafficRule `json:"upRule,omitempty"`

	// ForceSendFields is a list of field names (e.g. "DownRule") to
	// unconditionally include in API requests. By default, fields with
	// empty values are omitted from API requests. However, any non-pointer,
	// non-interface field appearing in ForceSendFields will be sent to the
	// server regardless of whether the field is empty or not. This may be
	// used to include empty fields in Patch requests.
	ForceSendFields []string `json:"-"`

	// NullFields is a list of field names (e.g. "DownRule") to include in
	// API requests with the JSON null value. By default, fields with empty
	// values are omitted from API requests. However, any field with an
	// empty value appearing in NullFields will be sent to the server as
	// null. It is an error if a field in this list has a non-empty value.
	// This may be used to include null fields in Patch requests.
	NullFields []string `json:"-"`
}

func (s *NetworkConfiguration) MarshalJSON() ([]byte, error) {
	type NoMethod NetworkConfiguration
	raw := NoMethod(*s)
	return gensupport.MarshalJSON(raw, s.ForceSendFields, s.NullFields)
}

type NetworkConfigurationCatalog struct {
	Configurations []*NetworkConfiguration `json:"configurations,omitempty"`

	// ForceSendFields is a list of field names (e.g. "Configurations") to
	// unconditionally include in API requests. By default, fields with
	// empty values are omitted from API requests. However, any non-pointer,
	// non-interface field appearing in ForceSendFields will be sent to the
	// server regardless of whether the field is empty or not. This may be
	// used to include empty fields in Patch requests.
	ForceSendFields []string `json:"-"`

	// NullFields is a list of field names (e.g. "Configurations") to
	// include in API requests with the JSON null value. By default, fields
	// with empty values are omitted from API requests. However, any field
	// with an empty value appearing in NullFields will be sent to the
	// server as null. It is an error if a field in this list has a
	// non-empty value. This may be used to include null fields in Patch
	// requests.
	NullFields []string `json:"-"`
}

func (s *NetworkConfigurationCatalog) MarshalJSON() ([]byte, error) {
	type NoMethod NetworkConfigurationCatalog
	raw := NoMethod(*s)
	return gensupport.MarshalJSON(raw, s.ForceSendFields, s.NullFields)
}

// ObbFile: An opaque binary blob file to install on the device before
// the test starts
type ObbFile struct {
	// Obb: Opaque Binary Blob (OBB) file(s) to install on the
	// device
	// Required
	Obb *FileReference `json:"obb,omitempty"`

	// ObbFileName: OBB file name which must conform to the format as
	// specified by
	// Android
	// e.g. [main|patch].0300110.com.example.android.obb
	// which will be installed into
	//   <shared-storage>/Android/obb/<package-name>/
	// on the device
	// Required
	ObbFileName string `json:"obbFileName,omitempty"`

	// ForceSendFields is a list of field names (e.g. "Obb") to
	// unconditionally include in API requests. By default, fields with
	// empty values are omitted from API requests. However, any non-pointer,
	// non-interface field appearing in ForceSendFields will be sent to the
	// server regardless of whether the field is empty or not. This may be
	// used to include empty fields in Patch requests.
	ForceSendFields []string `json:"-"`

	// NullFields is a list of field names (e.g. "Obb") to include in API
	// requests with the JSON null value. By default, fields with empty
	// values are omitted from API requests. However, any field with an
	// empty value appearing in NullFields will be sent to the server as
	// null. It is an error if a field in this list has a non-empty value.
	// This may be used to include null fields in Patch requests.
	NullFields []string `json:"-"`
}

func (s *ObbFile) MarshalJSON() ([]byte, error) {
	type NoMethod ObbFile
	raw := NoMethod(*s)
	return gensupport.MarshalJSON(raw, s.ForceSendFields, s.NullFields)
}

// Orientation: Screen orientation of the device.
type Orientation struct {
	// Id: The id for this orientation.
	// Example: "portrait"
	// @OutputOnly
	Id string `json:"id,omitempty"`

	// Name: A human-friendly name for this orientation.
	// Example: "portrait"
	// @OutputOnly
	Name string `json:"name,omitempty"`

	// Tags: Tags for this dimension.
	// Examples: "default"
	Tags []string `json:"tags,omitempty"`

	// ForceSendFields is a list of field names (e.g. "Id") to
	// unconditionally include in API requests. By default, fields with
	// empty values are omitted from API requests. However, any non-pointer,
	// non-interface field appearing in ForceSendFields will be sent to the
	// server regardless of whether the field is empty or not. This may be
	// used to include empty fields in Patch requests.
	ForceSendFields []string `json:"-"`

	// NullFields is a list of field names (e.g. "Id") to include in API
	// requests with the JSON null value. By default, fields with empty
	// values are omitted from API requests. However, any field with an
	// empty value appearing in NullFields will be sent to the server as
	// null. It is an error if a field in this list has a non-empty value.
	// This may be used to include null fields in Patch requests.
	NullFields []string `json:"-"`
}

func (s *Orientation) MarshalJSON() ([]byte, error) {
	type NoMethod Orientation
	raw := NoMethod(*s)
	return gensupport.MarshalJSON(raw, s.ForceSendFields, s.NullFields)
}

// ProvidedSoftwareCatalog: The currently provided software environment
// on the devices under test.
type ProvidedSoftwareCatalog struct {
	// OrchestratorVersion: A string representing the current version of
	// Android Test Orchestrator that
	// is provided by TestExecutionService. Example: "1.0.2 beta"
	OrchestratorVersion string `json:"orchestratorVersion,omitempty"`

	// ForceSendFields is a list of field names (e.g. "OrchestratorVersion")
	// to unconditionally include in API requests. By default, fields with
	// empty values are omitted from API requests. However, any non-pointer,
	// non-interface field appearing in ForceSendFields will be sent to the
	// server regardless of whether the field is empty or not. This may be
	// used to include empty fields in Patch requests.
	ForceSendFields []string `json:"-"`

	// NullFields is a list of field names (e.g. "OrchestratorVersion") to
	// include in API requests with the JSON null value. By default, fields
	// with empty values are omitted from API requests. However, any field
	// with an empty value appearing in NullFields will be sent to the
	// server as null. It is an error if a field in this list has a
	// non-empty value. This may be used to include null fields in Patch
	// requests.
	NullFields []string `json:"-"`
}

func (s *ProvidedSoftwareCatalog) MarshalJSON() ([]byte, error) {
	type NoMethod ProvidedSoftwareCatalog
	raw := NoMethod(*s)
	return gensupport.MarshalJSON(raw, s.ForceSendFields, s.NullFields)
}

// RegularFile: A file or directory to install on the device before the
// test starts
type RegularFile struct {
	// Content: Required
	Content *FileReference `json:"content,omitempty"`

	// DevicePath: Where to put the content on the device. Must be an
	// absolute, whitelisted
	// path. If the file exists, it will be replaced.
	// The following device-side directories and any of their subdirectories
	// are
	// whitelisted:
	// <p>${EXTERNAL_STORAGE}, or /sdcard</p>
	// <p>${ANDROID_DATA}/local/tmp, or /data/local/tmp</p>
	// <p>Specifying a path outside of these directory trees is
	// invalid.
	//
	// <p> The paths /sdcard and /data will be made available and treated
	// as
	// implicit path substitutions. E.g. if /sdcard on a particular device
	// does
	// not map to external storage, the system will replace it with the
	// external
	// storage path prefix for that device and copy the file there.
	//
	// <p> It is strongly advised to use the <a
	// href=
	// "http://developer.android.com/reference/android/os/Environment.h
	// tml">
	// Environment API</a> in app and test code to access files on the
	// device in a
	// portable way.
	// Required
	DevicePath string `json:"devicePath,omitempty"`

	// ForceSendFields is a list of field names (e.g. "Content") to
	// unconditionally include in API requests. By default, fields with
	// empty values are omitted from API requests. However, any non-pointer,
	// non-interface field appearing in ForceSendFields will be sent to the
	// server regardless of whether the field is empty or not. This may be
	// used to include empty fields in Patch requests.
	ForceSendFields []string `json:"-"`

	// NullFields is a list of field names (e.g. "Content") to include in
	// API requests with the JSON null value. By default, fields with empty
	// values are omitted from API requests. However, any field with an
	// empty value appearing in NullFields will be sent to the server as
	// null. It is an error if a field in this list has a non-empty value.
	// This may be used to include null fields in Patch requests.
	NullFields []string `json:"-"`
}

func (s *RegularFile) MarshalJSON() ([]byte, error) {
	type NoMethod RegularFile
	raw := NoMethod(*s)
	return gensupport.MarshalJSON(raw, s.ForceSendFields, s.NullFields)
}

// ResultStorage: Locations where the results of running the test are
// stored.
type ResultStorage struct {
	// GoogleCloudStorage: Required.
	GoogleCloudStorage *GoogleCloudStorage `json:"googleCloudStorage,omitempty"`

	// ToolResultsExecution: The tool results execution that results are
	// written to.
	// @OutputOnly
	ToolResultsExecution *ToolResultsExecution `json:"toolResultsExecution,omitempty"`

	// ToolResultsHistory: The tool results history that contains the tool
	// results execution that
	// results are written to.
	//
	// Optional, if not provided the service will choose an appropriate
	// value.
	ToolResultsHistory *ToolResultsHistory `json:"toolResultsHistory,omitempty"`

	// ForceSendFields is a list of field names (e.g. "GoogleCloudStorage")
	// to unconditionally include in API requests. By default, fields with
	// empty values are omitted from API requests. However, any non-pointer,
	// non-interface field appearing in ForceSendFields will be sent to the
	// server regardless of whether the field is empty or not. This may be
	// used to include empty fields in Patch requests.
	ForceSendFields []string `json:"-"`

	// NullFields is a list of field names (e.g. "GoogleCloudStorage") to
	// include in API requests with the JSON null value. By default, fields
	// with empty values are omitted from API requests. However, any field
	// with an empty value appearing in NullFields will be sent to the
	// server as null. It is an error if a field in this list has a
	// non-empty value. This may be used to include null fields in Patch
	// requests.
	NullFields []string `json:"-"`
}

func (s *ResultStorage) MarshalJSON() ([]byte, error) {
	type NoMethod ResultStorage
	raw := NoMethod(*s)
	return gensupport.MarshalJSON(raw, s.ForceSendFields, s.NullFields)
}

// RoboDirective: Directs Robo to interact with a specific UI element if
// it is encountered
// during the crawl. Currently, Robo can perform text entry or element
// click.
type RoboDirective struct {
	// ActionType: The type of action that Robo should perform on the
	// specified element.
	// Required.
	//
	// Possible values:
	//   "ACTION_TYPE_UNSPECIFIED" - DO NOT USE. For proto versioning only.
	//   "SINGLE_CLICK" - Direct Robo to click on the specified element.
	// No-op if specified element
	// is not clickable.
	//   "ENTER_TEXT" - Direct Robo to enter text on the specified element.
	// No-op if specified
	// element is not enabled or does not allow text entry.
	ActionType string `json:"actionType,omitempty"`

	// InputText: The text that Robo is directed to set. If left empty, the
	// directive will be
	// treated as a CLICK on the element matching the
	// resource_name.
	// Optional
	InputText string `json:"inputText,omitempty"`

	// ResourceName: The android resource name of the target UI element
	// For example,
	//    in Java: R.string.foo
	//    in xml: @string/foo
	// Only the “foo” part is needed.
	// Reference
	// doc:
	// https://developer.android.com/guide/topics/resources/accessing-re
	// sources.html
	// Required
	ResourceName string `json:"resourceName,omitempty"`

	// ForceSendFields is a list of field names (e.g. "ActionType") to
	// unconditionally include in API requests. By default, fields with
	// empty values are omitted from API requests. However, any non-pointer,
	// non-interface field appearing in ForceSendFields will be sent to the
	// server regardless of whether the field is empty or not. This may be
	// used to include empty fields in Patch requests.
	ForceSendFields []string `json:"-"`

	// NullFields is a list of field names (e.g. "ActionType") to include in
	// API requests with the JSON null value. By default, fields with empty
	// values are omitted from API requests. However, any field with an
	// empty value appearing in NullFields will be sent to the server as
	// null. It is an error if a field in this list has a non-empty value.
	// This may be used to include null fields in Patch requests.
	NullFields []string `json:"-"`
}

func (s *RoboDirective) MarshalJSON() ([]byte, error) {
	type NoMethod RoboDirective
	raw := NoMethod(*s)
	return gensupport.MarshalJSON(raw, s.ForceSendFields, s.NullFields)
}

// RoboStartingIntent: Message for specifying the start activities to
// crawl
type RoboStartingIntent struct {
	LauncherActivity *LauncherActivityIntent `json:"launcherActivity,omitempty"`

	StartActivity *StartActivityIntent `json:"startActivity,omitempty"`

	// ForceSendFields is a list of field names (e.g. "LauncherActivity") to
	// unconditionally include in API requests. By default, fields with
	// empty values are omitted from API requests. However, any non-pointer,
	// non-interface field appearing in ForceSendFields will be sent to the
	// server regardless of whether the field is empty or not. This may be
	// used to include empty fields in Patch requests.
	ForceSendFields []string `json:"-"`

	// NullFields is a list of field names (e.g. "LauncherActivity") to
	// include in API requests with the JSON null value. By default, fields
	// with empty values are omitted from API requests. However, any field
	// with an empty value appearing in NullFields will be sent to the
	// server as null. It is an error if a field in this list has a
	// non-empty value. This may be used to include null fields in Patch
	// requests.
	NullFields []string `json:"-"`
}

func (s *RoboStartingIntent) MarshalJSON() ([]byte, error) {
	type NoMethod RoboStartingIntent
	raw := NoMethod(*s)
	return gensupport.MarshalJSON(raw, s.ForceSendFields, s.NullFields)
}

// StartActivityIntent: A starting intent specified by an action, uri,
// and categories.
type StartActivityIntent struct {
	// Action: Action name.
	// Required for START_ACTIVITY.
	Action string `json:"action,omitempty"`

	// Categories: Intent categories to set on the intent.
	// Optional.
	Categories []string `json:"categories,omitempty"`

	// Uri: URI for the action.
	// Optional.
	Uri string `json:"uri,omitempty"`

	// ForceSendFields is a list of field names (e.g. "Action") to
	// unconditionally include in API requests. By default, fields with
	// empty values are omitted from API requests. However, any non-pointer,
	// non-interface field appearing in ForceSendFields will be sent to the
	// server regardless of whether the field is empty or not. This may be
	// used to include empty fields in Patch requests.
	ForceSendFields []string `json:"-"`

	// NullFields is a list of field names (e.g. "Action") to include in API
	// requests with the JSON null value. By default, fields with empty
	// values are omitted from API requests. However, any field with an
	// empty value appearing in NullFields will be sent to the server as
	// null. It is an error if a field in this list has a non-empty value.
	// This may be used to include null fields in Patch requests.
	NullFields []string `json:"-"`
}

func (s *StartActivityIntent) MarshalJSON() ([]byte, error) {
	type NoMethod StartActivityIntent
	raw := NoMethod(*s)
	return gensupport.MarshalJSON(raw, s.ForceSendFields, s.NullFields)
}

// TestDetails: Additional details about the progress of the running
// test.
type TestDetails struct {
	// ErrorMessage: If the TestState is ERROR, then this string will
	// contain human-readable
	// details about the error.
	// @OutputOnly
	ErrorMessage string `json:"errorMessage,omitempty"`

	// ProgressMessages: Human-readable, detailed descriptions of the test's
	// progress.
	// For example: "Provisioning a device", "Starting Test".
	//
	// During the course of execution new data may be appended
	// to the end of progress_messages.
	// @OutputOnly
	ProgressMessages []string `json:"progressMessages,omitempty"`

	// VideoRecordingDisabled: Indicates that video will not be recorded for
	// this execution either because
	// the user chose to disable it or the device does not support it.
	// See AndroidModel.video_recording_not_supported
	// @OutputOnly
	VideoRecordingDisabled bool `json:"videoRecordingDisabled,omitempty"`

	// ForceSendFields is a list of field names (e.g. "ErrorMessage") to
	// unconditionally include in API requests. By default, fields with
	// empty values are omitted from API requests. However, any non-pointer,
	// non-interface field appearing in ForceSendFields will be sent to the
	// server regardless of whether the field is empty or not. This may be
	// used to include empty fields in Patch requests.
	ForceSendFields []string `json:"-"`

	// NullFields is a list of field names (e.g. "ErrorMessage") to include
	// in API requests with the JSON null value. By default, fields with
	// empty values are omitted from API requests. However, any field with
	// an empty value appearing in NullFields will be sent to the server as
	// null. It is an error if a field in this list has a non-empty value.
	// This may be used to include null fields in Patch requests.
	NullFields []string `json:"-"`
}

func (s *TestDetails) MarshalJSON() ([]byte, error) {
	type NoMethod TestDetails
	raw := NoMethod(*s)
	return gensupport.MarshalJSON(raw, s.ForceSendFields, s.NullFields)
}

// TestEnvironmentCatalog: A description of a test environment.
type TestEnvironmentCatalog struct {
	// AndroidDeviceCatalog: Android devices suitable for running Android
	// Instrumentation Tests.
	AndroidDeviceCatalog *AndroidDeviceCatalog `json:"androidDeviceCatalog,omitempty"`

	// IosDeviceCatalog: Supported iOS devices
	IosDeviceCatalog *IosDeviceCatalog `json:"iosDeviceCatalog,omitempty"`

	// NetworkConfigurationCatalog: Supported network configurations
	NetworkConfigurationCatalog *NetworkConfigurationCatalog `json:"networkConfigurationCatalog,omitempty"`

	// SoftwareCatalog: The software test environment provided by
	// TestExecutionService.
	SoftwareCatalog *ProvidedSoftwareCatalog `json:"softwareCatalog,omitempty"`

	// ServerResponse contains the HTTP response code and headers from the
	// server.
	googleapi.ServerResponse `json:"-"`

	// ForceSendFields is a list of field names (e.g.
	// "AndroidDeviceCatalog") to unconditionally include in API requests.
	// By default, fields with empty values are omitted from API requests.
	// However, any non-pointer, non-interface field appearing in
	// ForceSendFields will be sent to the server regardless of whether the
	// field is empty or not. This may be used to include empty fields in
	// Patch requests.
	ForceSendFields []string `json:"-"`

	// NullFields is a list of field names (e.g. "AndroidDeviceCatalog") to
	// include in API requests with the JSON null value. By default, fields
	// with empty values are omitted from API requests. However, any field
	// with an empty value appearing in NullFields will be sent to the
	// server as null. It is an error if a field in this list has a
	// non-empty value. This may be used to include null fields in Patch
	// requests.
	NullFields []string `json:"-"`
}

func (s *TestEnvironmentCatalog) MarshalJSON() ([]byte, error) {
	type NoMethod TestEnvironmentCatalog
	raw := NoMethod(*s)
	return gensupport.MarshalJSON(raw, s.ForceSendFields, s.NullFields)
}

// TestExecution: Specifies a single test to be executed in a single
// environment.
type TestExecution struct {
	// Environment: How the host machine(s) are configured.
	// @OutputOnly
	Environment *Environment `json:"environment,omitempty"`

	// Id: Unique id set by the backend.
	// @OutputOnly
	Id string `json:"id,omitempty"`

	// MatrixId: Id of the containing TestMatrix.
	// @OutputOnly
	MatrixId string `json:"matrixId,omitempty"`

	// ProjectId: The cloud project that owns the test
	// execution.
	// @OutputOnly
	ProjectId string `json:"projectId,omitempty"`

	// State: Indicates the current progress of the test execution (e.g.,
	// FINISHED).
	// @OutputOnly
	//
	// Possible values:
	//   "TEST_STATE_UNSPECIFIED" - Do not use.  For proto versioning only.
	//   "VALIDATING" - The execution or matrix is being validated.
	//   "PENDING" - The execution or matrix is waiting for resources to
	// become available.
	//   "RUNNING" - The execution is currently being processed.
	//
	// Can only be set on an execution.
	//   "FINISHED" - The execution or matrix has terminated normally.
	//
	// On a matrix this means that the matrix level processing completed
	// normally,
	// but individual executions may be in an ERROR state.
	//   "ERROR" - The execution or matrix has stopped because it
	// encountered an
	// infrastructure failure.
	//   "UNSUPPORTED_ENVIRONMENT" - The execution was not run because it
	// corresponds to a unsupported
	// environment.
	//
	// Can only be set on an execution.
	//   "INCOMPATIBLE_ENVIRONMENT" - The execution was not run because the
	// provided inputs are incompatible with
	// the requested environment.
	//
	// Example: requested AndroidVersion is lower than APK's
	// minSdkVersion
	//
	// Can only be set on an execution.
	//   "INCOMPATIBLE_ARCHITECTURE" - The execution was not run because the
	// provided inputs are incompatible with
	// the requested architecture.
	//
	// Example: requested device does not support running the native code
	// in
	// the supplied APK
	//
	// Can only be set on an execution.
	//   "CANCELLED" - The user cancelled the execution.
	//
	// Can only be set on an execution.
	//   "INVALID" - The execution or matrix was not run because the
	// provided inputs are not
	// valid.
	//
	// Examples: input file is not of the expected type, is
	// malformed/corrupt, or
	// was flagged as malware
	State string `json:"state,omitempty"`

	// TestDetails: Additional details about the running test.
	// @OutputOnly
	TestDetails *TestDetails `json:"testDetails,omitempty"`

	// TestSpecification: How to run the test.
	// @OutputOnly
	TestSpecification *TestSpecification `json:"testSpecification,omitempty"`

	// Timestamp: The time this test execution was initially
	// created.
	// @OutputOnly
	Timestamp string `json:"timestamp,omitempty"`

	// ToolResultsStep: Where the results for this execution are
	// written.
	// @OutputOnly
	ToolResultsStep *ToolResultsStep `json:"toolResultsStep,omitempty"`

	// ForceSendFields is a list of field names (e.g. "Environment") to
	// unconditionally include in API requests. By default, fields with
	// empty values are omitted from API requests. However, any non-pointer,
	// non-interface field appearing in ForceSendFields will be sent to the
	// server regardless of whether the field is empty or not. This may be
	// used to include empty fields in Patch requests.
	ForceSendFields []string `json:"-"`

	// NullFields is a list of field names (e.g. "Environment") to include
	// in API requests with the JSON null value. By default, fields with
	// empty values are omitted from API requests. However, any field with
	// an empty value appearing in NullFields will be sent to the server as
	// null. It is an error if a field in this list has a non-empty value.
	// This may be used to include null fields in Patch requests.
	NullFields []string `json:"-"`
}

func (s *TestExecution) MarshalJSON() ([]byte, error) {
	type NoMethod TestExecution
	raw := NoMethod(*s)
	return gensupport.MarshalJSON(raw, s.ForceSendFields, s.NullFields)
}

// TestMatrix: A group of one or more TestExecutions, built by taking
// a
// product of values over a pre-defined set of axes.
type TestMatrix struct {
	// ClientInfo: Information about the client which invoked the
	// test.
	// Optional
	ClientInfo *ClientInfo `json:"clientInfo,omitempty"`

	// EnvironmentMatrix: How the host machine(s) are configured.
	// Required
	EnvironmentMatrix *EnvironmentMatrix `json:"environmentMatrix,omitempty"`

	// InvalidMatrixDetails: Describes why the matrix is considered
	// invalid.
	// Only useful for matrices in the INVALID state.
	// @OutputOnly
	//
	// Possible values:
	//   "INVALID_MATRIX_DETAILS_UNSPECIFIED" - Do not use. For proto
	// versioning only.
	//   "DETAILS_UNAVAILABLE" - The matrix is INVALID, but there are no
	// further details available.
	//   "MALFORMED_APK" - The input app APK could not be parsed.
	//   "MALFORMED_TEST_APK" - The input test APK could not be parsed.
	//   "NO_MANIFEST" - The AndroidManifest.xml could not be found.
	//   "NO_PACKAGE_NAME" - The APK manifest does not declare a package
	// name.
	//   "TEST_SAME_AS_APP" - The test package and app package are the same.
	//   "NO_INSTRUMENTATION" - The test apk does not declare an
	// instrumentation.
	//   "NO_SIGNATURE" - The input app apk does not have a signature.
	//   "INSTRUMENTATION_ORCHESTRATOR_INCOMPATIBLE" - The test runner class
	// specified by user or in the test APK's manifest file
	// is not compatible with Android Test Orchestrator.
	// Orchestrator is only compatible with AndroidJUnitRunner version 1.0
	// or
	// higher.
	// Orchestrator can be disabled by using
	// DO_NOT_USE_ORCHESTRATOR
	// OrchestratorOption.
	//   "NO_TEST_RUNNER_CLASS" - The test APK does not contain the test
	// runner class specified by user or in
	// the manifest file.
	// This can be caused by either of the following reasons:
	// - the user provided a runner class name that's incorrect, or
	// - the test runner isn't built into the test APK (might be in the app
	// APK
	// instead).
	//   "NO_LAUNCHER_ACTIVITY" - A main launcher activity could not be
	// found.
	//   "FORBIDDEN_PERMISSIONS" - The app declares one or more permissions
	// that are not allowed.
	//   "INVALID_ROBO_DIRECTIVES" - There is a conflict in the provided
	// robo_directives.
	//   "TEST_LOOP_INTENT_FILTER_NOT_FOUND" - There there is no test loop
	// intent filter, or the one that is given is
	// not formatted correctly.
	//   "SCENARIO_LABEL_NOT_DECLARED" - The request contains a scenario
	// label that was not declared in the
	// manifest.
	//   "SCENARIO_LABEL_MALFORMED" - There was an error when parsing a
	// label's value.
	//   "SCENARIO_NOT_DECLARED" - The request contains a scenario number
	// that was not declared in the
	// manifest.
	//   "DEVICE_ADMIN_RECEIVER" - Device administrator applications are not
	// allowed.
	//   "MALFORMED_XC_TEST_ZIP" - The zipped XCTest was malformed. The zip
	// did not contain a single
	// .xctestrun file and the contents of the
	// DerivedData/Build/Products
	// directory.
	//   "BUILT_FOR_IOS_SIMULATOR" - The zipped XCTest was built for the iOS
	// simulator rather than for a
	// physical device.
	//   "NO_TESTS_IN_XC_TEST_ZIP" - The .xctestrun file did not specify any
	// test targets.
	//   "USE_DESTINATION_ARTIFACTS" - One or more of the test targets
	// defined in the .xctestrun file specifies
	// "UseDestinationArtifacts", which is disallowed.
	//   "TEST_NOT_APP_HOSTED" - XC tests which run on physical devices must
	// have
	// "IsAppHostedTestBundle" == "true" in the xctestrun file.
	//   "TEST_ONLY_APK" - The APK is marked as "testOnly".
	// NOT USED
	//   "MALFORMED_IPA" - The input IPA could not be parsed.
	// NOT USED
	//   "NO_CODE_APK" - APK contains no code.
	// See
	// also
	// https://developer.android.com/guide/topics/manifest/application-e
	// lement.html#code
	//   "INVALID_INPUT_APK" - Either the provided input APK path was
	// malformed,
	// the APK file does not exist, or the user does not have permission
	// to
	// access the APK file.
	//   "INVALID_APK_PREVIEW_SDK" - APK is built for a preview SDK which is
	// unsupported
	InvalidMatrixDetails string `json:"invalidMatrixDetails,omitempty"`

	// ProjectId: The cloud project that owns the test matrix.
	// @OutputOnly
	ProjectId string `json:"projectId,omitempty"`

	// ResultStorage: Where the results for the matrix are written.
	// Required
	ResultStorage *ResultStorage `json:"resultStorage,omitempty"`

	// State: Indicates the current progress of the test matrix (e.g.,
	// FINISHED)
	// @OutputOnly
	//
	// Possible values:
	//   "TEST_STATE_UNSPECIFIED" - Do not use.  For proto versioning only.
	//   "VALIDATING" - The execution or matrix is being validated.
	//   "PENDING" - The execution or matrix is waiting for resources to
	// become available.
	//   "RUNNING" - The execution is currently being processed.
	//
	// Can only be set on an execution.
	//   "FINISHED" - The execution or matrix has terminated normally.
	//
	// On a matrix this means that the matrix level processing completed
	// normally,
	// but individual executions may be in an ERROR state.
	//   "ERROR" - The execution or matrix has stopped because it
	// encountered an
	// infrastructure failure.
	//   "UNSUPPORTED_ENVIRONMENT" - The execution was not run because it
	// corresponds to a unsupported
	// environment.
	//
	// Can only be set on an execution.
	//   "INCOMPATIBLE_ENVIRONMENT" - The execution was not run because the
	// provided inputs are incompatible with
	// the requested environment.
	//
	// Example: requested AndroidVersion is lower than APK's
	// minSdkVersion
	//
	// Can only be set on an execution.
	//   "INCOMPATIBLE_ARCHITECTURE" - The execution was not run because the
	// provided inputs are incompatible with
	// the requested architecture.
	//
	// Example: requested device does not support running the native code
	// in
	// the supplied APK
	//
	// Can only be set on an execution.
	//   "CANCELLED" - The user cancelled the execution.
	//
	// Can only be set on an execution.
	//   "INVALID" - The execution or matrix was not run because the
	// provided inputs are not
	// valid.
	//
	// Examples: input file is not of the expected type, is
	// malformed/corrupt, or
	// was flagged as malware
	State string `json:"state,omitempty"`

	// TestExecutions: The list of test executions that the service creates
	// for this matrix.
	// @OutputOnly
	TestExecutions []*TestExecution `json:"testExecutions,omitempty"`

	// TestMatrixId: Unique id set by the service.
	// @OutputOnly
	TestMatrixId string `json:"testMatrixId,omitempty"`

	// TestSpecification: How to run the test.
	// Required
	TestSpecification *TestSpecification `json:"testSpecification,omitempty"`

	// Timestamp: The time this test matrix was initially
	// created.
	// @OutputOnly
	Timestamp string `json:"timestamp,omitempty"`

	// ServerResponse contains the HTTP response code and headers from the
	// server.
	googleapi.ServerResponse `json:"-"`

	// ForceSendFields is a list of field names (e.g. "ClientInfo") to
	// unconditionally include in API requests. By default, fields with
	// empty values are omitted from API requests. However, any non-pointer,
	// non-interface field appearing in ForceSendFields will be sent to the
	// server regardless of whether the field is empty or not. This may be
	// used to include empty fields in Patch requests.
	ForceSendFields []string `json:"-"`

	// NullFields is a list of field names (e.g. "ClientInfo") to include in
	// API requests with the JSON null value. By default, fields with empty
	// values are omitted from API requests. However, any field with an
	// empty value appearing in NullFields will be sent to the server as
	// null. It is an error if a field in this list has a non-empty value.
	// This may be used to include null fields in Patch requests.
	NullFields []string `json:"-"`
}

func (s *TestMatrix) MarshalJSON() ([]byte, error) {
	type NoMethod TestMatrix
	raw := NoMethod(*s)
	return gensupport.MarshalJSON(raw, s.ForceSendFields, s.NullFields)
}

// TestSetup: A description of how to set up the Android device prior to
// running the test.
type TestSetup struct {
	// Account: The device will be logged in on this account for the
	// duration of the test.
	// Optional
	Account *Account `json:"account,omitempty"`

	// AdditionalApks: APKs to install in addition to those being directly
	// tested.
	// Currently capped at 100.
	// Optional
	AdditionalApks []*Apk `json:"additionalApks,omitempty"`

	// DirectoriesToPull: List of directories on the device to upload to GCS
	// at the end of the test;
	// they must be absolute paths under /sdcard or /data/local/tmp.
	// Path names are restricted to characters a-z A-Z 0-9 _ - . + and
	// /
	//
	// Note: The paths /sdcard and /data will be made available and treated
	// as
	// implicit path substitutions. E.g. if /sdcard on a particular device
	// does
	// not map to external storage, the system will replace it with the
	// external
	// storage path prefix for that device.
	//
	// Optional
	DirectoriesToPull []string `json:"directoriesToPull,omitempty"`

	// EnvironmentVariables: Environment variables to set for the test (only
	// applicable for
	// instrumentation tests).
	EnvironmentVariables []*EnvironmentVariable `json:"environmentVariables,omitempty"`

	// FilesToPush: List of files to push to the device before starting the
	// test.
	//
	// Optional
	FilesToPush []*DeviceFile `json:"filesToPush,omitempty"`

	// NetworkProfile: Optional. The network traffic profile used for
	// running the test.
	// Available network profiles can be queried by using
	// the
	// NETWORK_CONFIGURATION environment type when
	// calling
	// TestEnvironmentDiscoveryService.GetTestEnvironmentCatalog.
	NetworkProfile string `json:"networkProfile,omitempty"`

	// ForceSendFields is a list of field names (e.g. "Account") to
	// unconditionally include in API requests. By default, fields with
	// empty values are omitted from API requests. However, any non-pointer,
	// non-interface field appearing in ForceSendFields will be sent to the
	// server regardless of whether the field is empty or not. This may be
	// used to include empty fields in Patch requests.
	ForceSendFields []string `json:"-"`

	// NullFields is a list of field names (e.g. "Account") to include in
	// API requests with the JSON null value. By default, fields with empty
	// values are omitted from API requests. However, any field with an
	// empty value appearing in NullFields will be sent to the server as
	// null. It is an error if a field in this list has a non-empty value.
	// This may be used to include null fields in Patch requests.
	NullFields []string `json:"-"`
}

func (s *TestSetup) MarshalJSON() ([]byte, error) {
	type NoMethod TestSetup
	raw := NoMethod(*s)
	return gensupport.MarshalJSON(raw, s.ForceSendFields, s.NullFields)
}

// TestSpecification: A description of how to run the test.
type TestSpecification struct {
	// AndroidInstrumentationTest: An Android instrumentation test.
	AndroidInstrumentationTest *AndroidInstrumentationTest `json:"androidInstrumentationTest,omitempty"`

	// AndroidRoboTest: An Android robo test.
	AndroidRoboTest *AndroidRoboTest `json:"androidRoboTest,omitempty"`

	// AndroidTestLoop: An Android Application with a Test Loop
	AndroidTestLoop *AndroidTestLoop `json:"androidTestLoop,omitempty"`

	// AutoGoogleLogin: Enables automatic Google account login.
	// If set, the service will automatically generate a Google test account
	// and
	// add it to the device, before executing the test. Note that test
	// accounts
	// might be reused.
	// Many applications show their full set of functionalities when an
	// account is
	// present on the device. Logging into the device with these
	// generated
	// accounts allows testing more functionalities.
	// Default is false.
	// Optional
	AutoGoogleLogin bool `json:"autoGoogleLogin,omitempty"`

	// DisablePerformanceMetrics: Disables performance metrics recording;
	// may reduce test latency.
	DisablePerformanceMetrics bool `json:"disablePerformanceMetrics,omitempty"`

	// DisableVideoRecording: Disables video recording; may reduce test
	// latency.
	DisableVideoRecording bool `json:"disableVideoRecording,omitempty"`

	// IosTestSetup: Optional. Test setup requirements for iOS.
	IosTestSetup *IosTestSetup `json:"iosTestSetup,omitempty"`

	// IosXcTest: An iOS XCTest, via an .xctestrun file
	IosXcTest *IosXcTest `json:"iosXcTest,omitempty"`

	// TestSetup: Test setup requirements for Android e.g. files to install,
	// bootstrap
	// scripts.
	// Optional
	TestSetup *TestSetup `json:"testSetup,omitempty"`

	// TestTimeout: Max time a test execution is allowed to run before it
	// is
	// automatically cancelled.
	// Optional, default is 5 min.
	TestTimeout string `json:"testTimeout,omitempty"`

	// ForceSendFields is a list of field names (e.g.
	// "AndroidInstrumentationTest") to unconditionally include in API
	// requests. By default, fields with empty values are omitted from API
	// requests. However, any non-pointer, non-interface field appearing in
	// ForceSendFields will be sent to the server regardless of whether the
	// field is empty or not. This may be used to include empty fields in
	// Patch requests.
	ForceSendFields []string `json:"-"`

	// NullFields is a list of field names (e.g.
	// "AndroidInstrumentationTest") to include in API requests with the
	// JSON null value. By default, fields with empty values are omitted
	// from API requests. However, any field with an empty value appearing
	// in NullFields will be sent to the server as null. It is an error if a
	// field in this list has a non-empty value. This may be used to include
	// null fields in Patch requests.
	NullFields []string `json:"-"`
}

func (s *TestSpecification) MarshalJSON() ([]byte, error) {
	type NoMethod TestSpecification
	raw := NoMethod(*s)
	return gensupport.MarshalJSON(raw, s.ForceSendFields, s.NullFields)
}

// ToolResultsExecution: Represents a tool results execution
// resource.
//
// This has the results of a TestMatrix.
type ToolResultsExecution struct {
	// ExecutionId: A tool results execution ID.
	// @OutputOnly
	ExecutionId string `json:"executionId,omitempty"`

	// HistoryId: A tool results history ID.
	// @OutputOnly
	HistoryId string `json:"historyId,omitempty"`

	// ProjectId: The cloud project that owns the tool results
	// execution.
	// @OutputOnly
	ProjectId string `json:"projectId,omitempty"`

	// ForceSendFields is a list of field names (e.g. "ExecutionId") to
	// unconditionally include in API requests. By default, fields with
	// empty values are omitted from API requests. However, any non-pointer,
	// non-interface field appearing in ForceSendFields will be sent to the
	// server regardless of whether the field is empty or not. This may be
	// used to include empty fields in Patch requests.
	ForceSendFields []string `json:"-"`

	// NullFields is a list of field names (e.g. "ExecutionId") to include
	// in API requests with the JSON null value. By default, fields with
	// empty values are omitted from API requests. However, any field with
	// an empty value appearing in NullFields will be sent to the server as
	// null. It is an error if a field in this list has a non-empty value.
	// This may be used to include null fields in Patch requests.
	NullFields []string `json:"-"`
}

func (s *ToolResultsExecution) MarshalJSON() ([]byte, error) {
	type NoMethod ToolResultsExecution
	raw := NoMethod(*s)
	return gensupport.MarshalJSON(raw, s.ForceSendFields, s.NullFields)
}

// ToolResultsHistory: Represents a tool results history resource.
type ToolResultsHistory struct {
	// HistoryId: A tool results history ID.
	// Required
	HistoryId string `json:"historyId,omitempty"`

	// ProjectId: The cloud project that owns the tool results
	// history.
	// Required
	ProjectId string `json:"projectId,omitempty"`

	// ForceSendFields is a list of field names (e.g. "HistoryId") to
	// unconditionally include in API requests. By default, fields with
	// empty values are omitted from API requests. However, any non-pointer,
	// non-interface field appearing in ForceSendFields will be sent to the
	// server regardless of whether the field is empty or not. This may be
	// used to include empty fields in Patch requests.
	ForceSendFields []string `json:"-"`

	// NullFields is a list of field names (e.g. "HistoryId") to include in
	// API requests with the JSON null value. By default, fields with empty
	// values are omitted from API requests. However, any field with an
	// empty value appearing in NullFields will be sent to the server as
	// null. It is an error if a field in this list has a non-empty value.
	// This may be used to include null fields in Patch requests.
	NullFields []string `json:"-"`
}

func (s *ToolResultsHistory) MarshalJSON() ([]byte, error) {
	type NoMethod ToolResultsHistory
	raw := NoMethod(*s)
	return gensupport.MarshalJSON(raw, s.ForceSendFields, s.NullFields)
}

// ToolResultsStep: Represents a tool results step resource.
//
// This has the results of a TestExecution.
type ToolResultsStep struct {
	// ExecutionId: A tool results execution ID.
	// @OutputOnly
	ExecutionId string `json:"executionId,omitempty"`

	// HistoryId: A tool results history ID.
	// @OutputOnly
	HistoryId string `json:"historyId,omitempty"`

	// ProjectId: The cloud project that owns the tool results
	// step.
	// @OutputOnly
	ProjectId string `json:"projectId,omitempty"`

	// StepId: A tool results step ID.
	// @OutputOnly
	StepId string `json:"stepId,omitempty"`

	// ForceSendFields is a list of field names (e.g. "ExecutionId") to
	// unconditionally include in API requests. By default, fields with
	// empty values are omitted from API requests. However, any non-pointer,
	// non-interface field appearing in ForceSendFields will be sent to the
	// server regardless of whether the field is empty or not. This may be
	// used to include empty fields in Patch requests.
	ForceSendFields []string `json:"-"`

	// NullFields is a list of field names (e.g. "ExecutionId") to include
	// in API requests with the JSON null value. By default, fields with
	// empty values are omitted from API requests. However, any field with
	// an empty value appearing in NullFields will be sent to the server as
	// null. It is an error if a field in this list has a non-empty value.
	// This may be used to include null fields in Patch requests.
	NullFields []string `json:"-"`
}

func (s *ToolResultsStep) MarshalJSON() ([]byte, error) {
	type NoMethod ToolResultsStep
	raw := NoMethod(*s)
	return gensupport.MarshalJSON(raw, s.ForceSendFields, s.NullFields)
}

// TrafficRule: Network emulation parameters
type TrafficRule struct {
	// Bandwidth: Bandwidth in kbits/second
	Bandwidth float64 `json:"bandwidth,omitempty"`

	// Burst: Burst size in kbits
	Burst float64 `json:"burst,omitempty"`

	// Delay: Packet delay, must be >= 0
	Delay string `json:"delay,omitempty"`

	// PacketDuplicationRatio: Packet duplication ratio (0.0 - 1.0)
	PacketDuplicationRatio float64 `json:"packetDuplicationRatio,omitempty"`

	// PacketLossRatio: Packet loss ratio (0.0 - 1.0)
	PacketLossRatio float64 `json:"packetLossRatio,omitempty"`

	// ForceSendFields is a list of field names (e.g. "Bandwidth") to
	// unconditionally include in API requests. By default, fields with
	// empty values are omitted from API requests. However, any non-pointer,
	// non-interface field appearing in ForceSendFields will be sent to the
	// server regardless of whether the field is empty or not. This may be
	// used to include empty fields in Patch requests.
	ForceSendFields []string `json:"-"`

	// NullFields is a list of field names (e.g. "Bandwidth") to include in
	// API requests with the JSON null value. By default, fields with empty
	// values are omitted from API requests. However, any field with an
	// empty value appearing in NullFields will be sent to the server as
	// null. It is an error if a field in this list has a non-empty value.
	// This may be used to include null fields in Patch requests.
	NullFields []string `json:"-"`
}

func (s *TrafficRule) MarshalJSON() ([]byte, error) {
	type NoMethod TrafficRule
	raw := NoMethod(*s)
	return gensupport.MarshalJSON(raw, s.ForceSendFields, s.NullFields)
}

func (s *TrafficRule) UnmarshalJSON(data []byte) error {
	type NoMethod TrafficRule
	var s1 struct {
		Bandwidth              gensupport.JSONFloat64 `json:"bandwidth"`
		Burst                  gensupport.JSONFloat64 `json:"burst"`
		PacketDuplicationRatio gensupport.JSONFloat64 `json:"packetDuplicationRatio"`
		PacketLossRatio        gensupport.JSONFloat64 `json:"packetLossRatio"`
		*NoMethod
	}
	s1.NoMethod = (*NoMethod)(s)
	if err := json.Unmarshal(data, &s1); err != nil {
		return err
	}
	s.Bandwidth = float64(s1.Bandwidth)
	s.Burst = float64(s1.Burst)
	s.PacketDuplicationRatio = float64(s1.PacketDuplicationRatio)
	s.PacketLossRatio = float64(s1.PacketLossRatio)
	return nil
}

// method id "testing.applicationDetailService.getApkDetails":

type ApplicationDetailServiceGetApkDetailsCall struct {
	s             *Service
	filereference *FileReference
	urlParams_    gensupport.URLParams
	ctx_          context.Context
	header_       http.Header
}

// GetApkDetails: Request the details of an Android application APK.
func (r *ApplicationDetailServiceService) GetApkDetails(filereference *FileReference) *ApplicationDetailServiceGetApkDetailsCall {
	c := &ApplicationDetailServiceGetApkDetailsCall{s: r.s, urlParams_: make(gensupport.URLParams)}
	c.filereference = filereference
	return c
}

// Fields allows partial responses to be retrieved. See
// https://developers.google.com/gdata/docs/2.0/basics#PartialResponse
// for more information.
func (c *ApplicationDetailServiceGetApkDetailsCall) Fields(s ...googleapi.Field) *ApplicationDetailServiceGetApkDetailsCall {
	c.urlParams_.Set("fields", googleapi.CombineFields(s))
	return c
}

// Context sets the context to be used in this call's Do method. Any
// pending HTTP request will be aborted if the provided context is
// canceled.
func (c *ApplicationDetailServiceGetApkDetailsCall) Context(ctx context.Context) *ApplicationDetailServiceGetApkDetailsCall {
	c.ctx_ = ctx
	return c
}

// Header returns an http.Header that can be modified by the caller to
// add HTTP headers to the request.
func (c *ApplicationDetailServiceGetApkDetailsCall) Header() http.Header {
	if c.header_ == nil {
		c.header_ = make(http.Header)
	}
	return c.header_
}

func (c *ApplicationDetailServiceGetApkDetailsCall) doRequest(alt string) (*http.Response, error) {
	reqHeaders := make(http.Header)
	for k, v := range c.header_ {
		reqHeaders[k] = v
	}
	reqHeaders.Set("User-Agent", c.s.userAgent())
	var body io.Reader = nil
	body, err := googleapi.WithoutDataWrapper.JSONReader(c.filereference)
	if err != nil {
		return nil, err
	}
	reqHeaders.Set("Content-Type", "application/json")
	c.urlParams_.Set("alt", alt)
	c.urlParams_.Set("prettyPrint", "false")
	urls := googleapi.ResolveRelative(c.s.BasePath, "v1/applicationDetailService/getApkDetails")
	urls += "?" + c.urlParams_.Encode()
	req, _ := http.NewRequest("POST", urls, body)
	req.Header = reqHeaders
	return gensupport.SendRequest(c.ctx_, c.s.client, req)
}

// Do executes the "testing.applicationDetailService.getApkDetails" call.
// Exactly one of *GetApkDetailsResponse or error will be non-nil. Any
// non-2xx status code is an error. Response headers are in either
// *GetApkDetailsResponse.ServerResponse.Header or (if a response was
// returned at all) in error.(*googleapi.Error).Header. Use
// googleapi.IsNotModified to check whether the returned error was
// because http.StatusNotModified was returned.
func (c *ApplicationDetailServiceGetApkDetailsCall) Do(opts ...googleapi.CallOption) (*GetApkDetailsResponse, error) {
	gensupport.SetOptions(c.urlParams_, opts...)
	res, err := c.doRequest("json")
	if res != nil && res.StatusCode == http.StatusNotModified {
		if res.Body != nil {
			res.Body.Close()
		}
		return nil, &googleapi.Error{
			Code:   res.StatusCode,
			Header: res.Header,
		}
	}
	if err != nil {
		return nil, err
	}
	defer googleapi.CloseBody(res)
	if err := googleapi.CheckResponse(res); err != nil {
		return nil, err
	}
	ret := &GetApkDetailsResponse{
		ServerResponse: googleapi.ServerResponse{
			Header:         res.Header,
			HTTPStatusCode: res.StatusCode,
		},
	}
	target := &ret
	if err := gensupport.DecodeResponse(target, res); err != nil {
		return nil, err
	}
	return ret, nil
	// {
	//   "description": "Request the details of an Android application APK.",
	//   "flatPath": "v1/applicationDetailService/getApkDetails",
	//   "httpMethod": "POST",
	//   "id": "testing.applicationDetailService.getApkDetails",
	//   "parameterOrder": [],
	//   "parameters": {},
	//   "path": "v1/applicationDetailService/getApkDetails",
	//   "request": {
	//     "$ref": "FileReference"
	//   },
	//   "response": {
	//     "$ref": "GetApkDetailsResponse"
	//   },
	//   "scopes": [
	//     "https://www.googleapis.com/auth/cloud-platform"
	//   ]
	// }

}

// method id "testing.projects.testMatrices.cancel":

type ProjectsTestMatricesCancelCall struct {
	s            *Service
	projectId    string
	testMatrixId string
	urlParams_   gensupport.URLParams
	ctx_         context.Context
	header_      http.Header
}

// Cancel: Cancels unfinished test executions in a test matrix.
// This call returns immediately and cancellation proceeds
// asychronously.
// If the matrix is already final, this operation will have no
// effect.
//
// May return any of the following canonical error codes:
//
// - PERMISSION_DENIED - if the user is not authorized to read project
// - INVALID_ARGUMENT - if the request is malformed
// - NOT_FOUND - if the Test Matrix does not exist
func (r *ProjectsTestMatricesService) Cancel(projectId string, testMatrixId string) *ProjectsTestMatricesCancelCall {
	c := &ProjectsTestMatricesCancelCall{s: r.s, urlParams_: make(gensupport.URLParams)}
	c.projectId = projectId
	c.testMatrixId = testMatrixId
	return c
}

// Fields allows partial responses to be retrieved. See
// https://developers.google.com/gdata/docs/2.0/basics#PartialResponse
// for more information.
func (c *ProjectsTestMatricesCancelCall) Fields(s ...googleapi.Field) *ProjectsTestMatricesCancelCall {
	c.urlParams_.Set("fields", googleapi.CombineFields(s))
	return c
}

// Context sets the context to be used in this call's Do method. Any
// pending HTTP request will be aborted if the provided context is
// canceled.
func (c *ProjectsTestMatricesCancelCall) Context(ctx context.Context) *ProjectsTestMatricesCancelCall {
	c.ctx_ = ctx
	return c
}

// Header returns an http.Header that can be modified by the caller to
// add HTTP headers to the request.
func (c *ProjectsTestMatricesCancelCall) Header() http.Header {
	if c.header_ == nil {
		c.header_ = make(http.Header)
	}
	return c.header_
}

func (c *ProjectsTestMatricesCancelCall) doRequest(alt string) (*http.Response, error) {
	reqHeaders := make(http.Header)
	for k, v := range c.header_ {
		reqHeaders[k] = v
	}
	reqHeaders.Set("User-Agent", c.s.userAgent())
	var body io.Reader = nil
	c.urlParams_.Set("alt", alt)
	c.urlParams_.Set("prettyPrint", "false")
	urls := googleapi.ResolveRelative(c.s.BasePath, "v1/projects/{projectId}/testMatrices/{testMatrixId}:cancel")
	urls += "?" + c.urlParams_.Encode()
	req, _ := http.NewRequest("POST", urls, body)
	req.Header = reqHeaders
	googleapi.Expand(req.URL, map[string]string{
		"projectId":    c.projectId,
		"testMatrixId": c.testMatrixId,
	})
	return gensupport.SendRequest(c.ctx_, c.s.client, req)
}

// Do executes the "testing.projects.testMatrices.cancel" call.
// Exactly one of *CancelTestMatrixResponse or error will be non-nil.
// Any non-2xx status code is an error. Response headers are in either
// *CancelTestMatrixResponse.ServerResponse.Header or (if a response was
// returned at all) in error.(*googleapi.Error).Header. Use
// googleapi.IsNotModified to check whether the returned error was
// because http.StatusNotModified was returned.
func (c *ProjectsTestMatricesCancelCall) Do(opts ...googleapi.CallOption) (*CancelTestMatrixResponse, error) {
	gensupport.SetOptions(c.urlParams_, opts...)
	res, err := c.doRequest("json")
	if res != nil && res.StatusCode == http.StatusNotModified {
		if res.Body != nil {
			res.Body.Close()
		}
		return nil, &googleapi.Error{
			Code:   res.StatusCode,
			Header: res.Header,
		}
	}
	if err != nil {
		return nil, err
	}
	defer googleapi.CloseBody(res)
	if err := googleapi.CheckResponse(res); err != nil {
		return nil, err
	}
	ret := &CancelTestMatrixResponse{
		ServerResponse: googleapi.ServerResponse{
			Header:         res.Header,
			HTTPStatusCode: res.StatusCode,
		},
	}
	target := &ret
	if err := gensupport.DecodeResponse(target, res); err != nil {
		return nil, err
	}
	return ret, nil
	// {
	//   "description": "Cancels unfinished test executions in a test matrix.\nThis call returns immediately and cancellation proceeds asychronously.\nIf the matrix is already final, this operation will have no effect.\n\nMay return any of the following canonical error codes:\n\n- PERMISSION_DENIED - if the user is not authorized to read project\n- INVALID_ARGUMENT - if the request is malformed\n- NOT_FOUND - if the Test Matrix does not exist",
	//   "flatPath": "v1/projects/{projectId}/testMatrices/{testMatrixId}:cancel",
	//   "httpMethod": "POST",
	//   "id": "testing.projects.testMatrices.cancel",
	//   "parameterOrder": [
	//     "projectId",
	//     "testMatrixId"
	//   ],
	//   "parameters": {
	//     "projectId": {
	//       "description": "Cloud project that owns the test.",
	//       "location": "path",
	//       "required": true,
	//       "type": "string"
	//     },
	//     "testMatrixId": {
	//       "description": "Test matrix that will be canceled.",
	//       "location": "path",
	//       "required": true,
	//       "type": "string"
	//     }
	//   },
	//   "path": "v1/projects/{projectId}/testMatrices/{testMatrixId}:cancel",
	//   "response": {
	//     "$ref": "CancelTestMatrixResponse"
	//   },
	//   "scopes": [
	//     "https://www.googleapis.com/auth/cloud-platform"
	//   ]
	// }

}

// method id "testing.projects.testMatrices.create":

type ProjectsTestMatricesCreateCall struct {
	s          *Service
	projectId  string
	testmatrix *TestMatrix
	urlParams_ gensupport.URLParams
	ctx_       context.Context
	header_    http.Header
}

// Create: Request to run a matrix of tests according to the given
// specifications.
// Unsupported environments will be returned in the state
// UNSUPPORTED.
// Matrices are limited to at most 200 supported executions.
//
// May return any of the following canonical error codes:
//
// - PERMISSION_DENIED - if the user is not authorized to write to
// project
// - INVALID_ARGUMENT - if the request is malformed or if the matrix
// expands
//                      to more than 200 supported executions
func (r *ProjectsTestMatricesService) Create(projectId string, testmatrix *TestMatrix) *ProjectsTestMatricesCreateCall {
	c := &ProjectsTestMatricesCreateCall{s: r.s, urlParams_: make(gensupport.URLParams)}
	c.projectId = projectId
	c.testmatrix = testmatrix
	return c
}

// RequestId sets the optional parameter "requestId": A string id used
// to detect duplicated requests.
// Ids are automatically scoped to a project, so
// users should ensure the ID is unique per-project.
// A UUID is recommended.
//
// Optional, but strongly recommended.
func (c *ProjectsTestMatricesCreateCall) RequestId(requestId string) *ProjectsTestMatricesCreateCall {
	c.urlParams_.Set("requestId", requestId)
	return c
}

// Fields allows partial responses to be retrieved. See
// https://developers.google.com/gdata/docs/2.0/basics#PartialResponse
// for more information.
func (c *ProjectsTestMatricesCreateCall) Fields(s ...googleapi.Field) *ProjectsTestMatricesCreateCall {
	c.urlParams_.Set("fields", googleapi.CombineFields(s))
	return c
}

// Context sets the context to be used in this call's Do method. Any
// pending HTTP request will be aborted if the provided context is
// canceled.
func (c *ProjectsTestMatricesCreateCall) Context(ctx context.Context) *ProjectsTestMatricesCreateCall {
	c.ctx_ = ctx
	return c
}

// Header returns an http.Header that can be modified by the caller to
// add HTTP headers to the request.
func (c *ProjectsTestMatricesCreateCall) Header() http.Header {
	if c.header_ == nil {
		c.header_ = make(http.Header)
	}
	return c.header_
}

func (c *ProjectsTestMatricesCreateCall) doRequest(alt string) (*http.Response, error) {
	reqHeaders := make(http.Header)
	for k, v := range c.header_ {
		reqHeaders[k] = v
	}
	reqHeaders.Set("User-Agent", c.s.userAgent())
	var body io.Reader = nil
	body, err := googleapi.WithoutDataWrapper.JSONReader(c.testmatrix)
	if err != nil {
		return nil, err
	}
	reqHeaders.Set("Content-Type", "application/json")
	c.urlParams_.Set("alt", alt)
	c.urlParams_.Set("prettyPrint", "false")
	urls := googleapi.ResolveRelative(c.s.BasePath, "v1/projects/{projectId}/testMatrices")
	urls += "?" + c.urlParams_.Encode()
	req, _ := http.NewRequest("POST", urls, body)
	req.Header = reqHeaders
	googleapi.Expand(req.URL, map[string]string{
		"projectId": c.projectId,
	})
	return gensupport.SendRequest(c.ctx_, c.s.client, req)
}

// Do executes the "testing.projects.testMatrices.create" call.
// Exactly one of *TestMatrix or error will be non-nil. Any non-2xx
// status code is an error. Response headers are in either
// *TestMatrix.ServerResponse.Header or (if a response was returned at
// all) in error.(*googleapi.Error).Header. Use googleapi.IsNotModified
// to check whether the returned error was because
// http.StatusNotModified was returned.
func (c *ProjectsTestMatricesCreateCall) Do(opts ...googleapi.CallOption) (*TestMatrix, error) {
	gensupport.SetOptions(c.urlParams_, opts...)
	res, err := c.doRequest("json")
	if res != nil && res.StatusCode == http.StatusNotModified {
		if res.Body != nil {
			res.Body.Close()
		}
		return nil, &googleapi.Error{
			Code:   res.StatusCode,
			Header: res.Header,
		}
	}
	if err != nil {
		return nil, err
	}
	defer googleapi.CloseBody(res)
	if err := googleapi.CheckResponse(res); err != nil {
		return nil, err
	}
	ret := &TestMatrix{
		ServerResponse: googleapi.ServerResponse{
			Header:         res.Header,
			HTTPStatusCode: res.StatusCode,
		},
	}
	target := &ret
	if err := gensupport.DecodeResponse(target, res); err != nil {
		return nil, err
	}
	return ret, nil
	// {
	//   "description": "Request to run a matrix of tests according to the given specifications.\nUnsupported environments will be returned in the state UNSUPPORTED.\nMatrices are limited to at most 200 supported executions.\n\nMay return any of the following canonical error codes:\n\n- PERMISSION_DENIED - if the user is not authorized to write to project\n- INVALID_ARGUMENT - if the request is malformed or if the matrix expands\n                     to more than 200 supported executions",
	//   "flatPath": "v1/projects/{projectId}/testMatrices",
	//   "httpMethod": "POST",
	//   "id": "testing.projects.testMatrices.create",
	//   "parameterOrder": [
	//     "projectId"
	//   ],
	//   "parameters": {
	//     "projectId": {
	//       "description": "The GCE project under which this job will run.",
	//       "location": "path",
	//       "required": true,
	//       "type": "string"
	//     },
	//     "requestId": {
	//       "description": "A string id used to detect duplicated requests.\nIds are automatically scoped to a project, so\nusers should ensure the ID is unique per-project.\nA UUID is recommended.\n\nOptional, but strongly recommended.",
	//       "location": "query",
	//       "type": "string"
	//     }
	//   },
	//   "path": "v1/projects/{projectId}/testMatrices",
	//   "request": {
	//     "$ref": "TestMatrix"
	//   },
	//   "response": {
	//     "$ref": "TestMatrix"
	//   },
	//   "scopes": [
	//     "https://www.googleapis.com/auth/cloud-platform"
	//   ]
	// }

}

// method id "testing.projects.testMatrices.get":

type ProjectsTestMatricesGetCall struct {
	s            *Service
	projectId    string
	testMatrixId string
	urlParams_   gensupport.URLParams
	ifNoneMatch_ string
	ctx_         context.Context
	header_      http.Header
}

// Get: Check the status of a test matrix.
//
// May return any of the following canonical error codes:
//
// - PERMISSION_DENIED - if the user is not authorized to read project
// - INVALID_ARGUMENT - if the request is malformed
// - NOT_FOUND - if the Test Matrix does not exist
func (r *ProjectsTestMatricesService) Get(projectId string, testMatrixId string) *ProjectsTestMatricesGetCall {
	c := &ProjectsTestMatricesGetCall{s: r.s, urlParams_: make(gensupport.URLParams)}
	c.projectId = projectId
	c.testMatrixId = testMatrixId
	return c
}

// Fields allows partial responses to be retrieved. See
// https://developers.google.com/gdata/docs/2.0/basics#PartialResponse
// for more information.
func (c *ProjectsTestMatricesGetCall) Fields(s ...googleapi.Field) *ProjectsTestMatricesGetCall {
	c.urlParams_.Set("fields", googleapi.CombineFields(s))
	return c
}

// IfNoneMatch sets the optional parameter which makes the operation
// fail if the object's ETag matches the given value. This is useful for
// getting updates only after the object has changed since the last
// request. Use googleapi.IsNotModified to check whether the response
// error from Do is the result of In-None-Match.
func (c *ProjectsTestMatricesGetCall) IfNoneMatch(entityTag string) *ProjectsTestMatricesGetCall {
	c.ifNoneMatch_ = entityTag
	return c
}

// Context sets the context to be used in this call's Do method. Any
// pending HTTP request will be aborted if the provided context is
// canceled.
func (c *ProjectsTestMatricesGetCall) Context(ctx context.Context) *ProjectsTestMatricesGetCall {
	c.ctx_ = ctx
	return c
}

// Header returns an http.Header that can be modified by the caller to
// add HTTP headers to the request.
func (c *ProjectsTestMatricesGetCall) Header() http.Header {
	if c.header_ == nil {
		c.header_ = make(http.Header)
	}
	return c.header_
}

func (c *ProjectsTestMatricesGetCall) doRequest(alt string) (*http.Response, error) {
	reqHeaders := make(http.Header)
	for k, v := range c.header_ {
		reqHeaders[k] = v
	}
	reqHeaders.Set("User-Agent", c.s.userAgent())
	if c.ifNoneMatch_ != "" {
		reqHeaders.Set("If-None-Match", c.ifNoneMatch_)
	}
	var body io.Reader = nil
	c.urlParams_.Set("alt", alt)
	c.urlParams_.Set("prettyPrint", "false")
	urls := googleapi.ResolveRelative(c.s.BasePath, "v1/projects/{projectId}/testMatrices/{testMatrixId}")
	urls += "?" + c.urlParams_.Encode()
	req, _ := http.NewRequest("GET", urls, body)
	req.Header = reqHeaders
	googleapi.Expand(req.URL, map[string]string{
		"projectId":    c.projectId,
		"testMatrixId": c.testMatrixId,
	})
	return gensupport.SendRequest(c.ctx_, c.s.client, req)
}

// Do executes the "testing.projects.testMatrices.get" call.
// Exactly one of *TestMatrix or error will be non-nil. Any non-2xx
// status code is an error. Response headers are in either
// *TestMatrix.ServerResponse.Header or (if a response was returned at
// all) in error.(*googleapi.Error).Header. Use googleapi.IsNotModified
// to check whether the returned error was because
// http.StatusNotModified was returned.
func (c *ProjectsTestMatricesGetCall) Do(opts ...googleapi.CallOption) (*TestMatrix, error) {
	gensupport.SetOptions(c.urlParams_, opts...)
	res, err := c.doRequest("json")
	if res != nil && res.StatusCode == http.StatusNotModified {
		if res.Body != nil {
			res.Body.Close()
		}
		return nil, &googleapi.Error{
			Code:   res.StatusCode,
			Header: res.Header,
		}
	}
	if err != nil {
		return nil, err
	}
	defer googleapi.CloseBody(res)
	if err := googleapi.CheckResponse(res); err != nil {
		return nil, err
	}
	ret := &TestMatrix{
		ServerResponse: googleapi.ServerResponse{
			Header:         res.Header,
			HTTPStatusCode: res.StatusCode,
		},
	}
	target := &ret
	if err := gensupport.DecodeResponse(target, res); err != nil {
		return nil, err
	}
	return ret, nil
	// {
	//   "description": "Check the status of a test matrix.\n\nMay return any of the following canonical error codes:\n\n- PERMISSION_DENIED - if the user is not authorized to read project\n- INVALID_ARGUMENT - if the request is malformed\n- NOT_FOUND - if the Test Matrix does not exist",
	//   "flatPath": "v1/projects/{projectId}/testMatrices/{testMatrixId}",
	//   "httpMethod": "GET",
	//   "id": "testing.projects.testMatrices.get",
	//   "parameterOrder": [
	//     "projectId",
	//     "testMatrixId"
	//   ],
	//   "parameters": {
	//     "projectId": {
	//       "description": "Cloud project that owns the test matrix.",
	//       "location": "path",
	//       "required": true,
	//       "type": "string"
	//     },
	//     "testMatrixId": {
	//       "description": "Unique test matrix id which was assigned by the service.",
	//       "location": "path",
	//       "required": true,
	//       "type": "string"
	//     }
	//   },
	//   "path": "v1/projects/{projectId}/testMatrices/{testMatrixId}",
	//   "response": {
	//     "$ref": "TestMatrix"
	//   },
	//   "scopes": [
	//     "https://www.googleapis.com/auth/cloud-platform",
	//     "https://www.googleapis.com/auth/cloud-platform.read-only"
	//   ]
	// }

}

// method id "testing.testEnvironmentCatalog.get":

type TestEnvironmentCatalogGetCall struct {
	s               *Service
	environmentType string
	urlParams_      gensupport.URLParams
	ifNoneMatch_    string
	ctx_            context.Context
	header_         http.Header
}

// Get: Get the catalog of supported test environments.
//
// May return any of the following canonical error codes:
//
// - INVALID_ARGUMENT - if the request is malformed
// - NOT_FOUND - if the environment type does not exist
// - INTERNAL - if an internal error occurred
func (r *TestEnvironmentCatalogService) Get(environmentType string) *TestEnvironmentCatalogGetCall {
	c := &TestEnvironmentCatalogGetCall{s: r.s, urlParams_: make(gensupport.URLParams)}
	c.environmentType = environmentType
	return c
}

// ProjectId sets the optional parameter "projectId": For authorization,
// the cloud project requesting the TestEnvironmentCatalog.
// Optional
func (c *TestEnvironmentCatalogGetCall) ProjectId(projectId string) *TestEnvironmentCatalogGetCall {
	c.urlParams_.Set("projectId", projectId)
	return c
}

// Fields allows partial responses to be retrieved. See
// https://developers.google.com/gdata/docs/2.0/basics#PartialResponse
// for more information.
func (c *TestEnvironmentCatalogGetCall) Fields(s ...googleapi.Field) *TestEnvironmentCatalogGetCall {
	c.urlParams_.Set("fields", googleapi.CombineFields(s))
	return c
}

// IfNoneMatch sets the optional parameter which makes the operation
// fail if the object's ETag matches the given value. This is useful for
// getting updates only after the object has changed since the last
// request. Use googleapi.IsNotModified to check whether the response
// error from Do is the result of In-None-Match.
func (c *TestEnvironmentCatalogGetCall) IfNoneMatch(entityTag string) *TestEnvironmentCatalogGetCall {
	c.ifNoneMatch_ = entityTag
	return c
}

// Context sets the context to be used in this call's Do method. Any
// pending HTTP request will be aborted if the provided context is
// canceled.
func (c *TestEnvironmentCatalogGetCall) Context(ctx context.Context) *TestEnvironmentCatalogGetCall {
	c.ctx_ = ctx
	return c
}

// Header returns an http.Header that can be modified by the caller to
// add HTTP headers to the request.
func (c *TestEnvironmentCatalogGetCall) Header() http.Header {
	if c.header_ == nil {
		c.header_ = make(http.Header)
	}
	return c.header_
}

func (c *TestEnvironmentCatalogGetCall) doRequest(alt string) (*http.Response, error) {
	reqHeaders := make(http.Header)
	for k, v := range c.header_ {
		reqHeaders[k] = v
	}
	reqHeaders.Set("User-Agent", c.s.userAgent())
	if c.ifNoneMatch_ != "" {
		reqHeaders.Set("If-None-Match", c.ifNoneMatch_)
	}
	var body io.Reader = nil
	c.urlParams_.Set("alt", alt)
	c.urlParams_.Set("prettyPrint", "false")
	urls := googleapi.ResolveRelative(c.s.BasePath, "v1/testEnvironmentCatalog/{environmentType}")
	urls += "?" + c.urlParams_.Encode()
	req, _ := http.NewRequest("GET", urls, body)
	req.Header = reqHeaders
	googleapi.Expand(req.URL, map[string]string{
		"environmentType": c.environmentType,
	})
	return gensupport.SendRequest(c.ctx_, c.s.client, req)
}

// Do executes the "testing.testEnvironmentCatalog.get" call.
// Exactly one of *TestEnvironmentCatalog or error will be non-nil. Any
// non-2xx status code is an error. Response headers are in either
// *TestEnvironmentCatalog.ServerResponse.Header or (if a response was
// returned at all) in error.(*googleapi.Error).Header. Use
// googleapi.IsNotModified to check whether the returned error was
// because http.StatusNotModified was returned.
func (c *TestEnvironmentCatalogGetCall) Do(opts ...googleapi.CallOption) (*TestEnvironmentCatalog, error) {
	gensupport.SetOptions(c.urlParams_, opts...)
	res, err := c.doRequest("json")
	if res != nil && res.StatusCode == http.StatusNotModified {
		if res.Body != nil {
			res.Body.Close()
		}
		return nil, &googleapi.Error{
			Code:   res.StatusCode,
			Header: res.Header,
		}
	}
	if err != nil {
		return nil, err
	}
	defer googleapi.CloseBody(res)
	if err := googleapi.CheckResponse(res); err != nil {
		return nil, err
	}
	ret := &TestEnvironmentCatalog{
		ServerResponse: googleapi.ServerResponse{
			Header:         res.Header,
			HTTPStatusCode: res.StatusCode,
		},
	}
	target := &ret
	if err := gensupport.DecodeResponse(target, res); err != nil {
		return nil, err
	}
	return ret, nil
	// {
	//   "description": "Get the catalog of supported test environments.\n\nMay return any of the following canonical error codes:\n\n- INVALID_ARGUMENT - if the request is malformed\n- NOT_FOUND - if the environment type does not exist\n- INTERNAL - if an internal error occurred",
	//   "flatPath": "v1/testEnvironmentCatalog/{environmentType}",
	//   "httpMethod": "GET",
	//   "id": "testing.testEnvironmentCatalog.get",
	//   "parameterOrder": [
	//     "environmentType"
	//   ],
	//   "parameters": {
	//     "environmentType": {
	//       "description": "The type of environment that should be listed.\nRequired",
	//       "enum": [
	//         "ENVIRONMENT_TYPE_UNSPECIFIED",
	//         "ANDROID",
	//         "IOS",
	//         "NETWORK_CONFIGURATION",
	//         "PROVIDED_SOFTWARE"
	//       ],
	//       "location": "path",
	//       "required": true,
	//       "type": "string"
	//     },
	//     "projectId": {
	//       "description": "For authorization, the cloud project requesting the TestEnvironmentCatalog.\nOptional",
	//       "location": "query",
	//       "type": "string"
	//     }
	//   },
	//   "path": "v1/testEnvironmentCatalog/{environmentType}",
	//   "response": {
	//     "$ref": "TestEnvironmentCatalog"
	//   },
	//   "scopes": [
	//     "https://www.googleapis.com/auth/cloud-platform",
	//     "https://www.googleapis.com/auth/cloud-platform.read-only"
	//   ]
	// }

}
