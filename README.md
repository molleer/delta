# Delta

Yes, another account management service. But this one is small...hopefully. The only purpose of this application is for IT student being able to change their password in LDAP by logging in with Gamma.

## Development

Copy the `example.config.toml` to `config.toml`

Build dev environment

```
    docker-compose up
```

When all containers are running, you can access phpLdapAdmin at http://localhost:9000

Admin access LDAP:

- Login DN: `cn=admin,dc=chalmers,dc=it`
- Password: `password`

### Gamma

Run Gamma locally with the following command

```
    docker-compose -f gamma.docker-compose.yml up
```

Gamma can now be accessed via the browser at http://localhost:3000

Admin access Gamma:

- Username: `admin`
- Password: `password`

### LDAP Production

You may access the production LDAP server of the Chalmers IT student division [here](https://kamino.chalmers.it/phpldapadmin/).

User access:

- Login DN: `uid=[your cid],ou=people,dc=chalmers,dc=it`
- Password: Your password

Admin access:

- Login DN: `cn=admin,dc=chalmers,dc=it`
- Password: [found here](https://youtu.be/dQw4w9WgXcQ)

If the phpLdapAdmin site is down, you may run a local version with the following command

```
    docker-compose -f cthit.docker-compose.yml up
```

Your local phpLdapAdmin is now accessible in the browser at http://localhost:8080

## Existing problems

- A group may not be both an `prosixGroup` and `groupOfNames` `objectClass` at the same time since a group may only have ONE `structural` objectClass
  - The solution would be to change the `prosixGroup` schema to a `auxiliary` objectClass in the `mock/ldif/cthit.ldif` file, but this has not been done
- Cannot add groups under `ou=SUDOers`
  - The `sudoRole` objectclass has not been constructed (commendted out in `chalmersstudent`)
- `cn=digit,dc=chalmers,dc=it` is not the admin account and logging in as such will not give you any access
  - You may try it with the password: `password`
