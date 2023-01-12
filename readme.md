# Api Airline FPM

---

### api path

---
- auth 
```
[POST] /auth/login 
BODY 
client_id
client_secret
```
---

- common || master data

```
HEADER 
Authorization Bearer ...

[GET] /nationalities
[GET] /airports/original
[GET] /airports/original/:id
[GET] /airports/destination
[GET] /airports/destination/:id
```
---
- upload fpm 

```
HEADER 
Authorization Bearer ...

[POST] /fpm/create
BODY
{
    "flight_no" : "sdfdf",
    "origin_code" : "123213",
    "destination_code" : "AAA",
    "estimate_departure_date_time" : "2022-01-02 08:40",
    "estimate_arrival_date_time" : "2022-01-02 08:40",
    "passengers" : [
        {
            "first_name" : "AAA",
            "middle_name" : "",
            "last_name" : "AAA",
            "passport_no" : "123123123",
            "dateof_birth" : "2022-01-02",
            "gender" : "T",
            "nationality" : "THA",
            "phone_no" : "123123",
            "email" : "mail@mail.com",
            "seat_no" : "A23",
            "tax_code" : "12312"
        },
        ...
    ]
}
```