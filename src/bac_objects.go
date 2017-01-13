package bacnet

type Object interface {

}

type Accumulator struct {
	Default
	PresentValue                 string `json:"presentValue"`
	StatusFlags                  string `json:"statusFlags"`
	EventState                   string `json:"eventState"`
	OutOfService                 string `json:"outOfService"`
	Scale                        string `json:"scale"`
	Units                        string `json:"units"`
	MaxPresValue                 string `json:"maxPresValue"`
	Description                  string `json:"description"`
	DeviceType                   string `json:"deviceType"`
	Reliability                  string `json:"reliability"`
	Prescale                     string `json:"prescale"`
	ValueChangeTime              string `json:"valueChangeTime"`
	ValueBeforeChange            string `json:"valueBeforeChange"`
	ValueSet                     string `json:"valueSet"`
	LoggingRecord                string `json:"loggingRecord"`
	LoggingObject                string `json:"loggingObject"`
	PulseRate                    string `json:"pulseRate"`
	HighLimit                    string `json:"highLimit"`
	LowLimit                     string `json:"lowLimit"`
	LimitMonitoringInterval      string `json:"limitMonitoringInterval"`
	NotificationClass            string `json:"notificationClass"`
	TimeDelay                    string `json:"timeDelay"`
	LimitEnable                  string `json:"limitEnable"`
	EventEnable                  string `json:"eventEnable"`
	AckedTransitions             string `json:"ackedTransitions"`
	NotifyType                   string `json:"notifyType"`
	EventTimeStamps              string `json:"eventTimeStamps"`
	EventMessageTexts            string `json:"eventMessageTexts"`
	ReliabilityEvaluationInhibit string `json:"reliabilityEvaluationInhibit"`
	ProfileName                  string `json:"profileName"`
}

type AnalogInput struct {
	Default
	PresentValue                 string `json:"presentValue"`
	StatusFlags                  string `json:"statusFlags"`
	EventState                   string `json:"eventState"`
	OutOfService                 string `json:"outOfService"`
	Units                        string `json:"units"`
	Description                  string `json:"description"`
	DeviceType                   string `json:"deviceType"`
	Reliability                  string `json:"reliability"`
	UpdateInterval               string `json:"updateInterval"`
	MinPresValue                 string `json:"minPresValue"`
	MaxPresValue                 string `json:"maxPresValue"`
	Resolution                   string `json:"resolution"`
	CovIncrement                 string `json:"covIncrement"`
	TimeDelay                    string `json:"timeDelay"`
	NotificationClass            string `json:"notificationClass"`
	HighLimit                    string `json:"highLimit"`
	LowLimit                     string `json:"lowLimit"`
	Deadband                     string `json:"deadband"`
	LimitEnable                  string `json:"limitEnable"`
	EventEnable                  string `json:"eventEnable"`
	AckedTransitions             string `json:"ackedTransitions"`
	NotifyType                   string `json:"notifyType"`
	EventTimeStamps              string `json:"eventTimeStamps"`
	EventMessageTexts            string `json:"eventMessageTexts"`
	ReliabilityEvaluationInhibit string `json:"reliabilityEvaluationInhibit"`
	ProfileName                  string `json:"profileName"`
}

type AnalogOutput struct {
	Default
	PresentValue                 string `json:"presentValue"`
	StatusFlags                  string `json:"statusFlags"`
	EventState                   string `json:"eventState"`
	OutOfService                 string `json:"outOfService"`
	Units                        string `json:"units"`
	PriorityArray                string `json:"priorityArray"`
	RelinquishDefault            string `json:"relinquishDefault"`
	Description                  string `json:"description"`
	DeviceType                   string `json:"deviceType"`
	Reliability                  string `json:"reliability"`
	MinPresValue                 string `json:"minPresValue"`
	MaxPresValue                 string `json:"maxPresValue"`
	Resolution                   string `json:"resolution"`
	CovIncrement                 string `json:"covIncrement"`
	TimeDelay                    string `json:"timeDelay"`
	NotificationClass            string `json:"notificationClass"`
	HighLimit                    string `json:"highLimit"`
	LowLimit                     string `json:"lowLimit"`
	Deadband                     string `json:"deadband"`
	LimitEnable                  string `json:"limitEnable"`
	EventEnable                  string `json:"eventEnable"`
	AckedTransitions             string `json:"ackedTransitions"`
	NotifyType                   string `json:"notifyType"`
	EventTimeStamps              string `json:"eventTimeStamps"`
	EventMessageTexts            string `json:"eventMessageTexts"`
	ReliabilityEvaluationInhibit string `json:"reliabilityEvaluationInhibit"`
	ProfileName                  string `json:"profileName"`
}

type AnalogValue struct {
	Default
	PresentValue                 string `json:"presentValue"`
	StatusFlags                  string `json:"statusFlags"`
	EventState                   string `json:"eventState"`
	OutOfService                 string `json:"outOfService"`
	Units                        string `json:"units"`
	Description                  string `json:"description"`
	Reliability                  string `json:"reliability"`
	PriorityArray                string `json:"priorityArray"`
	RelinquishDefault            string `json:"relinquishDefault"`
	CovIncrement                 string `json:"covIncrement"`
	TimeDelay                    string `json:"timeDelay"`
	NotificationClass            string `json:"notificationClass"`
	HighLimit                    string `json:"highLimit"`
	LowLimit                     string `json:"lowLimit"`
	Deadband                     string `json:"deadband"`
	LimitEnable                  string `json:"limitEnable"`
	EventEnable                  string `json:"eventEnable"`
	AckedTransitions             string `json:"ackedTransitions"`
	NotifyType                   string `json:"notifyType"`
	EventTimeStamps              string `json:"eventTimeStamps"`
	EventMessageTexts            string `json:"eventMessageTexts"`
	ReliabilityEvaluationInhibit string `json:"reliabilityEvaluationInhibit"`
	ProfileName                  string `json:"profileName"`
}

type Averaging struct {
	Default
	MinimumValue            string `json:"minimumValue"`
	AverageValue            string `json:"averageValue"`
	MaximumValue            string `json:"maximumValue"`
	StatusFlags             string `json:"statusFlags"`
	EventState              string `json:"eventState"`
	OutOfService            string `json:"outOfService"`
	Units                   string `json:"units"`
	ProfileName             string `json:"profileName"`
	MinimumValueTimestamp   string `json:"minimumValueTimestamp"`
	VarianceValue           string `json:"varianceValue"`
	MaximumValueTimestamp   string `json:"maximumValueTimestamp"`
	Description             string `json:"description"`
	AttemptedSamples        string `json:"attemptedSamples"`
	ValidSamples            string `json:"validSamples"`
	ObjectPropertyReference string `json:"objectPropertyReference"`
	WindowInterval          string `json:"windowInterval"`
	WindowSamples           string `json:"windowSamples"`
}

type BinaryInput struct {
	Default
	PresentValue                 string `json:"presentValue"`
	StatusFlags                  string `json:"statusFlags"`
	EventState                   string `json:"eventState"`
	OutOfService                 string `json:"outOfService"`
	Polarity                     string `json:"polarity"`
	Description                  string `json:"description"`
	DeviceType                   string `json:"deviceType"`
	Reliability                  string `json:"reliability"`
	InactiveText                 string `json:"inactiveText"`
	ActiveText                   string `json:"activeText"`
	ChangeOfStateTime            string `json:"changeOfStateTime"`
	ChangeOfStateCount           string `json:"changeOfStateCount"`
	TimeOfStateCountReset        string `json:"timeOfStateCountReset"`
	ElapsedActiveTime            string `json:"elapsedActiveTime"`
	TimeOfActiveTimeReset        string `json:"timeOfActiveTimeReset"`
	TimeDelay                    string `json:"timeDelay"`
	NotificationClass            string `json:"notificationClass"`
	AlarmValue                   string `json:"alarmValue"`
	EventEnable                  string `json:"eventEnable"`
	AckedTransitions             string `json:"ackedTransitions"`
	NotifyType                   string `json:"notifyType"`
	EventTimeStamps              string `json:"eventTimeStamps"`
	EventMessageTexts            string `json:"eventMessageTexts"`
	ReliabilityEvaluationInhibit string `json:"reliabilityEvaluationInhibit"`
	ProfileName                  string `json:"profileName"`
}

type BinaryOutput struct {
	Default
	PresentValue                 string `json:"presentValue"`
	StatusFlags                  string `json:"statusFlags"`
	EventState                   string `json:"eventState"`
	OutOfService                 string `json:"outOfService"`
	Polarity                     string `json:"polarity"`
	PriorityArray                string `json:"priorityArray"`
	RelinquishDefault            string `json:"relinquishDefault"`
	Description                  string `json:"description"`
	DeviceType                   string `json:"deviceType"`
	Reliability                  string `json:"reliability"`
	InactiveText                 string `json:"inactiveText"`
	ActiveText                   string `json:"activeText"`
	ChangeOfStateTime            string `json:"changeOfStateTime"`
	ChangeOfStateCount           string `json:"changeOfStateCount"`
	TimeOfStateCountReset        string `json:"timeOfStateCountReset"`
	ElapsedActiveTime            string `json:"elapsedActiveTime"`
	TimeOfActiveTimeReset        string `json:"timeOfActiveTimeReset"`
	MinimumOffTime               string `json:"minimumOffTime"`
	MinimumOnTime                string `json:"minimumOnTime"`
	TimeDelay                    string `json:"timeDelay"`
	NotificationClass            string `json:"notificationClass"`
	FeedbackValue                string `json:"feedbackValue"`
	EventEnable                  string `json:"eventEnable"`
	AckedTransitions             string `json:"ackedTransitions"`
	NotifyType                   string `json:"notifyType"`
	EventTimeStamps              string `json:"eventTimeStamps"`
	EventMessageTexts            string `json:"eventMessageTexts"`
	ReliabilityEvaluationInhibit string `json:"reliabilityEvaluationInhibit"`
	ProfileName                  string `json:"profileName"`
}

type BinaryValue struct {
	Default
	PresentValue                 string `json:"presentValue"`
	StatusFlags                  string `json:"statusFlags"`
	EventState                   string `json:"eventState"`
	OutOfService                 string `json:"outOfService"`
	Description                  string `json:"description"`
	Reliability                  string `json:"reliability"`
	InactiveText                 string `json:"inactiveText"`
	ActiveText                   string `json:"activeText"`
	ChangeOfStateTime            string `json:"changeOfStateTime"`
	ChangeOfStateCount           string `json:"changeOfStateCount"`
	TimeOfStateCountReset        string `json:"timeOfStateCountReset"`
	ElapsedActiveTime            string `json:"elapsedActiveTime"`
	TimeOfActiveTimeReset        string `json:"timeOfActiveTimeReset"`
	MinimumOffTime               string `json:"minimumOffTime"`
	MinimumOnTime                string `json:"minimumOnTime"`
	PriorityArray                string `json:"priorityArray"`
	RelinquishDefault            string `json:"relinquishDefault"`
	TimeDelay                    string `json:"timeDelay"`
	NotificationClass            string `json:"notificationClass"`
	AlarmValue                   string `json:"alarmValue"`
	EventEnable                  string `json:"eventEnable"`
	AckedTransitions             string `json:"ackedTransitions"`
	NotifyType                   string `json:"notifyType"`
	EventTimeStamps              string `json:"eventTimeStamps"`
	EventMessageTexts            string `json:"eventMessageTexts"`
	ReliabilityEvaluationInhibit string `json:"reliabilityEvaluationInhibit"`
	ProfileName                  string `json:"profileName"`
}

type Calendar struct {
	Default
	PresentValue string `json:"presentValue"`
	DateList     string `json:"dateList"`
	Description  string `json:"description"`
	ProfileName  string `json:"profileName"`
}

type Channel struct {
	Default
	PresentValue                   string `json:"presentValue"`
	LastPriority                   string `json:"lastPriority"`
	WriteStatus                    string `json:"writeStatus"`
	StatusFlags                    string `json:"statusFlags"`
	OutOfService                   string `json:"outOfService"`
	ListOfObjectPropertyReferences string `json:"listOfObjectPropertyReferences"`
	ChannelNumber                  string `json:"channelNumber"`
	ControlGroups                  string `json:"controlGroups"`
	Description                    string `json:"description"`
	Reliability                    string `json:"reliability"`
	ExecutionDelay                 string `json:"executionDelay"`
	AllowGroupDelayInhibit         string `json:"allowGroupDelayInhibit"`
	EventDetectionEnable           string `json:"eventDetectionEnable"`
	NotificationClass              string `json:"notificationClass"`
	EventEnable                    string `json:"eventEnable"`
	EventState                     string `json:"eventState"`
	AckedTransitions               string `json:"ackedTransitions"`
	NotifyType                     string `json:"notifyType"`
	EventTimeStamps                string `json:"eventTimeStamps"`
	EventMessageTexts              string `json:"eventMessageTexts"`
	EventMessageTextsConfig        string `json:"eventMessageTextsConfig"`
	ReliabilityEvaluationInhibit   string `json:"reliabilityEvaluationInhibit"`
	ProfileName                    string `json:"profileName"`
}

type CharacterStringValue struct {
	Default
	PresentValue                 string `json:"presentValue"`
	StatusFlags                  string `json:"statusFlags"`
	Description                  string `json:"description"`
	EventState                   string `json:"eventState"`
	Reliability                  string `json:"reliability"`
	OutOfService                 string `json:"outOfService"`
	PriorityArray                string `json:"priorityArray"`
	RelinquishDefault            string `json:"relinquishDefault"`
	TimeDelay                    string `json:"timeDelay"`
	NotificationClass            string `json:"notificationClass"`
	AlarmValues                  string `json:"alarmValues"`
	FaultValues                  string `json:"faultValues"`
	EventEnable                  string `json:"eventEnable"`
	AckedTransitions             string `json:"ackedTransitions"`
	NotifyType                   string `json:"notifyType"`
	EventTimeStamps              string `json:"eventTimeStamps"`
	EventMessageTexts            string `json:"eventMessageTexts"`
	ReliabilityEvaluationInhibit string `json:"reliabilityEvaluationInhibit"`
	ProfileName                  string `json:"profileName"`
}

type Command struct {
	Default
	PresentValue        string `json:"presentValue"`
	InProcess           string `json:"inProcess"`
	AllWritesSuccessful string `json:"allWritesSuccessful"`
	Action              string `json:"action"`
	Description         string `json:"description"`
	ActionText          string `json:"actionText"`
	ProfileName         string `json:"profileName"`
}

type Default struct {
	ObjectIdentifier uint32 `json:"objectIdentifier"`
	ObjectName       string `json:"objectName"`
	ObjectType       string `json:"objectType"`
}

type Device struct {
	Default
	SystemStatus                     string `json:"systemStatus"`
	VendorName                       string `json:"vendorName"`
	VendorIdentifier                 string `json:"vendorIdentifier"`
	ModelName                        string `json:"modelName"`
	FirmwareRevision                 string `json:"firmwareRevision"`
	ApplicationSoftwareVersion       string `json:"applicationSoftwareVersion"`
	ProtocolVersion                  string `json:"protocolVersion"`
	ProtocolRevision                 string `json:"protocolRevision"`
	ProtocolServicesSupported        []string `json:"protocolServicesSupported"`
	ProtocolObjectTypesSupported     string `json:"protocolObjectTypesSupported"`
	ObjectList                       map[string]*Object `json:"objectList"`
	MaxAPDULengthAccepted            string `json:"maxAPDULengthAccepted"`
	SegmentationSupported            bool `json:"segmentationSupported"`
	APDUTimeout                      int32 `json:"APDUTimeout"`
	NumberOfAPDURetries              int `json:"numberOfAPDURetries"`
	DeviceAddressBinding             string `json:"deviceAddressBinding"`
	DatabaseRevision                 string `json:"databaseRevision"`
	Location                         string `json:"location"`
	Description                      string `json:"description"`
	StructuredObjectList             string `json:"structuredObjectList"`
	MaxSegmentsAccepted              string `json:"maxSegmentsAccepted"`
	VTClassesSupported               string `json:"VTClassesSupported"`
	ActiveVTSessions                 string `json:"activeVTSessions"`
	LocalTime                        string `json:"localTime"`
	LocalDate                        string `json:"localDate"`
	UTCOffset                        int `json:"UTCOffset"`
	DaylightSavingsStatus            bool `json:"daylightSavingsStatus"`
	APDUSegmentTimeout               int32 `json:"APDUSegmentTimeout"`
	TimeSynchronizationRecipients    string `json:"timeSynchronizationRecipients"`
	MaxMaster                        string `json:"maxMaster"`
	MaxInfoFrames                    string `json:"maxInfoFrames"`
	ConfigurationFiles               string `json:"configurationFiles"`
	LastRestoreTime                  string `json:"lastRestoreTime"`
	BackupFailureTimeout             string `json:"backupFailureTimeout"`
	BackupPreparationTime            string `json:"backupPreparationTime"`
	RestorePreparationTime           string `json:"restorePreparationTime"`
	RestoreCompletionTime            string `json:"restoreCompletionTime"`
	BackupAndRestoreState            string `json:"backupAndRestoreState"`
	ActiveCovSubscriptions           string `json:"activeCovSubscriptions"`
	SlaveProxyEnable                 string `json:"slaveProxyEnable"`
	ManualSlaveAddressBinding        string `json:"manualSlaveAddressBinding"`
	AutoSlaveDiscovery               string `json:"autoSlaveDiscovery"`
	SlaveAddressBinding              string `json:"slaveAddressBinding"`
	LastRestartReason                string `json:"lastRestartReason"`
	TimeOfDeviceRestart              string `json:"timeOfDeviceRestart"`
	RestartNotificationRecipients    string `json:"restartNotificationRecipients"`
	UTCTimeSynchronizationRecipients string `json:"UTCTimeSynchronizationRecipients"`
	TimeSynchronizationInterval      string `json:"timeSynchronizationInterval"`
	AlignIntervals                   string `json:"alignIntervals"`
	IntervalOffset                   string `json:"intervalOffset"`
	ProfileName                      string `json:"profileName"`
}

type File struct {
	Default
	FileType         string `json:"fileType"`
	FileSize         string `json:"fileSize"`
	ModificationDate string `json:"modificationDate"`
	Archive          string `json:"archive"`
	ReadOnly         string `json:"readOnly"`
	FileAccessMethod string `json:"fileAccessMethod"`
	Description      string `json:"description"`
	RecordCount      string `json:"recordCount"`
	ProfileName      string `json:"profileName"`
}

type IntegerValue struct {
	Default
	PresentValue                 string `json:"presentValue"`
	StatusFlags                  string `json:"statusFlags"`
	Units                        string `json:"units"`
	Description                  string `json:"description"`
	EventState                   string `json:"eventState"`
	Reliability                  string `json:"reliability"`
	OutOfService                 string `json:"outOfService"`
	PriorityArray                string `json:"priorityArray"`
	RelinquishDefault            string `json:"relinquishDefault"`
	CovIncrement                 string `json:"covIncrement"`
	TimeDelay                    string `json:"timeDelay"`
	NotificationClass            string `json:"notificationClass"`
	HighLimit                    string `json:"highLimit"`
	LowLimit                     string `json:"lowLimit"`
	Deadband                     string `json:"deadband"`
	LimitEnable                  string `json:"limitEnable"`
	EventEnable                  string `json:"eventEnable"`
	AckedTransitions             string `json:"ackedTransitions"`
	NotifyType                   string `json:"notifyType"`
	EventTimeStamps              string `json:"eventTimeStamps"`
	EventMessageTexts            string `json:"eventMessageTexts"`
	EventMessageTextsConfig      string `json:"eventMessageTextsConfig"`
	EventDetectionEnable         string `json:"eventDetectionEnable"`
	EventAlgorithmInhibitRef     string `json:"eventAlgorithmInhibitRef"`
	EventAlgorithmInhibit        string `json:"eventAlgorithmInhibit"`
	TimeDelayNormal              string `json:"timeDelayNormal"`
	ReliabilityEvaluationInhibit string `json:"reliabilityEvaluationInhibit"`
	MinPresValue                 string `json:"minPresValue"`
	MaxPresValue                 string `json:"maxPresValue"`
	Resolution                   string `json:"resolution"`
	ProfileName                  string `json:"profileName"`
	All                          string `json:"all"`
	Required                     string `json:"required"`
	Optional                     string `json:"optional"`
}

type LifeSafetyPoint struct {
	Default
	PresentValue                 string `json:"presentValue"`
	TrackingValue                string `json:"trackingValue"`
	StatusFlags                  string `json:"statusFlags"`
	EventState                   string `json:"eventState"`
	OutOfService                 string `json:"outOfService"`
	Reliability                  string `json:"reliability"`
	Mode                         string `json:"mode"`
	AcceptedModes                string `json:"acceptedModes"`
	Silenced                     string `json:"silenced"`
	OperationExpected            string `json:"operationExpected"`
	Description                  string `json:"description"`
	DeviceType                   string `json:"deviceType"`
	NotificationClass            string `json:"notificationClass"`
	LifeSafetyAlarmValues        string `json:"lifeSafetyAlarmValues"`
	AlarmValues                  string `json:"alarmValues"`
	FaultValues                  string `json:"faultValues"`
	EventEnable                  string `json:"eventEnable"`
	AckedTransitions             string `json:"ackedTransitions"`
	NotifyType                   string `json:"notifyType"`
	EventTimeStamps              string `json:"eventTimeStamps"`
	EventMessageTexts            string `json:"eventMessageTexts"`
	MaintenanceRequired          string `json:"maintenanceRequired"`
	Setting                      string `json:"setting"`
	DirectReading                string `json:"directReading"`
	Units                        string `json:"units"`
	MemberOf                     string `json:"memberOf"`
	ReliabilityEvaluationInhibit string `json:"reliabilityEvaluationInhibit"`
	ProfileName                  string `json:"profileName"`
}

type LightingOutput struct {
	Default
	PresentValue                   string `json:"presentValue"`
	TrackingValue                  string `json:"trackingValue"`
	LightingCommand                string `json:"lightingCommand"`
	InProgress                     string `json:"inProgress"`
	StatusFlags                    string `json:"statusFlags"`
	OutOfService                   string `json:"outOfService"`
	BlinkWarnEnable                string `json:"blinkWarnEnable"`
	EgressTime                     string `json:"egressTime"`
	EgressActive                   string `json:"egressActive"`
	DefaultFadeTime                string `json:"defaultFadeTime"`
	DefaultRampRate                string `json:"defaultRampRate"`
	DefaultStepIncrement           string `json:"defaultStepIncrement"`
	PriorityArray                  string `json:"priorityArray"`
	RelinquishDefault              string `json:"relinquishDefault"`
	LightingCommandDefaultPriority string `json:"lightingCommandDefaultPriority"`
	Description                    string `json:"description"`
	Reliability                    string `json:"reliability"`
	Transition                     string `json:"transition"`
	FeedbackValue                  string `json:"feedbackValue"`
	Power                          string `json:"power"`
	InstantaneousPower             string `json:"instantaneousPower"`
	MinActualValue                 string `json:"minActualValue"`
	MaxActualValue                 string `json:"maxActualValue"`
	CovIncrement                   string `json:"covIncrement"`
	ReliabilityEvaluationInhibit   string `json:"reliabilityEvaluationInhibit"`
	ProfileName                    string `json:"profileName"`
}

type LoadControl struct {
	Default
	PresentValue                 string `json:"presentValue"`
	StatusFlags                  string `json:"statusFlags"`
	EventState                   string `json:"eventState"`
	RequestedShedLevel           string `json:"requestedShedLevel"`
	StartTime                    string `json:"startTime"`
	ShedDuration                 string `json:"shedDuration"`
	DutyWindow                   string `json:"dutyWindow"`
	Enable                       string `json:"enable"`
	ExpectedShedLevel            string `json:"expectedShedLevel"`
	ActualShedLevel              string `json:"actualShedLevel"`
	ShedLevels                   string `json:"shedLevels"`
	ShedLevelDescriptions        string `json:"shedLevelDescriptions"`
	Description                  string `json:"description"`
	StateDescription             string `json:"stateDescription"`
	Reliability                  string `json:"reliability"`
	FullDutyBaseline             string `json:"fullDutyBaseline"`
	NotificationClass            string `json:"notificationClass"`
	TimeDelay                    string `json:"timeDelay"`
	EventEnable                  string `json:"eventEnable"`
	AckedTransitions             string `json:"ackedTransitions"`
	NotifyType                   string `json:"notifyType"`
	EventTimeStamps              string `json:"eventTimeStamps"`
	EventMessageTexts            string `json:"eventMessageTexts"`
	ReliabilityEvaluationInhibit string `json:"reliabilityEvaluationInhibit"`
	ProfileName                  string `json:"profileName"`
}

type MultistateInput struct {
	Default
	PresentValue                 string `json:"presentValue"`
	StatusFlags                  string `json:"statusFlags"`
	EventState                   string `json:"eventState"`
	OutOfService                 string `json:"outOfService"`
	NumberOfStates               string `json:"numberOfStates"`
	Description                  string `json:"description"`
	DeviceType                   string `json:"deviceType"`
	Reliability                  string `json:"reliability"`
	StateText                    string `json:"stateText"`
	TimeDelay                    string `json:"timeDelay"`
	NotificationClass            string `json:"notificationClass"`
	AlarmValues                  string `json:"alarmValues"`
	FaultValues                  string `json:"faultValues"`
	EventEnable                  string `json:"eventEnable"`
	AckedTransitions             string `json:"ackedTransitions"`
	NotifyType                   string `json:"notifyType"`
	EventTimeStamps              string `json:"eventTimeStamps"`
	EventMessageTexts            string `json:"eventMessageTexts"`
	ReliabilityEvaluationInhibit string `json:"reliabilityEvaluationInhibit"`
	ProfileName                  string `json:"profileName"`
}

type MultistateOutput struct {
	Default
	PresentValue                 string `json:"presentValue"`
	StatusFlags                  string `json:"statusFlags"`
	EventState                   string `json:"eventState"`
	OutOfService                 string `json:"outOfService"`
	NumberOfStates               string `json:"numberOfStates"`
	PriorityArray                string `json:"priorityArray"`
	RelinquishDefault            string `json:"relinquishDefault"`
	Description                  string `json:"description"`
	DeviceType                   string `json:"deviceType"`
	Reliability                  string `json:"reliability"`
	StateText                    string `json:"stateText"`
	TimeDelay                    string `json:"timeDelay"`
	NotificationClass            string `json:"notificationClass"`
	FeedbackValue                string `json:"feedbackValue"`
	EventEnable                  string `json:"eventEnable"`
	AckedTransitions             string `json:"ackedTransitions"`
	NotifyType                   string `json:"notifyType"`
	EventTimeStamps              string `json:"eventTimeStamps"`
	EventMessageTexts            string `json:"eventMessageTexts"`
	ReliabilityEvaluationInhibit string `json:"reliabilityEvaluationInhibit"`
	ProfileName                  string `json:"profileName"`
}

type MultistateValue struct {
	Default
	PresentValue                 string `json:"presentValue"`
	StatusFlags                  string `json:"statusFlags"`
	EventState                   string `json:"eventState"`
	OutOfService                 string `json:"outOfService"`
	NumberOfStates               string `json:"numberOfStates"`
	Description                  string `json:"description"`
	Reliability                  string `json:"reliability"`
	StateText                    string `json:"stateText"`
	PriorityArray                string `json:"priorityArray"`
	RelinquishDefault            string `json:"relinquishDefault"`
	TimeDelay                    string `json:"timeDelay"`
	NotificationClass            string `json:"notificationClass"`
	AlarmValues                  string `json:"alarmValues"`
	FaultValues                  string `json:"faultValues"`
	EventEnable                  string `json:"eventEnable"`
	AckedTransitions             string `json:"ackedTransitions"`
	NotifyType                   string `json:"notifyType"`
	EventTimeStamps              string `json:"eventTimeStamps"`
	EventMessageTexts            string `json:"eventMessageTexts"`
	ReliabilityEvaluationInhibit string `json:"reliabilityEvaluationInhibit"`
	ProfileName                  string `json:"profileName"`
}

type NotificationClass struct {
	Default
	NotificationClass string `json:"notificationClass"`
	Priority          string `json:"priority"`
	AckRequired       string `json:"ackRequired"`
	RecipientList     string `json:"recipientList"`
	Description       string `json:"description"`
	ProfileName       string `json:"profileName"`
}

type TrendLog struct {
	Default
	Enable                       string `json:"enable"`
	StopWhenFull                 string `json:"stopWhenFull"`
	BufferSize                   string `json:"bufferSize"`
	LogBuffer                    string `json:"logBuffer"`
	RecordCount                  string `json:"recordCount"`
	TotalRecordCount             string `json:"totalRecordCount"`
	EventState                   string `json:"eventState"`
	LoggingType                  string `json:"loggingType"`
	StatusFlags                  string `json:"statusFlags"`
	Description                  string `json:"description"`
	StartTime                    string `json:"startTime"`
	StopTime                     string `json:"stopTime"`
	LogDeviceObjectProperty      string `json:"logDeviceObjectProperty"`
	LogInterval                  string `json:"logInterval"`
	CovResubscriptionInterval    string `json:"covResubscriptionInterval"`
	ClientCovIncrement           string `json:"clientCovIncrement"`
	NotificationThreshold        string `json:"notificationThreshold"`
	RecordsSinceNotification     string `json:"recordsSinceNotification"`
	LastNotifyRecord             string `json:"lastNotifyRecord"`
	NotificationClass            string `json:"notificationClass"`
	EventEnable                  string `json:"eventEnable"`
	AckedTransitions             string `json:"ackedTransitions"`
	NotifyType                   string `json:"notifyType"`
	EventTimeStamps              string `json:"eventTimeStamps"`
	EventMessageTexts            string `json:"eventMessageTexts"`
	AlignIntervals               string `json:"alignIntervals"`
	IntervalOffset               string `json:"intervalOffset"`
	Trigger                      string `json:"trigger"`
	Reliability                  string `json:"reliability"`
	ReliabilityEvaluationInhibit string `json:"reliabilityEvaluationInhibit"`
	ProfileName                  string `json:"profileName"`
}