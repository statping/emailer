# Statping Email Notifier
This repo contains a dedicated service for Statping that allows users to receive notifications to their email address using our SMTP servers.

# Required Environment Variables
- `HOST` - SMTP host
- `USERNAME` - SMTP username
- `PASSWORD` - SMTP password
- `PORT` - SMTP port

> Server runs on port :8080

# API Endpoints

#### POST `/request`
This endpoint will send an email to the user with a confirmation link/key. Email's won't be sent until this confirmation link is clicked.
```json
{
  "email": "info@myaddress.com",
  "version": "0.90.33"
}
```

#### GET `/confirm/{id}`
Confirm an email address with the key provided in the request email.a

#### GET `/check?email=info@myaddress.com`
Check the status of your email address.

#### GET `/resend?email=info@myaddress.com`
Attempt to resend the confirmation request email.

#### GET `/unsubscribe?email=info@myaddress.com`
Unsubscribe the email address, and remove them from database.

#### POST `/send`
```json
{
  "email": "info@myaddress.com",
  "key": "<secret_key_here>",
  "version": "0.90.33",
  "service": {
    ...service JSON
  },
  "failure": {
      ...failure JSON
  }
}
```
The main endpoint that will send an email about uptime or downtime.



