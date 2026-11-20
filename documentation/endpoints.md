# endpoint

* post /products — Создать продукт 

1) Request: 
{

"title": "Молоко",

"price": 200.0

}

2) Response: 

{

"id": 1,

"title": "Молоко",

"price": 200.0

}

3) Status code: `201 Created`

* get /products — Получить все продукты

1) Request: None

2) Response: 
[

{

"id": 1,

"title": "Молоко",

"price": 200.0

},

{

"id": 2,

"title": "Хлеб",

"price": 50.0

}

]

3) Status code: `200 OK`

* get /products/{id} — Получить продукт по ID

1) Request: None

2) Response: 
{

"id": 1,

"title": "Молоко",

"price": 200.0

}
3) Status code: `200 OK`

* patch /products/{id} — Изменить продукт 

1) Request:

{

"title": "Молоко2",

"price": 22.0

}

2) Response:
{

"id": 1,

"title": "Молоко2",

"price": 22.0

}

3) Status code: `200 OK`

* delete /products/{id} — Удалить продукт

1) Request: None

2) Response: None

3) Status code: `204 No content`