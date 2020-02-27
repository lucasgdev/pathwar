// Code generated by protoc-gen-gogo. DO NOT EDIT.
// source: errcode.proto

package errcode

import (
	fmt "fmt"
	math "math"

	_ "github.com/gogo/protobuf/gogoproto"
	proto "github.com/gogo/protobuf/proto"
)

// Reference imports to suppress errors if they are not otherwise used.
var _ = proto.Marshal
var _ = fmt.Errorf
var _ = math.Inf

// This is a compile-time assertion to ensure that this generated file
// is compatible with the proto package it is being compiled against.
// A compilation error at this line likely means your copy of the
// proto package needs to be updated.
const _ = proto.GoGoProtoPackageIsVersion3 // please upgrade the proto package

type ErrCode int32

const (
	Undefined                               ErrCode = 0
	TODO                                    ErrCode = 666
	ErrNotImplemented                       ErrCode = 777
	ErrDeprecated                           ErrCode = 888
	ErrInternal                             ErrCode = 999
	ErrInvalidInput                         ErrCode = 101
	ErrMissingInput                         ErrCode = 102
	ErrUnauthenticated                      ErrCode = 103
	ErrSSOGetOIDC                           ErrCode = 1001
	ErrSSOInvalidPublicKey                  ErrCode = 1002
	ErrSSOFailedKeycloakRequest             ErrCode = 1003
	ErrSSOInvalidKeycloakResponse           ErrCode = 1004
	ErrSSOLogout                            ErrCode = 1005
	ErrSSOInitKeycloakClient                ErrCode = 1006
	ErrSSOInvalidBearer                     ErrCode = 1007
	ErrSSOKeycloakError                     ErrCode = 1008
	ErrDBNotFound                           ErrCode = 2001
	ErrDBInternal                           ErrCode = 2002
	ErrDBRunMigrations                      ErrCode = 2003
	ErrDBInit                               ErrCode = 2004
	ErrDBConnect                            ErrCode = 2005
	ErrDBAutoMigrate                        ErrCode = 2006
	ErrDBAddForeignKey                      ErrCode = 2007
	ErrComposeInvalidPath                   ErrCode = 3001
	ErrComposeDirectoryNotFound             ErrCode = 3002
	ErrComposeReadConfig                    ErrCode = 3003
	ErrComposeInvalidConfig                 ErrCode = 3004
	ErrComposeMarshalConfig                 ErrCode = 3005
	ErrComposeCreateTempFile                ErrCode = 3006
	ErrComposeWriteTempFile                 ErrCode = 3007
	ErrComposeCloseTempFile                 ErrCode = 3008
	ErrComposeBuild                         ErrCode = 3009
	ErrComposeBundle                        ErrCode = 30010
	ErrComposeReadDab                       ErrCode = 3011
	ErrComposeParseDab                      ErrCode = 3012
	ErrComposeParseConfig                   ErrCode = 3013
	ErrComposeCreateTempDir                 ErrCode = 3014
	ErrComposeUpdateTempFile                ErrCode = 3015
	ErrComposeForceRecreateDown             ErrCode = 3016
	ErrComposeRunCreate                     ErrCode = 3017
	ErrComposeRunUp                         ErrCode = 3018
	ErrGetPWInitBinary                      ErrCode = 3019
	ErrWritePWInitFileHeader                ErrCode = 3020
	ErrWritePWInitFile                      ErrCode = 3021
	ErrMarshalPWInitConfigFile              ErrCode = 3022
	ErrWritePWInitConfigFileHeader          ErrCode = 3023
	ErrWritePWInitConfigFile                ErrCode = 3024
	ErrWritePWInitCloseTarWriter            ErrCode = 3025
	ErrCopyPWInitToContainer                ErrCode = 3026
	ErrComposeGetContainersInfo             ErrCode = 3027
	ErrMissingPwinitConfig                  ErrCode = 3028
	ErrGetUserIDFromContext                 ErrCode = 4001
	ErrMissingChallengeValidation           ErrCode = 4002
	ErrInvalidSeason                        ErrCode = 4003
	ErrTeamNotInSeason                      ErrCode = 4004
	ErrGetChallengeSubscription             ErrCode = 4005
	ErrInitSnowflake                        ErrCode = 4006
	ErrUpdateChallengeSubscription          ErrCode = 4007
	ErrCreateChallengeValidation            ErrCode = 4008
	ErrGetChallengeValidation               ErrCode = 4009
	ErrInvalidTeam                          ErrCode = 4010
	ErrChallengeAlreadySubscribed           ErrCode = 4011
	ErrCreateChallengeSubscription          ErrCode = 4012
	ErrFindOrganizations                    ErrCode = 4013
	ErrGetSeasonFromSeasonChallenge         ErrCode = 4014
	ErrGetUserTeamFromSeason                ErrCode = 4015
	ErrInvalidSeasonID                      ErrCode = 4016
	ErrUserHasNoTeamForSeason               ErrCode = 4017
	ErrGetSeasonChallenges                  ErrCode = 4018
	ErrGetSeason                            ErrCode = 4019
	ErrSeasonDenied                         ErrCode = 4020
	ErrAlreadyHasTeamForSeason              ErrCode = 4021
	ErrReservedName                         ErrCode = 4022
	ErrCheckOrganizationUniqueName          ErrCode = 4023
	ErrCreateOrganization                   ErrCode = 4024
	ErrOrganizationAlreadyHasTeamForSeason  ErrCode = 4025
	ErrGetOrganization                      ErrCode = 4026
	ErrGetSeasonChallenge                   ErrCode = 4027
	ErrCannotCreateTeamForSoloOrganization  ErrCode = 4028
	ErrUserNotInOrganization                ErrCode = 4029
	ErrCreateTeam                           ErrCode = 4030
	ErrGetTeam                              ErrCode = 4031
	ErrGetTeams                             ErrCode = 4032
	ErrGetUser                              ErrCode = 4033
	ErrUpdateUser                           ErrCode = 4034
	ErrUpdateTeam                           ErrCode = 4035
	ErrUpdateOrganization                   ErrCode = 4036
	ErrNewUserFromClaims                    ErrCode = 4037
	ErrGetOAuthUser                         ErrCode = 4038
	ErrDifferentUserBetweenTokenAndDatabase ErrCode = 4039
	ErrLoadUserSeasons                      ErrCode = 4040
	ErrGetUserOrganizations                 ErrCode = 4041
	ErrGetSeasons                           ErrCode = 4042
	ErrGetUserBySubject                     ErrCode = 4043
	ErrEmailAddressNotVerified              ErrCode = 4044
	ErrGetDefaultSeason                     ErrCode = 4045
	ErrCommitUserTransaction                ErrCode = 4046
	ErrUpdateActiveSeason                   ErrCode = 4047
	ErrMissingContextMetadata               ErrCode = 4048
	ErrNoTokenProvided                      ErrCode = 4049
	ErrGetTokenWithClaims                   ErrCode = 4050
	ErrNoTokenInContext                     ErrCode = 4051
	ErrGetSubjectFromToken                  ErrCode = 4052
	ErrGetSubjectFromContext                ErrCode = 4053
	ErrGetActiveSeasonMembership            ErrCode = 4054
	ErrGetTokenFromContext                  ErrCode = 4055
	ErrChallengeAlreadyClosed               ErrCode = 4056
	ErrGetAgent                             ErrCode = 4057
	ErrSaveAgent                            ErrCode = 4058
	ErrListChallengeInstances               ErrCode = 4059
	ErrChallengeAdd                         ErrCode = 4060
	ErrServerListen                         ErrCode = 5001
	ErrServerRegisterGateway                ErrCode = 5002
	ErrInitLogger                           ErrCode = 6001
	ErrStartService                         ErrCode = 6002
	ErrInitServer                           ErrCode = 6003
	ErrGroupTerminated                      ErrCode = 6004
	ErrGetSSOClientFromFlags                ErrCode = 6005
	ErrDumpDatabase                         ErrCode = 6006
	ErrGetDBInfo                            ErrCode = 6007
	ErrGetSSOWhoami                         ErrCode = 6008
	ErrGetSSOLogout                         ErrCode = 6009
	ErrGetSSOClaims                         ErrCode = 6010
	ErrInitDockerClient                     ErrCode = 6011
	ErrInitDB                               ErrCode = 6012
	ErrConfigureDB                          ErrCode = 6013
	ErrInitSSOClient                        ErrCode = 6014
	ErrInitService                          ErrCode = 6015
	ErrAgentGetContainersInfo               ErrCode = 7001
	ErrCheckNginxContainer                  ErrCode = 7002
	ErrRemoveNginxContainer                 ErrCode = 7003
	ErrBuildNginxContainer                  ErrCode = 7004
	ErrStartNginxContainer                  ErrCode = 7005
	ErrParsingTemplate                      ErrCode = 7006
	ErrWriteConfigFileHeader                ErrCode = 7007
	ErrWriteConfigFile                      ErrCode = 7008
	ErrCloseTarWriter                       ErrCode = 7009
	ErrCopyNginxConfigToContainer           ErrCode = 7010
	ErrNginxNewConfigCheckFailed            ErrCode = 7011
	ErrNginxSendCommandNewConfigCheck       ErrCode = 7012
	ErrNginxSendCommandNewConfigRemove      ErrCode = 7013
	ErrNginxSendCommandConfigReplace        ErrCode = 7014
	ErrNginxSendCommandReloadConfig         ErrCode = 7015
	ErrNginxConnectNetwork                  ErrCode = 7016
	ErrContainerConnectNetwork              ErrCode = 7017
	ErrNatPortOpening                       ErrCode = 7018
	ErrBuildNginxConfig                     ErrCode = 7019
	ErrExecuteTemplate                      ErrCode = 7020
	ErrWriteBytesToHashBuilder              ErrCode = 7021
	ErrReadBytesFromHashBuilder             ErrCode = 7022
	ErrGeneratePrefixHash                   ErrCode = 7023
	ErrCleanPathwarInstances                ErrCode = 7024
	ErrParseInitConfig                      ErrCode = 7025
	ErrUpPathwarInstance                    ErrCode = 7026
	ErrUpdateNginx                          ErrCode = 7027
	ErrDockerAPIContainerList               ErrCode = 8001
	ErrDockerAPIContainerRemove             ErrCode = 8002
	ErrDockerAPIImageRemove                 ErrCode = 8003
	ErrDockerAPIContainerCreate             ErrCode = 8004
	ErrDockerAPIContainerExecCreate         ErrCode = 8005
	ErrDockerAPIContainerExecAttach         ErrCode = 8006
	ErrDockerAPIContainerExecStart          ErrCode = 8007
	ErrDockerAPIContainerExecInspect        ErrCode = 8008
	ErrDockerAPIImagePull                   ErrCode = 8009
	ErrDockerAPIImageInspect                ErrCode = 8010
	ErrDockerAPINetworkList                 ErrCode = 8011
	ErrDockerAPINetworkCreate               ErrCode = 8012
	ErrDockerAPINetworkRemove               ErrCode = 8013
	ErrDockerAPIExitCode                    ErrCode = 8014
	ErrExecuteOnInitHook                    ErrCode = 9001
	ErrRemoveInitConfig                     ErrCode = 9002
)

var ErrCode_name = map[int32]string{
	0:     "Undefined",
	666:   "TODO",
	777:   "ErrNotImplemented",
	888:   "ErrDeprecated",
	999:   "ErrInternal",
	101:   "ErrInvalidInput",
	102:   "ErrMissingInput",
	103:   "ErrUnauthenticated",
	1001:  "ErrSSOGetOIDC",
	1002:  "ErrSSOInvalidPublicKey",
	1003:  "ErrSSOFailedKeycloakRequest",
	1004:  "ErrSSOInvalidKeycloakResponse",
	1005:  "ErrSSOLogout",
	1006:  "ErrSSOInitKeycloakClient",
	1007:  "ErrSSOInvalidBearer",
	1008:  "ErrSSOKeycloakError",
	2001:  "ErrDBNotFound",
	2002:  "ErrDBInternal",
	2003:  "ErrDBRunMigrations",
	2004:  "ErrDBInit",
	2005:  "ErrDBConnect",
	2006:  "ErrDBAutoMigrate",
	2007:  "ErrDBAddForeignKey",
	3001:  "ErrComposeInvalidPath",
	3002:  "ErrComposeDirectoryNotFound",
	3003:  "ErrComposeReadConfig",
	3004:  "ErrComposeInvalidConfig",
	3005:  "ErrComposeMarshalConfig",
	3006:  "ErrComposeCreateTempFile",
	3007:  "ErrComposeWriteTempFile",
	3008:  "ErrComposeCloseTempFile",
	3009:  "ErrComposeBuild",
	30010: "ErrComposeBundle",
	3011:  "ErrComposeReadDab",
	3012:  "ErrComposeParseDab",
	3013:  "ErrComposeParseConfig",
	3014:  "ErrComposeCreateTempDir",
	3015:  "ErrComposeUpdateTempFile",
	3016:  "ErrComposeForceRecreateDown",
	3017:  "ErrComposeRunCreate",
	3018:  "ErrComposeRunUp",
	3019:  "ErrGetPWInitBinary",
	3020:  "ErrWritePWInitFileHeader",
	3021:  "ErrWritePWInitFile",
	3022:  "ErrMarshalPWInitConfigFile",
	3023:  "ErrWritePWInitConfigFileHeader",
	3024:  "ErrWritePWInitConfigFile",
	3025:  "ErrWritePWInitCloseTarWriter",
	3026:  "ErrCopyPWInitToContainer",
	3027:  "ErrComposeGetContainersInfo",
	3028:  "ErrMissingPwinitConfig",
	4001:  "ErrGetUserIDFromContext",
	4002:  "ErrMissingChallengeValidation",
	4003:  "ErrInvalidSeason",
	4004:  "ErrTeamNotInSeason",
	4005:  "ErrGetChallengeSubscription",
	4006:  "ErrInitSnowflake",
	4007:  "ErrUpdateChallengeSubscription",
	4008:  "ErrCreateChallengeValidation",
	4009:  "ErrGetChallengeValidation",
	4010:  "ErrInvalidTeam",
	4011:  "ErrChallengeAlreadySubscribed",
	4012:  "ErrCreateChallengeSubscription",
	4013:  "ErrFindOrganizations",
	4014:  "ErrGetSeasonFromSeasonChallenge",
	4015:  "ErrGetUserTeamFromSeason",
	4016:  "ErrInvalidSeasonID",
	4017:  "ErrUserHasNoTeamForSeason",
	4018:  "ErrGetSeasonChallenges",
	4019:  "ErrGetSeason",
	4020:  "ErrSeasonDenied",
	4021:  "ErrAlreadyHasTeamForSeason",
	4022:  "ErrReservedName",
	4023:  "ErrCheckOrganizationUniqueName",
	4024:  "ErrCreateOrganization",
	4025:  "ErrOrganizationAlreadyHasTeamForSeason",
	4026:  "ErrGetOrganization",
	4027:  "ErrGetSeasonChallenge",
	4028:  "ErrCannotCreateTeamForSoloOrganization",
	4029:  "ErrUserNotInOrganization",
	4030:  "ErrCreateTeam",
	4031:  "ErrGetTeam",
	4032:  "ErrGetTeams",
	4033:  "ErrGetUser",
	4034:  "ErrUpdateUser",
	4035:  "ErrUpdateTeam",
	4036:  "ErrUpdateOrganization",
	4037:  "ErrNewUserFromClaims",
	4038:  "ErrGetOAuthUser",
	4039:  "ErrDifferentUserBetweenTokenAndDatabase",
	4040:  "ErrLoadUserSeasons",
	4041:  "ErrGetUserOrganizations",
	4042:  "ErrGetSeasons",
	4043:  "ErrGetUserBySubject",
	4044:  "ErrEmailAddressNotVerified",
	4045:  "ErrGetDefaultSeason",
	4046:  "ErrCommitUserTransaction",
	4047:  "ErrUpdateActiveSeason",
	4048:  "ErrMissingContextMetadata",
	4049:  "ErrNoTokenProvided",
	4050:  "ErrGetTokenWithClaims",
	4051:  "ErrNoTokenInContext",
	4052:  "ErrGetSubjectFromToken",
	4053:  "ErrGetSubjectFromContext",
	4054:  "ErrGetActiveSeasonMembership",
	4055:  "ErrGetTokenFromContext",
	4056:  "ErrChallengeAlreadyClosed",
	4057:  "ErrGetAgent",
	4058:  "ErrSaveAgent",
	4059:  "ErrListChallengeInstances",
	4060:  "ErrChallengeAdd",
	5001:  "ErrServerListen",
	5002:  "ErrServerRegisterGateway",
	6001:  "ErrInitLogger",
	6002:  "ErrStartService",
	6003:  "ErrInitServer",
	6004:  "ErrGroupTerminated",
	6005:  "ErrGetSSOClientFromFlags",
	6006:  "ErrDumpDatabase",
	6007:  "ErrGetDBInfo",
	6008:  "ErrGetSSOWhoami",
	6009:  "ErrGetSSOLogout",
	6010:  "ErrGetSSOClaims",
	6011:  "ErrInitDockerClient",
	6012:  "ErrInitDB",
	6013:  "ErrConfigureDB",
	6014:  "ErrInitSSOClient",
	6015:  "ErrInitService",
	7001:  "ErrAgentGetContainersInfo",
	7002:  "ErrCheckNginxContainer",
	7003:  "ErrRemoveNginxContainer",
	7004:  "ErrBuildNginxContainer",
	7005:  "ErrStartNginxContainer",
	7006:  "ErrParsingTemplate",
	7007:  "ErrWriteConfigFileHeader",
	7008:  "ErrWriteConfigFile",
	7009:  "ErrCloseTarWriter",
	7010:  "ErrCopyNginxConfigToContainer",
	7011:  "ErrNginxNewConfigCheckFailed",
	7012:  "ErrNginxSendCommandNewConfigCheck",
	7013:  "ErrNginxSendCommandNewConfigRemove",
	7014:  "ErrNginxSendCommandConfigReplace",
	7015:  "ErrNginxSendCommandReloadConfig",
	7016:  "ErrNginxConnectNetwork",
	7017:  "ErrContainerConnectNetwork",
	7018:  "ErrNatPortOpening",
	7019:  "ErrBuildNginxConfig",
	7020:  "ErrExecuteTemplate",
	7021:  "ErrWriteBytesToHashBuilder",
	7022:  "ErrReadBytesFromHashBuilder",
	7023:  "ErrGeneratePrefixHash",
	7024:  "ErrCleanPathwarInstances",
	7025:  "ErrParseInitConfig",
	7026:  "ErrUpPathwarInstance",
	7027:  "ErrUpdateNginx",
	8001:  "ErrDockerAPIContainerList",
	8002:  "ErrDockerAPIContainerRemove",
	8003:  "ErrDockerAPIImageRemove",
	8004:  "ErrDockerAPIContainerCreate",
	8005:  "ErrDockerAPIContainerExecCreate",
	8006:  "ErrDockerAPIContainerExecAttach",
	8007:  "ErrDockerAPIContainerExecStart",
	8008:  "ErrDockerAPIContainerExecInspect",
	8009:  "ErrDockerAPIImagePull",
	8010:  "ErrDockerAPIImageInspect",
	8011:  "ErrDockerAPINetworkList",
	8012:  "ErrDockerAPINetworkCreate",
	8013:  "ErrDockerAPINetworkRemove",
	8014:  "ErrDockerAPIExitCode",
	9001:  "ErrExecuteOnInitHook",
	9002:  "ErrRemoveInitConfig",
}

var ErrCode_value = map[string]int32{
	"Undefined":                               0,
	"TODO":                                    666,
	"ErrNotImplemented":                       777,
	"ErrDeprecated":                           888,
	"ErrInternal":                             999,
	"ErrInvalidInput":                         101,
	"ErrMissingInput":                         102,
	"ErrUnauthenticated":                      103,
	"ErrSSOGetOIDC":                           1001,
	"ErrSSOInvalidPublicKey":                  1002,
	"ErrSSOFailedKeycloakRequest":             1003,
	"ErrSSOInvalidKeycloakResponse":           1004,
	"ErrSSOLogout":                            1005,
	"ErrSSOInitKeycloakClient":                1006,
	"ErrSSOInvalidBearer":                     1007,
	"ErrSSOKeycloakError":                     1008,
	"ErrDBNotFound":                           2001,
	"ErrDBInternal":                           2002,
	"ErrDBRunMigrations":                      2003,
	"ErrDBInit":                               2004,
	"ErrDBConnect":                            2005,
	"ErrDBAutoMigrate":                        2006,
	"ErrDBAddForeignKey":                      2007,
	"ErrComposeInvalidPath":                   3001,
	"ErrComposeDirectoryNotFound":             3002,
	"ErrComposeReadConfig":                    3003,
	"ErrComposeInvalidConfig":                 3004,
	"ErrComposeMarshalConfig":                 3005,
	"ErrComposeCreateTempFile":                3006,
	"ErrComposeWriteTempFile":                 3007,
	"ErrComposeCloseTempFile":                 3008,
	"ErrComposeBuild":                         3009,
	"ErrComposeBundle":                        30010,
	"ErrComposeReadDab":                       3011,
	"ErrComposeParseDab":                      3012,
	"ErrComposeParseConfig":                   3013,
	"ErrComposeCreateTempDir":                 3014,
	"ErrComposeUpdateTempFile":                3015,
	"ErrComposeForceRecreateDown":             3016,
	"ErrComposeRunCreate":                     3017,
	"ErrComposeRunUp":                         3018,
	"ErrGetPWInitBinary":                      3019,
	"ErrWritePWInitFileHeader":                3020,
	"ErrWritePWInitFile":                      3021,
	"ErrMarshalPWInitConfigFile":              3022,
	"ErrWritePWInitConfigFileHeader":          3023,
	"ErrWritePWInitConfigFile":                3024,
	"ErrWritePWInitCloseTarWriter":            3025,
	"ErrCopyPWInitToContainer":                3026,
	"ErrComposeGetContainersInfo":             3027,
	"ErrMissingPwinitConfig":                  3028,
	"ErrGetUserIDFromContext":                 4001,
	"ErrMissingChallengeValidation":           4002,
	"ErrInvalidSeason":                        4003,
	"ErrTeamNotInSeason":                      4004,
	"ErrGetChallengeSubscription":             4005,
	"ErrInitSnowflake":                        4006,
	"ErrUpdateChallengeSubscription":          4007,
	"ErrCreateChallengeValidation":            4008,
	"ErrGetChallengeValidation":               4009,
	"ErrInvalidTeam":                          4010,
	"ErrChallengeAlreadySubscribed":           4011,
	"ErrCreateChallengeSubscription":          4012,
	"ErrFindOrganizations":                    4013,
	"ErrGetSeasonFromSeasonChallenge":         4014,
	"ErrGetUserTeamFromSeason":                4015,
	"ErrInvalidSeasonID":                      4016,
	"ErrUserHasNoTeamForSeason":               4017,
	"ErrGetSeasonChallenges":                  4018,
	"ErrGetSeason":                            4019,
	"ErrSeasonDenied":                         4020,
	"ErrAlreadyHasTeamForSeason":              4021,
	"ErrReservedName":                         4022,
	"ErrCheckOrganizationUniqueName":          4023,
	"ErrCreateOrganization":                   4024,
	"ErrOrganizationAlreadyHasTeamForSeason":  4025,
	"ErrGetOrganization":                      4026,
	"ErrGetSeasonChallenge":                   4027,
	"ErrCannotCreateTeamForSoloOrganization":  4028,
	"ErrUserNotInOrganization":                4029,
	"ErrCreateTeam":                           4030,
	"ErrGetTeam":                              4031,
	"ErrGetTeams":                             4032,
	"ErrGetUser":                              4033,
	"ErrUpdateUser":                           4034,
	"ErrUpdateTeam":                           4035,
	"ErrUpdateOrganization":                   4036,
	"ErrNewUserFromClaims":                    4037,
	"ErrGetOAuthUser":                         4038,
	"ErrDifferentUserBetweenTokenAndDatabase": 4039,
	"ErrLoadUserSeasons":                      4040,
	"ErrGetUserOrganizations":                 4041,
	"ErrGetSeasons":                           4042,
	"ErrGetUserBySubject":                     4043,
	"ErrEmailAddressNotVerified":              4044,
	"ErrGetDefaultSeason":                     4045,
	"ErrCommitUserTransaction":                4046,
	"ErrUpdateActiveSeason":                   4047,
	"ErrMissingContextMetadata":               4048,
	"ErrNoTokenProvided":                      4049,
	"ErrGetTokenWithClaims":                   4050,
	"ErrNoTokenInContext":                     4051,
	"ErrGetSubjectFromToken":                  4052,
	"ErrGetSubjectFromContext":                4053,
	"ErrGetActiveSeasonMembership":            4054,
	"ErrGetTokenFromContext":                  4055,
	"ErrChallengeAlreadyClosed":               4056,
	"ErrGetAgent":                             4057,
	"ErrSaveAgent":                            4058,
	"ErrListChallengeInstances":               4059,
	"ErrChallengeAdd":                         4060,
	"ErrServerListen":                         5001,
	"ErrServerRegisterGateway":                5002,
	"ErrInitLogger":                           6001,
	"ErrStartService":                         6002,
	"ErrInitServer":                           6003,
	"ErrGroupTerminated":                      6004,
	"ErrGetSSOClientFromFlags":                6005,
	"ErrDumpDatabase":                         6006,
	"ErrGetDBInfo":                            6007,
	"ErrGetSSOWhoami":                         6008,
	"ErrGetSSOLogout":                         6009,
	"ErrGetSSOClaims":                         6010,
	"ErrInitDockerClient":                     6011,
	"ErrInitDB":                               6012,
	"ErrConfigureDB":                          6013,
	"ErrInitSSOClient":                        6014,
	"ErrInitService":                          6015,
	"ErrAgentGetContainersInfo":               7001,
	"ErrCheckNginxContainer":                  7002,
	"ErrRemoveNginxContainer":                 7003,
	"ErrBuildNginxContainer":                  7004,
	"ErrStartNginxContainer":                  7005,
	"ErrParsingTemplate":                      7006,
	"ErrWriteConfigFileHeader":                7007,
	"ErrWriteConfigFile":                      7008,
	"ErrCloseTarWriter":                       7009,
	"ErrCopyNginxConfigToContainer":           7010,
	"ErrNginxNewConfigCheckFailed":            7011,
	"ErrNginxSendCommandNewConfigCheck":       7012,
	"ErrNginxSendCommandNewConfigRemove":      7013,
	"ErrNginxSendCommandConfigReplace":        7014,
	"ErrNginxSendCommandReloadConfig":         7015,
	"ErrNginxConnectNetwork":                  7016,
	"ErrContainerConnectNetwork":              7017,
	"ErrNatPortOpening":                       7018,
	"ErrBuildNginxConfig":                     7019,
	"ErrExecuteTemplate":                      7020,
	"ErrWriteBytesToHashBuilder":              7021,
	"ErrReadBytesFromHashBuilder":             7022,
	"ErrGeneratePrefixHash":                   7023,
	"ErrCleanPathwarInstances":                7024,
	"ErrParseInitConfig":                      7025,
	"ErrUpPathwarInstance":                    7026,
	"ErrUpdateNginx":                          7027,
	"ErrDockerAPIContainerList":               8001,
	"ErrDockerAPIContainerRemove":             8002,
	"ErrDockerAPIImageRemove":                 8003,
	"ErrDockerAPIContainerCreate":             8004,
	"ErrDockerAPIContainerExecCreate":         8005,
	"ErrDockerAPIContainerExecAttach":         8006,
	"ErrDockerAPIContainerExecStart":          8007,
	"ErrDockerAPIContainerExecInspect":        8008,
	"ErrDockerAPIImagePull":                   8009,
	"ErrDockerAPIImageInspect":                8010,
	"ErrDockerAPINetworkList":                 8011,
	"ErrDockerAPINetworkCreate":               8012,
	"ErrDockerAPINetworkRemove":               8013,
	"ErrDockerAPIExitCode":                    8014,
	"ErrExecuteOnInitHook":                    9001,
	"ErrRemoveInitConfig":                     9002,
}

func (x ErrCode) String() string {
	return proto.EnumName(ErrCode_name, int32(x))
}

func (ErrCode) EnumDescriptor() ([]byte, []int) {
	return fileDescriptor_4240057316120df7, []int{0}
}

func init() {
	proto.RegisterEnum("pathwar.errcode.ErrCode", ErrCode_name, ErrCode_value)
}

func init() { proto.RegisterFile("errcode.proto", fileDescriptor_4240057316120df7) }

var fileDescriptor_4240057316120df7 = []byte{
	// 2133 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0x7c, 0x58, 0x59, 0x6f, 0x1c, 0xc7,
	0x11, 0x96, 0x80, 0xc4, 0x84, 0x27, 0xb1, 0x59, 0x1e, 0xd9, 0x5a, 0x5b, 0xb2, 0x39, 0xb6, 0x13,
	0x5b, 0x40, 0x8e, 0xd5, 0x43, 0x80, 0x05, 0xf2, 0x42, 0x80, 0xcb, 0x25, 0xa9, 0x85, 0x25, 0x92,
	0xe0, 0x61, 0x01, 0x79, 0x6b, 0xee, 0xd4, 0xce, 0x76, 0x38, 0xdb, 0xbd, 0xee, 0xe9, 0x21, 0xc5,
	0xfc, 0x03, 0xe7, 0x29, 0xcf, 0x79, 0x08, 0x90, 0x3b, 0x76, 0xee, 0x3b, 0x96, 0x4f, 0x59, 0xf7,
	0x69, 0x5d, 0xbe, 0x9d, 0xc4, 0x96, 0x92, 0x38, 0x96, 0x6f, 0x3b, 0x87, 0x73, 0x07, 0xdd, 0x53,
	0x3d, 0x3b, 0xb3, 0xa4, 0xfc, 0x46, 0xd6, 0x57, 0x55, 0xdd, 0xf5, 0xd5, 0xd1, 0x35, 0xeb, 0x5d,
	0x87, 0x4a, 0xb5, 0x64, 0x88, 0xd5, 0x9e, 0x92, 0x5a, 0xfa, 0xc3, 0x3d, 0xa6, 0x3b, 0xab, 0x4c,
	0x55, 0x49, 0xbc, 0xed, 0xb3, 0x11, 0xd7, 0x9d, 0x74, 0xa9, 0xda, 0x92, 0xdd, 0x9d, 0x91, 0x8c,
	0xe4, 0x4e, 0xab, 0xb7, 0x94, 0xb6, 0xed, 0x7f, 0xf6, 0x1f, 0xfb, 0x57, 0x66, 0xff, 0xa9, 0xaf,
	0xdd, 0xe5, 0x0d, 0x4d, 0x28, 0x35, 0x2e, 0x43, 0xf4, 0xaf, 0xf3, 0xae, 0x5d, 0x14, 0x21, 0xb6,
	0xb9, 0xc0, 0x10, 0x36, 0xf9, 0xd7, 0x7a, 0x1f, 0x59, 0x98, 0x69, 0xcc, 0xc0, 0x57, 0x3f, 0xea,
	0x6f, 0xf5, 0x6e, 0x98, 0x50, 0x6a, 0x5a, 0xea, 0x66, 0xb7, 0x17, 0x63, 0x17, 0x85, 0xc6, 0x10,
	0xee, 0xbf, 0xc6, 0xf7, 0xbd, 0xeb, 0x26, 0x94, 0x6a, 0x60, 0x4f, 0x61, 0x8b, 0x19, 0xd9, 0x07,
	0xd7, 0xf8, 0xe0, 0x7d, 0x6c, 0x42, 0xa9, 0xa6, 0xd0, 0xa8, 0x04, 0x8b, 0xe1, 0xd5, 0x21, 0x7f,
	0x8b, 0x37, 0x6c, 0x25, 0x2b, 0x2c, 0xe6, 0x61, 0x53, 0xf4, 0x52, 0x0d, 0x48, 0xc2, 0x3d, 0x3c,
	0x49, 0xb8, 0x88, 0x32, 0x61, 0xdb, 0xdf, 0xea, 0xf9, 0x13, 0x4a, 0x2d, 0x0a, 0x96, 0xea, 0x0e,
	0x0a, 0xcd, 0x33, 0xa7, 0x11, 0x9d, 0x33, 0x3f, 0x3f, 0x33, 0x85, 0x7a, 0xa6, 0xd9, 0x18, 0x87,
	0xd7, 0x86, 0xfc, 0xed, 0xde, 0xd6, 0x4c, 0x46, 0x8e, 0x67, 0xd3, 0xa5, 0x98, 0xb7, 0xee, 0xc1,
	0x35, 0xb8, 0x32, 0xe4, 0xdf, 0xee, 0x6d, 0xcf, 0xc0, 0x49, 0xc6, 0x63, 0x0c, 0xef, 0xc1, 0xb5,
	0x56, 0x2c, 0xd9, 0xf2, 0x1c, 0xde, 0x97, 0x62, 0xa2, 0xe1, 0xf5, 0x21, 0xff, 0x4e, 0xef, 0xb6,
	0x92, 0x79, 0x5f, 0x25, 0xe9, 0x49, 0x91, 0x20, 0xbc, 0x31, 0xe4, 0xdf, 0xe0, 0x7d, 0x3c, 0xd3,
	0xd9, 0x2d, 0x23, 0x99, 0x6a, 0x78, 0x73, 0xc8, 0xbf, 0xcd, 0xbb, 0xd9, 0x99, 0x71, 0xed, 0x6c,
	0xc6, 0x63, 0x8e, 0x42, 0xc3, 0x5b, 0x43, 0xfe, 0xcd, 0xde, 0x96, 0x92, 0xd7, 0x3a, 0x32, 0x85,
	0x0a, 0xde, 0x2e, 0x20, 0xce, 0x68, 0x42, 0x29, 0xa9, 0xe0, 0x9d, 0x21, 0x47, 0x62, 0x7d, 0x5a,
	0xea, 0x49, 0x99, 0x8a, 0x10, 0xce, 0x0d, 0xe7, 0xb2, 0x9c, 0xc6, 0xf3, 0xc3, 0x7e, 0xc5, 0x92,
	0xd3, 0xa8, 0xcf, 0xa5, 0x62, 0x0f, 0x8f, 0x14, 0xd3, 0x5c, 0x8a, 0x04, 0x2e, 0x0c, 0xfb, 0xd7,
	0x7b, 0xd7, 0x92, 0x32, 0xd7, 0x70, 0x71, 0x98, 0xae, 0xdd, 0xa8, 0x8f, 0x4b, 0x21, 0xb0, 0xa5,
	0xe1, 0xe9, 0x61, 0xff, 0x26, 0x0f, 0xac, 0x68, 0x2c, 0xd5, 0x32, 0x33, 0x46, 0x78, 0xa6, 0xef,
	0x72, 0x2c, 0x0c, 0x27, 0xa5, 0x42, 0x1e, 0x09, 0xc3, 0xdf, 0xb3, 0xc3, 0xfe, 0x36, 0xef, 0x26,
	0x5b, 0x15, 0xdd, 0x9e, 0x4c, 0xd0, 0x11, 0xcc, 0x74, 0x07, 0x1e, 0xaa, 0x10, 0xb7, 0x84, 0x35,
	0xb8, 0xc2, 0x96, 0x96, 0x6a, 0x2d, 0xbf, 0xfd, 0xfe, 0x8a, 0x7f, 0x8b, 0x77, 0x63, 0x5f, 0x63,
	0x0e, 0x59, 0x38, 0x2e, 0x45, 0x9b, 0x47, 0xf0, 0x70, 0xc5, 0xbf, 0xd5, 0xab, 0xac, 0x73, 0x4c,
	0xe8, 0x23, 0x03, 0xe8, 0x1e, 0xa6, 0x92, 0x0e, 0x8b, 0x09, 0x7d, 0xb4, 0x42, 0xdc, 0x13, 0x3a,
	0xae, 0x90, 0x69, 0x5c, 0xc0, 0x6e, 0x6f, 0x92, 0xc7, 0x08, 0x8f, 0x0d, 0x18, 0xef, 0x55, 0xbc,
	0x80, 0x3e, 0x3e, 0x80, 0x8e, 0xc7, 0x32, 0xe9, 0xa3, 0x4f, 0x54, 0xfc, 0x1b, 0x6d, 0x35, 0x12,
	0x5a, 0x4f, 0x79, 0x1c, 0xc2, 0x81, 0x8a, 0xbf, 0xd5, 0xb2, 0x96, 0x4b, 0x45, 0x18, 0x23, 0xec,
	0xbf, 0xb2, 0x99, 0xda, 0xa1, 0x10, 0x5f, 0x83, 0x2d, 0xc1, 0xc1, 0x0a, 0xd1, 0x49, 0xf2, 0x59,
	0xa6, 0x12, 0x34, 0xc0, 0xa1, 0x4a, 0x99, 0x4e, 0x0b, 0x50, 0x54, 0x87, 0x07, 0x2f, 0x96, 0x47,
	0xd5, 0xe0, 0x0a, 0x8e, 0x0c, 0xc4, 0xbc, 0xd8, 0x0b, 0x8b, 0x31, 0x1f, 0x1d, 0xc8, 0xc5, 0xa4,
	0x54, 0x2d, 0x9c, 0xc3, 0x96, 0xf5, 0xd1, 0x90, 0xab, 0x02, 0x8e, 0x55, 0xa8, 0xee, 0xdc, 0x5d,
	0x53, 0x91, 0x9d, 0x00, 0xc7, 0x07, 0x62, 0x9e, 0x4b, 0xc5, 0x62, 0x0f, 0x4e, 0xb8, 0x18, 0xa6,
	0x50, 0xcf, 0xee, 0x35, 0xf5, 0x54, 0xe7, 0x82, 0xa9, 0x35, 0x38, 0xe9, 0x6e, 0x62, 0x79, 0xcd,
	0x20, 0x73, 0x87, 0x5d, 0xc8, 0x42, 0x54, 0x70, 0xca, 0xd9, 0x0d, 0xc0, 0x70, 0xba, 0xe2, 0x07,
	0xde, 0x36, 0xd3, 0xe8, 0x59, 0x32, 0x33, 0x28, 0x0b, 0xde, 0x2a, 0x9c, 0xa9, 0xf8, 0x9f, 0xf0,
	0x46, 0xca, 0x96, 0x7d, 0x98, 0xdc, 0x3f, 0xb5, 0xc1, 0xe9, 0x05, 0x1f, 0x67, 0x2b, 0xfe, 0x1d,
	0xde, 0xad, 0x03, 0xb0, 0xcd, 0x30, 0xcb, 0x44, 0x0a, 0xce, 0xf5, 0x99, 0xec, 0xad, 0x65, 0x1a,
	0x0b, 0x72, 0x5c, 0x0a, 0xcd, 0xb8, 0x40, 0x05, 0xe7, 0x07, 0x98, 0x9c, 0x42, 0x9d, 0x83, 0x49,
	0x53, 0xb4, 0x25, 0x5c, 0xa8, 0xd0, 0xc0, 0xa1, 0x89, 0x35, 0xbb, 0xca, 0xf3, 0x4b, 0xc0, 0x45,
	0x97, 0xc5, 0x29, 0xd4, 0x8b, 0x09, 0xaa, 0x66, 0x63, 0x52, 0xc9, 0xae, 0xf1, 0x80, 0xfb, 0x34,
	0x7c, 0x3d, 0xa0, 0x61, 0x43, 0xa6, 0xe3, 0x1d, 0x16, 0xc7, 0x28, 0x22, 0xbc, 0xd7, 0x14, 0xbf,
	0x6d, 0x63, 0xf8, 0x46, 0x40, 0x2d, 0x4a, 0x2d, 0x31, 0x8f, 0x2c, 0x91, 0x02, 0xbe, 0x19, 0x10,
	0xaf, 0x0b, 0xc8, 0xba, 0x66, 0xfc, 0x0a, 0x02, 0xbe, 0x15, 0xd0, 0x85, 0xcd, 0x4d, 0x9d, 0xbf,
	0xf9, 0x74, 0x29, 0x69, 0x29, 0xde, 0xb3, 0x1e, 0xbf, 0xdd, 0xf7, 0xc8, 0xf5, 0xbc, 0x90, 0xab,
	0xed, 0x98, 0x2d, 0x23, 0x7c, 0x27, 0x20, 0xbe, 0xb3, 0x5a, 0xda, 0xd8, 0xf6, 0xbb, 0x01, 0x11,
	0x9a, 0x15, 0xcb, 0x46, 0x17, 0xfe, 0x5e, 0xe0, 0x8f, 0x78, 0xb7, 0x0c, 0x5c, 0xa0, 0x80, 0x3f,
	0x10, 0xf8, 0x5b, 0xbc, 0xeb, 0xfb, 0x01, 0x99, 0x00, 0xe0, 0x41, 0xc7, 0x44, 0x6e, 0x31, 0x16,
	0x2b, 0x64, 0xe1, 0x1a, 0x9d, 0xbe, 0x84, 0x21, 0x7c, 0xdf, 0x5d, 0x70, 0xe0, 0xec, 0xd2, 0x05,
	0x7f, 0x10, 0xd0, 0x8c, 0x99, 0xe4, 0x22, 0x9c, 0x51, 0x11, 0x13, 0xfc, 0x4b, 0x34, 0x0f, 0x7f,
	0x18, 0xf8, 0x9f, 0xf4, 0x82, 0xec, 0x62, 0x19, 0x59, 0x26, 0x17, 0xd9, 0x5f, 0xb9, 0x33, 0xf8,
	0x51, 0x40, 0xf5, 0x40, 0x19, 0x33, 0xd7, 0xeb, 0xeb, 0xc1, 0x8f, 0x1d, 0xef, 0xa5, 0x74, 0x34,
	0x1b, 0xf0, 0x13, 0x17, 0xb6, 0x31, 0xda, 0xc5, 0x92, 0x69, 0x69, 0x2d, 0xa5, 0x22, 0xc3, 0x9f,
	0x06, 0x54, 0x26, 0xf9, 0xe9, 0xf9, 0x99, 0x09, 0xfc, 0x2c, 0xa0, 0xd1, 0x9c, 0x83, 0xf0, 0xf3,
	0x80, 0xda, 0x30, 0xfb, 0xbf, 0x81, 0x82, 0x63, 0x08, 0xbf, 0x08, 0xa8, 0x6b, 0x88, 0x9e, 0x5d,
	0x2c, 0x29, 0x1f, 0xf3, 0x4b, 0x67, 0x36, 0x87, 0x09, 0xaa, 0x15, 0x0c, 0xa7, 0x59, 0x17, 0xe1,
	0x57, 0x39, 0x75, 0x1d, 0x6c, 0x2d, 0x17, 0x69, 0x59, 0x14, 0xfc, 0xbe, 0x14, 0xad, 0xd2, 0xaf,
	0x03, 0x37, 0x8d, 0x2c, 0xbf, 0x45, 0x2d, 0xf8, 0x4d, 0xe0, 0x7f, 0xda, 0xbb, 0x7b, 0x42, 0xa9,
	0xa2, 0xf4, 0x6a, 0x77, 0x78, 0x28, 0xe8, 0xcf, 0x8a, 0x92, 0x97, 0xfd, 0xee, 0x84, 0xf5, 0x1c,
	0xc0, 0xc3, 0xee, 0x84, 0x71, 0x26, 0x84, 0xd4, 0x6e, 0xdc, 0x65, 0x7e, 0x65, 0x2c, 0x4b, 0x8e,
	0x1e, 0x71, 0x49, 0x32, 0x64, 0xdb, 0xea, 0x2f, 0xc1, 0x8f, 0x06, 0xf4, 0x4c, 0xf6, 0xbd, 0xc0,
	0x63, 0x81, 0x3f, 0xec, 0x79, 0xd9, 0xd9, 0x56, 0xf0, 0x78, 0x40, 0x0b, 0x09, 0x09, 0x12, 0x78,
	0xa2, 0xa0, 0x62, 0x1c, 0xc3, 0x01, 0xe7, 0x27, 0x6b, 0x09, 0x2b, 0x7b, 0xb2, 0x2c, 0xb3, 0xae,
	0x0e, 0xba, 0xb8, 0x32, 0x59, 0xe9, 0x2e, 0x87, 0x5c, 0x41, 0x4e, 0xe3, 0xaa, 0x71, 0x60, 0xfb,
	0x3f, 0x66, 0xbc, 0x9b, 0xc0, 0x61, 0x97, 0x2b, 0xc3, 0xd3, 0x58, 0xaa, 0x3b, 0xf6, 0x80, 0x23,
	0x81, 0xff, 0x19, 0x6f, 0x87, 0x79, 0x7c, 0x79, 0xbb, 0x8d, 0x0a, 0x85, 0xbd, 0x4b, 0x1d, 0xf5,
	0x2a, 0xa2, 0x58, 0x90, 0xcb, 0x28, 0xc6, 0x44, 0xd8, 0x60, 0x9a, 0x2d, 0xb1, 0x04, 0xe1, 0xa8,
	0xe3, 0x7a, 0xb7, 0x64, 0xa1, 0x51, 0xcc, 0x78, 0x4d, 0xe0, 0x58, 0x50, 0x9e, 0x3c, 0xe5, 0x5e,
	0x38, 0xee, 0xa2, 0xc8, 0x33, 0x91, 0xc0, 0x89, 0x80, 0x9e, 0x04, 0xb2, 0xa8, 0x9b, 0xe6, 0xfb,
	0xa2, 0x59, 0x13, 0x4e, 0xba, 0xaa, 0x9b, 0xe8, 0x32, 0x1e, 0x8f, 0x85, 0xa1, 0xc2, 0x24, 0x99,
	0x96, 0xfa, 0x5e, 0x54, 0xbc, 0x6d, 0xca, 0xf2, 0x54, 0xc1, 0xb4, 0x81, 0x6d, 0x96, 0xc6, 0xae,
	0x8c, 0x4f, 0x07, 0xfd, 0x87, 0xaa, 0xcb, 0xb3, 0x8e, 0x52, 0x4c, 0x24, 0xac, 0x65, 0xd9, 0x39,
	0x53, 0x66, 0x6e, 0xac, 0xa5, 0xf9, 0x0a, 0x92, 0xe9, 0x53, 0xae, 0xa3, 0xdc, 0x74, 0xcc, 0xa6,
	0xe6, 0x1e, 0xd4, 0x2c, 0x64, 0x9a, 0xc1, 0x59, 0x17, 0xfa, 0xb4, 0xb4, 0xb4, 0xcc, 0x2a, 0xb9,
	0xc2, 0x43, 0x0c, 0xe1, 0x5c, 0xa1, 0xcc, 0x2c, 0xb2, 0x97, 0xeb, 0x0e, 0x71, 0x7e, 0xde, 0xdd,
	0x94, 0x8c, 0x9a, 0xc2, 0x0d, 0xe3, 0x0b, 0xc5, 0x06, 0xcd, 0x02, 0x37, 0xb9, 0xb2, 0x5a, 0x70,
	0xb1, 0x30, 0x15, 0x0a, 0xa0, 0xb3, 0x7d, 0xda, 0x8d, 0xc5, 0x29, 0xd4, 0xc5, 0x18, 0xf6, 0x60,
	0x77, 0x09, 0x55, 0xd2, 0xe1, 0x3d, 0x78, 0xa6, 0xe0, 0xde, 0xfa, 0x2c, 0xda, 0x3f, 0xeb, 0x42,
	0x1d, 0x1c, 0x7f, 0xf6, 0xb1, 0x0a, 0xe1, 0xb9, 0x42, 0xad, 0x8e, 0x45, 0x66, 0xa3, 0x7c, 0xde,
	0x4d, 0x8c, 0x79, 0xb6, 0x82, 0x99, 0xe8, 0x05, 0xe7, 0x64, 0x37, 0x4f, 0xfa, 0x93, 0xb7, 0x29,
	0x12, 0xcd, 0x44, 0x0b, 0x13, 0x78, 0xd1, 0x95, 0x5b, 0xff, 0x90, 0x30, 0x84, 0x97, 0xfa, 0x73,
	0x46, 0xad, 0xa0, 0xb5, 0x45, 0x01, 0xf7, 0xef, 0x70, 0xfb, 0xac, 0x95, 0xce, 0x61, 0x64, 0xe4,
	0x6a, 0x8a, 0x69, 0x5c, 0x65, 0x6b, 0xf0, 0xe5, 0x1d, 0x54, 0x3e, 0xe6, 0x09, 0xd9, 0x2d, 0xa3,
	0x08, 0x15, 0xbc, 0x5b, 0x75, 0x8e, 0x34, 0x53, 0xda, 0xd8, 0xf1, 0x16, 0xc2, 0x7b, 0xd5, 0x82,
	0x66, 0xe6, 0x0c, 0xde, 0xaf, 0xba, 0xf9, 0xa0, 0x64, 0xda, 0x5b, 0x40, 0xd5, 0xe5, 0xc2, 0xae,
	0xf3, 0x7f, 0xa9, 0x16, 0x58, 0x9e, 0x9f, 0xc9, 0x96, 0x67, 0xc3, 0xd3, 0x64, 0xcc, 0xa2, 0x04,
	0xfe, 0xea, 0x4e, 0x68, 0xa4, 0xdd, 0x5e, 0xde, 0x01, 0x7f, 0xab, 0xf6, 0x67, 0xa7, 0xd9, 0x74,
	0xdb, 0x12, 0xfe, 0x5e, 0xed, 0x37, 0xd6, 0xfc, 0xfc, 0xcc, 0xde, 0x8e, 0x64, 0x5d, 0x0e, 0x1f,
	0x94, 0xa5, 0xb4, 0xb9, 0xff, 0xa3, 0x2c, 0xa5, 0x32, 0xf9, 0x67, 0x95, 0xca, 0xc4, 0x5c, 0xbb,
	0x21, 0x5b, 0xcb, 0xa8, 0x68, 0x95, 0xff, 0x57, 0x95, 0xb6, 0x6a, 0x8b, 0xd4, 0xe1, 0xdf, 0x55,
	0x7a, 0xce, 0xb2, 0x17, 0x3f, 0x55, 0xd8, 0xa8, 0xc3, 0x7f, 0xaa, 0xc5, 0x27, 0xd6, 0x45, 0x02,
	0xff, 0xad, 0xe6, 0x4f, 0x1f, 0xcf, 0x19, 0xfa, 0x5f, 0x95, 0xd2, 0x66, 0xb3, 0xb8, 0x7e, 0xbf,
	0x78, 0xbe, 0x46, 0x85, 0x63, 0x67, 0xf7, 0x74, 0xc4, 0xc5, 0xbe, 0xfe, 0x7a, 0xf2, 0x42, 0x8d,
	0xba, 0x7c, 0x0e, 0xbb, 0x72, 0x05, 0x07, 0xd0, 0x17, 0x9d, 0xa9, 0xdd, 0x5b, 0x07, 0xc0, 0x97,
	0x1c, 0x68, 0xf3, 0x35, 0x00, 0xfe, 0xb6, 0x46, 0x29, 0x32, 0x2b, 0x29, 0x17, 0x91, 0xd9, 0x2c,
	0x63, 0xb3, 0x1d, 0xfe, 0xae, 0x56, 0x5c, 0xb8, 0xd6, 0xed, 0x63, 0xbf, 0xaf, 0x15, 0xd7, 0xbd,
	0xc2, 0x26, 0xf6, 0x72, 0xcd, 0xed, 0xc6, 0xe5, 0xf5, 0xeb, 0x95, 0x9a, 0x7b, 0xf8, 0x65, 0x6f,
	0xcd, 0x5d, 0xa2, 0xcd, 0xa3, 0xe2, 0x0e, 0x76, 0xa9, 0x46, 0xdd, 0x65, 0xf1, 0x69, 0x5c, 0xcd,
	0x54, 0x2c, 0x1f, 0xd9, 0x57, 0x1c, 0x5c, 0xae, 0xf9, 0x77, 0x7b, 0x77, 0x38, 0x95, 0x79, 0x14,
	0xa1, 0x99, 0x37, 0x4c, 0x84, 0x65, 0x6d, 0xf8, 0x43, 0xcd, 0xdf, 0xe1, 0xdd, 0xf9, 0x61, 0x7a,
	0x19, 0x91, 0xf0, 0xc7, 0x9a, 0x7f, 0x97, 0x77, 0xfb, 0x06, 0x8a, 0x4e, 0xab, 0x17, 0xb3, 0x16,
	0xc2, 0x9f, 0x6a, 0xb4, 0x53, 0x0c, 0xaa, 0xcd, 0x61, 0x2c, 0xf3, 0xaf, 0x9b, 0x57, 0x1d, 0xd5,
	0x2e, 0x40, 0xf3, 0xf1, 0x35, 0x8d, 0x7a, 0x55, 0xaa, 0x65, 0xf8, 0x73, 0x8d, 0x86, 0x6b, 0x1e,
	0xf0, 0x80, 0xc2, 0x6b, 0x8e, 0xba, 0x69, 0xa6, 0x67, 0xa5, 0xd2, 0x33, 0x3d, 0x14, 0x5c, 0x44,
	0x70, 0xa5, 0x46, 0x35, 0x5a, 0xca, 0xae, 0x39, 0xef, 0x75, 0x97, 0x85, 0x89, 0x7d, 0xd8, 0x4a,
	0xb3, 0xef, 0x02, 0x9b, 0xbd, 0x37, 0xdc, 0x59, 0x96, 0xfd, 0xfa, 0x9a, 0xc6, 0x64, 0x41, 0xee,
	0x62, 0x49, 0xc7, 0xba, 0x40, 0x05, 0x6f, 0xd6, 0x68, 0x7b, 0x34, 0xdf, 0x2e, 0x16, 0x37, 0xed,
	0x57, 0xd4, 0x78, 0xab, 0x96, 0x0f, 0x57, 0x81, 0xe6, 0x6b, 0x71, 0x56, 0x61, 0x9b, 0xef, 0x33,
	0x2a, 0xf0, 0xb6, 0x2b, 0x8e, 0xf1, 0x18, 0x99, 0x98, 0xcd, 0x7e, 0x7f, 0xe8, 0x0f, 0xa0, 0x77,
	0x8a, 0x45, 0x85, 0xfd, 0x55, 0x1d, 0xde, 0xad, 0xd1, 0x1b, 0xb9, 0xd8, 0x1b, 0x30, 0x82, 0xf7,
	0x6a, 0xd4, 0x32, 0xd9, 0x03, 0x61, 0xa3, 0x84, 0xf7, 0x6b, 0xd4, 0x32, 0x59, 0x67, 0x8e, 0xcd,
	0x36, 0x73, 0xee, 0xcc, 0xfc, 0x82, 0x03, 0xa3, 0x14, 0xc5, 0x7a, 0x9c, 0xd2, 0xfb, 0xe4, 0x28,
	0xf5, 0x4d, 0xae, 0xd1, 0xec, 0xb2, 0x08, 0x09, 0x3d, 0x78, 0x75, 0x7b, 0xfa, 0x48, 0x3a, 0x34,
	0x4a, 0x79, 0x5f, 0xaf, 0x61, 0x38, 0x27, 0xad, 0xc3, 0x1f, 0xae, 0x35, 0xa6, 0x35, 0x6b, 0x75,
	0xe0, 0xc8, 0x28, 0x2d, 0x67, 0x1b, 0x6b, 0xd9, 0xf6, 0x84, 0xa3, 0xa3, 0x54, 0x8f, 0x1b, 0x2b,
	0x35, 0x45, 0xd2, 0x33, 0x2f, 0xf5, 0xb1, 0x51, 0xca, 0x4e, 0x39, 0xae, 0xd9, 0x34, 0x8e, 0xe1,
	0xf8, 0x28, 0x65, 0xa7, 0x8c, 0x39, 0xd3, 0x13, 0xeb, 0x28, 0xa1, 0x02, 0xb4, 0x94, 0x9e, 0x1c,
	0x1d, 0xa4, 0x9c, 0x50, 0x0a, 0xf5, 0xd4, 0xd5, 0x70, 0xa2, 0xf4, 0xf4, 0x28, 0xa5, 0x38, 0xc7,
	0x27, 0xf6, 0x99, 0xfc, 0x87, 0x08, 0x67, 0x1c, 0x44, 0xd5, 0x3a, 0x23, 0x4c, 0x69, 0xec, 0x92,
	0x72, 0x19, 0x1e, 0x98, 0xa4, 0x12, 0xcf, 0xbc, 0x14, 0x4a, 0xe6, 0xc1, 0xc9, 0xfa, 0xe7, 0xcf,
	0xbe, 0x32, 0xb2, 0xe9, 0xd8, 0xa5, 0x91, 0xcd, 0x67, 0x2f, 0x8d, 0x6c, 0x7e, 0xf9, 0xd2, 0xc8,
	0xe6, 0xaf, 0x5c, 0x1e, 0xd9, 0x74, 0xf6, 0xf2, 0xc8, 0xa6, 0xe7, 0x2e, 0x8f, 0x6c, 0xfa, 0xc2,
	0x76, 0xf7, 0xf3, 0x57, 0xcc, 0x44, 0xb8, 0x33, 0x92, 0x3b, 0x7b, 0xcb, 0xd1, 0x4e, 0xfa, 0x29,
	0x6c, 0xe9, 0x1a, 0xfb, 0x13, 0xd7, 0xe7, 0xfe, 0x1f, 0x00, 0x00, 0xff, 0xff, 0xe7, 0xbe, 0x4d,
	0xb4, 0x33, 0x13, 0x00, 0x00,
}
