# Delta

Yes, another account management service. But this one is small...hopefully. The only purpose of this application is for IT student being able to change their password in LDAP by logging in with Gamma.

![demo](https://raw.githubusercontent.com/molleer/delta/main/demo_image.jpg)

## Development

Start local version of gamma, LDAP mock and redis with the following command

```
    docker-compose -f gamma.docker-compose.yml up -d
    docker-compose up -d
```

Start the backend with the following commands

```
    cd backend
    sh ./dev_startup.sh
```

If you change anything in the backend, the backend needs to be manually restarted

Start the frontend with the following commands

```
  cd frontend
  npm start
```

Any changes in the frontend will update the frontend automatically

## Testing

Testing has to be sone manually. Do the following:

1. See that you can log in as `svanni` on LDAP
    - Go to [phpLdapAdmin](http://localhost:9000) and login with the following credentials
    ```
      Login DN: uid=svanni,ou=people,dc=chalmers,dc=it
      Password: password
    ```
    - Log out
2. Change the password of `svanni`
    - Go to http://localhost:3001, follow the instructions and change the password to whatever you like. Use the following credentials to gamma
    ```
      cid: svanni
      password: password
    ```
3. Login to [phpLdapAdmin](http://localhost:9000) with your new password but with the same `Login DN`

If your are able to do this, the application works as it should.
