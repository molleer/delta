package services

import (
	"crypto/tls"
	"fmt"
	"os"

	"gopkg.in/ldap.v2"
)

var (
	LDAP_URL = os.Getenv("LDAP_URL")
	LDAP_SERVER_NAME = os.Getenv("LDAP_SERVER_NAME")
	LDAP_TLS = os.Getenv("LDAP_TLS") == "true"
	LDAP_USERS_DN = os.Getenv("LDAP_USERS_DN")
	admin_dn = os.Getenv("LDAP_ADMIN_DN")
	admin_pass = os.Getenv("LDAP_ADMIN_PASS")
)

type ServiceLDAP struct {
	Connection *ldap.Conn
}

func NewLDAPService() (*ServiceLDAP, error) {
	l, err := ldap.DialTLS("tcp",
		LDAP_URL,
		&tls.Config{
			ServerName: LDAP_SERVER_NAME,
			InsecureSkipVerify: !LDAP_TLS,
		})

	return &ServiceLDAP{
		Connection:   l,
	}, err
}

func (s *ServiceLDAP) AdminLogin() error {
	return s.LoginUser(admin_dn, admin_pass)
}

func (s *ServiceLDAP) LoginUser(userDN string, password string) error {
	return s.Connection.Bind(userDN, password)
}

func (s *ServiceLDAP) UserExist(cid string) bool {
	request := ldap.NewSearchRequest(
		LDAP_USERS_DN,
		ldap.ScopeWholeSubtree, ldap.NeverDerefAliases, 0, 0, false,
		fmt.Sprintf("(uid=%s)", cid),
		[]string{"uid"},
		nil,
	)

	user, err := s.Connection.Search(request)
	if err != nil || len(user.Entries) < 1 {
		return false
	}

	return user.Entries[0].GetAttributeValue("uid") == cid
}

func (s *ServiceLDAP) SetPassword(cid string, password string) error {
	_, err := s.Connection.PasswordModify(&ldap.PasswordModifyRequest{
		UserIdentity: fmt.Sprintf("uid=%s,%s", cid, LDAP_USERS_DN),
		NewPassword:  password,
	})
	return err
}