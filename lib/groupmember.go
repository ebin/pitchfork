package pitchfork

type PfGroupMember interface {
	Set(groupname, groupdesc, username, fullname, affiliation string, groupadmin bool, groupstate string, cansee bool, email, pgpkey_id, entered, activity, telephone, sms, airport string)
	GetGroupName() string
	GetGroupDesc() string
	GetUserName() string
	GetFullName() string
	GetEmail() string
	GetAffiliation() string
	GetGroupAdmin() bool
	GetGroupState() string
	GetGroupCanSee() bool
	GetPGPKeyID() string
	HasPGP() bool
	GetEntered() string
	GetActivity() string
	GetTel() string
	GetSMS() string
	GetAirport() string
}

type PfGroupMemberS struct {
	UserName    string
	FullName    string
	Email       string
	Affiliation string
	GroupAdmin  bool
	GroupState  string
	GroupCanSee bool
	GroupName   string
	GroupDesc   string
	PGPKeyID    string
	Entered     string
	Activity    string
	Tel         string
	SMS         string
	Airport     string
}

func NewPfGroupMember() PfGroupMember {
	return &PfGroupMemberS{}
}

func (grpm *PfGroupMemberS) Set(groupname, groupdesc, username, fullname, affiliation string, groupadmin bool, groupstate string, cansee bool, email, pgpkey_id, entered, activity, telephone, sms, airport string) {
	grpm.GroupName = groupname
	grpm.GroupDesc = groupdesc
	grpm.UserName = username
	grpm.FullName = fullname
	grpm.Affiliation = affiliation
	grpm.GroupAdmin = groupadmin
	grpm.GroupState = groupstate
	grpm.GroupCanSee = cansee
	grpm.Email = email
	grpm.PGPKeyID = pgpkey_id
	grpm.Entered = entered
	grpm.Activity = activity
	grpm.Tel = telephone
	grpm.SMS = sms
	grpm.Airport = airport
}

func (grpm *PfGroupMemberS) GetGroupName() string {
	return grpm.GroupName
}

func (grpm *PfGroupMemberS) GetGroupDesc() string {
	return grpm.GroupDesc
}

func (grpm *PfGroupMemberS) GetUserName() string {
	return grpm.UserName
}

func (grpm *PfGroupMemberS) GetFullName() string {
	return grpm.FullName
}

func (grpm *PfGroupMemberS) GetEmail() string {
	return grpm.Email
}

func (grpm *PfGroupMemberS) GetAffiliation() string {
	return grpm.Affiliation
}

func (grpm *PfGroupMemberS) GetGroupAdmin() bool {
	return grpm.GroupAdmin
}

func (grpm *PfGroupMemberS) GetGroupState() string {
	return grpm.GroupState
}

func (grpm *PfGroupMemberS) GetGroupCanSee() bool {
	return grpm.GroupCanSee
}

func (grpm *PfGroupMemberS) GetPGPKeyID() string {
	return grpm.PGPKeyID
}

func (grpm *PfGroupMemberS) HasPGP() bool {
	return grpm.PGPKeyID != ""
}

func (grpm *PfGroupMemberS) GetEntered() string {
	return grpm.Entered
}

func (grpm *PfGroupMemberS) GetActivity() string {
	return grpm.Activity
}

func (grpm *PfGroupMemberS) GetTel() string {
	return grpm.Tel
}

func (grpm *PfGroupMemberS) GetSMS() string {
	return grpm.SMS
}

func (grpm *PfGroupMemberS) GetAirport() string {
	return grpm.Airport
}
