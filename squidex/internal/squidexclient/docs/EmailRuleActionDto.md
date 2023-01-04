# EmailRuleActionDto

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**ServerHost** | **string** | The IP address or host to the SMTP server. | 
**ServerPort** | **int32** | The port to the SMTP server. | 
**ServerUsername** | **string** | The username for the SMTP server. | 
**ServerPassword** | **string** | The password for the SMTP server. | 
**MessageFrom** | **string** | The email sending address. | 
**MessageTo** | **string** | The email message will be sent to. | 
**MessageSubject** | **string** | The subject line for this email message. | 
**MessageBody** | **string** | The message body. | 

## Methods

### NewEmailRuleActionDto

`func NewEmailRuleActionDto(serverHost string, serverPort int32, serverUsername string, serverPassword string, messageFrom string, messageTo string, messageSubject string, messageBody string, ) *EmailRuleActionDto`

NewEmailRuleActionDto instantiates a new EmailRuleActionDto object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewEmailRuleActionDtoWithDefaults

`func NewEmailRuleActionDtoWithDefaults() *EmailRuleActionDto`

NewEmailRuleActionDtoWithDefaults instantiates a new EmailRuleActionDto object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetServerHost

`func (o *EmailRuleActionDto) GetServerHost() string`

GetServerHost returns the ServerHost field if non-nil, zero value otherwise.

### GetServerHostOk

`func (o *EmailRuleActionDto) GetServerHostOk() (*string, bool)`

GetServerHostOk returns a tuple with the ServerHost field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetServerHost

`func (o *EmailRuleActionDto) SetServerHost(v string)`

SetServerHost sets ServerHost field to given value.


### GetServerPort

`func (o *EmailRuleActionDto) GetServerPort() int32`

GetServerPort returns the ServerPort field if non-nil, zero value otherwise.

### GetServerPortOk

`func (o *EmailRuleActionDto) GetServerPortOk() (*int32, bool)`

GetServerPortOk returns a tuple with the ServerPort field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetServerPort

`func (o *EmailRuleActionDto) SetServerPort(v int32)`

SetServerPort sets ServerPort field to given value.


### GetServerUsername

`func (o *EmailRuleActionDto) GetServerUsername() string`

GetServerUsername returns the ServerUsername field if non-nil, zero value otherwise.

### GetServerUsernameOk

`func (o *EmailRuleActionDto) GetServerUsernameOk() (*string, bool)`

GetServerUsernameOk returns a tuple with the ServerUsername field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetServerUsername

`func (o *EmailRuleActionDto) SetServerUsername(v string)`

SetServerUsername sets ServerUsername field to given value.


### GetServerPassword

`func (o *EmailRuleActionDto) GetServerPassword() string`

GetServerPassword returns the ServerPassword field if non-nil, zero value otherwise.

### GetServerPasswordOk

`func (o *EmailRuleActionDto) GetServerPasswordOk() (*string, bool)`

GetServerPasswordOk returns a tuple with the ServerPassword field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetServerPassword

`func (o *EmailRuleActionDto) SetServerPassword(v string)`

SetServerPassword sets ServerPassword field to given value.


### GetMessageFrom

`func (o *EmailRuleActionDto) GetMessageFrom() string`

GetMessageFrom returns the MessageFrom field if non-nil, zero value otherwise.

### GetMessageFromOk

`func (o *EmailRuleActionDto) GetMessageFromOk() (*string, bool)`

GetMessageFromOk returns a tuple with the MessageFrom field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMessageFrom

`func (o *EmailRuleActionDto) SetMessageFrom(v string)`

SetMessageFrom sets MessageFrom field to given value.


### GetMessageTo

`func (o *EmailRuleActionDto) GetMessageTo() string`

GetMessageTo returns the MessageTo field if non-nil, zero value otherwise.

### GetMessageToOk

`func (o *EmailRuleActionDto) GetMessageToOk() (*string, bool)`

GetMessageToOk returns a tuple with the MessageTo field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMessageTo

`func (o *EmailRuleActionDto) SetMessageTo(v string)`

SetMessageTo sets MessageTo field to given value.


### GetMessageSubject

`func (o *EmailRuleActionDto) GetMessageSubject() string`

GetMessageSubject returns the MessageSubject field if non-nil, zero value otherwise.

### GetMessageSubjectOk

`func (o *EmailRuleActionDto) GetMessageSubjectOk() (*string, bool)`

GetMessageSubjectOk returns a tuple with the MessageSubject field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMessageSubject

`func (o *EmailRuleActionDto) SetMessageSubject(v string)`

SetMessageSubject sets MessageSubject field to given value.


### GetMessageBody

`func (o *EmailRuleActionDto) GetMessageBody() string`

GetMessageBody returns the MessageBody field if non-nil, zero value otherwise.

### GetMessageBodyOk

`func (o *EmailRuleActionDto) GetMessageBodyOk() (*string, bool)`

GetMessageBodyOk returns a tuple with the MessageBody field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMessageBody

`func (o *EmailRuleActionDto) SetMessageBody(v string)`

SetMessageBody sets MessageBody field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


