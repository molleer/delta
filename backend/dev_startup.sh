export COOKIE_DOMAIN=localhost
export SESSION_SECRET=secret
export REDIS_HOST=localhost
export REDIS_PASS=

export GAMMA_URL=http://localhost:8081
export GAMMA_REDIRECT_URL=http://localhost:8081
export GAMMA_CLIENT_ID=id
export GAMMA_CLIENT_SECRET=secret
export GAMMA_CALLBACK_URL=http://localhost:3001/callback

export LDAP_URL=localhost:636
export LDAP_SERVER_NAME=ldap.chalmers.it
export LDAP_TLS=false
export LDAP_ADMIN_DN=cn=admin,dc=chalmers,dc=it
export LDAP_ADMIN_PASS=password
export LDAP_USERS_DN=ou=people,dc=chalmers,dc=it

export GIN_MODE=debug
export PORT=8090

go run ./