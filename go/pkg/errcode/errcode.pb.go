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
	ErrRestrictedArea                       ErrCode = 104
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
	ErrCouponAlreadyValidatedBySameTeam     ErrCode = 4061
	ErrCouponExpired                        ErrCode = 4062
	ErrCouponNotFound                       ErrCode = 4063
	ErrUserDoesNotBelongToTeam              ErrCode = 4064
	ErrInactiveAgent                        ErrCode = 4065
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
	ErrInitTracer                           ErrCode = 6016
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
	ErrAgentUpdateState                     ErrCode = 7028
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
	104:   "ErrRestrictedArea",
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
	4061:  "ErrCouponAlreadyValidatedBySameTeam",
	4062:  "ErrCouponExpired",
	4063:  "ErrCouponNotFound",
	4064:  "ErrUserDoesNotBelongToTeam",
	4065:  "ErrInactiveAgent",
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
	6016:  "ErrInitTracer",
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
	7028:  "ErrAgentUpdateState",
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
	"ErrRestrictedArea":                       104,
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
	"ErrCouponAlreadyValidatedBySameTeam":     4061,
	"ErrCouponExpired":                        4062,
	"ErrCouponNotFound":                       4063,
	"ErrUserDoesNotBelongToTeam":              4064,
	"ErrInactiveAgent":                        4065,
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
	"ErrInitTracer":                           6016,
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
	"ErrAgentUpdateState":                     7028,
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
	// 2233 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0x7c, 0x58, 0x47, 0x70, 0x1c, 0xc7,
	0xd5, 0x26, 0xab, 0xfe, 0x5f, 0x28, 0x8d, 0x4d, 0xe1, 0x69, 0x28, 0x72, 0x25, 0x52, 0xc2, 0x28,
	0xd8, 0xa4, 0xcb, 0x61, 0x71, 0x70, 0xd5, 0x56, 0xf9, 0x82, 0x2a, 0x2c, 0x16, 0x00, 0xb7, 0x44,
	0x02, 0x28, 0x2c, 0x20, 0x56, 0xf9, 0xd6, 0xd8, 0x79, 0x3b, 0xdb, 0xc6, 0x6c, 0xf7, 0xa8, 0xa7,
	0x07, 0xc1, 0x27, 0x5f, 0xe5, 0x93, 0xcf, 0xbe, 0x39, 0x5b, 0x72, 0xce, 0x16, 0x15, 0x29, 0xe6,
	0x28, 0x26, 0xe5, 0xc4, 0x60, 0x5b, 0x16, 0x95, 0x25, 0xda, 0x96, 0xb3, 0xab, 0xd3, 0xec, 0xec,
	0x02, 0xd4, 0x6d, 0xf7, 0xa5, 0x7e, 0xef, 0x7b, 0xa1, 0x5f, 0x8f, 0xb7, 0x09, 0x85, 0x68, 0xf2,
	0x10, 0xcb, 0x89, 0xe0, 0x92, 0xfb, 0x83, 0x09, 0x91, 0xed, 0x65, 0x22, 0xca, 0x96, 0xbc, 0xed,
	0x0b, 0x11, 0x95, 0xed, 0x6c, 0xa1, 0xdc, 0xe4, 0x9d, 0xe1, 0x88, 0x47, 0x7c, 0x58, 0xcb, 0x2d,
	0x64, 0x2d, 0xfd, 0x4f, 0xff, 0xd1, 0xbf, 0x8c, 0xfe, 0x67, 0xaf, 0xed, 0xf0, 0x06, 0xc6, 0x85,
	0x18, 0xe3, 0x21, 0xfa, 0x9b, 0xbc, 0x1b, 0xe7, 0x59, 0x88, 0x2d, 0xca, 0x30, 0x84, 0x0d, 0xfe,
	0x8d, 0xde, 0xff, 0xcd, 0x4d, 0xd7, 0xa6, 0xe1, 0x9b, 0xff, 0xef, 0x6f, 0xf5, 0x6e, 0x1e, 0x17,
	0x62, 0x8a, 0xcb, 0x7a, 0x27, 0x89, 0xb1, 0x83, 0x4c, 0x62, 0x08, 0x0f, 0xdc, 0xe0, 0xfb, 0xde,
	0xa6, 0x71, 0x21, 0x6a, 0x98, 0x08, 0x6c, 0x12, 0x45, 0xfb, 0xe8, 0x06, 0x1f, 0xbc, 0x4f, 0x8c,
	0x0b, 0x51, 0x67, 0x12, 0x05, 0x23, 0x31, 0xbc, 0x3e, 0xe0, 0x6f, 0xf6, 0x06, 0x35, 0x65, 0x89,
	0xc4, 0x34, 0xac, 0xb3, 0x24, 0x93, 0x80, 0x96, 0xb8, 0x87, 0xa6, 0x29, 0x65, 0x91, 0x21, 0xb6,
	0xfc, 0xad, 0x9e, 0x3f, 0x2e, 0xc4, 0x3c, 0x23, 0x99, 0x6c, 0x23, 0x93, 0xd4, 0x18, 0x8d, 0xfc,
	0x2d, 0xfa, 0xfc, 0x59, 0x4c, 0xa5, 0xa0, 0x4d, 0x89, 0xe1, 0xa8, 0x40, 0x02, 0x6d, 0x7b, 0x7c,
	0xa3, 0x31, 0x3d, 0x89, 0x72, 0xba, 0x5e, 0x1b, 0x83, 0x37, 0x06, 0xfc, 0xed, 0xde, 0x56, 0x43,
	0xb3, 0xe7, 0xcd, 0x64, 0x0b, 0x31, 0x6d, 0xde, 0x8b, 0xab, 0x70, 0x75, 0xc0, 0xbf, 0xd3, 0xdb,
	0x6e, 0x98, 0x13, 0x84, 0xc6, 0x18, 0xde, 0x8b, 0xab, 0xcd, 0x98, 0x93, 0xc5, 0x59, 0xbc, 0x3f,
	0xc3, 0x54, 0xc2, 0x9b, 0x03, 0xfe, 0xdd, 0xde, 0x1d, 0x3d, 0xea, 0x5d, 0x91, 0x34, 0xe1, 0x2c,
	0x45, 0x78, 0x6b, 0xc0, 0xbf, 0xd9, 0xfb, 0xa4, 0x91, 0xd9, 0xcd, 0x23, 0x9e, 0x49, 0x78, 0x7b,
	0xc0, 0xbf, 0xc3, 0xbb, 0xd5, 0xa9, 0x51, 0xe9, 0x74, 0xc6, 0x62, 0x8a, 0x4c, 0xc2, 0x3b, 0x03,
	0xfe, 0xad, 0xde, 0xe6, 0x1e, 0xab, 0x55, 0x24, 0x02, 0x05, 0xbc, 0x5b, 0xe0, 0x38, 0xa5, 0x71,
	0x21, 0xb8, 0x80, 0xf7, 0x06, 0x1c, 0xb6, 0xd5, 0x29, 0x2e, 0x27, 0x78, 0xc6, 0x42, 0x38, 0x3b,
	0x98, 0xd3, 0x72, 0x74, 0xcf, 0x0d, 0xfa, 0x25, 0x8d, 0x59, 0xad, 0x3a, 0x9b, 0xb1, 0x3d, 0x34,
	0x12, 0x44, 0x52, 0xce, 0x52, 0x38, 0x3f, 0xe8, 0xdf, 0xe4, 0xdd, 0x68, 0x85, 0xa9, 0x84, 0x0b,
	0x83, 0xd6, 0xed, 0x5a, 0x75, 0x8c, 0x33, 0x86, 0x4d, 0x09, 0xcf, 0x0c, 0xfa, 0x5b, 0x3c, 0xd0,
	0xa4, 0xd1, 0x4c, 0x72, 0xa3, 0x8c, 0xf0, 0x6c, 0xd7, 0xe4, 0x68, 0x18, 0x4e, 0x70, 0x81, 0x34,
	0x62, 0x0a, 0xbf, 0xe7, 0x06, 0xfd, 0x6d, 0xde, 0x16, 0x5d, 0x2c, 0x9d, 0x84, 0xa7, 0xe8, 0x00,
	0x26, 0xb2, 0x0d, 0x0f, 0x97, 0x2c, 0xb6, 0x96, 0x57, 0xa3, 0x02, 0x9b, 0x92, 0x8b, 0xd5, 0xdc,
	0xfb, 0x7d, 0x25, 0xff, 0x36, 0xef, 0x96, 0xae, 0xc4, 0x2c, 0x92, 0x70, 0x8c, 0xb3, 0x16, 0x8d,
	0xe0, 0x91, 0x92, 0x7f, 0xbb, 0x57, 0x5a, 0x63, 0xd8, 0x72, 0x1f, 0xed, 0xe3, 0xee, 0x21, 0x22,
	0x6d, 0x93, 0xd8, 0x72, 0x1f, 0x2b, 0x59, 0xec, 0x2d, 0x77, 0x4c, 0x20, 0x91, 0x38, 0x87, 0x9d,
	0x64, 0x82, 0xc6, 0x08, 0x8f, 0xf7, 0x29, 0xef, 0x15, 0xb4, 0xc0, 0x7d, 0xa2, 0x8f, 0x3b, 0x16,
	0xf3, 0xb4, 0xcb, 0x7d, 0xb2, 0xe4, 0xdf, 0xa2, 0x8b, 0xd4, 0x72, 0xab, 0x19, 0x8d, 0x43, 0xd8,
	0x5f, 0xf2, 0xb7, 0x6a, 0xd4, 0x72, 0x2a, 0x0b, 0x63, 0x84, 0x7d, 0x57, 0x37, 0xda, 0x2e, 0x29,
	0xc4, 0x57, 0x23, 0x0b, 0x70, 0xa0, 0x64, 0xe1, 0xb4, 0xf4, 0x19, 0x22, 0x52, 0x54, 0x8c, 0x83,
	0xa5, 0x5e, 0x38, 0x35, 0xc3, 0x46, 0x75, 0xa8, 0xdf, 0xb1, 0x3c, 0xaa, 0x1a, 0x15, 0x70, 0xb8,
	0x2f, 0xe6, 0xf9, 0x24, 0x2c, 0xc6, 0x7c, 0xa4, 0x2f, 0x17, 0x13, 0x5c, 0x34, 0x71, 0x16, 0x9b,
	0xda, 0x46, 0x8d, 0x2f, 0x33, 0x38, 0x5a, 0xb2, 0x75, 0xe7, 0x7c, 0xcd, 0x98, 0x39, 0x01, 0x8e,
	0xf5, 0xc5, 0x3c, 0x9b, 0xb1, 0xf9, 0x04, 0x8e, 0xbb, 0x18, 0x26, 0x51, 0xce, 0xec, 0x55, 0xf5,
	0x54, 0xa5, 0x8c, 0x88, 0x55, 0x38, 0xe1, 0x3c, 0xd1, 0xb8, 0x1a, 0x96, 0xf2, 0x61, 0x17, 0x92,
	0x10, 0x05, 0x9c, 0x74, 0x7a, 0x7d, 0x6c, 0x38, 0x55, 0xf2, 0x03, 0x6f, 0x9b, 0xea, 0x7f, 0x93,
	0x4c, 0xc3, 0x32, 0xc1, 0x6b, 0x81, 0xd3, 0x25, 0xff, 0x1e, 0x6f, 0xa8, 0x57, 0xb3, 0xcb, 0xb6,
	0xe6, 0x9f, 0x5e, 0xe7, 0xf4, 0x82, 0x8d, 0x33, 0x25, 0xff, 0x2e, 0xef, 0xf6, 0x3e, 0xb6, 0xce,
	0x30, 0x31, 0x24, 0x01, 0x67, 0xbb, 0x48, 0x26, 0xab, 0x46, 0x62, 0x8e, 0x8f, 0x71, 0x26, 0x09,
	0x65, 0x28, 0xe0, 0x5c, 0x1f, 0x92, 0x93, 0x28, 0x73, 0x66, 0x5a, 0x67, 0x2d, 0x0e, 0xe7, 0x4b,
	0x76, 0xe0, 0xd8, 0x41, 0x36, 0xb3, 0x4c, 0x73, 0x27, 0xe0, 0x82, 0xcb, 0xe2, 0x24, 0xca, 0xf9,
	0x14, 0x45, 0xbd, 0x36, 0x21, 0x78, 0x47, 0x59, 0xc0, 0x15, 0x09, 0xdf, 0x0a, 0xec, 0xb0, 0xb1,
	0xaa, 0x63, 0x6d, 0x12, 0xc7, 0xc8, 0x22, 0xbc, 0x4f, 0x15, 0xbf, 0x6e, 0x63, 0xf8, 0x76, 0x60,
	0x5b, 0xd4, 0xb6, 0x44, 0x03, 0x49, 0xca, 0x19, 0x7c, 0x27, 0xb0, 0xb8, 0xce, 0x21, 0xe9, 0xa8,
	0xa9, 0xcc, 0x2c, 0xe3, 0xbb, 0x81, 0x75, 0x58, 0x79, 0xea, 0xec, 0x35, 0xb2, 0x85, 0xb4, 0x29,
	0x68, 0xa2, 0x2d, 0x7e, 0xaf, 0x6b, 0x91, 0xca, 0x06, 0xe3, 0xcb, 0xad, 0x98, 0x2c, 0x22, 0x7c,
	0x3f, 0xb0, 0x78, 0x9b, 0x5a, 0x5a, 0x5f, 0xf7, 0x07, 0x81, 0x05, 0xd4, 0x14, 0xcb, 0x7a, 0x0e,
	0xff, 0x30, 0xf0, 0x87, 0xbc, 0xdb, 0xfa, 0x1c, 0x28, 0xf0, 0x1f, 0x0c, 0xfc, 0xcd, 0xde, 0x4d,
	0xdd, 0x80, 0x54, 0x00, 0xf0, 0x90, 0x43, 0x22, 0xd7, 0x18, 0x8d, 0x05, 0x92, 0x70, 0xd5, 0x9e,
	0xbe, 0x80, 0x21, 0xfc, 0xc8, 0x39, 0xd8, 0x77, 0x76, 0x8f, 0x83, 0x3f, 0x0e, 0xec, 0x8c, 0x99,
	0xa0, 0x2c, 0x9c, 0x16, 0x11, 0x61, 0xf4, 0xab, 0x76, 0x1e, 0xfe, 0x24, 0xf0, 0x3f, 0xe5, 0x05,
	0xc6, 0x31, 0x03, 0x96, 0xca, 0x85, 0xf9, 0x95, 0x1b, 0x83, 0x9f, 0x06, 0xb6, 0x1e, 0x6c, 0xc6,
	0x94, 0x7b, 0x5d, 0x39, 0xf8, 0x99, 0xc3, 0xbd, 0x27, 0x1d, 0xf5, 0x1a, 0xfc, 0xdc, 0x85, 0xad,
	0x94, 0x76, 0x91, 0x74, 0x8a, 0x6b, 0x4d, 0x2e, 0xac, 0xe2, 0x2f, 0x02, 0x5b, 0x26, 0xf9, 0xe9,
	0xf9, 0x99, 0x29, 0xfc, 0x32, 0xb0, 0xa3, 0x39, 0x67, 0xc2, 0xaf, 0x02, 0xdb, 0x86, 0xe6, 0x7f,
	0x0d, 0x19, 0xc5, 0x10, 0x7e, 0x1d, 0xd8, 0xae, 0xb1, 0xf0, 0xec, 0x22, 0x69, 0xef, 0x31, 0xbf,
	0x71, 0x6a, 0xb3, 0x98, 0xa2, 0x58, 0xc2, 0x70, 0x8a, 0x74, 0x10, 0x7e, 0x9b, 0x43, 0xd7, 0xc6,
	0xe6, 0x62, 0x11, 0x96, 0x79, 0x46, 0xef, 0xcf, 0x50, 0x0b, 0xfd, 0x2e, 0x70, 0xd3, 0x48, 0xe3,
	0x5b, 0x94, 0x82, 0xdf, 0x07, 0xfe, 0xe7, 0xbc, 0x1d, 0xe3, 0x42, 0x14, 0xa9, 0xd7, 0xf3, 0xe1,
	0xe1, 0xa0, 0x3b, 0x2b, 0x7a, 0xac, 0xec, 0x73, 0x27, 0xac, 0xc5, 0x00, 0x1e, 0x71, 0x27, 0x8c,
	0x11, 0xc6, 0xb8, 0x74, 0xe3, 0xce, 0xd8, 0xe5, 0x31, 0xef, 0x31, 0xf4, 0xa8, 0x4b, 0x92, 0x02,
	0x5b, 0x57, 0x7f, 0x0f, 0xfb, 0xb1, 0xc0, 0x5e, 0x93, 0x5d, 0x2b, 0xf0, 0x78, 0xe0, 0x0f, 0x7a,
	0x9e, 0x39, 0x5b, 0x13, 0x9e, 0x08, 0xec, 0x9e, 0x62, 0x09, 0x29, 0x3c, 0x59, 0x10, 0x51, 0x86,
	0x61, 0xbf, 0xb3, 0x63, 0x5a, 0x42, 0xd3, 0x9e, 0xea, 0xa5, 0x69, 0x53, 0x07, 0x5c, 0x5c, 0x86,
	0xd6, 0xe3, 0xcb, 0x41, 0x57, 0x90, 0x53, 0xb8, 0xac, 0x0c, 0xe8, 0xfe, 0x8f, 0x09, 0xed, 0xa4,
	0x70, 0xc8, 0xe5, 0x4a, 0xe1, 0x34, 0x9a, 0xc9, 0xb6, 0x3e, 0xe0, 0x70, 0xe0, 0x7f, 0xde, 0xdb,
	0xa9, 0x2e, 0x5f, 0xda, 0x6a, 0xa1, 0x40, 0xa6, 0x7d, 0xa9, 0xa2, 0x5c, 0x46, 0x64, 0x73, 0x7c,
	0x11, 0xd9, 0x28, 0x0b, 0x6b, 0x44, 0x92, 0x05, 0x92, 0x22, 0x1c, 0x71, 0x58, 0xef, 0xe6, 0x24,
	0x54, 0x82, 0x06, 0xd7, 0x14, 0x8e, 0x06, 0xbd, 0x93, 0xa7, 0xb7, 0x17, 0x8e, 0xb9, 0x28, 0xf2,
	0x4c, 0xa4, 0x70, 0x3c, 0xb0, 0x57, 0x82, 0xd5, 0xa8, 0xaa, 0xe6, 0xfb, 0x8a, 0x5a, 0x13, 0x4e,
	0xb8, 0xaa, 0x1b, 0xef, 0x10, 0x1a, 0x8f, 0x86, 0xa1, 0xc0, 0x34, 0x9d, 0xe2, 0xf2, 0x3e, 0x14,
	0xb4, 0xa5, 0xca, 0xf2, 0x64, 0x41, 0xb5, 0x86, 0x2d, 0x92, 0xc5, 0xae, 0x8c, 0x4f, 0x05, 0xdd,
	0x8b, 0xaa, 0x43, 0x4d, 0x47, 0x09, 0xc2, 0x52, 0xd2, 0xd4, 0xe8, 0x9c, 0xee, 0x45, 0x6e, 0xb4,
	0x29, 0xe9, 0x12, 0x5a, 0xd5, 0xa7, 0x5d, 0x47, 0xb9, 0xe9, 0x68, 0xa6, 0xe6, 0x1e, 0x94, 0x24,
	0x24, 0x92, 0xc0, 0x19, 0x17, 0xfa, 0x14, 0xd7, 0xb0, 0xcc, 0x08, 0xbe, 0x44, 0x43, 0x0c, 0xe1,
	0x6c, 0xa1, 0xcc, 0x34, 0x67, 0x2f, 0x95, 0x6d, 0x8b, 0xf9, 0x39, 0xe7, 0xa9, 0x55, 0xaa, 0x33,
	0x37, 0x8c, 0xcf, 0x17, 0x1b, 0xd4, 0x04, 0xae, 0x72, 0xa5, 0xa5, 0xe0, 0x42, 0x61, 0x2a, 0x14,
	0x98, 0x4e, 0xf7, 0x19, 0x37, 0x16, 0x27, 0x51, 0x16, 0x63, 0xd8, 0x83, 0x9d, 0x05, 0x14, 0x69,
	0x9b, 0x26, 0xf0, 0x6c, 0xc1, 0xbc, 0xb6, 0x59, 0xd4, 0x7f, 0xce, 0x85, 0xda, 0x3f, 0xfe, 0xf4,
	0x65, 0x15, 0xc2, 0xf3, 0x85, 0x5a, 0x1d, 0x8d, 0xd4, 0x46, 0xf9, 0x82, 0x9b, 0x18, 0x0d, 0xb2,
	0x84, 0x86, 0xf4, 0xa2, 0x33, 0xb2, 0x9b, 0xa6, 0xdd, 0xc9, 0x5b, 0x67, 0xa9, 0x24, 0xac, 0x89,
	0x29, 0xbc, 0xe4, 0xca, 0xad, 0x7b, 0x48, 0x18, 0xc2, 0xcb, 0x81, 0xff, 0x19, 0xef, 0x1e, 0x9d,
	0xa0, 0x2c, 0xc9, 0x7b, 0xda, 0xce, 0x6b, 0x0c, 0xab, 0xab, 0x0d, 0xd2, 0x31, 0x55, 0xfe, 0x8a,
	0xbb, 0x37, 0x8c, 0xe4, 0xf8, 0x4a, 0x42, 0x05, 0x86, 0xf0, 0x6a, 0x90, 0x6f, 0x3d, 0x8a, 0x9c,
	0x6f, 0x7b, 0xaf, 0xb9, 0xa2, 0x51, 0x39, 0xaf, 0x71, 0x54, 0x05, 0x53, 0xc5, 0x98, 0xb3, 0x68,
	0x4e, 0x8f, 0x46, 0xb8, 0xd8, 0xbd, 0x87, 0x88, 0xc6, 0xcc, 0x84, 0x71, 0xa9, 0x3b, 0xf8, 0xc4,
	0x12, 0xea, 0x60, 0x90, 0xc1, 0x03, 0x3b, 0xdd, 0x82, 0xad, 0xa9, 0xb3, 0x18, 0x29, 0xba, 0x98,
	0x24, 0x12, 0x97, 0xc9, 0x2a, 0x7c, 0x7d, 0xa7, 0xad, 0x67, 0x75, 0xa7, 0xed, 0xe6, 0x51, 0x84,
	0x02, 0xde, 0x2f, 0x3b, 0x43, 0x92, 0x08, 0xa9, 0xf4, 0x68, 0x13, 0xe1, 0x83, 0x72, 0x41, 0xd2,
	0x18, 0x83, 0x0f, 0xcb, 0x6e, 0x60, 0x09, 0x9e, 0x25, 0x73, 0x28, 0x3a, 0x94, 0xe9, 0x67, 0xc7,
	0xb5, 0x72, 0x21, 0xed, 0x8d, 0x69, 0xb3, 0xcd, 0xab, 0xc4, 0x4d, 0xc4, 0x24, 0x4a, 0xe1, 0x2f,
	0xee, 0x84, 0x5a, 0xd6, 0x49, 0xf2, 0x96, 0xfc, 0x6b, 0xb9, 0x3b, 0xcc, 0xd5, 0xea, 0xdd, 0xe2,
	0xf0, 0xb7, 0x72, 0xb7, 0xd3, 0x1b, 0x8d, 0xe9, 0xbd, 0x6d, 0x4e, 0x3a, 0x14, 0x3e, 0xea, 0xa5,
	0xda, 0xa7, 0xc4, 0xdf, 0x7b, 0xa9, 0xb6, 0x6e, 0xff, 0x51, 0xb6, 0x75, 0xab, 0xdc, 0xae, 0xf1,
	0xe6, 0x22, 0x0a, 0xfb, 0xb6, 0xf8, 0x67, 0xd9, 0xae, 0xf9, 0x9a, 0x53, 0x85, 0x7f, 0x95, 0xed,
	0xfd, 0x6a, 0x56, 0x90, 0x4c, 0x60, 0xad, 0x0a, 0xff, 0x2e, 0x17, 0xef, 0x7c, 0x17, 0x09, 0xfc,
	0xa7, 0x9c, 0xdf, 0xc5, 0x34, 0x47, 0xe8, 0xbf, 0x45, 0x84, 0xe6, 0x04, 0x69, 0xa2, 0x80, 0xaf,
	0x0d, 0xdb, 0xda, 0xd2, 0x39, 0x5a, 0xbb, 0x04, 0xbd, 0x50, 0xb1, 0xd5, 0xad, 0x2f, 0x98, 0xa9,
	0x88, 0xb2, 0x95, 0xee, 0x0e, 0xf5, 0x62, 0xc5, 0x8e, 0xa2, 0x59, 0xec, 0xf0, 0x25, 0xec, 0xe3,
	0xbe, 0xe4, 0x54, 0xf5, 0x72, 0xdd, 0xc7, 0x7c, 0xd9, 0x31, 0x75, 0x0e, 0xfb, 0x98, 0xaf, 0x54,
	0x6c, 0xda, 0xd4, 0xde, 0x4c, 0x59, 0xa4, 0xd6, 0xdf, 0x58, 0xad, 0xb0, 0xaf, 0x56, 0x8a, 0x5b,
	0xe1, 0x9a, 0xa5, 0xf1, 0xb5, 0x4a, 0x71, 0x27, 0x2d, 0xac, 0x8b, 0x17, 0x2b, 0xae, 0x94, 0x7b,
	0x77, 0xc4, 0x4b, 0x15, 0xb7, 0x9d, 0xf0, 0x64, 0xd5, 0x39, 0xd1, 0xa2, 0x51, 0x71, 0x51, 0xbc,
	0x5c, 0xb1, 0x23, 0x40, 0xf3, 0xa7, 0x70, 0xd9, 0x88, 0x68, 0x3c, 0xcc, 0x53, 0x13, 0xae, 0x54,
	0xfc, 0x1d, 0xde, 0x5d, 0x4e, 0xa4, 0x81, 0x2c, 0x54, 0x43, 0x91, 0xb0, 0xb0, 0x57, 0x1a, 0xfe,
	0x50, 0xf1, 0x77, 0x7a, 0x77, 0x7f, 0x9c, 0x9c, 0x01, 0x12, 0xfe, 0x58, 0xf1, 0x3f, 0xed, 0xdd,
	0xb9, 0x8e, 0xa0, 0x93, 0x4a, 0x62, 0xd2, 0x44, 0xf8, 0x53, 0xc5, 0x2e, 0x3e, 0xfd, 0x62, 0xb3,
	0x18, 0xf3, 0xfc, 0x09, 0xf6, 0xba, 0x83, 0xda, 0x05, 0xa8, 0x5e, 0x88, 0x53, 0x28, 0x97, 0xb9,
	0x58, 0x84, 0x3f, 0x57, 0x6c, 0x33, 0xe7, 0x01, 0xf7, 0x09, 0xbc, 0xe1, 0xa0, 0x9b, 0x22, 0x72,
	0x86, 0x0b, 0x39, 0x9d, 0x20, 0xa3, 0x2c, 0x82, 0xab, 0x15, 0x5b, 0xb7, 0x3d, 0xd9, 0x55, 0xe7,
	0xbd, 0xe9, 0xb2, 0x30, 0xbe, 0x82, 0xcd, 0xcc, 0x3c, 0x5e, 0x74, 0xf6, 0xde, 0x72, 0x67, 0x69,
	0xf4, 0xab, 0xab, 0x12, 0xd3, 0x39, 0xbe, 0x8b, 0xa4, 0x6d, 0x6d, 0x02, 0x05, 0xbc, 0x5d, 0xb1,
	0x2b, 0xae, 0x7a, 0x60, 0x69, 0xbe, 0x6a, 0xc9, 0xa2, 0xc4, 0x3b, 0x95, 0xfc, 0x06, 0x60, 0xa8,
	0x9e, 0xb4, 0x33, 0x02, 0x5b, 0x74, 0x45, 0x89, 0xc0, 0xbb, 0xae, 0x38, 0xc6, 0x62, 0x24, 0x6c,
	0xc6, 0x7c, 0x3b, 0xe9, 0x4e, 0xc9, 0xf7, 0x8a, 0x45, 0x85, 0xdd, 0xf7, 0x04, 0xbc, 0x5f, 0xb1,
	0x17, 0xf9, 0x7c, 0xd2, 0xa7, 0x04, 0x1f, 0x54, 0x6c, 0x1b, 0x99, 0x5b, 0x4c, 0x47, 0x09, 0x1f,
	0xba, 0xc8, 0x75, 0xcb, 0x18, 0x4e, 0x43, 0xaa, 0x00, 0xaf, 0x55, 0x6c, 0x33, 0x99, 0x3e, 0x1e,
	0x9d, 0xa9, 0xe7, 0xa8, 0xaa, 0x69, 0x07, 0xfb, 0x47, 0x6c, 0x7c, 0x6b, 0xf9, 0x36, 0xf1, 0x4f,
	0x8d, 0xd8, 0x8e, 0xca, 0x25, 0xea, 0x1d, 0x12, 0xa1, 0xe5, 0x1e, 0xb8, 0xbe, 0xbe, 0x7d, 0xe3,
	0x1d, 0x1c, 0xb1, 0x15, 0xb1, 0x56, 0x42, 0x65, 0xc3, 0x4a, 0x1d, 0xfa, 0x78, 0xa9, 0x51, 0x29,
	0x49, 0xb3, 0x0d, 0x87, 0x47, 0xec, 0x6e, 0xb9, 0xbe, 0x94, 0x6e, 0x5c, 0x38, 0x32, 0x62, 0x2b,
	0x75, 0x7d, 0xa1, 0x3a, 0x4b, 0x13, 0xb5, 0x68, 0x1c, 0x1d, 0xb1, 0x79, 0xeb, 0x8d, 0x6b, 0x26,
	0x8b, 0x63, 0x38, 0x36, 0x62, 0xf3, 0xd6, 0xcb, 0x73, 0xaa, 0xc7, 0xd7, 0x40, 0x62, 0x4b, 0x53,
	0x43, 0x7a, 0x62, 0xa4, 0x1f, 0x72, 0xcb, 0xb5, 0xa1, 0x9e, 0xbc, 0x1e, 0xdf, 0x42, 0x7a, 0x6a,
	0xc4, 0x26, 0x3f, 0xe7, 0x8f, 0xaf, 0xa8, 0xca, 0x08, 0x11, 0x4e, 0x3b, 0x96, 0xad, 0xe3, 0x69,
	0xa6, 0x8a, 0x66, 0x17, 0xe7, 0x8b, 0xf0, 0xe0, 0x84, 0x2d, 0x01, 0x63, 0xa5, 0x50, 0x4c, 0x0f,
	0x4d, 0x54, 0xbf, 0x74, 0xe6, 0xd2, 0xd0, 0x86, 0xa3, 0x97, 0x87, 0x36, 0x9e, 0xb9, 0x3c, 0xb4,
	0xf1, 0xe2, 0xe5, 0xa1, 0x8d, 0xdf, 0xb8, 0x32, 0xb4, 0xe1, 0xcc, 0x95, 0xa1, 0x0d, 0xcf, 0x5f,
	0x19, 0xda, 0xf0, 0xe5, 0xed, 0xee, 0xa3, 0x5e, 0x4c, 0x58, 0x38, 0x1c, 0xf1, 0xe1, 0x64, 0x31,
	0x1a, 0xb6, 0x1f, 0xf8, 0x16, 0x6e, 0xd0, 0x1f, 0xee, 0xbe, 0xf8, 0xbf, 0x00, 0x00, 0x00, 0xff,
	0xff, 0xca, 0xf3, 0x5b, 0x52, 0x09, 0x14, 0x00, 0x00,
}
