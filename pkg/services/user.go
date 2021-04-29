package services

import (
	"fmt"
	"strconv"

	"gopkg.in/ldap.v2"
)

type Post struct {
	ID          string `json:"id"`
	Sv          string `json:"sv"`
	En          string `json:"en"`
	EmailPrefix string `json:"emailPrefix"`
}

type ITUser struct {
	ID                    string      `json:"id"`
	Cid                   string      `json:"cid"`
	Nick                  string      `json:"nick"`
	FirstName             string      `json:"firstName"`
	LastName              string      `json:"lastName"`
	Email                 string      `json:"email"`
	Phone                 string      `json:"phone"`
	Language              string      `json:"language"`
	AvatarURL             string      `json:"avatarUrl"`
	Gdpr                  bool        `json:"gdpr"`
	UserAgreement         bool        `json:"userAgreement"`
	AccountLocked         bool        `json:"accountLocked"`
	AcceptanceYear        int         `json:"acceptanceYear"`
	Authorities           []string    `json:"authorities"`
	Activated             bool        `json:"activated"`
	Enabled               bool        `json:"enabled"`
	Username              string      `json:"username"`
	AccountNonLocked      bool        `json:"accountNonLocked"`
	AccountNonExpired     bool        `json:"accountNonExpired"`
	CredentialsNonExpired bool        `json:"credentialsNonExpired"`
	Groups                []FKITGroup `json:"groups"`
	WebsiteURLs           string      `json:"websiteURLs"`
}

type FKITUser struct {
	Post                  Post         `json:"post"`
	FkitGroupDTO          FKITGroupDTO `json:"fkitGroupDTO"`
	UnofficialPostName    string       `json:"unofficialPostName"`
	ID                    string       `json:"id"`
	Cid                   string       `json:"cid"`
	Nick                  string       `json:"nick"`
	FirstName             string       `json:"firstName"`
	LastName              string       `json:"lastName"`
	Email                 string       `json:"email"`
	Phone                 string       `json:"phone"`
	Language              string       `json:"language"`
	AvatarURL             string       `json:"avatarUrl"`
	Gdpr                  bool         `json:"gdpr"`
	UserAgreement         bool         `json:"userAgreement"`
	AccountLocked         bool         `json:"accountLocked"`
	AcceptanceYear        int          `json:"acceptanceYear"`
	Authorities           []string     `json:"authorities"`
	Activated             bool         `json:"activated"`
	Username              string       `json:"username"`
	AccountNonExpired     bool         `json:"accountNonExpired"`
	AccountNonLocked      bool         `json:"accountNonLocked"`
	CredentialsNonExpired bool         `json:"credentialsNonExpired"`
	Enabled               bool         `json:"enabled"`
}

func (user *ITUser) ToLdapEntry(uidNumber int, genPassword bool) []ldap.Attribute {
	gdpr := ""
	if user.Gdpr {
		gdpr = "TRUE"
	} else {
		gdpr = "FALSE"
	}

	attributes := []ldap.Attribute{
		{Type: "accepteduseragreement", Vals: []string{gdpr}},
		{Type: "admissionyear", Vals: []string{strconv.FormatInt(int64(user.AcceptanceYear), 10)}},
		{Type: "cn", Vals: []string{"%{firstname} '%{nickname}' %{lastname}"}},
		{Type: "gidnumber", Vals: []string{"4500"}},
		{Type: "givenname", Vals: []string{user.FirstName}},
		{Type: "homedirectory", Vals: []string{fmt.Sprintf("/home/chalmersit/%s", user.Cid)}},
		{Type: "loginshell", Vals: []string{"/bin/bash"}},
		{Type: "mail", Vals: []string{user.Email}},
		{Type: "nickname", Vals: []string{user.Nick}},
		{Type: "objectclass", Vals: []string{"chalmersstudent", "posixAccount", "top"}},
		{Type: "sn", Vals: []string{user.LastName}},
		{Type: "telephonenumber", Vals: []string{user.Phone}},
		{Type: "uid", Vals: []string{user.Cid}},
		{Type: "uidnumber", Vals: []string{fmt.Sprintf("%v", uidNumber)}},
	}

	if genPassword {
		attributes = append(attributes, ldap.Attribute{
			Type: "userpassword",
			Vals: []string{fmt.Sprintf("{SSHA}%s", RandomString(30))},
		})
	}

	return attributes
}

func NewUser(entry *ldap.Entry) ITUser {
	userAgreement, _ := strconv.ParseBool(entry.GetAttributeValue("acceptedUserAgreement"))
	admissionYear, _ := strconv.Atoi(entry.GetAttributeValue("admissionYear"))
	return ITUser{
		Cid:                   entry.GetAttributeValue("uid"),
		FirstName:             entry.GetAttributeValue("givenName"),
		LastName:              entry.GetAttributeValue("sn"),
		UserAgreement:         userAgreement,
		AcceptanceYear:        admissionYear,
		Nick:                  entry.GetAttributeValue("nickname"),
		Email:                 entry.GetAttributeValue("mail"),
		Phone:                 entry.GetAttributeValue("telephoneNumber"),
		Language:              entry.GetAttributeValue("preferredLanguage"),
		AccountLocked:         false,
		Activated:             true,
		Enabled:               true,
		AccountNonLocked:      true,
		CredentialsNonExpired: true,
	}
}

// CRUD User =========================================

func (s *ServiceLDAP) GetITUsers() ([]ITUser, error) {
	request := ldap.NewSearchRequest(
		s.UsersConfig.BaseDN,
		ldap.ScopeWholeSubtree, ldap.NeverDerefAliases, 0, 0, false,
		s.UsersConfig.Filter,
		[]string{"uid", "givenName", "sn", "acceptedUserAgreement", "admissionYear",
			"nickname", "mail", "telephoneNumber", "preferredLanguage"},
		nil,
	)

	users, err := s.Connection.Search(request)
	if err != nil {
		return nil, err
	}

	itUsers := []ITUser{}
	for _, entry := range users.Entries {

		itUsers = append(itUsers, NewUser(entry))
	}

	return itUsers, nil
}

func (s *ServiceLDAP) AddITUser(user ITUser, uidNumber int) error {
	return s.Connection.Add(&ldap.AddRequest{
		DN:         fmt.Sprintf("uid=%s,%s", user.Cid, s.UsersConfig.BaseDN),
		Attributes: user.ToLdapEntry(uidNumber, true),
	})
}

func (s *ServiceLDAP) DeleteUser(cid string) error {
	return s.Connection.Del(
		ldap.NewDelRequest(fmt.Sprintf("uid=%s,%s", cid, s.UsersConfig.BaseDN),
			nil))
}

func (s *ServiceLDAP) getUser(cid string) (*ldap.Entry, error) {
	request := ldap.NewSearchRequest(
		s.UsersConfig.BaseDN,
		ldap.ScopeWholeSubtree, ldap.NeverDerefAliases, 0, 0, false,
		fmt.Sprintf("(uid=%s)", cid),
		[]string{"uid", "givenName", "sn", "acceptedUserAgreement", "admissionYear",
			"nickname", "mail", "telephoneNumber", "preferredLanguage"},
		nil,
	)

	user, err := s.Connection.Search(request)
	if err != nil {
		return nil, err
	}

	return user.Entries[0], nil
}

/**
* Returns the user information, but not group membership
 */
func (s *ServiceLDAP) GetITUser(cid string) (ITUser, error) {
	user, err := s.getUser(cid)
	if err != nil {
		return ITUser{}, err
	}

	return NewUser(user), nil
}

func (s *ServiceLDAP) UpdateUser(user ITUser) error {
	ldapUser, err := s.getUser(user.Cid)
	if err != nil {
		return err
	}
	uidNumber, _ := strconv.Atoi(ldapUser.GetAttributeValue("uidNumber"))

	return s.Connection.Modify(&ldap.ModifyRequest{
		DN:                fmt.Sprintf("uid=%s,%s", user.Cid, s.UsersConfig.BaseDN),
		ReplaceAttributes: ToPartial(user.ToLdapEntry(uidNumber, false)),
	})
}

func (s *ServiceLDAP) SetPassword(cid string, password string) error {
	/*UserIdentity string
	// OldPassword, if present, contains the user's current password
	OldPassword string
	// NewPassword, if present, contains the desired password for this user
	NewPassword string*/
	_, err := s.Connection.PasswordModify(&ldap.PasswordModifyRequest{
		UserIdentity: fmt.Sprintf("uid=%s,%s", cid, s.UsersConfig.BaseDN),
		NewPassword:  password,
	})
	return err
}
