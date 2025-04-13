package models

type AmoLeadPipeline int

const (
	AmoLeadPipelineIncomingRequests        AmoLeadPipeline = 7389450
	AmoLeadPipelineColdBaseCustomers       AmoLeadPipeline = 7395738
	AmoLeadPipelineColdBaseExecutors       AmoLeadPipeline = 7395742
	AmoLeadPipelineTenders                 AmoLeadPipeline = 7405518
	AmoLeadPipelineCustomers               AmoLeadPipeline = 7440246
	AmoLeadPipelineExecutors               AmoLeadPipeline = 7440250
	AmoLeadPipelineColdBaseTransportations AmoLeadPipeline = 7548018
)

type AmoLeadPipelineStatus int

const (
	AmoLeadIncomingRequestsPipelineStatusUnsorted                AmoLeadPipelineStatus = 61432862
	AmoLeadIncomingRequestsPipelineStatusNewLead                 AmoLeadPipelineStatus = 61432866
	AmoLeadIncomingRequestsPipelineStatusInWork                  AmoLeadPipelineStatus = 61432870
	AmoLeadIncomingRequestsPipelineStatusNoAnswer                AmoLeadPipelineStatus = 61432874
	AmoLeadIncomingRequestsPipelineStatusQualified               AmoLeadPipelineStatus = 61432878
	AmoLeadIncomingRequestsPipelineStatusInterestConfirmed       AmoLeadPipelineStatus = 61475670
	AmoLeadIncomingRequestsPipelineStatusDirectedToRegistration  AmoLeadPipelineStatus = 61475674
	AmoLeadIncomingRequestsPipelineStatusRegistrationCompleted   AmoLeadPipelineStatus = 61475678
	AmoLeadIncomingRequestsPipelineStatusVerificationCompleted   AmoLeadPipelineStatus = 61475682
	AmoLeadIncomingRequestsPipelineStatusSuccessfullyImplemented AmoLeadPipelineStatus = 142
	AmoLeadIncomingRequestsPipelineStatusClosedAndNotImplemented AmoLeadPipelineStatus = 143
)
