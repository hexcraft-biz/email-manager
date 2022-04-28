# Email-Manager
Current database schema version: 1.0

## email API
#### POST /email
>	- Description: Create a new email agent.
>	- Request Payload Body:
>	```json
>	{
>		"addr": "service-01@newmedia.pts.org.tw",
>		"password": "xxooxxoo"
>	}
>	```
>	- Success Response: 201
>	```json
>	{
>		"message": "Created",
>		"results": {
>			"id": "d353b4a6-73ac-45f0-a82a-58e94656a3d8"
>		}
>	}
>	```

#### PATCH /email/count `CRON: Daily`
>	- Description: Reset daily count. The system limit is 1950 per day. Reserving 50 per email for urgent case. See https://support.google.com/a/answer/166852
>	- Note: **This should be triggered ONCE EVERY DAY by cron.**
>	- Request Payload Body:
>	- Success Response: 200
>	```json
>	{
>		"message": "OK"
>	}
>	```

#### POST /content
>	- Description: Send email with given content
>	- Request Payload Body:
>	```json
>	{
>		"to": ["user1@example.com", "user2@example.com"],
>		"subject": "Hear this!",
>		"body": "My ggininder!"
>	}
>	```
>	- Success Response: 201
>	```json
>	{
>		"message": "Created"
>	}
>	```
