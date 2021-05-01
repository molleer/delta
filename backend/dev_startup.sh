export COOKIE_HOST=localhost
export COOKIE_SECRET=secret

export GAMMA_URL=http://localhost:8081
export GAMMA_REDIRECT_URL=http://localhost:3001
export GAMMA_CLIENT_ID=id
export GAMMA_CLIENT_KEY=key

export LDAP_URL=localhost:636
export LDAP_SERVER_NAME=ldap.chalmers.it
export LDAP_TLS=false
export LDAP_ADMIN_DN=cn=admin,dc=chalmers,dc=it
export LDAP_ADMIN_PASS=password
export LDAP_USERS_DN=ou=people,dc=chalmers,dc=it

export GIN_MODE=debug
export PORT=8081

~/go/bin/gin -p 8080 -a 8081 -i run main.go